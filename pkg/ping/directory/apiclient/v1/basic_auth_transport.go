package apiclient

import (
	"net/http"
)

// BasicAuthTransport is a struct to implement a http.Client transport
// function which injects a formatted Authorization header
type BasicAuthTransport struct {
	Username string
	Password string
}

func (bat BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	//req.Header.Set("Authorization", fmt.Sprintf("Basic %s",
	//	base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s",
	//		bat.Username, bat.Password)))))
	req.SetBasicAuth(bat.Username, bat.Password)
	return http.DefaultTransport.RoundTrip(req)
}
