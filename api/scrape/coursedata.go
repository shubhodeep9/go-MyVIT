package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/status"
	"os"
	"strings"
)

type CourseDataStruct struct {
	Uploads Upload              `json:"uploads"`
	Status  status.StatusStruct `json:"status"`
}

type Upload struct {
	TextMaterial  []Description             `json:"text_material"`
	QuestionPaper map[string]Questionstruct `json:"question_paper"`
	Assignments   []Description             `json:"assignments"`
	Lecture       []LectureStruct           `json:"lecture"`
}

type Questionstruct struct {
	Q Description `json:"question_paper"`
	A Description `json:"answer_key"`
}

type Description struct {
	Link string `json:"link"`
	Name string `json:"name"`
}

type LectureStruct struct {
	Date     string      `json:"date"`
	Day      string      `json:"day"`
	Topic    string      `json:"topic"`
	Material Description `json:"material"`
}

func CourseData(bow *browser.Browser, regno, password, baseuri, coursekey, slt, fac string, found bool) *CourseDataStruct {
	sem := os.Getenv("SEM")
	stats := status.Success()
	var upload Upload
	if !found {
		stats = status.SessionError()
	} else {
		bow.Open(baseuri + "/student/coursepage_plan_view.asp?sem=" + sem)
		if bow.Open(baseuri+"/student/coursepage_plan_view.asp?sem="+sem) == nil {
			fm, _ := bow.Form("form")
			fm.Input("sem", sem)
			fm.Set("course", coursekey)
			fm.Set("slot", slt)
			fm.Set("faculty", fac)
			fm.Submit()
			fm = bow.Forms()[3]
			fm.Set("sem", sem)
			classnbr, _ := bow.Find("input[name=classnbr]").Attr("value")
			crscd, _ := bow.Find("input[name=crscd]").Attr("value")
			crstp, _ := bow.Find("input[name=crstp]").Attr("value")
			fm.Set("classnbr", classnbr)
			fm.Set("crscd", crscd)
			fm.Set("crstp", crstp)
			fm.Submit()
			outer_table := bow.Find("table")
			inners := outer_table.Find("table")
			materials := inners.Eq(1)
			materials.Find("td[width='76%']").Each(func(i int, s *goquery.Selection) {
				s = s.Find("a")
				link, found := s.Attr("href")
				if found {
					u := Description{
						Link: baseuri + "/student/" + link,
						Name: strings.TrimSpace(s.Text()),
					}
					upload.TextMaterial = append(upload.TextMaterial, u)
				}
			})
			if inners.Length() == 6 {
				questions := inners.Eq(2)
				qpaper := make(map[string]Questionstruct)
				questions.Find("tr").Each(func(i int, s *goquery.Selection) {
					td := s.Find("td")
					title := strings.Split(td.Eq(0).Text(), " ")[0]
					a := td.Eq(1).Find("a")
					if link, found := a.Attr("href"); found {
						if i%2 == 0 {
							qpaper[title] = Questionstruct{
								Q: Description{
									Link: baseuri + "/student/" + link,
									Name: strings.TrimSpace(a.Text()),
								},
							}
						} else {
							qpaper[title] = Questionstruct{
								A: Description{
									Link: baseuri + "/student/" + link,
									Name: strings.TrimSpace(a.Text()),
								},
							}
						}
					}
				})
				upload.QuestionPaper = qpaper
			}

			assignment := inners.Eq(inners.Length() - 2)
			assignment.Find("td[width='76%']").Each(func(i int, s *goquery.Selection) {
				s = s.Find("a")
				link, found := s.Attr("href")
				if found {
					u := Description{
						Link: baseuri + "/student/" + link,
						Name: strings.TrimSpace(s.Text()),
					}
					upload.Assignments = append(upload.Assignments, u)
				}
			})

			lecture := inners.Eq(inners.Length() - 1)
			var lecstruct LectureStruct
			lecture.Find("tr[bgcolor='#EDEADE']").Each(func(i int, s *goquery.Selection) {
				td := s.Find("td")
				if td.Length() == 5 {
					lecstruct.Date = strings.TrimSpace(td.Eq(1).Text())
					lecstruct.Day = strings.TrimSpace(td.Eq(2).Text())
					lecstruct.Topic = strings.TrimSpace(td.Eq(3).Text())
					a := td.Eq(4).Find("a")
					link, found := a.Attr("href")
					if found {
						lecstruct.Material = Description{
							Link: baseuri + "/student/" + link,
							Name: strings.TrimSpace(a.Text()),
						}
					}
					upload.Lecture = append(upload.Lecture, lecstruct)
				}
			})
		}
	}
	return &CourseDataStruct{
		Uploads: upload,
		Status:  stats,
	}
}
