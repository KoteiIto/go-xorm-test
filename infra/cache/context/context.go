package context

import (
	"context"
	"fmt"
	"sort"

	projectContext "github.com/KoteiIto/go-xorm-test/context"
	"github.com/KoteiIto/go-xorm-test/domain/repository/db"
)

type ContextCacheAdapter struct {
	session db.Session
}

type contextCache struct {
	tableCacheMap map[string]tableCache
	CurrentOrder  int
}

type tableCache struct {
	changedDtoMap  map[string]db.CrudDto
	selectedDtoMap map[string]db.CrudDto
}

func (c *contextCache) nextOrder() int {
	c.CurrentOrder++
	return c.CurrentOrder
}

func NewContextCacheAdapter(sess db.Session) *ContextCacheAdapter {
	return &ContextCacheAdapter{session: sess}
}

func (sess *ContextCacheAdapter) WithContextCache(ctx context.Context) context.Context {
	return context.WithValue(ctx, projectContext.ContextCacheKey, &contextCache{
		tableCacheMap: make(map[string]tableCache),
	})
}

func (sess *ContextCacheAdapter) Reset(ctx context.Context) error {
	c, ok := ctx.Value(projectContext.ContextCacheKey).(*contextCache)
	if !ok || c == nil {
		return fmt.Errorf("context cache has not been initialized")
	}

	c.tableCacheMap = make(map[string]tableCache)
	c.CurrentOrder = 0
	return nil
}

func (sess *ContextCacheAdapter) Sync(ctx context.Context) error {
	c, ok := ctx.Value(projectContext.ContextCacheKey).(*contextCache)
	if !ok || c == nil {
		return fmt.Errorf("context cache has not been initialized")
	}

	defer sess.Reset(ctx)

	l := 0
	for _, tableCache := range c.tableCacheMap {
		l += len(tableCache.changedDtoMap)
	}

	dtos := make([]db.CrudDto, l)
	i := 0
	for _, tableCache := range c.tableCacheMap {
		for _, dto := range tableCache.changedDtoMap {
			dtos[i] = dto
			i++
		}
	}

	sort.Slice(dtos, func(e1, e2 int) bool {
		return dtos[e1].Order() < dtos[e2].Order()
	})

	err := sess.session.Begin()
	if err != nil {
		return err
	}

	for _, dto := range dtos {
		var (
			affected int64
			err      error
		)
		switch {
		case dto.IsCreated():
			affected, err = sess.session.Insert(ctx, dto)
		case dto.IsUpdated():
			affected, err = sess.session.Update(ctx, dto)
			if affected == 0 {
				err = fmt.Errorf("Concurrent update. table=[%s], cacheKey=[%s]", dto.Table(), dto.CacheKey())
			}
		case dto.IsDeleted():
			affected, err = sess.session.Delete(ctx, dto)
		}
		if err != nil {
			sess.session.Rollback()
			return err
		}
	}

	return sess.session.Commit()
}

func (sess *ContextCacheAdapter) Get(ctx context.Context, dto db.CrudDto, conditions ...db.Condition) (bool, error) {
	c, ok := ctx.Value(projectContext.ContextCacheKey).(*contextCache)
	if !ok || c == nil {
		return false, fmt.Errorf("context cache has not been initialized")
	}

	table := dto.Table()
	_tableCache, ok := c.tableCacheMap[table]
	if ok {

	}
}

func find(_tableCache tableCache, condtions []db.Condition) *db.CrudDto {
	return nil
}

func (sess *ContextCacheAdapter) Insert(ctx context.Context, dto db.CrudDto) (int64, error) {
	c, ok := ctx.Value(projectContext.ContextCacheKey).(*contextCache)
	if !ok || c == nil {
		return 0, fmt.Errorf("context cache has not been initialized")
	}

	table := dto.Table()
	cacheKey := dto.CacheKey()

	var (
		_tableCache tableCache
		hasTable    bool
		hasDto      bool
	)
	defer func() {
		if _tableCache.changedDtoMap != nil {
			c.tableCacheMap[table] = _tableCache
		}
	}()
	_tableCache, hasTable = c.tableCacheMap[table]
	if hasTable {
		if _, hasDto = _tableCache.selectedDtoMap[cacheKey]; hasDto {
			return 0, fmt.Errorf("duplicate entry. table=[%s] cacheKey=[%s]", table, cacheKey)
		} else if _, hasDto = _tableCache.changedDtoMap[cacheKey]; hasDto {
			switch {
			case dto.IsCreated():
				return 0, fmt.Errorf("duplicate entry. table=[%s] cacheKey=[%s]", table, cacheKey)
			case dto.IsUpdated():
				return 0, fmt.Errorf("duplicate entry. table=[%s] cacheKey=[%s]", table, cacheKey)
			case dto.IsDeleted():
				delete(_tableCache.changedDtoMap, cacheKey)
				return 1, nil
			default:
				return 0, fmt.Errorf("unexpected entry. table=[%s] cacheKey=[%s]", table, cacheKey)
			}
		}
	} else {
		_tableCache = tableCache{
			selectedDtoMap: make(map[string]db.CrudDto),
			changedDtoMap:  make(map[string]db.CrudDto),
		}
		c.tableCacheMap[table] = _tableCache
	}

	order := c.nextOrder()
	dto.SetOrder(order)
	dto.AsCreated()
	_tableCache.changedDtoMap[cacheKey] = dto

	return 1, nil
}

