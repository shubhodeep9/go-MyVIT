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
	"sync"
)

type RefreshStruct struct {
	Timetable    *Timetable      `json:"timetable"`
	Academics    *AcademicStruct `json:"academic_istory"`
	Advisor      *Advisor        `json:"advisor"`
	Attendance   *Attendance     `json:"attendance"`
	ExamSchedule *ExamSchedule   `json:"exam_schedule"`
}

func Refresh(bow *browser.Browser, regno, password, baseuri string) *RefreshStruct {
	var re sync.WaitGroup
	re.Add(5)
	var (
		timet *Timetable
		acad  *AcademicStruct
		adv   *Advisor
		att   *Attendance
		exam  *ExamSchedule
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
	re.Wait()
	return &RefreshStruct{
		Timetable:    timet,
		Academics:    acad,
		Advisor:      adv,
		Attendance:   att,
		ExamSchedule: exam,
	}
}
