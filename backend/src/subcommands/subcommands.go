package subcommands

import (
	"KnifeBot/src/bindings"
	"KnifeBot/src/constants"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Command uint8

const (
	SpyBalance Command = iota
	SpyTopBalance
	KnifeBalance
	KnifeTopBalance
	MooBalance
)

func SpyBalanceComm(target *common.Address) (uint64, error) {
	return getSpyBal(target)
}

func KnifeBalanceComm(target *common.Address) (uint64, error) {
	return getKnifeBal(target)
}

func MooBalanceComm(target *common.Address) (*big.Int, error) {
	return getMooBal(target)
}

// sybil cluster getters
func ViewSybilClusterComm(source *common.Address) ([]common.Address, error) {
	return getSybilCluster(source)
}

// total spies

// total knives

// total moo ? moo is non transferable tho
// --> calculate total moo but keep track 
// of each addr in the cluster's balance 

func getSpyBal(target *common.Address) (uint64, error) {
	client, _ := ethclient.Dial(constants.RPC)
	caller, err := bindings.NewSpyNFTCaller(common.HexToAddress(constants.SpyNFT),client)
	bal, err := caller.BalanceOf(&bind.CallOpts{}, *target)
	if err != nil {
		return 0, err
	}
	return bal.Uint64(), nil
}

func getKnifeBal(target *common.Address) (uint64, error) {
	client, _ := ethclient.Dial(constants.RPC)
	caller, err := bindings.NewKnifeNFTCaller(common.HexToAddress(constants.KnifeNFT),client)
	bal, err := caller.BalanceOf(&bind.CallOpts{}, *target)
	if err != nil {
		return 0, err
	}
	return bal.Uint64(), nil
}

func getMooBal(target *common.Address) (*big.Int, error) {
	client, _ := ethclient.Dial(constants.RPC)
	caller, err := bindings.NewSpyNFTCaller(common.HexToAddress(constants.SpyNFT),client)
	bal, err := caller.GooBalance(&bind.CallOpts{}, *target)
	if err != nil {
		return big.NewInt(0), err
	}
	return bal, nil
}

func getSybilCluster(source *common.Address) ([]common.Address, error) {
	client, _ := ethclient.Dial(constants.RPC)
	spyFilterer, err := bindings.NewSpyNFTFilterer(common.HexToAddress(constants.SpyNFT), client)
	if err != nil {
		return nil, err
	}
	return getSybilClusterBody(source, spyFilterer, make(map[common.Address]bool))
}

func getSybilClusterBody(source *common.Address, spyFilterer *bindings.SpyNFTFilterer, cluster map[common.Address]bool) ([]common.Address, error) {
	toItr, _ := spyFilterer.FilterTransfer(
		&bind.FilterOpts{
			Start: constants.SpyNFTGenesisBlock,
			End: nil,
		},
		[]common.Address{*source},
		nil,
		nil,
	)
	fromItr, _ := spyFilterer.FilterTransfer(
		&bind.FilterOpts{
			Start: constants.SpyNFTGenesisBlock,
			End: nil,
		},
		nil,
		[]common.Address{*source},
		nil,
	)
	dirty := false
	for toItr.Next() {
		if !cluster[toItr.Event.To] {
			cluster[toItr.Event.To] = true
			dirty = true
		}
	}
	for fromItr.Next() {
		if !cluster[fromItr.Event.From] {
			cluster[fromItr.Event.From] = true
			dirty = true
		}
	}
	if dirty {
		getSybilClusterBody(source, spyFilterer, cluster)
	}
	
	keys := make([]common.Address, 0, len(cluster))
	for a := range cluster {
		keys = append(keys, a)
	}
	return keys, nil
	
}