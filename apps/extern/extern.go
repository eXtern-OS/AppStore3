package extern

import (
	"extenos.io/AppStore3/app"
	"extenos.io/AppStore3/query"
	"sync"
)

func Search(q query.Query, res chan []app.App, wg *sync.WaitGroup) {
	wg.Done()
}
