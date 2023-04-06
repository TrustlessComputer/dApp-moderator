package helpers

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

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