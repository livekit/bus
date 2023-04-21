package interceptors

import (
	"context"

	"google.golang.org/protobuf/proto"

	"github.com/livekit/psrpc"
)

func ChainClientInterceptors[HandlerType any, InterceptorType ~func(psrpc.RPCInfo, HandlerType) HandlerType](
	interceptors []InterceptorType,
	info psrpc.RPCInfo,
	handler HandlerType,
) HandlerType {
	for i := len(interceptors) - 1; i >= 0; i-- {
		handler = interceptors[i](info, handler)
	}
	return handler
}

func ChainServerInterceptors(interceptors []psrpc.ServerRPCInterceptor) psrpc.ServerRPCInterceptor {
	switch n := len(interceptors); n {
	case 0:
		return nil
	case 1:
		return interceptors[0]
	default:
		return func(ctx context.Context, req proto.Message, info psrpc.RPCInfo, handler psrpc.ServerRPCHandler) (proto.Message, error) {
			// the struct ensures the variables are allocated together, rather than separately, since we
			// know they should be garbage collected together. This saves 1 allocation and decreases
			// time/call by about 10% on the microbenchmark.
			var state struct {
				i    int
				next psrpc.ServerRPCHandler
			}
			state.next = func(ctx context.Context, req proto.Message) (proto.Message, error) {
				if state.i == len(interceptors)-1 {
					return interceptors[state.i](ctx, req, info, handler)
				}
				state.i++
				return interceptors[state.i-1](ctx, req, info, state.next)
			}
			return state.next(ctx, req)
		}
	}
}
