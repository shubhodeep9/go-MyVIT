/*
@Author Shubhodeep Mukherjee
@Organization Google Developers Group VIT Vellore
	Isn't Go sexy?
	I know right!!
	Just like Shubhodeep
	I mean, have you seen the guy? xP
	#GDGSwag
*/

package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
Attendance structure json
*/
type Attendance struct {
	Average_Attendance float64            `json:"average_attendace"`
	AttendanceDet      map[string]Subject `json:"attendance_det"`
	Status             string             `json:"status"`
}

type Subject struct {
	Percentage int          `json:"attendance_percentage"`
	Classes    int          `json:"attended_classes"`
	Details    []DetsBranch `json:"details"`
	Date       string       `json:"registration_date"`
	TotalClass int          `json:"total_classes"`
}

type DetsBranch struct {
	Sl         int    `json:"sl"`
	ClassUnits int    `json:"class_units"`
	Date       string `json:"date"`
	Reason     string `json:"reason"`
	Slot       string `json:"slot"`
	Status     string `json:"status"`
}

/*
Function to get Course daily attendance,
@param classnbr baseuri bow (surf Browser)
@return List of DetsBranch struct
*/
func getDetails(semcode, classnbr, crscd, crstp, from_date, to_date, baseuri string, bow *browser.Browser) []DetsBranch {
	monthtoint := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}
	v := url.Values{}
	v.Set("semcode", semcode)
	v.Add("classnbr", classnbr)
	v.Add("from_date", from_date)
	v.Add("to_date", to_date)
	v.Add("crscd", crscd)
	v.Add("crstp", crstp)
	bow.PostForm(baseuri+"/student/attn_report_details.asp", v)
	table := bow.Find("table").Eq(2)
	tr := table.Find("tr")
	var detsbranchlis []DetsBranch
	var detsbranch DetsBranch
	tr.Each(func(i int, s *goquery.Selection) {
		if i > 1 {
			td := s.Find("td")
			date := strings.Split(td.Eq(1).Text(), "-")
			detsbranch = DetsBranch{
				Sl:         i - 1,
				ClassUnits: conver(td.Eq(4).Text()),
				Date:       date[2] + "-" + monthtoint[date[1]] + "-" + date[0],
				Reason:     td.Eq(5).Text(),
				Slot:       td.Eq(2).Text(),
				Status:     td.Eq(3).Text(),
			}
			detsbranchlis = append(detsbranchlis, detsbranch)
		}
	})
	return detsbranchlis
}

func conver(i string) int {
	l, _ := strconv.Atoi(i)
	return l
}

/*
Function to show Attendance,
Calls NewLogin to login to academics,
@param bow (surf Browser) registration_no password
@return Attendance struct
*/
func ShowAttendance(bow *browser.Browser, baseuri string) *Attendance {
	sem := os.Getenv("SEM")
	avg := 0
	status := "Success"
	tr_len := 0
	dets := make(map[string]Subject)
	if false {
		status = "Failure"
	} else {
		year, month, day := time.Now().Date()
		bow.Open(baseuri + "/student/attn_report.asp?sem=" + sem + "&fmdt=04-Jan-2017&todt=" + strconv.Itoa(day) + "-" + month.String()[:3] + "-" + strconv.Itoa(year))
		//Twice loading due to Redirect policy defined by academics.vit.ac.in
		if bow.Open(baseuri+"/student/attn_report.asp?sem="+sem+"&fmdt=04-Jan-2017&todt="+strconv.Itoa(day)+"-"+month.String()[:3]+"-"+strconv.Itoa(year)) == nil {
			table := bow.Find("table").Eq(4)
			tr := table.Find("tr")
			var wg sync.WaitGroup
			tr.Each(func(i int, s *goquery.Selection) {
				if i > 0 {
					wg.Add(1)
					go func() {
						defer wg.Done()
						td := s.Find("td")
						semcode, _ := s.Find("input[name=semcode]").Attr("value")
						classnbr, _ := s.Find("input[name=classnbr]").Attr("value")
						code := td.Eq(1).Text()
						crscd, _ := s.Find("input[name=crscd]").Attr("value")
						crstp, _ := s.Find("input[name=crstp]").Attr("value")
						from_date, _ := s.Find("input[name=from_date]").Attr("value")
						to_date, _ := s.Find("input[name=to_date]").Attr("value")
						if strings.Contains(td.Eq(3).Text(), "Lab") {
							code = code + "_L"
						}
						percent := td.Eq(8).Text()
						dets[code] = Subject{
							Percentage: conver(percent),
							Classes:    conver(td.Eq(6).Text()),
							Details:    getDetails(semcode, classnbr, crscd, crstp, from_date, to_date, baseuri, bow),
							Date:       td.Eq(5).Text(),
							TotalClass: conver(td.Eq(7).Text()),
						}
						perc := conver(td.Eq(8).Text()[:len(td.Eq(8).Text())-1])
						avg = avg + perc
					}()

				}
			})
			wg.Wait()
			tr_len = tr.Length() - 1
		}
	}
	return &Attendance{
		Average_Attendance: float64(avg / tr_len),
		AttendanceDet:      dets,
		Status:             status,
	}
}
