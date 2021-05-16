package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/address/:postCode", lookup_postcode)
	r.POST("/lookup/full", get_data_via_full_info)
	r.GET("/lookup/uprn/:uprn", get_data_via_uprn)
	r.GET("/lookup/telno/:telno", get_data_via_telno)
	r.GET("/lookup/alid/:alid", get_data_via_alid)
	r.Run("127.0.0.1:42069")
}
