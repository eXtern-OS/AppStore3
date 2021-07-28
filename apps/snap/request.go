package snap

import (
	"encoding/json"
	"github.com/eXtern-OS/common/app"
	"html"
	"io/ioutil"
	"net/http"
)

const BaseURL = "https://api.snapcraft.io/v2/snaps/find?fields=media,description,publisher,title,version&q="

func getData(param string) ([]byte, error) {
	res, err := http.Get(BaseURL + html.EscapeString(param))

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(res.Body)
}

func parseData(income []byte) ([]app.Snap, error) {
	var res []app.Snap

	return res, json.Unmarshal(income, &res)
}
