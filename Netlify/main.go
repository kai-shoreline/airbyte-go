package main

import (
  "fmt"
  "io"
  "strings"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

func main() {
	createSite("WdxR8VoUhXWRjc7O_KFB2YwMLLpK31BgdkuwduBqCQo")
	



}

func createSite(token string) {
	url := "https://app.netlify.com/access-control/bb-api/api/v1/cioc-shoreline/sites"
	payload := strings.NewReader(`{
	  "account_slug": "cioc-shoreline",
	  "repo": {
		  "provider": "github",
		  "installation_id": 28900764,
		  "id": 532657498,
		  "owner_type": "Organization",
		  "repo": "shorelinesoftware/incident-reporting-website",
		  "private": true,
		  "plugins": [
			  {
				  "package": "@netlify/plugin-nextjs"
			  }
		  ],
		  "branch": "main",
		  "frameworkName": "Next.js",
		  "dir": ".next",
		  "cmd": "printenv > .env && npm run build",
		  "plugins_recommended": [
			  "@netlify/plugin-nextjs"
		  ],
		  "framework": "next"
	  },
	  "default_hooks_data": {},
	  "plugins": [
		  {
			  "package": "@netlify/plugin-nextjs"
		  }
	  ],
	  "created_via": ""
  }`)
  
	res := postAPI(url, payload, token)
	fmt.Printf("%v", res)
}


func postAPI(url string, payload io.Reader, token string) map[string]interface{} {
	client := &http.Client {
	}
	req, err := http.NewRequest("POST", url, payload)
  
	if err != nil {
	  fmt.Println(err)
	  return nil
	}
	if token != "" {
	  req.Header.Add("Authorization", "Bearer " + token) 
	}
	req.Header.Add("Content-Type", "application/json")
  
	res, err := client.Do(req)
	if err != nil {
	  fmt.Println(err)
	  return nil
	}
	defer res.Body.Close()
  
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	  fmt.Println(err)
	  return nil
	}
	fmt.Println(string(body))
	var data map[string]interface{}
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		panic(err)
	}
	return data
  }