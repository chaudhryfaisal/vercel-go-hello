package api

import (
	"net/http"
	"net/http/httputil"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	 dump(w, r)
}

func dump(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(requestDump))
	fmt.Printf("------------------------- %s -------------------------\n", time.Now().Format("2006-01-02 3:4:5 PM"))
	_ = r.Write(w) // Save a copy of this request for debugging.
}
