package db

import (
	"github.com/jinzhu/gorm"
	//. "github.com/sqdron/sqdron-go-auth/db/query"
)

//var (
//	DB *gorm.DB
//)

type PostgreDB struct {
	Database
	DB *gorm.DB
}

type PostgreQuery struct {
	DB *gorm.DB
}

func (db *PostgreDB) For(model interface{}) *PostgreQuery {
	res := &PostgreQuery{};
	return res;
}

func (ctx *PostgreQuery) Bind(f func(db *gorm.DB) *PostgreQuery) *PostgreQuery {
	return f(ctx.DB);
}

//func (ctx *PostgreQuery) BindR(f func(db *gorm.DB) *IResultQuery) *IResultQuery {
//	result := &PostgreQuery{};
//	f(ctx.DB);
//	return result;
//}
//
//func (ctx *PostgreQuery) Where() *PostgreQuery {
//	return ctx.Bind(func(filter interface{}) *PostgreQuery {
//
//		return &PostgreQuery{DB: ctx.DB.Where("")};
//	});
//}
//
//func (ctx *PostgreQuery) Select() *PostgreQuery {
//	return ctx;
//}
//
//func (ctx *PostgreQuery) Order() *PostgreQuery {
//	return ctx;
//}
//
//func (ctx *PostgreQuery) First() *PostgreQuery {
//	return ctx.Bind(func(filter interface{}) *PostgreQuery {
//		return &PostgreQuery{DB: ctx.DB.First("")};
//	});
//}


//func (ctx *PostgreQuery) bind(func(filter interface{}) *PostgreQuery) *PostgreQuery{
//	return
//}