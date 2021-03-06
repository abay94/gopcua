// Copyright 2018 gopcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package uacp

import (
	"testing"

	"github.com/wmnsk/gopcua/utils/codectest"
)

func TestUACPMessage(t *testing.T) {
	cases := []codectest.Case{
		{
			Name: "Hello",
			Struct: NewHello(
				0,                                        //Version
				65280,                                    // ReceiveBufSize
				65535,                                    // SendBufSize
				4000,                                     // MaxMessageSize
				"opc.tcp://wow.its.easy:11111/UA/Server", // EndPointURL
			),
			Bytes: []byte{ // Hello message
				// MessageType: HEL
				0x48, 0x45, 0x4c,
				// Chunk Type: F
				0x46,
				// MessageSize: 70
				0x46, 0x00, 0x00, 0x00,
				// Version: 0
				0x00, 0x00, 0x00, 0x00,
				// ReceiveBufSize: 65280
				0x00, 0xff, 0x00, 0x00,
				// SendBufSize: 65535
				0xff, 0xff, 0x00, 0x00,
				// MaxMessageSize: 4000
				0xa0, 0x0f, 0x00, 0x00,
				// MaxChunkCount: 0
				0x00, 0x00, 0x00, 0x00,
				// EndPointURL
				0x26, 0x00, 0x00, 0x00, 0x6f, 0x70, 0x63, 0x2e,
				0x74, 0x63, 0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x6f,
				0x77, 0x2e, 0x69, 0x74, 0x73, 0x2e, 0x65, 0x61,
				0x73, 0x79, 0x3a, 0x31, 0x31, 0x31, 0x31, 0x31,
				0x2f, 0x55, 0x41, 0x2f, 0x53, 0x65, 0x72, 0x76,
				0x65, 0x72,
			},
		},
		{
			Name: "Acknowledge",
			Struct: NewAcknowledge(
				0,     //Version
				65280, // ReceiveBufSize
				65535, // SendBufSize
				4000,  // MaxMessageSize
			),
			Bytes: []byte{
				// MessageType: ACK
				0x41, 0x43, 0x4b,
				// Chunk Type: F
				0x46,
				// MessageSize: 28
				0x1c, 0x00, 0x00, 0x00,
				// Version: 0
				0x00, 0x00, 0x00, 0x00,
				// ReceiveBufSize: 65280
				0x00, 0xff, 0x00, 0x00,
				// SendBufSize: 65535
				0xff, 0xff, 0x00, 0x00,
				// MaxMessageSize: 4000
				0xa0, 0x0f, 0x00, 0x00,
				// MaxChunkCount: 0
				0x00, 0x00, 0x00, 0x00,
			},
		},
		{
			Name: "Error",
			Struct: NewError(
				BadSecureChannelClosed, // Error
				"foobar",
			),
			Bytes: []byte{
				// MessageType: ERR
				0x45, 0x52, 0x52,
				// Chunk Type: F
				0x46,
				// MessageSize: 22
				0x16, 0x00, 0x00, 0x00,
				// Error: BadSecureChannelClosed
				0x00, 0x00, 0x86, 0x80,
				// Reason: dummy
				0x06, 0x00, 0x00, 0x00, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72,
			},
		},
		{
			Name: "ReverseHello",
			Struct: NewReverseHello(
				"opc.tcp://wow.its.easy:11111/UA/Server", // ServerURI
				"opc.tcp://wow.its.easy:11111/UA/Server", // EndPointURL
			),
			Bytes: []byte{
				// MessageType: RHE
				0x52, 0x48, 0x45,
				// Chunk Type: F
				0x46,
				// MessageSize: 12
				0x5c, 0x00, 0x00, 0x00,
				// ServerURI
				0x26, 0x00, 0x00, 0x00, 0x6f, 0x70, 0x63, 0x2e,
				0x74, 0x63, 0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x6f,
				0x77, 0x2e, 0x69, 0x74, 0x73, 0x2e, 0x65, 0x61,
				0x73, 0x79, 0x3a, 0x31, 0x31, 0x31, 0x31, 0x31,
				0x2f, 0x55, 0x41, 0x2f, 0x53, 0x65, 0x72, 0x76,
				0x65, 0x72,
				// EndPointURL
				0x26, 0x00, 0x00, 0x00, 0x6f, 0x70, 0x63, 0x2e,
				0x74, 0x63, 0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x6f,
				0x77, 0x2e, 0x69, 0x74, 0x73, 0x2e, 0x65, 0x61,
				0x73, 0x79, 0x3a, 0x31, 0x31, 0x31, 0x31, 0x31,
				0x2f, 0x55, 0x41, 0x2f, 0x53, 0x65, 0x72, 0x76,
				0x65, 0x72,
			},
		},
		{
			Name: "Generic",
			Struct: NewGeneric(
				"XXX",
				"X",
				[]byte{0xde, 0xad, 0xbe, 0xef},
			),
			Bytes: []byte{
				// MessageType: XXX
				0x58, 0x58, 0x58,
				// Chunk Type: X
				0x58,
				// MessageSize: 12
				0x0c, 0x00, 0x00, 0x00,
				// dummy Payload
				0xde, 0xad, 0xbe, 0xef,
			},
		},
	}
	codectest.Run(t, cases, func(b []byte) (codectest.S, error) {
		v, err := Decode(b)
		if err != nil {
			return nil, err
		}
		switch got := v.(type) {
		case *Hello:
			got.Payload = nil
		case *Acknowledge:
			got.Payload = nil
		case *Error:
			got.Payload = nil
		case *ReverseHello:
			got.Payload = nil
		case *Generic:
			// do nothing
		default: // should not be called
			t.Fatalf("unexpected type: %T", v)
		}
		return v, nil
	})
}
