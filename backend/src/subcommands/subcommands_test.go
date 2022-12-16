package subcommands

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestViewSybilClusterComm(t *testing.T) {
	addr := common.HexToAddress("0x90F79bf6EB2c4f870365E785982E1f101E93b906")
	got, err := ViewSybilClusterComm(&addr)
	if err != nil {
		panic(err)
	}
	for _, a := range got {
		fmt.Printf("%s\n", a.String())
	}
}
