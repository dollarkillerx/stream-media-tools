package transcoding

import (
	"bytes"
	"github.com/dollarkillerx/easyutils/clog"
	"os/exec"
)

// ffmpeg  -i  uploadfile/video/test.wmv -c:v h264 -strict -2 uploadfile/mp4/real.mp4

// Ts to MP4
func ToH264(file, out string) error {
	e, _, i := toh264(file, out)
	if e != nil {
		clog.Println(e.Error())
		return e
	}
	// 判断是否异常
	clog.Println(i)

	return nil

}

func toh264(file string, out string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("ffmpeg", "-i", file, "-c:v", "h264", "-strict", "-2", out)
	cmd.Stdout = &stdout // 输出
	cmd.Stderr = &stderr // 输出错误
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
