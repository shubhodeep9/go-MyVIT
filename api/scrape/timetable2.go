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
	//"fmt"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/status"
	"strings"
	"os"
)

type Timetable2 struct {
	status status.StatusStruct `json:"status"`
	Tbl []Table `json:"timetable"`
}
type Table struct{
	Day string `json:"day"`
	Theory DayTheory `json:"theory"`
	Lab DayLab `json:"lab"`
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
	Ten    string `json:"17:50to18:40"`
	Eleven string `json:"19:00to19:50"`
	Twelve string `json:"20:00to20:50"`
}
type DayLab struct {
	One      string `json:"8:00to8:50"`
	Two      string `json:"8:50to9:40"`
	Three    string `json:"10:00to10:50"`
	Four     string `json:"10:50to11:40"`
	Five     string `json:"12:00to12:50"`
	Six      string `json:"12:40to13:30"`
	Seven    string `json:"14:00to14:50"`
	Eight    string `json:"14:50to15:40"`
	Nine     string `json:"16:00to16:50"`
	Ten      string `json:"16:50to17:40"`
	Eleven   string `json:"17:50to18:40"`
	Twelve   string `json:"18:40to19:30"`
	Thirteen string `json:"19:30to20:20"`
	Fourteen string `json:"20:20to21:10"`
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
func ShowTimetable2(bow *browser.Browser, reg, baseuri string, found bool) *Timetable2 {
	res := []Table{}
	//res2 := []DayLab{}

	sem := os.Getenv("SEM")
	//cnt := 1
	//sem = ""
	stat := status.Success()
	if !found {
		stat = status.SessionError()
	} else {
		bow.Open(baseuri + "/student/course_regular.asp?sem="+sem)
		//Twice loading due to Redirect policy defined by academics.vit.ac.in
		if bow.Open(baseuri+"/student/course_regular.asp?sem="+sem) == nil {
			table := bow.Find("table[width='95%']")
			table.Find("tr").Each(func(i int, s *goquery.Selection) {
				if i > 1 {
					td := s.Find("td")
					var temp Table			
					var temp1 DayTheory
					var temp2 DayLab
					temp.Day=td.Eq(0).Text()
					//temp1.Day = td.Eq(0).Text()
					//temp2.Day = td.Eq(0).Text()
					for x := 1; x <= 14; x++ {
						a := toKey(x)
						if strings.Contains(td.Eq(x).Text(), "ETH") ||  strings.Contains(td.Eq(x).Text(),"SS") && len(td.Eq(x).Text())!=0 {

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

						} else if strings.Contains(td.Eq(x).Text(), "ELA") && len(td.Eq(x).Text())!=0 {
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
					//fmt.Println(temp1)
					//fmt.Println(temp2)
					temp.Theory = temp1
					temp.Lab = temp2
					res = append(res, temp)
					//res2 = append(res2, temp2)

				}

			})

		}
	}

	return &Timetable2{
		status: stat,
		Tbl:res,
	}

}
