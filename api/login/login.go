/*
@Author Shubhodeep Mukherjee
@Organization Google Developers Group VIT Vellore
	Isn't Go sexy?
	I know right!!
	Just like Shubhodeep
	I mean, have you seen the guy? xP
	#GDGSwag
*/

package login

import (
	"github.com/patrickmn/go-cache"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/cache"
	"go-MyVIT/api/login/captcha"
	"go-MyVIT/api/status"
	"net/url"
	"os"
)

type Response struct {
	Regno  string              `json:"regno"`
	Campus string              `json:"campus"`
	Status status.StatusStruct `json:"status"`
}

var Sessionname string

/*
Creates a new StudLogin object and Starts logging in
@return Response struct
@param Registration_Number Password
*/
func NewLogin(bow *browser.Browser, reg, pass, baseuri string, cac *cache.Cache) *Response {
	stats := make(chan int)
	go DoLogin(bow, reg, pass, stats, baseuri, cac)
	success := <-stats
	var stt status.StatusStruct
	if success == 1 {
		stt = status.Success()
	} else {
		stt = status.CredentialError()
	}
	return &Response{
		Regno:  reg,
		Campus: "vellore",
		Status: stt,
	}
}

/*
Parses the captcha using parse.py and creates a session,
Using that session user is logged in.
@param bow(surf Browser) registration_no password status(channel for goroutine)
@return void
*/
func DoLogin(bow *browser.Browser, reg, pass string, stats chan int, baseuri string, cac *cache.Cache) {

	bow.Open("https://vtop.vit.ac.in/student/captcha.asp")
	out, err := os.Create("api/login/" + reg + ".bmp")
	bow.Download(out)
	out1 := captcha.Parse(reg)
	go os.Remove("api/login/" + reg + ".bmp")
	if err != nil {
		stats <- 0
	} else {
		v := url.Values{}
		v.Set("regno", reg)
		v.Add("passwd", pass)
		v.Add("vrfcd", out1)
		v.Add("message", "")
		bow.PostForm(baseuri+"/student/stud_login_submit.asp", v)
		stud_home := baseuri + "/student/stud_home.asp"
		home := baseuri + "/student/home.asp"
		u := bow.Url().String()
		if u == stud_home || u == home {
			cac.Set(reg, &cacheSession.MemCache{Regno: reg, MemCookie: bow.SiteCookies()}, cache.DefaultExpiration)
			stats <- 1
		} else {
			stats <- 0
		}
	}
}
