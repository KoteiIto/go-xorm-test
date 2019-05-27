package infra

import (
	"context"
	"testing"

	"github.com/KoteiIto/go-xorm-test/domain/model/condition"
	"github.com/KoteiIto/go-xorm-test/domain/model/user"
	contextCache "github.com/KoteiIto/go-xorm-test/infra/cache/context"
	"github.com/KoteiIto/go-xorm-test/infra/db/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func Benchmark_XormDB_SELECT(b *testing.B) {
	engine, err := xorm.NewEngine("mysql", "root:@tcp(0.0.0.0:4306)/xorm_user")
	if err != nil {
		b.Error(err)
	}

	tx := engine.NewSession()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u := &user.GroupMember{}
		tx.Where("user_id = ?", 1)
		has, err := tx.Get(u)
		if err != nil || !has {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func Benchmark_NoCacheDB_SELECT(b *testing.B) {
	mysql, err := mysql.NewXormMysqlDB("root:@tcp(0.0.0.0:4306)/xorm_user")
	if err != nil {
		b.Error(err)
	}

	tx := mysql.NewSession()
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u := user.NewGroupMemberDtoEmpty()
		has, err := tx.Get(ctx, u, user.GenGroupMemberUserIdCondition(condition.OperatorEQ, 1))
		if err != nil || !has {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func Benchmark_ContextCacheDB_SELECT(b *testing.B) {
	mysql, err := mysql.NewXormMysqlDB("root:@tcp(0.0.0.0:4306)/xorm_user")
	if err != nil {
		b.Error(err)
	}

	tx := mysql.NewSession()
	ctxCacheSession := contextCache.NewContextCacheAdapter(tx)
	ctx := ctxCacheSession.WithContextCache(context.Background())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u := user.NewGroupMemberDtoEmpty()
		has, err := ctxCacheSession.Get(ctx, u, user.GenGroupMemberUserIdCondition(condition.OperatorEQ, 1))
		if err != nil || !has {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func Benchmark_XormDB_UPDATE(b *testing.B) {
	engine, err := xorm.NewEngine("mysql", "root:@tcp(0.0.0.0:4306)/xorm_user")
	if err != nil {
		b.Error(err)
	}

	tx := engine.NewSession()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tx.Begin()
		u := &user.GroupMember{}
		tx.Where("user_id = ?", 1)
		has, err := tx.Get(u)
		if err != nil || !has {
			b.Error(err)
		}

		u.Version = 1
		tx.Where("user_id = ?", 1)
		affected, err := tx.Update(u)
		if err != nil || affected == 0 {
			b.Error(err)
		}

		u.Role = user.GroupMemberRoleAdmin
		tx.Where("user_id = ?", 1)
		affected, err = tx.Update(u)
		if err != nil || affected == 0 {
			b.Error(err)
		}

		u.Role = user.GroupMemberRoleGuest
		u.Version = 10
		tx.Where("user_id = ?", 1)
		affected, err = tx.Update(u)
		if err != nil || affected == 0 {
			b.Error(err)
		}

		err = tx.Commit()
		if err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func Benchmark_NoCacheDB_UPDATE(b *testing.B) {
	mysql, err := mysql.NewXormMysqlDB("root:@tcp(0.0.0.0:4306)/xorm_user")
	if err != nil {
		b.Error(err)
	}

	tx := mysql.NewSession()
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tx.Begin()
		u := user.NewGroupMemberDtoEmpty()
		has, err := tx.Get(ctx, u, user.GenGroupMemberUserIdCondition(condition.OperatorEQ, 1))
		if err != nil || !has {
			b.Error(err)
		}

		u.SetVersion(1)
		affected, err := tx.Update(ctx, u)
		if err != nil || affected == 0 {
			b.Error(err)
		}

		u.SetRole(user.GroupMemberRoleAdmin)
		affected, err = tx.Update(ctx, u)
		if err != nil || affected == 0 {
			b.Error(err)
		}

		u.SetVersion(10)
		u.SetRole(user.GroupMemberRoleGuest)
		affected, err = tx.Update(ctx, u)
		if err != nil || affected == 0 {
			b.Error(err)
		}

		err = tx.Commit()
		if err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func Benchmark_ContextCacheDB_UPDATE(b *testing.B) {
	mysql, err := mysql.NewXormMysqlDB("root:@tcp(0.0.0.0:4306)/xorm_user")
	if err != nil {
		b.Error(err)
	}

	tx := mysql.NewSession()
	ctxCacheSession := contextCache.NewContextCacheAdapter(tx)
	ctx := ctxCacheSession.WithContextCache(context.Background())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u := user.NewGroupMemberDtoEmpty()
		has, err := ctxCacheSession.Get(ctx, u, user.GenGroupMemberUserIdCondition(condition.OperatorEQ, 1))
		if err != nil || !has {
			b.Error(err)
		}

		u.SetVersion(1)
		affected, err := ctxCacheSession.Update(ctx, u)
		if err != nil || affected == 0 {
			b.Error(err)
		}

		u.SetRole(user.GroupMemberRoleAdmin)
		affected, err = ctxCacheSession.Update(ctx, u)
		if err != nil || affected == 0 {
			b.Error(err)
		}

		u.SetVersion(10)
		u.SetRole(user.GroupMemberRoleGuest)
		affected, err = ctxCacheSession.Update(ctx, u)
		if err != nil || affected == 0 {
			b.Error(err)
		}

		_, err = ctxCacheSession.Sync(ctx)
		if err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}
