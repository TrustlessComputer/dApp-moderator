package helpers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"regexp"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcutil"
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
	ciphertext = PKCS5UnPadding(ciphertext)

	return ciphertext, nil
}

// PKCS5UnPadding  pads a certain blob of data with necessary data to be used in AES block cipher
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}

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
