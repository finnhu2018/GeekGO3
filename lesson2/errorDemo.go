package lesson2

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

/*
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
*/

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

func SelectPolicy() {
	var person []Person
	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 1)
	switch err := errors.Cause(err).(type) {
	case sql.ErrNoRows:
		fmt.Println("exec failed, ", err)
		return
	default:
		panic("panic error!")
	}
	fmt.Println("select succ:", person)
}

var Db *sqlx.DB

func init() {

	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = database
	defer db.Close() // 注意这行代码要写在上面err判断的下面
}
