package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "os/exec"
  "time"
  "io"
  "encoding/json"
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

  // getSourceDefinitionID()

  // getSourceDefinitionID2()

  // getDestinationID()
  
  err = cmd.Process.Kill() // stop the command
  if err != nil {
      panic(err)
  }

}

func getWorkspaceId() {
  url := "http://localhost:3000/api/v1/workspaces/list"
  postAPI(url, nil)
}
 

func getSourceDefinitionID() {

  url := "http://localhost:3000/api/v1/source_definitions/create_custom"
  payload := strings.NewReader(`{
    "workspaceId": "b36bb3f7-e6f9-4105-9d2a-4abd07bc5c1a",
    "sourceDefinition": {
        "name": "PD hist main1669083192-test2",
        "documentationUrl": "",
        "dockerImageTag": "main1669083192",
        "dockerRepository": "892815091625.dkr.ecr.us-west-2.amazonaws.com/pagerduty-airbyte"
    }
  }`)
  postAPI(url, payload)
}


func getSourceDefinitionID2() {
  fmt.Println("test getSourceDefinitionID2")

  url := "http://localhost:3000/api/v1/source_definitions/create_custom"
  payload := strings.NewReader(`{
    "workspaceId": "b36bb3f7-e6f9-4105-9d2a-4abd07bc5c1a",
    "sourceDefinition": {
        "name": "opsgenie-initial1672253312",
        "documentationUrl": "",
        "dockerImageTag": "opsgenie-initial1672253312",
        "dockerRepository": "892815091625.dkr.ecr.us-west-2.amazonaws.com/opsgenie-airbyte"
    }
  }`)

  postAPI(url, payload)
}


func getDestinationID() {
  url := "http://localhost:3000/api/v1/destinations/create"
  payload := strings.NewReader(`{"name":"Postgres","destinationDefinitionId":"25c5221d-dce2-4163-ade9-739ef790f503","workspaceId":"b36bb3f7-e6f9-4105-9d2a-4abd07bc5c1a","connectionConfiguration":{"tunnel_method":{"tunnel_method":"NO_TUNNEL"},"username":"postgres","ssl_mode":{"mode":"disable"},"password":"f1PrBXKfKOxSObGbQLGQvs29ck7V6RLXEjyR9bZWUppxqdKpYO","database":"postgres","schema":"public","port":5432,"host":"insights-cluster-jk.cluster-c4wzjyxeyphq.us-east-1.rds.amazonaws.com","ssl":false}}`)
  postAPI(url, payload)
}


// type Person struct {
// 	workspaces string `json:"workspaces"`
// }

func postAPI(url string, payload io.Reader) {
  client := &http.Client {
  }
  req, err := http.NewRequest("POST", url, payload)

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

  var data map[string]interface{}
  err = json.Unmarshal([]byte(body), &data)
  if err != nil {
      panic(err)
  }
  // var data1 map[string]interface{}
  // data1 := data["workspaces"].([]map[string]interface{})
  // fmt.Printf("%+v\n", data1[0]["workspaceId"].(string))

  data1 := data["workspaces"].([]interface{})
  data2 := data1[0].(map[string]interface {})
  fmt.Printf("%+v\n", data2["workspaceId"])



	// var person Person
	// err = json.Unmarshal([]byte(body), &person)

  // fmt.Println(person.workspaces)


}