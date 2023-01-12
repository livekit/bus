package psrpc

import (
	"time"
)

const (
	DefaultClientTimeout = time.Second * 3
)

type ClientOption func(*clientOpts)

type clientOpts struct {
	timeout       time.Duration
	channelSize   int
	requestHooks  []ClientRequestHook
	responseHooks []ClientResponseHook
}

func WithClientTimeout(timeout time.Duration) ClientOption {
	return func(o *clientOpts) {
		o.timeout = timeout
	}
}

func WithClientChannelSize(size int) ClientOption {
	return func(o *clientOpts) {
		o.channelSize = size
	}
}

func WithClientRequestHooks(hooks ...ClientRequestHook) ClientOption {
	return func(o *clientOpts) {
		for _, hook := range hooks {
			if hook != nil {
				o.requestHooks = append(o.requestHooks, hook)
			}
		}
	}
}

func WithClientResponseHooks(hooks ...ClientResponseHook) ClientOption {
	return func(o *clientOpts) {
		for _, hook := range hooks {
			o.responseHooks = append(o.responseHooks, hook)
		}
	}
}

func getClientOpts(opts ...ClientOption) clientOpts {
	o := &clientOpts{
		timeout:     DefaultClientTimeout,
		channelSize: DefaultChannelSize,
	}
	for _, opt := range opts {
		opt(o)
	}
	return *o
}
