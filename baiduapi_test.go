// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wxspider

import (
	"fmt"
	"testing"
)

func Test_GetTags(t *testing.T) {
	url := fmt.Sprintf(`https://aip.baidubce.com/rpc/2.0/nlp/v1/keyword?access_token=%v`, GetToken())

	t.Fatal(url)
}
