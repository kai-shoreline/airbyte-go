package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "os/exec"
  "strings"
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

  workspaceId := getWorkspaceId()
  fmt.Printf("%+v\n",workspaceId)

  pugerdutyDefinitionId := getSourceDefinitionID(workspaceId, "PD hist main1669083192-test2", "main1669083192", "892815091625.dkr.ecr.us-west-2.amazonaws.com/pagerduty-airbyte")
  fmt.Printf("%+v\n",pugerdutyDefinitionId)

  opsgenieDefinitionId := getSourceDefinitionID(workspaceId, "opsgenie-initial1672253312", "opsgenie-initial1672253312", "892815091625.dkr.ecr.us-west-2.amazonaws.com/opsgenie-airbyte")
  fmt.Printf("%+v\n",opsgenieDefinitionId)

  postgresDefinitionId := getDestinationDefinitionID(workspaceId, "postgres")
  fmt.Printf("%+v\n",postgresDefinitionId)

  definitionId := getDestinationID("25c5221d-dce2-4163-ade9-739ef790f503", workspaceId, "f1PrBXKfKOxSObGbQLGQvs29ck7V6RLXEjyR9bZWUppxqdKpYO", "insights-cluster-jk.cluster-c4wzjyxeyphq.us-east-1.rds.amazonaws.com")
  fmt.Printf("%+v\n",definitionId)
  
  err = cmd.Process.Kill() // stop the command
  if err != nil {
      panic(err)
  }

}

func getDestinationDefinitionID(workspaceId string, database string) string {
  url := "http://localhost:3000/api/v1/destination_definitions/list_latest"
  payload := strings.NewReader(fmt.Sprintf(`{"workspaceId":"%s"}`, workspaceId))
  // fmt.Printf("%+v\n",payload)
  res := postAPI(url, payload)
  return res["destinationId"].(string)  
}


func getWorkspaceId() string {
  url := "http://localhost:3000/api/v1/workspaces/list"
  res := postAPI(url, nil)

  data1 := res["workspaces"].([]interface{})[0].(map[string]interface {})
  return data1["workspaceId"].(string)
}

func getSourceDefinitionID(workspaceId string, name string, dockerImageTag string, dockerRepository string) string {

  url := "http://localhost:3000/api/v1/source_definitions/create_custom"
  payload := strings.NewReader(fmt.Sprintf(`{
    "workspaceId": "%s",
    "sourceDefinition": {
        "name": "%s",
        "documentationUrl": "",
        "dockerImageTag": "%s",
        "dockerRepository": "%s"
    }
  }`, workspaceId, name, dockerImageTag, dockerRepository))
  res := postAPI(url, payload)
  return res["sourceDefinitionId"].(string)
}


func getDestinationID(destinationDefinitionId string, workspaceId string, password string, host string) string {
  url := "http://localhost:3000/api/v1/destinations/create"
  payload := strings.NewReader(fmt.Sprintf(`{"name":"Postgres","destinationDefinitionId":"%s","workspaceId":"%s","connectionConfiguration":{"tunnel_method":{"tunnel_method":"NO_TUNNEL"},"username":"postgres","ssl_mode":{"mode":"disable"},"password":"%s","database":"postgres","schema":"public","port":5432,"host":"%s","ssl":false}}`, destinationDefinitionId, workspaceId, password, host))
  // fmt.Printf("%+v\n",payload)
  res := postAPI(url, payload)
  return res["destinationId"].(string)
}


func postAPI(url string, payload io.Reader) map[string]interface{} {
  client := &http.Client {
  }
  req, err := http.NewRequest("POST", url, payload)

  if err != nil {
    fmt.Println(err)
    return nil
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