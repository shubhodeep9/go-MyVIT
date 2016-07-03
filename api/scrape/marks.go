/*
@Author Ujjwal Ayyangar
@Organization Google Developers Group VIT Vellore
Marks done!

*/

package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"math"
	"strconv"
	"strings"
	"sync"
)

type GetMarks struct {
	Status string          `json:"status"`
	Marks  map[string]Mrks `json:"marks"`
}

type Assessment struct {
	Title            string  `json:"title,omitempty"`
	Max_marks        int     `json:"max_marks,omitempty"`
	Weightage        int     `json:"weightage,omitempty"`
	Conducted_on     string  `json:"conducted_on,omitempty"`
	Status           string  `json:"status,omitempty"`
	ScoredMarks      float64 `json:"scored_marks"`
	ScoredPercentage float64 `json:"scored_percentage"`
}

type Mrks struct {
	Assessments       []Assessment `json:"assessments,omitempty"`
	Max_marks         int          `json:"max_marks,omitempty"`
	Max_percentage    int          `json:"max_percentage,omitempty"`
	Scored_Marks      float64      `json:"scored_marks,omitempty"`
	Scored_Percentage float64      `json:"scored_percentage,omitempty"`
}

/*
Function to show marks,
Calls NewLogin to login to academics,
@param bow (surf Browser) registration_no password
@return GetMarks struct
*/
func Value(inp string) (val float64) {
	if len(inp) == 0 {
		return 0
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

func ShowMarks(bow *browser.Browser, regno, password, baseuri string) *GetMarks {
	//conts := make(map[string]Contents)
	//type t make(map[string]Assessment)// "CAT: Assessment struct"
	type t2 []Assessment // t2 is a type for array of Assesments
	type m []Mrks
	rowdata := make(map[string]Mrks) // type :- "Subject name" : Array of Assesments of different exams
	status := "Success"
	if regno[1]-48 == 5 {

		if 1 != 1 {
			status = "Failure"
		} else {
			var wg sync.WaitGroup
			bow.Open(baseuri + "/student/marks.asp?sem=WS")

			//Twice loading due to Redirect policy defined by academics.vit.ac.in
			bow.Open(baseuri + "/student/marks.asp?sem=WS")
			tables := bow.Find("table")
			marks_table := tables.Eq(1)

			rows := marks_table.Find("tr")
			//tr_len := tr.Length()

			rows.Each(func(i int, s *goquery.Selection) {
				if i > 1 {
					wg.Add(1)

					go func() {
						defer wg.Done()
						td := s.Find("td") // all the columns of the row
						if td.Length() == 12 {
							fmarks := Value(td.Eq(6).Text())
							fmarksPer := (fmarks / 50) * 10
							cat1 := Assessment{
								Title:            "CAT-I",
								Max_marks:        50,
								Weightage:        10,
								Conducted_on:     "Check ExamSchedule",
								Status:           td.Eq(5).Text(),
								ScoredMarks:      fmarks,
								ScoredPercentage: toFixed(fmarksPer, 1),
							}
							fmarks2 := Value(td.Eq(8).Text())
							fmarks2Per := (fmarks2 / 50) * 10
							cat2 := Assessment{
								Title:            "CAT-II",
								Max_marks:        50,
								Weightage:        10,
								Conducted_on:     "Check ExamSchedule",
								Status:           td.Eq(7).Text(),
								ScoredMarks:      fmarks2,
								ScoredPercentage: toFixed(fmarks2Per, 1),
							}
							daMarks := Value(td.Eq(9).Text())
							daMarksPer := daMarks / 30
							da := Assessment{
								Title:            "Digital Assignment",
								Max_marks:        30,
								Weightage:        30,
								ScoredMarks:      daMarks,
								ScoredPercentage: toFixed(daMarksPer, 1),
							}
							fatmarks := Value(td.Eq(11).Text())
							fatPer := (fatmarks / 100) * 50
							fat := Assessment{
								Title:            "Final Assessment",
								Max_marks:        100,
								Weightage:        50,
								Conducted_on:     "Check ExamSchedule",
								Status:           td.Eq(10).Text(),
								ScoredMarks:      fatmarks,
								ScoredPercentage: toFixed(fatPer, 1),
							}
							Aments := []Assessment{cat1, cat2, da, fat}
							total := fatmarks + daMarks + fmarks2 + fmarks
							totalPer := fatPer + daMarksPer + fmarks2Per + fmarksPer
							rowdata[td.Eq(2).Text()] = Mrks{
								Assessments:       Aments,
								Max_marks:         230,
								Max_percentage:    100,
								Scored_Marks:      total,
								Scored_Percentage: toFixed(totalPer, 1),
							}

						} else if td.Length() == 8 {
							var title, code string
							if strings.Contains(td.Eq(4).Text(), "Lab") {
								code = td.Eq(2).Text() + "_L"
								title = "Lab_cam"
							} else {
								code = td.Eq(2).Text() + "_P"
								title = "Project"
							}
							score := Value(td.Eq(7).Text())
							other := Assessment{
								Title:            title,
								Max_marks:        100,
								Weightage:        50,
								Conducted_on:     "Tentative, set by faculty",
								Status:           td.Eq(6).Text(),
								ScoredMarks:      score,
								ScoredPercentage: score,
							}
							Aments := []Assessment{other}
							rowdata[code] = Mrks{
								Assessments:       Aments,
								Max_marks:         100,
								Max_percentage:    100,
								Scored_Marks:      score,
								Scored_Percentage: score,
							}

						} // else end

					}() // go func end

				} // i>2 end

			}) // go query end
			wg.Wait()

		}
	} else {

		if 1 != 1 {
			status = "Failure"
		} else {
			var wg sync.WaitGroup
			bow.Open(baseuri + "/student/marks.asp?sem=WS")
			//Twice loading due to Redirect policy defined by academics.vit.ac.in
			bow.Open(baseuri + "/student/marks.asp?sem=WS")
			tables := bow.Find("table")
			marks_table := tables.Eq(1)

			rows := marks_table.Find("tr")
			//tr_len := tr.Length()
			rows.Each(func(i int, s *goquery.Selection) {
				if i > 0 {
					wg.Add(1)

					go func() {
						defer wg.Done()
						td := s.Find("td") // all the columns of the row
						if td.Length() == 20 {
							fmarks := Value(td.Eq(6).Text())
							fmarksPer := (fmarks / 50) * 10
							cat1 := Assessment{
								Title:            "CAT-I",
								Max_marks:        50,
								Weightage:        15,
								Conducted_on:     "Check ExamSchedule",
								Status:           td.Eq(5).Text(),
								ScoredMarks:      fmarks,
								ScoredPercentage: toFixed(fmarksPer, 1),
							}
							fmarks2 := Value(td.Eq(8).Text())
							fmarks2Per := (fmarks2 / 50) * 10
							cat2 := Assessment{
								Title:            "CAT-II",
								Max_marks:        50,
								Weightage:        15,
								Conducted_on:     "Check ExamSchedule",
								Status:           td.Eq(7).Text(),
								ScoredMarks:      fmarks2,
								ScoredPercentage: toFixed(fmarks2Per, 1),
							}

							Q1marks := Value(td.Eq(10).Text())
							Q1marksPer := Q1marks
							quiz1 := Assessment{
								Title:            "Quiz-I",
								Max_marks:        5,
								Weightage:        5,
								Conducted_on:     "Check ExamSchedule",
								Status:           td.Eq(9).Text(),
								ScoredMarks:      Q1marks,
								ScoredPercentage: toFixed(Q1marksPer, 1),
							}

							Q2marks := Value(td.Eq(12).Text())
							Q2marksPer := Q1marks
							quiz2 := Assessment{
								Title:            "Quiz-II",
								Max_marks:        5,
								Weightage:        5,
								Conducted_on:     "Check ExamSchedule",
								Status:           td.Eq(11).Text(),
								ScoredMarks:      Q2marks,
								ScoredPercentage: toFixed(Q2marksPer, 1),
							}

							Q3marks := Value(td.Eq(14).Text())
							Q3marksPer := Q1marks
							quiz3 := Assessment{
								Title:            "Quiz-III",
								Max_marks:        5,
								Weightage:        5,
								Conducted_on:     "Check ExamSchedule",
								Status:           td.Eq(13).Text(),
								ScoredMarks:      Q3marks,
								ScoredPercentage: toFixed(Q3marksPer, 1),
							}
							daMarks := Value(td.Eq(16).Text())
							daMarksPer := daMarks / 30
							da := Assessment{
								Title:            "Assignment",
								Max_marks:        30,
								Weightage:        30,
								Status:           td.Eq(15).Text(),
								ScoredMarks:      daMarks,
								ScoredPercentage: toFixed(daMarksPer, 1),
							}
							fatmarks := Value(td.Eq(19).Text())
							fatPer := (fatmarks / 100) * 50
							fat := Assessment{
								Title:            "Final Assessment",
								Max_marks:        100,
								Weightage:        50,
								Conducted_on:     "Check ExamSchedule",
								Status:           td.Eq(18).Text(),
								ScoredMarks:      fatmarks,
								ScoredPercentage: toFixed(fatPer, 1),
							}
							Aments := []Assessment{cat1, cat2, quiz1, quiz2, quiz3, da, fat}
							total := fatmarks + daMarks + fmarks2 + fmarks + Q1marks + Q2marks + Q3marks
							totalPer := fatPer + daMarksPer + fmarks2Per + fmarksPer + Q1marksPer + Q2marksPer + Q3marksPer
							rowdata[td.Eq(2).Text()] = Mrks{
								Assessments:       Aments,
								Max_marks:         220,
								Max_percentage:    100,
								Scored_Marks:      total,
								Scored_Percentage: totalPer,
							}

						} else if td.Length() == 10 { // end of the 19 column if condition
							var title, code string
							if strings.Contains(td.Eq(4).Text(), "Lab") {
								code = td.Eq(2).Text() + "_L"
								title = "Lab_cam"
							}
							score := Value(td.Eq(7).Text())
							other := Assessment{
								Title:            title,
								Max_marks:        100,
								Weightage:        50,
								Conducted_on:     "Tentative, set by faculty",
								Status:           td.Eq(6).Text(),
								ScoredMarks:      score,
								ScoredPercentage: score,
							}
							Aments := []Assessment{other}
							rowdata[code] = Mrks{
								Assessments:       Aments,
								Max_marks:         100,
								Max_percentage:    100,
								Scored_Marks:      score,
								Scored_Percentage: score,
							}

						} // else end

					}() // go func end

				} // i>2 end

			}) // go query end
			wg.Wait()

		}

	}
	// end of  if condition for seniors/juniors
	return &GetMarks{
		Status: status,
		Marks:  rowdata,
	}
}
