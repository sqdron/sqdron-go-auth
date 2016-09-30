package main

import (
	//"github.com/sqdron/go-squad/endpoint"
	//"github.com/sqdron/go-squad/optionizer"
	//"github.com/sqdron/squad/configurator"

	"github.com/sqdron/squad/configurator"
	"os"
	"os/signal"
	//"fmt"
	//"time"
	"fmt"
	//"time"
	"sync"
	"github.com/sqdron/squad/endpoint/nats"
)

type Options struct {
	ApplicationID string `json:"application-id"`

	EndpointUrl   string `json:"endpoint-url"`
	EndpointToken string `json:"endpoint-token"`

	//TLSTimeout         float64       `json:"tls_timeout"`
	//TLS                bool          `json:"-"`
	//TLSVerify          bool          `json:"-"`
	//TLSCert            string        `json:"-"`
	//TLSKey             string        `json:"-"`
	//TLSCaCert          string        `json:"-"`
	//TLSConfig          *tls.Config   `json:"-"`
}

func main() {

	opt := &Options{}

	cfg := configurator.New()
	cfg.ReadFlags(opt)

	ep := nats.NatsEndpoint(opt.EndpointUrl)


	//c := make(chan os.Signal, 2)
	//signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	wg := sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(1)
	ch := ep.Listen("hello")
	go func() {
		for {
			select {
			case <-c:
				fmt.Println("Interrupting...")
				wg.Done()
				os.Exit(1)
			case res := <-ch:
				fmt.Println(res)
			}
		}
	}()

	wg.Wait()

	//log.Println(opt.ApplicationID)
	//log.Println(opt.EndpointUrl)

	//httpTrans := endpoint.Http()
	//cfg := optionizer.New().Range(httpTrans.Configuratoin())
	//cfg.ParseFlags()

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