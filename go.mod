module github.com/ipfsync/ipfsync

go 1.12

require (
	github.com/ipfs/interface-go-ipfs-core v0.0.6
	github.com/ipfsync/appserver v0.0.0
	github.com/ipfsync/ipfsmanager v0.0.0
	go.uber.org/fx v1.9.0
)

replace github.com/ipfsync/appserver => ../appserver

replace github.com/ipfsync/ipfsync => ../ipfsync

replace github.com/ipfsync/ipfsmanager => ../ipfsmanager
