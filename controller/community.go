/**
    @Author:     ZonzeeLi
    @Project:    bluebell
    @CreateDate: 11/29/2021
    @UpdateDate: xxx
    @Note:       xxxx
**/

package controller

import (
	"bluebell/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ---- 跟社区相关的 ----

// CommunityHandler 获取社区信息详情的处理函数
// @Summary 获取社区信息详情的处理函数
// @Description 获取社区信息详情的接口函数
// @Tags 社区信息详情
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseCommunity
// @Router /posts2 [get]
func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区(community_id, community_name) 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 社区分类详情
// @Summary 获取社区分类详情的处理函数
// @Description 获取社区分类详情的接口函数
// @Tags 社区分类详情
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param string query string false "id"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseCommunityDetail
// @Router /posts2 [get]
func CommunityDetailHandler(c *gin.Context) {
	// 1. 获取社区id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 查询到所有的社区(community_id, community_name) 以列表的形式返回
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}
