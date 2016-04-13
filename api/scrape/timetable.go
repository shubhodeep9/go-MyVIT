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
	"fmt"
	"github.com/patrickmn/go-cache"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/cache"
	"strings"
	"sync"
)

type Timetable struct {
	Status     string              `json:"status"`
	Time_table map[string]Contents `json:"time_table"`
}

type Contents struct {
	Class_number        string `json:"class_number"`
	Course_code         string `json:"course_code"`
	Course_mode         string `json:"course_mode"`
	Course_option       string `json:"course_option"`
	Course_title        string `json:"course_title"`
	Course_type         string `json:"course_type"`
	Faculty             string `json:"faculty"`
	Ltpjc               string `json:"ltpjc"`
	Registration_status string `json:"registration_status,omitempty"`
	Slot                string `json:"slot"`
	Venue               string `json:"venue"`
}

/*
Function to show Timetable,
Calls NewLogin to login to academics,
@param bow (surf Browser) registration_no password
@return Timtable struct
*/
func ShowTimetable(bow *browser.Browser, regno, password, baseuri string, cac *cache.Cache) *Timetable {
	cacheSession.SetSession(bow, cac, regno)
	conts := make(map[string]Contents)
	status := "Success"
	if 1 != 1 {
		status = "Failure"
	} else {
		bow.Open(baseuri + "/student/timetable_ws.asp")
		fmt.Println(bow.Url())
		//Twice loading due to Redirect policy defined by academics.vit.ac.in
		bow.Open(baseuri + "/student/timetable_ws.asp")
		tables := bow.Find("table")
		reg_table := tables.Eq(1)

		tr := reg_table.Find("tr")
		tr_len := tr.Length()
		var wg sync.WaitGroup
		tr.Each(func(i int, s *goquery.Selection) {
			if i > 0 && i < tr_len-2 {
				wg.Add(1)
				go func(conts map[string]Contents, s *goquery.Selection) {
					defer wg.Done()
					td := s.Find("td")
					code := td.Eq(3).Text()
					if code == "Embedded Lab" {
						code = td.Eq(1).Text() + "_L"
						conts[code] = Contents{
							Class_number:  td.Eq(0).Text(),
							Course_code:   td.Eq(1).Text(),
							Course_mode:   td.Eq(5).Text(),
							Course_option: td.Eq(6).Text(),
							Course_title:  td.Eq(2).Text(),
							Course_type:   td.Eq(3).Text(),
							Faculty:       td.Eq(9).Text(),
							Ltpjc:         strings.TrimSpace(td.Eq(4).Text()),
							Slot:          td.Eq(7).Text(),
							Venue:         td.Eq(8).Text(),
						}
					} else {
						conts[code] = Contents{
							Class_number:        td.Eq(2).Text(),
							Course_code:         td.Eq(3).Text(),
							Course_mode:         td.Eq(7).Text(),
							Course_option:       td.Eq(8).Text(),
							Course_title:        td.Eq(4).Text(),
							Course_type:         td.Eq(5).Text(),
							Faculty:             td.Eq(11).Text(),
							Ltpjc:               strings.TrimSpace(td.Eq(6).Text()),
							Registration_status: td.Eq(12).Text(),
							Slot:                td.Eq(9).Text(),
							Venue:               td.Eq(10).Text(),
						}
					}
				}(conts, s)
			}
		})
		wg.Wait()
		if len(conts) == 0 {
			status = "Failure"
		}
	}

	return &Timetable{
		Status:     status,
		Time_table: conts,
	}
}
