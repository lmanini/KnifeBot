package subcommands

import (
	"KnifeBot/src/bindings"
	"KnifeBot/src/constants"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Command uint8

const (
	SpyBalance Command = iota
	SpyTopBalance
)

func SpyBalanceComm(target *common.Address) (uint64, error) {
	client, _ := ethclient.Dial(constants.RPC)
	caller, err := bindings.NewSpyNFTCaller(common.HexToAddress(constants.SpyNFT),client)
	bal, err := caller.BalanceOf(&bind.CallOpts{}, *target)
	if err != nil {
		return 0, err
	}
	return bal.Uint64(), nil
}