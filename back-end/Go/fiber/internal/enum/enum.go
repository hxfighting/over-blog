package enum

type CacheKey string

func (c CacheKey) String() string {
	return string(c)
}

const (
	CacheConfig   CacheKey = "blog_config"
	CacheRotation CacheKey = "blog_rotation"
	CachePhoto    CacheKey = "blog_photo"
	CacheCategory CacheKey = "blog_category"
	CacheSocial   CacheKey = "blog_social"
	CacheTag      CacheKey = "blog_tag"
	CacheArticle  CacheKey = "blog_article"
	CacheComment  CacheKey = "blog_comment"
	CacheLink     CacheKey = "blog_link"
	CacheFooter   CacheKey = "blog_footer"
)
