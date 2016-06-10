package app

import "log"

type Runner interface  {
	Run() *Application
}

var app *Application

func (a *Application) Run(){

}


func GetApplication() *Application {
	if app == nil {
		app = new(Application)

		// Init
		app.Config = loadConfig("config.json")
		log.Println("Application config is loaded")
		//
		//app.Doc = make(AbstractPage)
		////app.routes = make(MapRoutes)
		//
		//// Настройка значений глобальных полей документа
		//app.Doc["Config"] = app.Config
		//app.Doc["Host"] = app.Config.GetString("site.host")
		//app.Doc["MetaTitle"] = app.Config.GetString("site.title")
		//app.Doc["MetaDescription"] = app.Config.GetString("site.description")
		//app.Doc["MetaAuthor"] = app.Config.GetString("site.author")
		//app.Doc["MetaCopyright"] = app.Config.GetString("site.copyright")
		//app.Doc["MetaKeywords"] = app.Config.GetString("site.keywords")
		//app.Doc["ContactPhone"] = app.Config.GetString("site.phone")
		//
		//if !app.Config.GetBool("db.disable") {
		//	db, err := NewDatabase(app.Config.GetString("db.driver"), app.Config.GetString("db.datasource"), app.Config.GetString("db.prefix"))
		//	if err != nil {
		//		panic("Database error: " + err.Error())
		//	}
		//
		//	app.DB = db
		//	log.Println("Database is connected")
		//}
	}

	return app
}