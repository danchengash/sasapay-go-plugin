package helpers

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

var contentyTypeHeaderJson = []byte("application/json")

func NewReq(url string, body *[]byte, headers *map[string]string, debug ...bool) (*fasthttp.Response, error) {
	readTimeout, _ := time.ParseDuration("15s")
	writeTimeout, _ := time.ParseDuration("15s")
	maxIdleConnDuration, _ := time.ParseDuration("30m")
	tlsConf := &tls.Config{InsecureSkipVerify: true}
	dial := (&fasthttp.TCPDialer{Concurrency: 100, DNSCacheDuration: time.Hour}).Dial
	client := fasthttp.Client{
		Name:                          "sasapay-sdk",
		ReadTimeout:                   readTimeout,
		WriteTimeout:                  writeTimeout,
		MaxIdleConnDuration:           maxIdleConnDuration,
		NoDefaultUserAgentHeader:      true,
		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        true,
		TLSConfig:                     tlsConf,
		Dial:                          dial,
	}
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	req.SetRequestURI(url)
	req.Header.SetContentTypeBytes(contentyTypeHeaderJson)
	if headers != nil {
		for key, value := range *headers {
			req.Header.Set(key, value)
		}
	}
	//GET REQUEST.
	if body == nil {
		req.Header.SetMethod(fasthttp.MethodGet)

	} else if body != nil {
		//POST REQUEST
		req.Header.SetMethod(fasthttp.MethodPost)
		req.SetBodyRaw(*body)
	}
	err := client.Do(req, resp)
	if err != nil {
		fmt.Printf("<ERROR ->>: %s\n", err)

	}
	if len(debug) != 0 {
		if debug[0] {
			fmt.Printf("-------------REQUEST START------------\n")
			fmt.Printf("[BODY]: %s\n", req.Body())
			fmt.Printf("[CODE]: %d\n[RESPONSE]: %s\n", resp.StatusCode(), resp.Body())
			fmt.Printf("-------------REQUEST END------------\n")
		}

	}

	// RELEASE RESOURCES.
	fasthttp.ReleaseRequest(req)
	return resp, nil
}
