package responses

import "winterchen.com/my-site-go/src/models"

type ContentResponse struct {
	Data  []models.Content `json:"data"`
	total int              `json:"total"`
}
