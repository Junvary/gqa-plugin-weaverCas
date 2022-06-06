package publicapi

import (
	"encoding/xml"
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/weavercas/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/weavercas/service/publicservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	gqaUtils "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

func ValidateTicket(c *gin.Context) {
	var validateTicket model.ValidateTicket
	if err := gqaModel.RequestShouldBindJSON(c, &validateTicket); err != nil {
		return
	}
	respBody, err := GetValidateResp(validateTicket)
	if err != nil {
		gqaModel.ResponseErrorMessage("认证失败!", c)
	}
	type AuthenticationSuccess struct {
		User string `xml:"user"`
	}
	type CasResp struct {
		XMLName               xml.Name              `xml:"serviceResponse"`
		AuthenticationSuccess AuthenticationSuccess `xml:"authenticationSuccess"`
	}
	var cr CasResp
	if err = xml.Unmarshal(respBody, &cr); err != nil {
		fmt.Println(err)
	}
	if cr.AuthenticationSuccess.User == "" {
		gqaModel.ResponseErrorMessage("认证失败!", c)
	} else {
		gqaModel.ResponseSuccessMessageData(gin.H{"records": cr.AuthenticationSuccess.User}, "认证成功!", c)
	}
}

func GetValidateResp(validateTicket model.ValidateTicket) (respBody []byte, err error) {
	var url = "http://192.168.44.121/sso/serviceValidate?appid=" + validateTicket.AppId + "&service=" + validateTicket.Service + "&ticket=" + validateTicket.Ticket
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	respBody, err = ioutil.ReadAll(resp.Body)
	return respBody, err
}

func CasLogin(c *gin.Context) {
	type LoginModel struct {
		Username string `json:"username"`
	}
	var lm LoginModel
	if err := gqaModel.RequestShouldBindJSON(c, &lm); err != nil {
		return
	}
	user, err := publicservice.CasLogin(lm.Username)
	if err != nil {
		gqaModel.ResponseErrorMessage("登录失败!", c)
	} else {
		ss := gqaUtils.CreateToken(lm.Username)
		if ss == "" {
			gqaGlobal.GqaLogger.Error("Jwt配置错误，请重新初始化数据库！")
			gqaModel.ResponseErrorMessage("Jwt配置错误，请重新初始化数据库！", c)
			gqaModel.ResponseErrorMessage("登录失败!", c)
		}
		if err = publicservice.LogLogin(lm.Username, c, "yes", "登录成功！"); err != nil {
			gqaGlobal.GqaLogger.Error("登录日志记录错误！", zap.Any("err", err))
		}
		if err := publicservice.SaveOnline(lm.Username, ss); err != nil {
			gqaGlobal.GqaLogger.Error("记录在线用户失败！", zap.Any("err", err))
		}
		gqaModel.ResponseSuccessMessageData(gqaModel.ResponseLogin{
			Avatar:   user.Avatar,
			Username: user.Username,
			Nickname: user.Nickname,
			RealName: user.RealName,
			Token:    ss,
		}, "登录成功！", c)
	}
}
