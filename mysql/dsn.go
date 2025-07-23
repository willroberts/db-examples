package mysqltest

import "fmt"

// DataSourceName.
type DSN struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

// Format: "username:password@protocol(address)/dbname?param=value".
func (d DSN) ToString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.Database,
	)
}
