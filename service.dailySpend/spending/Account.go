package spending

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Account(w http.ResponseWriter, r *http.Request) {

	cacheClient, err := GetCacheClient()

	if err != nil {
		log.Fatal(err)
	}

	if r.Method == "GET" {

		accs := cacheClient.GetAccounts()

		jsonString, err := json.Marshal(accs)

		if err != nil {
			http.Error(w, "can't parse accounts", http.StatusBadRequest)
			log.Fatal(err)
		}

		w.Write(jsonString)

	} else if r.Method == "POST" {

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		accName := string(body[:])

		if accName == "" {
			_ = NewCacheClient(cacheClient.CacheURL, "", false)
		} else {
			_ = NewCacheClient(cacheClient.CacheURL, accName, true)
		}

	}

}
