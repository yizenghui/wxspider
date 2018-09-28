package wxspider

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
)

//CheckImage 检查图片合法性
// 1. 宽高大于或等于320
func CheckImage(imageURL string) bool {
	resp, err := http.Get(imageURL)
	if err != nil {
		return false
	}
	c, _, err := image.DecodeConfig(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return false
	}
	width := c.Width
	height := c.Height
	// log.Println("width = ", width, "height = ", height)
	if width >= 320 && height >= 320 {
		return true
	}
	// file.Close()
	return false
}
