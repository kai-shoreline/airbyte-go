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
	// token := "WdxR8VoUhXWRjc7O_KFB2YwMLLpK31BgdkuwduBqCQo"
	// createSite(token)
	
	
	// createEnvVars("630fe0bf52cd4b08420e2721", "3774be8e-3f12-4231-b7b9-d79a2363e6d8", "tes2321", "tes3222") 

	createDnsZone("cioc-shoreline", "", "test5-insights.shoreline.io")

	// createDomain("3774be8e-3f12-4231-b7b9-d79a2363e6d8", "test2-insights.shoreline.io")

}

func createDnsZone(account_slug string, site_id string, name string) {

	url := "https://app.netlify.com/access-control/bb-api/api/v1/dns_zones"
	payload := strings.NewReader(`{
		"name": "test5-insights.shoreline.io",
		"account_slug": "cioc-shoreline"
	}`)
	body := postAPI(url, payload, "WdxR8VoUhXWRjc7O_KFB2YwMLLpK31BgdkuwduBqCQo")

	fmt.Println(string(body))
	var data map[string]interface{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", data)
	
}

func createDomain(site_id string, domain string) {
	url := "https://app.netlify.com/access-control/bb-api/api/v1/sites/" + site_id
	payload := strings.NewReader(`{"custom_domain":"test2-insights.shoreline.io"}`)

	body := putAPI(url, payload, "WdxR8VoUhXWRjc7O_KFB2YwMLLpK31BgdkuwduBqCQo")

	fmt.Println(string(body))
	var data map[string]interface{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", data)

}

func createEnvVars(account_id string, site_id string, key string, value string) {
	url := "https://app.netlify.com/access-control/bb-api/api/v1/accounts/" + account_id + "/env?site_id=" + site_id
  
	payload := strings.NewReader(fmt.Sprintf(`[
	  {
		  "key": "%s",
		  "scopes": [
			  "builds",
			  "functions",
			  "runtime",
			  "post_processing"
		  ],
		  "values": [
			  {
				  "context": "all",
				  "value": "%s"
			  }
		  ]
	  }
  	]`, key, value))
	body := postAPI(url, payload, "WdxR8VoUhXWRjc7O_KFB2YwMLLpK31BgdkuwduBqCQo")
	var data []interface{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", data)
	
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
  
  	body := postAPI(url, payload, token)

	fmt.Println(string(body))
	var data map[string]interface{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", data)
}


func postAPI(url string, payload io.Reader, token string) []byte {
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
	return body
  }

  func putAPI(url string, payload io.Reader, token string) []byte {
	client := &http.Client {
	}
	req, err := http.NewRequest("PUT", url, payload)
  
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
	return body
  }