package response

import "time"

type UploadResponse struct {
	FileName  string    `json:"file_name"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
