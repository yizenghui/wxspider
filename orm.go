package wxspider

// GetArticleByURL Article 如果没有的话进行初始化
func (article *Article) GetArticleByURL(url string) {
	DB().Where(Article{URL: url}).FirstOrCreate(&article)
}

// GetArticleByID 通过ID获取Article
func (article *Article) GetArticleByID(id int64) {
	DB().First(&article, id)
}

// Save Article
func (article *Article) Save() {
	DB().Save(&article)
}

// GetPlanPublushArticle 获取计划发布的 []Article
func (article *Article) GetPlanPublushArticle() []Article {
	var articles []Article
	DB().Where("publish_at = 0").Order("id asc").Limit(100).Find(&articles)
	return articles
}
