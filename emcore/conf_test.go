package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	parseFlags()
	loadConfig()
	os.Exit(m.Run())
}

func TestParsing(t *testing.T) {
	if len(ec.IRCGateways) != 1 {
		t.Log(fmt.Sprintf("expected 1 irc gateway in the config, got %d instead", len(ec.IRCGateways)))
		t.Fail()
	}
	if len(ec.TelegramGateways) != 1 {
		t.Log(fmt.Sprintf("expected 1 telegram gateway in the config, got %d instead", len(ec.TelegramGateways)))
		t.Fail()
	}
	if len(ec.Processors) != 2 {
		t.Log(fmt.Sprintf("expected 2 processors in the config, got %d instead", len(ec.Processors)))
		t.Fail()
	}
	if len(ec.Routes) != 2 {
		t.Log(fmt.Sprintf("expected 2 routes in the config, got %d instead", len(ec.Routes)))
		t.Fail()
	}
	if t.Failed() {
		return
	}

	proc := ec.Processors[0]
	if proc.PluginPath != "path/to/emi2t.so" {
		t.Log(fmt.Sprintf("incorrect processor plugin path for emi2t, got \"%s\"", proc.PluginPath))
		t.Fail()
	}
	if proc.Identifier != "emi2t" {
		t.Log(fmt.Sprintf("incorrect processor identifier for emi2t, got \"%s\"", proc.Identifier))
		t.Fail()
	}
	if proc.Config != "path/to/emi2t.toml" {
		t.Log(fmt.Sprintf("incorrect processor config file path for emi2t, got \"%s\"", proc.Config))
		t.Fail()
	}

	if ec.Router.PluginPath != "path/to/emrouter.so" {
		t.Log(fmt.Sprintf("incorrect router plugin path, got \"%s\"", ec.Router.PluginPath))
		t.Fail()
	}

	rt := ec.Routes[0]
	if rt.Source != "example_irc_id" {
		t.Log(fmt.Sprintf("incorrect values for the source, got \"%d\"", len(rt.Source)))
		t.Fail()
	}
	if len(rt.Destination) != 2 {
		t.Log(fmt.Sprintf("incorrect number of destinations for the example_irc_id route, expected 2, got %d", len(rt.Destination)))
		t.Fail()
	}
	if rt.Destination[0] != "emi2t" || rt.Destination[1] != "emirc_voice" {
		t.Log("incorrect values for destinations of the example_irc_id")
		t.Fail()
	}

	ircgw := ec.IRCGateways[0]
	if ircgw.PluginPath == nil {
		t.Log("the ircgw.PluginPath value is nil")
		t.Fail()
	} else if *ircgw.PluginPath != "path/to/emirc.so" {
		t.Log(fmt.Sprintf("the ircgw.PluginPath value is incorrect, got \"%s\"", *ircgw.PluginPath))
		t.Fail()
	}

	if ircgw.Identifier == nil {
		t.Log("the ircgw.Identifier value is nil")
		t.Fail()
	} else if *ircgw.Identifier != "example_irc_id" {
		t.Log(fmt.Sprintf("the ircgw.Identifier value is incorrect, got \"%s\"", *ircgw.Identifier))
		t.Fail()
	}

	if ircgw.Nick == nil {
		t.Log("the ircgw.Nick value is nil")
		t.Fail()
	} else if *ircgw.Nick != "emersyx" {
		t.Log(fmt.Sprintf("the ircgw.Nick value is incorrect, got \"%s\"", *ircgw.Nick))
		t.Fail()
	}

	if ircgw.Ident == nil {
		t.Log("the ircgw.Ident value is nil")
		t.Fail()
	} else if *ircgw.Ident != "emersyx" {
		t.Log(fmt.Sprintf("the ircgw.Ident value is incorrect, got \"%s\"", *ircgw.Ident))
		t.Fail()
	}

	if ircgw.Name == nil {
		t.Log("the ircgw.Name value is nil")
		t.Fail()
	} else if *ircgw.Name != "emersyx" {
		t.Log(fmt.Sprintf("the ircgw.Name value is incorrect, got \"%s\"", *ircgw.Name))
		t.Fail()
	}

	if ircgw.Version == nil {
		t.Log("the ircgw.Version value is nil")
		t.Fail()
	} else if *ircgw.Version != "alpha" {
		t.Log(fmt.Sprintf("the ircgw.Version value is incorrect, got \"%s\"", *ircgw.Version))
		t.Fail()
	}

	if ircgw.ServerAddress == nil {
		t.Log("the ircgw.ServerAddress value is nil")
		t.Fail()
	} else if *ircgw.ServerAddress != "chat.freenode.net" {
		t.Log(fmt.Sprintf("the ircgw.ServerAddress value is incorrect, got \"%s\"", *ircgw.ServerAddress))
		t.Fail()
	}

	if ircgw.ServerPort == nil {
		t.Log("the ircgw.ServerPort value is nil")
		t.Fail()
	} else if *ircgw.ServerPort != 6697 {
		t.Log(fmt.Sprintf("the ircgw.ServerPort value is incorrect, got %d", *ircgw.ServerPort))
		t.Fail()
	}

	if ircgw.ServerUseSSL == nil {
		t.Log("the ircgw.ServerUseSSL value is nil")
		t.Fail()
	} else if *ircgw.ServerUseSSL != true {
		t.Log(fmt.Sprintf("the ircgw.ServerUseSSL value is incorrect, got %t", *ircgw.ServerUseSSL))
		t.Fail()
	}

	if ircgw.QuitMessage == nil {
		t.Log("the ircgw.QuitMessage value is nil")
		t.Fail()
	} else if *ircgw.QuitMessage != "bye from emersyx!" {
		t.Log(fmt.Sprintf("the ircgw.QuitMessage value is incorrect, got \"%s\"", *ircgw.QuitMessage))
		t.Fail()
	}

	tggw := ec.TelegramGateways[0]
	if tggw.APIToken == nil {
		t.Log("the tggw.APIToken value is nil")
		t.Fail()
	} else if *tggw.APIToken != "Telegram Bot API token" {
		t.Log(fmt.Sprintf("the tggw.APIToken value is incorrect, got \"%s\"", *tggw.APIToken))
		t.Fail()
	}

	if tggw.UpdatesLimit == nil {
		t.Log("the tggw.UpdatesLimit value is nil")
		t.Fail()
	} else if *tggw.UpdatesLimit != 100 {
		t.Log(fmt.Sprintf("the tggw.UpdatesLimit value is incorrect, got %d", *tggw.UpdatesLimit))
		t.Fail()
	}

	if tggw.UpdatesTimeout == nil {
		t.Log("the tggw.UpdatesTimeout value is nil")
		t.Fail()
	} else if *tggw.UpdatesTimeout != 60 {
		t.Log(fmt.Sprintf("the tggw.UpdatesTimeout value is incorrect, got %d", *tggw.UpdatesTimeout))
		t.Fail()
	}

	if tggw.UpdatesAllowed == nil {
		t.Log("the tggw.UpdatesAllowed value is nil")
		t.Fail()
	} else if len(*tggw.UpdatesAllowed) != 9 {
		t.Log(fmt.Sprintf("the tggw.UpdatesAllowed array has invalid length, got %d", len(*tggw.UpdatesAllowed)))
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
		for i, u := range *tggw.UpdatesAllowed {
			if u != expectedUpdates[i] {
				t.Log(fmt.Sprintf("the tggw.UpdatesAllowed array has an invalid value at index %d, got \"%s\"", i, (*tggw.UpdatesAllowed)[i]))
				t.Fail()
			}
		}
	}
}
