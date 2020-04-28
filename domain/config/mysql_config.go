/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 18/11/2019, 20:50
 * Copyright (c) 2019
 */

package config

import "github.com/jinzhu/gorm"

func NewMysqlConfig(cfg Config) (err error, db *gorm.DB) {
	dbUser := cfg.GetString(`mysql.user`)
	dbPass := cfg.GetString(`mysql.pass`)
	dbName := cfg.GetString(`mysql.database`)
	dbHost := cfg.GetString(`mysql.address`)
	dbPort := cfg.GetString(`mysql.port`)

	db, err = gorm.Open("mysql", ""+dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed connect to database")
	}

	db.SingularTable(true)

	return err, db
}
