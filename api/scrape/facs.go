package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/login"
)

type FacStruct struct {
	Courselist map[string]string  `json:"courses"`
	Status     login.StatusStruct `json:"status"`
}

func Facs(bow *browser.Browser, regno, password, baseuri, coursekey, slt string, found bool) *FacStruct {

	status := login.StatusStruct{
		Message: "Successful execution",
		Code:    0,
	}
	courselist := make(map[string]string)
	if false {
		status = login.StatusStruct{
			Message: "Session Timed Out",
			Code:    11,
		}
	} else {
		bow.Open(baseuri + "/student/coursepage_view.asp?sem=WS")
		bow.Open(baseuri + "/student/coursepage_view.asp?sem=WS&crs=" + coursekey + "&slt=" + slt)
		options := bow.Find("select").Eq(2).Find("option")
		options.Each(func(i int, s *goquery.Selection) {
			if i > 0 {
				val, _ := s.Attr("value")
				courselist[val] = s.Text()
			}
		})
	}
	return &FacStruct{
		Courselist: courselist,
		Status:     status,
	}
}
