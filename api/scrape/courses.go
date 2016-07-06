package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/status"
)

type CourseStruct struct {
	Courselist map[string]string   `json:"courses"`
	Status     status.StatusStruct `json:"status"`
}

func Courses(bow *browser.Browser, regno, password, baseuri string, found bool) *CourseStruct {

	stats := status.Success()
	courselist := make(map[string]string)
	if !found {
		stats = status.SessionError()
	} else {
		bow.Open(baseuri + "/student/coursepage_view.asp?sem=WS")
		bow.Open(baseuri + "/student/coursepage_view.asp?sem=WS")
		options := bow.Find("select").Eq(0).Find("option")
		options.Each(func(i int, s *goquery.Selection) {
			if i > 0 {
				val, _ := s.Attr("value")
				courselist[val] = s.Text()
			}
		})
	}
	return &CourseStruct{
		Courselist: courselist,
		Status:     stats,
	}
}
