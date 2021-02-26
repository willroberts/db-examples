package databases

var (
	// CreateDatabase creates a test database in MariaDB.
	CreateDatabase = "CREATE DATABASE IF NOT EXISTS test_database"

	// CreateTable creates a table with two columns: id (int) and value (varchar(32)).
	CreateTable = "CREATE TABLE IF NOT EXISTS test_table (id int, value varchar(32))"

	// Insert writes a record into the test table.
	Insert = "INSERT INTO test_table VALUES (?, ?)"

	// Select reads all records from the test table.
	Select = "SELECT id, value FROM test_table"

	// SelectOne reads one record from the test table.
	SelectOne = "SELECT id, value FROM test_table LIMIT 1"
)
