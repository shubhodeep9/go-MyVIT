package scrape

import (
	"github.com/patrickmn/go-cache"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/cache"
)

type CourseStruct struct {
	Courselist map[string]string `json:"courses"`
	Status     string            `json:"status"`
}

func Courses(bow *browser.Browser, regno, password, baseuri string, cac *cache.Cache) *CourseStruct {
	cacheSession.SetSession(bow, cac, regno)
	status := "Success"
	courselist := make(map[string]string)
	if false {
		status = "Failure"
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
		Status:     status,
	}
}
