package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "os/exec"
  "time"
)

func main() {


  cmd := exec.Command("kubectl", "port-forward", "service/airbyte-webapp-svc", "3000:80")
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		return
	}
	fmt.Println("Command started successfully!")
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error waiting for command to finish:", err)
	}
  time.Sleep(10 * time.Second)

  getWorkspaceId()

  getSourceDefinitionID()

  getSourceDefinitionID2()

  getDestinationID()
  
  err = cmd.Process.Kill() // stop the command
  if err != nil {
      panic(err)
  }

}

func getWorkspaceId() {


  url := "http://localhost:3000/api/v1/workspaces/list"
  method := "POST"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("accept", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}
 

func getSourceDefinitionID() {

  url := "http://localhost:3000/api/v1/source_definitions/create_custom"
  method := "POST"

  payload := strings.NewReader(`{
    "workspaceId": "b36bb3f7-e6f9-4105-9d2a-4abd07bc5c1a",
    "sourceDefinition": {
        "name": "PD hist main1669083192-test2",
        "documentationUrl": "",
        "dockerImageTag": "main1669083192",
        "dockerRepository": "892815091625.dkr.ecr.us-west-2.amazonaws.com/pagerduty-airbyte"
    }
}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}


func getSourceDefinitionID2() {
  fmt.Println("test getSourceDefinitionID2")

  url := "http://localhost:3000/api/v1/source_definitions/create_custom"
  method := "POST"

  payload := strings.NewReader(`{
    "workspaceId": "b36bb3f7-e6f9-4105-9d2a-4abd07bc5c1a",
    "sourceDefinition": {
        "name": "opsgenie-initial1672253312",
        "documentationUrl": "",
        "dockerImageTag": "opsgenie-initial1672253312",
        "dockerRepository": "892815091625.dkr.ecr.us-west-2.amazonaws.com/opsgenie-airbyte"
    }
}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}


func getDestinationID() {
  url := "http://localhost:3000/api/v1/destinations/create"
  method := "POST"

  payload := strings.NewReader(`{"name":"Postgres","destinationDefinitionId":"25c5221d-dce2-4163-ade9-739ef790f503","workspaceId":"b36bb3f7-e6f9-4105-9d2a-4abd07bc5c1a","connectionConfiguration":{"tunnel_method":{"tunnel_method":"NO_TUNNEL"},"username":"postgres","ssl_mode":{"mode":"disable"},"password":"f1PrBXKfKOxSObGbQLGQvs29ck7V6RLXEjyR9bZWUppxqdKpYO","database":"postgres","schema":"public","port":5432,"host":"insights-cluster-jk.cluster-c4wzjyxeyphq.us-east-1.rds.amazonaws.com","ssl":false}}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}

func postAPI(url string, payload string) {
  client := &http.Client {
  }
  req, err := http.NewRequest("POST", url, strings.NewReader(payload))

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}