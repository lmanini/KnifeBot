package main

import (
	"KnifeBot/src/subcommands"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "KnifeBot",
		Usage: "CLI bot to play KnifeGame",
		Commands: []*cli.Command{
			{
				Name: "run",
				Aliases: []string{"r"},
				Usage: "run bot",
				Description: "Does nothing still :P",
				Action: func(ctx *cli.Context) error {
					return run()
				},
			},
			{
				Name: "spy-balance",
				Aliases: []string{"sb"},
				Description: "Fetch Spy balance",
				Action: func(ctx *cli.Context) error {
					return exec(subcommands.SpyBalance, common.HexToAddress(ctx.Args().Get(0)))
				},
			},
			{
				Name: "knife-balance",
				Aliases: []string{"kb"},
				Description: "Fetch Knife balance",
				Action: func(ctx *cli.Context) error {
					return exec(subcommands.KnifeBalance, common.HexToAddress(ctx.Args().Get(0)))
				},
			}, {
				Name: "moo-balance",
				Aliases: []string{"mb"},
				Description: "Fetch Moo balance",
				Action: func(ctx *cli.Context) error {
					return exec(subcommands.MooBalance, common.HexToAddress(ctx.Args().Get(0)))
				},
			},
			{
				Name: "view-sybil-cluster",
				Description: "Recursively looks through Transfer logs to see who has received/sent spies from/to the specified address",
				Aliases: []string{"cluster", "c"},
				Action: func(ctx *cli.Context) error {
					return exec(subcommands.ViewSybilCluster, common.HexToAddress(ctx.Args().Get(0)))
				},
			}, 
			{
				Name: "spy-price",
				Aliases: []string{"sp"},
				Action: func(ctx *cli.Context) error {
					return exec(subcommands.SpyPrice)
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	log.Println("GM")
	return nil
}

func exec(command subcommands.Command, v ...interface{}) error {
	switch command {
	case subcommands.SpyBalance:
		addr := v[0].(common.Address)
		bal, err := subcommands.SpyBalanceComm(&addr)
		if err != nil {
			return err
		}
		log.Printf("Address %s has Spy balance : %d\n", addr.Hex(), bal)
	case subcommands.KnifeBalance:
		addr := v[0].(common.Address)
		bal, err := subcommands.KnifeBalanceComm(&addr)
		if err != nil {
			return err
		}
		log.Printf("Address %s has Knife balance : %d\n", addr.Hex(), bal)
	case subcommands.MooBalance:
		addr := v[0].(common.Address)
		bal, err := subcommands.MooBalanceComm(&addr)
		if err != nil {
			return err
		}
		fBal := new(big.Float)
		fBal.SetString(bal.String())
		log.Printf("Address %s has Moo balance : %s\n", addr.Hex(), new(big.Float).Quo(fBal, big.NewFloat(math.Pow10(18))).String())
	case subcommands.ViewSybilCluster:
		addr := v[0].(common.Address)
		cluster, err := subcommands.ViewSybilClusterComm(&addr)
		if err != nil {
			return err
		}
		log.Printf("Address %s was found to have received/sent spies from/to:\n", addr)
		for _, a := range cluster {
			if a.String() == addr.String() { continue }
			log.Printf("\t%s\n", a.String())
		}
	case subcommands.SpyPrice:
		price, err := subcommands.ViewSpyPriceComm()
		if err != nil {
			return err
		}
		fPrice := new(big.Float)
		fPrice.SetString(price.String())
		log.Printf("Spy price is currently at %s", new(big.Float).Quo(fPrice, big.NewFloat(math.Pow10(18))).String())
	}
	
	return nil
}