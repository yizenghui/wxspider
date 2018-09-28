// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wxspider

import (
	"testing"
)

func Test_CheckImage(t *testing.T) {
	t.Fatal(CheckImage(`https://mmbiz.qpic.cn/mmbiz_jpg/oq1PymRl9D5kQRkOBLhsFymVWuZoEYqZbAdSUyKV8UOg4UMoibw8uUicZBsVxPVEBS5FNJYw3MnwGBVANYRTDhaA/640?wx_fmt=jpeg`))
}
