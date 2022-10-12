package httpclient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"strings"

	"github.com/kosha/freshservice-connector/pkg/models"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func makeHttpReq(apiKey string, req *http.Request) []byte {
	req.Header.Add("Authorization", "Basic "+basicAuth(apiKey, "X"))
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return bodyBytes
}

func getPageCount(apiKey string, path string) int {
	pageCount := 0
	for ok := true; ok; ok = pageCount > 0 {
		pageCount += 1
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			return 0
		}

		req.Header.Add("Authorization", "Basic "+basicAuth(apiKey, "X"))
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			return 0
		}
		defer resp.Body.Close()

		if len(resp.Header.Values("Link")) == 0 {
			return pageCount
		} else {
			link := strings.Split(resp.Header.Values("Link")[0], ";")
			path = strings.Trim(strings.Trim(link[0], "<"), ">")
		}
	}
	return 0
}

func GetAllTickets(url string, apiKey, pageNum string, getNumPages bool) (int, *models.Tickets) {
	if getNumPages {
		return getPageCount(apiKey, url+"/api/v2/tickets?page="+pageNum), nil
	}

	req, err := http.NewRequest("GET", url+"/api/v2/tickets?page="+pageNum, nil)

	if err != nil {
		return 0, nil
	}
	var tickets models.Tickets

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &tickets)
	if err != nil {
		return 0, nil
	}
	return 0, &tickets
}

func GetAgents(url string, apiKey string) *models.Agents {

	req, err := http.NewRequest("GET", url+"/api/v2/agents", nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var agents *models.Agents

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &agents)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return agents
}

func GetSingleAgent(url, id string, apiKey string) *models.SingleAgent {
	if id != "" {
		url = url + "/api/v2/agents/" + id

	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var agents *models.SingleAgent

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &agents)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return agents
}

func DeactivateSingleAgent(url, id string, apiKey string) *models.SingleAgent {
	if id != "" {
		url = url + "/api/v2/agents/" + id

	}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var agents *models.SingleAgent

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &agents)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return agents
}

func GetGroups(url string, apiKey string) *models.Groups {
	req, err := http.NewRequest("GET", url+"/api/v2/groups", nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var groups *models.Groups

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &groups)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return groups
}

func GetSingleGroup(url, id string, apiKey string) *models.SingleGroup {
	req, err := http.NewRequest("GET", url+"/api/v2/groups/"+id, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var group *models.SingleGroup

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &group)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return group
}

func GetAssets(url string, apiKey string) *models.Assets {
	req, err := http.NewRequest("GET", url+"/api/v2/assets", nil)
	if err != nil {
		return nil
	}
	var assets *models.Assets

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &assets)
	if err != nil {
		return nil
	}
	return assets
}

func GetSingleAsset(url, id string, apiKey string) *models.SingleAsset {
	req, err := http.NewRequest("GET", url+"/api/v2/assets/"+id, nil)
	if err != nil {
		return nil
	}
	var assets *models.SingleAsset

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &assets)
	if err != nil {
		return nil
	}
	return assets
}

func DeleteSingleAsset(url, id string, apiKey string) string {
	req, err := http.NewRequest("DELETE", url+"/api/v2/assets/"+id, nil)
	if err != nil {
		return ""
	}
	var assets string

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &assets)
	if err != nil {
		return ""
	}
	return assets
}

func GetProblems(url string, apiKey string) *models.Problems {
	req, err := http.NewRequest("GET", url+"/api/v2/problems", nil)
	if err != nil {
		return nil
	}
	var problems *models.Problems

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &problems)
	if err != nil {
		return nil
	}
	return problems
}

func GetSingleProblem(url, id string, apiKey string) *models.SingleProblem {
	req, err := http.NewRequest("GET", url+"/api/v2/problems/"+id, nil)
	if err != nil {
		return nil
	}
	var problem *models.SingleProblem

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &problem)
	if err != nil {
		return nil
	}
	return problem
}

