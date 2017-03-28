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

type FamilyStruct struct {
	Name            string `json:"name"`
	Qualification   string `json:"qualification"`
	Occupation      string `json:"occupation"`
	Organization    string `json:"organization"`
	EmpId           string `json:"empid"`
	MobileNo        string `json:"mobileNo"`
	EmailID         string `json:"email"`
	AnnualIncome    string `json:"annualIncome"`
	OfficialAddress string `json:"officalAddress"`
	CityName        string `json:"cityName"`
	StateName       string `json:"stateName"`
	Pincode         string `json:"pincode"`
	PhoneNo         string `json:"phoneNo"`
}

type FamilyDetailsStruct struct {
	Status status.StatusStruct `json:"status"`
	Dad    FamilyStruct        `json:"father"`
	Mom    FamilyStruct        `json:"mother"`
}

/*
Function to show family details,
Calls NewLogin to login to academics,
@param bow (surf Browser) registration_no password
@return FamilyDetailsStruct struct
*/
func ShowFamilyDetails(bow *browser.Browser, reg, baseuri string, found bool) *FamilyDetailsStruct {
	var (
		dad FamilyStruct
		mom FamilyStruct
	)
	sem := os.Getenv("SEM")
	sem = ""
	cnt := 1
	sem = ""
	stat := status.Success()
	if !found {
		stat = status.SessionError()
	} else {
		util := []string{}
		bow.Open(baseuri + "/student/profile_family_view.asp?sem=" + sem)
		//Twice loading due to Redirect policy defined by academics.vit.ac.in
		if bow.Open(baseuri+"/student/profile_family_view.asp?sem="+sem) == nil {
			table := bow.Find("table[width='600']")
			table.Find("tr").Each(func(i int, s *goquery.Selection) {
				td := s.Find("td")
				if cnt > 1 && cnt <= 14 {
					util = append(util, td.Eq(1).Text())
					if cnt == 14 {
						dad = FamilyStruct{
							Name:            util[0],
							Qualification:   util[1],
							Occupation:      util[2],
							Organization:    util[3],
							EmpId:           util[4],
							MobileNo:        util[5],
							EmailID:         util[6],
							AnnualIncome:    util[7],
							OfficialAddress: util[8],
							CityName:        util[9],
							StateName:       util[10],
							Pincode:         util[11],
							PhoneNo:         util[12],
						}
						util = []string{}
					} else if cnt > 15 && cnt <= 28 {
						util = append(util, td.Eq(1).Text())
						if cnt == 28 {
							mom = FamilyStruct{
								Name:            util[0],
								Qualification:   util[1],
								Occupation:      util[2],
								Organization:    util[3],
								EmpId:           util[4],
								MobileNo:        util[5],
								EmailID:         util[6],
								AnnualIncome:    util[7],
								OfficialAddress: util[8],
								CityName:        util[9],
								StateName:       util[10],
								Pincode:         util[11],
								PhoneNo:         util[12],
							}

						}
					}
				}

				cnt += 1
			})

		}
	}

	return &FamilyDetailsStruct{
		Status: stat,
		Dad:    dad,
		Mom:    mom,
	}

}
