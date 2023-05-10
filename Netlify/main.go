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
	//creat new team at first, and get token and team name
	
	teamName := "Shoreline"
	token := "WdxR8VoUhXWRjc7O_KFB2YwMLLpK31BgdkuwduBqCQo"
	domain := "test7-insights.shoreline.io"
	account := getAccount(token, teamName)
	
	account_slug := account[0]
	account_id := account[1]
	createDnsZone(account_slug, domain, token)
	res := createSite(account_slug, domain, token)
	site_id := res["site_id"].(string)
	// fmt.Printf("%v", site_id)
	API_BASE_URL := "https://insights-jk.api.project2022-staging.shoreline.io"
	API_URL := "https://insights-jk.api.project2022-staging.shoreline.io"
	AUTH0_AUDIENCE := "https://insights-jk.api.project2022-staging.shoreline.io"
	AUTH0_BASE_URL := "test-insights.shoreline.io"
	AUTH0_CLIENT_ID := "PRVTJ1MLUgKf9gks3HqcfAriRI7L43mz"
	AUTH0_CLIENT_SECRET := "1TNtUixWhqkhc5sFcplJI5Pe47DFGRWqNHv2SxSdosEdprXBzfIMLHMC08DDEcV1"
	AUTH0_DOMAIN := "jk-insights-dev.us.auth0.com"
	AUTH0_ISSUER_BASE_URL := "https://jk-insights-dev.us.auth0.com"
	AUTH0_SECRET := "c8dd15225354b1f35c0693b88a9e217bcdf69c8bdd95d133a1944ec4b2f7d147"
	AWS_INTEGRATION_ENABLED := "false"
	ENABLE_RUNBOOK := "true"
	SENTRY_AUTH_TOKEN := "7807ddf1645d4bdb8883eee8821636880b69b7bfbd794fe1ae33b4d861a7fb12"
	SHORELINE_INTEGRATION_ENABLED := "true"

	
	createEnvVars(account_id, site_id, "API_BASE_URL", API_BASE_URL)
	createEnvVars(account_id, site_id, "API_URL", API_URL)
	createEnvVars(account_id, site_id, "AUTH0_AUDIENCE", AUTH0_AUDIENCE)
	createEnvVars(account_id, site_id, "AUTH0_BASE_URL", AUTH0_BASE_URL)
	createEnvVars(account_id, site_id, "AUTH0_CLIENT_ID", AUTH0_CLIENT_ID)
	createEnvVars(account_id, site_id, "AUTH0_CLIENT_SECRET", AUTH0_CLIENT_SECRET)
	createEnvVars(account_id, site_id, "AUTH0_DOMAIN", AUTH0_DOMAIN)
	createEnvVars(account_id, site_id, "AUTH0_ISSUER_BASE_URL", AUTH0_ISSUER_BASE_URL)
	createEnvVars(account_id, site_id, "AUTH0_SECRET", AUTH0_SECRET)
	createEnvVars(account_id, site_id, "AWS_INTEGRATION_ENABLED", AWS_INTEGRATION_ENABLED)
	createEnvVars(account_id, site_id, "ENABLE_RUNBOOK", ENABLE_RUNBOOK)
	createEnvVars(account_id, site_id, "SENTRY_AUTH_TOKEN", SENTRY_AUTH_TOKEN)
	createEnvVars(account_id, site_id, "SHORELINE_INTEGRATION_ENABLED", SHORELINE_INTEGRATION_ENABLED)
}

func getAccount(token string, teamName string) [2]string {
	url := "https://app.netlify.com/access-control/bb-api/api/v1/accounts"
	body := useAPI(url, nil, token)
	// fmt.Println(string(body))
	var data []interface{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		panic(err)
	}
	var res [2]string
	for _, v := range data {
		if v.(map[string]interface {})["name"] == teamName {
			res[0] = v.(map[string]interface {})["slug"].(string)
			res[1] = v.(map[string]interface {})["id"].(string)
			return res
		}		
	}
	fmt.Printf("%+v", res)
	return res

}

func createDnsZone(account_slug string, name string, token string) {

	url := "https://app.netlify.com/access-control/bb-api/api/v1/dns_zones"
	payload := strings.NewReader(fmt.Sprintf(`{
		"name": "%s",
		"account_slug": "%s"
	}`,name, account_slug))
	body := useAPI(url, payload, token)

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
	body := useAPI(url, payload, "WdxR8VoUhXWRjc7O_KFB2YwMLLpK31BgdkuwduBqCQo")
	var data []interface{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", data)
	
}

func createSite(account_slug string, domain string, token string) map[string]interface{} {
	url := "https://app.netlify.com/access-control/bb-api/api/v1/" + account_slug + "/sites"
	payload := strings.NewReader(fmt.Sprintf(`{
	  "account_slug": "%s",
	  "custom_domain": "%s",
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
  }`, account_slug, domain))
  
  	body := useAPI(url, payload, token)

	fmt.Println(string(body))
	var data map[string]interface{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", data)
	return data
}


func useAPI(url string, payload io.Reader, token string) []byte {
	client := &http.Client {
	}
	
	action := "POST"
	if payload == nil {
		action = "GET"
	} 
	req, err := http.NewRequest(action, url, payload)
  
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
