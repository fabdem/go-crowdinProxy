package crowdinproxy

import (
	"fmt"
	"net/http"
	"time"
	"net/url"
	"github.com/medisafe/go-crowdin"
	"github.com/mreiferson/go-httpclient"
)

// Default values for timeouts
var connectionTOinSecs time.Duration = 5
var readwriteTOinSecs time.Duration = 40

// Set connection and read/write timeouts for the subsequent new connections
func SetTimeouts(cnctTOinSecs, rwTOinSecs int) {
    connectionTOinSecs = time.Duration(cnctTOinSecs)
    readwriteTOinSecs = time.Duration(rwTOinSecs)
}


// New - create a new instance of crowdin  use of a PROXY.
func New(token, project, proxy string) (*crowdin.Crowdin, error) {

	// transport httpclient.Transport

	api := crowdin.New(token, project)

	var proxyUrl *url.URL
	var err error

	if len(proxy) > 0 {     // If a proxy is defined
		proxyUrl, err = url.Parse(proxy)
		if err != nil {
			fmt.Println("Bad proxy URL", err)
			return nil,err
		}
	}

	// Set proxy and timeouts
	transport := &httpclient.Transport{
		ConnectTimeout:   connectionTOinSecs * time.Second,
		ReadWriteTimeout: readwriteTOinSecs * time.Second,
		Proxy: http.ProxyURL(proxyUrl),
	}



	api.SetClient(&http.Client{
		Transport: transport,
	})
	return api, nil
}
