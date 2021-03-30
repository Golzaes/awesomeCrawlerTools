package Requester

import (
	"bufio"
	"fmt"
	"github.com/Golzaes/awesomeCrawlerTools/Encoder"
	"github.com/Golzaes/awesomeCrawlerTools/Head"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

func Fetch(method, url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf(`NewRequest Error:%#v`, err)
	}
	req.Header.Add(`User-Agent`, Head.RandUa())
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf(` Client Error:%#v`, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(` Get invalid status code %s while scraping %s`, res.Status, url)
	}
	// Detection encode，eg: UTF8 GBK...
	bodyReader := bufio.NewReader(res.Body)
	PageEncode := Encoder.DetectionEncode(bodyReader)
	// recoding
	encodedReader := transform.NewReader(bodyReader, PageEncode.NewDecoder())
	return ioutil.ReadAll(encodedReader)
}
