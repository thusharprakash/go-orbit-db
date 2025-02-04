package replicator

import (
	"context"

	cid "github.com/ipfs/go-cid"
	coreiface "github.com/ipfs/kubo/core/coreiface"
	"github.com/libp2p/go-libp2p/core/event"
	ipfslog "github.com/thusharprakash/go-ipfs-log"
	"github.com/thusharprakash/go-ipfs-log/identityprovider"
	"github.com/thusharprakash/go-orbit-db/accesscontroller"
)

// storeInterface An interface used to avoid import cycles
type storeInterface interface {
	OpLog() ipfslog.Log
	IPFS() coreiface.CoreAPI
	Identity() *identityprovider.Identity
	AccessController() accesscontroller.Interface
	SortFn() ipfslog.SortFn
	IO() ipfslog.IO
}

// Replicator Replicates stores information among peers
type Replicator interface {
	// Stop Stops the replication
	Stop()

	// Load Loads new data to replicate
	Load(ctx context.Context, heads []ipfslog.Entry)

	// GetQueue Returns the list of CID in the queue
	GetQueue() []cid.Cid

	EventBus() event.Bus
}

// ReplicationInfo Holds information about the current replication state
type ReplicationInfo interface {
	// GetProgress Get the value of progress
	GetProgress() int

	// GetProgress Get the value of progress
	SetProgress(i int)

	// GetMax Get the value of max
	GetMax() int

	// SetMax Sets the value of max
	SetMax(i int)

	// Reset Resets all values to 0
	Reset()
}
