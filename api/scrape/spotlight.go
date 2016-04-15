// Status - Dome
/*
@Author :- Ujjwal Ayyangar
Finally getting a hang of this beautiful language :D
*/

package scrape

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	//"go-MyVIT/api/login"
	//"github.com/patrickmn/go-cache"

	//"sync"
	// "strings"
	//"fmt"

)

type Spotlight struct {
	Status string `json:"status"`
	Academics  []Base `json:"acad"`
	Coe []Base `json:"coe"`
	Research []Base  `json:"research"`
	}

type Base struct{
	Text string `json:"text"`
	Url string `json:"link"`
}


/*
Function ->Spoli to fetch the data of spotlight

@return Spoli struct
*/

func Spoli(bow *browser.Browser,regno, password, baseuri string) *Spotlight{

	status := "Success"
	var acad []Base
	var coe []Base
	var res []Base
	if 1 != 1 {
		status = "Failure"
	} else {
		bow.Open(baseuri+"/include_spotlight_part01.asp")
		bow.Open(baseuri+"/include_spotlight_part01.asp")
		tables := bow.Find("table")

		
		tables.Find("a").Each(func(_ int, s *goquery.Selection) {
			
			url,_:= s.Attr("href")
			temp:= Base {
				Text :s.Text(),
				Url:url,
			}
			acad = append(acad,temp)
    
})
		bow.Open(baseuri+"/include_spotlight_part02.asp")
		bow.Open(baseuri+"/include_spotlight_part02.asp")
		tables2 := bow.Find("table")

		
		tables2.Find("a").Each(func(_ int, s *goquery.Selection) {
			
			url,_:= s.Attr("href")
			temp:= Base {
				Text :s.Text(),
				Url:url,
			}
			coe = append(coe,temp)
    
})
		bow.Open(baseuri+"/include_spotlight_part03.asp")
		bow.Open(baseuri+"/include_spotlight_part03.asp")
		tables3 := bow.Find("table")

		
		tables3.Find("a").Each(func(_ int, s *goquery.Selection) {
			
			url,_:= s.Attr("href")
			temp:= Base {
				Text :s.Text(),
				Url:url,
			}
			res = append(res,temp)
    
})
		
}

	return &Spotlight{
		Status:     status,
		Academics: acad,
		Coe: coe,
		Research: res,
	}
}
