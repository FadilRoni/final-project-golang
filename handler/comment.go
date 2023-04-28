package handler

import (
	"final-project/helper"
	"final-project/model"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetAllComment godoc
// @Summary      List of Comment From One Photo
// @Description  List of Comment From One Photo
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Success      200  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /comment/{idPhoto} [get]
func (h HttpServer) GetAllComment(c *gin.Context) {

	userData := c.MustGet("userData").(jwt.MapClaims)
	in := model.Comment{}

	userID := int(userData["id"].(float64))

	in.UserID = userID

	res, err := h.app.GetAllComment(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.OkWithMessage(c, "Data Comment From One Photo", res)
}

// GetOneComment godoc
// @Summary      Get One Comment From One Photo for given id
// @Description  Get One Comment From One Photo for given id
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID of Comment"
// @Success      200  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /comment/{idPhoto}/{idComment} [get]
func (h HttpServer) GetOneComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)

	idPhoto, _ := strconv.Atoi(c.Param("idPhoto"))
	idComment, _ := strconv.Atoi(c.Param("idComment"))

	in := model.Comment{}
	userID := int(userData["id"].(float64))
	PhotoID := idPhoto
	CommentID := uint(idComment)

	in.UserID = userID
	in.PhotoID = PhotoID
	in.ID = CommentID

	res, err := h.app.GetOneComment(in)
	if err != nil {
		if err.Error() == helper.ErrNotFound {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.OkWithMessage(c, "Data Comment", res)
}

// Create Comment godoc
// @Summary      Created Comment
// @Description  Created Comment
// @Tags         Comment
// @Accept       json
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /comment/{idPhoto} [post]
func (h HttpServer) CreateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContent(c)

	idPhoto, _ := strconv.Atoi(c.Param("idPhoto"))

	in := model.Comment{}
	userID := int(userData["id"].(float64))
	PhotoID := idPhoto

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
	in.PhotoID = PhotoID

	res, err := h.app.CreateComment(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Created(c, "Success Created Comment", res)
}

// UpdateComment godoc
// @Summary      Update Comment From One Photo for given id
// @Description  Update Comment From One Photo by id
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID of Comment"
// @Success      200  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /comment/{idPhoto}/{idComment} [put]
func (h HttpServer) UpdateComment(c *gin.Context) {

	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContent(c)

	idPhoto, _ := strconv.Atoi(c.Param("idPhoto"))
	idComment, _ := strconv.Atoi(c.Param("idComment"))

	in := model.Comment{}
	userID := int(userData["id"].(float64))
	PhotoID := idPhoto
	CommentID := uint(idComment)

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
	in.PhotoID = PhotoID
	in.ID = CommentID

	res, err := h.app.UpdateComment(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.OkWithMessage(c, "Comment Updated: ", res)
}

// DeleteComment godoc
// @Summary      Delete Comment From One Photo fot given id
// @Description  Delete Comment From One Photo by id
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID of Comment"
// @Success      200  {object}  helper.Response
// @Failure      404  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /comment/{idPhoto}/{idComment} [delete]
func (h HttpServer) DeleteComment(c *gin.Context) {

	userData := c.MustGet("userData").(jwt.MapClaims)

	idPhoto, _ := strconv.Atoi(c.Param("idPhoto"))
	idComment, _ := strconv.Atoi(c.Param("idComment"))

	in := model.Comment{}
	userID := int(userData["id"].(float64))
	PhotoID := idPhoto
	CommentID := uint(idComment)

	in.UserID = userID
	in.PhotoID = PhotoID
	in.ID = CommentID

	err := h.app.DeleteComment(in)
	if err != nil {
		if err.Error() == helper.ErrNotFound {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.OkWithMessage(c, "Comment Has Been Deleted", nil)
}
