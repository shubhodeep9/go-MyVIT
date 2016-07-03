/*
@Author Shubhodeep Mukherjee and Ujjwal Ayyangar
@Organization Google Developers Group VIT Vellore
	Isn't Go sexy?
	I know right!!
	Just like Shubhodeep
	I mean, have you seen the guy? xP
	#GDGSwag
*/

package api

import (
	"github.com/patrickmn/go-cache"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf"
	"go-MyVIT/api/cache"
	"go-MyVIT/api/login"
	"go-MyVIT/api/scrape"
	"time"
)

var cac *cache.Cache = cache.New(2*time.Minute, 30*time.Second)

//Executable script to Login
func LogIn(regno, password, baseuri string) *login.Response {
	bow := surf.NewBrowser()
	return login.NewLogin(bow, regno, password, baseuri, cac)
}

func CourseCoursesPage(regno, password, baseuri string) *scrape.CourseStruct {
	bow := surf.NewBrowser()
	cacheSession.SetSession(bow, cac, regno)
	return scrape.Courses(bow, regno, password, baseuri)
}

func CourseSlotsPage(regno, password, baseuri, coursekey string) *scrape.SlotsStruct {
	bow := surf.NewBrowser()
	cacheSession.SetSession(bow, cac, regno)
	return scrape.Slots(bow, regno, password, baseuri, coursekey)
}

func CourseFacPage(regno, password, baseuri, coursekey, slt string) *scrape.FacStruct {
	bow := surf.NewBrowser()
	cacheSession.SetSession(bow, cac, regno)
	return scrape.Facs(bow, regno, password, baseuri, coursekey, slt)
}

func CourseDataPage(regno, password, baseuri, coursekey, slt, fac string) *scrape.CourseDataStruct {
	bow := surf.NewBrowser()
	cacheSession.SetSession(bow, cac, regno)
	return scrape.CourseData(bow, regno, password, baseuri, coursekey, slt, fac)
}

func Refresh(regno, password, baseuri string) *scrape.RefreshStruct {
	bow := surf.NewBrowser()

	return scrape.Refresh(bow, regno, password, baseuri, cacheSession.SetSession(bow, cac, regno))
}

func Spotlight(regno, password, baseuri string) *scrape.Spotlight {
	bow := surf.NewBrowser()
	return scrape.Spoli(bow, regno, password, baseuri)
}
