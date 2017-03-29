package scrape

import (
	//"encoding/json"
	"fmt"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/status"
	//"time"
	"bufio"
	"encoding/base64"
	"os"
	"runtime"
	"strconv"
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

/*
func convert(s string) []string {
	temp := ""
	t := []string{}
	for _, i := range s {
		if i != '\t' && i != '\n' {
			temp = fmt.Sprintf("%s%c", temp, i)
		}
	}

	if len(temp) == 0 {
		return t
	} else {
		temp = temp[14:]
		for len(temp) > 0 {
			fmt.Println(temp, len(temp))
			if len(temp)%19 == 0 {
				t = append(t, fmt.Sprintf("%s %s %s", temp[0:3], temp[3:11], temp[11:19]))
				temp = temp[19:]
			} else if len(temp) < 7 {
				return t
				continue
			} else if len(temp) == 20 {
				return t
			} else {
				t = append(t, fmt.Sprintf("%s %s %s"), temp[0:3], temp[3:11])
				temp = temp[11:]
			}
		}
		return t
	}

}
*/
func FacultySearch(bow *browser.Browser, reg, password, baseuri string, found bool) *AllFacs {
	runtime.GOMAXPROCS(runtime.NumCPU())
	faculties := []facInfo{}
	stat := status.ServerError()
	if found {
		stat = status.Success()
		//11550
		//15198
		for empId := 11400; empId <= 15198; empId++ {
			str := fmt.Sprintf("%s/student/official_detail_view.asp?empid=%s", baseuri, strconv.Itoa(empId))
			bow.Open(str)
			bow.Open(str)
			table := bow.Find("table[width='761']")
			util := []string{}
			yes := false
			var txt string
			util2 := []string{}
			table.Find("tr").Each(func(i int, s *goquery.Selection) {

				if i > 0 && i < 9 {
					txt = s.Find("td").Eq(1).Text()
					if len(txt) == 0 && i == 1 {
						yes = true
					}
					util = append(util, s.Find("td").Eq(1).Text())
				} else if i == 9 {
					tabl2 := bow.Find("table[width='250']")
					tabl2.Find("tr[bgcolor='#CCCCCC']").Each(func(i int, s2 *goquery.Selection) {
						td := s2.Find("td")
						if len(td.Text()) != 0 {
							l := fmt.Sprintf("%s %s %s", td.Eq(0).Text(), td.Eq(1).Text(), td.Eq(2).Text())
							util2 = append(util2, l)
						}

					})

				}
			})
			if yes {
				continue
			}
			fmt.Println(bow.Url())
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
				Photo:          util[8],
			}
			faculties = append(faculties, temp)

		}

	}
	return &AllFacs{
		Status:    stat,
		Faculties: faculties,
	}
}
