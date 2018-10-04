package mediaCtx

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	// "fmt"
	// "image/png"
	// "image/gif"
	// "image/svg"
	"log"
	"os"
	"strings"
)

func Resize(dir string, fn string) bool{
	checkExtFn := strings.Split(strings.ToLower(fn), ".")
	ExtFn := strings.ToLower(checkExtFn[len(checkExtFn)-1])

	// if ExtFn == "jpg" || ExtFn == "jpeg" || ExtFn == "gif" || ExtFn == "png" {
	if ExtFn == "jpg" || ExtFn == "jpeg" {
		oriImg := "." + dir + fn
		thumbImg := "." + dir + "thumbnails-128/" + fn

		if _, err := os.Stat(thumbImg); os.IsNotExist(err) {
			if _, err := os.Stat(oriImg); os.IsNotExist(err) {
				return false
			}
			return ResizeProcess(oriImg, thumbImg)
		}else{
			return true
		}
	}
	return false
}

func ResizeProcess(oriImg string, thumbImg string) bool{
	// open "test.jpg"
	file, err := os.Open(oriImg)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(128, 0, img, resize.Lanczos3)

	out, err := os.Create(thumbImg)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
	return true
}