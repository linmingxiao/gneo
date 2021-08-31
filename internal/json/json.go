// Copyright 2017 Bo-Yi Wu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !jsoniter && !go_json
// +build !jsoniter,!go_json

package json

import "encoding/json"

var (
	Marshal = json.Marshal
	Unmarshal = json.Unmarshal
	MarshalIndent = json.MarshalIndent
	NewDecoder = json.NewDecoder
	NewEncoder = json.NewEncoder
)
