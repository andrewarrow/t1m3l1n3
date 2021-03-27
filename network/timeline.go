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

func ShowTimelines(c *gin.Context) {

	ByFromLock.Lock()
	c.JSON(200, gin.H{"from": ByFrom})
	ByFromLock.Unlock()
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
	From map[string][]Timeline `json:"from"`
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

func (t *Timeline) ToKey() string {
	return fmt.Sprintf("%s_%d", t.From, t.PostedAt)
}

var ByFromLock sync.Mutex
var ByFrom map[string][]*Timeline = map[string][]*Timeline{}
var ByKey map[string]*Timeline = map[string]*Timeline{}

func NotifyTimeline(c *gin.Context) {
	m := mapJsonBody(c)
	fmt.Println("mapJsonBody", m)
	t := TimelineFromMap(m)
	t.AddToByKey()
}

func CreateTimeline(c *gin.Context) {
	m := mapBody(c)
	t := Timeline{}
	t.Text = m["text"]
	t.From = m["username"]
	t.PostedAt = time.Now().Unix()
	t.Origin = globalInOut.Name

	if t.AddToByKey() == true {
		TellOutAboutNewTimeline(&t, globalInOut.Out)
	}
}

func (t *Timeline) AddToByKey() bool {
	key := t.ToKey()
	ByFromLock.Lock()
	defer ByFromLock.Unlock()

	if ByKey[key] != nil && ByKey[key].Origin == globalInOut.Name {
		return false
	}

	ByKey[key] = t
	ByFrom = map[string][]*Timeline{}
	for k, v := range ByKey {
		tokens := strings.Split(k, "_")
		from := tokens[0]
		//ts := tokens[1]
		ByFrom[from] = append([]*Timeline{v}, ByFrom[from]...)
	}
	return true
}

func TellOutAboutNewTimeline(t *Timeline, out string) {
	os.Setenv("CLT_HOST", fmt.Sprintf("http://%s/", out))
	asBytes, _ := json.Marshal(t)
	DoPost("timelines/notify", asBytes)
}

func DisplayTimelines(s string) {
	var tw TimelineWrapper
	json.Unmarshal([]byte(s), &tw)
	for k, v := range tw.From {
		fmt.Println(k)
		for i, t := range v {
			fmt.Printf("%02d. %s\n", i+1, t.Text)
		}
	}
}
