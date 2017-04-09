package scrape

import (
	//"bytes"
	"bufio"
	"encoding/base64"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/status"
	"os"
	//"strconv"
)

type FacPhoto struct {
	Status status.StatusStruct `json:"status"`
	Photo  string              `json:"photo"`
}

//
//"Tue10:00 AM12:00 PM
//Fri10:00 AM12:00 PM"

func FacultyPhoto(bow *browser.Browser, reg, password, query, baseuri string, found bool) *FacPhoto {
	//faculties := []facInfo{}

	stat := status.ServerError()
	var temp FacPhoto
	if found {

		stat = status.Success()

		//getfacdet.asp?x=Wed,%2029%20Mar%202017%2013:17:17%20GMT&fac=SENTHIL
		bow.Open(baseuri + "/student/official_detail_view.asp?empid=" + query)
		if bow.Open(baseuri+"/student/official_detail_view.asp?empid="+query) == nil {
			bow.Open(baseuri + "/student/emp_photo.asp")
			out, _ := os.Create("api/" + reg + ".jpg")
			bow.Download(out)
			imgFile, _ := os.Open("api/" + reg + ".jpg")
			go os.Remove("api/" + reg + ".jpg")
			defer imgFile.Close()

			// create a new buffer base on file size
			fInfo, err := imgFile.Stat()
			if err == nil {
				var size int64 = fInfo.Size()
				buf := make([]byte, size)

				// read file content into buffer
				fReader := bufio.NewReader(imgFile)
				fReader.Read(buf)
				temp = FacPhoto{Photo: base64.StdEncoding.EncodeToString(buf)}

			}

		}
	}
	temp.Status = stat

	return &temp
}
