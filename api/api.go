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
	"crypto/tls"
	"github.com/alexjlockwood/gcm"
	"github.com/patrickmn/go-cache"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/cache"
	"go-MyVIT/api/login"
	"go-MyVIT/api/scrape"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"os"
	"time"
)

var cac *cache.Cache = cache.New(2*time.Minute, 30*time.Second)

//Executable script to Login
func LogIn(regno, password, baseuri string) *login.Response {
	var bow *browser.Browser = surf.NewBrowser()
	var client http.Client
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	client = login.LoginVtopBeta(client, regno, password)

	return login.NewLogin(bow, regno, password, baseuri, cac, &client)
}

/*

func LoginInVtopBeta(regno, password string) *http.Client {
	return login.LoginVtopBeta(regno, password)
}*/

func CourseCoursesPage(regno, password, baseuri string) *scrape.CourseStruct {
	var bow *browser.Browser = surf.NewBrowser()
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	return scrape.Courses(bow, regno, password, baseuri, cacheSession.SetSession(bow, cac, regno))
}

func CourseSlotsPage(regno, password, baseuri, coursekey string) *scrape.SlotsStruct {
	var bow *browser.Browser = surf.NewBrowser()
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	return scrape.Slots(bow, regno, password, baseuri, coursekey, cacheSession.SetSession(bow, cac, regno))
}
func FacPhotos(regno, password, query, baseuri string) *scrape.FacPhoto {
	var bow *browser.Browser = surf.NewBrowser()
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	return scrape.FacultyPhoto(bow, regno, password, query, baseuri, cacheSession.SetSession(bow, cac, regno))
}

func CalCourses(regno, password, baseuri string) *scrape.CalCourses {
	var bow *browser.Browser = surf.NewBrowser()
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	return scrape.CalCourseFunc(bow, regno, baseuri, cacheSession.SetSession(bow, cac, regno))
}

func ShowMenu(regno, password, baseuri string) *scrape.MenuStruct {
	var bow *browser.Browser = surf.NewBrowser()
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	return scrape.ShowMenu(bow, regno, baseuri, cacheSession.SetSession(bow, cac, regno))
}
func CourseFacPage(regno, password, baseuri, coursekey, slt string) *scrape.FacStruct {
	var bow *browser.Browser = surf.NewBrowser()
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	return scrape.Facs(bow, regno, password, baseuri, coursekey, slt, cacheSession.SetSession(bow, cac, regno))
}

func CourseDataPage(regno, password, baseuri, coursekey, slt, fac string) *scrape.CourseDataStruct {
	var bow *browser.Browser = surf.NewBrowser()
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	return scrape.CourseData(bow, regno, password, baseuri, coursekey, slt, fac, cacheSession.SetSession(bow, cac, regno))
}

func Refresh(regno, password, baseuri string) *scrape.RefreshStruct {
	var bow *browser.Browser = surf.NewBrowser()
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	return scrape.Refresh(bow, regno, password, baseuri, cacheSession.SetSession(bow, cac, regno))
}

func Spotlight(baseuri string) *scrape.Spotlight {
	var bow *browser.Browser = surf.NewBrowser()
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	return scrape.Spoli(bow, baseuri)
}

func ProfilePic(regno, password, baseuri string) string {
	var bow *browser.Browser = surf.NewBrowser()
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	cacheSession.SetSession(bow, cac, regno)
	return scrape.ProfilePhoto(bow, regno, baseuri)
}
func ShowMessages(regno, password, baseuri string) *scrape.MessagesStruct {
	var bow *browser.Browser = surf.NewBrowser()
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	return scrape.FacMessage(bow, regno, baseuri, cacheSession.SetSession(bow, cac, regno))
}
func LeaveRequest(regno, password, baseuri string) *scrape.LeaveRequestStruct {
	var bow *browser.Browser = surf.NewBrowser()
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	return scrape.LeaveRequest(bow, regno, baseuri, cacheSession.SetSession(bow, cac, regno))
}
func ShowExamSchedule(regno, password, baseuri string) *scrape.MainExamSchedule {
	var bow *browser.Browser = surf.NewBrowser()
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	return scrape.ExmSchedule(bow, regno, baseuri, cacheSession.SetSession(bow, cac, regno))
}

/* These are for vtopbeta2 */
func ShowTimetable(regno, password, baseuri string) *scrape.Timetable3 {
	client, _ := cacheSession.GetClient(cac, regno)
	return scrape.ShowTimetable(*client, regno, password, baseuri)
}

