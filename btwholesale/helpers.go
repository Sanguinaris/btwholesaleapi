package btwholesale

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/andreburgaud/crypt2go/padding"
	"github.com/martinlindhe/base36"
)

// this is garbage code, but so is theirs
func createCaptcha() string {
	str := "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	buff := bytes.NewBufferString("")
	buff.Grow(5)
	for i := 0; i < 6; i++ {
		buff.WriteByte(str[int(rand.Float64()*float64(len(str)))])
	}
	return buff.String()
}

func prependZeroForThings(d string) string {
	if len(d) == 1 {
		return "0" + d
	}
	return d
}

func seenBefore(what byte, idx int, arr *string) bool {
	arrKek := *arr
	for i := range arrKek {
		if i >= idx {
			return false
		}
		if what == arrKek[i] {
			return true
		}
	}
	return false
}

func generateUniqueString() string {
	time := time.Now()

	day := prependZeroForThings(strconv.Itoa(time.Day()))
	month := prependZeroForThings(strconv.Itoa(int(time.Month()) + 1))
	year := prependZeroForThings(strconv.Itoa(time.Year()))
	hours := prependZeroForThings(strconv.Itoa(time.Hour()))
	minutes := prependZeroForThings(strconv.Itoa(time.Minute()))
	seconds := prependZeroForThings(strconv.Itoa(time.Second()))

	concat := day + month + year + hours + minutes + seconds

	lol := "0." + strings.ToLower(base36.Encode(uint64(rand.Intn(99999999999999999-1000000000000000)+1000000000000000)))
	uniqueBois := []byte{}
	for i := range lol {
		if !seenBefore(lol[i], i, &lol) {
			uniqueBois = append(uniqueBois, lol[i])
		}
	}

	return string(uniqueBois[2:8]) + concat
}

func encryptUniqueString(unique string) string {
	key := []byte("GhB$skWWTrE27e=(")
	pt := []byte(unique)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBEncrypter(block)
	padder := padding.NewPkcs7Padding(mode.BlockSize())
	pt, err = padder.Pad(pt)
	if err != nil {
		panic(err.Error())
	}
	ct := make([]byte, len(pt))
	mode.CryptBlocks(ct, pt)

	return "bbac" + base64.StdEncoding.EncodeToString(ct)
}

func decrypt(ct, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBDecrypter(block)
	pt := make([]byte, len(ct))
	mode.CryptBlocks(pt, ct)
	padder := padding.NewPkcs7Padding(mode.BlockSize())
	pt, err = padder.Unpad(pt) // unpad plaintext after decryption
	if err != nil {
		panic(err.Error())
	}
	return pt
}

func decryptBtWhole(data string) []byte {
	key := []byte("GhB$skWWTrE27e=(")
	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	clear := decrypt(ciphertext, key)
	return clear
}

func request_setup() (client *http.Client, req *http.Request) {
	client = &http.Client{}

	req, err := http.NewRequest("GET", "https://api.wholesale.bt.com/bt-wholesale/v1/broadband/coverage/anonymous", nil)
	if err != nil {
		panic(err.Error())
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

	req.Header.Add("Authorization", "Basic MjFuZ0dHNks0bk9NY3gybEpzWkpWUFZZQXlZdTJIV3Y6Vjd1M1FQT2xITXF2TFRKbA==")
	uniqueThing := generateUniqueString()
	trackCap := encryptUniqueString(uniqueThing + createCaptcha())

	req.Header.Add("APIGW-Tracking-Header", uniqueThing)
	req.Header.Add("AuthTrackCap", trackCap)
	req.Header.Add("requestTime", time.Now().Format("2006-01-02 15:04:05"))
	return
}

func request_breakdown(client *http.Client, req *http.Request) string {
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
