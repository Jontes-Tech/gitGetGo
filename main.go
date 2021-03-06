// Notes to Opensourcers:
// - The code is really messy, but it works.
// - The code is not optimized, but it works.
package main
import (
	"fmt"
	"net/http"
	"log"
    "io/ioutil"
	"os/user"
	"os/exec"
	"os"
	"time"
)
func ssh_key(w http.ResponseWriter, req *http.Request) {
    data, ignore_me_1 := ioutil.ReadFile("/home/jonte/.ssh/id_rsa.pub")
    if ignore_me_1 != nil {
		exec.Command("ssh-keygen -t rsa -q -P \"\"").Run()
    }
	fmt.Fprint(w, string(data))
}
func main_page(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "GitGetGo is running. Please visit jontes.page/gitget")
}
func config(w http.ResponseWriter, req *http.Request) {
	setupCorsResponse(&w, req)
	mail := req.URL.Query().Get("mail")
    name := req.URL.Query().Get("name")
	exec.Command("git config --global user.name \"" + name + "\"")
	exec.Command("git config --global user.email " + mail)
	fmt.Println("Thanks for using GitGetGo by Jonte, have a nice day.")
	time.Sleep(2 * time.Second)
	os.Exit(0)
}
func handleRequests() {
	http.HandleFunc("/get/ssh_key", ssh_key)
	http.HandleFunc("/", main_page)
	http.HandleFunc("/config", config)
	http.Handle("/ping/", http.StripPrefix("/ping/", http.FileServer(http.Dir("./ping"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
}
func main() {
	user, username_err := user.Current()
    if user.Username == "root" {
        log.Fatal("Do not run as root.")
	}
	if username_err != nil {
		log.Fatalln(username_err)
	}
	fmt.Println("Please visit jontes.page/gitget")
	handleRequests()
}