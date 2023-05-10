package services

import (
	"fmt"
	"go-minimarket/models"
	"go-minimarket/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var DbConn *gorm.DB

func GeneratedSignatureHandler(c *gin.Context) {
	/*
		jika mengambil dari json variable
		secretKey := config.Config.Credentials.SecretKey
		clientKey := config.Config.Credentials.ClientKey
	*/
	//get data request
	//var dataSignature models.ParamSignature
	//c.ShouldBindJSON(&dataSignature)

	xpartner := c.Request.Header.Get("X-PARTNER")
	secretKey := c.Request.Header.Get("secretKey")
	clientKey := c.Request.Header.Get("clientKey")

	if xpartner == "" || secretKey == "" || clientKey == "" {
		c.JSON(utils.ResponseError(400, "Error", "X-PARTNER / clientKey / secretKey is mandatory!"))
		/*
			utils.Logger.Error(
				utils.LogRequest("400",c.Request.Method,c.Request.RequestURI,c.)
			)
			jika memakai log wolverine
			utils.LogResponse("241", "GET", "/blablaa", "{fwegrwgw}")
		*/
		return
	}

	var stringQuery string = fmt.Sprintf(`SELECT * FROM client WHERE client_name = %s`, "'"+xpartner+"'")
	var DataClient models.ClientData

	respDB := DbConn.Raw(stringQuery).Scan(&DataClient)

	if respDB.RowsAffected > 0 {

		secretKeyDB := DataClient.SecretKey
		clientKeyDB := DataClient.ClientKey

		if secretKey != secretKeyDB {
			c.JSON(utils.ResponseError(403, "Error", "Credentials - SecretKey invalid!"))
			return
		}
		if clientKey != clientKeyDB {
			c.JSON(utils.ResponseError(404, "Error", "Credentials - ClientKey invalid!"))
			return
		}

		signature := utils.CreateSignature(secretKeyDB, clientKeyDB)

		DataSign := map[string]interface{}{
			"clientKey": clientKeyDB,
			"secretKey": secretKeyDB,
			"Signature": signature,
		}
		c.JSON(http.StatusOK, models.SignatureResponse{
			ResponseCode:          "200",
			ResponseMessage:       "Success",
			ResponseMessageDetail: "Sukses generated signature",
			Data:                  DataSign,
		})

	} else if respDB.RowsAffected == 0 {

		c.JSON(utils.ResponseError(402, "Error", "X-PARTNER not found!"))
		return

	}

}
