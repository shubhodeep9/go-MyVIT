/*
@Author Ujjwal Ayyangar
@Organization Google Developers Group VIT Vellore
	VCC!
*/

package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"os"
	"sync"
)

type ExamSchedule struct {
	Status  string               `json:"status"`
	Cat1    map[string]Contents2 `json:"cat1"`
	Cat2    map[string]Contents2 `json:"cat2"`
	TermEnd map[string]Contents2 `json:"termend"`
}

type Contents2 struct {
	Course_Title string `json:"crTitle"`
	Slot         string `json:"slot"`
	Date         string `json:"date"`
	Day          string `json:"day"`
	Session      string `json:"session"`
	Time         string `json:"time"`
}

/*
Function ->ExmSchedule to fetch the exam schedule,

@return ExamSchedule struct
*/

func ExmSchedule(bow *browser.Browser, baseuri string) *ExamSchedule {
	sem := os.Getenv("SEM")
	status := "Success"
	var cat1 map[string]Contents2
	var cat2 map[string]Contents2
	var fat map[string]Contents2
	list := make([]map[string]Contents2, 100)
	list_count := 0
	count := 0
	if false {
		status = "Failure"
	} else {
		var wg sync.WaitGroup
		bow.Open(baseuri + "/student/exam_schedule.asp?sem=" + sem)
		//Reload
		if bow.Open(baseuri+"/student/exam_schedule.asp?sem="+sem) == nil {
			table := bow.Find("table").Eq(1)
			rows := table.Find("tr").Length()
			dets := make(map[string]Contents2)

			table.Find("tr").Each(func(i int, s *goquery.Selection) {
				if i >= 1 && i < rows-1 {
					wg.Add(1)
					go func() {
						defer wg.Done()
						td := s.Find("td")

						if td.Length() == 1 {
							// dets is a map of type signature - keyType = String and valueType = Contents2. ["Code":Contents2,"Code":Contents2]
							// dets map will be reset for every examType that is why it is being intialized again and again
							if count > 0 {
								list_count = list_count + 1
								dets = make(map[string]Contents2)
							}
							count = count + 1
						} else {
							dets[td.Eq(1).Text()] = Contents2{
								Course_Title: td.Eq(2).Text(),
								Slot:         td.Eq(4).Text(),
								Date:         td.Eq(5).Text(),
								Day:          td.Eq(6).Text(),
								Session:      td.Eq(7).Text(),
								Time:         td.Eq(8).Text()}
							list[list_count] = dets
						}
					}()
				}
			})
			wg.Wait()

			if len(list[0]) != 0 {
				cat1 = list[0]
			}
			if len(list[1]) != 0 {
				cat2 = list[1]
			}
			if len(list[2]) != 0 {
				fat = list[2]
			}
			if len(list) == 0 {
				status = "Failure"
			}
		}
	}
	return &ExamSchedule{Status: status, Cat1: cat1, Cat2: cat2, TermEnd: fat}
}
