package Controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hulingfeng211/golang-study/011-data-visual-service/Database"
	"github.com/hulingfeng211/golang-study/011-data-visual-service/Model"
	"log"
	"net/http"
	//"strings"
)

// @Summary 查找所有看板
// @Description 查找所有看板
// @Tags Dashboard
// @Accept  json
// @Produce  json
// @Success 200 {object} Model.R
// @Failure 400 {object}  Model.R "We need ID!!"
// @Failure 404 {object}  Model.R "Can not find ID"
// @Router /visual/list [get]
func ListVisual(c *gin.Context) {
	var count int64
	var visual []Model.Visual
	Database.DB.Where("1=1").Find(&Model.Visual{}).Count(&count)
	data := make(map[string]interface{})
	Database.DB.Where("1=1").Find(&visual).Limit(100)
	data["records"] = visual
	data["total"] = count
	data["size"] = 100
	data["pages"] = 1
	c.JSON(http.StatusOK, Model.R{Data: data})
}

// @Summary 创建看板
// @Description 创建看板
// @Tags Dashboard
// @Accept  json
// @Produce  json
// @Param    data    body    Model.Visual     true        "看板"
// @Success 200 {object} Model.R
// @Failure 400 {object}  Model.R "We need ID!!"
// @Failure 404 {object}  Model.R "Can not find ID"
// @Router /visual/save [post]
func SaveVisual(c *gin.Context) {
	var visual Model.Visual
	//Database.DB.Where("1=1").Find(&Model.VisualCategory{}).Count(&count)
	//data:=make(map[string]interface{})
	if err := c.ShouldBind(&visual); err == nil {
		//c.String(http.StatusOK, `the body should be formA`)
		// 因为现在 c.Request.Body 是 EOF，所以这里会报错。
		//Database.DB.Save(visualCategory)
		Database.DB.Create(&visual)
		c.JSON(http.StatusOK, Model.R{Data: "OK"})
	}
}

// @Summary 上传封面
// @Description 上传封面
// @Tags Dashboard
// @Accept   multipart/form-data
// @Produce  json
// @Param   file formData file true  "this is a test file"
// @Success 200 {object} Model.R
// @Failure 400 {object}  Model.R "We need ID!!"
// @Failure 404 {object}  Model.R "Can not find ID"
// @Router /visual/put-file [post]
func PutFileWithVisual(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	dst := "views/" + file.Filename
	// 上传文件至指定目录
	c.SaveUploadedFile(file, dst)
	//sitePrefix:="http://localhost:8080"
	url := "/views/" + file.Filename
	data := make(map[string]interface{})
	data["link"] = url
	c.JSON(http.StatusOK, Model.R{Data: data})
}
