package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/login"
)

type FacStruct struct{
	Courselist map[string]string `json:"courses"`
	Status string `json:"status"`
}

func Facs(bow *browser.Browser,regno, password, baseuri, coursekey, slt string) *FacStruct{
	response := login.NewLogin(bow,regno,password,baseuri)
	status := "Success"
	courselist := make(map[string]string)
	if response.Status == 0{
		status = "Failure"
	} else {
		bow.Open(baseuri+"/student/coursepage_view.asp?sem=WS")
		bow.Open(baseuri+"/student/coursepage_view.asp?sem=WS&crs="+coursekey+"&slt="+slt)
		options := bow.Find("select").Eq(2).Find("option")
		options.Each(func(i int, s *goquery.Selection){
			if i>0 {
				val, _ := s.Attr("value")
				courselist[val] = s.Text()
			}
		})
	}
	return &FacStruct{
		Courselist: courselist,
		Status: status,
	}
}