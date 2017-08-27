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

type EduDet struct {
	ApplDeg  string `json:"applDeg"`
	EduQual  string `json:"edu_qual"`
	Branch   string `json:"group_studied"`
	SchoName string `json:"school_name"`
	Medium   string `json:"medium"`
	Board    string `json:"board"`
	RegNo    string `json:"regNo"`
	Class    string `json:"classObtained"`
	Year     string `json:"passYear"`
	Month    string `json:"passMonth"`
}
type SchoolAddrress struct {
	Area    string `json:"areaName"`
	City    string `json:"cityName"`
	State   string `json:"state"`
	Pincode string `json:"pincode"`
	PhoneNo string `json:phoneNo"`
	Break   string `json:"break"`
	Reason  string `json:"reason"`
}

type EducationalDetailsStruct struct {
	Status       string         `json:"status"`
	EducationDet EduDet         `json:"educationalDetails"`
	SchoolAdd    SchoolAddrress `json:"schoolAddress"`
}

/*
Function to show educational details,
Calls NewLogin to login to academics,
@param bow (surf Browser) registration_no password
@return ShowEducationalDetails struct
*/
func ShowEducationalDetails(client http.Client, regNo, psswd, baseuri string) *EducationalDetailsStruct {
	dummyData := strings.NewReader("")
	reqEdu, _ := http.NewRequest("POST", "https://vtopbeta.vit.ac.in/vtop/studentsRecord/SearchRegnoStudent", dummyData)
	reqEdu.Header.Add("Content-Type", "text/html;charset=UTF-8")
	reqEdu.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Max OS X 10_10_5) AppleWebKit (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")

	resp, err := client.Do(reqEdu)
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
		e EduDet
		s SchoolAddrress
	)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader((html)))
	table := doc.Find("#2a .table")
	trow := table.Find("tr")
	util := []string{}
	trow.Each(func(i int, td *goquery.Selection) {
		td = td.Find("td")
		if i > 0 && i <= 10 {
			util = append(util, td.Eq(1).Text())
			if i == 10 {
				e = EduDet{
					ApplDeg:  trim(util[0]),
					EduQual:  trim(util[1]),
					Branch:   trim(util[2]),
					SchoName: trim(util[3]),
					Medium:   trim(util[4]),
					Board:    trim(util[5]),
					RegNo:    trim(util[6]),
					Class:    trim(util[7]),
					Year:     trim(util[8]),
					Month:    trim(util[9]),
				}
				util = []string{}

			}

		} else if i > 11 && i <= 18 {
			util = append(util, td.Eq(1).Text())
			if i == 18 {
				s = SchoolAddrress{
					Area:    util[0],
					City:    util[1],
					State:   util[2],
					Pincode: util[3],
					PhoneNo: util[4],
					Break:   util[5],
					Reason:  util[6],
				}
			}

		}

	})

	return &EducationalDetailsStruct{
		Status:       status,
		EducationDet: e,
		SchoolAdd:    s,
	}

}
