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

type cloneServerArrayRunner struct {
	id string
}

func (c *cloneServerArrayRunner) Run(cl *Client) (interface{}, error) {
	return nil, cl.CloneServerArray(c.id)
}

type createServerArrayRunner struct {
	serverArray                        *ServerArrayParam
	serverArrayJson                    string
	datacenterPolicyDatacenterHrefsPos []string
	datacenterPolicyDatacenterHrefs    []string
	datacenterPolicyMaxPos             []string
	datacenterPolicyMax                []string
	datacenterPolicyWeightPos          []string
	datacenterPolicyWeight             []string
	scheduleMaxBound                   []string
	scheduleMaxBoundPos                []string
	scheduleMinBound                   []string
	scheduleMinBoundPos                []string
	scheduleDay                        []string
	scheduleDayPos                     []string
}

func (c *createServerArrayRunner) Run(cl *Client) (interface{}, error) {
	maxDatacenterHrefsPos := 0
	var seenPos map[int]bool
	for _, p := range c.datacenterPolicyDatacenterHrefsPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position '%s'", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Datacenter Href for DatacenterPolicy #%s defined multiple times", p)
		}
		seen[pos] = true
		if pos > maxDatacenterHrefsPos {
			maxDatacenterHrefsPos = pos
		}
	}
	if len(c.datacenterPolicyDatacenterHrefs) != maxDatacenterHrefsPos {
		return nil, fmt.Errof("Invalid DatacenterPolicy DatacenterHrefs options, %s were defined but max position is %s",
			len(c.datacenterPolicyDatacenterHrefs), maxDatacenterHrefsPos)
	}
	// SAME AS ABOVE FOR MAX AND WEIGHT
	datacenterPolicy := make([]*DatacenterPolicy, maxDatacenterHrefsPos+1)
	for i := 0; i < maxDatacenterHrefsPos+1; i++ {
		datacenterPolicy[i] = &DatacenterPolicy{
			DatacenterHref: c.datacenterPolicyDatacenterHrefs[c.datacenterPolicyDatacenterHrefsPos[i]],
			Max:            c.datacenterPolicyDatacenterMax[c.datacenterPolicyDatacenterMaxPos[i]],
			Weight:         c.datacenterPolicyDatacenterWeigth[c.datacenterPolicyDatacenterWeigthPos[i]],
		}
	}
	c.serverArray.DatacenterPolicy = datacenterPolicy

	return cl.CreateServerArray(c.serverArray)
}

func registerServerArrayCmds(app *kingpin.Application) {
	resCmd := app.Command("ServerArray", "")

	cloneRunner := cloneServerArrayRunner{}
	cloneCmd := resCmd.Command("clone", "Clones a given server array")
	cloneCmd.Flag("id", "").Required().StringVar(&cloneRunner.id)
	registry[cloneServerArrayCmd.FullCommand()] = &cloneRunner

	createElasticityParamSchedules := []*Schedule{}
	createElasticityParam := ElasticityParam{Schedule: createElasticityParamSchedules}
	createServerArrayParam := ServerArrayParam{elasticityParam: &createElasticityParam}
	createRunner := createServerArrayRunner{serverArray: &createServerArrayParam}

	createCmd := resCmd.Command("create", "Creates a new server array, and configures its corresponding \"next\" instance with the received parameters")
	createCmd.Flag("serverArray", "JSON").StringVar(&createRunner.serverArrayJson)
	createCmd.Flag("serverArray.ArrayType", "").Required().StringVar(&serverArrayParam.ArrayType)
	createCmd.FlagPattern(`serverArray\.DatacenterPolicy\.(\d+)\.DatacenterHef`, "").Required().
		Capture(&createRunner.datacenterPolicyDatacenterHrefsPos).
		StringsVar(&createRunner.datacenterPolicyDatacenterHrefs)
	createCmd.FlagPattern(`serverArray\.DatacenterPolicy\.(\d+)\.Max`, "").Required().
		Capture(&createRunner.datacenterPolicyMaxPos).
		StringsVar(&createRunner.datacenterPolicyMax)
	createCmd.FlagPattern(`serverArray\.DatacenterPolicy\.(\d+)\.Weight`, "").Required().
		Capture(&createRunner.datacenterPolicyWeightPos).
		StringsVar(&createRunner.datacenterPolicyWeight)
	createCmd.Flag("serverArray.ElasticityParam.Name", "").StringVar(&elasticityParam.Name)
	createCmd.FlagPattern(`serverArray\.ElasticityParam\.Schedule\.(\d+)\.MinBound`, "").Required().
		Capture(&createRunner.scheduleMinBoundPos).
		StringsVar(&createRunner.scheduleMinBound)
	createCmd.FlagPattern(`serverArray\.ElasticityParam\.Schedule\.(\d+)\.MaxBound`, "").Required().
		Capture(&createRunner.scheduleMaxBoundPos).
		StringsVar(&createRunner.scheduleMaxBound)
	createCmd.FlagPattern(`serverArray\.ElasticityParam\.Schedule\.(\d+)\.Day`, "").Required().
		Capture(&createRunner.scheduleDayPos).
		StringsVar(&createRunner.scheduleDay)
}
