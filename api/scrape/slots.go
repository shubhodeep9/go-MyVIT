package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/login"
)

type SlotsStruct struct {
	Courselist map[string]string  `json:"courses"`
	Status     login.StatusStruct `json:"status"`
}

func Slots(bow *browser.Browser, regno, password, baseuri, coursekey string, found bool) *SlotsStruct {

	status := login.StatusStruct{
		Message: "Successful execution",
		Code:    0,
	}
	courselist := make(map[string]string)
	if !found {
		status = login.StatusStruct{
			Message: "Session Timed Out",
			Code:    11,
		}
	} else {
		bow.Open(baseuri + "/student/coursepage_view.asp?sem=WS")
		bow.Open(baseuri + "/student/coursepage_view.asp?sem=WS&crs=" + coursekey)
		options := bow.Find("select").Eq(1).Find("option")
		options.Each(func(i int, s *goquery.Selection) {
			if i > 0 {
				val, _ := s.Attr("value")
				courselist[val] = s.Text()
			}
		})
	}
	return &SlotsStruct{
		Courselist: courselist,
		Status:     status,
	}
}
