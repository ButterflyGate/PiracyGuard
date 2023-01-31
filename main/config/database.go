package config

import "fmt"

type Database struct {
	user string
	pass string
	host string
	port int
	name string
}

func NewDatabase() *Database {
	d := &Database{
		user: GetEnvStr("DB_USER", d.db.dbUser),
		pass: GetEnvStr("DB_PASS", d.db.dbPass),
		host: GetEnvStr("DB_HOST", d.db.dbHost),
		port: GetEnvInt("DB_PORT", d.db.dbPort),
		name: GetEnvStr("DB_NAME", d.db.dbName),
	}
	return d
}

func (d *Database) DSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", d.user, d.pass, d.host, d.port, d.name)
}
