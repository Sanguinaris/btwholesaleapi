package main

import (
	"math/rand"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	godotenv.Load()

	//boi := btwholesale.BTBroadBandChecker()
	//log.Println(boi)
}
