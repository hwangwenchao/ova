package msssql

import (
	"database/sql"
	"strings"

	_ "github.com/mattn/go-adodb"
)

type SA struct {
	User     string
	Password string
}

type Msssql struct {
	*sql.DB
	DataSource string
	Database   string
	IsWindows  bool
	Sa         SA
}

func (m *Msssql) Open() (err error) {
	var conf []string
	conf = append(conf, "Provider=SQLOLEDB")
	conf = append(conf, "Data Source="+m.DataSource)
	if m.IsWindows {
		conf = append(conf, "integrated security=SSPI")
	}
	conf = append(conf, "Initial Catalog="+m.Database)
	conf = append(conf, "user id="+m.Sa.User)
	conf = append(conf, "password="+m.Sa.Password)

	m.DB, err = sql.Open("adodb", strings.Join(conf, ";"))
	if err != nil {
		return err
	}
	return nil
}
