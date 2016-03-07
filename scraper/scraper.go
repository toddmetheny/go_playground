//http://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html

package main

import (
  "net/http"

  "golang.org/x/net/html"

  "fmt"
)

var url = "https://clutchprep.com"

//~~~~~~~~~~~~~~~~~~~~~~//
// Make an HTTP request //
//~~~~~~~~~~~~~~~~~~~~~~//

func httpRequest(url string){
  resp, _ := http.Get(url)
  bytes, _ := ioutil.ReadAll(resp.Body)

  fmt.Println("HTML:\n\n", string(bytes))

  resp.Body.Close()
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~//
// Parse HTML for Anchor Tags //
//~~~~~~~~~~~~~~~~~~~~~~~~~~~~//

func parseHtml(){
  z := html.NewTokenizer(response.Body)

  for {
    tt := z.Next()

    switch {
    case tt == html.ErrorToken:
        // End of the document, we're done
        return
    case tt == html.StartTagToken:
        t := z.Token()

        isAnchor := t.Data == "a"
        if isAnchor {
            fmt.Println("We found a link!")
        }
    }
  }
}
