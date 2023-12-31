package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type Punches struct {
	gorm.Model
	UserID uint
}

// ShowPunch show punch state
func ShowPunch(context *gin.Context) {
	var punches []Punches
	weekAgo := time.Now().AddDate(0, 0, -7)
	DB.Where("created_at > ?", weekAgo).Find(&punches)
	context.JSON(200, gin.H{
		"punches": punches,
	})
}

// FormPost  post punches tables createTime
func FormPost(context *gin.Context) {

	formData := context.PostForm("punch")
	var user User
	DB.Where("name = ?", formData).Find(&user)

	punch := Punches{
		Model:  gorm.Model{},
		UserID: user.ID,
	}
	DB.Create(&punch)
	context.JSON(200, gin.H{
		"code": punch.UserID,
	})

}
