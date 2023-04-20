package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type AbiData struct {
	Name     string      `json:"contractName"`
	ByteCode string      `json:"bytecode"`
	Abi      interface{} `json:"abi"`
}

type UploadedFile struct {
	Name    string `json:"name"`
	Content []byte `json:"content"`
}

const (
	PATH_PROJECT_TEMPLATE = "/app/tools/compile-contract/base-project"
	PATH_RUNTIME_PROJECT  = "/app/project/"

	dirPerm = 0777
)

func (u *Usecase) CompileContract(r *http.Request) ([]AbiData, error) {

	var abiData []AbiData

	err := r.ParseMultipartForm(1048576) // 1 MB
	if err != nil {
		fmt.Println("ParseMultipartForm: ", err)
		return abiData, err
	}

	// Get the files from the form data
	files := make([]UploadedFile, 0)
	for _, fheader := range r.MultipartForm.File["files"] {
		file, err := fheader.Open()
		if err != nil {
			fmt.Println("MultipartForm: ", err)
			return abiData, err
		}
		defer file.Close()

		content, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println("ReadAll: ", err)
			return abiData, err
		}

		// check: name
		// Check that the file name ends with .sol
		if !strings.HasSuffix(fheader.Filename, ".sol") {
			err := errors.New("Invalid file extension, file must be .sol")
			return abiData, err
		}

		files = append(files, UploadedFile{
			Name:    fheader.Filename,
			Content: content,
		})
	}
	if len(files) > 0 {

		// clone ./tools/compile-contract/base-project to: /app/user-contracts/<uuid>/
		pathSaveTo := PATH_RUNTIME_PROJECT + uuid.New().String()

		if _, err := os.Stat(pathSaveTo); os.IsNotExist(err) {
			fmt.Println("pathSaveTo not exits")
			if err := os.MkdirAll(pathSaveTo, os.FileMode(dirPerm)); err != nil {
				fmt.Println("MkdirAll err", err)
				return abiData, err
			}
			fmt.Println("pathSaveTo created: ", pathSaveTo)

		}

		if _, err := os.Stat(PATH_PROJECT_TEMPLATE); os.IsNotExist(err) {
			fmt.Println("PATH_PROJECT_TEMPLATE not exits")

			fmt.Println("PATH_PROJECT_TEMPLATE not exits err", err)
			return abiData, err

		}

		cmd := exec.Command("cp", "-R", PATH_PROJECT_TEMPLATE, pathSaveTo)

		fmt.Println("Command.cp: ", PATH_PROJECT_TEMPLATE, pathSaveTo)

		err = cmd.Run()
		if err != nil {
			fmt.Println("Command.cp: ", err)
			return abiData, err
		}
		if _, err := os.Stat(pathSaveTo); os.IsNotExist(err) {
			fmt.Println("pathSaveTo IsNotExist: ", err)
			return abiData, err
		}
		// upload file:
		folderToCompile := pathSaveTo + "/base-project/"
		for _, file := range files {
			err := ioutil.WriteFile(folderToCompile+"contracts/"+file.Name, file.Content, 0644)
			if err != nil {
				fmt.Println("WriteFile: ", err)
				return abiData, err
			}
		}

		// dir, err := os.Getwd()
		// if err != nil {
		// 	fmt.Println("Error getting current directory:", err)
		// 	return abiData, err
		// }

		// fmt.Println("current dir: ", dir)

		// // change the working directory to the current directory
		// err = os.Chdir(folderToCompile)
		// if err != nil {
		// 	fmt.Println("Error changing directory:", err)
		// 	return abiData, err
		// }

		// dir, err = os.Getwd()
		// if err != nil {
		// 	fmt.Println("Error getting current directory:", err)
		// 	return abiData, err
		// }

		// fmt.Println("current dir 2: ", dir)

		// run compile:
		// compileCmd := fmt.Sprintf("cd %s && npx hardhat compile", folderToCompile)

		// compileCmd := fmt.Sprintf("npx hardhat compile")

		// fmt.Println("Command.compile: ", compileCmd)

		cmd = exec.Command("npx", "hardhat", "compile")
		cmd.Dir = folderToCompile

		// err = cmd.Run()

		output, err := cmd.CombinedOutput()

		fmt.Println("Output:", string(output))

		if err != nil {
			fmt.Println("Command.compile: ", err)
			return abiData, err
		}
		// find list json in contract folder:
		folderToGetAbi := folderToCompile + "artifacts/contracts/"

		if _, err := os.Stat(folderToGetAbi); os.IsNotExist(err) {
			fmt.Println("folderToGetAbi IsNotExist: ", err)
			return abiData, err
		}

		for _, file := range files {
			folderToGetAbiChild := folderToGetAbi + file.Name
			if _, err := os.Stat(folderToGetAbiChild); os.IsNotExist(err) {
				fmt.Println("folderToGetAbiChild IsNotExist: ", err)
				return abiData, err
			}
			// list file json:
			directory, err := os.Open(folderToGetAbiChild)
			if err != nil {
				return abiData, err
			}
			defer directory.Close()

			objects, err := directory.Readdir(-1)
			if err != nil {
				return abiData, err
			}

			for _, obj := range objects {

				if strings.Contains(obj.Name(), "dbg") {
					continue
				}
				abiFile := filepath.Join(folderToGetAbiChild+"/", obj.Name())

				fmt.Println("file to get json abi: ", abiFile)

				if !obj.IsDir() && strings.HasSuffix(obj.Name(), ".json") {
					//read file
					fileContent, err := ioutil.ReadFile(abiFile)
					if err != nil {
						fmt.Println("Error reading file:", err)
						continue
					}

					var abiInfo AbiData
					err = json.Unmarshal(fileContent, &abiInfo)
					if err != nil {
						fmt.Println("Error decoding JSON:", err)
						continue
					}
					abiData = append(abiData, abiInfo)
				}
			}

		}

	}

	return abiData, nil
}
