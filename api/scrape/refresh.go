/*
@Author Shubhodeep Mukherjee & Ujjwal Ayyangar
@Organization Google Developers Group VIT Vellore
	Isn't Go sexy?
	I know right!!
	Just like Shubhodeep
	I mean, have you seen the guy? xP
	#GDGSwag
*/

package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/login"
	"sync"
)

type RefreshStruct struct {
	RegNo    string             `json:"reg_no"`
	Campus   string             `json:"campus"`
	Semester string             `json:"semester"`
	Courses  []Contents         `json:"courses"`
	Status   login.StatusStruct `json:"status"`
}

func Refresh(bow *browser.Browser, regno, password, baseuri string) *RefreshStruct {
	var re sync.WaitGroup
	re.Add(6)
	var (
		timet *Timetable
		acad  *AcademicStruct
		adv   *Advisor
		att   *Attendance
		exam  *ExamSchedule
		marks *GetMarks
	)
	go func() {
		defer re.Done()
		timet = ShowTimetable(bow, regno, password, baseuri)
	}()
	go func() {
		defer re.Done()
		acad = Academics(bow, regno, password, baseuri)
	}()
	go func() {
		defer re.Done()
		adv = FacultyAdvisor(bow, regno, password, baseuri)
	}()
	go func() {
		defer re.Done()
		att = ShowAttendance(bow, regno, password, baseuri)
	}()
	go func() {
		defer re.Done()
		exam = ExmSchedule(bow, regno, password, baseuri)
	}()
	go func() {
		defer re.Done()
		marks = ShowMarks(bow, regno, password, baseuri)
	}()
	re.Wait()
	var courses []Contents
	timetable := timet.Time_table
	var course Contents
	var timings []TimeStruct
	time := TimeStruct{
		Day:       0,
		StartTime: "03:30:00Z",
		EndTime:   "03:30:00Z",
	}
	timings = append(timings, time)
	for i := range timetable {
		course = Contents{
			Class_number:        timetable[i].Class_number,
			Course_code:         timetable[i].Course_code,
			Course_mode:         timetable[i].Course_mode,
			Course_option:       timetable[i].Course_option,
			Course_title:        timetable[i].Course_title,
			Course_type:         timetable[i].Course_type,
			Faculty:             timetable[i].Faculty,
			Ltpjc:               timetable[i].Ltpjc,
			Registration_status: timetable[i].Registration_status,
			Slot:                timetable[i].Slot,
			Venue:               timetable[i].Venue,
			Timings:             timings,
			Attendance:          att.AttendanceDet[i],
			Marks:               marks.Marks[i],
		}
		courses = append(courses, course)
	}
	stt := login.StatusStruct{
		Message: "Successful Execution",
		Code:    0,
	}
	return &RefreshStruct{
		RegNo:    regno,
		Campus:   "vellore",
		Semester: "WS",
		Courses:  courses,
		Status:   stt,
	}
}
