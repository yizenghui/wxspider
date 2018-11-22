// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wxspider

import (
	"testing"
)

func Test_GetToken(t *testing.T) {
	cf := GetConf()
	c := NewClient(cf.BaiDuAiConf.APIKey, cf.BaiDuAiConf.SecretKey)
	c.Auth()
	t.Fatal(c.AccessToken)
}
