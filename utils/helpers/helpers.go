package helpers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
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

func ReplaceNonUTF8(filename string) string  {
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