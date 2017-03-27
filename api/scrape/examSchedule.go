/*
@Author Ujjwal Ayyangar
@Organization Google Developers Group VIT Vellore
	VCC!
*/

package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/status"
	//"os"
	//"sync"
)

type ExamSchedule struct {
	Status  string              `json:"status"`
	Cat1    map[string]Content3 `json:"cat1"`
	Cat2    map[string]Content3 `json:"cat2"`
	TermEnd map[string]Content3 `json:"termend"`
}
type MainExamSchedule struct {
	Exam_Schedule FinalExamSchedule `json:"exam_schedule"`
}

type FinalExamSchedule struct {
	Status status.StatusStruct `json:"status"`
	Exams  []exms              `json:"exams"`
}
type exms struct {
	Name     string     `json:"name"`
	Schedule []Content3 `json:"schedule"`
}

type Content3 struct {
	CourseCode   string `json:'courseCode'`
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

func ExmSchedule(bow *browser.Browser, reg, baseuri string, found bool) *MainExamSchedule {
	result := []exms{}
	var b FinalExamSchedule

	//sem := os.Getenv("SEM")
	//sem = "WS"
	stat := status.Success()
	var cat1 map[string]Content3
	var cat2 map[string]Content3
	var fat map[string]Content3
	list := make([]map[string]Content3, 100)
	list_count := 0
	count := 0
	if !found {
		stat = status.SessionError()
	} else {
		//var wg sync.WaitGroup
		//student/exam_schedule.asp?
		//exam_schedule.asp?sem=WS

		bow.Open(baseuri + "/student/exam_schedule.asp?sem=WS")
		//Reload
		if bow.Open("https://vtop.vit.ac.in/student/exam_schedule.asp?sem=WS") == nil {
			table := bow.Find("table").Eq(1)
			rows := table.Find("tr").Length()
			dets := make(map[string]Content3)

			table.Find("tr").Each(func(i int, s *goquery.Selection) {
				if i >= 1 && i <= rows-1 {
					//wg.Add(1)
					//go func() {
					//defer wg.Done()
					td := s.Find("td")

					if td.Length() == 1 {
						// dets is a map of type signature - keyType = String and valueType = Contents2. ["Code":Contents2,"Code":Contents2]
						// dets map will be reset for every examType that is why it is being intialized again and again
						if count > 0 {
							list_count = list_count + 1
							dets = make(map[string]Content3)
						}
						count = count + 1
					} else {
						dets[td.Eq(1).Text()] = Content3{
							CourseCode:   td.Eq(1).Text(),
							Course_Title: td.Eq(2).Text(),
							Slot:         td.Eq(4).Text(),
							Date:         td.Eq(5).Text(),
							Day:          td.Eq(6).Text(),
							Session:      td.Eq(7).Text(),
							Time:         td.Eq(8).Text()}
						list[list_count] = dets
					}
					//}()
				}
			})
			//wg.Wait()

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
				stat = status.SessionError()
			}
			//sched1 := []Content3{}

			sched1 := []Content3{}
			for _, v := range cat1 {
				sched1 = append(sched1, v)
			}
			exm1 := exms{
				Name:     "CAT 1",
				Schedule: sched1,
			}
			sched2 := []Content3{}
			for _, v := range cat2 {
				sched2 = append(sched2, v)
			}
			exm2 := exms{
				Name:     "CAT 2",
				Schedule: sched2,
			}
			sched3 := []Content3{}
			for _, v := range fat {
				sched3 = append(sched3, v)
			}
			exm3 := exms{
				Name:     "FAT",
				Schedule: sched3,
			}
			result = append(result, exm1, exm2, exm3)
			b = FinalExamSchedule{
				Status: stat,
				Exams:  result,
			}

		}
	}

	return &MainExamSchedule{Exam_Schedule: b}

}
