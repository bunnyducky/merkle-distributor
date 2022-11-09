package gotest

import (
	"context"

	ag_binary "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type UnmarshalWithBorshDecoder[B any] interface {
	UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error)
	*B // non-interface type constraint element, https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#pointer-method-example
}

// anchor account loader (with 8 byte account discriminator)
type AccountLoader[T any, PT UnmarshalWithBorshDecoder[T]] struct {
	programID     solana.PublicKey
	discriminator [8]byte
	rpc           *rpc.Client
}

func (l *AccountLoader[T, PT]) SetProgramID(id solana.PublicKey) {
	l.programID = id
}

func (l *AccountLoader[T, PT]) GetProgramID() solana.PublicKey {
	return l.programID
}

func NewAccountLoader[T any, PT UnmarshalWithBorshDecoder[T]](programID solana.PublicKey, discriminator [8]byte, client *rpc.Client) *AccountLoader[T, PT] {
	return &AccountLoader[T, PT]{
		programID:     programID,
		discriminator: discriminator,
		rpc:           client,
	}
}

type AccountState[T any] struct {
	Address solana.PublicKey
	State   T
}

func (l *AccountLoader[T, PT]) DecodeAccount(data []byte) (*T, error) {
	t := new(T)
	pt := PT(t)
	err := pt.UnmarshalWithDecoder(ag_binary.NewBorshDecoder(data))
	return t, err
}

func (l *AccountLoader[T, PT]) Account(ctx context.Context, account solana.PublicKey) (AccountState[*T], error) {
	ret, err := l.rpc.GetAccountInfo(ctx, account)
	if err != nil {
		return AccountState[*T]{}, err
	}

	acc := ret.Value

	state, err := l.DecodeAccount(acc.Data.GetBinary())
	if err != nil {
		return AccountState[*T]{}, err
	}

	return AccountState[*T]{
		Address: account,
		State:   state,
	}, nil

}

// func (l *AccountLoader[T, PT]) Accounts(ctx context.Context, filters ...[]byte) ([]AccountState[*T], error) {
// 	filterBytes := l.discriminator[:]
// 	for i := 0; i < len(filters); i++ {
// 		filterBytes = append(filterBytes, filters[i]...)
// 	}

// 	accounts, err := l.rpc.GetProgramAccountsWithOpts(ctx, l.programID, &rpc.GetProgramAccountsOpts{
// 		Filters: []rpc.RPCFilter{
// 			{Memcmp: &rpc.RPCFilterMemcmp{
// 				Offset: 0,
// 				Bytes:  solana.Base58(filterBytes),
// 			}},
// 		},
// 	})
// 	if err != nil {
// 		return nil, errors.Wrap(err, "load accounts")
// 	}

// 	return slice.TryMap(accounts, func(t *rpc.KeyedAccount) (AccountState[*T], error) {
// 		state, err := l.DecodeAccount(t.Account.Data.GetBinary())
// 		if err != nil {
// 			return AccountState[*T]{}, err
// 		}

// 		return AccountState[*T]{
// 			Address: t.Pubkey,
// 			State:   state,
// 		}, nil
// 	})
// }
