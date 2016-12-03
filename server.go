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
)
func handler(w http.ResponseWriter, r *http.Request) {  
    fmt.Fprintf(w, "Hi there - this page was served using Go \\o/")
}
func main() {  
    http.HandleFunc("/", handler)
    http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
}