package scrape

import (
	//"encoding/json"
	"fmt"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/status"
	//"time"
	//"bufio"
	//"encoding/base64"
	//"os"
	"runtime"
	//"strconv"
	//"strings"
)

type facInfo struct {
	Name           string   `json:"name"`
	School         string   `json:"school"`
	Designation    string   `json:"designation"`
	Venue          string   `json:"venue"`
	Intercom       string   `json:"intercom"`
	EmailId        string   `json:"email"`
	Division       string   `json:"division"`
	AdditionalRole string   `json:"additionalRole"`
	OpenHour       []string `json:"openHours"`
	Photo          string   `json:"photo"`
}

//
//"Tue10:00 AM12:00 PM
//Fri10:00 AM12:00 PM"
type AllFacs struct {
	Status    status.StatusStruct `json:"status"`
	Faculties []facInfo           `json:"allFaculties"`
}

func FacultySearch(bow *browser.Browser, reg, password, query, baseuri string, found bool) *AllFacs {
	runtime.GOMAXPROCS(runtime.NumCPU())
	faculties := []facInfo{}
	stat := status.ServerError()
	if found {
		stat = status.Success()
		//getfacdet.asp?x=Wed,%2029%20Mar%202017%2013:17:17%20GMT&fac=SENTHIL
		bow.Open("https://vtop.vit.ac.in/student/getfacdet.asp?x=Wed,%2029%20Mar%202017%2013:17:17%20GMT&fac=" + query)
		bow.Open("https://vtop.vit.ac.in/student/getfacdet.asp?x=Wed,%2029%20Mar%202017%2013:17:17%20GMT&fac=" + query)
		table := bow.Find("table")
		table.Find("tr[bgcolor='#EDEADE']").Each(func(i int, s *goquery.Selection) {
			s = s.Find("a")
			link, _ := s.Attr("href")

			bow.Open(baseuri + "/student/" + link)
			bow.Open(baseuri + "/student/" + link)
			//fmt.Print(bow.Body())
			facTable := bow.Find("table[width='761']")
			util := []string{}
			util2 := []string{}
			facTable.Find("tr").Each(func(i int, s *goquery.Selection) {
				if i > 0 && i < 9 {
					util = append(util, s.Find("td").Eq(1).Text())
				} else if i == 9 {
					table2 := bow.Find("table[width='250']")
					table2.Find("tr[bgcolor='#CCCCCC']").Each(func(i int, s2 *goquery.Selection) {
						td := s2.Find("td")
						if len(td.Text()) != 0 {
							l := fmt.Sprintf("%s %s %s", td.Eq(0).Text(), td.Eq(1).Text(), td.Eq(2).Text())
							util2 = append(util2, l)
						}

					})

				}

			})
			/*
				bow.Open(baseuri + "/student/emp_photo.asp")
				out, _ := os.Create("api/" + reg + ".jpg")
				bow.Download(out)
				imgFile, _ := os.Open("api/" + reg + ".jpg")
				go os.Remove("api/" + reg + ".jpg")
				defer imgFile.Close()
				util2 = []string{}
				// create a new buffer base on file size
				fInfo, err := imgFile.Stat()
				if err == nil {
					var size int64 = fInfo.Size()
					buf := make([]byte, size)

					// read file content into buffer
					fReader := bufio.NewReader(imgFile)
					fReader.Read(buf)
					util = append(util, base64.StdEncoding.EncodeToString(buf))
				} else {
					util = append(util, "")
				}
			*/
			temp := facInfo{
				Name:           util[0],
				School:         util[1],
				Designation:    util[2],
				Venue:          util[3],
				Intercom:       util[4],
				EmailId:        util[5],
				Division:       util[6],
				AdditionalRole: util[7],
				OpenHour:       util2,
				//Photo:          util[8],
			}
			faculties = append(faculties, temp)

		})
	}

	return &AllFacs{
		Status:    stat,
		Faculties: faculties,
	}
}
