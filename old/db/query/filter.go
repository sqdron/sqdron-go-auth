package query

//import (
//	. "github.com/sqdron/sqdron-go-auth/common"
//)
//
//type Predicate func(Any) bool;
//
//
//func CreateFilter() IFilterQuery{
//	return &FilterQuery{};
//}
//
//type FilterQuery struct{
//	Filter interface{}
//}
//
//type IFilterQuery interface {
//	Where() *FilterQuery
//	Select() *FilterQuery
//	Order() *FilterQuery
//}
//
//func (ctx *FilterQuery) Bind(f func(filter interface{}) *FilterQuery) *FilterQuery{
//	return f(ctx.Filter);
//}
//
//func (ctx *FilterQuery) Where() *FilterQuery {
//	return ctx;
//}
//
//func (ctx *FilterQuery) Select() *FilterQuery {
//	return ctx;
//}
//
//func (ctx *FilterQuery) Order() *FilterQuery {
//	return ctx;
//}