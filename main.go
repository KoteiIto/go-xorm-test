package main

import (
	"fmt"
	"os"

	"github.com/KoteiIto/go-xorm-test/domain/model/user"
	"github.com/KoteiIto/go-xorm-test/domain/repository/db"
	"github.com/KoteiIto/go-xorm-test/infra/db/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/golog"
)

//go:generate xorm reverse mysql "root:@tcp(0.0.0.0:4306)/xorm_master" _template/model/master domain/model/master
//go:generate xorm reverse mysql "root:@tcp(0.0.0.0:4306)/xorm_user" _template/model/user domain/model/user

func main() {
	golog.Println("hello world")

	mysql, err := mysql.NewDB("root:@tcp(0.0.0.0:4306)/xorm_user")
	if err != nil {
		golog.Println(err)
		os.Exit(1)
	}

	result, err := mysql.Transaction(func(tx db.Session) (interface{}, error) {
		user1 := user.NewGroupMemberDto(user.GroupMember{
			UserId:  1,
			GroupId: 100,
			Role:    user.GroupMemberRoleAdmin,
		})

		_, err := tx.Insert(user1)
		if err != nil {
			return nil, err
		}

		user1.SetRole(user.GroupMemberRoleGuest)

		_, err = tx.Update(user1)
		if err != nil {
			return nil, err
		}

		_, err = tx.Delete(user1)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("reset db")
	})

	if err != nil {
		golog.Println(err)
		os.Exit(1)
	}
	golog.Println(result)
}
