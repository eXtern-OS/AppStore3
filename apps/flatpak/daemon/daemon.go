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

// We update 12 times / day
func sleep() bool {
	time.Sleep(2 * time.Hour)

	return true
}

const AppsURL = "https://flathub.org/api/v1/apps"

// obviously, it gets data from the URL
func getData() ([]byte, error) {
	res, err := http.Get(AppsURL)

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(res.Body)
}


// The same thing here
func parseData(income []byte) ([]app.Flatpak, error) {
	var res []app.Flatpak

	return res, json.Unmarshal(income, &res)
}

func updateApps() {

	go log.Println("Started updateApps")

	d, err := getData()
	if err != nil {
		go beatrix.SendError("Failed to make request: "+err.Error(), "flatpak.daemon.updateApps")
		return
	}

	go log.Println("Got data")

	apps, err := parseData(d)

	go log.Println("Parsed data")

	if err != nil {
		go beatrix.SendError("Failed to parse apps response: "+err.Error(), "flatpak.daemon.updateApps")
		go log.Println(string(d))
		return
	}

	status.ReasonableLimit = len(apps)

	var insertData []interface{}

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

	go log.Println("Finished update apps")

	return
}

func StartDaemon() {
	log.SetPrefix("[FLATPAK/DAEMON] ")
	updateApps()
	for sleep() {
		updateApps()
	}
}
