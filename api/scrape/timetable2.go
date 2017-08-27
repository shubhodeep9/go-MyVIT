//Hostel_LAB_Permission.asp
//https://vtop.vit.ac.in/student/student_history.asp
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

type Timetable2 struct {
	status string  `json:"status"`
	Tbl    []Table `json:"timetable"`
}
type Table struct {
	Day    string    `json:"day"`
	Theory DayTheory `json:"theory"`
	Lab    DayLab    `json:"lab"`
}

type DayTheory struct {
	One    string `json:"8:00to8:50"`
	Two    string `json:"9:00to9:50"`
	Three  string `json:"10:00to10:50"`
	Four   string `json:"11:00to11:50"`
	Five   string `json:"12:00to12:50"`
	Six    string `json:"14:00to14:50"`
	Seven  string `json:"15:00to15:50"`
	Eight  string `json:"16:00to16:50"`
	Nine   string `json:"17:00to17:50"`
	Ten    string `json:"18:00to18:50"`
	Eleven string `json:"19:00to19:50"`
	Twelve string `json:"20:00to20:50"`
}
type DayLab struct {
	One      string `json:"8:00to8:50"`
	Two      string `json:"8:51to9:40"`
	Three    string `json:"10:00to10:50"`
	Four     string `json:"10:51to11:40"`
	Five     string `json:"11:50to12:40"`
	Six      string `json:"12:41to13:30"`
	Seven    string `json:"14:00to14:50"`
	Eight    string `json:"14:51to15:40"`
	Nine     string `json:"16:00to16:50"`
	Ten      string `json:"16:51to17:40"`
	Eleven   string `json:"17:50to18:40"`
	Twelve   string `json:"18:41to19:30"`
	Thirteen string `json:"19:31to20:20"`
	Fourteen string `json:"20:21to21:10"`
}

func toKey(n int) string {
	switch n {
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	case 10:
		return "Ten"
	case 11:
		return "Eleven"
	case 12:
		return "Twelve"
	case 13:
		return "Thirteen"
	case 14:
		return "Fourteen"

	}
	return ""

}

/*
Function to show academic history,
Calls NewLogin to login to academics,
@param bow (surf Browser) registration_no password
@return AcademicHistory struct
*/
func ShowTimetable2(client http.Client, regNo, psswd, baseuri string) *Timetable2 {
	res := []Table{}
	PostData3 := strings.NewReader("semesterSubId=VL2017181")
	req3, _ := http.NewRequest("POST", "https://vtopbeta.vit.ac.in/vtop/processViewTimeTable", PostData3)
	req3.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req3.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Max OS X 10_10_5) AppleWebKit (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")
	//fmt.Println("Post Data for the time table = ",PostData3)

	resp, err := client.Do(req3)
	var status string
	if err != nil {
		status = "Failure"
	} else {
		status = "Success"
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	html := string(body)
	//fmt.Println("html is", html)

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader((html)))

	table := doc.Find("table#timeTableStyle")
	trow := table.Find("tr")
	var temp Table
	var temp2 DayLab
	var temp1 DayTheory
	trow.Each(func(i int, td *goquery.Selection) {
		td = td.Find("td")
		if i > 3 && i%2 == 0 {
			temp1 = DayTheory{}
			temp.Day = td.Eq(0).Text()
			lC := 0 // used for not taking lunch
			for x := 2; x <= 16; x++ {
				if x > 8 {
					lC = 1
				} else {
					lC = 0
				}
				a := toKey(x - lC - 2)
				if len(td.Eq(x).Text()) > 4 && x != 8 {
					switch a {
					case "One":
						temp1.One = td.Eq(x).Text()
					case "Two":
						temp1.Two = td.Eq(x).Text()
					case "Three":
						temp1.Three = td.Eq(x).Text()
					case "Four":
						temp1.Four = td.Eq(x).Text()
					case "Five":
						temp1.Five = td.Eq(x).Text()
					case "Six":
						temp1.Six = td.Eq(x).Text()
					case "Seven":
						temp1.Seven = td.Eq(x).Text()
					case "Eight":
						temp1.Eight = td.Eq(x).Text()
					case "Nine":
						temp1.Nine = td.Eq(x).Text()
					case "Ten":
						temp1.Ten = td.Eq(x).Text()
					case "Eleven":
						temp1.Eleven = td.Eq(x).Text()
					case "Twelve":
						temp1.Twelve = td.Eq(x).Text()

					}
				}
			}

			temp.Theory = temp1

		} else {
			temp2 = DayLab{}
			lC := 0 // used for not taking lunch
			for x := 1; x <= 15; x++ {
				if x > 7 {
					lC = 1
				} else {
					lC = 0
				}
				a := toKey(x - lC)
				if len(td.Eq(x).Text()) > 4 && x != 7 {
					switch a {
					case "One":
						temp2.One = td.Eq(x).Text()
					case "Two":
						temp2.Two = td.Eq(x).Text()
					case "Three":
						temp2.Three = td.Eq(x).Text()
					case "Four":
						temp2.Four = td.Eq(x).Text()
					case "Five":
						temp2.Five = td.Eq(x).Text()
					case "Six":
						temp2.Six = td.Eq(x).Text()
					case "Seven":
						temp2.Seven = td.Eq(x).Text()
					case "Eight":
						temp2.Eight = td.Eq(x).Text()
					case "Nine":
						temp2.Nine = td.Eq(x).Text()
					case "Ten":
						temp2.Ten = td.Eq(x).Text()
					case "Eleven":
						temp2.Eleven = td.Eq(x).Text()
					case "Twelve":
						temp2.Twelve = td.Eq(x).Text()
					case "Thirteen":
						temp2.Thirteen = td.Eq(x).Text()
					case "Fourteen":
						temp2.Fourteen = td.Eq(x).Text()

					}
				}
			}
			if i > 3 {
				temp.Lab = temp2
				res = append(res, temp)
			}
		}

	})

	return &Timetable2{
		status: status,
		Tbl:    res,
	}

}
