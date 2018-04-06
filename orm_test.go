// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wxspider

import (
	"testing"
)

func Test_GetPlanPublushArticle(t *testing.T) {

	var article Article
	articles := article.GetPlanPublushArticle()
	t.Fatal(articles)
}
