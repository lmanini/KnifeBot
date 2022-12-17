package subcommands

import (
	"KnifeBot/src/bindings"
	"KnifeBot/src/constants"
	"math/big"
	"sort"

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
	MooTopBalance
	SpyPrice
	KnifePrice
	ViewSybilCluster
	ViewAllBalances
)

/** COMMANDS */

func SpyBalanceComm(target *common.Address) (uint64, error) {
	return getSpyBal(target)
}

func KnifeBalanceComm(target *common.Address) (uint64, error) {
	return getKnifeBal(target)
}

func MooBalanceComm(target *common.Address) (*big.Int, error) {
	return getMooBal(target)
}

func SpyTopBalanceComm(limit int) ([]common.Address, []int, error) {
	return topSpyBalance(limit)
}

func KnifeTopBalanceComm(limit int) ([]common.Address, []int, error) {
	return topKnifeBalance(limit)
}

func MooTopBalanceComm(limit int) ([]common.Address, []*big.Int, error) {
	return topMooBalance(limit)
}

func ViewAllBalancesComm(source *common.Address) (uint64, uint64, *big.Int, error) {
	return getAllBalances(source)
}

func ViewSybilClusterComm(source *common.Address) ([]common.Address, error) {
	return getSybilCluster(source)
}

func ViewSpyPriceComm() (*big.Int, error) {
	return getSpyPrice()
}

func ViewKnifePriceComm() (*big.Int, error) {
	return getKnifePrice()
}

/** BODIES */

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

func getSpyPrice() (*big.Int, error) {
	client, _ := ethclient.Dial(constants.RPC)	
	caller, _ := bindings.NewKnifeGameCaller(common.HexToAddress(constants.KnifeGame), client)
	price, err := caller.SpyPrice(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return price, nil
}

func getKnifePrice() (*big.Int, error) {
	client, _ := ethclient.Dial(constants.RPC)	
	caller, _ := bindings.NewKnifeGameCaller(common.HexToAddress(constants.KnifeGame), client)
	price, err := caller.KnifePrice(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return price, nil
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
	sources := []common.Address{}
	for toItr.Next() {
		if toItr.Event.To.String() != constants.ZeroAddr && !cluster[toItr.Event.To] {
			cluster[toItr.Event.To] = true
			dirty = true
			sources = append(sources, toItr.Event.To)
		}
	}
	for fromItr.Next() {
		if fromItr.Event.From.String() != constants.ZeroAddr && !cluster[fromItr.Event.From] {
			cluster[fromItr.Event.From] = true
			dirty = true
			sources = append(sources, fromItr.Event.From)
		}
	}
	if dirty {
		for _, s := range sources {
			getSybilClusterBody(&s, spyFilterer, cluster)
		}
	}
	
	keys := make([]common.Address, 0, len(cluster))
	for a := range cluster {
		keys = append(keys, a)
	}
	return keys, nil
}

func topSpyBalance(limit int) ([]common.Address, []int, error) {
	client, _ := ethclient.Dial(constants.RPC)
	spyNFT, _ := bindings.NewSpyNFT(common.HexToAddress(constants.SpyNFT), client)
	it, err := spyNFT.SpyNFTFilterer.FilterTransfer(
		&bind.FilterOpts{
			Start: constants.SpyNFTGenesisBlock,
			End: nil,
		},
		[]common.Address{common.HexToAddress(constants.ZeroAddr)},
		nil,
		nil,
	)
	if err != nil {
		return nil, nil, err
	}
	addresses := []common.Address{}
	addrToBal := make(map[common.Address]uint64)
	for it.Next() {
		if _, ok := addrToBal[it.Event.To]; !ok {
			addresses = append(addresses, it.Event.To)
			spyBal, _ := getSpyBal(&(it.Event.To))
			addrToBal[it.Event.To] += spyBal
		}
	}
	sort.SliceStable(addresses, func(i, j int) bool {
		return addrToBal[addresses[i]] > addrToBal[addresses[j]]
	})

	topBal := make([]int, limit)
	for i, a := range addresses[:limit] {topBal[i] = int(addrToBal[a])}
	return addresses[:limit], topBal, nil
}

func topKnifeBalance(limit int) ([]common.Address, []int, error) {
	client, _ := ethclient.Dial(constants.RPC)
	knifeNFT, _ := bindings.NewKnifeNFT(common.HexToAddress(constants.KnifeNFT), client)
	it, err := knifeNFT.KnifeNFTFilterer.FilterTransfer(
		&bind.FilterOpts{
			Start: constants.KnifeNFTGenesisBlock,
			End: nil,
		},
		[]common.Address{common.HexToAddress(constants.ZeroAddr)},
		nil,
		nil,
	)
	if err != nil {
		return nil, nil, err
	}
	addresses := []common.Address{}
	addrToBal := make(map[common.Address]uint64)
	for it.Next() {
		if _, ok := addrToBal[it.Event.To]; !ok {
			addresses = append(addresses, it.Event.To)
			knifeBal, _ := getKnifeBal(&(it.Event.To))
			addrToBal[it.Event.To] += knifeBal
		}
	}
	sort.SliceStable(addresses, func(i, j int) bool {
		return addrToBal[addresses[i]] > addrToBal[addresses[j]]
	})

	topBal := make([]int, limit)
	for i, a := range addresses[:limit] {topBal[i] = int(addrToBal[a])}
	return addresses[:limit], topBal, nil
}

func topMooBalance(limit int) ([]common.Address, []*big.Int, error) {
	client, _ := ethclient.Dial(constants.RPC)
	spyNFT, _ := bindings.NewSpyNFT(common.HexToAddress(constants.SpyNFT), client)
	it, err := spyNFT.SpyNFTFilterer.FilterTransfer(
		&bind.FilterOpts{
			Start: constants.SpyNFTGenesisBlock,
			End: nil,
		},
		[]common.Address{common.HexToAddress(constants.ZeroAddr)},
		nil,
		nil,
	)
	if err != nil {
		return nil, nil, err
	}
	addresses := []common.Address{}
	addrToBal := make(map[common.Address]*big.Int)
	for it.Next() {
		if _, ok := addrToBal[it.Event.To]; !ok {
			addresses = append(addresses, it.Event.To)
			mooBal, _ := getMooBal(&(it.Event.To))
			if addrToBal[it.Event.To] == nil {
				addrToBal[it.Event.To] = big.NewInt(0)
			}
			addrToBal[it.Event.To].Add(addrToBal[it.Event.To], mooBal)
		}
	}
	sort.SliceStable(addresses, func(i, j int) bool {
		return addrToBal[addresses[i]].Cmp(addrToBal[addresses[j]]) > 0
	})

	topBal := make([]*big.Int, limit)
	for i, a := range addresses[:limit] {topBal[i] = addrToBal[a]}
	return addresses[:limit], topBal, nil
}

func getAllBalances(source *common.Address) (uint64, uint64, *big.Int, error) {
	spyBal, err := getSpyBal(source)
	knifeBal, err := getKnifeBal(source)
	mooBal, err := getMooBal(source)
	if err != nil {
		return 0, 0, nil, err
	}
	return spyBal, knifeBal, mooBal, nil
}