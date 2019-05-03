package core

import (
	"github.com/ipfsync/ipfsmanager"
)

type Api struct {
	mgr *ipfsmanager.IpfsManager
}

func NewApi(mgr *ipfsmanager.IpfsManager) *Api {
	return &Api{mgr: mgr}
}

func (api *Api) Peers() {

}
