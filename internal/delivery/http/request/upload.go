package request

import "errors"

type CreateMultipartUploadRequest struct {
	FileName string `json:"fileName"`
	Group    string `json:"group"`
}

type CompleteMultipartUploadRequest struct {
	WalletAddress string `json:"walletAddress"`
}

func (g CreateMultipartUploadRequest) SelfValidate() error {
	if g.FileName == "" {
		return errors.New("fileName is required")
	}

	if g.Group == "" {
		return errors.New("group should not be empty")
	}

	return nil
}

type CompressFileSize struct {
	FileContent string `json:"file_content"`
}
