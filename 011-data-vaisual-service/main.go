package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hulingfeng211/golang-study/011-data-visual-service/Controller"
	"github.com/hulingfeng211/golang-study/011-data-visual-service/Database"
	"github.com/hulingfeng211/golang-study/011-data-visual-service/Lib"
	"github.com/hulingfeng211/golang-study/011-data-visual-service/Model"
	_ "github.com/hulingfeng211/golang-study/011-data-visual-service/docs"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"strconv"
	"time"
)

func mapDetailHandler(c *gin.Context) {
	mapId := c.Query("id")
	var visualMap Model.VisualMap
	if s, err := strconv.Atoi(mapId); err == nil {
		Database.DB.Where("id=?", s).First(&visualMap)
		//if err := Database.DB.Find(&visualMap).Error; err != nil {
		//	log.Fatal(err.Error())
		//}
		c.JSON(http.StatusOK, Model.R{Data: visualMap})
	}

}
func routeGroup(router *gin.Engine) {

	mapRoute := router.Group("/map")
	{
		//mapRoute.GET("/detail",mapDetailHandler)
		mapRoute.GET("/detail", mapDetailHandler)
		mapRoute.PUT("/save", Controller.CreateVisualMap)
		mapRoute.POST("/update", Controller.UpdateVisualMap)
		mapRoute.GET("/list", Controller.ListVisualMap)
		mapRoute.POST("/remove", Controller.DeleteVisualMap)
		//mapRoute.GET("/detail",mapDetailHandler)
		//mapRoute.GET("/detail",mapDetailHandler)

	}
	mapCategory := router.Group("/category")
	{
		//mapRoute.GET("/detail",mapDetailHandler)
		//mapCategory.GET("/data",mapDetailHandler)
		//mapCategory.PUT("/save",Controller.CreateVisualMap)
		mapCategory.GET("/list", Controller.ListVisualCategory)
		mapCategory.GET("/detail", Controller.VisualCategoryDetail)
		mapCategory.POST("/update", Controller.UpdateVisualCategory)
		mapCategory.POST("/save", Controller.SaveVisualCategory)
		mapCategory.POST("/remove", Controller.DeleteVisualCategory)
		//mapRoute.GET("/detail",mapDetailHandler)
		//mapRoute.GET("/detail",mapDetailHandler)

	}

	visualRouteGroup := router.Group("/visual")
	{
		//mapRoute.GET("/detail",mapDetailHandler)
		//mapCategory.GET("/data",mapDetailHandler)
		//mapCategory.PUT("/save",Controller.CreateVisualMap)
		visualRouteGroup.GET("/list", Controller.ListVisual)
		visualRouteGroup.POST("/save", Controller.SaveVisual)
		visualRouteGroup.POST("/put-file", Controller.PutFileWithVisual)
		//mapRoute.GET("/detail",mapDetailHandler)
		//mapRoute.GET("/detail",mapDetailHandler)

	}

}

func main() {
	gin.SetMode(gin.DebugMode)
	//migrate()
	dsn := "root:123mmm@tcp(192.168.99.100:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	Database.DB = Database.Migrate(dsn, &Model.VisualMap{}, &Model.VisualCategory{}, &Model.Visual{})
	r := gin.Default()
	r.Use(Lib.StaticServe("/views", Lib.LocalFile("views", true)))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))
	routeGroup(r)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run(":8080")

}
