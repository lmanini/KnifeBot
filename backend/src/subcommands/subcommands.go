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