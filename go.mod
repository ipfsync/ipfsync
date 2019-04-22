module github.com/ipfsync/ipfsync

go 1.12

require (
	github.com/ipfsync/appserver v0.0.0
	github.com/ipfsync/ipfsmanager v0.0.0
	go.uber.org/atomic v1.3.2 // indirect
	go.uber.org/dig v1.7.0 // indirect
	go.uber.org/fx v1.9.0
	go.uber.org/goleak v0.10.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
)

replace github.com/ipfsync/appserver => ../appserver

replace github.com/ipfsync/ipfsmanager => ../ipfsmanager
