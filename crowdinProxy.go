package crowdin

import (
	//"encoding/json"
	//"errors"
	"fmt"
	//"io"
	"net/http"
	//"os"
	"time"
	"net/url"
	"github.com/medisafe/go-crowdin"

	"github.com/mreiferson/go-httpclient"
)


// New - create new instances of Crowdin API making use of a PROXY.
func New_prxy(token, project string, proxy string) (*crowdin.Crowdin, error) {

	proxyUrl, err := url.Parse("proxy")
	if err != nil {
		fmt.Println("Bad proxy URL", err)
		return nil,err
	} 

	transport := &httpclient.Transport{
		ConnectTimeout:   5 * time.Second,
		ReadWriteTimeout: 40 * time.Second,
		Proxy: http.ProxyURL(proxyUrl),
		
	}
	defer transport.Close()

	s := &crowdin.Crowdin{}
	s.config.apiBaseURL = apiBaseURL
	s.config.apiAccountBaseURL = apiAccountBaseURL
	s.config.token = token
	s.config.project = project
	s.config.client = &http.Client{
		Transport: transport,
	}
	return s, nil
}


