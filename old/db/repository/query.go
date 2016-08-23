package repository

import (
	//. "github.com/sqdron/sqdron-go-auth/types"
)
//
//type Query struct {
//	source Any
//}
//
//type IQuery interface {
//	Bind(func(Any, IQuery) IQuery) IQuery
//	ToQuery(interface{}) IQuery
//	First();
//}
//
//func (m Query) ToQuery(a Any) IQuery {
//	return Query{&a};
//}
////
//func (m Query) Bind(f func(Any, IQuery) IQuery) IQuery {
//	var i Iterator = func(m Query, it func(valye Any)) Query {
//		return m;
//	};
//
//	return i(m, func(x Any) {
//		f(x, m);
//	});
//	//return Query{&a};
//}
//
//type Iterator func(Query, func(value Any)) Query;
//
//type Predicate func(Any) bool;
//

//func (m Query) Where(p Predicate){
//
//	m.Bind(func(x Any, q Query){
//
//	})
//}

//
//func Where(model Any, predicate Predicate, m IQuery) IQuery {
//	return m.ToQuery(2 * i.(int))
//}

//
//func (m Query) Bind(f func(interface{}, IQuery) IQuery) IQuery {
//	if m == None {
//		return None
//	}
//	return f(*m.val, m)
//}
//
//
//
//type DbContext struct {
//
//}
//
//type IDbProvider interface {
//	ToQuery(model interface{}) *Query;
//}
//
//
//
//func Q(model *model.Entity) *Query {
//	return &Query{};
//}
//
//func (p *DbContext) ToQuery(model model.Entity) *Query {
//	return &Query{};
//}
//
//func (q *Query) Where(p Predicate) *Query {
//	return q;
//}
//func (q *Query) Select() *Query {
//	return q;
//}
//
//func (q *Query) Count() int {
//	return 0;
//}
//
//func (q *Query) First() Any {
//	return q;
//}
//
//func (q *Query) Last() *Query {
//	return q;
//}
//
//func (q *Query) Take(count int) *Query {
//	return q;
//}
//
//func (q *Query) Skip(count int) *Query {
//	return q;
//}
//
