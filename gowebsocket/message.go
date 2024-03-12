package gowebsocket

import (
	"context"
)

type Message interface {
	GetID() string
	GetRoute() Route
	GetPayload() string
}

type Route string

type MessageParser func(ctx context.Context, raw []byte) (Message, error)

var messageParser MessageParser

func SetMessageParser(parser MessageParser) {
	messageParser = parser
}
