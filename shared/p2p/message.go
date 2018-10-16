package p2p

import (
	"context"
	"reflect"

	"github.com/golang/protobuf/proto"
)

// Message represents a message received from an external peer.
type Message struct {
	// Ctx message context.
	Ctx context.Context
	// Peer represents the sender of the message.
	Peer Peer
	// Data can be any type of message found in sharding/p2p/proto package.
	Data proto.Message
}

// messageType returns the underlying struct type for a given proto.message.
func messageType(msg proto.Message) reflect.Type {
	// proto.Message is a pointer and we need to dereference the pointer
	// and take the type of the original struct. Otherwise reflect.TypeOf will
	// create a new value of type **pb.BeaconBlockHashAnnounce for example.
	return reflect.ValueOf(msg).Elem().Type()
}
