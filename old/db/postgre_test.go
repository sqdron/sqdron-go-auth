package db

import (
	//"testing"
	//. "github.com/sqdron/sqdron-go-auth/db/query"
	//. "github.com/sqdron/sqdron-go-auth/db/model"
	//. "github.com/smartystreets/goconvey/convey"
	//"github.com/jinzhu/gorm"
	//"time"
	//"fmt"
	//"os"
	//"testing"
)

//func TestQueryWorkflow2(t *testing.T) {
	//var qr = CreateQuery();
	//
	//Convey("Getting table for query", t, func() {
	//	user := &User{};
	//	user.ID = "1";
	//	user.Username = "Denis"
	//	modelQuery := qr.For(user);
	//	Convey("Make where query", func() {
	//		filterQuery := modelQuery.Where();
	//		Convey("Make Select", func() {
	//			selectQuery := filterQuery.Select();
	//			Convey("Make Order", func() {
	//				orderQuery := selectQuery.Order();
	//				Convey("Make ALL result", func() {
	//					result := orderQuery.All();
	//					So(result, ShouldNotBeNil)
	//				})
	//
	//				Convey("Make First result", func() {
	//					result := orderQuery.First();
	//					So(result, ShouldNotBeNil)
	//				})
	//
	//				Convey("Make Single result", func() {
	//					result := orderQuery.Single();
	//					So(result, ShouldNotBeNil)
	//				})
	//			})
	//		})
	//	})
	//})
//}
//
//var (
//	DB                 *gorm.DB
//	t1, t2, t3, t4, t5 time.Time
//)
//
//func init() {
//	var err error
//
//	if DB, err = OpenTestConnection("postgres"); err != nil {
//		panic(fmt.Sprintf("No error should happen when connecting to test database, but got err=%+v", err))
//	}
//
//	// DB.SetLogger(Logger{log.New(os.Stdout, "\r\n", 0)})
//	// DB.SetLogger(log.New(os.Stdout, "\r\n", 0))
//	if os.Getenv("DEBUG") == "true" {
//		DB.LogMode(true)
//	}
//	//
//	//DB.DB().SetMaxIdleConns(10)
//
//	//runMigration()
//}
////
//func OpenTestConnection(dialect string, ) (db *gorm.DB, err error) {
//	switch dialect {
//	//case "mysql":
//	//	// CREATE USER 'gorm'@'localhost' IDENTIFIED BY 'gorm';
//	//	// CREATE DATABASE gorm;
//	//	// GRANT ALL ON gorm.* TO 'gorm'@'localhost';
//	//	fmt.Println("testing mysql...")
//	//	dbhost := os.Getenv("GORM_DBADDRESS")
//	//	if dbhost != "" {
//	//		dbhost = fmt.Sprintf("tcp(%v)", dbhost)
//	//	}
//	//	db, err = gorm.Open("mysql", fmt.Sprintf("gorm:gorm@%v/gorm?charset=utf8&parseTime=True", dbhost))
//	case "postgres":
//		fmt.Println("testing postgres...")
//		db, err = gorm.Open("postgres", "user=gorm dbname=gorm");
//		//dbhost := os.Getenv("GORM_DBHOST")
//		//if dbhost != "" {
//		//	dbhost = fmt.Sprintf("host=%v ", dbhost)
//		//}
//		//db, err = gorm.Open("postgres", fmt.Sprintf("%vuser=gorm password=gorm DB.name=gorm sslmode=disable", dbhost))
//	//case "foundation":
//	//	fmt.Println("testing foundation...")
//	//	db, err = gorm.Open("foundation", "dbname=gorm port=15432 sslmode=disable")
//	//case "mssql":
//	//	fmt.Println("testing mssql...")
//	//	db, err = gorm.Open("mssql", "server=SERVER_HERE;database=rogue;user id=USER_HERE;password=PW_HERE;port=1433")
//	//default:
//	//	fmt.Println("testing sqlite3...")
//	//	db, err = gorm.Open("sqlite3", filepath.Join(os.TempDir(), "gorm.db"))
//	}
//	return
//}
////
//func TestStringPrimaryKey(t *testing.T) {
//	type UUIDStruct struct {
//		ID   string `gorm:"primary_key"`
//		Name string
//	}
//	DB.AutoMigrate(&UUIDStruct{})
//
//	data := UUIDStruct{ID: "uuid", Name: "hello"}
//	Convey("Test connection", t, func() {
//		if err := DB.Save(&data).Error; err != nil || data.ID != "uuid" {
//			t.Errorf("string primary key should not be populated")
//		}
//		So(data, ShouldNotBeNil)
//	});
//}
//
//func TestPostgreDB(t *testing.T) {
//	var db = &PostgreDB{};
//	db.For(&User{}).Where().First();
//
//	Convey("Getting table for query", t, func() {
//		user := &User{};
//		user.ID = "1";
//		user.Username = "Denis"
//		So(user.ID, ShouldNotBeEmpty)
//	})
//}
//
