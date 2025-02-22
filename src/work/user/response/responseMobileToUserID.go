package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseMobileToUserID struct {
	*response.ResponseWork

	UserID string `json:"userid"`
}

type ResponseConvertToUserID struct {
	*response.ResponseWork

	UserID string `json:"userid"`
}
