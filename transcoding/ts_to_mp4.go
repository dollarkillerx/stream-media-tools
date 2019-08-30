package transcoding

import (
	"bytes"
	"github.com/dollarkillerx/easyutils/clog"
	"os/exec"
)

// Ts to MP4
func TsToMp4(file, out string) error {
	e, _, i := tomp4(file, out)
	if e != nil {
		clog.Println(e.Error())
		return e
	}
	// 判断是否异常
	clog.Println(i)

	return nil

}

func tomp4(file string, out string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("ffmpeg", "-i", file, "-acodec", "copy", "-vcodec", "copy", out)
	cmd.Stdout = &stdout // 输出
	cmd.Stderr = &stderr // 输出错误
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
