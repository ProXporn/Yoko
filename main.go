package main

import (
	"log"
	"net/http"
	"time"
        "os/exec"
        "bytes"
        "fmt"
	tb "gopkg.in/tucnak/telebot.v2"
        eval "github.com/apaxa-go/eval"

)


var (
 menu = &tb.ReplyMarkup{}
)


func Shellout(command string) string {
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd := exec.Command("bash", "-c", command)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    cmd.Run()
    return stdout.String()
}

    

func main() {
	b, err := tb.NewBot(tb.Settings{
		URL:       "",
		Token:     "5050904599:AAG-YrM6KN4EJJx8peQOn901qHhLCkFo5QA",
		Updates:   0,
		Poller:    &tb.LongPoller{Timeout: 10 * time.Second},
		ParseMode: "HTML",
		Reporter: func(error) {
		},
		Client: &http.Client{},
	})

	if err != nil {
		log.Fatal(err)
		return
        }
        b.Handle("/eval", func(m *tb.Message) {
          expr, err:= eval.ParseString(string(m.Payload), "")
          out := ""
          if err != nil{
             out := err.Error()
             }
          a := eval.Args{"fmt.Sprint": eval.MakeDataRegularInterface(fmt.Sprint), "bot": eval.MakeDataRegularInterface(b)}
          r, err := expr.EvalToInterface(a)
          if err != nil{
         	out := err.Error()
          } else {
             out := fmt.Sprint(r)
          }
          eval_out := fmt.Sprintf("<b>► EVALGo</b>\n%s\n\n<b>► OUTPUT\n<code>%d</code>", string(m.Payload), out)
          b.Reply(m, eval_out)
        })
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

        b.Handle("/sh", func(m *tb.Message) {
                if m.Sender.ID != 1833850637 {return}
                if string(m.Payload) == string("") {
                   b.Reply(m, "No CMD given.")
                   return
                  }
                out := Shellout(m.Payload)
                output := fmt.Sprintf("<code>Go#~: %s</code>", string(out))
                b.Reply(m, output)
        })
	b.Start()
}
