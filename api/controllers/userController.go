package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/miguelanselmo/my-api-demo/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type params struct {
	PageID   int8 `form:"page_id" binding:"required,min=1"`
	PageSize int8 `form:"page_size" binding:"required,min=5,max=10"`
}

//https://github.com/swaggo/swag

// ShowAccount godoc
// @Summary      get accounts
// @Description  get all account of user
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  models.User
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /accounts/{id} [get]
func (ctrl *Controller) GetUserAll(c *gin.Context) {
	//Get all users
	var params params
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("PageID: ", params.PageID)
	log.Println("PageSize: ", params.PageSize)
	users, err := ctrl.uc.GetUsersAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (ctrl *Controller) GetUserById(c *gin.Context) {
	//Get a user by id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := ctrl.uc.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ctrl *Controller) CreateUser(c *gin.Context) {
	//Create a user
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userCreate, err := ctrl.uc.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, userCreate)
}

func (ctrl *Controller) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Id = id
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userUpdate, err := ctrl.uc.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userUpdate)
}
