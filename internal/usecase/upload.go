package usecase

import (
	"dapp-moderator/utils/googlecloud"
	"fmt"
	"mime/multipart"
	"os"
)

func (u *Usecase) UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	gf := googlecloud.GcsFile{
		FileHeader: fileHeader,
	}
	uploaded, err := u.Storage.FileUploadToBucket(gf)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", os.Getenv("GCS_DOMAIN"), uploaded.Name), nil
}
