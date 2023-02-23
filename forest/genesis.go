package forest

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/greenchainearth/seed-base/hash"
	"github.com/greenchainearth/seed-base/inter/idx"

	"github.com/greenchainearth/go-forest/forest/genesis"
	"github.com/greenchainearth/go-forest/forest/genesis/gpos"
	"github.com/greenchainearth/go-forest/inter"
)

type Genesis struct {
	Accounts    genesis.Accounts
	Storage     genesis.Storage
	Delegations genesis.Delegations
	Blocks      genesis.Blocks
	RawEvmItems genesis.RawEvmItems
	Validators  gpos.Validators

	FirstEpoch    idx.Epoch
	PrevEpochTime inter.Timestamp
	Time          inter.Timestamp
	ExtraData     []byte

	TotalSupply *big.Int

	DriverOwner common.Address

	Rules Rules

	Hash func() hash.Hash
}
