package query
//
//type modelQuery struct {
//	model interface{}
//}
//
//func CreateQuery() IModelQuery{
//	return &modelQuery{};
//}
//
//type IModelQuery interface {
//	For(interface{}) IFilterQuery
//}
//
//func (ctx *modelQuery) bind(f func(m interface{}) IFilterQuery) IFilterQuery{
//	return f(ctx.model);
//}
//
//func (ctx *modelQuery) For(interface{}) IFilterQuery {
//	f:= func(m interface{}) IFilterQuery {
//		return &FilterQuery{};
//	}
//	return ctx.bind(f);
//}