package scrape

import (
	//"fmt"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/status"
	"os"
)

type RoomAllotStruct struct {
	AllotedRooms []det               `json:"allotedRooms"`
	Status       status.StatusStruct `json:"status"`
}
type det struct {
	Year        string `json:"year"`
	Block       string `json:"block"`
	RoomNo      string `json:"roomNo"`
	RoomType    string `json:"roomType"`
	BedType     string `json:"bedType"`
	MessName    string `json:"messName"`
	CattersName string `json:"cattersName"`
}

func RoomAllot(bow *browser.Browser, reg, baseuri string, found bool) *RoomAllotStruct {
	sem := os.Getenv("SEM")
	rooms := []det{}
	stats := status.Success()
	if !found {
		stats = status.SessionError()
	} else {
		bow.Open(baseuri + "/student/room_alloted_status_new.asp?=" + sem)
		if bow.Open(baseuri+"/student/room_alloted_status_new.asp??sem="+sem) == nil {
			table := bow.Find("table")
			table.Find("tr").Each(func(i int, s *goquery.Selection) {
				if i > 2 {
					td := s.Find("td")
					temp := det{
						Year:        td.Eq(1).Text(),
						Block:       td.Eq(2).Text(),
						RoomNo:      td.Eq(3).Text(),
						RoomType:    td.Eq(4).Text(),
						BedType:     td.Eq(5).Text(),
						MessName:    td.Eq(6).Text(),
						CattersName: td.Eq(7).Text(),
					}
					rooms = append(rooms, temp)

				}
			})

		}
	}
	return &RoomAllotStruct{
		AllotedRooms: rooms[:len(rooms)-1],
		Status:       stats,
	}
}
