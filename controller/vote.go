/*
   @Author:     ZonzeeLi
   @Project:    bluebell
   @CreateTime: 2021/12/1
   @Note:
*/

package controller

import (
	"bluebell/logic"
	"bluebell/models"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// 投票

type VoteData struct {
	// UserID 从请求中获取当前的用户
	PostID    int64 `json:"post_id,string"`   // 帖子ID
	Direction int   `json:"direction,string"` // 赞成票(1)or反对票(-1)
}

// PostVoteController 帖子投票处理函数
// @Summary 帖子投票处理函数
// @Description 帖子投票的接口函数
// @Tags 帖子投票详情
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query models.ParamVoteData false "投票详情"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseSuccess
// @Router /posts2 [get]
func PostVoteController(c *gin.Context) {
	// 参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) // 翻译并去除掉错误提示中的结构体提示
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	userID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 具体投票的业务逻辑
	if err = logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
