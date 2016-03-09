/*
@Author Ujjwal Ayyangar
@Organization Google Developers Group VIT Vellore
	VCC!




*/


package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/login"
	"strings"
	"sync"
)

type ExamSchedule struct {
	Status string `json:"status"`
	Cat1 Contents `json:"cat1"`
	Cat2 Contents `json:"cat2"`
	TermEnd Contents `json:"termend"`
	//Eschedule map[string]Contents `json:"eSchedule"`
}

type Contents struct {

	Course_Title string `json:"crTitle"`
	Slot string `json:"slot"`
	Date string `json:"date"`
	Day string `json:"day"`
	Session string `json:"session"`
	Time string `json:"time"`
}

/*
Function ->ExmSchedule to fetch the exam schedule,

@return ExamSchedule struct
*/

func ExmSchedule(bow *browser.Browser,regno, password, baseuri string) *ExamSchedule{
	response := login.NewLogin(bow,regno,password,baseuri)
	status := "Success"
	dets := make(map[string]Contents)
	list := make(map[int]string)
	if response.Status == 0 {
		status = "Failure"
	} else {
		var wg sync.WaitGroup
		bow.Open(baseuri+"/student/exam_schedule.asp")
		//Reload
		bow.Open(baseuri+"/student/exam_schedule.asp")
		tables := bow.Find("table")

		schedTable := tables.Eq(1)
		tr:=schedTable.Find("tr")
		tr_len=tr.Length()
		tr.Find("tr").Each(func(i int, s *goquery.Selection){
			if i>0 && i<tr_len-2 {
				wg.Add(1)
				var head string
				go func(){
					defer wg.Done()
					td := s.Find("td")
					if(td.length()==1){
						head :=td.Eq(1).Text()
						list=Append(list,head)
					}
					if(td.length() !=1){
					dets[head] = Contents {
						Course_Title: td.Eq(2).Text()
						Slot : td.Eq(4).Text()
						Date : td.Eq(5).Text()
						Day : td.Eq(6).Text()
						Session : td.Eq(7).Text()
						Time : td.Eq(8).Text()
					}

				}
				}()
			}
		})
		wg.Wait()

	}
	return &ExamSchedule{
		Status: status,
		Cat1:dets[list[0]]
		Cat2:dets[list[1]]
		TermEnd:dets[list[2]]
	}
}
