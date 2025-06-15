package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/Aritiaya50217/GolangBook/tree/main/BuildingMicroservices/Chapter01IntroductionToMicroservices/rpc_http_json/contract"
)

const port = 1234

type HttpConn struct {
	in  io.Reader
	out io.Writer
}

func (c *HttpConn) Read(p []byte) (n int, err error) {
	return c.in.Read(p)
}

func (c *HttpConn) Writer(d []byte) (n int, err error) {
	return c.out.Write(d)
}

func (c *HttpConn) Close() error {
	return nil
}

type HelloworldHandler struct{}

func (h *HelloworldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}

func StartServer() {
	helloWorld := new(HelloworldHandler)
	rpc.Register(helloWorld)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}

	http.Serve(l, http.HandlerFunc(httpHandler))
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	serverCodec := jsonrpc.NewServerCodec(&HttpConn{in: r.Body, out: w})
	err := rpc.ServeRequest(serverCodec)
	if err != nil {
		log.Printf("Error while serving JSON request: %v", err)
		http.Error(w, "Error while serving JSON request, details have been logged.", 500)
		return
	}
}
