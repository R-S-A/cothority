package skipchain

import (
	"github.com/dedis/cothority/lib/network"
	"github.com/dedis/cothority/lib/sda"
)

func init() {
	network.RegisterMessageType(&RNBRet{})
}

// This file holds all messages that can be sent to the SkipChain,
// both from the outside and between instances of this service

// External calls

// RequestNewBlock - Requests a new skipblock to be appended to
// the given SkipBlock. If the given SkipBlock has Index 0 (which
// is invalid), a new SkipChain will be created.
// The AppId will be used to call the corresponding verification-
// routines who will have to sign off on the new Tree.
type RequestNewBlock struct {
	AppId string
	*SkipBlock
	*sda.EntityList
}

// AddRet - returns the signed SkipBlock with updated backlinks
type RNBRet struct {
	*SkipBlock
}

// GetUpdateChain - the client sends the hash of the last known
// Skipblock and will get back a list of all necessary SkipBlocks
// to get to the latest.
type GetUpdateChain struct {
	*SkipBlock
}

// GetUpdateChainRet - returns the shortest chain to the current SkipBlock,
// starting from the SkipBlock the client sent
type GUCRet struct {
	Update []*SkipBlock
}

// Internal calls

// PropagateSkipBlock sends a newly signed SkipBlock to all members of
// the Cothority
type PropagateSkipBlock struct {
	*SkipBlock
}

// ForwardSignature asks this responsible for a SkipChain to sign off
// a new ForwardLink. This will probably be sent to all members of any
// SkipChain-definition at time 'n'
type ForwardSignature struct {
	Old    *SkipBlock
	Latest *SkipBlock
}