func ShowTimetable2(regno, password, baseuri string) *scrape.Timetable2 {
	client, _ := cacheSession.GetClient(cac, regno)
	return scrape.ShowTimetable2(*client, regno, password, baseuri)
}

func ShowAttendance(regno, password, baseuri string) *scrape.Attendance2 {
	client, _ := cacheSession.GetClient(cac, regno)
	return scrape.ScrapeAttendance(*client, regno, password, baseuri)
}

func ShowExamScheduleVtopBeta(regno, password, baseuri string) *scrape.MainExamSchedule2 {
	client, _ := cacheSession.GetClient(cac, regno)
	return scrape.ShowExamScheduleVtopBeta(*client, regno, password, baseuri)
}

func ShowPersonalDetails(regno, password, baseuri string) *scrape.PersonalDetailsStruct {
	client, _ := cacheSession.GetClient(cac, regno)
	return scrape.ShowPersonalDetails(*client, regno, password, baseuri)
}

/* Till here */

func ShowEducationalDetails(regno, password, baseuri string) *scrape.EducationalDetailsStruct {
	var bow *browser.Browser = surf.NewBrowser()
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	return scrape.ShowEducationalDetails(bow, regno, baseuri, cacheSession.SetSession(bow, cac, regno))
}
func ShowFamilyDetails(regno, password, baseuri string) *scrape.FamilyDetailsStruct {
	var bow *browser.Browser = surf.NewBrowser()
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	return scrape.ShowFamilyDetails(bow, regno, baseuri, cacheSession.SetSession(bow, cac, regno))
}
func RoomAllot(regno, password, baseuri string) *scrape.RoomAllotStruct {
	var bow *browser.Browser = surf.NewBrowser()
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	return scrape.RoomAllot(bow, regno, baseuri, cacheSession.SetSession(bow, cac, regno))
}
func ShowHostelDetails(regno, password, baseuri string) *scrape.HostelDetailsStruct {
	var bow *browser.Browser = surf.NewBrowser()
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	return scrape.ShowHostelDetails(bow, regno, baseuri, cacheSession.SetSession(bow, cac, regno))
}

func ShowStats() map[string]int {
	stat := make(map[string]int)
	stat["current_users"] = len(cac.Items())
	return stat
}

func FacultyInformation(regno, password, query, baseuri string) *scrape.AllFacs {
	var bow *browser.Browser = surf.NewBrowser()
	var tr *http.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	bow.SetTransport(tr)
	return scrape.FacultySearch(bow, regno, password, query, baseuri, cacheSession.SetSession(bow, cac, regno))
}

func CookieReturn(regno string) string {
	val, found := cac.Get(regno)
	if found {
		cookies := val.(*cacheSession.MemCache)
		result := ""
		for i := range cookies.MemCookieOld {
			result = result + cookies.MemCookieOld[i].Name + "=" + cookies.MemCookieOld[i].Value + ";"
		}
		return result[:len(result)-1]
	} else {
		return ""
	}
}

type Registrations struct {
	Regid string
}

func GcmSender(message string) *gcm.Response {
	session, _ := mgo.Dial(os.Getenv("VITMONGOURL"))
	defer session.Close()
	var registrations []*Registrations
	c := session.DB("analyticsweekly").C("gcm")
	c.Find(bson.M{}).All(&registrations)
	var regIDs []string
	for _, val := range registrations {
		regIDs = append(regIDs, val.Regid)
	}
	data := map[string]interface{}{"message": message}

	//query for more than 1000 regs
	for i := 0; i < len(regIDs)/1000; i++ {
		msg := gcm.NewMessage(data, regIDs[1000*i:(1000*i)+1000]...)
		sender := &gcm.Sender{ApiKey: os.Getenv("VITKEY")}
		// Send the message and receive the response after at most two retries.
		sender.Send(msg, 2)
	}
	sender := &gcm.Sender{ApiKey: os.Getenv("VITKEY")}
	//for remainder
	div := len(regIDs) / 1000
	msg := gcm.NewMessage(data, regIDs[div*1000:]...)

	// Send the message and receive the response after at most two retries.
	response, _ := sender.Send(msg, 2)
	return response
}

func GcmRegister(regID string) string {
	session, _ := mgo.Dial(os.Getenv("VITMONGOURL"))
	defer session.Close()
	c := session.DB("analyticsweekly").C("gcm")
	n, _ := c.Find(bson.M{"regid": regID}).Count()
	if n == 0 {
		c.Insert(&Registrations{Regid: regID})
		return "Success"
	} else {
		return "Failure"
	}
}
