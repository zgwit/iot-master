package web

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/zgwit/swagger-files"
)

func RegisterSwaggerDocs(app *gin.Engine) {
	//注册接口文档
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
