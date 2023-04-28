package middleware

import (
	"final-project/config"
	"final-project/helper"
	"final-project/model"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// func AllBookAuthorization() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		userData := c.MustGet("userData").(jwt.MapClaims)
// 		userID := uint(userData["id"].(float64))
// 		book := model.Book{}

// 		_ = config.GORM.Gorm.DB.Select("user_id").First(&book.UserID, uint(userID)).Error

// 		if book.UserID != userID {
// 			helper.Unauthorized(c, string(book.UserID))
// 			return
// 		}

// 		c.Next()
// 	}
// }

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			helper.BadRequest(c, "Invalid Parameter")
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := int(userData["id"].(float64))
		in := model.SocialMedia{}

		err = config.GORM.Gorm.DB.Select("user_id").First(&in, int(id)).Error
		if err != nil {
			helper.NotFound(c, "Data Not Found")
			return
		}

		if in.UserID != userID {
			helper.Unauthorized(c, "you are not allowrd to access this data")
			return
		}

		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			helper.BadRequest(c, "Invalid Parameter")
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := int(userData["id"].(float64))
		in := model.Photo{}

		err = config.GORM.Gorm.DB.Select("user_id").First(&in, int(id)).Error
		if err != nil {
			helper.NotFound(c, "Data Not Found")
			return
		}

		if in.UserID != userID {
			helper.Unauthorized(c, "you are not allowrd to access this data")
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		idPhoto, err := strconv.Atoi(c.Param("idPhoto"))
		if err != nil {
			helper.BadRequest(c, "Invalid Parameter")
			return
		}

		idComment, err := strconv.Atoi(c.Param("idComment"))
		if err != nil {
			helper.BadRequest(c, "Invalid Parameter")
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := int(userData["id"].(float64))
		in := model.Comment{}

		err = config.GORM.Gorm.DB.Select("photo_id").First(&in, int(idPhoto)).Error
		if err != nil {
			helper.NotFound(c, "Data Not Found")
			return
		}

		err = config.GORM.Gorm.DB.Select("user_id").First(&in, int(idComment)).Error
		if err != nil {
			helper.NotFound(c, "Data Not Found")
			return
		}

		if in.UserID != userID {
			helper.Unauthorized(c, "you are not allowrd to access this data")
			return
		}

		c.Next()
	}
}
