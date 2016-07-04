/*
@Author Shubhodeep Mukherjee
@Organization Google Developers Group VIT Vellore
	Isn't Go sexy?
	I know right!!
	Just like Shubhodeep
	I mean, have you seen the guy? xP
	#GDGSwag
*/

package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"

	"strings"
	"sync"
)

type Advisor struct {
	Status  string            `json:"status"`
	Details map[string]string `json:"faculty_det"`
}

/*
Function to fetch faculty advisor details,
Calls NewLogin to login to academics,
@param bow(surf Browser) registration_no password
@return Advisor struct
*/
func FacultyAdvisor(bow *browser.Browser, baseuri string) *Advisor {

	status := "Success"
	dets := make(map[string]string)
	if false {
		status = "Failure"
	} else {
		var wg sync.WaitGroup
		bow.Open(baseuri + "/student/faculty_advisor_view.asp")
		//Reload
		bow.Open(baseuri + "/student/faculty_advisor_view.asp")
		table := bow.Find("table").Eq(1)
		rows := table.Find("tr").Length()
		table.Find("tr").Each(func(i int, s *goquery.Selection) {
			if i > 0 && i < rows-1 {
				wg.Add(1)
				go func() {
					defer wg.Done()
					td := s.Find("td")
					dets[strings.TrimSpace(td.Eq(0).Text())] = strings.TrimSpace(td.Eq(1).Text())

				}()
			}
		})
		wg.Wait()
	}
	return &Advisor{
		Status:  status,
		Details: dets,
	}
}
