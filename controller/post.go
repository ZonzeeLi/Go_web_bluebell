/**
    @Author:     ZonzeeLi
    @Project:    bluebell
    @CreateDate: 11/30/2021
    @UpdateDate: xxx
    @Note:       xxxx
**/

package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CreatePostHandler 创建帖子的处理函数
// @Summary 创建帖子的处理函数
// @Description 创建帖子的接口函数
// @Tags 创建帖子
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query models.Post false "创建参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseSuccess
// @Router /posts2 [get]
func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数及参数的校验

	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 从 c 取到当前发请求的用户ID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 2. 创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子详情的处理函数
// @Summary 获取帖子详情的处理函数
// @Description 获取帖子详情的接口函数
// @Tags 帖子详情接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param string query string false "id"
// @Security ApiKeyAuth
// @Success 200 {object} models.ApiPostDetail
// @Router /posts2 [get]
func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取参数(从URL中获取帖子的id)
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 2. 根据id取出帖子数据
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler 获取帖子列表的处理函数
// @Summary 获取帖子列表的处理函数
// @Description 获取帖子列表的接口函
// @Tags 帖子列表接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param string query string false "page"
// @Param string query string false "size"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func GetPostListHandler(c *gin.Context) {
	page, size := GetPageInfo(c)
	// 获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func GetPostListHandler2(c *gin.Context) {
	// GET请求参数(query string): /api/v1/posts2?page=1&size=10&order=time
	// 初始化结构体时指定初始参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  100,
		Order: models.OrderTime, // magic string
	}
	// c.ShouldBind() 根据请求的数据类型选择相应的方法去获取数据
	// c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 可以这么做，但是这种判断最好放在logic里
	//var data []*models.ApiPostDetail
	//var err error
	//if p.CommunityID == 0 {
	//	// 获取数据
	//	data, err = logic.GetPostList2(p)
	//} else {
	//	data, err = logic.GetCommunityPostList(p)
	//}
	//if err != nil {
	//	zap.L().Error("logic.xxx failed", zap.Error(err))
	//	ResponseError(c, CodeServerBusy)
	//	return
	//}
	//// 返回响应
	//ResponseSuccess(c, data)

	data, err := logic.GetPostListNew(p)
	// 获取数据
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GetCommunityPostListHandler 根据社区去查询帖子列表(由于该部分代码和上述的getpost2类似，所以可以整合一下接口，结构体和后续调用参数均以整合)
//func GetCommunityPostListHandler(c *gin.Context) {
//	p := &models.ParamCommunityPostList{
//		ParamPostList: models.ParamPostList{
//			Page:  1,
//			Size:  10,
//			Order: models.OrderTime,
//		},
//		CommunityID: 1,
//	}
//	if err := c.ShouldBindQuery(p); err != nil {
//		zap.L().Error("GetCommunityPostListHandler with invalid params", zap.Error(err))
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//
//	// 获取参数
//	data, err := logic.GetCommunityPostList(p)
//	if err != nil {
//		zap.L().Error("logic.GetCommunityPostList() failed", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//	ResponseSuccess(c, data)
//}
