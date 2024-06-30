package controller

import (
	"belajar/golang/database"
	"belajar/golang/helper"
	"belajar/golang/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)
	SocialMedia := model.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserID,
		"created_at":       SocialMedia.CreatedAt,
	})
}

func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := []model.SocialMedia{}

	db.Find(&SocialMedia)
	if len(SocialMedia) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "No user found",
			"error_message": "There are no user found.",
		})
		return
	}
	User := model.User{}
	db.Find(&User)

	c.JSON(http.StatusOK, gin.H{
		"social_medias": SocialMedia,
		"User": gin.H{
			"id":             User.ID,
			"username":       User.Username,
			"profil_img_url": User.ProfilImgUrl,
		},
	})
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)
	SocialMedia := model.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaId)

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaId).Updates(model.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserID,
		"updated_at":       SocialMedia.UpdatedAt,
	})
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := model.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	err := db.Where("id = ?", socialMediaId).Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Delete Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Your social media has been successfully deleted",
	})
}
