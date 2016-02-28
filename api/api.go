package api

import (
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
	"github.com/headzoo/surf/jar"
	"go-MyVIT/api/login"
)

/*
Developer Note:
Main package for now => api package for implementation
*/

func LogIn(regno, password string) *login.Response {
	bow := surf.NewBrowser()
	bow.SetAttribute(browser.FollowRedirects, true)
	bow.SetAttribute(browser.SendReferer, true)
	bow.SetCookieJar(jar.NewMemoryCookies())
	bow.Open("https://academics.vit.ac.in/student/stud_login.asp")
	/*
		@TODO retrieve details from GET URL
	*/
	return login.NewLogin(bow, regno, password)
}
