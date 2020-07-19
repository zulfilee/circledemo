package main

import (
	"fmt"
	"net/http"
)

//var airbrake = gobrake.NewNotifierWithOptions(&gobrake.NotifierOptions{
//	ProjectId:   252405,
//	ProjectKey:  "c032c12b1096cb8fec0dabf7c186ad9a",
//	Environment: "dev",
//})

//func init() {
//    airbrake.AddFilter(func(notice *gobrake.Notice) *gobrake.Notice {
//       notice.Params["user"] = map[string]string{
//          "id": "1",
//        "username": "johnsmith",
//      "name": "John Smith",
// }
//return notice
// })
//}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from Spectro!")
}

func main() {
	//	defer airbrake.Close()
	//	defer airbrake.NotifyOnPanic()

	//	airbrake.Notify(errors.New("operation failed"), nil)
	http.HandleFunc("/hello-backend", handler)
	fmt.Println("Server running...")

	http.ListenAndServe(":8080", nil)
	// new comment added

}
