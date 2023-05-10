package services

import (
	"encoding/json"
	"go-minimarket/models"
	"go-minimarket/utils"

	"github.com/gin-gonic/gin"
)

func SendWAHandler(c *gin.Context) {
	var dataSendWA models.ParamSendWA
	c.ShouldBindJSON(&dataSendWA)

	Auth := c.Request.Header.Get("Authorization")
	header := map[string]string{
		"Authorization": Auth,
		"Content-Type":  "application/json",
	}

	params := map[string]string{}

	body := map[string]string{
		"target":  dataSendWA.Target,
		"message": dataSendWA.Message,
	}

	url := "https://api.fonnte.com/send"

	resp, err := utils.HttpRequest("POST", url, header, params, body, 300)
	_ = err

	var responseBody map[string]interface{}
	err = json.Unmarshal(resp.Body(), &responseBody)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, responseBody)

}
