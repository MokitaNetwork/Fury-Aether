package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	servercmd "github.com/cosmos/cosmos-sdk/server/cmd"

	aether "github.com/mokitanetwork/aether/app"
)

func main() {
	aether.SetAccountAddressPrefixes()

	root, _ := NewRootCmd()
	if err := servercmd.Execute(root, aether.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)
		default:
			os.Exit(1)
		}
	}
}
