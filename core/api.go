package core

import (
	"context"
	"sort"
	"time"

	"github.com/ipfsync/resource"

	"github.com/spf13/viper"

	net "github.com/libp2p/go-libp2p-net"

	"github.com/ipfsync/ipfsmanager"
)

type Api struct {
	mgr *ipfsmanager.IpfsManager
	cfg *viper.Viper
	ds  *resource.Datastore
}

func NewApi(mgr *ipfsmanager.IpfsManager, cfg *viper.Viper, ds *resource.Datastore) *Api {
	return &Api{mgr: mgr, cfg: cfg, ds: ds}
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

func (api *Api) NewCollection(name, address string) (*resource.Collection, error) {
	if address == "" {
		keyName := "ipfsync_ipnskey"
		ctx := context.TODO()

		// Remove possible existed key
		_, _ = api.mgr.API.Key().Remove(ctx, keyName)

		// Generate new key
		k, err := api.mgr.API.Key().Generate(context.TODO(), keyName)
		if err != nil {
			return nil, err
		}

		address = k.ID().Pretty()

		// Rename new key to ID string
		_, _, err = api.mgr.API.Key().Rename(ctx, keyName, address)
		if err != nil {
			return nil, err
		}
	}

	if name == "" {
		name = address
	}

	// Insert data into datastore
	c := &resource.Collection{Name: name, IPNSAddress: address}
	err := api.ds.CreateOrUpdateCollection(c)
	if err != nil {
		return c, err
	}

	return c, nil
}
