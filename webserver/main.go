package main

import (
	"errors"
	"fmt"
	"github.com/airbrake/gobrake"
	"net/http"
)

var airbrake = gobrake.NewNotifierWithOptions(&gobrake.NotifierOptions{
	ProjectId:   252403,
	ProjectKey:  "302de0f68f1987b32b080b1ccf647ea4",
	Environment: "production",
})

func init() {
    airbrake.AddFilter(func(notice *gobrake.Notice) *gobrake.Notice {
        notice.Params["user"] = map[string]string{
            "id": "1",
            "username": "johnsmith",
            "name": "John Smith",
        }
        return notice
    })
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from Spectro!")
}

func main() {
	defer airbrake.Close()
	defer airbrake.NotifyOnPanic()

	airbrake.Notify(errors.New("operation failed"), nil)
	http.HandleFunc("/hello-backend", handler)
	fmt.Println("Server running...")

	http.ListenAndServe(":8080", nil)
	// new comment added

}
