package testdata

import "github.com/gagliardetto/solana-go"

var (
	TestKey = solana.MustPublicKeyFromBase58("4wVKadFD3y35mMBCHmhwceHheWSxaTdnWqZ24qwEouxN")

	Wallet = map[solana.PublicKey]solana.PrivateKey{
		TestKey: solana.PrivateKey{
			173, 94, 6, 124, 251, 220, 160, 27, 139, 138, 44, 149, 133, 175, 186, 89, 184,
			222, 242, 17, 142, 0, 184, 155, 184, 50, 124, 160, 223, 123, 2, 121, 58, 136,
			114, 69, 246, 14, 3, 178, 53, 192, 225, 125, 103, 128, 94, 242, 161, 94, 109,
			81, 1, 129, 57, 16, 152, 210, 60, 140, 64, 61, 105, 27,
		},
	}
)
