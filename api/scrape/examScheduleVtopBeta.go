package scrape

import (
	//"crypto/tls"
	"fmt"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/status"
	"io/ioutil"
	//"log"
	"net/http"
	//"net/http/cookiejar"
	//"net/url"
	//"os"
	//"os/exec"
	//"path/filepath"
	"strings"
)

type MainExamSchedule2 struct {
	Exam_Schedule FinalExamSchedule2 `json:"exam_schedule"`
}

type FinalExamSchedule2 struct {
	Status status.StatusStruct `json:"status"`
	Exams  []exms              `json:"exams"`
}
type exms struct {
	Name     string        `json:"name"`
	Schedule []ExamContent `json:"schedule"`
}

type ExamContent struct {
	CourseCode  string `json:'courseCode'`
	CourseTitle string `json:"course_title"`
	CourseType  string `json:"course_type"`
	ClassID     string `json:"classId"`
	Slot        string `json:"slot"`
	Date        string `json:"date"`
	Session     string `json:"session"`
	Time        string `json:"time"`
	Venue       string `json:"venue"`
	Seat        string `json:"seat"`
	SeatNo      string `json:"seatNo"`
}

func ShowExamScheduleVtopBeta(client http.Client, regNo, psswd, baseuri string) *MainExamSchedule2 {
	fmt.Println("HERE")

	PostData3 := strings.NewReader("semesterSubId=VL2017181")
	req3, _ := http.NewRequest("POST", "https://vtopbeta.vit.ac.in/vtop/examinations/doSearchExamScheduleForStudent", PostData3)
	req3.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req3.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Max OS X 10_10_5) AppleWebKit (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")

	resp, err := client.Do(req3)
	//fmt.Println("response", resp)
	stat := status.Success()
	if err != nil {
		stat = status.SessionError()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	html := string(body)
	//	fmt.Println("html is", html)

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader((html)))
	table := doc.Find(".table")
	trow := table.Find("tr")
	var Exam string
	cat1 := []ExamContent{}
	cat2 := []ExamContent{}
	fat := []ExamContent{}
	trow.Each(func(i int, td *goquery.Selection) {
		td = td.Find("td")
		fmt.Println("ESSS")
		if td.Length() == 1 {
			Exam = trim(td.Text())
			//fmt.Println("EXAM IS", Exam)
		} else {
			if i > 0 {
				//Course Code Course title CourseType ClassID Slot ExamDate ExamSession ExamTime Venue SeatLocation SeatNo
				t := ExamContent{
					CourseCode:  trim(td.Eq(1).Text()),
					CourseTitle: trim(td.Eq(2).Text()),
					CourseType:  trim(td.Eq(3).Text()),
					ClassID:     trim(td.Eq(4).Text()),
					Slot:        trim(td.Eq(5).Text()),
					Date:        trim(td.Eq(6).Text()),
					Session:     trim(td.Eq(7).Text()),
					Time:        trim(td.Eq(8).Text()),
					Venue:       trim(td.Eq(9).Text()),
					Seat:        trim(td.Eq(10).Text()),
					SeatNo:      trim(td.Eq(11).Text()),
				}
				if Exam == "CAT1" {
					cat1 = append(cat1, t)
				} else if Exam == "CAT2" {
					cat2 = append(cat2, t)
				} else {
					fat = append(fat, t)
				}
			}
		}

	})

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

	ret := FinalExamSchedule2{
		Status: stat,
		Exams:  finalSchedule,
	}

	return &MainExamSchedule2{Exam_Schedule: ret}

}

//https://vtopbeta.vit.ac.in/vtop/examinations/doSearchExamScheduleForStudent