func (sess *ContextCacheAdapter) Update(ctx context.Context, dto db.CrudDto) (int64, error) {
	c, ok := ctx.Value(projectContext.ContextCacheKey).(*contextCache)
	if !ok || c == nil {
		return 0, fmt.Errorf("context cache has not been initialized")
	}

	table := dto.Table()
	cacheKey := dto.CacheKey()

	var (
		_tableCache tableCache
		existDto    db.CrudDto
		hasTable    bool
		hasDto      bool
	)
	defer func() {
		if _tableCache.changedDtoMap != nil {
			c.tableCacheMap[table] = _tableCache
		}
	}()
	_tableCache, hasTable = c.tableCacheMap[table]
	if hasTable {
		if existDto, hasDto = _tableCache.selectedDtoMap[cacheKey]; hasDto {
			delete(_tableCache.selectedDtoMap, cacheKey)
		} else if existDto, hasDto = _tableCache.changedDtoMap[cacheKey]; hasDto {
			switch {
			case dto.IsCreated():
				existDto.SetEntity(dto.Entity())
				return 1, nil
			case dto.IsUpdated():
				existDto.SetEntity(dto.Entity())
				return 1, nil
			case dto.IsDeleted():
				return 0, nil
			default:
				return 0, fmt.Errorf("unexpected entry. table=[%s] cacheKey=[%s]", table, cacheKey)
			}
		}
	} else {
		_tableCache = tableCache{
			selectedDtoMap: make(map[string]db.CrudDto),
			changedDtoMap:  make(map[string]db.CrudDto),
		}
		c.tableCacheMap[table] = _tableCache
	}

	order := c.nextOrder()
	dto.SetOrder(order)
	dto.AsUpdated()
	_tableCache.changedDtoMap[cacheKey] = dto

	return 1, nil
}

func (sess *ContextCacheAdapter) Delete(ctx context.Context, dto db.CrudDto) (int64, error) {
	c, ok := ctx.Value(projectContext.ContextCacheKey).(*contextCache)
	if !ok || c == nil {
		return 0, fmt.Errorf("context cache has not been initialized")
	}

	table := dto.Table()
	cacheKey := dto.CacheKey()

	var (
		_tableCache tableCache
		hasTable    bool
		hasDto      bool
	)
	defer func() {
		if _tableCache.changedDtoMap != nil {
			c.tableCacheMap[table] = _tableCache
		}
	}()
	_tableCache, hasTable = c.tableCacheMap[table]
	if hasTable {
		if _, hasDto = _tableCache.selectedDtoMap[cacheKey]; hasDto {
			delete(_tableCache.selectedDtoMap, cacheKey)
		} else if _, hasDto = _tableCache.changedDtoMap[cacheKey]; hasDto {
			switch {
			case dto.IsCreated():
				delete(_tableCache.changedDtoMap, cacheKey)
				return 1, nil
			case dto.IsUpdated():
				delete(_tableCache.changedDtoMap, cacheKey)
			case dto.IsDeleted():
				return 0, nil
			default:
				return 0, fmt.Errorf("unexpected entry. table=[%s] cacheKey=[%s]", table, cacheKey)
			}
		}
	} else {
		_tableCache = tableCache{
			selectedDtoMap: make(map[string]db.CrudDto),
			changedDtoMap:  make(map[string]db.CrudDto),
		}
		c.tableCacheMap[table] = _tableCache
	}

	order := c.nextOrder()
	dto.SetOrder(order)
	dto.AsDeleted()
	_tableCache.changedDtoMap[cacheKey] = dto
	return 1, nil
}
