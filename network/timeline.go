package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func ShowInbox(c *gin.Context) {
	username := c.Request.Header["Username"]
	UniverseLock.Lock()
	fromIndex := universe.UsernameToIndex(username[0]) - 1
	c.JSON(200, gin.H{"inbox": universe.Inboxes[fromIndex]})
	UniverseLock.Unlock()
}
func ShowTimelines(c *gin.Context) {

	username := c.Param("username")
	UniverseLock.Lock()
	fromIndex := universe.UsernameToIndex(username) - 1
	c.JSON(200, gin.H{"profile": universe.Profile[fromIndex]})
	UniverseLock.Unlock()
}

func mapIt(tokens []string) (key, val string) {
	if len(tokens) < 2 {
		return "", ""
	}
	return tokens[0], tokens[1]
}
func mapJsonBody(c *gin.Context) map[string]string {
	defer c.Request.Body.Close()
	body, _ := ioutil.ReadAll(c.Request.Body)
	var j map[string]interface{}
	json.Unmarshal(body, &j)
	m := map[string]string{}
	for k, v := range j {
		m[k] = fmt.Sprintf("%v", v)
	}
	return m
}
func mapBody(c *gin.Context) map[string]string {
	defer c.Request.Body.Close()
	body, _ := ioutil.ReadAll(c.Request.Body)
	m := map[string]string{}
	for _, line := range strings.Split(string(body), "\n") {
		k, v := mapIt(strings.Split(line, "="))
		if k != "" {
			m[k] = v
		}
	}
	return m
}

type TimelineWrapper struct {
	// {"profile":[{"text":"hi2","from":"andrew","posted_at":1616960612,"origin":"localhost:8080"},{"text":"hi","from":"andrew","posted_at":1616960598,"origin":"localhost:8080"}]}
	Profile []Timeline `json:"profile"`
}

type Timeline struct {
	Text     string `json:"text"`
	From     string `json:"from"`
	PostedAt int64  `json:"posted_at"`
	Origin   string `json:"origin"`
}

func TimelineFromMap(m map[string]string) *Timeline {
	t := Timeline{}
	t.Text = m["text"]
	t.From = m["from"]
	t.PostedAt, _ = strconv.ParseInt(m["posted_at"], 10, 64)
	t.Origin = m["origin"]

	return &t
}

var UniverseLock sync.Mutex

func NotifyTimeline(c *gin.Context) {
	m := mapJsonBody(c)
	fmt.Println("mapJsonBody", m)
	//t := TimelineFromMap(m)
	//t.AddToByKey()
}

func CreateTimeline(c *gin.Context) {
	m := mapBody(c)
	t := Timeline{}
	t.Text = m["text"]
	t.From = m["username"]
	t.PostedAt = time.Now().Unix()
	t.Origin = globalInOut.Name

	if t.AddToUniverse() == true {
		//TellOutAboutNewTimeline(&t, globalInOut.Out)
	}
}

func (t *Timeline) AddToUniverse() bool {
	UniverseLock.Lock()
	defer UniverseLock.Unlock()
	universe.BroadcastNewTimeline(t)
	return true
}

func TellOutAboutNewTimeline(t *Timeline, out string) {
	os.Setenv("CLT_HOST", fmt.Sprintf("http://%s/", out))
	asBytes, _ := json.Marshal(t)
	DoPost("timelines/notify", asBytes)
	os.Setenv("CLT_HOST", "")
}

func DisplayTimelines(s string) {
	var tw TimelineWrapper
	json.Unmarshal([]byte(s), &tw)
	fmt.Println("Profile")
	for i, t := range tw.Profile {
		fmt.Printf("%02d. %20s %s\n", i+1, t.From, t.Text)
	}
}
func PostNewTimeline(text, from string) {
	s := `text=%s
username=%s
`
	payload := fmt.Sprintf(s, text, from)
	DoPost("timelines", []byte(payload))
}
