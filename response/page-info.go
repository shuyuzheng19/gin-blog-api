package response

import "encoding/json"

// PageInfo 分页
// @Description 全局分页
type PageInfo struct {
	Page  int         `json:"page"`  //页码
	Count int64       `json:"total"` //总共多少条
	Size  int         `json:"size"`  //每页大小
	Data  interface{} `json:"data"`  //分页数据
}

func (p *PageInfo) MarshalBinary() (data []byte, err error) {
	return json.Marshal(p)
}

func (p *PageInfo) UnmarshalBinary(data []byte) (err error) {
	return json.Unmarshal(data, p)
}
