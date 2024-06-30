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

func CreateComment(c *gin.Context) {

	var commentRequest struct {
		PhotoID uint   `json:"photo_id" form:"photo_id"`
		Message string `json:"message" form:"message" binding:"required"`
	}
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&commentRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&commentRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	comment := model.Comment{
		PhotoID: uint(commentRequest.PhotoID),
		Message: commentRequest.Message,
		UserID:  userID,
	}

	err := db.Debug().Create(&comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        comment.ID,
		"message":   comment.Message,
		"photo_id":  comment.PhotoID,
		"user_id":   comment.UserID,
		"create_at": comment.CreatedAt,
	})

}

func GetComment(c *gin.Context) {
	db := database.GetDB()
	GetComment := []model.Comment{}

	db.Find(&GetComment)

	if len(GetComment) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "No Photos found",
			"error_message": "There are no photos found.",
		})
		return
	}

	User := model.User{}
	db.Find(&User)
	Photo := model.Photo{}
	db.Find(&Photo)

	c.JSON(http.StatusOK, gin.H{
		"Comment": GetComment,
		"User": gin.H{
			"id":       User.ID,
			"email":    User.Email,
			"username": User.Username,
		},
		"Photo": gin.H{
			"id":        Photo.ID,
			"title":     Photo.Title,
			"caption":   Photo.Caption,
			"photo_url": Photo.PhotoURL,
			"user_id":   Photo.UserID,
		},
	})
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)
	Comment := model.Comment{}
	Photo := model.Photo{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(model.Comment{Message: Comment.Message}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	db.Find(&Photo)
	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"user_id":    Comment.UserID,
		"updated_at": Comment.UpdatedAt,
	})
}

func DeleteComent(c *gin.Context) {
	db := database.GetDB()
	Comment := model.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))

	err := db.Where("id = ?", commentId).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Delete Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
