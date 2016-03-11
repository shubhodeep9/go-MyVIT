/*
@Author Shubhodeep Mukherjee
@Organization Google Developers Group VIT Vellore
	Isn't Go sexy?
	I know right!!
	Just like Shubhodeep
	I mean, have you seen the guy? xP
	#GDGSwag
*/


package api

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf"
	"go-MyVIT/api/login"
	"go-MyVIT/api/scrape"
)

//Executable script to Login
func LogIn(regno, password, baseuri string) *login.Response {
	bow := surf.NewBrowser()
	return login.NewLogin(bow, regno, password, baseuri)
}

//Executable script to show timetable
func TimeTable(regno, password, baseuri string) *scrape.Timetable {
	bow := surf.NewBrowser()
	return scrape.ShowTimetable(bow,regno,password, baseuri)
}

//Executable script to show Faculty Advisor details
func Advisor(regno, password, baseuri string) *scrape.Advisor{
	bow := surf.NewBrowser()
	return scrape.FacultyAdvisor(bow,regno,password, baseuri)
}

//Executable script to show Attendance
func Attendance(regno, password, baseuri string) *scrape.Attendance{
	bow := surf.NewBrowser()
	return scrape.ShowAttendance(bow,regno,password,baseuri)
}
func Schedule(regno, password, baseuri string) *scrape.ExamSchedule{
	bow := surf.NewBrowser()
	return scrape.ExmSchedule(bow,regno,password,baseuri)
}


