/*
@Author Ujjwal Ayyangar
@Organization Google Developers Group VIT Vellore
Marks done!

*/

package scrape

import (
	//"fmt"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"math"
	"os"
	"strconv"
	"strings"
	//"sync"
)

type GetMarks struct {
	Status string          `json:"status"`
	Marks  map[string]Mrks `json:"marks"`
}

type Assessment struct {
	Title            string  `json:"title"`
	Max_marks        int     `json:"max_marks"`
	Weightage        int     `json:"weightage"`
	Conducted_on     string  `json:"conducted_on"`
	Status           string  `json:"status"`
	ScoredMarks      float64 `json:"scored_marks"`
	ScoredPercentage float64 `json:"scored_percentage"`
}

type Mrks struct {
	Assessments       []Assessment `json:"assessments"`
	Max_marks         int          `json:"max_marks"`
	Max_percentage    int          `json:"max_percentage"`
	Scored_Marks      float64      `json:"scored_marks"`
	Scored_Percentage float64      `json:"scored_percentage"`
}

/*
Function to show marks,
Calls NewLogin to login to academics,
@param bow (surf Browser) registration_no password
@return GetMarks struct
*/
func Value(inp string) (val float64) {
	if len(inp) == 0 {
		return 0.0
	} else {
		ret, _ := strconv.ParseFloat(inp, 64)
		return ret
	}
}
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func ShowMarks(bow *browser.Browser, regno, baseuri string) *GetMarks {
	sem := os.Getenv("SEM")
	rowdata := make(map[string]Mrks)
	status := "Success"
	bow.Open(baseuri + "/student/marks.asp?sem=" + sem)
	if bow.Open(baseuri+"/student/marks.asp?sem="+sem) == nil {

		tables := bow.Find("table")

		/**
		* Store the subjects
		 */
		tables.Eq(1).Find("tr[bgcolor='#EDEADE']").Each(func(i int, s *goquery.Selection) {
			td := s.Find("td")
			code := td.Eq(2).Text()
			if strings.Contains(td.Eq(4).Text(), "Lab") {
				code = code + "_L"
			}
			var assessment_list []Assessment
			marksection := s.Next().Find("table").Find("tr[bgcolor='#CCCCCC']")
			totalmaxmarks := 0
			totalweight := 0
			totalscored := float64(0)
			totalscoredper := float64(0)
			marksection.Each(func(j int, sel *goquery.Selection) {
				seltd := sel.Find("td")
				maxmark, _ := strconv.Atoi(seltd.Eq(2).Text())
				weight, _ := strconv.Atoi(seltd.Eq(3).Text())
				scored := Value(seltd.Eq(5).Text())
				per := Value(seltd.Eq(6).Text())
				assess := Assessment{
					Title:            seltd.Eq(1).Text(),
					Max_marks:        maxmark,
					Weightage:        weight,
					Conducted_on:     "Check ExamSchedule",
					Status:           seltd.Eq(4).Text(),
					ScoredMarks:      scored,
					ScoredPercentage: per,
				}
				totalmaxmarks = totalmaxmarks + maxmark
				totalweight = totalweight + weight
				totalscored = totalscored + scored
				totalscoredper = totalscoredper + per
				assessment_list = append(assessment_list, assess)
			})
			rowdata[code] = Mrks{
				Assessments:       assessment_list,
				Max_marks:         totalmaxmarks,
				Max_percentage:    totalweight,
				Scored_Marks:      totalscored,
				Scored_Percentage: totalscoredper,
			}
		})
	}
	// end of  if condition for seniors/juniors
	return &GetMarks{
		Status: status,
		Marks:  rowdata,
	}
}
