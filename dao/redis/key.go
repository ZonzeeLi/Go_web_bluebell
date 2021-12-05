/*
   @Author:     ZonzeeLi
   @Project:    bluebell
   @CreateTime: 2021/12/1
   @Note:
*/

package redis

// redis key

// redis key注意使用命名空间的方式，方便查询和拆分

const (
	Prefix             = "bluebell:"
	KeyPostTimeZSet    = "post:time"   // zset:帖子及发帖时间
	KeyPostScoreZSet   = "post:score"  // zset:帖子及投票的分数
	KeyPostVotedZSetPF = "post:voted:" // zset:记录用户及投票的类型;参数是post id
	KeyCommunitySetPF  = "community:"  // set:保存每个分区下帖子的id
)

// GetRedisKey 给redis key加上前缀
func GetRedisKey(key string) string {
	return Prefix + key
}
