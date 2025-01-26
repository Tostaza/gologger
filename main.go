package main

import (
    "net/http"
    "time"
    "os"
    "log"
)

var (
    ADDR_VARNAME = "GOLOGGER_LISTEN_ADDR"
    DEFAULT_ADDR = "0.0.0.0"
    PORT_VARNAME = "GOLOGGER_LISTEN_PORT"
    DEFAULT_PORT = "8080"
)


func logRequest(w http.ResponseWriter, r *http.Request) {
    // Loop to get all headers
    header := ""
    for key, values := range r.Header {
        header = header + key + ": "
        for i := 0; i < len(values); i++ {
            header = header + values[i] + " "
        }
    }

    log.Printf("[INFO]: HTTP request received at [%s]: %s %s %s %v %v Host:%s Header:{ %s}",
        time.Now().UTC(), 
        r.Method, 
        r.URL.Scheme+r.URL.Host+r.URL.Path, 
        r.Proto,
        r.ProtoMajor,
        r.ProtoMinor,
        r.Host,
        header,
    )
}

// ping returns a "pong" message
func ping(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong"))
}


func main() {
    addr, ok := os.LookupEnv(ADDR_VARNAME)
    if !ok {
        addr = DEFAULT_ADDR
    }

    port, ok := os.LookupEnv(PORT_VARNAME);
    if !ok {
        port = DEFAULT_PORT
    }

    http.HandleFunc("/ping", ping)

    http.HandleFunc("/", logRequest)

    log.Printf("[INFO]: Starting server at %s:%s", addr, port)
    // Running proxy server
    if err := http.ListenAndServe(addr+":"+port, nil); err != nil {
        log.Printf("[ERROR]: Could not start the server: %v", err)
    }
}
