//https://vtop.vit.ac.in/student/profile_education_view.asp
/*
@Author Ujjwal Ayyangar
@Organization Google Developers Group VIT Vellore
	I err, therefore I am
	#GDGSwag
*/

package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/status"
	"os"
)

type StudentInfoStruct struct {
	BlockID     string `json:"blockId"`
	BlockName   string `json:"blockName"`
	RoomNo      string `json:"roomNo"`
	RoomType    string `json:"roomType"`
	MessType    string `json:"messType"`
	CatererName string `json:"catererName"`
}
type RoommatesStruct struct {
	RegisterNo string `json:"regno"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Program    string `json:"program"`
	School     string `json:"school"`
}



type HostelDetailsStruct struct {
	Status  status.StatusStruct `json:"status"`
	You     StudentInfoStruct   `json:"you"`
	Roomies []RoommatesStruct   `json:"roomMates"`
}

/*
Function to show hostel details,
Calls NewLogin to login to academics,
@param bow (surf Browser) registration_no password
@return ShowHostelDetails struct
*/
func ShowHostelDetails(bow *browser.Browser, reg, baseuri string, found bool) *HostelDetailsStruct {
	var (
		stud StudentInfoStruct
	)
	roomies := []RoommatesStruct{}
	sem := os.Getenv("SEM")
	sem = ""
	cnt := 1
	cnt2 := 0
	sem = ""
	stat := status.Success()
	if !found {
		stat = status.SessionError()
	} else {
		util := []string{}
		bow.Open(baseuri + "/student/hostel_info.asp?sem=" + sem)
		//Twice loading due to Redirect policy defined by academics.vit.ac.in
		if bow.Open(baseuri+"/student/hostel_info.asp?sem="+sem) == nil {
			table := bow.Find("table[width='700']")
			table.Find("tr").Each(func(i int, s *goquery.Selection) {
				td := s.Find("td")
				if cnt > 1 && cnt <= 8 {
					util = append(util, td.Eq(1).Text())
					if cnt == 8 {
						stud = StudentInfoStruct{
							BlockID:     util[0],
							BlockName:   util[1],
							RoomNo:      util[2],
							RoomType:    util[3],
							MessType:    util[4],
							CatererName: util[5],
						}
						util = []string{}

					}
				} else if cnt > 9 {
					ncnt := cnt - (9 + 6*cnt2)
					if ncnt > 0 && ncnt <= 5 {
						util = append(util, td.Eq(1).Text())
					} else if ncnt == 6 {
						cnt2 += 1
						temp := RoommatesStruct{
							RegisterNo: util[0],
							Name:       util[1],
							Gender:     util[2],
							Program:    util[3],
							School:     util[4],
						}

						util = []string{}
						roomies = append(roomies, temp)

					}

				}

				cnt += 1
			})

		}
	}

	return &HostelDetailsStruct{
		Status:  stat,
		You:     stud,
		Roomies: roomies,
	}

}
