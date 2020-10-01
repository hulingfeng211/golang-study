package Controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hulingfeng211/golang-study/011-data-visual-service/Database"
	"github.com/hulingfeng211/golang-study/011-data-visual-service/Model"
	"log"
	"net/http"
	"strings"
)

func CreateVisualMap(c *gin.Context) {
	visualMap := Model.VisualMap{}
	if err := c.ShouldBindBodyWith(&visualMap, binding.JSON); err == nil {
		//visualMap.CreatedAt=time.Now()
		//visualMap.UpdatedAt=time.Now()
		Database.DB.Create(&visualMap)
		c.JSON(http.StatusOK, Model.R{
			Code: 0,
			Data: visualMap,
		})
	}
}

/**
更新地图数据
*/
func UpdateVisualMap(c *gin.Context) {
	var visualMap Model.VisualMap
	//Database.DB.Where("1=1").Find(&Model.VisualCategory{}).Count(&count)
	//data:=make(map[string]interface{})
	if err := c.ShouldBind(&visualMap); err == nil {
		//c.String(http.StatusOK, `the body should be formA`)
		// 因为现在 c.Request.Body 是 EOF，所以这里会报错。
		Database.DB.Updates(visualMap)
		c.JSON(http.StatusOK, Model.R{Data: "OK"})
	}
}

/**
查找所有的地图列表
*/
func ListVisualMap(c *gin.Context) {

	var count int64
	var visualMaps []Model.VisualMap
	Database.DB.Where("1=1").Find(&Model.VisualMap{}).Count(&count)
	//db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	//Database.DB.Where("1=1").Count()
	//Database.DB.Count()
	/**
	Model.IPage{
		Total:   0,
		Size:    0,
		Current: 0,
		Pages:   1,
		Records: visualMaps,
	}}
	*/
	data := make(map[string]interface{})
	Database.DB.Where("1=1").Find(&visualMaps).Limit(100)
	data["records"] = visualMaps
	data["total"] = count
	data["size"] = 100
	data["pages"] = 1
	c.JSON(http.StatusOK, Model.R{Data: data})

}

/**
删除地图
*/
func DeleteVisualMap(c *gin.Context) {
	if cids, ok := c.GetQuery("ids"); ok {
		//db.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
		Database.DB.Where("id in (?)", strings.Split(cids, ",")).Delete(&Model.VisualMap{})
		//Database.DB.Delete(visualCategory)
		c.JSON(http.StatusOK, Model.R{Data: "OK"})
	} else {
		c.JSON(http.StatusOK, Model.R{Data: "OK"})
	}
}

/**
查找所有的分类列表
*/
func ListVisualCategory(c *gin.Context) {

	var count int64
	var visualCategory []Model.VisualCategory
	Database.DB.Where("1=1").Find(&Model.VisualCategory{}).Count(&count)
	//data:=make(map[string]interface{})
	Database.DB.Where("1=1").Find(&visualCategory).Limit(100)
	//data["records"]=visualCategory
	//data["total"]=count
	//data["size"]=100
	//data["pages"]=1
	c.JSON(http.StatusOK, Model.R{Data: visualCategory})

}

/**
根据分类信息获取分类明细
*/
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
	//data["records"]=visualCategory
	//data["total"]=count
	//data["size"]=100
	//data["pages"]=1

}

/**
更新分类
*/
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
	//if cid,ok:=c.GetQuery("id"); ok {
	//	Database.DB.Where("id=?",cid).Find(&visualCategory).Limit(1)
	//	c.JSON(http.StatusOK,Model.R{Data:visualCategory})
	//}else {
	//	c.JSON(http.StatusOK,Model.R{Data:visualCategory})
	//}

}

/**
创建新的分类
*/
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

/**
删除分类
*/
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

func PutFileWithVisual(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	dst := "views/" + file.Filename
	// 上传文件至指定目录
	c.SaveUploadedFile(file, dst)
	url := "http://localhost:8080/views/" + file.Filename
	data := make(map[string]interface{})
	data["link"] = url
	c.JSON(http.StatusOK, Model.R{Data: data})
}
