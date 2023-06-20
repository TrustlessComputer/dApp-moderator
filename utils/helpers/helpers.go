package helpers

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcutil"
	"github.com/ethereum/go-ethereum/params"
	"go.mongodb.org/mongo-driver/bson"
)

func ToDoc(v interface{}) (*bson.D, error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return nil, err
	}

	doc := &bson.D{}
	err = bson.Unmarshal(data, doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func GenerateSlug(key string) string {
	key = strings.ReplaceAll(key, " ", "-")
	key = strings.ReplaceAll(key, "#", "")
	key = strings.ReplaceAll(key, "@", "")
	key = strings.ReplaceAll(key, `%`, "")
	key = strings.ReplaceAll(key, `?`, "")
	key = strings.ReplaceAll(key, `(`, "")
	key = strings.ReplaceAll(key, `)`, "")
	key = strings.ReplaceAll(key, `[`, "")
	key = strings.ReplaceAll(key, `]`, "")
	key = strings.ReplaceAll(key, `{`, "")
	key = strings.ReplaceAll(key, `}`, "")
	key = strings.ReplaceAll(key, `!`, "")
	key = strings.ReplaceAll(key, `=`, "")
	//key = regexp.MustCompile(`[^a-zA-Z0-9?:-]+`).ReplaceAllString(key, "")
	key = strings.ToLower(key)
	key = ReplaceNonUTF8(key)
	return key
}

func ReplaceNonUTF8(filename string) string {
	re := regexp.MustCompile("[^a-zA-Z0-9./:]")
	return fmt.Sprintf(re.ReplaceAllString(filename, ""))
}

func JsonTransform(from interface{}, to interface{}) error {
	bytes, err := json.Marshal(from)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, to)
	if err != nil {
		return err
	}

	return nil
}

func ParseData(from []byte, to interface{}) error {
	err := json.Unmarshal(from, to)
	if err != nil {
		return err
	}

	return nil
}

func Transform(from interface{}, to interface{}) error {
	bytes, err := bson.Marshal(from)
	if err != nil {
		return err
	}

	err = bson.Unmarshal(bytes, to)
	if err != nil {
		return err
	}

	return nil
}

func GenerateMd5String(input string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(input)))
}

func MagicHash(msg, messagePrefix string) (chainhash.Hash, error) {
	if messagePrefix == "" {
		messagePrefix = "\u0018Bitcoin Signed Message:\n"
	}

	bytes := append([]byte(messagePrefix), []byte(msg)...)
	return chainhash.DoubleHashH(bytes), nil
}

func GetAddressFromPubKey(publicKey *btcec.PublicKey, compressed bool) (*btcutil.AddressPubKeyHash, error) {
	temp, err := btcutil.NewAddressPubKeyHash(btcutil.Hash160(publicKey.SerializeCompressed()), &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}
	return temp, nil
}

func PubKeyFromSignature(sig, msg string, prefix string) (pubKey *btcec.PublicKey, wasCompressed bool, err error) {
	// var decodedSig []byte
	// if decodedSig, err = base64.StdEncoding.DecodeString(sig); err != nil {
	// 	return nil, false, err
	// }

	// temp, err := MagicHash(msg, prefix)
	// if err != nil {
	// 	return nil, false, err
	// }
	// k, c, err := ecdsa.RecoverCompact(decodedSig, temp[:])
	// return k, c, err

	//TODO - implement me
	return nil, false, nil
}

func ReplaceToken(token string) string {
	token = strings.ReplaceAll(token, "Bearer", "")
	token = strings.ReplaceAll(token, "bearer", "")
	token = strings.ReplaceAll(token, " ", "")
	return token
}

func NftsOfContractPageKey(contract string) string {
	return fmt.Sprintf("contract.%s.nfts.page", contract)
}

func EtherToWei(eth *big.Float) *big.Int {
	if eth == nil {
		return big.NewInt(0.0)
	}

	truncInt, _ := eth.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))
	fracStr := strings.Split(fmt.Sprintf("%.18f", eth), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}

func ConvertWeiToBigFloat(amt *big.Int, decimals uint) *big.Float {
	if amt == nil {
		return big.NewFloat(0.0)
	}

	if amt.Cmp(big.NewInt(0)) < 0 {
		panic(errors.New("amount is small than 0"))
	}
	amtFloat := new(big.Float).SetPrec(1024).SetInt(amt)
	decimalFloat := new(big.Float).SetPrec(1024).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
	retFloat := new(big.Float).Quo(amtFloat, decimalFloat)
	return retFloat
}

func ConvertWeiToBigFloatNegative(amt *big.Int, decimals uint) *big.Float {
	if amt == nil {
		return big.NewFloat(0.0)
	}
	amtFloat := new(big.Float).SetPrec(1024).SetInt(amt)
	decimalFloat := new(big.Float).SetPrec(1024).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
	retFloat := new(big.Float).Quo(amtFloat, decimalFloat)
	return retFloat
}

