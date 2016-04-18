// Status - Dome
/*
@Author :- Ujjwal Ayyangar
Finally getting a hang of this beautiful language :D
*/

package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/login"
	"github.com/patrickmn/go-cache"
	"fmt"


	"sync"
	// "strings"
	

)

type Message struct {
	From string `json:"from"`
	Course string  `json:"course"`
	Message string `json:"message"`
	PostedOn string `json:"posted"`
}

type ViewMessages struct{
	Status string `json:"status"`
	Messages []Message `json:"messages"`
}


/*
Function ->Messg to fetch the data of messages sent by the faculty

@return ViewMessages struct
*/
func setMessageSession(bow *browser.Browser, cac *cache.Cache, regno string) {
	cacheval, _ := cac.Get(regno)
	cachevalue := cacheval.(*login.MemCache)

	bow.SetSiteCookies(cachevalue.MemCookie)
}
func Messg(bow *browser.Browser, regno, password, baseuri string, cac *cache.Cache) *ViewMessages{
	setMessageSession(bow, cac, regno)
	var msg []Message
	status := "Success"
	if 1 != 1 {
		status = "Failure"
	} else {
		    var wg sync.WaitGroup

		bow.Open(baseuri+"/student/class_message_view.asp?sem=WS")
		bow.Open(baseuri+"/student/class_message_view.asp?sem=WS")
		tables := bow.Find("table")
		messageTable:= tables.Eq(1)
		tr:= messageTable.Find("tr")
		tr_len:= tr.Length()
		tr.Each(func(i int,s *goquery.Selection) {
			if i>0 && i<tr_len-1 {
	        wg.Add(1)
	        go func(){
	        	defer wg.Done()
			fmt.Println(s.Text())
			td:=s.Find("td")
			x:=Message{
				From:td.Eq(0).Text(),
				Course:td.Eq(1).Text(),
				Message:td.Eq(2).Text(),
				PostedOn:td.Eq(3).Text(),

			}
		msg=append(msg,x)
		}()
		wg.Wait()
	}
			})


		
		}


	return &ViewMessages{
		Status:     status,
		Messages : msg,
	}
}
