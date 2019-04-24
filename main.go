package main

import (
	"context"
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/ipfsync/appserver"

	"github.com/ipfsync/ipfsmanager"
	"go.uber.org/fx"
)

func NewIpfsManager(lc fx.Lifecycle) (*ipfsmanager.IpfsManager, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	im, err := ipfsmanager.NewIpfsManager(filepath.Join(usr.HomeDir, "ipfshare"))
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			err := im.StartNode()
			keys, err := im.API.Key().List(context.TODO())
			if err != nil {
				return err
			}

			for _, key := range keys {
				fmt.Printf("Key ID: %s", key.ID())
			}

			return err
		},
		OnStop: func(ctx context.Context) error {
			return im.StopNode()
		},
	})
	return im, nil
}

func NewAppServer(lc fx.Lifecycle) (*appserver.AppServer, error) {
	srv := appserver.NewAppServer()
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			srv.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Stop()
		},
	})

	return srv, nil
}

func main() {

	app := fx.New(
		fx.Invoke(NewIpfsManager),
		fx.Invoke(NewAppServer),
	)

	app.Run()

}
