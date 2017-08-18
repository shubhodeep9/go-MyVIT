package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/http/cookiejar"
	//"github.com/headzoo/surf/browser"
	"crypto/tls"
	//"github.com/headzoo/surf"
	"io/ioutil"
	"net/url"
	//"strconv"
	"strings"
	//"os"
	"os/exec"
)

type Timetable struct {
	Status     string              `json:"status"`
	Time_table map[string]Contents `json:"time_table"`
}

func trim(s string) string{
    return strings.TrimSpace(s)
}

type Contents struct {
	Class_number        string    `json:"class_number"`
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

type TimeStruct struct {
	Day       int    `json:"day"`
	StartTime string `json:"day"`
	EndTime   string `json:"end_time"`
}

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar:       jar,
		Transport: tr,
	}
	postData := url.Values{}

	req, _ := http.NewRequest("GET", "https://vtopbeta.vit.ac.in/vtop/", strings.NewReader(postData.Encode()))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Max OS X 10_10_5) AppleWebKit (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")
	resp, err := client.Do(req)
	//fmt.Println("JAr", jar, "\n")
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	html := string(body)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader((html)))
	var base64 string
	doc.Find("img[alt='vtopCaptcha']").Each(func(i int, inp *goquery.Selection) {

		b64, _ := inp.Attr("src")
		b64 = strings.Split(b64, ",")[1]
		base64 = b64
	})


	_, err2 := exec.Command("python", "parseIt.py", base64).Output()
	var captcha string
	if err2 == nil {
		captchaFile, err3 := ioutil.ReadFile("output.txt")
		if err3 == nil {
			captcha = string(captchaFile)
			//fmt.Println("Captcha is", captcha)
		}

	}
	postData.Add("uname", "15BCB0064")
	postData.Add("passwd", "Arsenal@1997")
	postData.Add("captchaCheck", captcha)

	PostData2 := strings.NewReader("uname=15BCB0064&passwd=Arsenal@1997&captchaCheck=" + captcha)
	req2, _ := http.NewRequest("POST", "https://vtopbeta.vit.ac.in/vtop/processLogin", PostData2)

	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err = client.Do(req2)

	body, _ = ioutil.ReadAll(resp.Body)
	PostData3 := strings.NewReader("semesterSubId=VL2017181")
	req3, _ := http.NewRequest("POST", "https://vtopbeta.vit.ac.in/vtop/processViewTimeTable", PostData3)
	req3.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req3.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Max OS X 10_10_5) AppleWebKit (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")

	resp, err = client.Do(req3)
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	html = string(body)

	doc, _ = goquery.NewDocumentFromReader(strings.NewReader((html)))
	table := doc.Find(".table")
	trow := table.Find("tr")

	conts := make(map[string]Contents)
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
                fmt.Println("SS baby")
			}


			conts[code] = Contents{
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
    for k,_:= range(conts){
        fmt.Println(k,conts[k])
    }

}
