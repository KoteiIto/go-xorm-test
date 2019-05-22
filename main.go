package main

import "github.com/kataras/golog"

//go:generate xorm reverse mysql "root:@tcp(0.0.0.0:4306)/xorm_master" _template/model/master domain/model/master
//go:generate xorm reverse mysql "root:@tcp(0.0.0.0:4306)/xorm_user" _template/model/user domain/model/user

func main() {
	golog.Println("hello world")
}
