// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wxspider

import (
	"testing"
)

func Test_SpiderArticle(t *testing.T) {
	t.Fatal(SpiderArticle(`http://mp.weixin.qq.com/s/qlU4E2WzvYrnmuxiBwxavw`))
}

func Test_PostArticle(t *testing.T) {
	err := PublishArticle()
	t.Fatal(err)
}

func Test_PostArticleOne(t *testing.T) {
	rows := GetArticles()
	for _, row := range rows {
		_, e := PostArticle(row)
		// t.Fatal(row)
		t.Fatal(e)
	}
	t.Fatal(rows)
}
