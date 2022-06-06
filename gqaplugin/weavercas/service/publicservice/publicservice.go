package publicservice

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

func CasLogin(username string) (user gqaModel.SysUser, err error) {
	err = gqaGlobal.GqaDb.Where("username =?", username).Find(&user).Error
	return user, err
}

func LogLogin(username string, c *gin.Context, status string, detail string) (err error) {
	var loginLog gqaModel.SysLogLogin
	ua := user_agent.New(c.Request.UserAgent())

	loginLog.LoginUsername = username
	loginLog.LoginIp = c.ClientIP()
	name, version := ua.Browser()
	loginLog.LoginBrowser = name + " " + version
	loginLog.LoginOs = ua.OS()
	loginLog.LoginPlatform = ua.Platform()
	loginLog.LoginSuccess = status
	loginLog.Memo = detail + c.Request.UserAgent()

	err = gqaGlobal.GqaDb.Create(&loginLog).Error
	return err
}

func SaveOnline(username string, token string) error {
	var online = gqaModel.SysUserOnline{
		Username: username,
		Token:    token,
	}
	err := gqaGlobal.GqaDb.Where("username = ?", online.Username).Delete(&gqaModel.SysUserOnline{}).Error
	err = gqaGlobal.GqaDb.Create(&online).Error
	return err
}
