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
					subcommand(subcommands.SpyBalance, common.HexToAddress(ctx.Args().Get(0)))
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

func subcommand(command subcommands.Command, addr common.Address) {
	switch command {
	case subcommands.SpyBalance:
		log.Printf("Address %s has Spy balance : %d", addr.Hex(), subcommands.SpyBalanceComm(&addr))
	}
}