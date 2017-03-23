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

type general struct {
	Name           string `json:'name'`
	DOB            string `json:'dob'`
	Gender         string `json:'gender'`
	NativeLanguage string `json:'nativelang'`
	Nationality    string `json:'nationality'`
	BloodGroup     string `json:'bloodGroup'`
	VitMail        string `json:'vitmail'`
	Hosteler       string `json:'hosteler'`
}
type hostelAddr struct {
	BlockName string `json:'blockname'`
	Room      string `json:'room'`
	Mess      string `json:'mess'`
}

type school struct {
	RegisterNumber string `json:'regno'`
	AtmNo          string `json:'atmno'`
	School         string `json:'school'`
	Prog           string `json:'programme'`
}

type permanentAddr struct {
	Street       string `json:'street'`
	Area         string `json:'area'`
	City         string `json:'city'`
	Pincode      string `json:'pincode'`
	State        string `json:'state'`
	Country      string `json:'country'`
	PhoneNumber  string `json:'phoneno'`
	MobileNumber string `json:'mobileno'`
	EmailID      string `json:'emailid'`
}

type PersonalDetailsStruct struct {
	Status           string        `json:"status"`
	General          general       `json:"general"`
	Hostel           hostelAddr    `json:"hostel"`
	School           school        `json:"school"`
	PermanentAddress permanentAddr `json:"permanentAddress"`
}

/*
Function to show personal details,
Calls NewLogin to login to academics,
@param bow (surf Browser) registration_no password
@return Attendance struct
*/
func ShowPersonalDetails(bow *browser.Browser, reg, baseuri string, found bool) *PersonalDetailsStruct {
	sem := os.Getenv("SEM")
	status := "Success"
	var (
		g  general
		h  hostelAddr
		so school
		p  permanentAddr
	)
	cnt := 1
	if !found {
		status = "Failure"
	} else {

		util := []string{}
		bow.Open(baseuri + "/student/profile_personal_view.asp?sem=" + sem)
		//Twice loading due to Redirect policy defined by academics.vit.ac.in
		if bow.Open(baseuri+"/student/profile_personal_view.asp?sem="+sem) == nil {
			table := bow.Find("table[width='720']")
			table.Find("tr").Each(func(i int, s *goquery.Selection) {
				td := s.Find("td")
				if cnt > 1 && cnt <= 9 {
					util = append(util, td.Eq(1).Text())
					if cnt == 9 {
						g = general{
							Name:           util[0],
							DOB:            util[1],
							Gender:         util[2],
							NativeLanguage: util[3],
							Nationality:    util[4],
							BloodGroup:     util[5],
							VitMail:        util[6],
							Hosteler:       util[7],
						}
						util = []string{}
					}
				} else if cnt > 10 && cnt <= 13 {
					util = append(util, td.Eq(1).Text())
					if cnt == 13 {
						h = hostelAddr{
							BlockName: util[0],
							Room:      util[1],
							Mess:      util[2],
						}
						util = []string{}
					}
				} else if cnt > 14 && cnt <= 18 {
					util = append(util, td.Eq(1).Text())
					if cnt == 18 {
						so = school{RegisterNumber: util[0], AtmNo: util[1], School: util[2], Prog: util[3]}
						//fmt.Println(util[0], util[1], util[2], util[3])
						util = []string{}
					}
				} else if cnt > 19 && cnt <= 28 {
					util = append(util, td.Eq(1).Text())
					if cnt == 28 {
						p = permanentAddr{
							Street:       util[0],
							Area:         util[1],
							City:         util[2],
							Pincode:      util[3],
							State:        util[4],
							Country:      util[5],
							PhoneNumber:  util[6],
							MobileNumber: util[7],
							EmailID:      util[8],
						}
						util = []string{}

					}
				}

				//fmt.Println(td.Text(), cnt)
				cnt += 1
			})

			//fmt.Println(bow.Body())

		}
	}
	/*else if cnt > 14 && cnt <= 18 {
		util = append(util, td.Eq(1).Text())
		if cnt == 18 {
			s = schoolInfo{
				RegisterNumber: util[0],
				AtmNo:          util[1],
				School:         util[2],
				Programme:      util[3],
			}
			util = []string{}
		}
	} */

	return &PersonalDetailsStruct{
		Status:           status,
		General:          g,
		Hostel:           h,
		School:           so,
		PermanentAddress: p,
	}

}
