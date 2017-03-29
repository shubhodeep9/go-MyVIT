//https://vtop.vit.ac.in/student/profile_education_view.asp
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
	"go-MyVIT/api/status"
	"os"
)

type EduDet struct {
	QualifyingExam string `json:"qualifyingExam"`
	School         string `json:"school"`
	Medium         string `json:"medium"`
	BoardName      string `json:"boardName"`
	Regno          string `json:"regno"`
	ClassObtained  string `json:"classObtained"`
	PassingDate    string `json:"passingDate"`
	BreakOS        string `json:"breakOfStudy"`
	ReasonBOS      string `json:"reasonBOS"`
}
type SchoolAddrress struct {
	Area    string `json:"areaName"`
	City    string `json:"cityName"`
	State   string `json:"state"`
	Pincode string `json:"pincode"`
	PhoneNo string `json:phoneNo"`
}

type EducationalDetailsStruct struct {
	Status       status.StatusStruct `json:"status"`
	EducationDet EduDet              `json:"educationalDetails"`
	SchoolAdd    SchoolAddrress      `json:"schoolAddress"`
}

/*
Function to show educational details,
Calls NewLogin to login to academics,
@param bow (surf Browser) registration_no password
@return ShowEducationalDetails struct
*/
func ShowEducationalDetails(bow *browser.Browser, reg, baseuri string, found bool) *EducationalDetailsStruct {
	var (
		edudet EduDet
		saddr  SchoolAddrress
	)

	sem := os.Getenv("SEM")
	cnt := 1
	sem = ""
	stat := status.Success()
	if !found {
		stat = status.SessionError()
	} else {
		util := []string{}
		bow.Open(baseuri + "/student/profile_education_view.asp?sem=" + sem)
		//Twice loading due to Redirect policy defined by academics.vit.ac.in
		if bow.Open(baseuri+"/student/profile_education_view.asp?sem="+sem) == nil {
			table := bow.Find("table[width='600']")
			table.Find("tr").Each(func(i int, s *goquery.Selection) {
				td := s.Find("td")
				if cnt > 1 && cnt <= 10 {
					util = append(util, td.Eq(1).Text())
					if cnt == 10 {
						edudet = EduDet{
							QualifyingExam: util[0],
							School:         util[1],
							Medium:         util[2],
							BoardName:      util[3],
							Regno:          util[4],
							ClassObtained:  util[5],
							PassingDate:    util[6],
							BreakOS:        util[7],
							ReasonBOS:      util[8],
						}
						util = []string{}

					}

				} else if cnt > 11 && cnt <= 16 {
					util = append(util, td.Eq(1).Text())
					if cnt == 16 {
						saddr = SchoolAddrress{
							Area:    util[0],
							City:    util[1],
							State:   util[2],
							Pincode: util[3],
							PhoneNo: util[4],
						}
					}

				}
				cnt += 1

			})

		}
	}

	return &EducationalDetailsStruct{
		Status:       stat,
		EducationDet: edudet,
		SchoolAdd:    saddr,
	}

}
