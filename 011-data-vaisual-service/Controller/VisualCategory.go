package Controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hulingfeng211/golang-study/011-data-visual-service/Database"
	"github.com/hulingfeng211/golang-study/011-data-visual-service/Model"
	"net/http"
	"strings"
)

// @Summary 查找所有的分类列表
// @Description 查找所有的分类列表
// @Tags VisualCategory
// @Accept  json
// @Produce  json
// @Success 200 {object} Model.R
// @Failure 400 {object}  Model.R "We need ID!!"
// @Failure 404 {object}  Model.R "Can not find ID"
// @Router /category/list [get]
func ListVisualCategory(c *gin.Context) {
	var count int64
	var visualCategory []Model.VisualCategory
	Database.DB.Where("1=1").Find(&Model.VisualCategory{}).Count(&count)
	Database.DB.Where("1=1").Find(&visualCategory).Limit(100)
	c.JSON(http.StatusOK, Model.R{Data: visualCategory})
}

// @Summary 根据分类信息获取分类明细
// @Description 根据分类信息获取分类明细
// @Tags VisualCategory
// @Accept  json
// @Produce  json
// @Param   id     query    int     true        "Some ID"
// @Success 200 {object} Model.R
// @Failure 400 {object}  Model.R "We need ID!!"
// @Failure 404 {object}  Model.R "Can not find ID"
// @Router /category/detail [get]
func VisualCategoryDetail(c *gin.Context) {
	//var count int64
	var visualCategory Model.VisualCategory
	//Database.DB.Where("1=1").Find(&Model.VisualCategory{}).Count(&count)
	//data:=make(map[string]interface{})
	if cid, ok := c.GetQuery("id"); ok {
		Database.DB.Where("id=?", cid).Find(&visualCategory).Limit(1)
		c.JSON(http.StatusOK, Model.R{Data: visualCategory})
	} else {
		c.JSON(http.StatusOK, Model.R{Data: visualCategory})
	}
}

// @Summary 更新分类
// @Description 更新分类
// @Tags VisualCategory
// @Accept  json
// @Produce  json
// @Param    data    body    Model.VisualCategory     true        "更新分类"
// @Success 200 {object} Model.R
// @Failure 400 {object}  Model.R "We need ID!!"
// @Failure 404 {object}  Model.R "Can not find ID"
// @Router /category/update [post]
func UpdateVisualCategory(c *gin.Context) {
	var visualCategory Model.VisualCategory
	//Database.DB.Where("1=1").Find(&Model.VisualCategory{}).Count(&count)
	//data:=make(map[string]interface{})
	if err := c.ShouldBind(&visualCategory); err == nil {
		//c.String(http.StatusOK, `the body should be formA`)
		// 因为现在 c.Request.Body 是 EOF，所以这里会报错。
		Database.DB.Updates(visualCategory)
		c.JSON(http.StatusOK, Model.R{Data: "OK"})
	}

}

// @Summary 创建新的分类
// @Description 创建新的分类
// @Tags VisualCategory
// @Accept  json
// @Produce  json
// @Param    data    body    Model.VisualMap     true        "分类数据"
// @Success 200 {object} Model.R
// @Failure 400 {object}  Model.R "We need ID!!"
// @Failure 404 {object}  Model.R "Can not find ID"
// @Router /category/save [post]
func SaveVisualCategory(c *gin.Context) {
	var visualCategory Model.VisualCategory
	//Database.DB.Where("1=1").Find(&Model.VisualCategory{}).Count(&count)
	//data:=make(map[string]interface{})
	if err := c.ShouldBind(&visualCategory); err == nil {
		//c.String(http.StatusOK, `the body should be formA`)
		// 因为现在 c.Request.Body 是 EOF，所以这里会报错。
		//Database.DB.Save(visualCategory)
		Database.DB.Create(&visualCategory)
		c.JSON(http.StatusOK, Model.R{Data: "OK"})
	}
}

// @Summary 删除分类
// @Description 删除分类
// @Tags VisualCategory
// @Accept  json
// @Produce  json
// @Param ids body    string     true        "ID列表，多个以,分隔"
// @Success 200 {object} Model.R
// @Failure 400 {object}  Model.R "We need ID!!"
// @Failure 404 {object}  Model.R "Can not find ID"
// @Router /category/list [post]
func DeleteVisualCategory(c *gin.Context) {
	//var visualCategory Model.VisualCategory
	//Database.DB.Where("1=1").Find(&Model.VisualCategory{}).Count(&count)
	//data:=make(map[string]interface{})
	if cids, ok := c.GetQuery("ids"); ok {
		//db.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
		Database.DB.Where("id in (?)", strings.Split(cids, ",")).Delete(&Model.VisualCategory{})
		//Database.DB.Delete(visualCategory)
		c.JSON(http.StatusOK, Model.R{Data: "OK"})
	} else {
		c.JSON(http.StatusOK, Model.R{Data: "OK"})
	}
}
