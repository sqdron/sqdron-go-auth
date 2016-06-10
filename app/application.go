package app

import (
	//"log"
	//"net/http"
	//"fmt"
	//"golang.org/x/tools/go/gcimporter15/testdata"
)

type Application struct {
	Config Config
}

//func (app *Application) Run() {
	//r := mux.NewRouter()
	//r.StrictSlash(true)
	//
	//for _, val := range app.routes {
	//	for url, ctrl := range val {
	//		r.HandleFunc(url, obs(ctrl))
	//	}
	//}
	//
	//http.Handle("/", r)
	//listen := fmt.Sprintf("%s:%d", app.Config.GetString("net.listen.host"), app.Config.GetInt("net.listen.port"))
	//
	//log.Println("Server is started on", listen)
	//if err := http.ListenAndServe(listen, nil); err != nil {
	//	log.Println(err)
	//}

//}
func (a *Application) LoadConfig(){
	a.Config = nil;
	//a.Run();
}
