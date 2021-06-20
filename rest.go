package gomc

// https://docs.google.com/presentation/d/1enuUT-sQcIA8vFLQtc-Hthj_xsbRx88mG1BMrnRSMhk/edit
// https://journal.lampetty.net/entry/understanding-http-handler-in-go
// multiple handler
// static content
// json
// parameter
// etc?

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"

	starterListener "github.com/lestrrat-go/server-starter/listener"
)
func handleHello(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello DEMO")
}

func newHandler() http.Handler {
	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/", handleHello)
	return multiplexer
}

func RestMain(_ []string) {
	var l net.Listener
	if os.Getenv("SERVER_STARTER_PORT") != "" {
		listeners, err := starterListener.ListenAll()
		if err != nil {
			fmt.Println(err)
			return
		}
		if 0 < len(listeners) {
			l = listeners[0]
		}
	}
	if l == nil {
		var err error
		l, err = net.Listen("tcp", fmt.Sprintf(":8080"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	log.Printf("Start to serve")
	myOS, myArch := runtime.GOOS, runtime.GOARCH
	log.Printf("I'm running on %s/%s.\n", myOS, myArch)

	// よくある ListenAndServe() では、listener が出てこない
	// この例では Server Starter のために listener を切り替える
	// 可能性があるので
	fmt.Println(http.Serve(l, newHandler()))
}
