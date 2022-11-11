package gotest

import (
	"context"
	"time"

	"github.com/gagliardetto/solana-go"
	associatedtokenaccount "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/pkg/errors"
)

type Wallet interface {
	GetPrivateKey(solana.PublicKey) *solana.PrivateKey
}

type PrivateKeys map[solana.PublicKey]solana.PrivateKey

func (keys PrivateKeys) GetPrivateKey(pk solana.PublicKey) *solana.PrivateKey {
	k := keys[pk]
	return &k
}

func AppendToWallet(w Wallet, key solana.PrivateKey) Wallet {
	if m, ok := w.(PrivateKeys); ok {
		m[key.PublicKey()] = key
		return w
	}
	panic("not PrivateKeys, append to wallet failed")
}

func SignAndSendTransaction(ctx context.Context, client *rpc.Client, wallet Wallet, feePayer solana.PublicKey, instrs ...solana.Instruction) (*solana.Signature, error) {
	hash, err := client.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return nil, err
	}
	builder := solana.NewTransactionBuilder().
		SetFeePayer(feePayer).
		SetRecentBlockHash(hash.Value.Blockhash)

	for i := 0; i < len(instrs); i++ {
		builder.AddInstruction(instrs[i])
	}

	tx, err := builder.Build()
	if err != nil {
		return nil, err
	}

	_, err = tx.Sign(wallet.GetPrivateKey)
	if err != nil {
		return nil, err
	}

	// for _, s := range tx.Signatures {
	// 	fmt.Println("debug: tx signature ", s)
	// }

	sig, err := client.SendTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}
	return &sig, nil
}

func ATA(mint, owner solana.PublicKey) solana.PublicKey {
	ata, _, _ := solana.FindAssociatedTokenAddress(owner, mint)
	return ata
}

func NewATAInstrs(fund uint64, mint, owner, feePayer solana.PublicKey) []solana.Instruction {
	ata, _, _ := solana.FindAssociatedTokenAddress(owner, mint)

	instrs := make([]solana.Instruction, 0)
	if solana.SolMint.Equals(mint) {
		instrs = append(instrs, system.NewTransferInstruction(fund, feePayer, ata).Build())
	}
	instrs = append(instrs, associatedtokenaccount.NewCreateInstruction(feePayer, owner, mint).Build())
	return instrs
}

func PollingWaitTXConfirm(ctx context.Context, client *rpc.Client, sig *solana.Signature) error {
	const timeout = time.Minute
	const sleepInterval = 1 * time.Second

	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for {
		result, err := client.GetTransaction(ctx, *sig, &rpc.GetTransactionOpts{
			Commitment: rpc.CommitmentFinalized,
		})
		if err == nil {
			if result.Meta.Err == nil {
				return nil
			} else {
				return errors.New("Transaction failed")
			}
		}
		select {
		case <-timeoutCtx.Done():
			return errors.Errorf("tx(%s) not confirmed during %s, or context done", sig, timeout)
		default:
		}
		time.Sleep(sleepInterval)
	}
}

func SimulateTransaction(ctx context.Context, client *rpc.Client, wallet Wallet, feePayer solana.PublicKey, instrs ...solana.Instruction) (*rpc.SimulateTransactionResponse, error) {
	hash, err := client.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return nil, err
	}
	builder := solana.NewTransactionBuilder().
		SetFeePayer(feePayer).
		SetRecentBlockHash(hash.Value.Blockhash)

	for i := 0; i < len(instrs); i++ {
		builder.AddInstruction(instrs[i])
	}

	tx, err := builder.Build()
	if err != nil {
		return nil, err
	}

	_, err = tx.Sign(wallet.GetPrivateKey)
	if err != nil {
		return nil, err
	}

	return client.SimulateTransactionWithOpts(ctx, tx, &rpc.SimulateTransactionOpts{
		SigVerify: true,
	})
}
