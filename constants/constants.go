package constants

import (
	"github.com/go-sql-driver/mysql"
)

var MySQLConfig = mysql.Config{
	User:                 "root",
	Passwd:               "root",
	DBName:               "db",
	Net:                  "tcp",
	Addr:                 "profile-db:3306",
	AllowNativePasswords: true,
	CheckConnLiveness:    true,
	MaxAllowedPacket:     64 << 20,
}

const (
	baseBankApi   string = "http://bank:8080"
	CreateCard    string = baseBankApi + "/new-card" // GET
	TransferMoney string = baseBankApi + "/transfer" // POST
	GetValue      string = baseBankApi + "/get-card" // GET with route argument :num
)

const DatabaseSchema = `
CREATE TABLE IF NOT EXISTS users(
	id int primary key auto_increment,
	name text not null,
	card_number int
);
`
