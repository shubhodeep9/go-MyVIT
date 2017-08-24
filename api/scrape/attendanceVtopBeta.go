package scrape

import (
	//"crypto/tls"
	//"fmt"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
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

type Attendance2 struct {
	Attend map[string]SubjectAttendance `json:"attendance"`
	Status string                       `json:"status"`
}

type SubjectAttendance struct {
	CourseCode  string `json:"course_code"`
	CourseTitle string `json:"course_title"`
	CourseType  string `json:"subject_type"`
	Slot        string `json:"slot"`
	FacultyName string `json:"faculty"`
	Attended    string `json:"attended"`
	TotalClass  string `json:"totalClasses"`
	AttendPer   string `json:"attendPer"`
}

func ScrapeAttendance(client http.Client, regNo, psswd, baseuri string) *Attendance2 {
	//fmt.Println("HERE")

	PostData3 := strings.NewReader("semesterSubId=VL2017181")
	req3, _ := http.NewRequest("POST", "https://vtopbeta.vit.ac.in/vtop/processViewStudentAttendance", PostData3)
	req3.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req3.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Max OS X 10_10_5) AppleWebKit (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")

	resp, err := client.Do(req3)
	var status string
	if err != nil {
		status = "Failure"
	} else {
		status = "Success"
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	html := string(body)
	//fmt.Println("html is", html)

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader((html)))
	table := doc.Find(".table")
	trow := table.Find("tr")

	attendance := make(map[string]SubjectAttendance)
	trow.EachWithBreak(func(i int, td *goquery.Selection) bool {
		td = td.Find("td")
		if td.Length() == 1 {
			return false
		}
		if i > 0 {
			code := trim(td.Eq(2).Text())
			ctype := trim(td.Eq(3).Text())
			if ctype == "ETH" {
				code = code + "_ETH"
			} else if ctype == "EPJ" {
				code = code + "_EPJ"
			} else if ctype == "ELA" {
				code = code + "_ELA"
			} else if ctype == "TH" {
				code = code + "_TH"
			} else if ctype == "SS" {
				code = code + "_SS"
				//fmt.Println("SS baby")
			}
            //fmt.Println(code)

			attendance[code] = SubjectAttendance{
				CourseCode:  trim(td.Eq(1).Text()),
				CourseTitle: trim(td.Eq(2).Text()),
				CourseType:  trim(td.Eq(3).Text()),
				Slot:        trim(td.Eq(4).Text()),
				FacultyName: trim(td.Eq(5).Text()),
				Attended:    trim(td.Eq(6).Text()),
				TotalClass:  trim(td.Eq(7).Text()),
				AttendPer:   trim(td.Eq(8).Text()),
			}
		}
		return true
	})

	return &Attendance2{
		Status: status,
		Attend: attendance,
	}

}

//https://vtopbeta.vit.ac.in/vtop/examinations/doSearchExamScheduleForStudent
