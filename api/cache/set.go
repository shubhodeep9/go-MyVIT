package cacheSession

import (
	"github.com/patrickmn/go-cache"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"net/http"
)

type MemCache struct {
	Regno        string
	MemCookieOld []*http.Cookie
	BetaClient *http.Client
}

func SetSession(bow *browser.Browser, cac *cache.Cache, regno string) bool {
	cacheval, found := cac.Get(regno)
	if found {
		cachevalue := cacheval.(*MemCache)
		bow.SetSiteCookies(cachevalue.MemCookieOld)
	}
	return found
}

func GetClient(cac *cache.Cache, regno string) (*http.Client,bool) {
    cacheval, found := cac.Get(regno)
    if found {
        cachevalue := cacheval.(*MemCache)
        return cachevalue.BetaClient,true
    }
    return &http.Client{},false
}

        
