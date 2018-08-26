// Code generated by cdpgen. DO NOT EDIT.

package animation

import (
	"github.com/mafredri/cdp/rpcc"
)

// CanceledClient is a client for AnimationCanceled events. Event for when an animation has been canceled.
type CanceledClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*CanceledReply, error)
	rpcc.Stream
}

// CanceledReply is the reply for AnimationCanceled events.
type CanceledReply struct {
	ID string `json:"id"` // Id of the animation that was canceled.
}

// CreatedClient is a client for AnimationCreated events. Event for each animation that has been created.
type CreatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*CreatedReply, error)
	rpcc.Stream
}

// CreatedReply is the reply for AnimationCreated events.
type CreatedReply struct {
	ID string `json:"id"` // Id of the animation that was created.
}

// StartedClient is a client for AnimationStarted events. Event for animation that has been started.
type StartedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*StartedReply, error)
	rpcc.Stream
}

// StartedReply is the reply for AnimationStarted events.
type StartedReply struct {
	Animation Animation `json:"animation"` // Animation that was started.
}
