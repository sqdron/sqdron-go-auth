package query
//
//import (
//	"testing"
//	. "github.com/smartystreets/goconvey/convey"
//	. "github.com/sqdron/sqdron-go-auth/db/model"
//)
//
//func TestQueryWorkflow(t *testing.T) {
//	var qr = CreateQuery();
//	Convey("Getting table for query", t, func() {
//		user := &User{};
//		user.ID = "1";
//		user.Username = "Denis"
//		modelQuery := qr.For(user);
//		Convey("Make where query", func() {
//			filterQuery := modelQuery.Where();
//			Convey("Make Select", func() {
//				selectQuery := filterQuery.Select();
//				Convey("Make Order", func() {
//					orderQuery := selectQuery.Order();
//					Convey("Make ALL result", func() {
//						result := orderQuery.All();
//						So(result, ShouldNotBeNil)
//					})
//
//					Convey("Make First result", func() {
//						result := orderQuery.First();
//						So(result, ShouldNotBeNil)
//					})
//
//					Convey("Make Single result", func() {
//						result := orderQuery.Single();
//						So(result, ShouldNotBeNil)
//					})
//				})
//			})
//		})
//	})
//}