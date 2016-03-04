package main

import(
        "flag"
        "fmt"
        "net/http"
        "strings"
        "io/ioutil"
        "encoding/json"
        "gopkg.in/mgo.v2/bson"
        "text/tabwriter"
        "os"
      )
// The main structure for database
type Panda struct {
    Name string        `json:"name"`                    // Name of Panda
}

func main(){
   url := flag.String("url", "", "API endpoint URL")
   method := flag.String("method", "", "API Method ( get, post, put, delete)")
   data := flag.String("data", "", "JSON Body")
    flag.Parse()
   req := new(http.Request)
    client := &http.Client{}
    
switch *method{
    case "get":
        req, _ = http.NewRequest("GET", *url, nil)
    case "post":
        req, _ = http.NewRequest("POST", *url, strings.NewReader(*data))
    case "delete":
        req, _ = http.NewRequest("DELETE", *url, nil)
    case "put":
        req, _ = http.NewRequest("PUT", *url,  strings.NewReader(*data))
    default:
    }
        fmt.Println("\n\n")
        resp, _ := client.Do(req)
        fmt.Println("\n\n---------------------------------Response Headers------------------------------")
        fmt.Println(resp.Header)
        body, _ := ioutil.ReadAll(resp.Body)
    m := *method
    if m == "get" {
        var pandas []Panda
        err := json.Unmarshal(body, &pandas)
        if err != nil {
            panic(err)
        }
        w := new(tabwriter.Writer)
        w.Init(os.Stdout, 0, 8, 0, '\t', 0)
        fmt.Println("\n\n---------------------------------Response Body---------------------------------")
        fmt.Fprintln(w, "Id\t Name\t")
        for _, panda := range pandas{
            fmt.Fprintf(w,"%s  \t  %s  \t\n",panda.Id.Hex(), panda.Name)
        }
        w.Flush()
    }
        fmt.Println("\n\n")
}
