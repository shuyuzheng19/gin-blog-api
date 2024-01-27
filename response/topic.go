package response

// TopicResponse 专题返回概要
// @Description 专题返回概要
type TopicResponse struct {
	Id          int                `json:"id"`          //专题id
	Name        string             `json:"name"`        //专题名
	Description string             `json:"description"` //专题描述
	CoverImage  string             `json:"cover"`       //专题封面
	CreatedAt   timeStamp          `json:"timeStamp"`   //专题创建时间戳
	User        SimpleUserResponse `json:"user"`
	UserId      int                `json:"-"`
}

// UserTopicResponse 用户专题返回概要
// @Description 用户专题返回概要
type UserTopicResponse struct {
	Id         int    `json:"id"`    //专题id
	Name       string `json:"name"`  //专题名子
	CoverImage string `json:"cover"` //专题封面
}

// SimpleTopicResponse 简洁的专题返回概要
// @Description 简洁的专题返回概要信息
type SimpleTopicResponse struct {
	Id   int    `json:"id"`   //专题id
	Name string `json:"name"` //专题名子
}

// AdminTopicResponse 后台管理专题模型
// @Description 后台管理专题模型
type AdminTopicResponse struct {
	Id          int                `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	CreatedAt   myTime             `json:"createAt"`
	UpdatedAt   myTime             `json:"updateAt"`
	CoverImage  string             `json:"coverImage"`
	User        SimpleUserResponse `json:"user"`
	UserId      int                `json:"-"`
}
