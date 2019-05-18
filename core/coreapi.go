package core

import (
	"context"
	"sort"
	"time"

	net "github.com/libp2p/go-libp2p-net"

	"github.com/ipfsync/ipfsmanager"
)

type Api struct {
	mgr *ipfsmanager.IpfsManager
}

func NewApi(mgr *ipfsmanager.IpfsManager) *Api {
	return &Api{mgr: mgr}
}

type Peerinfo struct {
	Address   string
	Direction net.Direction
	Latency   time.Duration
}

var oldpeersinfo []Peerinfo

// Peers returns peers that IPFS is currently connecting to
func (api *Api) Peers() ([]Peerinfo, bool, error) {
	peers, err := api.mgr.API.Swarm().Peers(context.TODO())
	if err != nil {
		return nil, false, err
	}

	var peersinfo []Peerinfo
	for _, p := range peers {
		l, _ := p.Latency()
		peersinfo = append(peersinfo, Peerinfo{
			Address:   p.Address().String(),
			Direction: p.Direction(),
			Latency:   l,
		})
	}

	// Sort
	sort.Slice(peersinfo, func(i, j int) bool {
		return peersinfo[i].Address < peersinfo[j].Address
	})

	changed := false
	if len(oldpeersinfo) != len(peersinfo) {
		changed = true
	} else {
		for i, p := range peersinfo {
			if oldpeersinfo[i].Address != p.Address {
				changed = true
				break
			}
		}
	}

	oldpeersinfo = peersinfo

	return peersinfo, changed, nil
}
