package cacheSession

import (
	"github.com/patrickmn/go-cache"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"net/http"
)

type MemCache struct {
	Regno     string
	MemCookie []*http.Cookie
}

func SetSession(bow *browser.Browser, cac *cache.Cache, regno string) {
	cacheval, _ := cac.Get(regno)
	cachevalue := cacheval.(*MemCache)

	bow.SetSiteCookies(cachevalue.MemCookie)
}
