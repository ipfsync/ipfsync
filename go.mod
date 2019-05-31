module github.com/ipfsync/ipfsync

go 1.12

require (
	github.com/ipfsync/appserver v0.0.0
	github.com/ipfsync/common v0.0.0
	github.com/ipfsync/ipfsmanager v0.0.0
	github.com/ipfsync/resource v0.0.0
	github.com/libp2p/go-libp2p-net v0.0.2
	github.com/spf13/viper v1.4.0
	go.uber.org/fx v1.9.0
	google.golang.org/genproto v0.0.0-20180831171423-11092d34479b // indirect
)

replace github.com/ipfsync/appserver => ../appserver

replace github.com/ipfsync/ipfsync => ../ipfsync

replace github.com/ipfsync/ipfsmanager => ../ipfsmanager

replace github.com/ipfsync/common => ../common

replace github.com/ipfsync/resource => ../resource
