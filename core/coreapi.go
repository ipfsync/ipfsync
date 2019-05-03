package core

import (
	"context"

	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfsync/ipfsmanager"
)

type Api struct {
	mgr *ipfsmanager.IpfsManager
}

func NewApi(mgr *ipfsmanager.IpfsManager) *Api {
	return &Api{mgr: mgr}
}

// Peers returns peers that IPFS is currently connecting to
func (api *Api) Peers() ([]iface.ConnectionInfo, error) {
	return api.mgr.API.Swarm().Peers(context.TODO())
}
