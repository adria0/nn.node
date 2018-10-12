package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"io/ioutil"
	"bytes"
	"os"
)

func main() {

	if len(os.Args)!=2 {
		log.Error("No key argument")
		return
	}


	http.HandleFunc("/"+os.Args[1], func(w http.ResponseWriter, r *http.Request) {

		fail := func(err error) {
			w.Header().Set("Content-Type", "text/plain")
			log.Error(err)
			w.Write([]byte(err.Error()))
		}

		inbody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fail(err)
			return
		}
		req, err := http.NewRequest("POST", "http://localhost:8545", bytes.NewReader(inbody))
		if err != nil {
			fail(err)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fail(err)
			return
		}

		defer resp.Body.Close()
		outbody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fail(err)
			return
		}
    		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With");
   		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(outbody)
		log.Info(string(inbody)," -> ",string(outbody))
	})

	log.Info("Starting serving requests")
	http.ListenAndServeTLS(":5000",
		"/etc/letsencrypt/live/nou.network/fullchain.pem",
		"/etc/letsencrypt/live/nou.network/privkey.pem",
		 nil)
}
