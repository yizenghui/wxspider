// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wxspider

import (
	"testing"
)

func Test_GetToken(t *testing.T) {

	c := NewClient(`FRVsHFvKNIaKn0jNbW33jftt`, `4pRzbjduWISmvC66HfoC1vPpGEdz7294`)
	c.Auth()
	t.Fatal(c.AccessToken)
}
