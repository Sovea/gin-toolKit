package routers
import(
	"net/http"
	"github.com/gin-gonic/gin"
)
type Option func(*gin.Engine)

var options = []Option{}

//Cors cross-domain solution(Gin Middleware)
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,X-Nideshop-Token, x-nideshop-token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
// Include will combine some options.
func Include(opts ...Option) {
	options = append(options, opts...)
}
// IncludeWith will combine some options to a specified Gin instance.
func IncludeWith(r *gin.Engine, opts ...Option) {
	options = append(options, opts...)
	for _, opt := range options {
		opt(r)
	}
}
// AssembledInit will create a Gin instance with some options.
func AssembledInit() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())
	for _, opt := range options {
		opt(r)
	}
	return r
}
//RawInit will create a default Gin instance.
func RawInit() *gin.Engine {
	r := gin.Default()
	return r
}