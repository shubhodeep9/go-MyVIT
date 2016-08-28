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
	"fmt"
	"github.com/alexjlockwood/gcm"
	"github.com/patrickmn/go-cache"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf"
	"go-MyVIT/api/cache"
	"go-MyVIT/api/login"
	"go-MyVIT/api/scrape"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"crypto/tls"
	"net/http"
)

var cac *cache.Cache = cache.New(2*time.Minute, 30*time.Second)

//Executable script to Login
func LogIn(regno, password, baseuri string) *login.Response {
	bow := surf.NewBrowser()
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }
	bow.SetTransport(tr)
	return login.NewLogin(bow, regno, password, baseuri, cac)
}

func CourseCoursesPage(regno, password, baseuri string) *scrape.CourseStruct {
	bow := surf.NewBrowser()
	return scrape.Courses(bow, regno, password, baseuri, cacheSession.SetSession(bow, cac, regno))
}

func CourseSlotsPage(regno, password, baseuri, coursekey string) *scrape.SlotsStruct {
	bow := surf.NewBrowser()
	return scrape.Slots(bow, regno, password, baseuri, coursekey, cacheSession.SetSession(bow, cac, regno))
}

func CourseFacPage(regno, password, baseuri, coursekey, slt string) *scrape.FacStruct {
	bow := surf.NewBrowser()
	return scrape.Facs(bow, regno, password, baseuri, coursekey, slt, cacheSession.SetSession(bow, cac, regno))
}

func CourseDataPage(regno, password, baseuri, coursekey, slt, fac string) *scrape.CourseDataStruct {
	bow := surf.NewBrowser()
	return scrape.CourseData(bow, regno, password, baseuri, coursekey, slt, fac, cacheSession.SetSession(bow, cac, regno))
}

func Refresh(regno, password, baseuri string) *scrape.RefreshStruct {
	bow := surf.NewBrowser()
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }
	bow.SetTransport(tr)
	return scrape.Refresh(bow, regno, password, baseuri, cacheSession.SetSession(bow, cac, regno))
}

func Spotlight(baseuri string) *scrape.Spotlight {
	bow := surf.NewBrowser()
	return scrape.Spoli(bow, baseuri)
}

func ProfilePic(regno, password, baseuri string) string {
	bow := surf.NewBrowser()
	cacheSession.SetSession(bow, cac, regno)
	return scrape.ProfilePhoto(bow, regno, baseuri)
}

func ShowStats() map[string]int {
	stat := make(map[string]int)
	stat["current_users"] = len(cac.Items())
	return stat
}

func FacultyInformation(regno, password, baseuri, query string) string {
	bow := surf.NewBrowser()
	return scrape.FacultySearch(bow, regno, password, baseuri, query, cacheSession.SetSession(bow, cac, regno))
}

func CookieReturn(regno string) string {
	val, found := cac.Get(regno)
	fmt.Println(found)
	if found {
		cookies := val.(*cacheSession.MemCache)
		result := ""
		for i := range cookies.MemCookie {
			result = result + cookies.MemCookie[i].Name + "=" + cookies.MemCookie[i].Value + ";"
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
	session, _ := mgo.Dial("mongodb://shubho:deep@ds047865.mlab.com:47865/analyticsweekly")
	defer session.Close()
	var registrations []*Registrations
	c := session.DB("analyticsweekly").C("gcm")
	c.Find(bson.M{}).All(&registrations)
	var regIDs []string
	for _, val := range registrations {
		regIDs = append(regIDs, val.Regid)
	}
	data := map[string]interface{}{"message": message}
	msg := gcm.NewMessage(data, regIDs...)

	// Create a Sender to send the message.
	sender := &gcm.Sender{ApiKey: "AIzaSyBMGB6Mk4SJCH3hP5f_r8OUiL0mjxEhuWk"}

	// Send the message and receive the response after at most two retries.
	response, err := sender.Send(msg, 2)
	fmt.Println(err)
	return response
}

func GcmRegister(regID string) string {
	session, _ := mgo.Dial("mongodb://shubho:deep@ds047865.mlab.com:47865/analyticsweekly")
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
