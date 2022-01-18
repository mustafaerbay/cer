package internal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/viper"

	"cer/modals"
	"fmt"
)

func Joinstr(element...string)string{
    return strings.Join(element, "&")
}


type AddHeaderTransport struct {
	T http.RoundTripper
}

func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("PRIVATE-TOKEN", viper.GetString("personal_access_token"))
	return adt.T.RoundTrip(req)
}

func NewAddHeaderTransport(T http.RoundTripper) *AddHeaderTransport {
	if T == nil {
		T = http.DefaultTransport
	}
	return &AddHeaderTransport{T}
}

func HttpClient() *http.Client {
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: NewAddHeaderTransport(nil),
	}
	return client
}

func SendRequest(client *http.Client, method string, endpoint string) []byte {
	// endpoint := "https://rnd-gitlab-eu.huawei.com/api/v4/projects/5674/issues?per_page=25"
      values := map[string]string{"foo": "baz"}
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Fatalf("Failed to marshal request: %v", err)
	}

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	return body
}

// Golang Variable Declaration Block
var (
	Issue modals.Issue
	Issues []modals.Issue
	IssueList modals.IssueList
)

//Convert *ItemList to *Item

// TODO: https://stackoverflow.com/questions/55440854/how-do-i-return-data-from-a-for-loop-in-go
func GetIssuesByName(client *http.Client, method string, endpoint string , perPage string, userid string, username string) {
	// m := make(map[string]string)
	as := "per_page="+perPage+"&"+"assignee_username="+userid +"&not[labels]=Verified,Status%3A+Invalid"
	endpoint = endpoint+as
	fmt.Println(endpoint)

	values := map[string]string{"foo": "baz"}
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Fatalf("Failed to marshal request: %v", err)
	}

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}
	json.Unmarshal(responseBody,&Issues)

	// fmt.Println("ISSUE OWNER",username)
	fmt.Printf("OWNER: %v [%v]",username,len(Issues))
	fmt.Println("")
	now := time.Now()
	
	for i, v := range Issues {
		fmt.Println("NO:\t",i+1)
		fmt.Println("ID:\t",v.IID)
		fmt.Println("TITLE:\t",v.Title)
		fmt.Println("PAST:\t",now.Sub(v.CreatedAt).Hours() / 24)
		fmt.Println("LINK:\t",v.WebURL)
		fmt.Println("-----")
	}
	// list := make([]*modals.Issue, len(responseBody))
	// json.Unmarshal(responseBody, &IssueList.ManyIssues)
	//  fmt.Println(issueBody)
	//  return responseBody
	fmt.Println("========================================================================================")
}


// func (s *modals.Issues) GetIssuesByName(repo_url string, c *http.Client) *modals.Issues {


// 	responseBody := SendRequest(c, http.MethodGet, repo_url)
// 	json.Unmarshal(responseBody, []modals.Issue)


// 	return modals.Issues{IssueList: issueBody}
// }

// func (b GroupIssueBoard) String() string {
// 	return Stringify(b)
// }