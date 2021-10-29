package status

import "sync"

// How many apps should user expect rEaSoNaBly ((Clarkson ((awesome))))  
var ReasonableLimit int

// Mutex to avoid collisions when searching
var Mutex *sync.Mutex

func init() {
	Mutex = &sync.Mutex{}
}
