package main

import (
	"os"

	"github.com/shomali11/slacker"
)

//handle used to operate on commands sent to the bot
func handle(request *slacker.Request, response slacker.ResponseWriter) {
	response.Reply("Hey!")
}

func main() {
	// set an environment variable on heroku for the api key
	api := os.Getenv("api")
	bot := slacker.NewClient(api)
	// default reply if command not valid, testbot is name of bot
	dr := "For a list of my function please type @testbot help"
	bot.Default(func(request *slacker.Request, response slacker.ResponseWriter) {
		response.Reply(dr)
	})
	// hello is what you say, description, then func to handle this
	bot.Command("hello", "Say hello", handle)
	err := bot.Listen()
	if err != nil {
		panic(err)
	}
}
