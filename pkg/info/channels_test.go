package info

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/livekit/psrpc"
)

func TestChannelFormatters(t *testing.T) {
	i := &RequestInfo{
		RPCInfo: psrpc.RPCInfo{
			Service: "foo",
			Method:  "bar",
			Topic:   nil,
		},
	}

	require.Equal(t, "foo|bar|RES", GetResponseChannel("foo", "bar"))
	require.Equal(t, "foo|bar|CLAIM", GetClaimRequestChannel("foo", "bar"))
	require.Equal(t, "foo|bar|STR", GetStreamChannel("foo", "bar"))

	require.Equal(t, "foo|bar|REQ", i.GetRPCChannel())
	require.Equal(t, "foo|bar|RCLAIM", i.GetClaimResponseChannel())
	require.Equal(t, "foo|bar|STR", i.GetStreamServerChannel())

	i.Topic = []string{"a", "b", "c"}

	require.Equal(t, "foo|bar|a|b|c|REQ", i.GetRPCChannel())
	require.Equal(t, "bar|a|b|c", i.GetHandlerKey())
	require.Equal(t, "foo|bar|a|b|c|RCLAIM", i.GetClaimResponseChannel())
	require.Equal(t, "foo|bar|a|b|c|STR", i.GetStreamServerChannel())

	require.Equal(t, "U+0001f680_u+00c9|U+0001f6f0_bar|u+8f6fu+4ef6|END", formatChannel("🚀_É", "🛰_bar", []string{"软件"}, "END"))
}
