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
		Campus string `json:"campus"`
		Status string `json:"status"`
		Spot Spotlight1 `json:"spotlight"`

}
type Spotlight1 struct {
	Academics  []Base `json:"academics"`
	Coe []Base `json:"coe"`
	Research []Base  `json:"research"`
	}

type Base struct{
	Text string `json:"text"`
	Url string `json:"url"`
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
		
		if len(res) ==0{
			res=make([]Base,0)
		}
		if len(acad) ==0{
			acad=make([]Base,0)
		}
		if len(coe) == 0{
			coe=make([]Base,0)
		}
				
}
x:= Spotlight1{
			Academics:acad,
			Coe:coe,
			Research:res,
		}


	return &Spotlight{
		Campus : "Vellore",
		Status:     status,
		Spot: x,

	}
}
