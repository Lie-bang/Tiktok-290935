package Service

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
)

func ffmpegPicture(filepath string) (string, error) {
	//filepath = filepath + ".mp4"
	buf := bytes.NewBuffer(nil)
	err := ffmpeg_go.Input(filepath+".mp4").Filter("select", ffmpeg_go.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg_go.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).WithOutput(buf, os.Stdout).Run()

	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	img, err1 := imaging.Decode(buf)

	if err1 != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	err2 := imaging.Save(img, filepath+".png")
	if err2 != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}
	return filepath + ".png", nil
}
