package routes

import (
	"ginblog/api"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("front", "web/front/index.html")
	p.AddFromFiles("admin", "web/admin/index.html")
	return p
}

func InitRouter() { // 返回 *gin.Engine
	gin.SetMode(utils.AppMode) // 设置模式
	r := gin.New()             // 初始化路由，Default自带两个中间件（New()没有）
	r.HTMLRender = createMyRender()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.Static("admin/static", "./web/admin/static")
	r.Static("static", "./web/front/static")
	r.StaticFile("admin/favicon.ico", "./web/admin/favicon.ico")

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	auth := r.Group("api")
	auth.Use(middleware.JwtToken())
	{
		// User 模块路由接口
		auth.PUT("user/:id", api.EditUser)
		auth.PUT("admin/changepw/:id", api.ChangePassword)
		auth.DELETE("user/:id", api.DeleteUser)

		// Category 模块路由接口
		auth.POST("category/add", api.AddCategory)
		auth.PUT("category/:id", api.EditCategory)
		auth.DELETE("category/:id", api.DeleteCategory)

		// Article 模块路由接口
		auth.POST("article/add", api.AddArticle)
		auth.PUT("article/:id", api.EditArticle)
		auth.DELETE("article/:id", api.DeleteArticle)

		// Upload 文件上传路由接口
		auth.POST("upload", api.Upload)

		auth.PUT("profile/:id", api.UpdateProfile)
	}

	router := r.Group("api") // 路由组
	{
		router.POST("login", api.Login)

		// User 模块路由接口
		router.POST("user/add", api.AddUser)
		router.GET("user/:id", api.GetUserInfo)
		router.GET("users", api.GetUserList)

		// Category 模块路由接口
		router.GET("categories", api.GetCategoryList)
		router.GET("category/:id", api.GetCategoryInfo)

		// Article 模块路由接口
		router.GET("articles", api.GetArticleList)
		router.GET("article/info/:id", api.GetArticleById)
		router.GET("article/cateList/:cid", api.GetArticleUnderCategory)

		router.GET("profile/:id", api.GetProfile)
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
