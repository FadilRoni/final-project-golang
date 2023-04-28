package handler

import (
	"final-project/helper"
	"final-project/model"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetAll godoc
// @Summary      List of Social Media
// @Description  List of Social Media
// @Tags         Social Media
// @Accept       json
// @Produce      json
// @Success      200  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /social_media [get]
func (h HttpServer) GetAllSocialMedia(c *gin.Context) {

	userData := c.MustGet("userData").(jwt.MapClaims)
	in := model.SocialMedia{}

	userID := int(userData["id"].(float64))

	in.UserID = userID

	res, err := h.app.GetAllSocialMedia(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// GetOne godoc
// @Summary      Get detail for given id
// @Description  Get detail of Social Media by id
// @Tags         Social Media
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID of Social Media"
// @Success      200  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /social_media/{id} [get]
func (h HttpServer) GetOneSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	in := model.SocialMedia{}
	userID := int(userData["id"].(float64))
	in.UserID = userID

	res, err := h.app.GetOneSocialMedia(int64(idInt))
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

// Create Social Media godoc
// @Summary      Created an Social Media
// @Description  Created an Social Media
// @Tags         Social Media
// @Accept       json
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /social_media [post]
func (h HttpServer) CreateSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContent(c)

	in := model.SocialMedia{}
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

	res, err := h.app.CreateSocialMedia(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Created(c, "Success Created Social Media", res)
}

// UpdateSocialMedia godoc
// @Summary      Update detail of Social Media for given id
// @Description  update detail of Social Media by id
// @Tags         Social Media
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID of Social Media"
// @Success      200  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /social_media/{id} [put]
func (h HttpServer) UpdateSocialMedia(c *gin.Context) {

	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContent(c)

	in := model.SocialMedia{}

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

	res, err := h.app.UpdateSocialMedia(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.OkWithMessage(c, "Social Media Updated: ", res)
}

// Delete Social Media godoc
// @Summary      Delete of Social Media data
// @Description  Delete of Social Media by id
// @Tags         Social Media
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID of Social Media"
// @Success      200  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /social_media/{id} [delete]
func (h HttpServer) DeleteSocialMedia(c *gin.Context) {

	userData := c.MustGet("userData").(jwt.MapClaims)

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	in := model.SocialMedia{}
	userID := int(userData["id"].(float64))
	in.UserID = userID

	err = h.app.DeleteSocialMedia(int64(idInt))
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
