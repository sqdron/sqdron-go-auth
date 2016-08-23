package http
//
//func Register(srv *micro.IService) {
//
//	module := httpModule.NewHttp(&httpModule.Options{Port:"8080"});
//	module.Route("/auth").
//		GET("/token", func(rw http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(rw, "Hello request ID %s\n", 211)
//	}).
//		GET("/stop", func(rw http.ResponseWriter, r *http.Request) {
//		module.Stop();
//	}).
//		POST("/login", func(rw http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(rw, "login user")
//	});
//
//	//module.Run("8000");
//	log.Println("stopped server")
//}
//
//func logger(fn http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		log.Println("Before")
//		fn(w, r)
//		log.Println("After")
//	}
//}

//
//
//
