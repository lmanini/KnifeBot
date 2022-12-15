package main

import (
	"KnifeBot/src/subcommands"
	"log"
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
				Action: func(ctx *cli.Context) error {
					return run()
				},
			},
			{
				Name: "spy-balance",
				Aliases: []string{"sb"},
				Action: func(ctx *cli.Context) error {
					exec(subcommands.SpyBalance, common.HexToAddress(ctx.Args().Get(0)))
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	
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
		log.Printf("Address %s has Spy balance : %d", addr.Hex(), bal)
	}
	return nil
}