func DeleteSingleProblem(url, id string, apiKey string) string {
	req, err := http.NewRequest("DELETE", url+"/api/v2/problems/"+id, nil)
	if err != nil {
		return err.Error()
	}
	var resp string

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return err.Error()
	}
	return resp
}

func SearchTickets(url string, apiKey, query string, page string) *models.Tickets {

	if page == "" {
		page = "1"
	}
	baseUrl := url + "/api/v2/tickets/filter?page=" + page + "&query=" + neturl.QueryEscape(query)
	/*if agentId != "" && status == 0 {
		baseUrl = baseUrl + "\"agent_id:" + agentId + "\""
	}
	if status != 0 && agentId == "" {
		baseUrl = baseUrl + "\"status:" + strconv.Itoa(status) + "\""
	}
	if status != 0 && agentId != "" {
		baseUrl = baseUrl + "\"agent_id:" + agentId + "%20OR%20status:" + strconv.Itoa(status) + "\""
	}
		if query != "" {
			baseUrl += query + "\""
		}
	*/
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil
	}
	var searchResult *models.Tickets

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &searchResult)
	if err != nil {
		return nil
	}
	return searchResult
}

func GetTicketsPerAgent(url string, apiKey, agentId string, page string) *models.SearchResults {

	if page == "" {
		page = "1"
	}
	baseUrl := url + "/api/v2/tickets/filter?page=" + page + "&query="

	if agentId != "" {
		baseUrl = baseUrl + "\"agent_id:" + agentId + "\""
	}
	fmt.Println(baseUrl)
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil
	}
	var searchResult *models.SearchResults

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &searchResult)
	if err != nil {
		return nil
	}
	return searchResult
}

func GetTicketsCreatedAt(url string, apiKey, createdAt string, page string) *models.SearchResults {

	if page == "" {
		page = "1"
	}
	baseUrl := url + "/api/v2/search/tickets?page=" + page + "&query="

	if createdAt != "" {
		baseUrl = baseUrl + "\"created_at:>'" + createdAt + "'\""
	}
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil
	}
	var searchResult *models.SearchResults

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &searchResult)
	if err != nil {
		return nil
	}
	return searchResult
}

func GetTicketsPerGroup(url string, apiKey, groupId string, page string) *models.SearchResults {

	if page == "" {
		page = "1"
	}
	baseUrl := url + "/api/v2/search/tickets?page=" + page + "&query="

	if groupId != "" {
		baseUrl = baseUrl + "\"group_id:" + groupId + "\""
	}
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil
	}
	var searchResult *models.SearchResults

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &searchResult)
	if err != nil {
		return nil
	}
	return searchResult
}

func GetSingleTicket(url, id string, apiKey string) *models.SingleTicket {
	req, err := http.NewRequest("GET", url+"/api/v2/tickets/"+id+"+?include=tags,requester,conversations,stats", nil)
	if err != nil {
		return nil
	}
	var ticket *models.SingleTicket

	res := makeHttpReq(apiKey, req)
	// Convert response body to target struct
	err = json.Unmarshal(res, &ticket)
	if err != nil {
		return nil
	}
	return ticket
}

func CreateTicket(url, apiKey string, ticket *models.Ticket) (*models.Ticket, error) {

	jsonReq, err := json.Marshal(ticket)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url+"/api/v2/tickets", bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp := makeHttpReq(apiKey, req)
	fmt.Println(string(resp))

	var ticket1 *models.Ticket
	err = json.Unmarshal(resp, &ticket1)

	if err != nil {
		return nil, err
	}
	return ticket1, nil
}

func RemoveTicket(url, id, apiKey string) (string, error) {
	req, err := http.NewRequest(http.MethodDelete, url+"/api/v2/tickets/"+id, nil)
	if err != nil {
		return "", err
	}
	bodyString := makeHttpReq(apiKey, req)
	return string(bodyString), nil
}
