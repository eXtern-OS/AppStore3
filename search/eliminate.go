package search

import (
	"github.com/eXtern-OS/common/app"
	"sync"
)


/* Eliminator provides smart queue which removes (eliminates) duplicates. The priority is following:
	* eXtern OS apps
	* Flatpak apps
	* Snap apps
*/
type eliminator struct {
	waiting int

	queueFlatLocked bool
	queueSnapLocked bool
	queueSnap       []*app.ExportedApp
	queueFlat       []*app.ExportedApp

	mutexMap *sync.RWMutex
	apps     map[string]*app.ExportedApp
}

// check provides method to check whether appName has already been seen
func (e *eliminator) check(appName string) bool {
	e.mutexMap.RLock()
	_, ok := e.apps[appName]
	e.mutexMap.RUnlock()
	return ok
}

// add adds app into map
func (e *eliminator) add(app *app.ExportedApp) {
	e.mutexMap.Lock()
	e.apps[app.Name] = app
	e.mutexMap.Unlock()
}

/* start is a core function. You pass a chan that consumes apps and a waitgroup for snap queries
 * Once we are done with all apps incoming, we can begin sorting them
 * Therefore, we launch flatpak queue (by releasing locker)
*/
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


// checkAndAdd adds apps if name wasn't met before. This was done to avoid situations when snap or flatpak apps erase eXtern OS
func (e *eliminator) checkAndAdd(income []*app.ExportedApp) {
	for _, x := range income {
		if !e.check(x.Name) {
			e.add(x)
		}
	}
}


// queueSnapDaemon checks and adds apps from the queue to the result
func (e *eliminator) queueSnapDaemon(wg *sync.WaitGroup) {
	for e.queueSnapLocked {

	}

	e.checkAndAdd(e.queueSnap)

	wg.Done()
}


// queueFlatDaemon checks and adds apps from the flatpak queue. After that, launches corresponding function for snap by releasing snap locker
func (e *eliminator) queueFlatDaemon() {
	for e.queueFlatLocked {

	}

	e.checkAndAdd(e.queueFlat)

	e.queueSnapLocked = false
}


// get provides method to extract all sorted apps
func (e *eliminator) get() []app.ExportedApp {
	var res []app.ExportedApp

	e.mutexMap.Lock()
	for _, x := range e.apps {
		res = append(res, *x)
	}
	e.mutexMap.Unlock()
	return res
}


// newEliminator creates new eliminator. Default params are understandable
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
