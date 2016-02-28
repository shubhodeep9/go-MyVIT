package main

import (
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/jar"
	"github.com/headzoo/surf/browser"
	"github.com/headzoo/surf/agent"
	"go-MyVIT/api/login"
	"fmt"
)


/*
Developer Note:
Main package for now => api package for implementation
*/

func logsIn(bow *browser.Browser,regno,password string) {
	newLogin := login.NewLogin(regno,password)
	status := make(chan int)
	go newLogin.DoLogin(status)
	success := <-status
	if success==1 {
		fmt.Println("Success")
	} else {
		fmt.Println("Try again")
		logsIn(bow,regno,password)
	}
	bow.SetSiteCookies(newLogin.GetCookies())
}

func main() {
	bow := surf.NewBrowser()
	bow.SetUserAgent(agent.Chrome())
	bow.SetAttribute(browser.FollowRedirects, true)
	bow.SetAttribute(browser.SendReferer, true)
	bow.SetCookieJar(jar.NewMemoryCookies())
	err := bow.Open("https://academics.vit.ac.in/student/stud_login.asp")
	if err != nil {
		fmt.Println(err)
	}
	/*
	@TODO retrieve details from GET URL
	*/
	regno := ""
	password := ""
	logsIn(bow,regno,password)
	bow.Open("https://academics.vit.ac.in/student/home.asp")
	fmt.Println(bow.Url())
}