// Copyright 2018 gopcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package services

import (
	"testing"
	"time"

	"github.com/wmnsk/gopcua/utils/codectest"
)

func TestChannelSecurityToken(t *testing.T) {
	cases := []codectest.Case{
		{
			Name: "Normal",
			Struct: NewChannelSecurityToken(
				1, 2, time.Date(2018, time.August, 10, 23, 0, 0, 0, time.UTC), 6000000,
			),
			Bytes: []byte{
				// ChannelID
				0x01, 0x00, 0x00, 0x00,
				// TokenID
				0x02, 0x00, 0x00, 0x00,
				// CreatedAt
				0x00, 0x98, 0x67, 0xdd, 0xfd, 0x30, 0xd4, 0x01,
				// RevisedLifetime
				0x80, 0x8d, 0x5b, 0x00,
			},
		},
	}
	codectest.Run(t, cases, func(b []byte) (codectest.S, error) {
		return DecodeChannelSecurityToken(b)
	})
}
