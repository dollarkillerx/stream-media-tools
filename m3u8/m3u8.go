package m3u8

import (
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/stream-media-tools/m3u8/dl"
)

func M3U8Download(output, url string) error {
	downloader, err := dl.NewTask(output, url)
	if err != nil {
		clog.Println(err.Error())
		return err
	}
	if err := downloader.Start(25); err != nil {
		clog.Println(err.Error())
		return err
	}

	return nil
}
