package extern

import (
	"externos.io/AppStore3/query"
	"github.com/eXtern-OS/common/app"
	"sync"
)

func Search(q query.Query, res chan []app.App, wg *sync.WaitGroup, limit int) {
	wg.Done()
}
