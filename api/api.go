package api

import (
	"net/http"

	"github.com/Sanguinaris/btwholesaleapi/btwholesale"
	"github.com/gin-gonic/gin"
)

func lookup_postcode(c *gin.Context) {
	town := c.Query("town")
	street := c.Query("street")
	buildingNumber := c.Query("buildingNumber")
	buildingName := c.Query("buildingName")
	postCode := c.Param("postCode")

	lookup, boi := btwholesale.LookupAddress(postCode, town, street, buildingNumber, buildingName)

	if lookup != nil {
		c.JSON(http.StatusOK, lookup)
	} else if boi != nil {
		c.JSON(http.StatusOK, boi)
	} else {
		c.JSON(http.StatusInternalServerError, nil)
	}
}

func get_data_via_full_info(c *gin.Context) {
	var body FullAddress
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lookup := btwholesale.BTBroadBandChecker(body.PostCode, body.Town, body.Refnumber, body.Street, body.BuildingNumber, body.BuildingName, body.DistrictId)
	c.JSON(http.StatusOK, lookup)
}

func get_data_via_uprn(c *gin.Context) {
	uprn := c.Param("uprn")
	lookup := btwholesale.LookupUPRN(uprn)
	c.JSON(http.StatusOK, lookup)
}

func get_data_via_telno(c *gin.Context) {
	telno := c.Param("telno")
	lookup := btwholesale.LookupTelephone(telno)
	c.JSON(http.StatusOK, lookup)
}

func get_data_via_alid(c *gin.Context) {
	alid := c.Param("alid")
	lookup := btwholesale.LookupAlId(alid)
	c.JSON(http.StatusOK, lookup)
}
