package swagger_routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/example/basic/docs" // docs is generated by Swag CLI, you have to import it.
)

func Routers(e *gin.Engine) {
	router_swagger := e.Group("/swagger")
	{
		url := ginSwagger.URL("https://localhost:8779/public/docs/swagger.json") // The url pointing to API definition
		router_swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
}
