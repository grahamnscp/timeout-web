package main

import (
  "fmt"
  "net/http"
  "os"
  "time"
  "strconv"
  "strings"
)

var hostname, oserr = os.Hostname()

func main() {
  if oserr != nil {
    fmt.Println("io.hostname error: ", oserr, " hostname value:", hostname)
  }
  http.HandleFunc("/", Handler)
  http.HandleFunc("/health", healthHandler)
  http.HandleFunc("/ping", pingHandler)
  http.HandleFunc("/favicon.ico", nullHandler)
  fmt.Println("Serving HTTP on port :8080") 

  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    fmt.Println("Serve Http:", err)
  }
}

func Handler(w http.ResponseWriter, r *http.Request) {

  dt := time.Now()
  fmt.Printf("[%s timeout-web] Host: %s, Client: %s, Request: %s%s\n", dt.Format("01-02-2006 15:04:05.00"), 
              hostname, r.RemoteAddr, r.Host, r.URL.Path) 

  reqHostname := r.Host
  reqURI := strings.Trim(r.URL.Path, "/")
  reqSleep := 0
  if reqURI == "" {
    reqURI = "nil" 
  }
  reqSleep, err := strconv.Atoi(reqURI)
  if err != nil {
    fmt.Printf("[%s timeout-web] Host: %s, Client: %s requested sleep duration URL: '%s' is not a number, defaulting to: %ds\n", 
                dt.Format("01-02-2006 15:04:05.00"), hostname, r.RemoteAddr, reqURI, reqSleep)
    fmt.Fprintf(w,"[%s timeout-web] Please enter a valid sleep duration as URI, '/%s' is not valid.\n",
                   dt.Format("01-02-2006 15:04:05.00"), reqURI)
  } else {
    fmt.Printf("[%s timeout-web] Host: %s, Client: %s, sleeping for %ds..\n", dt.Format("01-02-2006 15:04:05.00"), 
                hostname, r.RemoteAddr, reqSleep)

    // sleep
    time.Sleep(time.Duration(reqSleep)*time.Second)

    fmt.Fprintf(w,"[%s timeout-web] Host: %s, Recieved Request: %s/%s\n", 
                   dt.Format("01-02-2006 15:04:05.00"), hostname, reqHostname, reqURI)
  }
  dt = time.Now()
  fmt.Printf("[%s timeout-web] Host: %s, Client: %s done. Slept for: %ds\n", dt.Format("01-02-2006 15:04:05.00"),
              hostname, r.RemoteAddr, reqSleep)
  fmt.Fprintf(w,"[%s timeout-web] Host: %s, slept for: %ds\n", dt.Format("01-02-2006 15:04:05.00"), hostname, reqSleep)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w,"ok\n")
  return
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
  dt := time.Now()
  fmt.Printf("[%s timeout-web] Host: %s, Client: %s, Request: %s%s\n", dt.Format("01-02-2006 15:04:05.00"), 
              hostname, r.RemoteAddr, r.Host, r.URL.Path) 
  fmt.Fprintf(w,"ack\n")
  return
}

func nullHandler(w http.ResponseWriter, r *http.Request) {
  return
}

