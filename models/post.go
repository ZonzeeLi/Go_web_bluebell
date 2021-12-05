/**
    @Author:     ZonzeeLi
    @Project:    bluebell
    @CreateDate: 11/30/2021
    @UpdateDate: xxx
    @Note:       xxxx
**/

package models

import (
	"time"
)

// Go 内存对齐概念(可以节省结构体占用内存)

type Post struct {
	ID          int64     `json:"id,string" db:"post_id" ` // tag里json加上string，解决前端数字id失真问题
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`     // 帖子题目
	Content     string    `json:"content" db:"content" binding:"required"` // 帖子详情
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// ApiPostDetail 帖子详情借口的结构体
type ApiPostDetail struct {
	AuthorName       string             `json:"author_name"`
	VoteNum          int64              `json:"vote_num"`
	*Post                               // 嵌入帖子结构体
	*CommunityDetail `json:"community"` // 嵌入社区信息
}
