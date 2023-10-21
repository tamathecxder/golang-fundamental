package database

var connection string

func init() {
	connection = "PostgreSQL"

	// fmt.Println("init is called")
}

func GetDatabase() string {
	return connection
}
