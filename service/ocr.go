package service

import (
	"fmt"
	"github.com/deloz/baiduocr"
	"os"
)

func Ocr() {
	ocr := &baiduocr.Ocr{
		APIKey:       "your app key",

		//如果以下参数省略,会自动使用默认值
		FromDevice:   baiduocr.FROM_DEVICE_API,
		ClientIP:     "your client IP",
		DetectType:   baiduocr.DETECT_TYPE_LOCATE_RECOGNIZE,
		LanguageType: baiduocr.LANGUAGE_TYPE_CHN_ENG,
		ImageType:    baiduocr.IMAGE_TYPE_ORIGINAL,
		//或者 使用base64编码
		//ImageType: baiduocr.IMAGE_TYPE_BASE64,
	}

	words, err := ocr.Scan("D:\\meetyouGo\\wingo\\bg1.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	txt := ""
	for _, word := range words {
		txt += word
	}

	fmt.Println("--------")
	fmt.Println(txt)
	fmt.Println("--------")
}
