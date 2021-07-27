/*
	Since I haven't found normal API we'll have to manually update it once a day

*/

package daemon

import (
	"encoding/json"
	"externos.io/AppStore3/apps/flatpak/status"
	"github.com/eXtern-OS/Beatrix"
	"github.com/eXtern-OS/common/app"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func sleep() bool {
	time.Sleep(2 * time.Hour)

	return true
}

const AppsURL = "https://flathub.org/api/v1/apps"

func getData() ([]byte, error) {
	res, err := http.Get(AppsURL)

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(res.Body)
}

func parseData(income []byte) ([]app.Flatpak, error) {
	var res []app.Flatpak

	return res, json.Unmarshal(income, &res)
}

func updateApps() {

	d, err := getData()
	if err != nil {
		go beatrix.SendError("Failed to make request: "+err.Error(), "flatpak.daemon.updateApps")
		return
	}

	apps, err := parseData(d)

	if err != nil {
		go beatrix.SendError("Failed to parse apps response: "+err.Error(), "flatpak.daemon.updateApps")
		go log.Println(string(d))
		return
	}

	var insertData []interface{}

	// Wtf

	for _, x := range apps {
		insertData = append(insertData, x)
	}

	status.Mutex.Lock()
	if err := deleteEverything(); err != nil {
		go beatrix.SendError("Failed to delete flathub apps: "+err.Error(), "flatpak.daemon.updateApps")
	} else {
		if err = dbc.InsertMany(insertData, DatabaseName, CollectionName); err != nil {
			go beatrix.SendError("Failed to insert flathub apps: "+err.Error(), "flatpak.daemon.updateApps")
		}
	}
	status.Mutex.Unlock()
	return
}

func StartDaemon() {
	updateApps()
	for sleep() {
		updateApps()
	}
}
