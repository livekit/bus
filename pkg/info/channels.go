// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package info

import (
	"unicode"

	"github.com/livekit/psrpc/internal/bus"
)

const lowerHex = "0123456789abcdef"

var channelChar = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x0030, 0x0039, 1}, // 0-9
		{0x0041, 0x005a, 1}, // A-Z
		{0x005f, 0x005f, 1}, // _
		{0x0061, 0x007a, 1}, // a-z
	},
	LatinOffset: 4,
}

type wildcard byte

var wc wildcard = '*'

func GetClaimRequestChannel(service, clientID string) bus.Channel {
	return bus.Channel{
		Legacy:  formatChannel('|', service, clientID, "CLAIM"),
		Primary: formatChannel('.', "CLI", service, clientID, "CLAIM"),
	}
}

func GetStreamChannel(service, nodeID string) bus.Channel {
	return bus.Channel{
		Legacy:  formatChannel('|', service, nodeID, "STR"),
		Primary: formatChannel('.', "CLI", service, nodeID, "STR"),
	}
}

func GetResponseChannel(service, clientID string) bus.Channel {
	return bus.Channel{
		Legacy:  formatChannel('|', service, clientID, "RES"),
		Primary: formatChannel('.', "CLI", service, clientID, "RES"),
	}
}

func (i *RequestInfo) GetRPCChannel() bus.Channel {
	return bus.Channel{
		Legacy:   formatChannel('|', i.Service, i.Method, i.Topic, "REQ"),
		Primary:  formatChannel('.', "SRV", i.Service, i.Method, i.Topic, "REQ"),
		Wildcard: formatChannel('.', "SRV", i.Service, wc, i.Topic, wc),
	}
}

func (i *RequestInfo) GetHandlerKey() string {
	return formatChannel('.', i.Method, i.Topic)
}

func (i *RequestInfo) GetClaimResponseChannel() bus.Channel {
	return bus.Channel{
		Legacy:   formatChannel('|', i.Service, i.Method, i.Topic, "RCLAIM"),
		Primary:  formatChannel('.', "SRV", i.Service, i.Method, i.Topic, "RCLAIM"),
		Wildcard: formatChannel('.', "SRV", i.Service, wc, i.Topic, wc),
	}
}

func (i *RequestInfo) GetStreamServerChannel() bus.Channel {
	return bus.Channel{
		Legacy:   formatChannel('|', i.Service, i.Method, i.Topic, "STR"),
		Primary:  formatChannel('.', "SRV", i.Service, i.Method, i.Topic, "STR"),
		Wildcard: formatChannel('.', "SRV", i.Service, wc, i.Topic, wc),
	}
}

func formatChannel(delim byte, parts ...any) string {
	buf := make([]byte, 0, 4*channelPartsLen(parts...)/3)
	return string(appendChannelParts(delim, buf, parts...))
}

func channelPartsLen[T any](parts ...T) int {
	var n int
	for _, t := range parts {
		switch v := any(t).(type) {
		case string:
			n += len(v) + 1
		case wildcard:
			n++
		case []string:
			n += channelPartsLen(v...)
		}
	}
	return n
}

func appendChannelParts[T any](delim byte, buf []byte, parts ...T) []byte {
	var prefix bool
	for _, t := range parts {
		if prefix {
			buf = append(buf, delim)
		}
		l := len(buf)
		switch v := any(t).(type) {
		case string:
			buf = appendSanitizedChannelPart(buf, v)
		case wildcard:
			buf = append(buf, '*')
		case []string:
			buf = appendChannelParts(delim, buf, v...)
		}
		prefix = len(buf) > l
	}
	return buf
}

func appendSanitizedChannelPart(buf []byte, s string) []byte {
	for _, r := range s {
		if unicode.Is(channelChar, r) {
			buf = append(buf, byte(r))
		} else if r < 0x10000 {
			buf = append(buf, `u+`...)
			for s := 12; s >= 0; s -= 4 {
				buf = append(buf, lowerHex[r>>uint(s)&0xF])
			}
		} else {
			buf = append(buf, `U+`...)
			for s := 28; s >= 0; s -= 4 {
				buf = append(buf, lowerHex[r>>uint(s)&0xF])
			}
		}
	}
	return buf
}
