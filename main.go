package main

import (
	"context"
	"os"

	"github.com/KoteiIto/go-xorm-test/domain/model/condition"
	"github.com/KoteiIto/go-xorm-test/domain/model/user"
	contextCache "github.com/KoteiIto/go-xorm-test/infra/cache/context"
	"github.com/KoteiIto/go-xorm-test/infra/db/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/golog"
)

//go:generate xorm reverse mysql "root:@tcp(0.0.0.0:4306)/xorm_master" _template/model/master domain/model/master
//go:generate xorm reverse mysql "root:@tcp(0.0.0.0:4306)/xorm_user" _template/model/user domain/model/user

func main() {
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

		_, err = ctxCacheSession.Get(
			ctx,
			user.NewGroupMemberDto(user.GroupMember{}),
			user.GenGroupMemberUserIdCondition(condition.OperatorEQ, 10),
		)
		if err != nil {
			return err
		}

		_, err = ctxCacheSession.Get(
			ctx,
			user.NewGroupMemberDto(user.GroupMember{}),
			user.GenGroupMemberUserIdCondition(condition.OperatorGTE, 20),
			user.GenGroupMemberRoleCondition(condition.OperatorEQ, user.GroupMemberRoleGuest),
		)
		if err != nil {
			return err
		}

		user1.SetRole(user.GroupMemberRoleAdmin)
		_, err = ctxCacheSession.Insert(ctx, user1)
		if err != nil {
			return err
		}

		_, err := ctxCacheSession.Sync(ctx)
		return err
	}()

	if err != nil {
		golog.Println(err)
		os.Exit(1)
	}

	err = func() error {
		user2 := user.NewGroupMemberDtoEmpty()
		_, err = ctxCacheSession.Get(
			ctx,
			user2,
			user.GenGroupMemberUserIdCondition(condition.OperatorEQ, 2),
		)
		if err != nil {
			return err
		}

		user2.SetRole(user.GroupMemberRoleAdmin)
		ctxCacheSession.Update(ctx, user2)

		user2 = user.NewGroupMemberDtoEmpty()
		_, err = ctxCacheSession.Get(
			ctx,
			user2,
			user.GenGroupMemberUserIdCondition(condition.OperatorEQ, 2),
		)
		if err != nil {
			return err
		}

		_, err := ctxCacheSession.Delete(ctx, user2)
		if err != nil {
			return err
		}

		user2 = user.NewGroupMemberDtoEmpty()
		has, err := ctxCacheSession.Get(
			ctx,
			user2,
			user.GenGroupMemberUserIdCondition(condition.OperatorEQ, 2),
		)
		if err != nil {
			return err
		}
		golog.Println(has)

		user1 := user.NewGroupMemberDtoEmpty()
		has, err = ctxCacheSession.Get(
			ctx,
			user1,
			user.GenGroupMemberUserIdCondition(condition.OperatorEQ, 1),
		)
		if err != nil {
			return err
		}
		golog.Println(has)

		user1.SetRole(user.GroupMemberRoleGuest)
		_, err = ctxCacheSession.Update(ctx, user1)
		if err != nil {
			return err
		}

		_, err = ctxCacheSession.Sync(ctx)
		return err
	}()

	if err != nil {
		golog.Println(err)
		os.Exit(1)
	}
}
