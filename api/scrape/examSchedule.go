/*
@Author Ujjwal Ayyangar
@Organization Google Developers Group VIT Vellore
	I err, therefore I am.
*/

package scrape

import (
	//"fmt"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/status"
	//"os"
	//"sync"
)

/*

type ExamSchedule struct {
	Status  string              `json:"status"`
	Cat1    map[string]Content3 `json:"cat1"`
	Cat2    map[string]Content3 `json:"cat2"`
	TermEnd map[string]Content3 `json:"termend"`
}
*/
type MainExamSchedule struct {
	Exam_Schedule FinalExamSchedule `json:"exam_schedule"`
}

type FinalExamSchedule struct {
	Status status.StatusStruct `json:"status"`
	Exams  []exms              `json:"exams"`
}
type exms struct {
	Name     string    `json:"name"`
	Schedule []Content `json:"schedule"`
}

type Content struct {
	CourseCode   string `json:'courseCode'`
	Course_Title string `json:"crTitle"`
	Slot         string `json:"slot"`
	Date         string `json:"date"`
	Day          string `json:"day"`
	Session      string `json:"session"`
	Time         string `json:"time"`
	Venue        string `json:"venue"`
	Table        string `json:"table"`
}

/*
Function ->ExmSchedule to fetch the exam schedule,
@return ExamSchedule struct
*/

func ExmSchedule(bow *browser.Browser, reg, baseuri string, found bool) *MainExamSchedule {

	cat1 := []Content{}
	cat2 := []Content{}
	fat := []Content{}
	//sem := os.Getenv("SEM")
	//sem = "WS"
	stat := status.Success()

	if !found {
		stat = status.SessionError()
	} else {
		var Exam string

		bow.Open(baseuri + "/student/exam_schedule.asp?sem=WS")
		//Reload
		if bow.Open(baseuri+"/student/exam_schedule.asp?sem=WS") == nil {

			table := bow.Find("table[width='897']")
			rows := table.Find("tr").Length()
			table.Find("tr").Each(func(i int, s *goquery.Selection) {
				td := s.Find("td")
				if i >= 1 && i <= rows-1 {
					if td.Length() == 1 {
						Exam = td.Text()
					} else {
						t := Content{
							CourseCode:   td.Eq(1).Text(),
							Course_Title: td.Eq(2).Text(),
							Slot:         td.Eq(4).Text(),
							Date:         td.Eq(5).Text(),
							Day:          td.Eq(6).Text(),
							Session:      td.Eq(7).Text(),
							Time:         td.Eq(8).Text(),
							Venue:        td.Eq(9).Text(),
							Table:        td.Eq(10).Text(),
						}
						if Exam == "CAT - I" {
							cat1 = append(cat1, t)
						} else if Exam == "CAT - II" {
							cat2 = append(cat2, t)
						} else {
							fat = append(fat, t)
						}

					}
				}
			})
		}

	}
	CAT1 := exms{
		Name:     "CAT - I",
		Schedule: cat1,
	}
	CAT2 := exms{
		Name:     "CAT - II",
		Schedule: cat2,
	}
	FAT := exms{
		Name:     "FAT",
		Schedule: fat,
	}
	finalSchedule := []exms{CAT1, CAT2, FAT}

	ret := FinalExamSchedule{
		Status: stat,
		Exams:  finalSchedule,
	}

	return &MainExamSchedule{Exam_Schedule: ret}

}
