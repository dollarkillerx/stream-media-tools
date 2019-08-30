package tool

import (
	"fmt"
	"github.com/dollarkillerx/easyutils"
	"io"
	"net/http"
)

var cookies []*http.Cookie
var one int

func Get(url string) (io.ReadCloser, error) {
	st, e := easyutils.InitProxy()
	if e != nil {
		panic(e.Error())
	}

	if one == 0 {
		response, e := st.ReptileUserRequestFrom("https://www.crunchyroll.com/military/episode-12-military-668525", nil, nil)
		if e != nil {
			panic(e.Error())
		}

		cookies = response.Cookies()
		one += 1
	}

	resp, e := st.ReptileUserRequestFrom(url, nil, cookies)
	if e != nil {
		panic(e.Error())
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", resp.StatusCode)
	}
	return resp.Body, nil
}
