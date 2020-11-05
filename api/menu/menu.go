package menu

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-go/dto"
	"simple-go/requests"
	"simple-go/services/menu"
)

func CreateMenu(c *gin.Context) {

	request := requests.MenuRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service menu.MenuContract = &menu.MenuService{}

	u := dto.MenuDto{
		Username: request.UserName,
		Password: request.Password,
	}
	result := service.CreateMenu(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetMenuList(c *gin.Context) {
	request := requests.MenuRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service menu.MenuContract = &menu.MenuService{}

	result := service.GetMenuList("u")
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func UpdateMenu(c *gin.Context) {

	request := requests.MenuRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service menu.MenuContract = &menu.MenuService{}
	u := dto.MenuDto{
		ID:       request.ID,
		Username: request.UserName,
		Password: request.Password,
	}
	result := service.UpdateMenu(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func DeleteMenuById(c *gin.Context) {

	request := requests.MenuRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service menu.MenuContract = &menu.MenuService{}

	result := service.DeletMenuById(request.ID)
	c.JSON(http.StatusOK, gin.H{"data": result})
}
