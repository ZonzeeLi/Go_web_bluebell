/*
   @Author:     ZonzeeLi
   @Project:    bluebell
   @CreateTime: 2021/12/1
   @Note:
*/

package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"
)

// VoteForPost 为帖子投票的函数
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("PostID", p.PostID),
		zap.Int8("Direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
