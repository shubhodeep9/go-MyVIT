/*
@Author Ujjwal Ayyangar
@Organization Google Developers Group VIT Vellore
	I err, therefore I am
	#GDGSwag
*/

package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"os"
)

type MessagesStructUtil struct {
	From   string `json:"from"`
	Course string `json:"course"`
	Msg    string `json:"message"`
	Posted string `json:"posted"`
}
type MessagesStruct struct {
	FacultyMsgs []MessagesStructUtil `json:"faculty_messages"`
	Status      string               `json:"status"`
}

/*
Function to show messages,
Calls NewLogin to login to academics,
@param bow (surf Browser) registration_no password
@return Attendance struct
*/
func FacMessage(bow *browser.Browser, reg, baseuri string, found bool) *MessagesStruct {
	facmess := []MessagesStructUtil{}
	sem := os.Getenv("SEM")
	status := "Success"
	//tr_len := 0
	//dets := make(map[string]Subject)
	if !found {
		status = "Failure"
	} else {

		bow.Open(baseuri + "/student/class_message_view.asp?sem=" + sem)
		cnt := 1
		util := []string{}
		//Twice loading due to Redirect policy defined by academics.vit.ac.in
		if bow.Open(baseuri+"/student/class_message_view.asp?sem="+sem) == nil {
			tables := bow.Find("table[cellpadding='3']")
			tables.Find("tr").Each(func(i int, s *goquery.Selection) {
				td1 := s.Find("td[width='90']")
				td2 := s.Find("td[width='350']")
				if td1.Text() == "Advisor" {
					util = append(util, "Faculty Advisor", "-")
					cnt = 3

				} else if td1.Text() == "Faculty" {
					util = append(util, "Class Faculty")
					cnt = 2
				} else if len(td2.Text()) != 0 {
					if cnt < 4 {

						util = append(util, td2.Text())
						cnt += 1

					} else {
						util = append(util, td2.Text())
						cnt = 1
						sandesha := MessagesStructUtil{
							From:   util[0],
							Course: util[1],
							Msg:    util[2],
							Posted: util[3],
						}
						util = []string{}
						facmess = append(facmess, sandesha)
					}
				}

			})

		}
	}

	return &MessagesStruct{
		FacultyMsgs: facmess,
		Status:      status,
	}

}
