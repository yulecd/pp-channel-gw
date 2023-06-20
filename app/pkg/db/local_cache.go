package db

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	CountriesLCache    *cache.Cache
	StatesLCache       *cache.Cache
	CityLCache         *cache.Cache
	ProjectsLCache     *cache.Cache
	CustomerGroupCache *cache.Cache
)

func InitLocalCache() {
	CountriesLCache = cache.New(30*time.Minute, 60*time.Minute)
	StatesLCache = cache.New(30*time.Minute, 60*time.Minute)
	CityLCache = cache.New(30*time.Minute, 60*time.Minute)
	ProjectsLCache = cache.New(30*time.Minute, 60*time.Minute)
	CustomerGroupCache = cache.New(30*time.Minute, 60*time.Minute)
}
