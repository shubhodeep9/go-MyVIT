package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/status"
	"os"
)

type LeaveRequestStruct struct {
	ApprovingAuthority []string            `json:"approvingAuthority"`
	LeaveType          []string            `json:"leaveType"`
	Status             status.StatusStruct `json:"status"`
}

func LeaveRequest(bow *browser.Browser, reg, baseuri string, found bool) *LeaveRequestStruct {
	sem := os.Getenv("SEM")
	authorities := []string{}
	leaves := []string{}
	stats := status.Success()
	if !found {
		stats = status.SessionError()
	} else {
		//https://vtop.vit.ac.in/student/leave_request.asp
		bow.Open(baseuri + "/student/leave_request.asp?sem=" + sem)
		if bow.Open(baseuri+"/student/leave_request.asp?sem="+sem) == nil {
			options := bow.Find("select").Eq(0).Find("option")
			options.Each(func(i int, s *goquery.Selection) {
				if i > 0 {
					authorities = append(authorities, s.Text())
				}
			})

			option2 := bow.Find("select").Eq(1).Find("option")
			option2.Each(func(i int, s *goquery.Selection) {
				if i > 0 {
					leaves = append(leaves, s.Text())
				}
			})

		}
	}
	return &LeaveRequestStruct{
		ApprovingAuthority: authorities,
		LeaveType:          leaves,
		Status:             stats,
	}
}
