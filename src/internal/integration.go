package internal

import (
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/p2p/simulations/adapters"

	"github.com/Fantom-foundation/go-lachesis/src/gossip"
	"github.com/Fantom-foundation/go-lachesis/src/lachesis"
	"github.com/Fantom-foundation/go-lachesis/src/poset"
)

func NewIntegration(cfg *adapters.NodeConfig, config *lachesis.Net) *gossip.Service {
	makeDb := dbProducer(cfg.DataDir)
	gdb, cdb := makeStorages(makeDb)

	err := cdb.ApplyGenesis(config.Genesis)
	if err != nil {
		panic(err)
	}

	c := posposet.New(cdb, gdb)

	svc, err := gossip.NewService(config, new(event.TypeMux), gdb, c)
	if err != nil {
		panic(err)
	}

	return svc
}