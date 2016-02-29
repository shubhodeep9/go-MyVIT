package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"strings"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/login"
	"fmt"
)

type Timetable struct {
	Status string	`json:"status"`
	Time_table map[string]Contents `json:"time_table"`
}

type Contents struct {
	Class_number string `json:"class_number"`
	Course_code string	`json:"course_code"`
	Course_mode string	`json:"course_mode"`
	Course_option string `json:"course_option"`
	Course_title string `json:"course_title"`
	Course_type string `json:"course_type"`
	Faculty string `json:"faculty"`
	Ltpjc string `json:"ltpjc"`
	Registration_status string `json:"registration_status,omitempty"`
	Slot string `json:"slot"`
	Venue string `json:"venue"`
}

func ShowTimetable(bow *browser.Browser,regno, password string) *Timetable {
	login.NewLogin(bow,regno,password)
	bow.Open("https://academics.vit.ac.in/student/timetable_ws.asp")
	//Twice loading due to Redirect policy defined by academics.vit.ac.in
	bow.Open("https://academics.vit.ac.in/student/timetable_ws.asp")
	tables := bow.Find("table")
	reg_table := tables.Eq(1)
	conts := make(map[string]Contents)
	reg_table.Find("tr").Each(func(i int, s *goquery.Selection){
		if i>0 && i<reg_table.Find("tr").Length()-1 {
			code := s.Find("td").Eq(3).Text()
			if code == "Embedded Lab" {
				code = s.Find("td").Eq(1).Text()+"_L"
				conts[code] = Contents{
					Class_number: s.Find("td").Eq(0).Text(),
					Course_code:s.Find("td").Eq(1).Text(),
					Course_mode:s.Find("td").Eq(5).Text(),
					Course_option:s.Find("td").Eq(6).Text(),
					Course_title:s.Find("td").Eq(2).Text(),
					Course_type:s.Find("td").Eq(3).Text(),
					Faculty:s.Find("td").Eq(9).Text(),
					Ltpjc:strings.TrimSpace(s.Find("td").Eq(4).Text()),
					Slot:s.Find("td").Eq(7).Text(),
					Venue:s.Find("td").Eq(8).Text(),
				}
			} else {
				conts[code] = Contents{
					Class_number: s.Find("td").Eq(2).Text(),
					Course_code:s.Find("td").Eq(3).Text(),
					Course_mode:s.Find("td").Eq(7).Text(),
					Course_option:s.Find("td").Eq(8).Text(),
					Course_title:s.Find("td").Eq(4).Text(),
					Course_type:s.Find("td").Eq(5).Text(),
					Faculty:s.Find("td").Eq(11).Text(),
					Ltpjc:strings.TrimSpace(s.Find("td").Eq(6).Text()),
					Registration_status:s.Find("td").Eq(12).Text(),
					Slot:s.Find("td").Eq(9).Text(),
					Venue:s.Find("td").Eq(10).Text(),
				}
			}
		}
	})
	return &Timetable{
		Status: "Success",
		Time_table: conts,
	}
}

