package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := parseFlags()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		// run the tests
		os.Exit(m.Run())
	}
}

func TestParsing(t *testing.T) {
	if len(ec.Receptors) != 2 {
		t.Log(fmt.Sprintf("expected 2 receptors in the config, got %d instead", len(ec.Receptors)))
		t.Fail()
	} else if len(ec.Processors) != 2 {
		t.Log(fmt.Sprintf("expected 2 processors in the config, got %d instead", len(ec.Processors)))
		t.Fail()
	} else if len(ec.Routes) != 2 {
		t.Log(fmt.Sprintf("expected 2 routes in the config, got %d instead", len(ec.Routes)))
		t.Fail()
	}

	if rec, ok := ec.Receptors["example_irc_id"]; ok {
		if rec.ReceptorType != "IRCBot" {
			t.Log(fmt.Sprintf("incorrect receptor type for example_irc_id, expected \"IRCBot\", got \"%s\" instead", rec.ReceptorType))
			t.Fail()
		}
		if rec.Plugin != "path/to/emirc.so" {
			t.Log(fmt.Sprintf("incorrect receptor plugin path for example_irc_id, got \"%s\"", rec.Plugin))
			t.Fail()
		}
	} else {
		t.Log("did not find the example_irc_id receptor configuration")
		t.Fail()
	}

	if proc, ok := ec.Processors["emirc2tg"]; ok {
		if proc.Plugin != "path/to/emirc2tg.so" {
			t.Log(fmt.Sprintf("incorrect processor plugin path for emirc2tg, got \"%s\"", proc.Plugin))
			t.Fail()
		}
		if proc.Config != "path/to/emirc2tg.toml" {
			t.Log(fmt.Sprintf("incorrect processor config file path for emirc2tg, got \"%s\"", proc.Config))
			t.Fail()
		}
	} else {
		t.Log("did not find the emirc2tg processor configuration")
		t.Fail()
	}

	if ec.Router.Plugin != "path/to/emrouter.so" {
		t.Log(fmt.Sprintf("incorrect router plugin path, got \"%s\"", ec.Router.Plugin))
		t.Fail()
	}

	if routes, ok := ec.Routes["example_irc_id"]; ok {
		if len(routes.Destination) != 2 {
			t.Log(fmt.Sprintf("incorrect number of destinations for the example_irc_id route, expected 2, got %d", len(routes.Destination)))
			t.Fail()
		}
		if routes.Destination[0] != "emirc2tg" || routes.Destination[1] != "emirc_voice" {
			t.Log("incorrect values for destinations of the example_irc_id")
			t.Fail()
		}
	} else {
		t.Log("did not find the routes for the example_irc_id receptor")
		t.Fail()
	}

	if ircbot, ok := ec.IRCBots["example_irc_id"]; ok {
		if ircbot.Nick == nil {
			t.Log("the ircbot.Nick value is nil")
			t.Fail()
		} else if *ircbot.Nick != "emersyx" {
			t.Log(fmt.Sprintf("the ircbot.Nick value is incorrect, got \"%s\"", *ircbot.Nick))
			t.Fail()
		}

		if ircbot.Ident == nil {
			t.Log("the ircbot.Ident value is nil")
			t.Fail()
		} else if *ircbot.Ident != "emersyx" {
			t.Log(fmt.Sprintf("the ircbot.Ident value is incorrect, got \"%s\"", *ircbot.Ident))
			t.Fail()
		}

		if ircbot.Name == nil {
			t.Log("the ircbot.Name value is nil")
			t.Fail()
		} else if *ircbot.Name != "emersyx" {
			t.Log(fmt.Sprintf("the ircbot.Name value is incorrect, got \"%s\"", *ircbot.Name))
			t.Fail()
		}

		if ircbot.Version == nil {
			t.Log("the ircbot.Version value is nil")
			t.Fail()
		} else if *ircbot.Version != "alpha" {
			t.Log(fmt.Sprintf("the ircbot.Version value is incorrect, got \"%s\"", *ircbot.Version))
			t.Fail()
		}

		if ircbot.ServerAddress == nil {
			t.Log("the ircbot.ServerAddress value is nil")
			t.Fail()
		} else if *ircbot.ServerAddress != "chat.freenode.net" {
			t.Log(fmt.Sprintf("the ircbot.ServerAddress value is incorrect, got \"%s\"", *ircbot.ServerAddress))
			t.Fail()
		}

		if ircbot.ServerPort == nil {
			t.Log("the ircbot.ServerPort value is nil")
			t.Fail()
		} else if *ircbot.ServerPort != 6697 {
			t.Log(fmt.Sprintf("the ircbot.ServerPort value is incorrect, got %d", *ircbot.ServerPort))
			t.Fail()
		}

		if ircbot.ServerUseSSL == nil {
			t.Log("the ircbot.ServerUseSSL value is nil")
			t.Fail()
		} else if *ircbot.ServerUseSSL != true {
			t.Log(fmt.Sprintf("the ircbot.ServerUseSSL value is incorrect, got %t", *ircbot.ServerUseSSL))
			t.Fail()
		}

		if ircbot.QuitMessage == nil {
			t.Log("the ircbot.QuitMessage value is nil")
			t.Fail()
		} else if *ircbot.QuitMessage != "bye from emersyx!" {
			t.Log(fmt.Sprintf("the ircbot.QuitMessage value is incorrect, got \"%s\"", *ircbot.QuitMessage))
			t.Fail()
		}
	} else {
		t.Log("did not find the config for the example_irc_id IRCBot")
		t.Fail()
	}

	if telegrambot, ok := ec.TelegramBots["example_telegram_id"]; ok {
		if telegrambot.APIToken == nil {
			t.Log("the telegrambot.APIToken value is nil")
			t.Fail()
		} else if *telegrambot.APIToken != "Telegram Bot API token" {
			t.Log(fmt.Sprintf("the telegrambot.APIToken value is incorrect, got \"%s\"", *telegrambot.APIToken))
			t.Fail()
		}

		if telegrambot.UpdatesLimit == nil {
			t.Log("the telegrambot.UpdatesLimit value is nil")
			t.Fail()
		} else if *telegrambot.UpdatesLimit != 100 {
			t.Log(fmt.Sprintf("the telegrambot.UpdatesLimit value is incorrect, got %d", *telegrambot.UpdatesLimit))
			t.Fail()
		}

		if telegrambot.UpdatesTimeout == nil {
			t.Log("the telegrambot.UpdatesTimeout value is nil")
			t.Fail()
		} else if *telegrambot.UpdatesTimeout != 60 {
			t.Log(fmt.Sprintf("the telegrambot.UpdatesTimeout value is incorrect, got %d", *telegrambot.UpdatesTimeout))
			t.Fail()
		}

		if telegrambot.UpdatesAllowed == nil {
			t.Log("the telegrambot.UpdatesAllowed value is nil")
			t.Fail()
		} else if len(*telegrambot.UpdatesAllowed) != 9 {
			t.Log(fmt.Sprintf("the telegrambot.UpdatesAllowed array has invalid length, got %d", len(*telegrambot.UpdatesAllowed)))
			t.Fail()
		} else {
			expectedUpdates := []string{
				"message",
				"edited_message",
				"channel_post",
				"edited_channel_post",
				"inline_query",
				"chosen_inline_result",
				"callback_query",
				"shipping_query",
				"pre_checkout_query",
			}
			for i, u := range *telegrambot.UpdatesAllowed {
				if u != expectedUpdates[i] {
					t.Log(fmt.Sprintf("the telegrambot.UpdatesAllowed array has an invalid value at index %d, got \"%s\"", i, (*telegrambot.UpdatesAllowed)[i]))
					t.Fail()
				}
			}
		}
	} else {
		t.Log("did not find the config for the example_telegram_id TelegramBot")
		t.Fail()
	}
}
