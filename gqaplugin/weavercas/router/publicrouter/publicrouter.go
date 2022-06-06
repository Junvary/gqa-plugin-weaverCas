package publicrouter

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/weavercas/api/publicapi"
	"github.com/gin-gonic/gin"
)

func InitPublicRouter(publicGroup *gin.RouterGroup) {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。
	{
		//验证ticket
		publicGroup.POST("validate-ticket", publicapi.ValidateTicket)
		//cas登录
		publicGroup.POST("cas-login", publicapi.CasLogin)
	}
}
