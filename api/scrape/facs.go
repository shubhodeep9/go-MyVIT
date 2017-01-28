package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/status"
	"os"
)

type FacStruct struct {
	Courselist map[string]string   `json:"courses"`
	Status     status.StatusStruct `json:"status"`
}

func Facs(bow *browser.Browser, regno, password, baseuri, coursekey, slt string, found bool) *FacStruct {
	sem := os.Getenv("SEM")
	stats := status.Success()
	courselist := make(map[string]string)
	if !found {
		stats = status.SessionError()
	} else {
		bow.Open(baseuri + "/student/coursepage_plan_view.asp?sem=" + sem)
		if bow.Open(baseuri+"/student/coursepage_plan_view.asp?sem="+sem) == nil {
			fm, _ := bow.Form("form")
			fm.Input("sem", sem)
			fm.Set("course", coursekey)
			fm.Set("slot", slt)
			fm.Submit()
			options := bow.Find("select").Eq(2).Find("option")
			options.Each(func(i int, s *goquery.Selection) {
				if i > 0 {
					val, _ := s.Attr("value")
					courselist[val] = s.Text()
				}
			})
		}
	}
	return &FacStruct{
		Courselist: courselist,
		Status:     stats,
	}
}
