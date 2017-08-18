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

type Timetable3 struct {
	Status     string               `json:"status"`
	Time_table map[string]Contents3 `json:"time_table"`
}

func trim(s string) string {
	return strings.TrimSpace(s)
}

type Contents3 struct {
	Class_number        string `json:"class_number"`
	Course_code         string `json:"course_code"`
	Course_mode         string `json:"course_mode"`
	Course_option       string `json:"course_option"`
	Course_title        string `json:"course_title"`
	Course_type         string `json:"subject_type"`
	Faculty             string `json:"faculty"`
	Ltpjc               string `json:"ltpc"`
	Registration_status string `json:"registration_status"`
	Slot                string `json:"slot"`
	Venue               string `json:"venue"`
	//BillDate string `json:"bill_date"`
	//BillNumber string `json:"bill_number"`
	//ProjectTitle string `json:"project_title"`
	//Timings []TimeStruct `json:"timings"`
	//Attendance Subject `json:"attendace"`
	//Marks Mrks `json:"marks"`
}

type TimeStruct3 struct {
	Day       int    `json:"day"`
	StartTime string `json:"day"`
	EndTime   string `json:"end_time"`
}

func ShowTimetable(client http.Client, regNo, psswd, baseuri string) *Timetable3 {
	//fmt.Println("HERE")

	PostData3 := strings.NewReader("semesterSubId=VL2017181")
	req3, _ := http.NewRequest("POST", "https://vtopbeta.vit.ac.in/vtop/processViewTimeTable", PostData3)
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

	conts := make(map[string]Contents3)
	trow.EachWithBreak(func(i int, td *goquery.Selection) bool {
		td = td.Find("td")
		if td.Length() == 1 {
			return false
		}
		if i > 0 {
			code := trim(td.Eq(2).Text())
			ctype := trim(td.Eq(4).Text())
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

			conts[code] = Contents3{
				Class_number:  trim(td.Eq(1).Text()),
				Course_code:   trim(td.Eq(2).Text()),
				Course_title:  trim(td.Eq(3).Text()),
				Course_type:   trim(td.Eq(4).Text()),
				Ltpjc:         trim(td.Eq(5).Text()) + trim(td.Eq(6).Text()) + trim(td.Eq(7).Text()) + trim(td.Eq(8).Text()) + trim(td.Eq(9).Text()),
				Course_option: trim(td.Eq(10).Text()),
				Slot:          trim(td.Eq(11).Text()),
				Venue:         trim(td.Eq(12).Text()),
				Faculty:       trim(td.Eq(13).Text()),
			}
		}
		return true
	})

	return &Timetable3{
		Status:     status,
		Time_table: conts,
	}

}
