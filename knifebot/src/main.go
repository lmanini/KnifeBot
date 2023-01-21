package main

import (
	"KnifeBot/src/subcommands"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
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
			{
				Name: "knife-price",
				Aliases: []string{"kp"},
				Action: func(ctx *cli.Context) error {
					return exec(subcommands.KnifePrice)
				},
			}, {
				Name: "spy-balance-top",
				Aliases: []string{"sbt"},
				Action: func(ctx *cli.Context) error {
					limit, err := strconv.Atoi(ctx.Args().Get(0))
					if err != nil {
						return err
					}
					return exec(subcommands.SpyTopBalance, limit)
				},
			}, {
				Name: "knife-balance-top",
				Aliases: []string{"kbt"},
				Action: func(ctx *cli.Context) error {
					limit, err := strconv.Atoi(ctx.Args().Get(0))
					if err != nil {
						return err
					}
					return exec(subcommands.KnifeTopBalance, limit)
				},
			},  {
				Name: "moo-balance-top",
				Aliases: []string{"mbt"},
				Action: func(ctx *cli.Context) error {
					limit, err := strconv.Atoi(ctx.Args().Get(0))
					if err != nil {
						return err
					}
					return exec(subcommands.MooTopBalance, limit)
				},
			}, {
				Name: "view-all-balances",
				Aliases: []string{"all", "vab"},
				Action: func(ctx *cli.Context) error {
					return exec(subcommands.ViewAllBalances, common.HexToAddress(ctx.Args().Get(0)))
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// init
	err := godotenv.Load(".env")
	if err != nil { return err }
	pks := getPrivateKeys()
	addrs := getAddresses(pks)

	//main body
	for {
		playKnifeGame(pks, addrs)
	}

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
		log.Printf("Address %s has Moo balance : %s\n", addr.Hex(), getWithDecimals(bal, 18).String())
	case subcommands.SpyTopBalance:
		top, topBal, err := subcommands.SpyTopBalanceComm(v[0].(int))
		if err != nil {
			return err
		}
		log.Printf("Spy leaderboard : \n")
		for i, a := range top {
			log.Printf("\t%d. : %s has %d spies", i + 1, a.Hex(), topBal[i])
		}
	case subcommands.KnifeTopBalance:
		top, topBal, err := subcommands.KnifeTopBalanceComm(v[0].(int))
		if err != nil {
			return err
		}
		log.Printf("Knife leaderboard : \n")
		for i, a := range top {
			log.Printf("\t%d. : %s has %d knives", i + 1, a.Hex(), topBal[i])
		}
	case subcommands.MooTopBalance:
		top,topBal, err := subcommands.MooTopBalanceComm(v[0].(int))
		if err != nil {
			return err
		}
		log.Printf("Moo leaderboard : \n")
		for i, a := range top {
			log.Printf("\t%d. : %s has %s moo", i + 1, a.Hex(), getWithDecimals(topBal[i], 18).String())
		}
	
	case subcommands.SpyPrice:
		price, err := subcommands.ViewSpyPriceComm()
		if err != nil {
			return err
		}
		fPrice := new(big.Float)
		fPrice.SetString(price.String())
		log.Printf("Spy price is currently at %s", new(big.Float).Quo(fPrice, big.NewFloat(math.Pow10(18))).String())
	case subcommands.KnifePrice:
		price, err := subcommands.ViewKnifePriceComm()
		if err != nil {
			return err
		}
		fPrice := new(big.Float)
		fPrice.SetString(price.String())
		log.Printf("Knife price is currently at %s", new(big.Float).Quo(fPrice, big.NewFloat(math.Pow10(18))).String())

	case subcommands.ViewSybilCluster:
		addr := v[0].(common.Address)
		cluster, err := subcommands.ViewSybilClusterComm(&addr)
		if err != nil {
			return err
		}
		log.Printf("Address %s was found to have dealt spies & knives with:\n", addr)
		for _, a := range cluster {
			if a.String() == addr.String() { continue }
			log.Printf("\t%s\n", a.String())
		}
	case subcommands.ViewAllBalances:
		addr := v[0].(common.Address)
		spyBal, knifeBal, mooBal, err := subcommands.ViewAllBalancesComm(&addr)
		if err != nil {
			return err
		}
		log.Printf("Address %s has : \n\t%d spies \n\t%d knives \n\t%s moo", addr.Hex(), spyBal, knifeBal, getWithDecimals(mooBal, 18).String())
	}
	return nil
}

func getWithDecimals(num *big.Int, decimals int) *big.Float {
	fBal := new(big.Float)
	fBal.SetString(num.String())
	return new(big.Float).Quo(fBal, big.NewFloat(math.Pow10(decimals)))
} 

func getPrivateKeys() []string {
	pk1, _ := os.LookupEnv("PK1")
	pk2, _ := os.LookupEnv("PK2")
	pk3, _ := os.LookupEnv("PK3")
	pk4, _ := os.LookupEnv("PK4")

	return []string{pk1, pk2, pk3, pk4}
}

func getAddresses(pks []string) []common.Address{
	addresses := make([]common.Address, 0, len(pks))
	for _, pk := range pks {
		ecdsa, _ := crypto.HexToECDSA(pk)
		addr := crypto.PubkeyToAddress(ecdsa.PublicKey)
		addresses = append(addresses, addr)
	}
	return addresses
}

func playKnifeGame(pks []string, addrs []common.Address) {

	//initial log : commanding these addrs using private keys pks

	//main loop
}