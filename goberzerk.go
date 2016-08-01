package main

import (
	"log"
	"os"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/leap"
	"github.com/nlopes/slack"
)

func main() {

	//Configure the Slack Api interface for the nlopes/slack
	api := slack.New("SLACK API KEY")
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(true)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	// Use The Slack api to get group information Groups are the chat rooms.
	//groups, err := api.GetGroups(false)
	//if err != nil {
	//	fmt.Printf("%s\n", err)
	//	return
	//}

	// Print out the group information use the group number for your chat room in the api call that is sending data to the slack channel
	//for _, group := range groups {
	//	fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	//}

	//configure a new gobot with drivers and adapter for leap motion
	gbot := gobot.NewGobot()
	leapMotionAdaptor := leap.NewLeapMotionAdaptor("leap", "127.0.0.1:6437")
	l := leap.NewLeapMotionDriver(leapMotionAdaptor, "leap")

	//The work that actually gets done on the leap motion
	work := func() {
		count := 0
		//Runs this loop every time it gets an event back from magic leap that a hand is detected
		gobot.On(l.Event("hand"), func(data interface{}) {
			count++
			//Change this number in order to increase or decrease response time
			n := 100
			//Since the hand events are verbose only do the following if we recieve data 100 times from the leap motion device
			if (count % n) == 0 {
				log.Println("Hit Trigger at count", count)
				reply(rtm)
			}
			log.Println("Tiggers every ", n)
			//log.Println("Not within ", n ," times")
		})
	}

	// create a new gobot device and assign drivers and work functions
	robot1 := gobot.NewRobot("leapBot",
		[]gobot.Connection{leapMotionAdaptor},
		[]gobot.Device{l},
		work,
	)

	gbot.AddRobot(robot1)
	//start the device
	gbot.Start()

}

// This function
// Place the ID you printed out from the commented out function above here
func reply(rtm *slack.RTM) {
	rtm.SendMessage(rtm.NewOutgoingMessage("Intruder alert! Intruder alert! ", "GROUP ID for Room you want to send the message to"))
}