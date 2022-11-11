package gotest

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/pngfi/merkle-distributor/generated/merkle_distributor"
	"github.com/stretchr/testify/require"
)

func TestAdminWithdraw(t *testing.T) {
	t.Skip()

	wsolATA := ATA(solana.SolMint, TestFeePayer)
	basePrivk := solana.NewWallet().PrivateKey
	base := basePrivk.PublicKey()
	distributor, bump, _ := solana.FindProgramAddress([][]byte{
		[]byte("MerkleDistributor"), base[:],
	}, ProgramID)
	distributorATA := ATA(solana.SolMint, distributor)

	{ // create wsol ata
		instr1 := NewATAInstrs(1e9, solana.SolMint, TestFeePayer, TestFeePayer)
		instr2 := NewATAInstrs(1e9, solana.SolMint, distributor, TestFeePayer)

		instrs := append(instr1, instr2...)

		sig, err := SignAndSendTransaction(ctx, LocalTestClient, TestWallet, TestFeePayer, instrs...)
		require.NoError(t, err)
		require.NoError(t, PollingWaitTXConfirm(ctx, LocalTestClient, sig))
	}

	{ // new distributor
		instr := merkle_distributor.NewNewDistributorInstruction(
			bump,
			[32]uint8{},
			2,
			2,
			base,
			TestFeePayer,
			distributor,
			solana.SolMint,
			TestFeePayer,
			solana.SystemProgramID,
		).Build()
		sig, err := SignAndSendTransaction(ctx, LocalTestClient, AppendToWallet(TestWallet, basePrivk), TestFeePayer, instr)
		require.NoError(t, err)

		require.NoError(t, PollingWaitTXConfirm(ctx, LocalTestClient, sig))

		state, err := distributorLoader.Account(ctx, distributor)
		require.NoError(t, err)
		fmt.Println("load mint:", state.State.Mint)
	}

	{ // admin withdraw + assert
		instr := merkle_distributor.NewAdminWithdrawInstruction(1e8, distributor, TestFeePayer, distributorATA, wsolATA, solana.TokenProgramID).Build()
		sig, err := SignAndSendTransaction(ctx, LocalTestClient, AppendToWallet(TestWallet, basePrivk), TestFeePayer, instr)
		require.NoError(t, err)

		require.NoError(t, PollingWaitTXConfirm(ctx, LocalTestClient, sig))

		tx, err := LocalTestClient.GetTransaction(ctx, *sig, &rpc.GetTransactionOpts{})
		require.NoError(t, err)
		fmt.Println("tx logs:", tx.Meta.LogMessages)

		bal, err := LocalTestClient.GetBalance(ctx, wsolATA, rpc.CommitmentConfirmed)
		require.NoError(t, err)
		fmt.Println("bal::", bal)
	}
}

func TestUpdateConfig(t *testing.T) {
	// t.Skip()

	wsolATA := ATA(solana.SolMint, TestFeePayer)
	basePrivk := solana.NewWallet().PrivateKey
	base := basePrivk.PublicKey()
	distributor, bump, _ := solana.FindProgramAddress([][]byte{
		[]byte("MerkleDistributor"), base[:],
	}, ProgramID)
	distributorATA := ATA(solana.SolMint, distributor)

	_, _ = wsolATA, distributorATA

	{ // create wsol ata
		instr1 := NewATAInstrs(1e9, solana.SolMint, TestFeePayer, TestFeePayer)
		instr2 := NewATAInstrs(1e9, solana.SolMint, distributor, TestFeePayer)

		instrs := append(instr1, instr2...)

		sig, err := SignAndSendTransaction(ctx, LocalTestClient, TestWallet, TestFeePayer, instrs...)
		require.NoError(t, err)
		require.NoError(t, PollingWaitTXConfirm(ctx, LocalTestClient, sig))
	}

	{ // new distributor
		instr := merkle_distributor.NewNewDistributorInstruction(
			bump,
			[32]uint8{},
			2,
			2,
			base,
			TestFeePayer,
			distributor,
			solana.SolMint,
			TestFeePayer,
			solana.SystemProgramID,
		).Build()
		sig, err := SignAndSendTransaction(ctx, LocalTestClient, AppendToWallet(TestWallet, basePrivk), TestFeePayer, instr)
		require.NoError(t, err)

		require.NoError(t, PollingWaitTXConfirm(ctx, LocalTestClient, sig))

		state, err := distributorLoader.Account(ctx, distributor)
		require.NoError(t, err)
		fmt.Println("load mint:", state.State.Mint)
	}

	config, _, _ := solana.FindProgramAddress([][]byte{
		[]byte("distributor_config"), distributor[:],
	}, ProgramID)

	claimStatus, claimBump, _ := solana.FindProgramAddress([][]byte{
		[]byte("ClaimStatus"),
		distributor[:],
		TestFeePayer[:],
	}, ProgramID)
	{ //claim should return InvalidProofError
		claimInstr := merkle_distributor.NewClaimInstruction(
			claimBump,
			0,
			1,
			[][32]uint8{},
			distributor,
			config,
			claimStatus,
			distributorATA,
			wsolATA,
			TestFeePayer,
			TestFeePayer,
			solana.SystemProgramID,
			solana.TokenProgramID,
		).Build()

		ret, err := SimulateTransaction(ctx, LocalTestClient, TestWallet, TestFeePayer, claimInstr)
		require.NoError(t, err)
		// fmt.Println("should claim failed with err invalid proof", ret.Value.Logs) //0x1770
		require.Contains(t, strings.Join(ret.Value.Logs, ""), "0x1770")
	}

	{
		deadline := time.Now().Unix() - 10
		instr := merkle_distributor.NewUpdateConfigInstruction(deadline, TestFeePayer, distributor, config, solana.SystemProgramID).Build()
		sig, err := SignAndSendTransaction(ctx, LocalTestClient, AppendToWallet(TestWallet, basePrivk), TestFeePayer, instr)
		require.NoError(t, err)
		require.NoError(t, PollingWaitTXConfirm(ctx, LocalTestClient, sig))

		state, err := configLoader.Account(ctx, config)
		require.NoError(t, err)

		require.Equal(t, *state.State.ClaimDeadline, deadline)
		require.Equal(t, state.State.Distributor, distributor)
	}

	{ //claim should return InvalidProofError
		claimInstr := merkle_distributor.NewClaimInstruction(
			claimBump,
			0,
			1,
			[][32]uint8{},
			distributor,
			config,
			claimStatus,
			distributorATA,
			wsolATA,
			TestFeePayer,
			TestFeePayer,
			solana.SystemProgramID,
			solana.TokenProgramID,
		).Build()

		ret, err := SimulateTransaction(ctx, LocalTestClient, TestWallet, TestFeePayer, claimInstr)
		require.NoError(t, err)
		// fmt.Println("should claim failed with err ExceededDeadline", ret.Value.Logs)

		require.Contains(t, strings.Join(ret.Value.Logs, ""), "0x177a")
	}
}
