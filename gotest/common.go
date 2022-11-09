package gotest

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/pngfi/merkle-distributor/generated/merkle_distributor"
	"github.com/pngfi/merkle-distributor/gotest/testdata"
)

var (
	LocalTestClient = rpc.New("http://localhost:8899")
	TestWallet      = PrivateKeys(testdata.Wallet)
	TestFeePayer    = testdata.TestKey

	ctx = context.Background()

	ProgramID = solana.MustPublicKeyFromBase58("PMRKTWvK9f1cPkQuXvvyDPmyCSoq8FdedCimXrXJp8M")

	distributorLoader = NewAccountLoader[merkle_distributor.MerkleDistributor](
		ProgramID, merkle_distributor.MerkleDistributorDiscriminator, LocalTestClient,
	)
	configLoader = NewAccountLoader[merkle_distributor.Config](
		ProgramID, merkle_distributor.ConfigDiscriminator, LocalTestClient,
	)
)
