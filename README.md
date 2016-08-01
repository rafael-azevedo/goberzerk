**goberzerk**
----
  Slack Bot that uses *github.com/hybridgroup/gobot/platforms/leap* and *github.com/nlopes/slack* to send a slack bot message to a channel when someone is at your computer.

**General Setup**

This guide assumes you have Go language environment setup. If you do not please refer to this guide [here](https://golang.org/doc/install)

To setup a quick IDE for Go development use the guide [here](https://www.wolfe.id.au/2015/03/05/using-sublime-text-for-go-development/)
 
You will then need to get a slack api key you can create your own channel and have up to 10 bots per channel 

*  **Create a Slack Channel and Get an API key**
	* Go to Slack.com. Fill out an email and hit *Create New Team*
	* Once in your slack account hit the drop down next to your user and select *Apps & Integration*
	* Search and select *Bots*
	* Hit *Add Configuration* and give your bot a unique name
	* Click *Add bot integration* to finish adding the integration and View settings to get its API key

* **Add the bot to a channel**

	Once you have an api key and bot created go to the channel you wish to notify with the bot and add it to the channel using the following command

	`/invite @<your_bot_name>`

*  **Configure Code**
	*  Clone repo to a directory