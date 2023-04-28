package handler

import (
	"final-project/helper"
	"final-project/model"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetAll godoc
// @Summary      List of Photo
// @Description  List of Photo
// @Tags         Photo
// @Accept       json
// @Produce      json
// @Success      200  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /photo [get]
func (h HttpServer) GetAllPhoto(c *gin.Context) {

	userData := c.MustGet("userData").(jwt.MapClaims)
	in := model.Photo{}

	userID := int(userData["id"].(float64))

	in.UserID = userID

	res, err := h.app.GetAllPhoto(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// GetOne godoc
// @Summary      Get detail for given id
// @Description  Get detail of Photo by id
// @Tags         Photo
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID of Photo"
// @Success      200  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /photo/{id} [get]
func (h HttpServer) GetOnePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	in := model.Photo{}
	userID := int(userData["id"].(float64))
	in.UserID = userID

	res, err := h.app.GetOnePhoto(int64(idInt))
	if err != nil {
		if err.Error() == helper.ErrNotFound {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.OkWithMessage(c, "Data Social Media", res)
}

// Create Photo godoc
// @Summary      Created an Photo
// @Description  Created an Photo
// @Tags         Photo
// @Accept       json
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /photo [post]
func (h HttpServer) CreatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContent(c)

	in := model.Photo{}
	userID := int(userData["id"].(float64))

	if contentType == appJSON {
		err := c.ShouldBindJSON(&in)
		if err != nil {
			helper.BadRequest(c, err.Error())
			return
		}
	} else {
		err := c.ShouldBind(&in)
		if err != nil {
			helper.BadRequest(c, err.Error())
			return
		}
	}

	in.UserID = userID

	res, err := h.app.CreatePhoto(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Created(c, "Success Created Social Media", res)
}

// UpdatePhoto godoc
// @Summary      Update detail of Photo for given id
// @Description  update detail of Photo by id
// @Tags         Photo
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID of Photo"
// @Success      200  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /photo/{id} [put]
func (h HttpServer) UpdatePhoto(c *gin.Context) {

	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContent(c)

	in := model.Photo{}

	id, _ := strconv.Atoi(c.Param("id"))
	userID := int(userData["id"].(float64))

	if contentType == appJSON {
		err := c.ShouldBindJSON(&in)
		if err != nil {
			helper.BadRequest(c, err.Error())
			return
		}
	} else {
		err := c.ShouldBind(&in)
		if err != nil {
			helper.BadRequest(c, err.Error())
			return
		}
	}

	in.UserID = userID
	in.ID = uint(id)

	res, err := h.app.UpdatePhoto(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.OkWithMessage(c, "Social Media Updated: ", res)
}

// Delete Photo godoc
// @Summary      Delete of Photo data
// @Description  Delete of Photo by id
// @Tags         Photo
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID of Photo"
// @Success      200  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /photo/{id} [delete]
func (h HttpServer) DeletePhoto(c *gin.Context) {

	userData := c.MustGet("userData").(jwt.MapClaims)

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	in := model.Photo{}
	userID := int(userData["id"].(float64))
	in.UserID = userID

	err = h.app.DeletePhoto(int64(idInt))
	if err != nil {
		if err.Error() == helper.ErrNotFound {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.OkWithMessage(c, "Social Media Has Been Deleted", nil)
}
