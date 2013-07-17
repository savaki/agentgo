package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os/exec"
	"regexp"
	"time"
)

type UserData struct {
	Server string `json:"go-server"`
	Port   string `json:"go-port"`
}

var (
	GO_AGENT_CONFIGFILE = "/etc/defaults/go-agent"
	TIMEOUT             = time.Duration(3 * time.Second)
	ERR_FILE            = "/GO_AGENT_FAILED"
	GO_SERVER           = "GO_SERVER="
	GO_SERVER_RE        = regexp.MustCompile("(?m)^GO_SERVER=(.+)")
	GO_SERVER_PORT      = "GO_SERVER_PORT="
	GO_SERVER_PORT_RE   = regexp.MustCompile("(?m)^GO_SERVER_PORT=(.+)")
)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, TIMEOUT)
}

func getAmazonUserData() (io.ReadCloser, error) {
	transport := http.Transport{Dial: dialTimeout}
	client := http.Client{Transport: &transport}
	request, err := http.NewRequest("GET", "http://169.254.169.254/latest/user-data", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	} else {
		return response.Body, nil
	}
}

func getUserData(dataSource func() (io.ReadCloser, error)) (*UserData, error) {
	readCloser, err := dataSource()
	if err != nil {
		return nil, err
	}
	defer readCloser.Close()

	data, err := ioutil.ReadAll(readCloser)
	if err != nil {
		return nil, err
	}

	userData := new(UserData)
	err = json.Unmarshal(data, userData)
	return userData, err
}

func writeErrorMessage(err error, filename string) {
	ioutil.WriteFile(filename, []byte(err.Error()), 0644)
}

func writeGoAgentConfig(config *UserData, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	data = GO_SERVER_RE.ReplaceAllLiteral(data, []byte(GO_SERVER+config.Server))
	data = GO_SERVER_PORT_RE.ReplaceAllLiteral(data, []byte(GO_SERVER_PORT+config.Port))

	return ioutil.WriteFile(filename, data, 0644)
}

func startGoAgent() error {
	command := exec.Command("/etc/init.d/go-agent", "start")
	return command.Wait()
}

// agent.go is a script that runs at startup that performs the following tasks:
// read the user-data url, http://169.254.169.254/latest/user-data
//
func main() {
	// check amazon's user-data service
	config, err := getUserData(getAmazonUserData)
	if err != nil {
		writeErrorMessage(err, ERR_FILE)
		return
	}

	// rewrite /etc/defaults/go-agent
	err = writeGoAgentConfig(config, GO_AGENT_CONFIGFILE)
	if err != nil {
		writeErrorMessage(err, ERR_FILE)
	}

	// start up the go-agent
	err = startGoAgent()
	if err != nil {
		writeErrorMessage(err, ERR_FILE)
	}
}
