package core

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/ipfsync/ipfsync/core/api"

	"github.com/ipfsync/resource"

	"github.com/ipfsync/appserver"
	"github.com/ipfsync/ipfsmanager"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func NewIpfsManager(lc fx.Lifecycle, cfg *viper.Viper) (*ipfsmanager.IpfsManager, error) {
	im, err := ipfsmanager.NewIpfsManager(cfg.GetString("repoDir"))
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

func NewAppServer(lc fx.Lifecycle, api *api.Api, cfg *viper.Viper) (*appserver.AppServer, error) {
	srv := appserver.NewAppServer(api, cfg)
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

func NewDataStore(lc fx.Lifecycle, cfg *viper.Viper) (*resource.Datastore, error) {
	ds, err := resource.NewDatastore(filepath.Join(cfg.GetString("dataDir"), "db"))
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return ds.Close()
		},
	})

	return ds, nil
}
