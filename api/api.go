package api

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/jar"
	"go-MyVIT/api/login"
	"go-MyVIT/api/scrape"
)

//Executable script to Login
func LogIn(regno, password string) *login.Response {
	bow := surf.NewBrowser()
	bow.SetAttribute(browser.FollowRedirects, true)
	bow.SetAttribute(browser.SendReferer, true)
	bow.SetCookieJar(jar.NewMemoryCookies())
	bow.Open("https://academics.vit.ac.in/student/stud_login.asp")
	return login.NewLogin(bow, regno, password)
}

//Executable script to show timetable
func TimeTable(regno, password string) *scrape.Timetable {
	bow := surf.NewBrowser()
	bow.SetAttribute(browser.FollowRedirects, true)
	bow.SetAttribute(browser.SendReferer, true)
	bow.SetCookieJar(jar.NewMemoryCookies())
	bow.Open("https://academics.vit.ac.in/student/stud_login.asp")
	return scrape.ShowTimetable(bow,regno,password)
}


