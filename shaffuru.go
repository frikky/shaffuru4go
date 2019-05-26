package shaffuru

import (
	"fmt"
	"github.com/levigross/grequests"
	"time"
)

// Stores login data
type Data struct {
	Url    string
	Apikey string
	Ro     grequests.RequestOptions
}

// Defines API login principles that can be reused in requests
// Takes three parameters:
//  1. URL string
//  2. API key
//  3. Verify boolean that should be true in order to verify the servers certificate
// Returns Hivedata struct
func CreateLogin(inurl string, apikey string, verify bool) Data {
	formattedApikey := fmt.Sprintf("Bearer %s", apikey)
	return Data{
		Url:    inurl,
		Apikey: apikey,
		Ro: grequests.RequestOptions{
			Headers: map[string]string{
				"Authorization": formattedApikey,
			},
			RequestTimeout:     time.Duration(30) * time.Second,
			InsecureSkipVerify: !verify,
		},
	}
}

// Gets a raw json query and returns all data
// Takes one parameter:
//  1. search []bytes - Raw marshalled JSON string
// Returns multiple alerts and the request response
func (login *Data) UploadResult(id string, inputdata []byte) (*grequests.Response, error) {
	var url string
	url = fmt.Sprintf("%s/api/v1/uploadResult/%s", login.Url, id)

	// Dunno how this works for other types
	login.Ro.JSON = inputdata

	ret, err := grequests.Post(url, &login.Ro)
	return ret, err
}
