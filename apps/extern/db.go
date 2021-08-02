package extern

import "github.com/eXtern-OS/common/db"

var dbc *db.DBClient

const (
	DatabaseName   = "Apps"
	CollectionName = "Extern"
)

func Init() {
	dbc = db.NewClient()
}
