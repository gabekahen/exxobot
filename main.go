package main

import (
	"flag"
	"fmt"
	"time"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "html"
  "regexp"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
  quoteGenUrl string = "http://quotesondesign.com/wp-json/posts?filter[orderby]=rand&filter[posts_per_page]=1"
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
	return
}

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
    fmt.Println(m.Author.Username, m.ID, c.ID, m.Content)
    qTitle, qContent := getQuote()
    myMessage := "Now Exxo, you know what " + qTitle + " always says...\n" + "_" + qContent + "_"
    s.ChannelMessageSend(c.ID, myMessage)
  }
	// Print message to stdout.
	fmt.Printf("%20s %20s %20s > %s\n", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)

}

func getQuote() (qTitle, qContent string) {
  type randQuote struct {
    Title string
    Content string
  }
  htmlTagRegex := regexp.MustCompile("<.+?>")
  resp, err := http.Get(quoteGenUrl)
  if err != nil {
    return
  }
  var myQuote randQuote
  body, err := ioutil.ReadAll(resp.Body)
  body = body[1:len(body)-1]
  json.Unmarshal(body, &myQuote)
  return myQuote.Title, htmlTagRegex.ReplaceAllString(html.UnescapeString(myQuote.Content), "")
}
