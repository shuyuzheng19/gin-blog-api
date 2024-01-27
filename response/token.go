package response

// TokenResponse 登陆成功返回的token概要
// @Description 登陆成功返回的token概要
type TokenResponse struct {
	Token  string `json:"token"`  //token
	Expire string `json:"expire"` //过期时间
	Create string `json:"create"` //创建时间
}
