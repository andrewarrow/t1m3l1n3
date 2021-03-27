package persist

import (
	"io/ioutil"
	"os"
	"runtime"
)

const DIRNAME = ".command-line-timelines"

func Init() {
	home := UserHomeDir()
	os.Mkdir(home+"/"+DIRNAME, 0755)
}

func RemoveList(files []string) {
	home := UserHomeDir()
	for _, f := range files {
		path := home + "/" + DIRNAME + "/" + f
		os.Remove(path)
	}
}
func SaveToFile(name, value string) {
	home := UserHomeDir()
	path := home + "/" + DIRNAME + "/" + name
	os.Remove(path)
	f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	defer f.Close()
	f.WriteString(value)
}

func ReadFromFile(name string) string {
	home := UserHomeDir()
	path := home + "/" + DIRNAME + "/" + name
	b, _ := ioutil.ReadFile(path)
	return string(b)
}
func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
