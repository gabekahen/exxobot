package main

import (
  "fmt"
  "net/http"
  "html"
  "regexp"
  "ioutil"
)

func getQuote() (qTitle, qContent string) {
  type randQuote struct {
    Title string
    Content string
  }
  htmlTagRegex := regexp.MustCompile("<.+?>")
  resp, err := http.Get(quoteGenUrl)
  if err != nil {
    fmt.Println("Unable to get quote from " + quoteGenUrl)
    fmt.Println(err.Error())
    return
  }
  var myQuote randQuote
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println("Unable to read response from " + quoteGenUrl)
    fmt.Println(err.Error())
    return
  }
  err = json.Unmarshal(body, &myQuote)
  if err != nil {
    fmt.Println("Unable to parse JSON from " + quoteGenUrl)
    fmt.Println(err.Error())
    return
  }
  return myQuote.Title, htmlTagRegex.ReplaceAllString(html.UnescapeString(myQuote.Content), "")
}
