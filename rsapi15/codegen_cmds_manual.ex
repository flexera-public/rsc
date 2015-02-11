package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/raphael/kingpin"
)

type CommandRunner interface {
	Run(cl *Client) (interface{}, error)
}

var registry map[string]CommandRunner = make(map[string]CommandRunner)

func main() {
	app := kingpin.New("goo", "gar")
	registerInstanceCmds(app)
	cmdName, err := app.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}
	if cmdName == "" {
		fmt.Println("Missing coommand")
		return
	}
	registry[cmdName].Run()
}

type createServerArrayRunner struct {
	Name                         string
	ServerArrayJson              string
	ServerArrayArrayType         string
	ServerArrayInstanceCloudHref string
	ServerArrayInstanceSshKey    string
	ServerArraySchedulesNames    []string
	ServerArraySchedulesNamesPos []string
	ServerArrayInputsValues      []string
	ServerArrayInputsNames       []string
}

func (r *createServerArrayRunner) Run(cl *Client) (interface{}, error) {

	/** Handle serverArray parameter **/

	var serverArray ServerArrayParam

	// Load JSON if provided
	if len(r.ServerArrayJson) > 0 {
		if err := Json.Unmarshal(r.ServerArrayJson, &serverArray); err != nil {
			return nil, fmt.Errorf("Could not load ServerArray json: %s", err.Error())
		}
	}

	// Override with any provided basic fields
	if len(r.ServerArrayArrayType) > 0 {
		serverArray.ArrayType = r.ServerArrayArrayType
	}
	if len(r.ServerArrayInstanceCloudHref) > 0 {
		serverArray.Instance.CloudHref = r.ServerArrayInstanceCloudHref
	}
	if len(r.ServerArrayInstanceSshKey) > 0 {
		serverArray.Instance.SshKey = r.ServerArrayInstanceSshKey
	}

	// Create array fields from value and position flags
	maxServerArraySchedulesNamesPos := 0
	var seenPos map[int]bool
	for _, p := range r.ServerArraySchedulesNamesPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position '%s'", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("SchedulesNames defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerArraySchedulesNamesPos {
			maxServerArraySchedulesNamesPos = pos
		}
	}
	if len(r.ServerArraySchedulesNames) != maxServerArraySchedulesNamesPos {
		return nil, fmt.Errof("Invalid SchedulesNames options, %s were defined but max position is %s",
			len(r.ServerArraySchedulesNames), maxServerArraySchedulesNamesPos)
	}
	serverArraySchedules := make([]*ServerArraySchedules, maxServerArraySchedulesNamesPos+1)
	for i := 0; i < maxServerArraySchedulesNamesPos+1; i++ {
		serverArraySchedules[i] = &Schedule{
			Name: r.ServerArraySchedulesNames[r.ServerArraySchedulesNamesPos[i]],
		}
	}
	if maxServerArraySchedulesNamesPos > 0 {
		serverArray.Schedules = serverArraySchedules
	}

	// Create Enum fields from name and value flags
	serverArrayInputs := make(map[string]string, len(r.ServerArrayInputsNames))
	for i, n := range ServerArrayInputsNames {
		serverArrayInputs[n] = r.ServerArrayInputsValues[i]
	}
	serverArray.Inputs = serverArrayInputs

	// All done, call API
	return cl.CreateServerArray(r.Name, &serverArray)
}

func registerServerArrayCmds(app *kingpin.Application) {
	resCmd := app.Command("ServerArray", "")

	createRunner := new(createServerArrayRunner)

	createCmd := resCmd.Command("create", "Creates a new server array, and configures its corresponding \"next\" instance with the received parameters")
	createCmd.Flag("name", "Name").StringVar(&createRunner.Name)
	createCmd.Flag("serverArray", "JSON").StringVar(&createRunner.ServerArrayJson)
	createCmd.Flag("serverArray.ArrayType", "").Required().StringVar(&createRunner.ServerArrayArrayType)
	createCmd.Flag("serverArray.Instance.CloudHref", "").Required().StringVar(&createRunner.ServerArrayInstanceCloudHref)
	createCmd.Flag("serverArray.Instance.SshKey", "").StringVar(&createRunner.ServerArrayInstanceSshKey)
	createCmd.FlagPattern(`serverArray\.Schedules\.(\d+)\.Name`, "").Required().
		Capture(&createRunner.ServerArraySchedulesNamesPos).StringsVar(&createRunner.ServerArraySchedulesNames)
	createCmd.FlagPattern(`serverArray\.Inputs\.([a-z0-9_]+)`, "").
		Capture(&createRunner.ServerArrayInputsNames).StringsVar(&createRunner.ServerArrayInputsValues)

	registry[createCmd.FullCommand()] = createRunner
}
