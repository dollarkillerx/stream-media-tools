package transcoding

import (
	"bytes"
	"github.com/dollarkillerx/easyutils/clog"
	"os/exec"
)

// ffmpeg -i foo.mp4 -codec copy -vbsf h264_mp4toannexb -map 0 -f segment -segment_list out.m3u8 -segment_time 10 out%03d.ts

func ToM3u8(file, out string) error {
	e, _, i := tom3u8(file, out)
	if e != nil {
		clog.Println(e.Error())
		return e
	}
	// 判断是否异常
	clog.Println(i)

	return nil

}

func tom3u8(file string, out string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("ffmpeg", "-i", file, "-codec", "copy", "-vbsf", "h264_mp4toannexb", "-map", "0", "-f", "segment", "-segment_list", out, "-segment_time", "10", "out%03d.ts")
	cmd.Stdout = &stdout // 输出
	cmd.Stderr = &stderr // 输出错误
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
