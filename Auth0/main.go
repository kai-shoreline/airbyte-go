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
	domain := "https://test3-insight.us.auth0.com"
	clientId := "yHCr3AuTyFDzCpLIEm5bGDBlwCkXUnli"
	clientSecret := "3TD7qs3JicpWOlQ-C1lL1DAnxRE4nLQvD1KhybF8hjxvZqg3iSTQI41LhrHUwhQr"
	token := getAPIToken(domain, clientId, clientSecret)
    // fmt.Printf(token)	

	rule1 := "add email to token"
	script1 := `function (user, context, callback) {\n\tcontext.accessToken['https://insights.shoreline.io/email'] = user.email;\n  context.accessToken['https://insights.shoreline.io/name'] = user.name;\n  return callback(null, user, context);\n}`
	addRules(domain, rule1, script1, token)

	rule2 := "add teammate"
	hostApi := "https://insights-jk.api.project2022-staging.shoreline.io"
	script2 := fmt.Sprintf(`function (user, context, callback) {\n  const axios = require('axios@0.19.2');\n  \n  if (context.request.query.invited === 'true') {\n    const options = { method: 'POST',\n    \turl: \"%s/api/v2/users/teams/default/invitations/accept\",\n    \theaders: { 'content-type': 'application/json' },\n    \tdata: JSON.stringify({ \n        user_id: context.request.query.user_id, \n        auth0_id: user.user_id, \n        name: user.name,\n        invitation_id: context.request.query.invitation_id\n      })\n     };\n  \taxios(options);\n  }\n  return callback(null, user, context);\n}`, hostApi)
	addRules(domain, rule2, script2, token)






}

func addRules(domain string, name string, script string, token string) {
	url := domain + "/api/v2/rules"
	payload := strings.NewReader(fmt.Sprintf(`{
		"name": "%s",
		"script": "%s",
		"enabled": true
	}`, name, script))
	fmt.Printf("%+v\n",payload)
	res := postAPI(url, payload, token)
	fmt.Printf("%+v\n",res)
}

func getAPIToken(domain string, clientId string, clientSecret string ) string {
	payload := strings.NewReader(fmt.Sprintf("{\"client_id\":\"%s\",\"client_secret\":\"%s\",\"audience\":\"%s/api/v2/\",\"grant_type\":\"client_credentials\"}", clientId, clientSecret, domain))
	url := domain + "/oauth/token"
	res := postAPI(url, payload, "")
	return res["access_token"].(string)
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

  var data map[string]interface{}
  err = json.Unmarshal([]byte(body), &data)
  if err != nil {
      panic(err)
  }
  return data
}