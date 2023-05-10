package services

import (
	"fmt"
	"go-minimarket/models"
	"go-minimarket/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListBarangHandler(c *gin.Context) {

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

	limitStart := c.Query("limitStart")
	limitEnd := c.Query("limitEnd")
	var stringQuery string = fmt.Sprintf(`SELECT * FROM barang order by nama,kode ASC LIMIT %s,%s`, ""+limitStart+"", ""+limitEnd+"")
	var Listbarang []models.Barang

	respDB := DbConn.Raw(stringQuery).Scan(&Listbarang)
	fmt.Println(respDB.Error)

	if respDB.Error != nil {
		c.JSON(utils.ResponseError(401, "Error", "Data barang not found!"))
		return
	}

	c.JSON(http.StatusOK, models.ListBarangResponse{
		ResponseCode:          "200",
		ResponseMessage:       "Success",
		ResponseMessageDetail: "Sukses get list barang",
		ListData:              Listbarang,
	})
}

func SearchBarangHandler(c *gin.Context) {

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

	var searchbarang models.ParamListBarang
	//Validasi request mandatory
	c.ShouldBindJSON(&searchbarang)

	//validasi kolom empty
	if searchbarang.Nama_barang == "" {
		c.JSON(utils.ResponseError(400, "Error", "Nama barang is mandatory!"))
		return
	}
	if searchbarang.Kategori == "" {
		c.JSON(utils.ResponseError(400, "Error", "Kategori is mandatory!"))
		return
	}

	//validasi length minimal
	if len(searchbarang.Nama_barang) < 4 || len(searchbarang.Kategori) < 4 {
		c.JSON(utils.ResponseError(402, "Error", "Minimal panjang karakter = 4!"))
		return
	}

	var stringQuery string = fmt.Sprintf(`SELECT * FROM barang WHERE nama like %s AND kategori like %s`,
		"'%"+searchbarang.Nama_barang+"%'", "'%"+searchbarang.Kategori+"%'")
	var Listbarang []models.Barang

	respDB := DbConn.Raw(stringQuery).Scan(&Listbarang)
	fmt.Println(respDB.Error)

	if respDB.RowsAffected == 0 {
		c.JSON(utils.ResponseError(401, "Error", "Data barang not found!"))
		return
	}

	c.JSON(http.StatusOK, models.ListBarangResponse{
		ResponseCode:          "200",
		ResponseMessage:       "Success",
		ResponseMessageDetail: "Sukses get list barang",
		ListData:              Listbarang,
	})
}
