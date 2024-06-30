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

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)

	Photo := model.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":        Photo.ID,
		"title":     Photo.Title,
		"caption":   Photo.Caption,
		"photo_url": Photo.PhotoURL,
		"user_id":   Photo.UserID,
		"create_at": Photo.CreatedAt,
	})

}

func GetPhoto(c *gin.Context) {
	db := database.GetDB()
	getPhotos := []model.Photo{}

	db.Find(&getPhotos)

	if len(getPhotos) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "No Photos found",
			"error_message": "There are no photos found.",
		})
		return
	}
	Photo := model.Photo{}
	db.Find(&Photo)
	User := model.User{}
	db.Find(&User)

	c.JSON(http.StatusOK, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoURL,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
		"update_at":  Photo.UpdatedAt,
		"User": gin.H{
			"email":    User.Email,
			"username": User.Username,
		},
	})
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)
	Photo := model.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(model.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoURL: Photo.PhotoURL}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        Photo.ID,
		"title":     Photo.Title,
		"caption":   Photo.Caption,
		"photo_url": Photo.PhotoURL,
		"user_id":   Photo.UserID,
		"update_at": Photo.UpdatedAt,
	})
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := model.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))

	err := db.Where("id = ?", photoId).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Delete Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
