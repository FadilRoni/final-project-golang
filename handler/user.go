package handler

import (
	"final-project/helper"
	"final-project/model"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

// CreateUser godoc
// @Summary      Created an Account
// @Description  Created an Account
// @Tags         Users
// @Accept       json
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /users/register [post]
func (h HttpServer) Register(c *gin.Context) {

	contentType := helper.GetContent(c)
	in := model.User{}

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

	res, err := h.app.Register(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Created(c, "Berhasil Melakukan Restrasi Akun", res)
}

// LoginUser godoc
// @Summary      Login an Account
// @Description  Login an Account
// @Tags         Users
// @Accept       json
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /users/login [post]
func (h HttpServer) Login(c *gin.Context) {

	contentType := helper.GetContent(c)
	in := model.User{}
	password := ""

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

	res, err := h.app.Login(in)
	if err != nil {
		helper.BadRequest(c, "Wrong Email")
		return
	}

	password = in.Password

	comparePass := helper.ComparePass([]byte(res.Password), []byte(password))
	if !comparePass {
		helper.Unauthorized(c, "Invalid Email/Password")
		return
	}

	helper.Ok(c, helper.GenerateToken(res.ID, res.Email))

	// log.Fatalln(res)
	// os.Exit(1)

	// helper.OkWithMessage(c, "Login Berhasil", res)
}
