package common

import "time"

// 用户缓存键集合
const (
	UserTokenKey       = "USER-TOKEN:"    //缓存用户Token的Key
	EmailCodeKey       = "EMAIL-CODE:"    //缓存注册邮箱验证码的key
	EmailCodeKeyExpire = time.Minute      //邮箱验证码过期实际
	UserInfoKey        = "USER-INFO:"     //缓存用户信息的key
	UserInfoKeyExpire  = time.Minute * 30 //用户信息过期时间
)

// 博客缓存键集合
const (
	PageInfoPrefixKey  = "PAGE-INFO"       //缓存博客列表的key
	PageInfoExpire     = time.Hour * 3     //博客列表过期时间
	LatestBlogKey      = "LATEST-BLOG"     //缓存最新博客的key
	LatestBlogExpire   = time.Hour * 10    //最新博客过期时间
	HotBlogKey         = "HOT-BLOG"        //缓存热门博客的key
	HotBlogExpire      = time.Hour * 10    //热门博客过期时间
	RecommendKey       = "RECOMMEND-BLOG"  //缓存推荐博客的key
	WebSiteConfigKey   = "WEB-SITE-CONFIG" //缓存网站配置信息
	BlogMapKey         = "BLOG-MAP"        //缓存博客详情的key
	BlogEyeCountMapKey = "EYE-MAP"         //缓存博客的浏览量
)

// 分类缓存键集合
const (
	CategoryListKey    = "CATEGORY-LIST"    //缓存分类列表的key
	CategoryListExpire = time.Hour * 24 * 3 //分类列表过期时间
)

// 标签缓存键集合
const (
	RandomTagListKey    = "RANDOM-TAG"       //随机获取标签的数量
	RandomTagCount      = 25                 //随机获取标签的数量
	RandomTagListExpire = time.Hour * 24 * 3 //随机标签的过期时间
	TagMapKey           = "TAG-MAP"          //标签简要信息的key
)

// 专题缓存集合
const (
	TopicPageKey    = "TOPIC-PAGE:"      //缓存专题页的key
	TopicPageExpire = time.Hour * 24 * 3 //专题页过期时间
	TopicMapKey     = "TOPIC-MAP"        //专题简要信息的key
)

// 其他缓存集合
const (
	UserEditKey = "USER-EDIT-MAP"
)
