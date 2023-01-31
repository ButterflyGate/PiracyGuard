package config

type defaultValue struct {
	db   defaultValueDB
	http defaultValueHTTP
	log  defaultValueLog
}

type defaultValueDB struct {
	dbUser string
	dbPass string
	dbHost string
	dbPort int
	dbName string
}
type defaultValueHTTP struct {
	httpListenPort int
}
type defaultValueLog struct {
	logLevel int
}

var d defaultValue = defaultValue{
	db: defaultValueDB{
		dbUser: "postgres",
		dbPass: "postgres",
		dbHost: "butterfalygate.com",
		dbPort: 5432,
		dbName: "piracy",
	},
	http: defaultValueHTTP{
		httpListenPort: 30000,
	},
	log: defaultValueLog{
		logLevel: 7,
	},
}
