package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"

	"net/url"
)

type CourseDataStruct struct {
	Uploads []string `json:"uploads"`
	Status  string   `json:"status"`
}

func CourseData(bow *browser.Browser, regno, password, baseuri, coursekey, slt, fac string) *CourseDataStruct {

	status := "Success"
	var upload []string
	//courselist := make(map[string]string)
	if false {
		status = "Failure"
	} else {

		bow.Open(baseuri + "/student/coursepage_view.asp?sem=WS")
		bow.Open(baseuri + "/student/coursepage_view.asp?sem=WS&crs=" + coursekey + "&slt=" + slt + "&fac=" + fac)
		v := url.Values{}
		v.Set("sem", "WS")
		crsplancode, _ := bow.Find("input[name=crsplancode]").Attr("value")
		v.Add("crsplancode", crsplancode)
		v.Add("crpnvwcmd", "View")
		bow.PostForm(baseuri+"/student/coursepage_view3.asp", v)
		bow.Find("a").Each(func(i int, s *goquery.Selection) {
			href, _ := s.Attr("href")
			upload = append(upload, href)
		})
	}
	return &CourseDataStruct{
		Uploads: upload,
		Status:  status,
	}
}
