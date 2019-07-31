// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: message.proto

package message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Message service

type MessageService interface {
	Create(ctx context.Context, in *MessageCreateRequest, opts ...client.CallOption) (*MessageCreateResponse, error)
}

type messageService struct {
	c    client.Client
	name string
}

func NewMessageService(name string, c client.Client) MessageService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "message"
	}
	return &messageService{
		c:    c,
		name: name,
	}
}

func (c *messageService) Create(ctx context.Context, in *MessageCreateRequest, opts ...client.CallOption) (*MessageCreateResponse, error) {
	req := c.c.NewRequest(c.name, "Message.Create", in)
	out := new(MessageCreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Message service

type MessageHandler interface {
	Create(context.Context, *MessageCreateRequest, *MessageCreateResponse) error
}

func RegisterMessageHandler(s server.Server, hdlr MessageHandler, opts ...server.HandlerOption) error {
	type message interface {
		Create(ctx context.Context, in *MessageCreateRequest, out *MessageCreateResponse) error
	}
	type Message struct {
		message
	}
	h := &messageHandler{hdlr}
	return s.Handle(s.NewHandler(&Message{h}, opts...))
}

type messageHandler struct {
	MessageHandler
}

func (h *messageHandler) Create(ctx context.Context, in *MessageCreateRequest, out *MessageCreateResponse) error {
	return h.MessageHandler.Create(ctx, in, out)
}
