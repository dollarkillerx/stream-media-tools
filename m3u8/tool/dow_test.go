package tool

import (
	"github.com/dollarkillerx/easyutils"
	"io/ioutil"
	"log"
	"testing"
)

func TestDow(t *testing.T) {
	tag := "https://a-vrv.akamaized.net/evs/c54c039c654c2a7d459447cc5774d825/assets/epkupkj4q8gmbir_1076103.mp4/index-v1-a1.m3u8?t=st=1567063114~exp=1567149514~acl=/evs/c54c039c654c2a7d459447cc5774d825/assets/epkupkj4q8gmbir_*~hmac=1032be4ddd0b223bb2255f7ab3e5a6c7f9a3e180af14e51698ce4bfa572c61b2"

	st, e := easyutils.InitProxy()
	if e != nil {
		panic(e.Error())
	}

	response, e := st.ReptileUserRequestFrom(tag, nil, nil)
	defer response.Body.Close()
	if e != nil {
		panic(e.Error())
	}

	t.Log(response.StatusCode)

	bytes, e := ioutil.ReadAll(response.Body)
	if e != nil {
		panic(e.Error())
	}

	log.Println(string(bytes))
}
