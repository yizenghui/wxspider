// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wxspider

import (
	"fmt"
	"testing"
)

func init() {
	DB().AutoMigrate(&Article{})

}
func Test_SpiderArticle(t *testing.T) {
	fmt.Println(SpiderArticle(`http://mp.weixin.qq.com/s/qlU4E2WzvYrnmuxiBwxavw`))
}

func Test_PostArticle(t *testing.T) {
	err := PublishArticle()
	t.Fatal(err)
}
