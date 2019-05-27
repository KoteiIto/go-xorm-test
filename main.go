package main

import (
	"context"
	"os"

	"github.com/KoteiIto/go-xorm-test/domain/model/user"
	contextCache "github.com/KoteiIto/go-xorm-test/infra/cache/context"
	"github.com/KoteiIto/go-xorm-test/infra/db/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/golog"
)

//go:generate xorm reverse mysql "root:@tcp(0.0.0.0:4306)/xorm_master" _template/model/master domain/model/master
//go:generate xorm reverse mysql "root:@tcp(0.0.0.0:4306)/xorm_user" _template/model/user domain/model/user

func main() {
	golog.Println("hello world")

	mysql, err := mysql.NewXormMysqlDB("root:@tcp(0.0.0.0:4306)/xorm_user")
	if err != nil {
		golog.Println(err)
		os.Exit(1)
	}

	err = mysql.Truncate("group_member")
	if err != nil {
		golog.Println(err)
		os.Exit(1)
	}

	tx := mysql.NewSession()
	ctxCacheSession := contextCache.NewContextCacheAdapter(tx)

	ctx := ctxCacheSession.WithContextCache(context.Background())

	err = func() error {
		user1 := user.NewGroupMemberDto(user.GroupMember{
			UserId:  1,
			GroupId: 100,
			Role:    user.GroupMemberRoleAdmin,
		})

		user2 := user.NewGroupMemberDto(user.GroupMember{
			UserId:  2,
			GroupId: 200,
			Role:    user.GroupMemberRoleAdmin,
		})

		_, err = ctxCacheSession.Insert(ctx, user1)
		if err != nil {
			return err
		}

		_, err = ctxCacheSession.Insert(ctx, user2)
		if err != nil {
			return err
		}

		user1.SetRole(user.GroupMemberRoleGuest)
		_, err = ctxCacheSession.Update(ctx, user1)
		if err != nil {
			return err
		}

		user2.SetRole(user.GroupMemberRoleGuest)
		_, err = ctxCacheSession.Update(ctx, user1)
		if err != nil {
			return err
		}

		_, err = ctxCacheSession.Delete(ctx, user1)
		if err != nil {
			return err
		}

		user1 = user.NewGroupMemberDto(user.GroupMember{
			UserId:  1,
			GroupId: 300,
			Role:    user.GroupMemberRoleAdmin,
		})

		_, err = ctxCacheSession.Insert(ctx, user1)
		if err != nil {
			return err
		}

		user1.SetRole(user.GroupMemberRoleGuest)
		_, err = ctxCacheSession.Update(ctx, user1)
		if err != nil {
			return err
		}

		_, err = ctxCacheSession.Delete(ctx, user1)
		if err != nil {
			return err
		}

		return ctxCacheSession.Sync(ctx)
	}()

	if err != nil {
		golog.Println(err)
		os.Exit(1)
	}
}
