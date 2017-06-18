/*
@Author Ujjwal Ayyangar
@Organization Google Developers Group VIT Vellore
	VCC!
*/

package scrape

import (
	//"fmt"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/status"
	"net/url"
	"os"
	//"sync"
)

type CalCourses struct {
	Status status.StatusStruct `json:"status"`
	Course []Course            `json:"course"`
}

type Course struct {
	CourseType  string    `json:"coursetype"`
	CourseType2 string    `json:"crstp"`
	ClassNbr    string    `json:"classnbr"`
	CourseCode  string    `json:"crscd"`
	CourseTitle string    `json:"crstitle"`
	Faculty     string    `json:"faculty"`
	Sem         string    `json:"sem"`
	Assignments []Assigns `json:"assignments"`
	Reviews     []Project `json:"reviews"`
}
type Assigns struct {
	Number       string `json:"Num"`
	Title        string `json:"Title"`
	DueDate      string `json:"duedate"`
	MaxMark      string `json:"maxMark"`
	Question     string `json:"question"`
	Answer       string `json:"answer"`
	StatusAssign string `json:"assignStatus"`
	Score        string `json:"score"`
	Detail       string `json:"detail"`
}
type Project struct {
	Number   string `json:"Num"`
	ReviewNo string `json:"reviewNo"`
	MaxMark  string `json:"maxmark"`
	Status   string `json:"status"`
	Score    string `json:"score"`
	Comments string `json:"comments"`
}

/*
Function ->CalCourseFunc to fetch the cal courses of the student,
@return CalCourses struct
*/

func CalCourseFunc(bow *browser.Browser, reg, baseuri string, found bool) *CalCourses {

	sem := os.Getenv("SEM")
	//sem = "WS"
	stat := status.Success()
	courses := []Course{}

	if !found {
		stat = status.SessionError()
	} else {
		//https://vtop.vit.ac.in/student/marks_da.asp?sem=WS
		bow.Open(baseuri + "/student/marks_da.asp?sem=" + sem)
		//Reload
		if bow.Open(baseuri+"/student/marks_da.asp?sem="+sem) == nil {

			table := bow.Find("table[width='900']")
			table.Find("tr").Each(func(i int, s *goquery.Selection) {

				if i > 0 {

					td := s.Find("td")
					//s.Find("input[name=semcode]").Attr("value")
					sem, _ := s.Find("input[name=sem]").Attr("value")
					classnbr, _ := s.Find("input[name=classnbr]").Attr("value")
					crscd, _ := s.Find("input[name=crscd]").Attr("value")
					crstp, _ := s.Find("input[name=crstp]").Attr("value")

					crsType := ""
					switch crstp {
					case "ETH":
						crsType = "Embedded Theory"
					case "ELA":
						crsType = "Embedded Lab"
					case "EPJ":
						crsType = "Embedded Project"
					}
					action, _ := s.Find("form[method=post]").Attr("action")
					data := url.Values{}
					data.Set("sem", sem)
					data.Set("classnbr", classnbr)
					data.Set("crscd", crscd)
					data.Set("crstp", crstp)
					bow.PostForm(baseuri+"/student/"+action, data)
					if crsType != "Embedded Project" {
						assignmentTable := bow.Find("table[width='1000']")
						assignments := []Assigns{}
						rows := assignmentTable.Find("tr")
						rows.Each(func(i2 int, s2 *goquery.Selection) {
							if i2 > 1 && i2 < rows.Length()-1 {
								td2 := s2.Find("td")
								t := Assigns{
									Number:       td2.Eq(0).Text(),
									Title:        td2.Eq(1).Text(),
									DueDate:      td2.Eq(2).Text(),
									MaxMark:      td2.Eq(3).Text(),
									Question:     td2.Eq(4).Text(),
									Answer:       td2.Eq(5).Text(),
									StatusAssign: td2.Eq(6).Text(),
									Score:        td2.Eq(7).Text(),
									Detail:       td2.Eq(8).Text(),
								}
								assignments = append(assignments, t)

							}

						})

						temp := Course{
							CourseType:  crsType,
							CourseType2: crstp,
							ClassNbr:    classnbr,
							CourseCode:  crscd,
							CourseTitle: td.Eq(3).Text(),
							Faculty:     td.Eq(5).Text(),
							Sem:         sem,
							Assignments: assignments,
						}
						courses = append(courses, temp)
					} else {

						marksTable := bow.Find("table[width='800']")
						reviews := []Project{}
						rows := marksTable.Find("tr")
						rows.Each(func(i2 int, s2 *goquery.Selection) {
							if i2 > 1 && i2 < rows.Length() {
								td2 := s2.Find("td")
								t := Project{
									Number:   td2.Eq(0).Text(),
									ReviewNo: td2.Eq(1).Text(),
									MaxMark:  td2.Eq(2).Text(),
									Status:   td2.Eq(4).Text(),
									Score:    td2.Eq(5).Text(),
									Comments: td2.Eq(6).Text(),
								}
								reviews = append(reviews, t)
							}

						})
						temp := Course{
							CourseType:  crsType,
							CourseType2: crstp,
							ClassNbr:    classnbr,
							CourseCode:  crscd,
							CourseTitle: td.Eq(3).Text(),
							Faculty:     td.Eq(5).Text(),
							Sem:         sem,
							Reviews:     reviews,
						}
						//fmt.Println(temp)
						courses = append(courses, temp)
					}

				}
			})

			//fmt.Println(bow.Body())
		}

	}

	return &CalCourses{Status: stat, Course: courses}

}
