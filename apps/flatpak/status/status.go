package status

import "sync"

var ReasonableLimit int

var Mutex *sync.Mutex

func init() {
	Mutex = &sync.Mutex{}
}
