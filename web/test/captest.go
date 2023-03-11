package main

import (
	"github.com/afocus/captcha"
	"image/png"
	"net/http"
)

func main() {
	cap := captcha.New()
	// 设置字体
	cap.SetFont("comic.ttf")

	// 创建验证码 4个字符
	http.HandleFunc("/r", func(w http.ResponseWriter, r *http.Request){
		img, str := cap.Create(4, captcha.NUM)
		png.Encode(w, img)

		println(str)
	})
	http.ListenAndServe(":8085", nil)
}