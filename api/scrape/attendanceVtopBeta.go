package scrape

import (
	//"crypto/tls"
	//"fmt"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"io/ioutil"
	//"log"
	"net/http"
	//"net/http/cookiejar"
	"net/url"
	//"os"
	//"os/exec"
	//"path/filepath"
	"strings"
)

type Attendance2 struct {
	Attend map[string]SubjectAttendance `json:"attendance"`
	Status string                       `json:"status"`
}

type SubjectAttendanceDetails struct {
	Sno      string `json:"sno"`
	Date     string `json:"date"`
	Slot     string `json:"slot"`
	DayTime  string `json:"day_time"`
	Attended string `json:"status"`
}

type SubjectAttendance struct {
	CourseCode  string                     `json:"course_code"`
	CourseTitle string                     `json:"course_title"`
	CourseType  string                     `json:"subject_type"`
	Slot        string                     `json:"slot"`
	FacultyName string                     `json:"faculty"`
	Attended    string                     `json:"attended"`
	TotalClass  string                     `json:"totalClasses"`
	AttendPer   string                     `json:"attendPer"`
	Details     []SubjectAttendanceDetails `json:"details"`
}

func Details(td *goquery.Selection, client http.Client) []SubjectAttendanceDetails {
	var details []SubjectAttendanceDetails
	slot := strings.Split(trim(td.Eq(4).Text()), "+")
	a, _ := td.Eq(9).Html()
	t_classID := strings.Split(a, "&#39") // Temporary classID store
	if len(t_classID) > 1 {
		classID := string(t_classID[1][1:])
		form := url.Values{}
		form.Add("classId", classID)
		if len(slot) > 1 {
			form.Add("slotName", slot[0]+" "+slot[1])
		} else {
			form.Add("slotName", slot[0])
		}
		attReq, _ := http.NewRequest("POST", "https://vtopbeta.vit.ac.in/vtop/processViewAttendanceDetail", strings.NewReader(form.Encode()))
		attReq.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		attReq.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Max OS X 10_10_5) AppleWebKit (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")
		respDet, _ := client.Do(attReq)
		body, _ := ioutil.ReadAll(respDet.Body)
		respDet.Body.Close()
		html := string(body)
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader((html)))
		table2 := doc.Find(".table")
		trow2 := table2.Find("tr")
		trow2.Each(func(i int, tr *goquery.Selection) {
			if i > 0 {
				td := tr.Find("td")
				temp := SubjectAttendanceDetails{
					Sno:      trim(td.Eq(0).Text()),
					Date:     trim(td.Eq(1).Text()),
					DayTime:  trim(td.Eq(3).Text()),
					Slot:     trim(td.Eq(2).Text()),
					Attended: trim(td.Eq(4).Text()),
				}
				details = append(details, temp)
			}

		})

	}
	return details
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
			tempDet := Details(td, client)
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
				Details:     tempDet,
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
