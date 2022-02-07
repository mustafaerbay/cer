package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/viper"
	"github.com/xuri/excelize/v2"

	_ "cer/config"
	"cer/modals"
	"fmt"
)

var (
	Info       = Teal
	Warn       = Yellow
	Fata       = Red
	Suggestion = Green
	Bug        = Red
	Dev        = Teal
	Test       = Yellow
	Us         = Blue
	Title      = White
	Sep        = Lightblue
)

var (
	Black     = Color("\033[1;30m%s\033[0m")
	Red       = Color("\033[1;31m%s\033[0m")
	Green     = Color("\033[1;32m%s\033[0m")
	Yellow    = Color("\033[1;33m%s\033[0m")
	Blue      = Color("\033[1;34m%s\033[0m")
	Magenta   = Color("\033[1;35m%s\033[0m")
	Teal      = Color("\033[1;36m%s\033[0m")
	White     = Color("\033[1;37m%s\033[0m")
	Lightblue = Color("\033[1;96m%s\033[0m")
)

type Person struct {
	Name string `*Info,json:"name"`
}

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func Joinstr(element ...string) string {
	return strings.Join(element, "&")
}

type AddHeaderTransport struct {
	T http.RoundTripper
}

func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("PRIVATE-TOKEN", viper.GetString("project.personal_access_token"))
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
		Timeout:   10 * time.Second,
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
	Issues    []modals.Issue
)

// TODO: https://stackoverflow.com/questions/55440854/how-do-i-return-data-from-a-for-loop-in-go
func GetIssuesByName(client *http.Client, method string, endpoint string, perPage string, userid string, username string) {
	// m := make(map[string]string)
	as := "per_page=" + perPage + "&" + "assignee_username=" + userid + "&not[labels]=Verified,Status%3A+Invalid"
	endpoint = endpoint + as
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
	json.Unmarshal(responseBody, &Issues)
	// fmt.Println(Info("hello, world!"))

	fmt.Printf("%v [%v]", White(username), Bug(len(Issues)))
	fmt.Println("")
	now := time.Now()

	for i, v := range Issues {
		fmt.Printf("#%v - ID: %v, Title: %v \n", i+1, v.IID, Title(v.Title))

		// fmt.Printf("ID:\t",v.IID)
		for _, k := range v.Labels {
			switch k {
			case "Type: Bug":
				fmt.Println(Bug("BUG"))
			case "Suggestion":
				fmt.Println(Suggestion("Suggestion"))
			case "Status: In Dev":
				fmt.Println(Dev("Status: In Dev"))
			case "Status: In Test":
				fmt.Println(Test("Status: In Test"))
			case "US":
				fmt.Println(Us("US"))
			default:
				fmt.Println(k)
			}
		}
		// d := time.
		// fmt.Println(now.Sub(v.CreatedAt).Round(d))
		fmt.Println("Created at: \t", v.CreatedAt.Format("2006-01-01"), "--->", Red(math.Round(now.Sub(v.CreatedAt).Hours()/24), " Days Ago"))
		fmt.Println("Last Updated at: \t\t", Red(math.Round(now.Sub(v.UpdatedAt).Hours()/24), " Days Ago"))
		// fmt.Println("Updated at: ", v.UpdatedAt.Format("2006-01-01"))
		fmt.Println("Created by: ", v.Author.Name)
		fmt.Println("LINK:\t", v.WebURL)
		fmt.Println(Sep("-----------------------------------------------------------------------------------------"))

	}
	// list := make([]*modals.Issue, len(responseBody))
	// json.Unmarshal(responseBody, &IssueList.ManyIssues)
	//  fmt.Println(issueBody)
	//  return responseBody
	fmt.Println(Red("==========================================================================================================================================="))
}

func GetCol(f *excelize.File, sheet string, col rune) ([]string, error) {
    colIndex := int(col) - 'A'
    cols, err := f.GetCols(sheet)
    if err != nil {
        return nil, err
    }
    if colIndex >= len(cols) {
        return nil, errors.New("column index out of bound")
    }
    return cols[colIndex], nil
}

