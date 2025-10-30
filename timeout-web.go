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
  http.HandleFunc("/favicon.ico", nullHandler)
  fmt.Println("Serving HTTP on port :8080") 

  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    fmt.Println("Serve Http:", err)
  }
}

func Handler(w http.ResponseWriter, r *http.Request) {

  dt := time.Now()
  fmt.Printf("[%s timeout-web] Host: %s, Request: %s%s\n", dt.Format("01-02-2006 15:04:05.00"), 
              hostname, r.Host, r.URL.Path) 

  reqHostname := r.Host
  reqURI := strings.Trim(r.URL.Path, "/")
  reqSleep := 0
  if reqURI == "" {
    reqURI = "nil" 
  }
  reqSleep, err := strconv.Atoi(reqURI)
  if err != nil {
    fmt.Printf("Requested sleep duration URL: '%s' is not a number, defaulting to: %ds\n", reqURI, reqSleep)
    fmt.Fprintf(w,"[%s timeout-web] Please enter a valid sleep duration as URI, '/%s' is not valid.\n",
                   dt.Format("01-02-2006 15:04:05.00"), reqURI)
  } else {
    fmt.Printf("[%s timeout-web] Sleeping for %ds..\n", dt.Format("01-02-2006 15:04:05.00"), reqSleep)

    // sleep
    time.Sleep(time.Duration(reqSleep)*time.Second)

    fmt.Fprintf(w,"[%s timeout-web] Host: %s, Recieved Request: %s/%s\n", 
                   dt.Format("01-02-2006 15:04:05.00"), hostname, reqHostname, reqURI)
  }
  dt = time.Now()
  fmt.Printf("[%s timeout-web] done. Slept for: %ds\n", dt.Format("01-02-2006 15:04:05.00"), reqSleep)
  fmt.Fprintf(w,"[%s timeout-web] Slept for: %ds\n", dt.Format("01-02-2006 15:04:05.00"), reqSleep)
}

func nullHandler(w http.ResponseWriter, r *http.Request) {
  return
}
