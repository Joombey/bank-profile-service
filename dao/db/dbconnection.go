package db

import (
	"database/sql"
	"fmt"
	"log"

	cts "farukh.go/profile/constants"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var localDb *sql.DB

func Init() {
	db, err := sql.Open("mysql", cts.MySQLConfig.FormatDSN())
	defer func() { localDb = db }()
	if err != nil {
		log.Panicf("error opening db %s", err.Error())
	}

	stmt, err := db.Prepare(cts.DatabaseSchemaMySQL)

	if err != nil {
		log.Panicf("error creation tables %s", err.Error())
	}

	stmt.Exec()
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}

type UserTable struct {
	Id         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	CardNumber int    `db:"card_number" json:"card_number"`
	Password   string `db:"password" json:"password"`
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func (r UserRepositoryImpl) New() *UserRepositoryImpl {
	return &UserRepositoryImpl{db: localDb}
}

func (repo *UserRepositoryImpl) CreateUser(name string, cardNumber int) <-chan UserTable {
	tableChan := make(chan UserTable)
	go func() {
		repo.db.Exec("INSERT INTO users (name, card_number) values (?, ?)", name, cardNumber)
		var user = UserTable{}
		repo.db.QueryRow("SELECT * FROM users ORDER BY id DESC LIMIT 1").Scan(&user.Id, &user.Name, &user.CardNumber)
		tableChan <- user
		close(tableChan)
	}()
	return tableChan
}

func (repo *UserRepositoryImpl) GetUserById(id int) <-chan UserTable {
	tableChan := make(chan UserTable)
	go func() {
		var user = UserTable{}
		repo.db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.CardNumber)
		println(user.Id, user.Name, user.CardNumber)
		tableChan <- user
		close(tableChan)
	}()
	return tableChan
}


func (repo *UserRepositoryImpl) Delete(id int) {
	repo.db.Exec("DELETE FROM users WHERE id = ?", id)
}