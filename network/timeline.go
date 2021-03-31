package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/justincampbell/timeago"

	"github.com/gin-gonic/gin"
)

func ShowInbox(c *gin.Context) {
	from := c.Request.Header["Username"]
	fmt.Println(from)
	UniverseLock.Lock()
	defer UniverseLock.Unlock()
	fromIndex := universes[uids[uidIndex]].UsernameToIndex(from[0]) - 1
	if fromIndex < 255 {
		c.JSON(200, gin.H{"inbox": universes[uids[uidIndex]].Inboxes[fromIndex]})
		return
	}
	fromIndex = universes[uids[uidIndex+1]].UsernameToIndex(from[0]) - 1
	c.JSON(200, gin.H{"inbox": universes[uids[uidIndex+1]].Inboxes[fromIndex]})
}

func ToggleFollowPost(c *gin.Context) {
	from := c.Request.Header["Username"]
	to := c.Param("username")
	UniverseLock.Lock()
	bin := universes[uids[uidIndex]].ToggleFollow(from[0], to)
	c.JSON(200, gin.H{"mask": bin})
	UniverseLock.Unlock()
}

func ShowTimelines(c *gin.Context) {

	username := c.Param("username")
	UniverseLock.Lock()
	fromIndex := universes[uids[uidIndex]].UsernameToIndex(username) - 1
	c.JSON(200, gin.H{"profile": universes[uids[uidIndex]].Profile[fromIndex]})
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

type TimelineProfileWrapper struct {
	Profile []Timeline `json:"profile"`
}
type TimelineInboxWrapper struct {
	Inbox []Timeline `json:"inbox"`
}

type Timeline struct {
	Text     string `json:"text"`
	From     string `json:"from"`
	PostedAt int64  `json:"posted_at"`
	Origin   string `json:"origin"`
}

func (t *Timeline) AsTime() time.Time {
	return time.Unix(t.PostedAt, 0)
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

func VerifySig(msg, s, from string) bool {
	UniverseLock.Lock()
	defer UniverseLock.Unlock()
	pubKey := universes[uids[uidIndex]].UsernameKeys[from]
	if len(pubKey) == 0 {
		return false
	}

	return true
}

func CreateTimeline(c *gin.Context) {
	m := mapJsonBody(c)
	t := Timeline{}
	t.Text = m["text"]
	t.From = m["username"]
	s := m["s"]

	if VerifySig(t.Text, s, t.From) == false {
		c.JSON(422, gin.H{"ok": false, "sig": "failed"})
		return
	}

	t.PostedAt = time.Now().Unix()
	t.Origin = globalInOut.Name

	if t.AddToUniverse() == true {
		//TellOutAboutNewTimeline(&t, globalInOut.Out)
		c.JSON(200, gin.H{"ok": true})
		return
	}
	c.JSON(422, gin.H{"ok": false})
}

func (t *Timeline) AddToUniverse() bool {
	UniverseLock.Lock()
	defer UniverseLock.Unlock()
	if universes[uids[uidIndex]].BroadcastNewTimeline(t) {
		fmt.Println("Add User or Existing User", uidIndex)
		return true
	}
	if universes[uids[uidIndex+1]].BroadcastNewTimeline(t) {
		fmt.Println("Add User or Existing User", uidIndex+1)
		return true
	}
	return false
}

func TellOutAboutNewTimeline(t *Timeline, out string) {
	os.Setenv("CLT_HOST", fmt.Sprintf("http://%s/", out))
	asBytes, _ := json.Marshal(t)
	DoPost("timelines/notify", asBytes)
	os.Setenv("CLT_HOST", "")
}

func DisplayInboxTimelines(s string) {
	var tw TimelineInboxWrapper
	json.Unmarshal([]byte(s), &tw)
	fmt.Println("Inbox")
	for i, t := range tw.Inbox {
		fmt.Printf("%02d. %20s %20s %s\n", i+1, t.From,
			timeago.FromDuration(time.Since(t.AsTime())), t.Text)
		if i > 20 {
			break
		}
	}
}
func DisplayProfileTimelines(s string) {
	var tw TimelineProfileWrapper
	json.Unmarshal([]byte(s), &tw)
	fmt.Println("Profile")
	for i, t := range tw.Profile {
		fmt.Printf("%02d. %20s %20s %s\n", i+1, t.From,
			timeago.FromDuration(time.Since(t.AsTime())), t.Text)
	}
}
func PostNewTimeline(text, from, s string) {
	m := map[string]string{"text": text, "username": from, "s": s}
	asBytes, _ := json.Marshal(m)
	DoPost("timelines", asBytes)
}
