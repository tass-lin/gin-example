package controllers

import (
	"gin-example/models"
	"gin-example/repository"
	"github.com/gin-gonic/gin"
)

func MysqlGet(c *gin.Context) {
	data, err := repository.MysqlGet()

	if err != nil {
		c.Error(err)
		c.Set("statusCode",500)
		return
	}
	c.Set("statusCode",200)
	c.Set("data",data)
}

func MysqlPost(c *gin.Context) {
	mysqlPost := models.Test{}
	if err := c.ShouldBindJSON(&mysqlPost); err == nil {
		err := repository.MysqlPost(mysqlPost)
		if err != nil {
			c.Error(err)
			c.Set("statusCode",500)
			return
		}
		c.Set("statusCode",200)
		c.Set("data",nil)
	}else {
		c.Error(err)
		c.Set("statusCode", 400)
		return
	}

}


func MysqlPatch(c *gin.Context) {
	mysqlPatch := models.Test{}
	if err := c.ShouldBindJSON(&mysqlPatch); err == nil {
		err := repository.MysqlUpdate(mysqlPatch)
		if err != nil {
			c.Error(err)
			c.Set("statusCode",500)
			return
		}
		c.Set("statusCode",200)
		c.Set("data",nil)
	}else {
		c.Error(err)
		c.Set("statusCode", 400)
		return
	}

}


func MysqlDelete(c *gin.Context) {
	mysqlDelete := models.Test{}
	if err := c.ShouldBindJSON(&mysqlDelete); err == nil {
		err := repository.MysqlDelete(mysqlDelete)
		if err != nil {
			c.Error(err)
			c.Set("statusCode",500)
			return
		}
		c.Set("statusCode",200)
		c.Set("data",nil)
	}else {
		c.Error(err)
		c.Set("statusCode", 400)
		return
	}

}



