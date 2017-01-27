package main

import (
	"fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "html"
  "regexp"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  c, err := s.State.Channel(m.ChannelID)
  if err != nil {
    return
  }
  // func (s *Session) ChannelMessageEdit(channelID, messageID, content string) (st *Message, err error) {
  //g, err := s.State.Guild(c.GuildID)
  if m.Author.ID == "90670438945951744" {
    qTitle, qContent := getQuote()
    myMessage := "Now Exxo, you know what " + qTitle + " always says...\n" + "_" + qContent + "_"
    _, err := s.ChannelMessageSend(c.ID, myMessage)
    if err != nil {
      fmt.Println("Unable to send message to channel ", c.ID, err.Error())
      return
    }
  }
}

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
