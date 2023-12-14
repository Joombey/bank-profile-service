package constants

import (
	"github.com/go-sql-driver/mysql"
)

var MySQLConfig = mysql.Config{
	User:                 "root",
	Passwd:               "root",
	DBName:               "maria_db",
	Net:                  "tcp",
	Addr:                 "localhost:3306",
	AllowNativePasswords: true,
	CheckConnLiveness:    true,
	MaxAllowedPacket:     64 << 20,
}

const (
	baseBankApi     string = "http://bank:8080"
	CreateCard      string = baseBankApi + "/new-card" // GET
	TransferMoney   string = baseBankApi + "/transfer" // POST
	GetValue        string = baseBankApi + "/get-card" // GET with route argument :num
	DeleteCard      string = baseBankApi + "/delete"   // GET with route argument :num
	LocalConfigPath string = "I:/dev/go-projects/bank-profile-service/configs/local.yaml"
)

const DatabaseSchemaMySQL = `
CREATE TABLE IF NOT EXISTS users(
	id int primary key auto_increment,
	name text not null,
	card_number int
);
`
const DatabaseSchemaPostgres = `
CREATE TABLE IF NOT EXISTS users(
	id SERIAL primary key,
	name text not null,
	card_number int
);
`
