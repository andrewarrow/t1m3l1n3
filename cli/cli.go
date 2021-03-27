package cli

import (
	"clt/persist"
	"fmt"
	"os"
	"strings"
)

var ArgMap = map[string]string{}
var Username string

func EnsureParamPass(vars ...string) {
	for _, v := range vars {
		if ArgMap[v] == "" {
			fmt.Printf("error: Missing argument. Pass --%s=%s_ID\n", v, strings.ToUpper(v))
			os.Exit(1)
		}
	}
}

func ReadInGlobalVars() {
	ArgMap = argsToMap()
	Username = persist.ReadFromFile("USERNAME")
}

func DisplayString(s string, size int) string {
	if len(s) > size {
		return s[0:size]
	}
	return s
}
func LeftAligned(thing interface{}, size int) string {
	s := fmt.Sprintf("%v", thing)

	if len(s) > size {
		return s[0:size]
	}
	fill := size - len(s)
	spaces := []string{}
	for {
		spaces = append(spaces, " ")
		if len(spaces) >= fill {
			break
		}
	}
	return s + strings.Join(spaces, "")
}

func argsToMap() map[string]string {
	m := map[string]string{}
	if len(os.Args) == 1 {
		return m
	}

	for _, a := range os.Args[1:] {
		if strings.HasPrefix(a, "--") {
			tokens := strings.Split(a, "=")
			key := strings.Split(tokens[0], "--")
			if len(tokens) == 2 {
				m[key[1]] = tokens[1]
			} else {
				m[key[1]] = "true"
			}
		} else if strings.Contains(a, "=") {
			tokens := strings.Split(a, "=")
			if len(tokens) == 2 {
				m[tokens[0]] = tokens[1]
			}
		}
	}
	return m
}
