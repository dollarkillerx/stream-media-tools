package test

import (
	"github.com/dollarkillerx/stream-media-tools/transcoding"
	"testing"
)

func TestTsToMp4(t *testing.T) {
	err := transcoding.TsToMp4("./dow/main.ts", "./out/1.mp4")
	if err != nil {

	}
}
