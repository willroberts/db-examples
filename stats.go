package databases

import (
	"database/sql"
	"log"
)

// PrintStats pretty-prints a sql.DBStats object.
func PrintStats(dbname string, stats sql.DBStats) {
	log.Println("database:", dbname)
	log.Println("- max open connections:", stats.MaxOpenConnections)
	log.Println("- open connections:", stats.OpenConnections)
	log.Println("- connections in use:", stats.InUse)
	log.Println("- idle connections", stats.Idle)
	log.Println("- waiting connections:", stats.WaitCount)
	log.Println("- wait time:", stats.WaitDuration)
	log.Println("- closed because idle:", stats.MaxIdleClosed)
	log.Println("- closed because idle too long:", stats.MaxIdleTimeClosed)
	log.Println("- closed because exceeded lifetime:", stats.MaxLifetimeClosed)
}
