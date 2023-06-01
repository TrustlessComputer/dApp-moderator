package helpers

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

func CreateFile(fileName string, data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile(fileName string) ([]byte, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return content, nil
}

type CSVLine struct {
	TxHash string
}

func ReadCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	resp := [][]string{}

	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		// do something with read line
		resp = append(resp, rec)
	}

	return resp, nil
}
