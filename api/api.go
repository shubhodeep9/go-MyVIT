package main

import (
	"fmt"
	"go-MyVIT/Godeps/_workspace/src/github.com/headzoo/surf"
	"go-MyVIT/Godeps/_workspace/src/github.com/headzoo/surf/agent"
	"go-MyVIT/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/Godeps/_workspace/src/github.com/headzoo/surf/jar"
	"go-MyVIT/api/login"
)

/*
Developer Note:
Main package for now => api package for implementation
*/

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
	var (
		regno    string
		password string
	)
	fmt.Println("Enter")
	fmt.Scanf("%s %s", &regno, &password)
	login.NewLogin(bow, regno, password)
	bow.Open("https://academics.vit.ac.in/student/home.asp")
}
