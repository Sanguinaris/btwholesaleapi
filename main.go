package main

import (
	"math/rand"
	"time"

	"github.com/Sanguinaris/btwholesaleapi/api"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	api.Start()
	//boi := btwholesale.BTBroadBandChecker()
	//log.Println(boi)
}
