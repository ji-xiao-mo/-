package controllers

import (
	"github.com/gin-gonic/gin"
)

type OrderController struct{}

type Search struct {
	Name string `json:"name"`
	Cid  int    `json:"cid"`
}

func (u OrderController) GetList(c *gin.Context) {
	// cid := c.PostForm("cid")
	// name := c.DefaultPostForm("name", "wang")
	// param := make(map[string]interface{})
	// err := c.BindJSON(&param)
	search := &Search{}
	err := c.BindJSON(&search)

	if err != nil {
		ReturnError(c, 4001, gin.H{"err": err})
		return
	}

	ReturnSuccess(c, 0, search.Name, search.Cid, 1)
}
