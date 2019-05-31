package main

import (
	"github.com/ipfsync/ipfsync/core"
	"github.com/ipfsync/ipfsync/core/api"

	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(api.NewApi, core.NewIpfsManager, core.NewConfig, core.NewDataStore),
		fx.Invoke(core.NewAppServer),
	)

	app.Run()

}
