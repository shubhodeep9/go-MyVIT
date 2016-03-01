package api

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf"
	"go-MyVIT/api/login"
	"go-MyVIT/api/scrape"
)

//Executable script to Login
func LogIn(regno, password string) *login.Response {
	bow := surf.NewBrowser()
	return login.NewLogin(bow, regno, password)
}

//Executable script to show timetable
func TimeTable(regno, password string) *scrape.Timetable {
	bow := surf.NewBrowser()
	return scrape.ShowTimetable(bow,regno,password)
}


