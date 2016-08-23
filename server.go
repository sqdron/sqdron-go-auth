package main

import (
	"github.com/sqdron/go-squad/endpoint"
	"github.com/sqdron/go-squad/optionizer"
)

type Options struct {
}

//
//
//// Options block for gnatsd server.
//func main() {
//
//	server:= micro.NewServer()
//	server.Module(nil);
//
//	server.Start();
//}

func main() {

	httpTrans := endpoint.Http()
	cfg := optionizer.New().Range(httpTrans.Options())
	cfg.ParseFlags()

	//
	//cfg.Describe("EndpointSecurity", "SecutrityKey", "password")
	//
	//ep:= endpoint.Http(nil).Use();
	////t:= transport.Http(context.Background());
	//
	//squad.New().Transport(t).Activate();
	//// Parse flags
	////sOpts, nOpts := parseFlags()
	//// override the NoSigs for NATS since we have our own signal handler below
	////stand.ConfigureLogger(sOpts, nOpts)
	//s := micro.Server()
	//
	//s.UseHttp()
	////s.Module()
	////http.Register(s);
	////s.Start();
	//
	//
	//
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt)
	//go func() {
	//	s.Run(<-c)
	//	os.Exit(0)
	//}()
	//runtime.Goexit()

}
//
//func parseFlags() (*interface{}, interface{}) {
//
//}