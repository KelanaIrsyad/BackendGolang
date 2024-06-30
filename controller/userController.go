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

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType
	User := model.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"email":    User.Email,
		"username": User.Username,
		"age":      User.Age,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType
	User := model.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Find(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helper.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}
	token := helper.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UpdateUser(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)
	User := model.User{}
	id, _ := strconv.Atoi(c.Param("userId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.ID = userID
	err := db.Model(&User).Where("id=?", id).Updates(model.User{Username: User.Username, Email: User.Email, ProfilImgUrl: User.ProfilImgUrl}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":             User.ID,
		"email":          User.Email,
		"username":       User.Username,
		"age":            User.Age,
		"profil_img_url": User.ProfilImgUrl,
		"update_at":      User.UpdatedAt,
	})
}

func DeleteUser(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	id, _ := strconv.Atoi(c.Param("userId"))
	userID := uint(userData["id"].(float64))

	if userID != uint(id) {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "Forbidden",
			"message": "You don't have permission to delete this user",
		})
		return
	}

	err := db.Debug().Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfuly deleted",
	})
}
