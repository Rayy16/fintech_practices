package tools

import (
	"bytes"
	"fintechpractices/global"
	"fmt"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ExtractVedioToImage(src, dest string) error {
	return ExtractVedioToImageWithFrameNum(src, dest, 1)
}

func ExtractVedioToImageWithFrameNum(src, dest string, frameNum int) error {
	log := global.Log.Sugar()

	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(src).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n, %d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "png"}).
		WithOutput(buf).Run()
	if err != nil {
		log.Errorf("ffmpeg stream operation failed: %s", err.Error())
		return err
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		log.Errorf("imaging.Decode err: %s", err.Error())
		return err
	}

	err = imaging.Save(img, dest)
	if err != nil {
		log.Errorf("imaging.Save err: %s", err.Error())
		return err
	}
	return nil
}
