package cwmp

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CwmpService() {
	http.HandleFunc("/", myHandle)
	http.ListenAndServe(":7547", nil)
}

func myHandle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	con, _ := ioutil.ReadAll(r.Body) //获取post的数据
	fmt.Println(string(con))

}