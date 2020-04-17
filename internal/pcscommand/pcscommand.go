// Package pcscommand 命令包
package pcscommand

import (
	"github.com/konglong87/BaiduPCS-Go/baidupcs"
	"github.com/konglong87/BaiduPCS-Go/internal/pcsconfig"
	"github.com/konglong87/BaiduPCS-Go/pcsverbose"
)

var (
	pcsCommandVerbose = pcsverbose.New("PCSCOMMAND")
)

// GetActiveUser 获取当前登录的百度帐号
func GetActiveUser() *pcsconfig.Baidu {
	return pcsconfig.Config.ActiveUser()
}

// GetBaiduPCS 从配置读取BaiduPCS
func GetBaiduPCS() *baidupcs.BaiduPCS {
	return pcsconfig.Config.ActiveUserBaiduPCS()
}
