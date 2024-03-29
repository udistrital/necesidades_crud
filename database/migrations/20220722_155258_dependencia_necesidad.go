package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type DependenciaNecesidad_20220722_155258 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &DependenciaNecesidad_20220722_155258{}
	m.Created = "20220722_155258"

	migration.Register("DependenciaNecesidad_20220722_155258", m)
}

// Run the migrations
func (m *DependenciaNecesidad_20220722_155258) Up() {
	file, err := ioutil.ReadFile("../script/20220722_155258_dependencia_necesidad_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *DependenciaNecesidad_20220722_155258) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../script/20220722_155258_dependencia_necesidad_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
