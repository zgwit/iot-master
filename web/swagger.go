package web

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/zgwit/swagger-files"
)

func RegisterSwaggerDocs(app *gin.RouterGroup, instance string) {
	//注册接口文档
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.InstanceName(instance), ginSwagger.DefaultModelsExpandDepth(-1)))
}
