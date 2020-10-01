package Controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hulingfeng211/golang-study/011-data-visual-service/Database"
	"github.com/hulingfeng211/golang-study/011-data-visual-service/Model"
	"net/http"
	"strconv"
	"strings"
)

// @Summary 创建可视化地图
// @Description 通过地图编号获取地图数据明细
// @Tags VisualMap
// @Accept  json
// @Produce  json
// @Param    data    body    Model.VisualMap     true        "可视化地图数据"
// @Success 200 {object} Model.R
// @Failure 400 {object}  Model.R "We need ID!!"
// @Failure 404 {object}  Model.R "Can not find ID"
// @Router /map/save [post]
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

// @Summary 获取地图明细数据
// @Description 通过地图编号获取地图数据明细
// @Tags VisualMap
// @Accept  json
// @Produce  json
// @Param   id     query    int     true        "Some ID"
// @Success 200 {object} Model.R
// @Failure 400 {object}  Model.R "We need ID!!"
// @Failure 404 {object}  Model.R "Can not find ID"
// @Router /map/detail [get]
func VisualMapDetail(c *gin.Context) {
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

// @Summary 更新地图数据
// @Description 更新地图数据
// @Tags VisualMap
// @Accept  json
// @Produce  json
// @Param    data    body    Model.VisualMap     true        "可视化地图数据"
// @Success 200 {object} Model.R
// @Failure 400 {object}  Model.R "We need ID!!"
// @Failure 404 {object}  Model.R "Can not find ID"
// @Router /map/update [post]
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

// @Summary 查找所有的地图列表
// @Description 查找所有的地图列表
// @Tags VisualMap
// @Accept  json
// @Produce  json
// @Success 200 {object} Model.R
// @Failure 400 {object}  Model.R "We need ID!!"
// @Failure 404 {object}  Model.R "Can not find ID"
// @Router /map/list [get]
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

// @Summary 删除地图
// @Description 删除地图
// @Tags VisualMap
// @Accept  json
// @Produce  json
// @Param ids body    string     true        "ID列表，多个以,分隔"
// @Success 200 {object} Model.R
// @Failure 400 {object}  Model.R "We need ID!!"
// @Failure 404 {object}  Model.R "Can not find ID"
// @Router /map/list [post]
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
