package services

import (
	"fmt"
	"go-minimarket/models"
	"go-minimarket/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var dataAccount models.AccountData

	xpartner := c.Request.Header.Get("X-PARTNER")
	xsignature := c.Request.Header.Get("X-SIGNATURE")

	if xpartner == "" {
		c.JSON(utils.ResponseError(400, "Error", "X-PARTNER is mandatory!"))
		return
	}
	if xsignature == "" {
		c.JSON(utils.ResponseError(400, "Error", "X-SiGNATURE is mandatory!"))
		return
	}

	var QueryClient string = fmt.Sprintf(`SELECT * FROM client WHERE client_name = %s`, "'"+xpartner+"'")
	var DataClient models.ClientData

	respDBClient := DbConn.Raw(QueryClient).Scan(&DataClient)

	if respDBClient.RowsAffected > 0 {

		secretKeyDB := DataClient.SecretKey
		clientKeyDB := DataClient.ClientKey

		isValid := utils.ValidateSignature(secretKeyDB, clientKeyDB, xsignature)

		if !isValid {
			c.JSON(utils.ResponseError(403, "Error", "Signature invalid!"))
			return
		}

	} else {
		c.JSON(utils.ResponseError(402, "Error", "X-PARTNER not found!"))
		return
	}

	//Validasi request mandatory
	c.ShouldBindJSON(&dataAccount)

	if dataAccount.Username == "" {
		c.JSON(utils.ResponseError(400, "Error", "Username is mandatory!"))
		return
	}

	var searchAccount string = fmt.Sprintf(`SELECT * FROM account WHERE username = %s`, "'"+dataAccount.Username+"'")
	var DataAccount models.AccountData

	Username := DataAccount.Username

	respDBAccount := DbConn.Raw(searchAccount).Scan(&DataAccount)

	if respDBAccount.RowsAffected > 0 {
		c.JSON(utils.ResponseError(400, "Error", "Username "+Username+" already exist!"))
		return
	}

	if dataAccount.Password == "" {
		c.JSON(utils.ResponseError(400, "Error", "Password is mandatory!"))
		return
	}
	if len(dataAccount.Password) != 8 {
		c.JSON(utils.ResponseError(400, "Error", "Panjang password adalah 8!"))
		return
	}
	specialChars := "~!@#$%^&*()_+-={}[]\\|<,>.?/\"';:`"
	for _, char := range specialChars {
		if strings.Contains(dataAccount.Password, string(char)) {
			c.JSON(utils.ResponseError(400, "Error", "Password tidak boleh mengandung karakter khusus!"))
			return
		}
	}
	if dataAccount.Nama_lengkap == "" {
		c.JSON(utils.ResponseError(400, "Error", "Nama_lengkap is mandatory!"))
		return
	}
	if dataAccount.Telp == "" {
		c.JSON(utils.ResponseError(400, "Error", "Telp is mandatory!"))
		return
	}
	if len(dataAccount.Telp) < 10 {
		c.JSON(utils.ResponseError(400, "Error", "Min length kolom Telp adalah 10!"))
		return
	}

	var insertQuery string = fmt.Sprintf(`INSERT INTO account (username,password,nama_lengkap,email,telp) VALUES(%s,%s,%s,%s,%s)`,
		"'"+dataAccount.Username+"'", "'"+dataAccount.Password+"'", "'"+dataAccount.Nama_lengkap+"'", "'"+dataAccount.Email+"'", "'"+dataAccount.Telp+"'")
	var DataAccountInsert []models.AccountData

	respDBInsert := DbConn.Raw(insertQuery).Scan(&DataAccountInsert)
	if respDBInsert.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": respDBInsert.Error.Error()})
		return
	}

	DataNewAccount := map[string]interface{}{
		"username":     dataAccount.Username,
		"password":     dataAccount.Password,
		"nama_lengkap": dataAccount.Nama_lengkap,
		"email":        dataAccount.Email,
		"telp":         dataAccount.Telp,
	}

	c.JSON(http.StatusOK, models.RegisterResponse{
		ResponseCode:          "200",
		ResponseMessage:       "Success",
		ResponseMessageDetail: "Sukses insert new account",
		DataAccount:           DataNewAccount,
	})

}
