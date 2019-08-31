package tool

import (
	"fmt"
	"github.com/dollarkillerx/easyutils"
	"io"
	"net/http"
	"sync"
)

var cookies []*http.Cookie
var one int

func Get(url string) (io.ReadCloser, error) {
	st, e := easyutils.InitProxy()
	if e != nil {
		panic(e.Error())
	}
	sy := sync.RWMutex{}

	sy.Lock()
	if one == 0 {
		response, e := st.ReptileUserRequestFrom("https://www.crunchyroll.com/", nil, nil)
		if e != nil {
			panic(e.Error())
		}

		cookies = response.Cookies()
		one += 1
	}
	sy.Unlock()

	resp, e := st.ReptileUserRequestFrom(url, nil, cookies)


	if e != nil {
		panic(e.Error())
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", resp.StatusCode)
	}
	return resp.Body, nil
}
