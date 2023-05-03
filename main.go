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

  err = cmd.Process.Kill() // stop the command
  if err != nil {
      panic(err)
  }

}

func getWorkspaceId() {
  fmt.Println("test workspace")

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

