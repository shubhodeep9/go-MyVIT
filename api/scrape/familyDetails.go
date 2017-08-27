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
	"io/ioutil"
	"net/http"
	"strings"
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
	Designation     string `json:"designation"`
	GuardianInfo    string `json:"guardian_info"`
}
type GeneralFamStruct struct {
	Nob      string `json:"NumOfBros"`
	Nos      string `json:"NumOfSis"`
	SibVit   string `json:"BorS_InVit"`
	StudDet  string `json:"studying_details"`
	SibVited string `json:"BorS_wasInVit"`
	StudDet2 string `json:"studying_details2"`
}

type FamilyDetailsStruct struct {
	Status  string           `json:"status"`
	Dad     FamilyStruct     `json:"father"`
	Mom     FamilyStruct     `json:"mother"`
	General GeneralFamStruct `json:"general"`
}

/*
Function to show family details,
Calls NewLogin to login to academics,
@param bow (surf Browser) registration_no password
@return FamilyDetailsStruct struct
*/
func ShowFamilyDetails(client http.Client, regNo, psswd, baseuri string) *FamilyDetailsStruct {

	dummyData := strings.NewReader("")
	reqFam, _ := http.NewRequest("POST", "https://vtopbeta.vit.ac.in/vtop/studentsRecord/SearchRegnoStudent", dummyData)
	reqFam.Header.Add("Content-Type", "text/html;charset=UTF-8")
	reqFam.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Max OS X 10_10_5) AppleWebKit (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")

	resp, err := client.Do(reqFam)
	var status string
	if err != nil {
		status = "Failure"
	} else {
		status = "Success"
	}

	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	html := string(body)
	var (
		dad FamilyStruct
		mom FamilyStruct
		gen GeneralFamStruct
	)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader((html)))
	table := doc.Find("#3a .table")
	trow := table.Find("tr")
	util := []string{}
	trow.Each(func(i int, td *goquery.Selection) {
		td = td.Find("td")
		if i > 0 && i <= 6 {
			util = append(util, td.Eq(1).Text())
			if i == 6 {
				gen = GeneralFamStruct{
					Nob:      trim(util[0]),
					Nos:      trim(util[1]),
					SibVit:   trim(util[2]),
					StudDet:  trim(util[3]),
					SibVited: trim(util[4]),
					StudDet2: trim(util[5]),
				}
				util = []string{}

			}
		} else if i > 7 && i <= 21 {
			util = append(util, td.Eq(1).Text())
			if i == 21 {
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
					Designation:     util[13],
				}
				util = []string{}
			}
		} else if i > 22 && i <= 37 {
			util = append(util, td.Eq(1).Text())
			if i == 37 {
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
					Designation:     util[13],
					GuardianInfo:    util[14],
				}

			}
		}

	})

	return &FamilyDetailsStruct{
		Status:  status,
		Dad:     dad,
		Mom:     mom,
		General: gen,
	}

}
