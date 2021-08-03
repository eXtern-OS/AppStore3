package search

import (
	"github.com/eXtern-OS/common/app"
	"sync"
)

type eliminator struct {
	waiting int

	queueFlatLocked bool
	queueSnapLocked bool
	queueSnap       []*app.ExportedApp
	queueFlat       []*app.ExportedApp

	mutexMap *sync.RWMutex
	apps     map[string]*app.ExportedApp
}

func (e *eliminator) check(appName string) bool {
	e.mutexMap.RLock()
	_, ok := e.apps[appName]
	e.mutexMap.RUnlock()
	return ok
}

func (e *eliminator) add(app *app.ExportedApp) {
	e.mutexMap.Lock()
	e.apps[app.Name] = app
	e.mutexMap.Unlock()
}

func (e *eliminator) start(income chan *app.ExportedApp, wg *sync.WaitGroup) {
	go e.queueFlatDaemon()
	go e.queueSnapDaemon(wg)
	for a := range income {
		switch a.AppType {
		case app.ExternApp:
			go e.add(a)
			break
		case app.FlatpakApp:
			if !e.check(a.Name) {
				e.queueFlat = append(e.queueFlat, a)
			}
			break
		case app.SnapApp:
			if !e.check(a.Name) {
				e.queueSnap = append(e.queueSnap, a)
			}
			break
		}
	}
	e.queueFlatLocked = false
}

func (e *eliminator) checkAndAdd(income []*app.ExportedApp) {
	for _, x := range income {
		if !e.check(x.Name) {
			e.add(x)
		}
	}
}

func (e *eliminator) queueSnapDaemon(wg *sync.WaitGroup) {
	for e.queueSnapLocked {

	}

	e.checkAndAdd(e.queueSnap)

	wg.Done()
}

func (e *eliminator) queueFlatDaemon() {
	for e.queueFlatLocked {

	}

	e.checkAndAdd(e.queueFlat)

	e.queueSnapLocked = false
}

func (e *eliminator) get() []app.ExportedApp {
	var res []app.ExportedApp

	e.mutexMap.Lock()
	for _, x := range e.apps {
		res = append(res, *x)
	}
	e.mutexMap.Unlock()
	return res
}

func newEliminator() eliminator {
	var e = eliminator{
		queueFlatLocked: true,
		queueSnapLocked: true,
		waiting:         0,
		queueFlat:       []*app.ExportedApp{},
		queueSnap:       []*app.ExportedApp{},
		mutexMap:        &sync.RWMutex{},
	}

	e.mutexMap.Lock()
	e.apps = make(map[string]*app.ExportedApp)
	e.mutexMap.Unlock()
	return e
}
