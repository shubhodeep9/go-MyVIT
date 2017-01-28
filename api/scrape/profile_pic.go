package scrape

import (
	"bufio"
	"encoding/base64"
	// "fmt"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"os"
)

func ProfilePhoto(bow *browser.Browser, reg, baseuri string) string {
	bow.Open(baseuri + "/student/view_photo.asp")
	if bow.Open(baseuri+"/student/view_photo.asp") == nil {
		out, _ := os.Create("api/" + reg + ".jpg")
		bow.Download(out)
		imgFile, _ := os.Open("api/" + reg + ".jpg")
		go os.Remove("api/" + reg + ".jpg")
		defer imgFile.Close()

		// create a new buffer base on file size
		fInfo, _ := imgFile.Stat()
		var size int64 = fInfo.Size()
		buf := make([]byte, size)

		// read file content into buffer
		fReader := bufio.NewReader(imgFile)
		fReader.Read(buf)

		// if you create a new image instead of loading from file, encode the image to buffer instead with png.Encode()

		// png.Encode(&buf, image)

		// convert the buffer bytes to base64 string - use buf.Bytes() for new image
		imgBase64Str := base64.StdEncoding.EncodeToString(buf)
		return imgBase64Str
	} else {
		return ""
	}
}
