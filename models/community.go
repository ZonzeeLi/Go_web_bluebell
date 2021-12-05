/**
    @Author:     ZonzeeLi
    @Project:    bluebell
    @CreateDate: 11/29/2021
    @UpdateDate: xxx
    @Note:       xxxx
**/

package models

import "time"

type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	ID           int64     `json:"id" db:"community_id"`
	Name         string    `json:"name" db:"community_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
}
