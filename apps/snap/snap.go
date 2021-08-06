package snap

import (
	beatrix "github.com/eXtern-OS/Beatrix"
	"github.com/eXtern-OS/common/app"
	"log"
	"sync"
)

type CacheMap struct {
	Mutex sync.Mutex
	Map   map[string]bool
}

func (c *CacheMap) get(key string) bool {
	c.Mutex.Lock()
	v := c.Map[key]
	c.Mutex.Unlock()
	return v
}

func (c *CacheMap) set(key string) {
	c.Mutex.Lock()
	c.Map[key] = true
	c.Mutex.Unlock()
}

var cmap CacheMap

func init() {
	cmap = CacheMap{
		Mutex: sync.Mutex{},
	}
	cmap.Mutex.Lock()
	cmap.Map = make(map[string]bool)
	cmap.Mutex.Unlock()
}

func Search(q string, res chan *app.ExportedApp, limit int, wg *sync.WaitGroup) {
	if cmap.get(q) {
		if apps, err := LoadFromCache(q); err != nil {
			go beatrix.SendError("Failed to load from cache: "+err.Error(), "snap.Search")
			log.Println("Failed to load from cache: " + err.Error())
		} else {
			if len(apps) > limit {
				apps = apps[:limit]
			}
			for _, x := range apps {
				m := x.Export()
				res <- &m
			}
			wg.Done()
			return
		}
	}

	d, err := getData(q)

	if err != nil {
		go beatrix.SendError("Failed to get data: "+err.Error(), "apps.snap.Search")
		go log.Println("Failed to get data: " + err.Error())
	} else {
		snapApps, err := parseData(d)

		if err != nil {
			go beatrix.SendError("Failed to parse data: "+err.Error(), "apps.snap.Search")
			go log.Println(err)
		} else {
			if len(snapApps) > limit {
				snapApps = snapApps[:(limit - 1)]
			}

			for _, x := range snapApps {
				m := x.Export()
				res <- &m
			}
		}

		go AddToCache(q, snapApps)
	}

	//log.Println("Snap finished")
	wg.Done()
}
