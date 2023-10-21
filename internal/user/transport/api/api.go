package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-webapp/internal/user/entity"
	"net/http"
)

type Service interface {
	CreateUser(c context.Context, req *entity.CreateUserReq) (*entity.CreateUserRes, error)
	Login(c context.Context, req *entity.LoginUserReq) (*entity.LoginUserRes, error)
}

type Api struct {
	service Service
}

func NewHandler(s Service) *Api {
	return &Api{
		service: s,
	}
}

func (api *Api) CreateUser(c *gin.Context) {
	var u entity.CreateUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := api.service.CreateUser(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (api *Api) Login(c *gin.Context) {
	var user entity.LoginUserReq
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := api.service.Login(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("jwt", u.AccessToken, 60*60*24, "/", "localhost", false, true)
	c.JSON(http.StatusOK, u)
}

func (api *Api) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}
