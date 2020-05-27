package controllers

import (
	"github.com/gin-gonic/gin"
	"gin-example/models"
	"gin-example/repository"
)

func SampleGet(c *gin.Context) {
	data, err := repository.SampleGet()

	if err != nil {
		c.Error(err)
		c.Set("statusCode",500)
		return
	}
	c.Set("statusCode",200)
	c.Set("data",data)
}

func SamplePost(c *gin.Context) {
	sampleInsert := models.SampleInsert{}
	if err := c.ShouldBindJSON(&sampleInsert); err == nil {
		err := repository.SamplePost(sampleInsert)
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

func SamplePatch(c *gin.Context) {
	samplePatch := models.SamplePatch{}
	if err := c.ShouldBindJSON(&samplePatch); err == nil {
		err := repository.SamplePatch(samplePatch)
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

func SampleDelete(c *gin.Context) {
	sampleDelete := models.SampleDelete{}
	if err := c.ShouldBindJSON(&sampleDelete); err == nil {
		err := repository.SampleDelete(sampleDelete)
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