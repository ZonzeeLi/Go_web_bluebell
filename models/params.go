package models

// 定义请求参数结构体

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`                     // 输入用户名
	Password   string `json:"password" binding:"required"`                     // 输入密码
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // 再次输入密码
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"` // 输入用户名
	Password string `json:"password" binding:"required"` // 输入密码
}

// ParamVoteData 投票数据
type ParamVoteData struct {
	// UserID 从请求中获取当前的用户
	PostID    string `json:"post_id" binding:"required"`                          // 帖子ID
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1" example:"1"` // 赞成票(1)or反对票(-1),取消投票(0)
}

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`   // 可以为空，如果传了社区id则根据社区id去查找，否则返回全部帖子
	Page        int64  `json:"page" form:"page"`                   // 页码
	Size        int64  `json:"size" form:"size"`                   // 每页数据量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}

// ParamCommunityPostList 按社区获取帖子列表query string参数
type ParamCommunityPostList struct {
	ParamPostList
	CommunityID int64 `json:"community_id" form:"community_id"`
}
