package extern

import "github.com/eXtern-OS/common/db"

var dbc *db.DBClient

// -- Default DB implementation, nothing to see here --
const (
	DatabaseName   = "Apps"
	CollectionName = "Extern"
)

func Init() {
	dbc = db.NewClient()
}
