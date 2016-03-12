package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/login"
	"time"
	"strconv"
)

type Attendance struct {
	AttendanceDet map[string]Subject `json:"attendance_det"`
	Status string `json:"status"`
}

type Subject struct {
	Percentage string `json:"attendance_percentage"`
	Classes string `json:"attended_classes"`
	Date string `json:"registration_date"`
	TotalClass string `json:"total_classes"`
}

func ShowAttendance(bow *browser.Browser,regno, password, baseuri string) *Attendance{
	response := login.NewLogin(bow,regno,password,baseuri)
	status := "Success" 
	dets := make(map[string]Subject)
	if response.Status != 1 {
		status = "Failure"
	} else {
		year, month, day := time.Now().Date()
		bow.Open(baseuri+"/student/attn_report.asp?sem=WS&fmdt=04-Jan-2016&todt="+strconv.Itoa(day)+"-"+month.String()[:3]+"-"+strconv.Itoa(year))
		//Twice loading due to Redirect policy defined by academics.vit.ac.in
		bow.Open(baseuri+"/student/attn_report.asp?sem=WS&fmdt=04-Jan-2016&todt="+strconv.Itoa(day)+"-"+month.String()[:3]+"-"+strconv.Itoa(year))
		table := bow.Find("table").Eq(3)
		tr := table.Find("tr")
		
		tr.Each(func(i int, s *goquery.Selection){
			if i>0 {
				td := s.Find("td")
				dets[td.Eq(1).Text()] = Subject{
					Percentage: td.Eq(8).Text(),
					Classes: td.Eq(6).Text(),
					Date: td.Eq(5).Text(),
					TotalClass: td.Eq(7).Text(),
				}
			}
		})
	}
	return &Attendance{
		AttendanceDet: dets,
		Status: status,
	}
}
