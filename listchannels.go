package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func main() {

	//Configure the Slack Api interface for the nlopes/slack
	api := slack.New("SLACK API KEY")
	//Uncomment the logging and debuging if you are having issues
	//logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	//slack.SetLogger(logger)
	//api.SetDebug(true)

	// Use The Slack api to get group information Groups are the chat rooms.
	channels, err := api.GetChannels(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, channel := range channels {
		fmt.Println(channel.Name, ":", channel.ID)
	}

}
