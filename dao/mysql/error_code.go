/**
    @Author:     ZonzeeLi
    @Project:    bluebell
    @CreateDate: 11/30/2021
    @UpdateDate: xxx
    @Note:       xxxx
**/

package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
	ErrorInvaildID       = errors.New("无效的ID")
)
