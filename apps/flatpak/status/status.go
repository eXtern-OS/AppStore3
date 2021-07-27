package status

import "sync"

var Mutex *sync.Mutex

func init() {
	Mutex = &sync.Mutex{}
}
