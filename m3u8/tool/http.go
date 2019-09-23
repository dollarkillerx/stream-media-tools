package tool

import (
	"fmt"
	"github.com/dollarkillerx/easyutils/httplib"
	"io"
)

func Get(url string) (io.ReadCloser, error) {

	get := httplib.Get(url)
	resp, e := get.Response()

	if e != nil {
		return nil, e
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", resp.StatusCode)
	}
	return resp.Body, nil
}
