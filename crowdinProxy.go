package crowdinproxy

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
func New(token, project, proxy string) (*crowdin.Crowdin, error) {

	proxyUrl, err := url.Parse("proxy")
	if err != nil {
		fmt.Println("Bad proxy URL", err)
		return nil,err
	} 

    api := crowdin.New(token, project)
	
	transport := &httpclient.Transport{
		ConnectTimeout:   5 * time.Second,
		ReadWriteTimeout: 40 * time.Second,
		Proxy: http.ProxyURL(proxyUrl),
	}
	
	api.SetClient(&http.Client{
		Transport: transport,
	})
	
	return api, nil
}


