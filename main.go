package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/Sanguinaris/btwholesaleapi/btwholesale"
	"github.com/joho/godotenv"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	godotenv.Load()

	boi := btwholesale.BTBroadBandChecker()
	log.Println(boi)
}
