package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/singleflight"
)

//singleflight已经作为sync的标准库了
//参考文章：https://medium.com/@vCabbage/go-avoid-duplicate-requests-with-sync-singleflight-311601b3068b

func main() {
	var requestGroup singleflight.Group

	http.HandleFunc("/github", func(w http.ResponseWriter, r *http.Request) {
		v, err, shared := requestGroup.Do("github", func() (interface{}, error) {
			return githubStatus()
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		status := v.(string)

		log.Printf("/github handler requst: status %q, shared result %t", status, shared)

		fmt.Fprintf(w, "GitHub Status: %q", status)
	})

	http.HandleFunc("/bitbucket", func(w http.ResponseWriter, r *http.Request) {
		// We can use the same singleflight.Group as long as we use a different key
		v, err, shared := requestGroup.Do("bitbucket", func() (interface{}, error) {
			return bitbucketStatus()
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		status := v.(string)

		log.Printf("/bitbucket handler requst: status %q, shared result %t", status, shared)

		fmt.Fprintf(w, "BitBucket Status: %q", status)
	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}

// githubStatus retrieves GitHub's API status
func githubStatus() (string, error) {
	log.Println("Making request to GitHub API")
	defer log.Println("Request to GitHub API Complete")

	time.Sleep(1 * time.Second)

	resp, err := http.Get("https://status.github.com/api/status.json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("github response: %s", resp.Status)
	}

	r := struct{ Status string }{}

	err = json.NewDecoder(resp.Body).Decode(&r)

	return r.Status, err
}

// bitbucketStatus retrieves BitBucket's API status
func bitbucketStatus() (string, error) {
	log.Println("Making request to BitBucket API")
	defer log.Println("Request to BitBucket API Complete")

	time.Sleep(1 * time.Second)

	resp, err := http.Get("https://status.bitbucket.org/api/v2/status.json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("github response: %s", resp.Status)
	}

	r := struct{ Status struct{ Description string } }{}

	err = json.NewDecoder(resp.Body).Decode(&r)

	return r.Status.Description, err
}
