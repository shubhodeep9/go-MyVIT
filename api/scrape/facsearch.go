package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
)

type FacultySearc struct {
	Name        string `json:"name"`
	Designation string `json:"designation"`
	School      string `json:"school"`
	Venue       string `json:"venue"`
}

func FacultySearch(bow *browser.Browser, regno, password, baseuri, query string, found bool) string {
	if found {
		bow.Open(baseuri + "/student/getfacdet.asp?fac=" + query)
		bow.Open(baseuri + "/student/getfacdet.asp?fac=" + query)
	}
	return "hey"
}
