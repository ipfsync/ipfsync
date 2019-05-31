package main

import (
	"github.com/ipfsync/ipfsync/core"

	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(core.NewApi, core.NewIpfsManager, core.NewConfig),
		fx.Invoke(core.NewAppServer),
	)

	app.Run()

}
