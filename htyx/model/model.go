package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

//将第二个数据库注释掉

type Database struct {
	Self *sqlx.DB //第一个数据库
	//Docker *sqlx.DB
}

var DB *Database

/*
func setupDB(db *sqlx.DB) {
	db.DB.SetMaxOpenConns() // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB.SetMaxIdleConns() // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}
*/
func openDB(username, password, addr, name string) *sqlx.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, addr, name)
	db, err := sqlx.Open("mysql", config)
	if err != nil {

		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}
	//setupDB(db)

	return db
}
func InitSelfDB() *sqlx.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

/*
func InitDockerDB() *sqlx.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}
*/
func Init() {
	DB = &Database{
		Self: InitSelfDB(),
		//Docker: InitDockerDB(),
	}
}
func Close() {
	DB.Self.Close()
	//DB.Docker.Close()
}
