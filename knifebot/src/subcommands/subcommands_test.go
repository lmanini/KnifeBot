package subcommands

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestViewSybilClusterComm(t *testing.T) {
	addr := common.HexToAddress("0x9c9190635d46c36452da89c1603216e0377e0e57")
	got, err := ViewSybilClusterComm(&addr)
	if err != nil {
		panic(err)
	}
	for _, a := range got {
		fmt.Printf("%s\n", a.String())
	}
}
