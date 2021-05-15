package btwholesale

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func BTBroadBandChecker() BTWholeBoi {
	return deDaDiDo(decryptBtWhole(getData()))
}

func getData() string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.wholesale.bt.com/bt-wholesale/v1/broadband/coverage/anonymous", nil)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "en-GB,en;q=0.9,en-US;q=0.8")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"90\", \"Google Chrome\";v=\"90\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "cross-site")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
	req.Header.Add("Origin", "https://www.broadbandchecker.btwholesale.com")
	req.Header.Add("Referer", "https://www.broadbandchecker.btwholesale.com/")

	req.Header.Add("requestType", "request2")
	req.Header.Add("Authorization", "Basic MjFuZ0dHNks0bk9NY3gybEpzWkpWUFZZQXlZdTJIV3Y6Vjd1M1FQT2xITXF2TFRKbA==")
	uniqueThing := generateUniqueString()
	trackCap := encryptUniqueString(uniqueThing + createCaptcha())

	req.Header.Add("APIGW-Tracking-Header", uniqueThing)
	req.Header.Add("AuthTrackCap", trackCap)
	req.Header.Add("requestTime", time.Now().Format("2006-01-02 15:04:05"))

	req.Header.Add("postCode", os.Getenv("postCode"))
	req.Header.Add("town", os.Getenv("town"))
	req.Header.Add("refNumber", os.Getenv("refNumber"))
	req.Header.Add("street", os.Getenv("street"))
	req.Header.Add("buildingNumber", os.Getenv("buildingNumber"))
	req.Header.Add("buildingName", os.Getenv("buildingName"))
	req.Header.Add("districtId", os.Getenv("districtId"))

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return ""
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(string(body[1 : len(body)-1]))
}
