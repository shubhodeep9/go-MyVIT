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

type AttendanceDetails struct {
	Attend map[string][]SubjectAttendanceDetails2 `json:"attendance"`
	Status string                                 `json:"status"`
}

type SubjectAttendanceDetails2 struct {
	Sno      string `json:"sno"`
	Date     string `json:"date"`
	Slot     string `json:"slot"`
	DayTime  string `json:"day_time"`
	Attended string `json:"status"`
}

func ShowAttendanceDetails(client http.Client, regNo, psswd, baseuri string) *AttendanceDetails {
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

	attendance := make(map[string][]SubjectAttendanceDetails2)

	trow.Each(func(i int, tr *goquery.Selection) {
		var details []SubjectAttendanceDetails2
		td := tr.Find("td")
		if td.Length() > 1 {
			code := trim(td.Eq(1).Text())
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
			}

			slot := strings.Split(trim(td.Eq(4).Text()), "+")
			a, _ := td.Eq(10).Html()
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
				//postData := strings.NewReader("classId=" + classID + "slotName=" + slot[0] + " " + slot[1])
				//dummy := strings.NewReader()
				attReq, _ := http.NewRequest("POST", "https://vtopbeta.vit.ac.in/vtop/processViewAttendanceDetail", strings.NewReader(form.Encode()))
				attReq.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
				attReq.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Max OS X 10_10_5) AppleWebKit (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")
				//attReq.PostForm = form
				respDet, _ := client.Do(attReq)
				body, _ = ioutil.ReadAll(respDet.Body)
				respDet.Body.Close()
				html = string(body)
				//fmt.Println("html is", html)

				doc, _ = goquery.NewDocumentFromReader(strings.NewReader((html)))
				table2 := doc.Find(".table")
				trow2 := table2.Find("tr")

				trow2.Each(func(i int, tr *goquery.Selection) {
					if i > 0 {
						td := tr.Find("td")
						temp := SubjectAttendanceDetails2{
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
			attendance[code] = details
		}

	})

	return &AttendanceDetails{
		Status: status,
		Attend: attendance,
	}

}

//https://vtopbeta.vit.ac.in/vtop/examinations/doSearchExamScheduleForStudent
