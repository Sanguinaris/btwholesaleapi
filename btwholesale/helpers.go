package btwholesale

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"encoding/json"
	"log"
	"math/rand"
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

func deDaDiDo(data []byte) BTWholeBoi {
	var boi BTWholeBoi
	json.Unmarshal(data, &boi)
	return boi
}
