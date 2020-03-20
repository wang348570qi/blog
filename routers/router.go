package routers

import (
	"blog/middleware/jwt"
	"blog/pkg/export"
	"blog/pkg/qrcode"
	"blog/pkg/setting"
	"blog/pkg/upload"
	"blog/routers/api"
	"net/http"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	v1 "blog/routers/api/v1"

	_ "blog/docs"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//token
	r.GET("/auth", api.GetAuth)
	r.POST("/upload", api.UploadImage)

	//注册路由
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT()) //引入中间件
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新增
		apiv1.POST("/tags", v1.AddTag)
		//编辑
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		//二维码
		apiv1.POST("/article/poster/generate", v1.GenerateArticlePoster)

		//导出
		//r.POST("/tags/export", v1.ExportTag)
		//导入tag
		//r.POST("/tags/import", v1.ImportTag)

	}
	return r
}
