/*
@Author Ujjwal Ayyangar
@Organization Google Developers Group VIT Vellore
	I err, therefore I am
*/

package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"strings"
)

type general struct {
	AppNo          string `json:"ApplNum"`
	Name           string `json:"name"`
	DOB            string `json:"dob"`
	Gender         string `json:"gender"`
	NativeLanguage string `json:"native_lang"`
	NativeState    string `json:"native_state"`
	Nationality    string `json:"nationality"`
	BloodGroup     string `json:"bloodGroup"`
	VitMail        string `json:"vitmail"`
	Hosteler       string `json:"hosteler"`
	PhysicChal     string `json:"phy_challenged"`
	Community      string `json:"community"`
	Religion       string `json:"religion"`
	Caste          string `json:"caste"`
	AadharNumber   string `json:"aadhar_num"`
}

type school struct {
	RegisterNumber string `json:"regno"`
	School         string `json:"school"`
	Branch         string `json:"branch"`
	Prog           string `json:"programme"`
}

type permanentAddr struct {
	Street       string `json:"street"`
	Area         string `json:"area"`
	City         string `json:"city"`
	Pincode      string `json:"pincode"`
	State        string `json:"state"`
	Country      string `json:"country"`
	PhoneNumber  string `json:"phoneno"`
	MobileNumber string `json:"mobileno"`
	EmailID      string `json:"email_id"`
	FriendMob    string `json:"friend_mob"`
}

type PersonalDetailsStruct struct {
	Status           string        `json:"status"`
	General          general       `json:"general"`
	School           school        `json:"school"`
	PermanentAddress permanentAddr `json:"permanentAddress"`
}

/*
Function to show personal details,
Calls NewLogin to login to academics,
@param bow (surf Browser) registration_no password
@return PersonalDetails struct
*/
func ShowPersonalDetails(client http.Client, regNo, psswd, baseuri string) *PersonalDetailsStruct {
	PersonalDetData := strings.NewReader("")
	reqPer, _ := http.NewRequest("POST", "https://vtopbeta.vit.ac.in/vtop/studentsRecord/SearchRegnoStudent", PersonalDetData)
	reqPer.Header.Add("Content-Type", "text/html;charset=UTF-8")
	reqPer.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Max OS X 10_10_5) AppleWebKit (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")

	resp, err := client.Do(reqPer)
	var status string
	if err != nil {
		status = "Failure"
	} else {
		status = "Success"
	}

	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	html := string(body)
	count := 0
	var (
		g  general
		so school
		p  permanentAddr
	)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader((html)))
	table := doc.Find("#1a .table")
	trow := table.Find("tr")
	temp := []string{}
	trow.Each(func(i int, td *goquery.Selection) {
		td = td.Find("td")
		if td.Length() == 1 {
			count += 1
			if count > 1 {
				if count == 2 {
					g = general{
						AppNo:          temp[0],
						Name:           temp[1],
						DOB:            temp[2],
						Gender:         temp[3],
						NativeLanguage: temp[4],
						NativeState:    temp[5],
						BloodGroup:     temp[6],
						PhysicChal:     temp[7],
						Community:      temp[8],
						Religion:       temp[9],
						Caste:          temp[10],
						Nationality:    temp[11],
						Hosteler:       temp[12],
						AadharNumber:   temp[13],
					}
				} else if count == 3 {
					so = school{
						RegisterNumber: temp[0],
						Prog:           temp[1],
						Branch:         temp[2],
						School:         temp[3],
					}
				}

				temp = []string{}
			}
		} else {
			temp = append(temp, trim(td.Eq(1).Text()))
		}
	})
	p = permanentAddr{
		Street:       temp[0],
		Area:         temp[1],
		City:         temp[2],
		State:        temp[3],
		Country:      temp[4],
		Pincode:      temp[5],
		PhoneNumber:  temp[6],
		MobileNumber: temp[7],
		EmailID:      temp[8],
		FriendMob:    temp[9],
	}

	return &PersonalDetailsStruct{
		Status:           status,
		General:          g,
		School:           so,
		PermanentAddress: p,
	}

}
