package main

import (
	"io"
	"net/http"
	"strconv"
	"github.com/Maoyis/GO-Study"
	//logging "github.com/op/go-logging"
)

func heartbreakerHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, GO_Study.Study())
}

var LISTENING_PORT = 1024
//var logger = logging.MustGetLogger("hugmachine.log")

func main() {
	//logging.NewLogBackend(os.Stderr, "", 0)
	http.HandleFunc("/study", heartbreakerHandler)
	//logger.Infof("Listening on port %d", LISTENING_PORT)
	err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(LISTENING_PORT), nil)
	if err != nil {
		//logger.Fatal("ListenAndServe: " + err.Error())
	}
}
