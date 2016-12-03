// package main

// import (
//     "os"
//     "fmt"
//     "net/http"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Hello, World!!!!")
// }



// func main() {
//     PORT := ":"
//     PORT += "3000"
//     if os.Getenv("PORT") != "" {
//         PORT += os.Getenv("PORT")
//     }
//     fmt.Println("port is", PORT)
    
//     http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
//     http.ListenAndServe(PORT, nil)
// }

package main  
import (  
    "fmt"
    "net/http"
    "os" 
    "log"

	"github.com/line/line-bot-sdk-go/linebot"
)

func handler(w http.ResponseWriter, r *http.Request) {  
    fmt.Fprintf(w, "Hi there - this page was served using Go \\o/")
}
func main() {
    bot, err := linebot.New(
		"b81bfdb52cc7c587f9865f2a1c9b7939",
		"YKwBLutO/SIk5zq88hi/+C/fWMmRxBBQ2jUhDa789Rrjaiotb/Bc7fjeicFpLrWrulVDve4pQNSl0x6K/Nbite0B46fC/P8lQfCm3ub6NLFsJG9cUbp39RzVJ/3kQbaiCzcYYUEChWatthPq8G0QXwdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}

    http.HandleFunc("/", handler)
	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
        fmt.Printf("ping\n")
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
                    fmt.Printf("%v", message)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
    
	// This is just a sample code.
	// For actually use, you must support HTTPS by using `ListenAndServeTLS`, reverse proxy or etc.
    fmt.Printf("サーバーを起動しています...")
    
	if err := http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
