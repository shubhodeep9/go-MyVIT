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
	//"strings"
	"fmt"
	"sync"
)

type ExamSchedule struct {
	Status string `json:"status"`
	Cat1 Contents2 `json:"cat1"`
	Cat2 Contents2 `json:"cat2"`
	TermEnd Contents2 `json:"termend"`
	}//Eschedule map[string]Contents `json:"eSchedule"`
//}

type Contents2 struct {

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
	dets := make(map[string]Contents2)
	var list []string
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
		tr_len:=tr.Length()
		tr.Find("tr").Each(func(i int, s *goquery.Selection){
			if i>0 && i<tr_len-2 {
				wg.Add(1)
				var head string
				go func(dets map[string]Contents2,s *goquery.Selection){
					defer wg.Done()
					td := s.Find("td")
					if(td.Length()==1){
						head :=td.Eq(1).Text()
						list=append(list,head)
					}
					if(td.Length() !=1){
					dets[head] = Contents2 {
						Course_Title: td.Eq(2).Text(),
						Slot: td.Eq(4).Text(),
						Date: td.Eq(5).Text(),
						Day: td.Eq(6).Text(),
						Session: td.Eq(7).Text(),
						Time: td.Eq(8).Text()	}

				}
				}(dets,s)
			}
		})
		wg.Wait()
		fmt.Println(list)
		if len(dets)==0 {
			status = "Failure"
		}

	}
	return &ExamSchedule{Status: status,Cat1: dets["CAT-I"],Cat2: dets["CAT-II"],TermEnd: dets["FAT"]}
}
