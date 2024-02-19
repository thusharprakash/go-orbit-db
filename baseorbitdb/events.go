package baseorbitdb

import (
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/thusharprakash/go-orbit-db/iface"
)

type EventExchangeHeads struct {
	Peer    peer.ID
	Message *iface.MessageExchangeHeads
}

func NewEventExchangeHeads(p peer.ID, msg *iface.MessageExchangeHeads) EventExchangeHeads {
	return EventExchangeHeads{
		Peer:    p,
		Message: msg,
	}
}
