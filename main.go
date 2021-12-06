package main

import (
	"log"

	tb "gopkg.in/tucnak/telebot.v2"
        tl "github.com/xelaj/mtproto/telegram"
)

var (
	menu = &tb.ReplyMarkup{}
)

var (
	b, err = tb.NewBot(tb.Settings{
		Token:       "5050904599:AAG-YrM6KN4EJJx8peQOn901qHhLCkFo5QA",
		Synchronous: false,
		Poller:      &tb.LongPoller{Timeout: 10},
		ParseMode:   "HTML",
	})
)

const (
 appID 4529547
 appHash "55bc2f0ca39d588ce5471e52acbf5a69"
 botToken "5050904599:AAG-YrM6KN4EJJx8peQOn901qHhLCkFo5QA"
)


func main() {
	if err != nil {
		log.Fatal(err)
		return
	}
	b.Handle("/eval", evaluate)
	b.Handle("/start", func(m *tb.Message) {
		if m.Private() {
			menu.Inline(
				menu.Row(menu.URL("Support", "t.me/roseloverx_support"), menu.URL("Updates", "t.me/roseloverx_support")),
				menu.Row(menu.Data("Commands", "help_menu")),
				menu.Row(menu.URL("Add me to your group", "https://t.me/Yoko_Robot?startgroup=true")))
			b.Send(m.Sender, "Hey there! I am <b>Yoko</b>.\nIm an Anime themed Group Management Bot, feel free to add me to your groups!", menu)
			return
		}
		b.Reply(m, "Hey I'm Alive.")
	})
	b.Handle("/info", info)
	b.Handle("/sh", execute)
	b.Handle("/ban", ban)
        b.Handle("/ux", u)
	b.Start()
}