func SlackHook(channel, content string) error {
	slackURL := "https://hooks.slack.com/services/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"
	// slackURL := "https://hooks.slack.com/services/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"
	go func() error {
		bodyRequest, err := json.Marshal(map[string]interface{}{
			"channel":  channel,
			"username": "tc-report",
			"text":     content,
			"icon_url": "http://www.hopabot.com/img/intro-carousel/f2.png",
		})
		if err != nil {
			return err
		}
		req, err := http.NewRequest("POST", slackURL, bytes.NewBuffer(bodyRequest))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			return errors.New(res.Status)
		}
		return nil
	}()
	return nil
}

// GetAESDecrypted decrypts given text in AES 256 CBC
func GetAESDecrypted(key, iv, encrypted string) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return nil, err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("block size cant be zero")
	}

	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, ciphertext)
	// ciphertext = PKCS5UnPadding(ciphertext)

	return ciphertext, nil
}

// // PKCS5UnPadding  pads a certain blob of data with necessary data to be used in AES block cipher
// func PKCS5UnPadding(src []byte) []byte {
// 	length := len(src)
// 	unpadding := int(src[length-1])

// 	return src[:(length - unpadding)]
// }

// GetAESEncrypted encrypts given text in AES 256 CBC
func GetAESEncrypted(key, iv, plaintext string) (string, error) {
	var plainTextBlock []byte
	length := len(plaintext)
	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}

	copy(plainTextBlock, plaintext)
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, plainTextBlock)

	str := base64.StdEncoding.EncodeToString(ciphertext)

	return str, nil
}

func EncryptAES(key []byte, plaintext string) string {
	c, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}
	out := make([]byte, len(plaintext))
	c.Encrypt(out, []byte(plaintext))
	return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string) string {
	ciphertext, _ := hex.DecodeString(ct)
	c, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}
	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)
	s := string(pt[:])
	fmt.Println("DECRYPTED:", s)
	return s
}

func GetGoogleSecretKey(name string) (string, error) {
	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name + "/versions/latest",
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", err
	}

	return string(result.Payload.Data), nil
}

func Base64Decode(base64Str string) ([]byte, error) {
	sDec, err := b64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}
	return sDec, nil
}

func Base64Encode(base64Str string) string {
	sDec := b64.StdEncoding.EncodeToString([]byte(base64Str))
	return sDec
}

func TxHashInfo(txhash string) ([]byte, *http.Header, int, error) {
	txhash = strings.ToLower(txhash)
	url := os.Getenv("TC_ENDPOINT")
	requestBody := make(map[string]interface{})
	requestBody["version"] = "2.0"
	requestBody["method"] = "eth_getTransactionByHash"
	requestBody["params"] = []string{txhash}
	return HttpRequest(url, "POST", make(map[string]string), requestBody)
}

func BnsTokenNameKey(token string) string {
	return fmt.Sprintf("bns.token.%s", token)
}

func TokenRateKey(address string) string {
	return fmt.Sprintf("token.rate.%s", strings.ToLower(address))
}

func GetExternalPrice(tokenSymbol string) (float64, error) {
	binanceAPI := os.Getenv("BINANCE_API")
	if binanceAPI == "" {
		binanceAPI = "https://api.binance.com"
	}
	binancePriceURL := fmt.Sprintf("%v/api/v3/ticker/price?symbol=", binanceAPI)
	var price struct {
		Symbol string `json:"symbol"`
		Price  string `json:"price"`
	}
	var jsonErr struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	retryTimes := 0
retry:
	retryTimes++
	if retryTimes > 2 {
		return 0, nil
	}
	tk := strings.ToUpper(tokenSymbol)
	resp, err := http.Get(binancePriceURL + tk + "USDT")
	if err != nil {
		log.Println(err)
		goto retry
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &price)
	if err != nil {
		err = json.Unmarshal(body, &jsonErr)
		if err != nil {
			log.Println(err)
			goto retry
		}
	}
	resp.Body.Close()
	value, err := strconv.ParseFloat(price.Price, 32)
	if err != nil {
		log.Println("getExternalPrice", tokenSymbol, err)
		return 0, nil
	}
	return value, nil
}

func GetValue(amount string, decimal float64) float64 {
	amountBig := new(big.Float)
	amountBig.SetString(amount)

	pow10 := math.Pow10(int(decimal))
	pow10Big := big.NewFloat(pow10)

	result := amountBig.Quo(amountBig, pow10Big) //divide
	amountInt, _ := result.Float64()
	return amountInt
}

func ConvertAmountString(amount float64) string {

	result := ConvertAmount(amount)
	amountInt, _ := result.Int64()
	return fmt.Sprintf("%d", amountInt)
}

func ConvertAmount(amount float64) *big.Float {
	decimal := 18
	amountBig := big.NewFloat(amount)

	pow10 := math.Pow10(int(decimal))
	pow10Big := big.NewFloat(pow10)

	result := amountBig.Mul(amountBig, pow10Big) //divide
	return result
}

func InArray(text string, arrayText []string) bool {
	for _, item := range arrayText {
		if strings.ToLower(item) == strings.ToLower(text) {
			return true
		}
	}

	return false
}

func ParseUintToUnixTime(number uint64) *time.Time {
	t := time.Unix(int64(number), 0)
	return &t
}
