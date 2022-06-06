package weavercas

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/weaverCas/router/publicrouter"
	"github.com/gin-gonic/gin"
)

var PluginWeaverCas = new(weaverCas)

type weaverCas struct{}

func (*weaverCas) PluginCode() string { //实现接口方法，插件编码。返回值：请使用 "plugin-"前缀开头。
	return "plugin-weaverCas"
}

func (*weaverCas) PluginName() string { //实现接口方法，插件名称
	return "泛微CAS"
}

func (*weaverCas) PluginVersion() string { //实现接口方法，插件版本
	return "v0.0.1"
}

func (*weaverCas) PluginMemo() string { //实现接口方法，插件描述
	return "这是泛微CAS插件"
}

func (p *weaverCas) PluginRouterPublic(publicGroup *gin.RouterGroup) { //实现接口方法，公开路由初始化
	publicrouter.InitPublicRouter(publicGroup)
}

func (p *weaverCas) PluginRouterPrivate(privateGroup *gin.RouterGroup) { //实现接口方法，鉴权路由初始化
}

func (p *weaverCas) PluginMigrate() []interface{} { //实现接口方法，迁移插件数据表
	return nil
}

func (p *weaverCas) PluginData() []interface{ LoadData() (err error) } { //实现接口方法，初始化数据
	return nil
}
