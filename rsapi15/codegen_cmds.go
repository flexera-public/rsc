//************************************************************************//
//                RightScale API 1.5 command line client
//
// Generated Feb 11, 2015 at 1:26pm (PST)
// Command:
// $ api15gen -metadata=../../rsapi15 -output=../../rsapi15
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package rsapi15

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v1"
)

// Registry of all sub-commands
var registry map[string]CommandRunner = make(map[string]CommandRunner)

// Register all sub-commands
func RegisterCommands(app *kingpin.Application) {
	registerAccountCommands(app)
	registerAccountGroupCommands(app)
	registerAlertCommands(app)
	registerAlertSpecCommands(app)
	registerAuditEntryCommands(app)
	registerBackupCommands(app)
	registerChildAccountCommands(app)
	registerCloudCommands(app)
	registerCloudAccountCommands(app)
	registerCookbookCommands(app)
	registerCookbookAttachmentCommands(app)
	registerCredentialCommands(app)
	registerDatacenterCommands(app)
	registerDeploymentCommands(app)
	registerHealthCheckCommands(app)
	registerIdentityProviderCommands(app)
	registerImageCommands(app)
	registerInputCommands(app)
	registerInstanceCommands(app)
	registerInstanceCustomLodgementCommands(app)
	registerInstanceTypeCommands(app)
	registerIpAddressCommands(app)
	registerIpAddressBindingCommands(app)
	registerMonitoringMetricCommands(app)
	registerMultiCloudImageCommands(app)
	registerMultiCloudImageSettingCommands(app)
	registerNetworkCommands(app)
	registerNetworkGatewayCommands(app)
	registerNetworkOptionGroupCommands(app)
	registerNetworkOptionGroupAttachmentCommands(app)
	registerOauth2Commands(app)
	registerPermissionCommands(app)
	registerPlacementGroupCommands(app)
	registerPreferenceCommands(app)
	registerPublicationCommands(app)
	registerPublicationLineageCommands(app)
	registerRecurringVolumeAttachmentCommands(app)
	registerRepositoryCommands(app)
	registerRepositoryAssetCommands(app)
	registerRightScriptCommands(app)
	registerRouteCommands(app)
	registerRouteTableCommands(app)
	registerRunnableBindingCommands(app)
	registerSecurityGroupCommands(app)
	registerSecurityGroupRuleCommands(app)
	registerServerCommands(app)
	registerServerArrayCommands(app)
	registerServerTemplateCommands(app)
	registerServerTemplateMultiCloudImageCommands(app)
	registerSessionCommands(app)
	registerSshKeyCommands(app)
	registerSubnetCommands(app)
	registerTagCommands(app)
	registerTaskCommands(app)
	registerUserCommands(app)
	registerUserDataCommands(app)
	registerVolumeCommands(app)
	registerVolumeAttachmentCommands(app)
	registerVolumeSnapshotCommands(app)
	registerVolumeTypeCommands(app)
}

/****** Account ******/

type ShowAccountAccountRunner struct {
	id string
}

func (r *ShowAccountAccountRunner) Run(c *Client) (interface{}, error) {
	return c.ShowAccount(r.id)
}

// Register all Account actions
func registerAccountCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Account", ``)

	ShowAccountRunner := new(ShowAccountAccountRunner)
	ShowAccountCmd := resCmd.Command("ShowAccount", `Show information about a single Account.`)
	ShowAccountRunner.Flag(`id`, ``).Required().StringVar(&ShowAccountRunner.id)
	registry[ShowAccountCmd.FullCommand()] = ShowAccountRunner
}

/****** AccountGroup ******/

type IndexAccountGroupsAccountGroupRunner struct {
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexAccountGroupsAccountGroupRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexAccountGroups(filter, r.view)
}

type ShowAccountGroupAccountGroupRunner struct {
	id   string
	view string
}

func (r *ShowAccountGroupAccountGroupRunner) Run(c *Client) (interface{}, error) {
	return c.ShowAccountGroup(r.id, r.view)
}

// Register all AccountGroup actions
func registerAccountGroupCmds(app *kingpin.Application) {
	resCmd := app.Cmd("AccountGroup", `An Account Group specifies which RightScale accounts will have access to import a shared RightScale component (e.g. ServerTemplate, RightScript, etc.) from the MultiCloud Marketplace.`)

	IndexAccountGroupsRunner := new(IndexAccountGroupsAccountGroupRunner)
	IndexAccountGroupsCmd := resCmd.Command("IndexAccountGroups", `Lists the AccountGroups owned by this Account.`)
	IndexAccountGroupsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexAccountGroupsRunner.filterPos).StringsVar(&IndexAccountGroupsRunner.filter)
	IndexAccountGroupsRunner.Flag(`view`, ``).StringVar(&IndexAccountGroupsRunner.view)
	registry[IndexAccountGroupsCmd.FullCommand()] = IndexAccountGroupsRunner

	ShowAccountGroupRunner := new(ShowAccountGroupAccountGroupRunner)
	ShowAccountGroupCmd := resCmd.Command("ShowAccountGroup", `Show information about a single AccountGroup.`)
	ShowAccountGroupRunner.Flag(`id`, ``).Required().StringVar(&ShowAccountGroupRunner.id)
	ShowAccountGroupRunner.Flag(`view`, ``).StringVar(&ShowAccountGroupRunner.view)
	registry[ShowAccountGroupCmd.FullCommand()] = ShowAccountGroupRunner
}

/****** Alert ******/

type DisableAlertAlertRunner struct {
	cloudId    string
	id         string
	instanceId string
}

func (r *DisableAlertAlertRunner) Run(c *Client) (interface{}, error) {
	return c.DisableAlert(r.cloudId, r.id, r.instanceId)
}

type EnableAlertAlertRunner struct {
	cloudId    string
	id         string
	instanceId string
}

func (r *EnableAlertAlertRunner) Run(c *Client) (interface{}, error) {
	return c.EnableAlert(r.cloudId, r.id, r.instanceId)
}

type IndexAlertsAlertRunner struct {
	cloudId    string
	filter     []string
	filterPos  []string
	instanceId string
	view       string
}

func (r *IndexAlertsAlertRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexAlerts(r.cloudId, filter, r.instanceId, r.view)
}

type QuenchAlertAlertRunner struct {
	cloudId    string
	duration   string
	id         string
	instanceId string
}

func (r *QuenchAlertAlertRunner) Run(c *Client) (interface{}, error) {
	return c.QuenchAlert(r.cloudId, r.duration, r.id, r.instanceId)
}

type ShowAlertAlertRunner struct {
	cloudId    string
	id         string
	instanceId string
	view       string
}

func (r *ShowAlertAlertRunner) Run(c *Client) (interface{}, error) {
	return c.ShowAlert(r.cloudId, r.id, r.instanceId, r.view)
}

// Register all Alert actions
func registerAlertCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Alert", `An Alert represents an AlertSpec bound to a running Instance.`)

	DisableAlertRunner := new(DisableAlertAlertRunner)
	DisableAlertCmd := resCmd.Command("DisableAlert", `Disables the Alert indefinitely. Idempotent.`)
	DisableAlertRunner.Flag(`cloudId`, ``).Required().StringVar(&DisableAlertRunner.cloudId)
	DisableAlertRunner.Flag(`id`, ``).Required().StringVar(&DisableAlertRunner.id)
	DisableAlertRunner.Flag(`instanceId`, ``).Required().StringVar(&DisableAlertRunner.instanceId)
	registry[DisableAlertCmd.FullCommand()] = DisableAlertRunner

	EnableAlertRunner := new(EnableAlertAlertRunner)
	EnableAlertCmd := resCmd.Command("EnableAlert", `Enables the Alert indefinitely. Idempotent.`)
	EnableAlertRunner.Flag(`cloudId`, ``).Required().StringVar(&EnableAlertRunner.cloudId)
	EnableAlertRunner.Flag(`id`, ``).Required().StringVar(&EnableAlertRunner.id)
	EnableAlertRunner.Flag(`instanceId`, ``).Required().StringVar(&EnableAlertRunner.instanceId)
	registry[EnableAlertCmd.FullCommand()] = EnableAlertRunner

	IndexAlertsRunner := new(IndexAlertsAlertRunner)
	IndexAlertsCmd := resCmd.Command("IndexAlerts", `Lists all Alerts.`)
	IndexAlertsRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexAlertsRunner.cloudId)
	IndexAlertsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexAlertsRunner.filterPos).StringsVar(&IndexAlertsRunner.filter)
	IndexAlertsRunner.Flag(`instanceId`, ``).Required().StringVar(&IndexAlertsRunner.instanceId)
	IndexAlertsRunner.Flag(`view`, ``).StringVar(&IndexAlertsRunner.view)
	registry[IndexAlertsCmd.FullCommand()] = IndexAlertsRunner

	QuenchAlertRunner := new(QuenchAlertAlertRunner)
	QuenchAlertCmd := resCmd.Command("QuenchAlert", `Suppresses the Alert from being triggered for a given time period. Idempotent.`)
	QuenchAlertRunner.Flag(`cloudId`, ``).Required().StringVar(&QuenchAlertRunner.cloudId)
	QuenchAlertRunner.Flag(`duration`, `The time period in seconds to suppress Alert from being triggered.`).Required().StringVar(&QuenchAlertRunner.duration)
	QuenchAlertRunner.Flag(`id`, ``).Required().StringVar(&QuenchAlertRunner.id)
	QuenchAlertRunner.Flag(`instanceId`, ``).Required().StringVar(&QuenchAlertRunner.instanceId)
	registry[QuenchAlertCmd.FullCommand()] = QuenchAlertRunner

	ShowAlertRunner := new(ShowAlertAlertRunner)
	ShowAlertCmd := resCmd.Command("ShowAlert", `Shows the attributes of a specified Alert.`)
	ShowAlertRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowAlertRunner.cloudId)
	ShowAlertRunner.Flag(`id`, ``).Required().StringVar(&ShowAlertRunner.id)
	ShowAlertRunner.Flag(`instanceId`, ``).Required().StringVar(&ShowAlertRunner.instanceId)
	ShowAlertRunner.Flag(`view`, ``).StringVar(&ShowAlertRunner.view)
	registry[ShowAlertCmd.FullCommand()] = ShowAlertRunner
}

/****** AlertSpec ******/

type CreateAlertSpecAlertSpecRunner struct {
	alertSpecCondition      string
	alertSpecDescription    string
	alertSpecDuration       string
	alertSpecEscalationName string
	alertSpecFile           string
	alertSpecName           string
	alertSpecSubjectHref    string
	alertSpecThreshold      string
	alertSpecVariable       string
	alertSpecVoteTag        string
	alertSpecVoteType       string
	serverId                string
}

func (r *CreateAlertSpecAlertSpecRunner) Run(c *Client) (interface{}, error) {

	/** Handle alertSpec parameter **/
	var alertSpec AlertSpecParam

	// Load JSON if provided
	if len(r.alertSpecJson) > 0 {
		if err := Json.Unmarshal(r.alertSpecJson, &alertSpec); err != nil {
			return nil, fmt.Errorf("Could not load alertSpec JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.alertSpeccondition) > 0 {
		alertSpec.alertSpec.condition = r.alertSpeccondition
	}

	if len(r.alertSpecdescription) > 0 {
		alertSpec.alertSpec.description = r.alertSpecdescription
	}

	if len(r.alertSpecduration) > 0 {
		alertSpec.alertSpec.duration = r.alertSpecduration
	}

	if len(r.alertSpecescalationName) > 0 {
		alertSpec.alertSpec.escalationName = r.alertSpecescalationName
	}

	if len(r.alertSpecfile) > 0 {
		alertSpec.alertSpec.file = r.alertSpecfile
	}

	if len(r.alertSpecname) > 0 {
		alertSpec.alertSpec.name = r.alertSpecname
	}

	if len(r.alertSpecsubjectHref) > 0 {
		alertSpec.alertSpec.subjectHref = r.alertSpecsubjectHref
	}

	if len(r.alertSpecthreshold) > 0 {
		alertSpec.alertSpec.threshold = r.alertSpecthreshold
	}

	if len(r.alertSpecvariable) > 0 {
		alertSpec.alertSpec.variable = r.alertSpecvariable
	}

	if len(r.alertSpecvoteTag) > 0 {
		alertSpec.alertSpec.voteTag = r.alertSpecvoteTag
	}

	if len(r.alertSpecvoteType) > 0 {
		alertSpec.alertSpec.voteType = r.alertSpecvoteType
	}

	return c.CreateAlertSpec(&alertSpec, r.serverId)
}

type DestroyAlertSpecAlertSpecRunner struct {
	id       string
	serverId string
}

func (r *DestroyAlertSpecAlertSpecRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyAlertSpec(r.id, r.serverId)
}

type IndexAlertSpecsAlertSpecRunner struct {
	filter        []string
	filterPos     []string
	serverId      string
	view          string
	withInherited string
}

func (r *IndexAlertSpecsAlertSpecRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexAlertSpecs(filter, r.serverId, r.view, r.withInherited)
}

type ShowAlertSpecAlertSpecRunner struct {
	id       string
	serverId string
	view     string
}

func (r *ShowAlertSpecAlertSpecRunner) Run(c *Client) (interface{}, error) {
	return c.ShowAlertSpec(r.id, r.serverId, r.view)
}

type UpdateAlertSpecAlertSpecRunner struct {
	alertSpecCondition      string
	alertSpecDescription    string
	alertSpecDuration       string
	alertSpecEscalationName string
	alertSpecFile           string
	alertSpecName           string
	alertSpecThreshold      string
	alertSpecVariable       string
	alertSpecVoteTag        string
	alertSpecVoteType       string
	id                      string
	serverId                string
}

func (r *UpdateAlertSpecAlertSpecRunner) Run(c *Client) (interface{}, error) {

	/** Handle alertSpec parameter **/
	var alertSpec AlertSpecParam2

	// Load JSON if provided
	if len(r.alertSpecJson) > 0 {
		if err := Json.Unmarshal(r.alertSpecJson, &alertSpec); err != nil {
			return nil, fmt.Errorf("Could not load alertSpec JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.alertSpeccondition) > 0 {
		alertSpec.alertSpec.condition = r.alertSpeccondition
	}

	if len(r.alertSpecdescription) > 0 {
		alertSpec.alertSpec.description = r.alertSpecdescription
	}

	if len(r.alertSpecduration) > 0 {
		alertSpec.alertSpec.duration = r.alertSpecduration
	}

	if len(r.alertSpecescalationName) > 0 {
		alertSpec.alertSpec.escalationName = r.alertSpecescalationName
	}

	if len(r.alertSpecfile) > 0 {
		alertSpec.alertSpec.file = r.alertSpecfile
	}

	if len(r.alertSpecname) > 0 {
		alertSpec.alertSpec.name = r.alertSpecname
	}

	if len(r.alertSpecthreshold) > 0 {
		alertSpec.alertSpec.threshold = r.alertSpecthreshold
	}

	if len(r.alertSpecvariable) > 0 {
		alertSpec.alertSpec.variable = r.alertSpecvariable
	}

	if len(r.alertSpecvoteTag) > 0 {
		alertSpec.alertSpec.voteTag = r.alertSpecvoteTag
	}

	if len(r.alertSpecvoteType) > 0 {
		alertSpec.alertSpec.voteType = r.alertSpecvoteType
	}

	return c.UpdateAlertSpec(&alertSpec, r.id, r.serverId)
}

// Register all AlertSpec actions
func registerAlertSpecCmds(app *kingpin.Application) {
	resCmd := app.Cmd("AlertSpec", `An AlertSpec defines the conditions under which an Alert is triggered and escalated.`)

	CreateAlertSpecRunner := new(CreateAlertSpecAlertSpecRunner)
	CreateAlertSpecCmd := resCmd.Command("CreateAlertSpec", `Creates a new AlertSpec with the given parameters.`)
	CreateAlertSpecRunner.Flag(`alertSpec.condition`, `The condition (operator) in the condition sentence.`).Required().StringVar(&CreateAlertSpecRunner.alertSpecCondition)
	CreateAlertSpecRunner.Flag(`alertSpec.description`, `The description of the AlertSpec.`).StringVar(&CreateAlertSpecRunner.alertSpecDescription)
	CreateAlertSpecRunner.Flag(`alertSpec.duration`, `The duration in minutes of the condition sentence.`).Required().StringVar(&CreateAlertSpecRunner.alertSpecDuration)
	CreateAlertSpecRunner.Flag(`alertSpec.escalationName`, `Escalate to the named alert escalation when the alert is triggered. Must either escalate or vote.`).StringVar(&CreateAlertSpecRunner.alertSpecEscalationName)
	CreateAlertSpecRunner.Flag(`alertSpec.file`, `The RRD path/file_name of the condition sentence.`).Required().StringVar(&CreateAlertSpecRunner.alertSpecFile)
	CreateAlertSpecRunner.Flag(`alertSpec.name`, `The name of the AlertSpec.`).Required().StringVar(&CreateAlertSpecRunner.alertSpecName)
	CreateAlertSpecRunner.Flag(`alertSpec.subjectHref`, `The href of the resource that this AlertSpec should be associated with. The subject can be a ServerTemplate, Server, ServerArray, or Instance.`).StringVar(&CreateAlertSpecRunner.alertSpecSubjectHref)
	CreateAlertSpecRunner.Flag(`alertSpec.threshold`, `The threshold of the condition sentence.`).Required().StringVar(&CreateAlertSpecRunner.alertSpecThreshold)
	CreateAlertSpecRunner.Flag(`alertSpec.variable`, `The RRD variable of the condition sentence.`).Required().StringVar(&CreateAlertSpecRunner.alertSpecVariable)
	CreateAlertSpecRunner.Flag(`alertSpec.voteTag`, `Should correspond to a vote tag on a ServerArray if vote to grow or shrink.`).StringVar(&CreateAlertSpecRunner.alertSpecVoteTag)
	CreateAlertSpecRunner.Flag(`alertSpec.voteType`, `Vote to grow or shrink a ServerArray when the alert is triggered. Must either escalate or vote.`).StringVar(&CreateAlertSpecRunner.alertSpecVoteType)
	CreateAlertSpecRunner.Flag(`serverId`, ``).Required().StringVar(&CreateAlertSpecRunner.serverId)
	registry[CreateAlertSpecCmd.FullCommand()] = CreateAlertSpecRunner

	DestroyAlertSpecRunner := new(DestroyAlertSpecAlertSpecRunner)
	DestroyAlertSpecCmd := resCmd.Command("DestroyAlertSpec", `Deletes a given AlertSpec.`)
	DestroyAlertSpecRunner.Flag(`id`, ``).Required().StringVar(&DestroyAlertSpecRunner.id)
	DestroyAlertSpecRunner.Flag(`serverId`, ``).Required().StringVar(&DestroyAlertSpecRunner.serverId)
	registry[DestroyAlertSpecCmd.FullCommand()] = DestroyAlertSpecRunner

	IndexAlertSpecsRunner := new(IndexAlertSpecsAlertSpecRunner)
	IndexAlertSpecsCmd := resCmd.Command("IndexAlertSpecs", `<no description> -- Optional parameters: 	filter 	view 	withInherited: Flag indicating whether or not to include AlertSpecs from the ServerTemplate in the index...`)
	IndexAlertSpecsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexAlertSpecsRunner.filterPos).StringsVar(&IndexAlertSpecsRunner.filter)
	IndexAlertSpecsRunner.Flag(`serverId`, ``).Required().StringVar(&IndexAlertSpecsRunner.serverId)
	IndexAlertSpecsRunner.Flag(`view`, ``).StringVar(&IndexAlertSpecsRunner.view)
	IndexAlertSpecsRunner.Flag(`withInherited`, `Flag indicating whether or not to include AlertSpecs from the ServerTemplate in the index.`).StringVar(&IndexAlertSpecsRunner.withInherited)
	registry[IndexAlertSpecsCmd.FullCommand()] = IndexAlertSpecsRunner

	ShowAlertSpecRunner := new(ShowAlertSpecAlertSpecRunner)
	ShowAlertSpecCmd := resCmd.Command("ShowAlertSpec", `<no description> -- Optional parameters: 	view...`)
	ShowAlertSpecRunner.Flag(`id`, ``).Required().StringVar(&ShowAlertSpecRunner.id)
	ShowAlertSpecRunner.Flag(`serverId`, ``).Required().StringVar(&ShowAlertSpecRunner.serverId)
	ShowAlertSpecRunner.Flag(`view`, ``).StringVar(&ShowAlertSpecRunner.view)
	registry[ShowAlertSpecCmd.FullCommand()] = ShowAlertSpecRunner

	UpdateAlertSpecRunner := new(UpdateAlertSpecAlertSpecRunner)
	UpdateAlertSpecCmd := resCmd.Command("UpdateAlertSpec", `Updates an AlertSpec with the given parameters.`)
	UpdateAlertSpecRunner.Flag(`alertSpec.condition`, `The condition (operator) in the condition sentence.`).StringVar(&UpdateAlertSpecRunner.alertSpecCondition)
	UpdateAlertSpecRunner.Flag(`alertSpec.description`, `The description of the AlertSpec.`).StringVar(&UpdateAlertSpecRunner.alertSpecDescription)
	UpdateAlertSpecRunner.Flag(`alertSpec.duration`, `The duration in minutes of the condition sentence.`).StringVar(&UpdateAlertSpecRunner.alertSpecDuration)
	UpdateAlertSpecRunner.Flag(`alertSpec.escalationName`, `Escalate to the named alert escalation when the alert is triggered.`).StringVar(&UpdateAlertSpecRunner.alertSpecEscalationName)
	UpdateAlertSpecRunner.Flag(`alertSpec.file`, `The RRD path/file_name of the condition sentence.`).StringVar(&UpdateAlertSpecRunner.alertSpecFile)
	UpdateAlertSpecRunner.Flag(`alertSpec.name`, `The name of the AlertSpec.`).StringVar(&UpdateAlertSpecRunner.alertSpecName)
	UpdateAlertSpecRunner.Flag(`alertSpec.threshold`, `The threshold of the condition sentence.`).StringVar(&UpdateAlertSpecRunner.alertSpecThreshold)
	UpdateAlertSpecRunner.Flag(`alertSpec.variable`, `The RRD variable of the condition sentence.`).StringVar(&UpdateAlertSpecRunner.alertSpecVariable)
	UpdateAlertSpecRunner.Flag(`alertSpec.voteTag`, `Should correspond to a vote tag on a ServerArray if vote to grow or shrink.`).StringVar(&UpdateAlertSpecRunner.alertSpecVoteTag)
	UpdateAlertSpecRunner.Flag(`alertSpec.voteType`, `Vote to grow or shrink a ServerArray when the alert is triggered.`).StringVar(&UpdateAlertSpecRunner.alertSpecVoteType)
	UpdateAlertSpecRunner.Flag(`id`, ``).Required().StringVar(&UpdateAlertSpecRunner.id)
	UpdateAlertSpecRunner.Flag(`serverId`, ``).Required().StringVar(&UpdateAlertSpecRunner.serverId)
	registry[UpdateAlertSpecCmd.FullCommand()] = UpdateAlertSpecRunner
}

/****** AuditEntry ******/

type AppendAuditEntryAuditEntryRunner struct {
	detail  string
	id      string
	notify  string
	offset  int
	summary string
}

func (r *AppendAuditEntryAuditEntryRunner) Run(c *Client) (interface{}, error) {
	return c.AppendAuditEntry(r.detail, r.id, r.notify, r.offset, r.summary)
}

type CreateAuditEntryAuditEntryRunner struct {
	auditEntryAuditeeHref string
	auditEntryDetail      string
	auditEntrySummary     string
	notify                string
	userEmail             string
}

func (r *CreateAuditEntryAuditEntryRunner) Run(c *Client) (interface{}, error) {

	/** Handle auditEntry parameter **/
	var auditEntry AuditEntryParam

	// Load JSON if provided
	if len(r.auditEntryJson) > 0 {
		if err := Json.Unmarshal(r.auditEntryJson, &auditEntry); err != nil {
			return nil, fmt.Errorf("Could not load auditEntry JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.auditEntryauditeeHref) > 0 {
		auditEntry.auditEntry.auditeeHref = r.auditEntryauditeeHref
	}

	if len(r.auditEntrydetail) > 0 {
		auditEntry.auditEntry.detail = r.auditEntrydetail
	}

	if len(r.auditEntrysummary) > 0 {
		auditEntry.auditEntry.summary = r.auditEntrysummary
	}

	return c.CreateAuditEntry(&auditEntry, r.notify, r.userEmail)
}

type DetailAuditEntryAuditEntryRunner struct {
	id string
}

func (r *DetailAuditEntryAuditEntryRunner) Run(c *Client) (interface{}, error) {
	return c.DetailAuditEntry(r.id)
}

type IndexAuditEntriesAuditEntryRunner struct {
	endDate   string
	filter    []string
	filterPos []string
	limit     string
	startDate string
	view      string
}

func (r *IndexAuditEntriesAuditEntryRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexAuditEntries(r.endDate, filter, r.limit, r.startDate, r.view)
}

type ShowAuditEntryAuditEntryRunner struct {
	id   string
	view string
}

func (r *ShowAuditEntryAuditEntryRunner) Run(c *Client) (interface{}, error) {
	return c.ShowAuditEntry(r.id, r.view)
}

type UpdateAuditEntryAuditEntryRunner struct {
	auditEntryOffset  int
	auditEntrySummary string
	id                string
	notify            string
}

func (r *UpdateAuditEntryAuditEntryRunner) Run(c *Client) (interface{}, error) {

	/** Handle auditEntry parameter **/
	var auditEntry AuditEntryParam2

	// Load JSON if provided
	if len(r.auditEntryJson) > 0 {
		if err := Json.Unmarshal(r.auditEntryJson, &auditEntry); err != nil {
			return nil, fmt.Errorf("Could not load auditEntry JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.auditEntryoffset) > 0 {
		auditEntry.auditEntry.offset = r.auditEntryoffset
	}

	if len(r.auditEntrysummary) > 0 {
		auditEntry.auditEntry.summary = r.auditEntrysummary
	}

	return c.UpdateAuditEntry(&auditEntry, r.id, r.notify)
}

// Register all AuditEntry actions
func registerAuditEntryCmds(app *kingpin.Application) {
	resCmd := app.Cmd("AuditEntry", `An Audit Entry can be used to track various activities of a resource.`)

	AppendAuditEntryRunner := new(AppendAuditEntryAuditEntryRunner)
	AppendAuditEntryCmd := resCmd.Command("AppendAuditEntry", `Updates the summary and appends more details to a given AuditEntry`)
	AppendAuditEntryRunner.Flag(`detail`, `The details to be appended to the audit entry record.`).StringVar(&AppendAuditEntryRunner.detail)
	AppendAuditEntryRunner.Flag(`id`, ``).Required().StringVar(&AppendAuditEntryRunner.id)
	AppendAuditEntryRunner.Flag(`notify`, `The event notification category. Defaults to 'None'.`).StringVar(&AppendAuditEntryRunner.notify)
	AppendAuditEntryRunner.Flag(`offset`, `The offset where the new details should be appended to in the audit entry's existing details section. Also used in ordering of summary updates. Defaults to end.`).IntVar(&AppendAuditEntryRunner.offset)
	AppendAuditEntryRunner.Flag(`summary`, `The updated summary for the audit entry, maximum length is 255 characters.`).StringVar(&AppendAuditEntryRunner.summary)
	registry[AppendAuditEntryCmd.FullCommand()] = AppendAuditEntryRunner

	CreateAuditEntryRunner := new(CreateAuditEntryAuditEntryRunner)
	CreateAuditEntryCmd := resCmd.Command("CreateAuditEntry", `Creates a new AuditEntry with the given parameters.`)
	CreateAuditEntryRunner.Flag(`auditEntry.auditeeHref`, `The href of the resource that this audit entry should be associated with (e.g. an instance's href).`).Required().StringVar(&CreateAuditEntryRunner.auditEntryAuditeeHref)
	CreateAuditEntryRunner.Flag(`auditEntry.detail`, `The initial details of the audit entry to be created.`).StringVar(&CreateAuditEntryRunner.auditEntryDetail)
	CreateAuditEntryRunner.Flag(`auditEntry.summary`, `The summary of the audit entry to be created, maximum length is 255 characters.`).Required().StringVar(&CreateAuditEntryRunner.auditEntrySummary)
	CreateAuditEntryRunner.Flag(`notify`, `The event notification category. Defaults to 'None'.`).StringVar(&CreateAuditEntryRunner.notify)
	CreateAuditEntryRunner.Flag(`userEmail`, `The email of the user (who created/triggered the audit entry). Only usable with instance role.`).StringVar(&CreateAuditEntryRunner.userEmail)
	registry[CreateAuditEntryCmd.FullCommand()] = CreateAuditEntryRunner

	DetailAuditEntryRunner := new(DetailAuditEntryAuditEntryRunner)
	DetailAuditEntryCmd := resCmd.Command("DetailAuditEntry", `shows the details of a given AuditEntry.`)
	DetailAuditEntryRunner.Flag(`id`, ``).Required().StringVar(&DetailAuditEntryRunner.id)
	registry[DetailAuditEntryCmd.FullCommand()] = DetailAuditEntryRunner

	IndexAuditEntriesRunner := new(IndexAuditEntriesAuditEntryRunner)
	IndexAuditEntriesCmd := resCmd.Command("IndexAuditEntries", `Lists AuditEntries of the account`)
	IndexAuditEntriesRunner.Flag(`endDate`, `The end date for retrieving audit entries (the format must be the same as start date). The time period between start and end date must be less than 3 months (93 days).`).Required().StringVar(&IndexAuditEntriesRunner.endDate)
	IndexAuditEntriesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexAuditEntriesRunner.filterPos).StringsVar(&IndexAuditEntriesRunner.filter)
	IndexAuditEntriesRunner.Flag(`limit`, `Limit the audit entries to this number. The limit should >= 1 and <= 1000`).Required().StringVar(&IndexAuditEntriesRunner.limit)
	IndexAuditEntriesRunner.Flag(`startDate`, `The start date for retrieving audit entries, the format must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g., 2011/06/25 00:00:00 +0000`).Required().StringVar(&IndexAuditEntriesRunner.startDate)
	IndexAuditEntriesRunner.Flag(`view`, ``).StringVar(&IndexAuditEntriesRunner.view)
	registry[IndexAuditEntriesCmd.FullCommand()] = IndexAuditEntriesRunner

	ShowAuditEntryRunner := new(ShowAuditEntryAuditEntryRunner)
	ShowAuditEntryCmd := resCmd.Command("ShowAuditEntry", `Lists the attributes of a given audit entry.`)
	ShowAuditEntryRunner.Flag(`id`, ``).Required().StringVar(&ShowAuditEntryRunner.id)
	ShowAuditEntryRunner.Flag(`view`, ``).StringVar(&ShowAuditEntryRunner.view)
	registry[ShowAuditEntryCmd.FullCommand()] = ShowAuditEntryRunner

	UpdateAuditEntryRunner := new(UpdateAuditEntryAuditEntryRunner)
	UpdateAuditEntryCmd := resCmd.Command("UpdateAuditEntry", `Updates the summary of a given AuditEntry.`)
	UpdateAuditEntryRunner.Flag(`auditEntry.offset`, `The offset where the next details will be appended. Used in ordering of summary updates.`).IntVar(&UpdateAuditEntryRunner.auditEntryOffset)
	UpdateAuditEntryRunner.Flag(`auditEntry.summary`, `The updated summary for the audit entry, maximum length is 255 characters.`).Required().StringVar(&UpdateAuditEntryRunner.auditEntrySummary)
	UpdateAuditEntryRunner.Flag(`id`, ``).Required().StringVar(&UpdateAuditEntryRunner.id)
	UpdateAuditEntryRunner.Flag(`notify`, `The event notification category. Defaults to 'None'.`).StringVar(&UpdateAuditEntryRunner.notify)
	registry[UpdateAuditEntryCmd.FullCommand()] = UpdateAuditEntryRunner
}

/****** Backup ******/

type CleanupBackupBackupRunner struct {
	cloudHref string
	dailies   string
	keepLast  string
	lineage   string
	monthlies string
	weeklies  string
	yearlies  string
}

func (r *CleanupBackupBackupRunner) Run(c *Client) (interface{}, error) {
	return c.CleanupBackup(r.cloudHref, r.dailies, r.keepLast, r.lineage, r.monthlies, r.weeklies, r.yearlies)
}

type CreateBackupBackupRunner struct {
	backupDescription string
	backupFromMaster  string
	backupLineage     string
	backupName        string
	backupItem        []string
	backupItemPos     []string
}

func (r *CreateBackupBackupRunner) Run(c *Client) (interface{}, error) {

	/** Handle backup parameter **/
	var backup BackupParam

	// Load JSON if provided
	if len(r.backupJson) > 0 {
		if err := Json.Unmarshal(r.backupJson, &backup); err != nil {
			return nil, fmt.Errorf("Could not load backup JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.backupdescription) > 0 {
		backup.backup.description = r.backupdescription
	}

	if len(r.backupfromMaster) > 0 {
		backup.backup.fromMaster = r.backupfromMaster
	}

	if len(r.backuplineage) > 0 {
		backup.backup.lineage = r.backuplineage
	}

	if len(r.backupname) > 0 {
		backup.backup.name = r.backupname
	}

	if len(r.backupvolumeAttachmentHrefsitem) > 0 {
		backup.backup.volumeAttachmentHrefs.item = r.backupvolumeAttachmentHrefsitem
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxbackupvolumeAttachmentHrefsitemPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.backupvolumeAttachmentHrefsitemPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for backup.volumeAttachmentHrefs.item field of volumeAttachmentHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of backup.volumeAttachmentHrefs.item field of volumeAttachmentHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxbackupvolumeAttachmentHrefsitemPos {
			maxbackupvolumeAttachmentHrefsitemPos = pos
		}
	}
	if len(r.backupvolumeAttachmentHrefsitem) != maxbackupvolumeAttachmentHrefsitemPos {
		return nil, fmt.Errof("Invalid flags for backup.volumeAttachmentHrefs.item field of volumeAttachmentHrefs array, %s were defined but max position is %s",
			len(r.backupvolumeAttachmentHrefsitem), maxbackupvolumeAttachmentHrefsitemPos)
	}
	if maxbackupvolumeAttachmentHrefsitemPos > -1 {
		backupvolumeAttachmentHrefs := make([]*string, maxbackupvolumeAttachmentHrefsitemPos+1)
		for i := 0; i < maxbackupvolumeAttachmentHrefsitemPos+1; i++ {
			backupvolumeAttachmentHrefs[i] = string{
			//TBD
			}
		}
		backup.volumeAttachmentHrefs = backupvolumeAttachmentHrefs
	}

	return c.CreateBackup(&backup)
}

type DestroyBackupBackupRunner struct {
	id string
}

func (r *DestroyBackupBackupRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyBackup(r.id)
}

type IndexBackupsBackupRunner struct {
	filter    []string
	filterPos []string
	lineage   string
}

func (r *IndexBackupsBackupRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexBackups(filter, r.lineage)
}

type RestoreBackupBackupRunner struct {
	backupDescription    string
	backupIops           string
	backupName           string
	backupSize           string
	backupVolumeTypeHref string
	id                   string
	instanceHref         string
}

func (r *RestoreBackupBackupRunner) Run(c *Client) (interface{}, error) {

	/** Handle backup parameter **/
	var backup BackupParam2

	// Load JSON if provided
	if len(r.backupJson) > 0 {
		if err := Json.Unmarshal(r.backupJson, &backup); err != nil {
			return nil, fmt.Errorf("Could not load backup JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.backupdescription) > 0 {
		backup.backup.description = r.backupdescription
	}

	if len(r.backupiops) > 0 {
		backup.backup.iops = r.backupiops
	}

	if len(r.backupname) > 0 {
		backup.backup.name = r.backupname
	}

	if len(r.backupsize) > 0 {
		backup.backup.size = r.backupsize
	}

	if len(r.backupvolumeTypeHref) > 0 {
		backup.backup.volumeTypeHref = r.backupvolumeTypeHref
	}

	return c.RestoreBackup(&backup, r.id, r.instanceHref)
}

type ShowBackupBackupRunner struct {
	id string
}

func (r *ShowBackupBackupRunner) Run(c *Client) (interface{}, error) {
	return c.ShowBackup(r.id)
}

type UpdateBackupBackupRunner struct {
	backupCommitted string
	id              string
}

func (r *UpdateBackupBackupRunner) Run(c *Client) (interface{}, error) {

	/** Handle backup parameter **/
	var backup BackupParam2

	// Load JSON if provided
	if len(r.backupJson) > 0 {
		if err := Json.Unmarshal(r.backupJson, &backup); err != nil {
			return nil, fmt.Errorf("Could not load backup JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.backupcommitted) > 0 {
		backup.backup.committed = r.backupcommitted
	}

	return c.UpdateBackup(&backup, r.id)
}

// Register all Backup actions
func registerBackupCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Backup", ``)

	CleanupBackupRunner := new(CleanupBackupBackupRunner)
	CleanupBackupCmd := resCmd.Command("CleanupBackup", `Deletes old backups that meet the given criteria`)
	CleanupBackupRunner.Flag(`cloudHref`, `Backups belonging to only this cloud are considered for cleanup. Otherwise, all backups in the account with the same lineage will be considered.`).StringVar(&CleanupBackupRunner.cloudHref)
	CleanupBackupRunner.Flag(`dailies`, `The number of daily backups(the latest one in each day) that should be kept.`).StringVar(&CleanupBackupRunner.dailies)
	CleanupBackupRunner.Flag(`keepLast`, `The number of backups that should be kept.`).Required().StringVar(&CleanupBackupRunner.keepLast)
	CleanupBackupRunner.Flag(`lineage`, `The lineage of the backups that are to be cleaned-up.`).Required().StringVar(&CleanupBackupRunner.lineage)
	CleanupBackupRunner.Flag(`monthlies`, `The number of monthly backups(the latest one in each month) that should be kept.`).StringVar(&CleanupBackupRunner.monthlies)
	CleanupBackupRunner.Flag(`weeklies`, `The number of weekly backups(the latest one in each week) that should be kept.`).StringVar(&CleanupBackupRunner.weeklies)
	CleanupBackupRunner.Flag(`yearlies`, `The number of yearly backups(the latest one in each year) that should be kept.`).StringVar(&CleanupBackupRunner.yearlies)
	registry[CleanupBackupCmd.FullCommand()] = CleanupBackupRunner

	CreateBackupRunner := new(CreateBackupBackupRunner)
	CreateBackupCmd := resCmd.Command("CreateBackup", `Takes in an array of volumeattachmenthrefs and takes a snapshot of each.`)
	CreateBackupRunner.Flag(`backup.description`, `The description to be set on each of the volume snapshots`).StringVar(&CreateBackupRunner.backupDescription)
	CreateBackupRunner.Flag(`backup.fromMaster`, `Setting this to 'true' will create a tag 'rs_backup:from_master=true' on the snapshots so that one can filter them later.`).StringVar(&CreateBackupRunner.backupFromMaster)
	CreateBackupRunner.Flag(`backup.lineage`, `A unique value to create backups belonging to a particular system.`).Required().StringVar(&CreateBackupRunner.backupLineage)
	CreateBackupRunner.Flag(`backup.name`, `The name to be set on each of the volume snapshots.`).Required().StringVar(&CreateBackupRunner.backupName)
	CreateBackupRunner.FlagPattern(`backup\.item\.(\d+)`, `List of volume attachment hrefs that are to be backed-up.`).Required().Capture(&CreateBackupRunner.backupItemPos).StringsVar(&CreateBackupRunner.backupItem)
	registry[CreateBackupCmd.FullCommand()] = CreateBackupRunner

	DestroyBackupRunner := new(DestroyBackupBackupRunner)
	DestroyBackupCmd := resCmd.Command("DestroyBackup", `Deletes a given backup by deleting all of its snapshots, this call will succeed even if the backup has not completed.`)
	DestroyBackupRunner.Flag(`id`, ``).Required().StringVar(&DestroyBackupRunner.id)
	registry[DestroyBackupCmd.FullCommand()] = DestroyBackupRunner

	IndexBackupsRunner := new(IndexBackupsBackupRunner)
	IndexBackupsCmd := resCmd.Command("IndexBackups", `Lists all of the backups with the given lineage tag`)
	IndexBackupsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexBackupsRunner.filterPos).StringsVar(&IndexBackupsRunner.filter)
	IndexBackupsRunner.Flag(`lineage`, `Backups belonging to this lineage.`).Required().StringVar(&IndexBackupsRunner.lineage)
	registry[IndexBackupsCmd.FullCommand()] = IndexBackupsRunner

	RestoreBackupRunner := new(RestoreBackupBackupRunner)
	RestoreBackupCmd := resCmd.Command("RestoreBackup", `Restores the given Backup.`)
	RestoreBackupRunner.Flag(`backup.description`, `Each volume is created with this description instead of the volume snapshot's description`).StringVar(&RestoreBackupRunner.backupDescription)
	RestoreBackupRunner.Flag(`backup.iops`, `The number of IOPS (I/O Operations Per Second) each volume should support. Only available on clouds supporting performance provisioning.`).StringVar(&RestoreBackupRunner.backupIops)
	RestoreBackupRunner.Flag(`backup.name`, `Each volume is created with this name instead of the volume snapshot's name`).StringVar(&RestoreBackupRunner.backupName)
	RestoreBackupRunner.Flag(`backup.size`, `Each volume is created with this size in gigabytes (GB) instead of the volume snapshot's size (must be equal or larger). Some volume types have predefined sizes and do not allow selecting a custom size on volume creation.`).StringVar(&RestoreBackupRunner.backupSize)
	RestoreBackupRunner.Flag(`backup.volumeTypeHref`, `The href of the volume type. Each volume is created with this volume type instead of the default volume type for the cloud. A Name, Resource UID and optional Size is associated with a volume type.`).StringVar(&RestoreBackupRunner.backupVolumeTypeHref)
	RestoreBackupRunner.Flag(`id`, ``).Required().StringVar(&RestoreBackupRunner.id)
	RestoreBackupRunner.Flag(`instanceHref`, `The instance href that the backup will be restored to.`).Required().StringVar(&RestoreBackupRunner.instanceHref)
	registry[RestoreBackupCmd.FullCommand()] = RestoreBackupRunner

	ShowBackupRunner := new(ShowBackupBackupRunner)
	ShowBackupCmd := resCmd.Command("ShowBackup", `Lists the attributes of a given backup`)
	ShowBackupRunner.Flag(`id`, ``).Required().StringVar(&ShowBackupRunner.id)
	registry[ShowBackupCmd.FullCommand()] = ShowBackupRunner

	UpdateBackupRunner := new(UpdateBackupBackupRunner)
	UpdateBackupCmd := resCmd.Command("UpdateBackup", `Updates the committed tag for all of the VolumeSnapshots in the given Backup to the given value.`)
	UpdateBackupRunner.Flag(`backup.committed`, `Setting this to 'true' will update the 'rs_backup:committed=false' tag to 'rs_backup:committed=true' on all the snapshots.`).Required().StringVar(&UpdateBackupRunner.backupCommitted)
	UpdateBackupRunner.Flag(`id`, ``).Required().StringVar(&UpdateBackupRunner.id)
	registry[UpdateBackupCmd.FullCommand()] = UpdateBackupRunner
}

/****** ChildAccount ******/

type CreateChildAccountChildAccountRunner struct {
	childAccountClusterHref string
	childAccountName        string
}

func (r *CreateChildAccountChildAccountRunner) Run(c *Client) (interface{}, error) {

	/** Handle childAccount parameter **/
	var childAccount ChildAccountParam

	// Load JSON if provided
	if len(r.childAccountJson) > 0 {
		if err := Json.Unmarshal(r.childAccountJson, &childAccount); err != nil {
			return nil, fmt.Errorf("Could not load childAccount JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.childAccountclusterHref) > 0 {
		childAccount.childAccount.clusterHref = r.childAccountclusterHref
	}

	if len(r.childAccountname) > 0 {
		childAccount.childAccount.name = r.childAccountname
	}

	return c.CreateChildAccount(&childAccount)
}

type IndexChildAccountsChildAccountRunner struct {
	filter    []string
	filterPos []string
}

func (r *IndexChildAccountsChildAccountRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexChildAccounts(filter)
}

type UpdateChildAccountChildAccountRunner struct {
	childAccountName string
	id               string
}

func (r *UpdateChildAccountChildAccountRunner) Run(c *Client) (interface{}, error) {

	/** Handle childAccount parameter **/
	var childAccount ChildAccountParam2

	// Load JSON if provided
	if len(r.childAccountJson) > 0 {
		if err := Json.Unmarshal(r.childAccountJson, &childAccount); err != nil {
			return nil, fmt.Errorf("Could not load childAccount JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.childAccountname) > 0 {
		childAccount.childAccount.name = r.childAccountname
	}

	return c.UpdateChildAccount(&childAccount, r.id)
}

// Register all ChildAccount actions
func registerChildAccountCmds(app *kingpin.Application) {
	resCmd := app.Cmd("ChildAccount", ``)

	CreateChildAccountRunner := new(CreateChildAccountChildAccountRunner)
	CreateChildAccountCmd := resCmd.Command("CreateChildAccount", `Create an enterprise ChildAccount for this Account`)
	CreateChildAccountRunner.Flag(`childAccount.clusterHref`, `The href of the cluster in which to create the account. If not specified, will default to the cluster of the parent account.`).StringVar(&CreateChildAccountRunner.childAccountClusterHref)
	CreateChildAccountRunner.Flag(`childAccount.name`, ``).Required().StringVar(&CreateChildAccountRunner.childAccountName)
	registry[CreateChildAccountCmd.FullCommand()] = CreateChildAccountRunner

	IndexChildAccountsRunner := new(IndexChildAccountsChildAccountRunner)
	IndexChildAccountsCmd := resCmd.Command("IndexChildAccounts", `Lists the enterprise ChildAccounts available for this Account.`)
	IndexChildAccountsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexChildAccountsRunner.filterPos).StringsVar(&IndexChildAccountsRunner.filter)
	registry[IndexChildAccountsCmd.FullCommand()] = IndexChildAccountsRunner

	UpdateChildAccountRunner := new(UpdateChildAccountChildAccountRunner)
	UpdateChildAccountCmd := resCmd.Command("UpdateChildAccount", `Update an enterprise ChildAccount for this Account.`)
	UpdateChildAccountRunner.Flag(`childAccount.name`, `The updated name for the account.`).StringVar(&UpdateChildAccountRunner.childAccountName)
	UpdateChildAccountRunner.Flag(`id`, ``).Required().StringVar(&UpdateChildAccountRunner.id)
	registry[UpdateChildAccountCmd.FullCommand()] = UpdateChildAccountRunner
}

/****** Cloud ******/

type IndexCloudsCloudRunner struct {
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexCloudsCloudRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexClouds(filter, r.view)
}

type ShowCloudCloudRunner struct {
	id   string
	view string
}

func (r *ShowCloudCloudRunner) Run(c *Client) (interface{}, error) {
	return c.ShowCloud(r.id, r.view)
}

// Register all Cloud actions
func registerCloudCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Cloud", `Represents a Cloud (within the context of the account in the session).`)

	IndexCloudsRunner := new(IndexCloudsCloudRunner)
	IndexCloudsCmd := resCmd.Command("IndexClouds", `Lists the clouds available to this account.`)
	IndexCloudsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexCloudsRunner.filterPos).StringsVar(&IndexCloudsRunner.filter)
	IndexCloudsRunner.Flag(`view`, ``).StringVar(&IndexCloudsRunner.view)
	registry[IndexCloudsCmd.FullCommand()] = IndexCloudsRunner

	ShowCloudRunner := new(ShowCloudCloudRunner)
	ShowCloudCmd := resCmd.Command("ShowCloud", `Show information about a single cloud`)
	ShowCloudRunner.Flag(`id`, ``).Required().StringVar(&ShowCloudRunner.id)
	ShowCloudRunner.Flag(`view`, ``).StringVar(&ShowCloudRunner.view)
	registry[ShowCloudCmd.FullCommand()] = ShowCloudRunner
}

/****** CloudAccount ******/

type CreateCloudAccountCloudAccountRunner struct {
	cloudAccountCloudHref   string
	cloudAccountCredsValues []string
	cloudAccountCredsNames  []string
	cloudAccountToken       string
}

func (r *CreateCloudAccountCloudAccountRunner) Run(c *Client) (interface{}, error) {

	/** Handle cloudAccount parameter **/
	var cloudAccount CloudAccountParam

	// Load JSON if provided
	if len(r.cloudAccountJson) > 0 {
		if err := Json.Unmarshal(r.cloudAccountJson, &cloudAccount); err != nil {
			return nil, fmt.Errorf("Could not load cloudAccount JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.cloudAccountcloudHref) > 0 {
		cloudAccount.cloudAccount.cloudHref = r.cloudAccountcloudHref
	}

	if len(r.cloudAccounttoken) > 0 {
		cloudAccount.cloudAccount.token = r.cloudAccounttoken
	}

	// Create enumarable fields from flags
	cloudAccountcreds := make(map[string]string, len(r.cloudAccountcredsNames))
	for i, n := range r.cloudAccountcredsNames {
		cloudAccountcreds[n] = r.cloudAccountcredsValues[i]
	}
	cloudAccount.cloudAccount.creds = cloudAccountcreds

	return c.CreateCloudAccount(&cloudAccount)
}

type DestroyCloudAccountCloudAccountRunner struct {
	id string
}

func (r *DestroyCloudAccountCloudAccountRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyCloudAccount(r.id)
}

type IndexCloudAccountsCloudAccountRunner struct {
}

func (r *IndexCloudAccountsCloudAccountRunner) Run(c *Client) (interface{}, error) {
	return c.IndexCloudAccounts()
}

type ShowCloudAccountCloudAccountRunner struct {
	id string
}

func (r *ShowCloudAccountCloudAccountRunner) Run(c *Client) (interface{}, error) {
	return c.ShowCloudAccount(r.id)
}

// Register all CloudAccount actions
func registerCloudAccountCmds(app *kingpin.Application) {
	resCmd := app.Cmd("CloudAccount", `Represents a Cloud Account (An association between the account and a cloud).`)

	CreateCloudAccountRunner := new(CreateCloudAccountCloudAccountRunner)
	CreateCloudAccountCmd := resCmd.Command("CreateCloudAccount", `Create a CloudAccount by passing in the respective credentials for each cloud.`)
	CreateCloudAccountRunner.Flag(`cloudAccount.cloudHref`, `The href of the cloud if it is known. For valid values see support portal link above.`).StringVar(&CreateCloudAccountRunner.cloudAccountCloudHref)
	CreateCloudAccountRunner.FlagPattern(`cloudAccount\.creds\.([a-z0-9_]+)`, ``).Required().Capture(&CreateCloudAccountRunner.cloudAccountCredsNames).StringVar(&CreateCloudAccountRunner.cloudAccountCredsValues)
	CreateCloudAccountRunner.Flag(`cloudAccount.token`, `The cloud token to identify a private cloud`).StringVar(&CreateCloudAccountRunner.cloudAccountToken)
	registry[CreateCloudAccountCmd.FullCommand()] = CreateCloudAccountRunner

	DestroyCloudAccountRunner := new(DestroyCloudAccountCloudAccountRunner)
	DestroyCloudAccountCmd := resCmd.Command("DestroyCloudAccount", `Delete a CloudAccount.`)
	DestroyCloudAccountRunner.Flag(`id`, ``).Required().StringVar(&DestroyCloudAccountRunner.id)
	registry[DestroyCloudAccountCmd.FullCommand()] = DestroyCloudAccountRunner

	IndexCloudAccountsRunner := new(IndexCloudAccountsCloudAccountRunner)
	IndexCloudAccountsCmd := resCmd.Command("IndexCloudAccounts", `Lists the CloudAccounts (non-aws) available to this Account.`)
	registry[IndexCloudAccountsCmd.FullCommand()] = IndexCloudAccountsRunner

	ShowCloudAccountRunner := new(ShowCloudAccountCloudAccountRunner)
	ShowCloudAccountCmd := resCmd.Command("ShowCloudAccount", ``)
	ShowCloudAccountRunner.Flag(`id`, ``).Required().StringVar(&ShowCloudAccountRunner.id)
	registry[ShowCloudAccountCmd.FullCommand()] = ShowCloudAccountRunner
}

/****** Cookbook ******/

type DestroyCookbookCookbookRunner struct {
	id string
}

func (r *DestroyCookbookCookbookRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyCookbook(r.id)
}

type FollowCookbookCookbookRunner struct {
	id    string
	value string
}

func (r *FollowCookbookCookbookRunner) Run(c *Client) (interface{}, error) {
	return c.FollowCookbook(r.id, r.value)
}

type FreezeCookbookCookbookRunner struct {
	id    string
	value string
}

func (r *FreezeCookbookCookbookRunner) Run(c *Client) (interface{}, error) {
	return c.FreezeCookbook(r.id, r.value)
}

type IndexCookbooksCookbookRunner struct {
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexCookbooksCookbookRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexCookbooks(filter, r.view)
}

type ObsoleteCookbookCookbookRunner struct {
	id    string
	value string
}

func (r *ObsoleteCookbookCookbookRunner) Run(c *Client) (interface{}, error) {
	return c.ObsoleteCookbook(r.id, r.value)
}

type ShowCookbookCookbookRunner struct {
	id   string
	view string
}

func (r *ShowCookbookCookbookRunner) Run(c *Client) (interface{}, error) {
	return c.ShowCookbook(r.id, r.view)
}

// Register all Cookbook actions
func registerCookbookCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Cookbook", `Represents a given instance of a single cookbook.`)

	DestroyCookbookRunner := new(DestroyCookbookCookbookRunner)
	DestroyCookbookCmd := resCmd.Command("DestroyCookbook", `Destroys a Cookbook. Only available for cookbooks that have no Cookbook Attachments.`)
	DestroyCookbookRunner.Flag(`id`, ``).Required().StringVar(&DestroyCookbookRunner.id)
	registry[DestroyCookbookCmd.FullCommand()] = DestroyCookbookRunner

	FollowCookbookRunner := new(FollowCookbookCookbookRunner)
	FollowCookbookCmd := resCmd.Command("FollowCookbook", `Follows (or unfollows) a Cookbook. Only available for cookbooks that are in the Alternate namespace.`)
	FollowCookbookRunner.Flag(`id`, ``).Required().StringVar(&FollowCookbookRunner.id)
	FollowCookbookRunner.Flag(`value`, `Indicates if this action should follow (true) or unfollow (false) a Cookbook.`).Required().StringVar(&FollowCookbookRunner.value)
	registry[FollowCookbookCmd.FullCommand()] = FollowCookbookRunner

	FreezeCookbookRunner := new(FreezeCookbookCookbookRunner)
	FreezeCookbookCmd := resCmd.Command("FreezeCookbook", `Freezes (or unfreezes) a Cookbook. Only available for cookbooks that are in the Primary namespace.`)
	FreezeCookbookRunner.Flag(`id`, ``).Required().StringVar(&FreezeCookbookRunner.id)
	FreezeCookbookRunner.Flag(`value`, `Indicates if this action should freeze (true) or unfreeze (false) a Cookbook.`).Required().StringVar(&FreezeCookbookRunner.value)
	registry[FreezeCookbookCmd.FullCommand()] = FreezeCookbookRunner

	IndexCookbooksRunner := new(IndexCookbooksCookbookRunner)
	IndexCookbooksCmd := resCmd.Command("IndexCookbooks", `Lists the Cookbooks available to this account.`)
	IndexCookbooksRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexCookbooksRunner.filterPos).StringsVar(&IndexCookbooksRunner.filter)
	IndexCookbooksRunner.Flag(`view`, ``).StringVar(&IndexCookbooksRunner.view)
	registry[IndexCookbooksCmd.FullCommand()] = IndexCookbooksRunner

	ObsoleteCookbookRunner := new(ObsoleteCookbookCookbookRunner)
	ObsoleteCookbookCmd := resCmd.Command("ObsoleteCookbook", `Marks a Cookbook as obsolete (or un-obsolete).`)
	ObsoleteCookbookRunner.Flag(`id`, ``).Required().StringVar(&ObsoleteCookbookRunner.id)
	ObsoleteCookbookRunner.Flag(`value`, `Indicates if this action should obsolete (true) or un-obsolete (false) a Cookbook.`).Required().StringVar(&ObsoleteCookbookRunner.value)
	registry[ObsoleteCookbookCmd.FullCommand()] = ObsoleteCookbookRunner

	ShowCookbookRunner := new(ShowCookbookCookbookRunner)
	ShowCookbookCmd := resCmd.Command("ShowCookbook", `Show information about a single Cookbook.`)
	ShowCookbookRunner.Flag(`id`, ``).Required().StringVar(&ShowCookbookRunner.id)
	ShowCookbookRunner.Flag(`view`, ``).StringVar(&ShowCookbookRunner.view)
	registry[ShowCookbookCmd.FullCommand()] = ShowCookbookRunner
}

/****** CookbookAttachment ******/

type CreateCookbookAttachmentCookbookAttachmentRunner struct {
	cookbookAttachmentCookbookHref       string
	cookbookAttachmentServerTemplateHref string
	cookbookId                           string
}

func (r *CreateCookbookAttachmentCookbookAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle cookbookAttachment parameter **/
	var cookbookAttachment CookbookAttachmentParam

	// Load JSON if provided
	if len(r.cookbookAttachmentJson) > 0 {
		if err := Json.Unmarshal(r.cookbookAttachmentJson, &cookbookAttachment); err != nil {
			return nil, fmt.Errorf("Could not load cookbookAttachment JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.cookbookAttachmentcookbookHref) > 0 {
		cookbookAttachment.cookbookAttachment.cookbookHref = r.cookbookAttachmentcookbookHref
	}

	if len(r.cookbookAttachmentserverTemplateHref) > 0 {
		cookbookAttachment.cookbookAttachment.serverTemplateHref = r.cookbookAttachmentserverTemplateHref
	}

	return c.CreateCookbookAttachment(&cookbookAttachment, r.cookbookId)
}

type DestroyCookbookAttachmentCookbookAttachmentRunner struct {
	cookbookId string
	id         string
}

func (r *DestroyCookbookAttachmentCookbookAttachmentRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyCookbookAttachment(r.cookbookId, r.id)
}

type IndexCookbookAttachmentsCookbookAttachmentRunner struct {
	cookbookId string
	view       string
}

func (r *IndexCookbookAttachmentsCookbookAttachmentRunner) Run(c *Client) (interface{}, error) {
	return c.IndexCookbookAttachments(r.cookbookId, r.view)
}

type MultiAttachCookbookAttachmentsCookbookAttachmentRunner struct {
	cookbookAttachmentsItem               []string
	cookbookAttachmentsItemPos            []string
	cookbookAttachmentsServerTemplateHref string
	serverTemplateId                      string
}

func (r *MultiAttachCookbookAttachmentsCookbookAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle cookbookAttachments parameter **/
	var cookbookAttachments CookbookAttachments

	// Load JSON if provided
	if len(r.cookbookAttachmentsJson) > 0 {
		if err := Json.Unmarshal(r.cookbookAttachmentsJson, &cookbookAttachments); err != nil {
			return nil, fmt.Errorf("Could not load cookbookAttachments JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.cookbookAttachmentscookbookHrefsitem) > 0 {
		cookbookAttachments.cookbookAttachments.cookbookHrefs.item = r.cookbookAttachmentscookbookHrefsitem
	}

	if len(r.cookbookAttachmentsserverTemplateHref) > 0 {
		cookbookAttachments.cookbookAttachments.serverTemplateHref = r.cookbookAttachmentsserverTemplateHref
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxcookbookAttachmentscookbookHrefsitemPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.cookbookAttachmentscookbookHrefsitemPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for cookbookAttachments.cookbookHrefs.item field of cookbookHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of cookbookAttachments.cookbookHrefs.item field of cookbookHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxcookbookAttachmentscookbookHrefsitemPos {
			maxcookbookAttachmentscookbookHrefsitemPos = pos
		}
	}
	if len(r.cookbookAttachmentscookbookHrefsitem) != maxcookbookAttachmentscookbookHrefsitemPos {
		return nil, fmt.Errof("Invalid flags for cookbookAttachments.cookbookHrefs.item field of cookbookHrefs array, %s were defined but max position is %s",
			len(r.cookbookAttachmentscookbookHrefsitem), maxcookbookAttachmentscookbookHrefsitemPos)
	}
	if maxcookbookAttachmentscookbookHrefsitemPos > -1 {
		cookbookAttachmentscookbookHrefs := make([]*string, maxcookbookAttachmentscookbookHrefsitemPos+1)
		for i := 0; i < maxcookbookAttachmentscookbookHrefsitemPos+1; i++ {
			cookbookAttachmentscookbookHrefs[i] = string{
			//TBD
			}
		}
		cookbookAttachments.cookbookHrefs = cookbookAttachmentscookbookHrefs
	}

	return c.MultiAttachCookbookAttachments(&cookbookAttachments, r.serverTemplateId)
}

type MultiDetachCookbookAttachmentsCookbookAttachmentRunner struct {
	cookbookAttachmentsItem    []string
	cookbookAttachmentsItemPos []string
	serverTemplateId           string
}

func (r *MultiDetachCookbookAttachmentsCookbookAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle cookbookAttachments parameter **/
	var cookbookAttachments CookbookAttachments2

	// Load JSON if provided
	if len(r.cookbookAttachmentsJson) > 0 {
		if err := Json.Unmarshal(r.cookbookAttachmentsJson, &cookbookAttachments); err != nil {
			return nil, fmt.Errorf("Could not load cookbookAttachments JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.cookbookAttachmentscookbookAttachmentHrefsitem) > 0 {
		cookbookAttachments.cookbookAttachments.cookbookAttachmentHrefs.item = r.cookbookAttachmentscookbookAttachmentHrefsitem
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxcookbookAttachmentscookbookAttachmentHrefsitemPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.cookbookAttachmentscookbookAttachmentHrefsitemPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for cookbookAttachments.cookbookAttachmentHrefs.item field of cookbookAttachmentHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of cookbookAttachments.cookbookAttachmentHrefs.item field of cookbookAttachmentHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxcookbookAttachmentscookbookAttachmentHrefsitemPos {
			maxcookbookAttachmentscookbookAttachmentHrefsitemPos = pos
		}
	}
	if len(r.cookbookAttachmentscookbookAttachmentHrefsitem) != maxcookbookAttachmentscookbookAttachmentHrefsitemPos {
		return nil, fmt.Errof("Invalid flags for cookbookAttachments.cookbookAttachmentHrefs.item field of cookbookAttachmentHrefs array, %s were defined but max position is %s",
			len(r.cookbookAttachmentscookbookAttachmentHrefsitem), maxcookbookAttachmentscookbookAttachmentHrefsitemPos)
	}
	if maxcookbookAttachmentscookbookAttachmentHrefsitemPos > -1 {
		cookbookAttachmentscookbookAttachmentHrefs := make([]*string, maxcookbookAttachmentscookbookAttachmentHrefsitemPos+1)
		for i := 0; i < maxcookbookAttachmentscookbookAttachmentHrefsitemPos+1; i++ {
			cookbookAttachmentscookbookAttachmentHrefs[i] = string{
			//TBD
			}
		}
		cookbookAttachments.cookbookAttachmentHrefs = cookbookAttachmentscookbookAttachmentHrefs
	}

	return c.MultiDetachCookbookAttachments(&cookbookAttachments, r.serverTemplateId)
}

type ShowCookbookAttachmentCookbookAttachmentRunner struct {
	cookbookId string
	id         string
	view       string
}

func (r *ShowCookbookAttachmentCookbookAttachmentRunner) Run(c *Client) (interface{}, error) {
	return c.ShowCookbookAttachment(r.cookbookId, r.id, r.view)
}

// Register all CookbookAttachment actions
func registerCookbookAttachmentCmds(app *kingpin.Application) {
	resCmd := app.Cmd("CookbookAttachment", `Cookbook Attachment is used to associate a particular cookbook with a ServerTemplate. A Cookbook Attachment must be in place before a recipe can be bound to a runlist using RunnableBinding.`)

	CreateCookbookAttachmentRunner := new(CreateCookbookAttachmentCookbookAttachmentRunner)
	CreateCookbookAttachmentCmd := resCmd.Command("CreateCookbookAttachment", `Attach a cookbook to a given resource.`)
	CreateCookbookAttachmentRunner.Flag(`cookbookAttachment.cookbookHref`, `The href of the cookbook to attach.`).StringVar(&CreateCookbookAttachmentRunner.cookbookAttachmentCookbookHref)
	CreateCookbookAttachmentRunner.Flag(`cookbookAttachment.serverTemplateHref`, `The href of the server template to attach the cookbook to.`).StringVar(&CreateCookbookAttachmentRunner.cookbookAttachmentServerTemplateHref)
	CreateCookbookAttachmentRunner.Flag(`cookbookId`, ``).Required().StringVar(&CreateCookbookAttachmentRunner.cookbookId)
	registry[CreateCookbookAttachmentCmd.FullCommand()] = CreateCookbookAttachmentRunner

	DestroyCookbookAttachmentRunner := new(DestroyCookbookAttachmentCookbookAttachmentRunner)
	DestroyCookbookAttachmentCmd := resCmd.Command("DestroyCookbookAttachment", `Detach a cookbook from a given resource.`)
	DestroyCookbookAttachmentRunner.Flag(`cookbookId`, ``).Required().StringVar(&DestroyCookbookAttachmentRunner.cookbookId)
	DestroyCookbookAttachmentRunner.Flag(`id`, ``).Required().StringVar(&DestroyCookbookAttachmentRunner.id)
	registry[DestroyCookbookAttachmentCmd.FullCommand()] = DestroyCookbookAttachmentRunner

	IndexCookbookAttachmentsRunner := new(IndexCookbookAttachmentsCookbookAttachmentRunner)
	IndexCookbookAttachmentsCmd := resCmd.Command("IndexCookbookAttachments", `Lists Cookbook Attachments.`)
	IndexCookbookAttachmentsRunner.Flag(`cookbookId`, ``).Required().StringVar(&IndexCookbookAttachmentsRunner.cookbookId)
	IndexCookbookAttachmentsRunner.Flag(`view`, ``).StringVar(&IndexCookbookAttachmentsRunner.view)
	registry[IndexCookbookAttachmentsCmd.FullCommand()] = IndexCookbookAttachmentsRunner

	MultiAttachCookbookAttachmentsRunner := new(MultiAttachCookbookAttachmentsCookbookAttachmentRunner)
	MultiAttachCookbookAttachmentsCmd := resCmd.Command("MultiAttachCookbookAttachments", `Attach multiple cookbooks to a given resource.`)
	MultiAttachCookbookAttachmentsRunner.FlagPattern(`cookbookAttachments\.item\.(\d+)`, `The hrefs of the cookbooks to be attached`).Capture(&MultiAttachCookbookAttachmentsRunner.cookbookAttachmentsItemPos).StringsVar(&MultiAttachCookbookAttachmentsRunner.cookbookAttachmentsItem)
	MultiAttachCookbookAttachmentsRunner.Flag(`cookbookAttachments.serverTemplateHref`, `The href of the server template to attach the cookbooks to.`).StringVar(&MultiAttachCookbookAttachmentsRunner.cookbookAttachmentsServerTemplateHref)
	MultiAttachCookbookAttachmentsRunner.Flag(`serverTemplateId`, ``).Required().StringVar(&MultiAttachCookbookAttachmentsRunner.serverTemplateId)
	registry[MultiAttachCookbookAttachmentsCmd.FullCommand()] = MultiAttachCookbookAttachmentsRunner

	MultiDetachCookbookAttachmentsRunner := new(MultiDetachCookbookAttachmentsCookbookAttachmentRunner)
	MultiDetachCookbookAttachmentsCmd := resCmd.Command("MultiDetachCookbookAttachments", `Detach multiple cookbooks from a given resource.`)
	MultiDetachCookbookAttachmentsRunner.FlagPattern(`cookbookAttachments\.item\.(\d+)`, `The hrefs of the cookbook attachments to be detached`).Capture(&MultiDetachCookbookAttachmentsRunner.cookbookAttachmentsItemPos).StringsVar(&MultiDetachCookbookAttachmentsRunner.cookbookAttachmentsItem)
	MultiDetachCookbookAttachmentsRunner.Flag(`serverTemplateId`, ``).Required().StringVar(&MultiDetachCookbookAttachmentsRunner.serverTemplateId)
	registry[MultiDetachCookbookAttachmentsCmd.FullCommand()] = MultiDetachCookbookAttachmentsRunner

	ShowCookbookAttachmentRunner := new(ShowCookbookAttachmentCookbookAttachmentRunner)
	ShowCookbookAttachmentCmd := resCmd.Command("ShowCookbookAttachment", `Displays information about a single cookbook attachment to a ServerTemplate.`)
	ShowCookbookAttachmentRunner.Flag(`cookbookId`, ``).Required().StringVar(&ShowCookbookAttachmentRunner.cookbookId)
	ShowCookbookAttachmentRunner.Flag(`id`, ``).Required().StringVar(&ShowCookbookAttachmentRunner.id)
	ShowCookbookAttachmentRunner.Flag(`view`, ``).StringVar(&ShowCookbookAttachmentRunner.view)
	registry[ShowCookbookAttachmentCmd.FullCommand()] = ShowCookbookAttachmentRunner
}

/****** Credential ******/

type CreateCredentialCredentialRunner struct {
	credentialDescription string
	credentialName        string
	credentialValue       string
}

func (r *CreateCredentialCredentialRunner) Run(c *Client) (interface{}, error) {

	/** Handle credential parameter **/
	var credential CredentialParam

	// Load JSON if provided
	if len(r.credentialJson) > 0 {
		if err := Json.Unmarshal(r.credentialJson, &credential); err != nil {
			return nil, fmt.Errorf("Could not load credential JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.credentialdescription) > 0 {
		credential.credential.description = r.credentialdescription
	}

	if len(r.credentialname) > 0 {
		credential.credential.name = r.credentialname
	}

	if len(r.credentialvalue) > 0 {
		credential.credential.value = r.credentialvalue
	}

	return c.CreateCredential(&credential)
}

type DestroyCredentialCredentialRunner struct {
	id string
}

func (r *DestroyCredentialCredentialRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyCredential(r.id)
}

type IndexCredentialsCredentialRunner struct {
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexCredentialsCredentialRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexCredentials(filter, r.view)
}

type ShowCredentialCredentialRunner struct {
	id   string
	view string
}

func (r *ShowCredentialCredentialRunner) Run(c *Client) (interface{}, error) {
	return c.ShowCredential(r.id, r.view)
}

type UpdateCredentialCredentialRunner struct {
	credentialDescription string
	credentialName        string
	credentialValue       string
	id                    string
}

func (r *UpdateCredentialCredentialRunner) Run(c *Client) (interface{}, error) {

	/** Handle credential parameter **/
	var credential CredentialParam2

	// Load JSON if provided
	if len(r.credentialJson) > 0 {
		if err := Json.Unmarshal(r.credentialJson, &credential); err != nil {
			return nil, fmt.Errorf("Could not load credential JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.credentialdescription) > 0 {
		credential.credential.description = r.credentialdescription
	}

	if len(r.credentialname) > 0 {
		credential.credential.name = r.credentialname
	}

	if len(r.credentialvalue) > 0 {
		credential.credential.value = r.credentialvalue
	}

	return c.UpdateCredential(&credential, r.id)
}

// Register all Credential actions
func registerCredentialCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Credential", `A Credential provides an abstraction for sensitive textual information, such as passphrases or cloud credentials, whose value should be encrypted when stored in the database and not generally shown in the UI or through the API...`)

	CreateCredentialRunner := new(CreateCredentialCredentialRunner)
	CreateCredentialCmd := resCmd.Command("CreateCredential", `Creates a new Credential with the given parameters.`)
	CreateCredentialRunner.Flag(`credential.description`, `The description of the Credential to be created.`).StringVar(&CreateCredentialRunner.credentialDescription)
	CreateCredentialRunner.Flag(`credential.name`, `The name of the Credential to be created.`).Required().StringVar(&CreateCredentialRunner.credentialName)
	CreateCredentialRunner.Flag(`credential.value`, `The value of the Credential to be created.`).Required().StringVar(&CreateCredentialRunner.credentialValue)
	registry[CreateCredentialCmd.FullCommand()] = CreateCredentialRunner

	DestroyCredentialRunner := new(DestroyCredentialCredentialRunner)
	DestroyCredentialCmd := resCmd.Command("DestroyCredential", `Deletes a Credential.`)
	DestroyCredentialRunner.Flag(`id`, ``).Required().StringVar(&DestroyCredentialRunner.id)
	registry[DestroyCredentialCmd.FullCommand()] = DestroyCredentialRunner

	IndexCredentialsRunner := new(IndexCredentialsCredentialRunner)
	IndexCredentialsCmd := resCmd.Command("IndexCredentials", `Lists the Credentials available to this account.`)
	IndexCredentialsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexCredentialsRunner.filterPos).StringsVar(&IndexCredentialsRunner.filter)
	IndexCredentialsRunner.Flag(`view`, ``).StringVar(&IndexCredentialsRunner.view)
	registry[IndexCredentialsCmd.FullCommand()] = IndexCredentialsRunner

	ShowCredentialRunner := new(ShowCredentialCredentialRunner)
	ShowCredentialCmd := resCmd.Command("ShowCredential", `Show information about a single Credential. NOTE: Credential values may be updated through the API, but values cannot be retrieved via the API.`)
	ShowCredentialRunner.Flag(`id`, ``).Required().StringVar(&ShowCredentialRunner.id)
	ShowCredentialRunner.Flag(`view`, ``).StringVar(&ShowCredentialRunner.view)
	registry[ShowCredentialCmd.FullCommand()] = ShowCredentialRunner

	UpdateCredentialRunner := new(UpdateCredentialCredentialRunner)
	UpdateCredentialCmd := resCmd.Command("UpdateCredential", `Updates attributes of a Credential.`)
	UpdateCredentialRunner.Flag(`credential.description`, `The updated description of the Credential.`).StringVar(&UpdateCredentialRunner.credentialDescription)
	UpdateCredentialRunner.Flag(`credential.name`, `The updated name of the Credential.`).StringVar(&UpdateCredentialRunner.credentialName)
	UpdateCredentialRunner.Flag(`credential.value`, `The updated value of the Credential.`).StringVar(&UpdateCredentialRunner.credentialValue)
	UpdateCredentialRunner.Flag(`id`, ``).Required().StringVar(&UpdateCredentialRunner.id)
	registry[UpdateCredentialCmd.FullCommand()] = UpdateCredentialRunner
}

/****** Datacenter ******/

type IndexDatacentersDatacenterRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexDatacentersDatacenterRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexDatacenters(r.cloudId, filter, r.view)
}

type ShowDatacenterDatacenterRunner struct {
	cloudId string
	id      string
	view    string
}

func (r *ShowDatacenterDatacenterRunner) Run(c *Client) (interface{}, error) {
	return c.ShowDatacenter(r.cloudId, r.id, r.view)
}

// Register all Datacenter actions
func registerDatacenterCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Datacenter", `Datacenters represent isolated facilities within a cloud`)

	IndexDatacentersRunner := new(IndexDatacentersDatacenterRunner)
	IndexDatacentersCmd := resCmd.Command("IndexDatacenters", `Lists all Datacenters for a particular cloud.`)
	IndexDatacentersRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexDatacentersRunner.cloudId)
	IndexDatacentersRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexDatacentersRunner.filterPos).StringsVar(&IndexDatacentersRunner.filter)
	IndexDatacentersRunner.Flag(`view`, ``).StringVar(&IndexDatacentersRunner.view)
	registry[IndexDatacentersCmd.FullCommand()] = IndexDatacentersRunner

	ShowDatacenterRunner := new(ShowDatacenterDatacenterRunner)
	ShowDatacenterCmd := resCmd.Command("ShowDatacenter", `Displays information about a single Datacenter.`)
	ShowDatacenterRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowDatacenterRunner.cloudId)
	ShowDatacenterRunner.Flag(`id`, ``).Required().StringVar(&ShowDatacenterRunner.id)
	ShowDatacenterRunner.Flag(`view`, ``).StringVar(&ShowDatacenterRunner.view)
	registry[ShowDatacenterCmd.FullCommand()] = ShowDatacenterRunner
}

/****** Deployment ******/

type CloneDeploymentDeploymentRunner struct {
	deploymentDescription    string
	deploymentName           string
	deploymentServerTagScope string
	id                       string
}

func (r *CloneDeploymentDeploymentRunner) Run(c *Client) (interface{}, error) {

	/** Handle deployment parameter **/
	var deployment DeploymentParam

	// Load JSON if provided
	if len(r.deploymentJson) > 0 {
		if err := Json.Unmarshal(r.deploymentJson, &deployment); err != nil {
			return nil, fmt.Errorf("Could not load deployment JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.deploymentdescription) > 0 {
		deployment.deployment.description = r.deploymentdescription
	}

	if len(r.deploymentname) > 0 {
		deployment.deployment.name = r.deploymentname
	}

	if len(r.deploymentserverTagScope) > 0 {
		deployment.deployment.serverTagScope = r.deploymentserverTagScope
	}

	return c.CloneDeployment(&deployment, r.id)
}

type CreateDeploymentDeploymentRunner struct {
	deploymentDescription    string
	deploymentName           string
	deploymentServerTagScope string
}

func (r *CreateDeploymentDeploymentRunner) Run(c *Client) (interface{}, error) {

	/** Handle deployment parameter **/
	var deployment DeploymentParam2

	// Load JSON if provided
	if len(r.deploymentJson) > 0 {
		if err := Json.Unmarshal(r.deploymentJson, &deployment); err != nil {
			return nil, fmt.Errorf("Could not load deployment JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.deploymentdescription) > 0 {
		deployment.deployment.description = r.deploymentdescription
	}

	if len(r.deploymentname) > 0 {
		deployment.deployment.name = r.deploymentname
	}

	if len(r.deploymentserverTagScope) > 0 {
		deployment.deployment.serverTagScope = r.deploymentserverTagScope
	}

	return c.CreateDeployment(&deployment)
}

type DestroyDeploymentDeploymentRunner struct {
	id string
}

func (r *DestroyDeploymentDeploymentRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyDeployment(r.id)
}

type IndexDeploymentsDeploymentRunner struct {
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexDeploymentsDeploymentRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexDeployments(filter, r.view)
}

type LockDeploymentDeploymentRunner struct {
	id string
}

func (r *LockDeploymentDeploymentRunner) Run(c *Client) (interface{}, error) {
	return c.LockDeployment(r.id)
}

type ServersDeploymentDeploymentRunner struct {
	id string
}

func (r *ServersDeploymentDeploymentRunner) Run(c *Client) (interface{}, error) {
	return c.ServersDeployment(r.id)
}

type ShowDeploymentDeploymentRunner struct {
	id   string
	view string
}

func (r *ShowDeploymentDeploymentRunner) Run(c *Client) (interface{}, error) {
	return c.ShowDeployment(r.id, r.view)
}

type UnlockDeploymentDeploymentRunner struct {
	id string
}

func (r *UnlockDeploymentDeploymentRunner) Run(c *Client) (interface{}, error) {
	return c.UnlockDeployment(r.id)
}

type UpdateDeploymentDeploymentRunner struct {
	deploymentDescription    string
	deploymentName           string
	deploymentServerTagScope string
	id                       string
}

func (r *UpdateDeploymentDeploymentRunner) Run(c *Client) (interface{}, error) {

	/** Handle deployment parameter **/
	var deployment DeploymentParam

	// Load JSON if provided
	if len(r.deploymentJson) > 0 {
		if err := Json.Unmarshal(r.deploymentJson, &deployment); err != nil {
			return nil, fmt.Errorf("Could not load deployment JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.deploymentdescription) > 0 {
		deployment.deployment.description = r.deploymentdescription
	}

	if len(r.deploymentname) > 0 {
		deployment.deployment.name = r.deploymentname
	}

	if len(r.deploymentserverTagScope) > 0 {
		deployment.deployment.serverTagScope = r.deploymentserverTagScope
	}

	return c.UpdateDeployment(&deployment, r.id)
}

// Register all Deployment actions
func registerDeploymentCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Deployment", `Deployments represent logical groupings of related assets such as servers, server arrays, default configuration settings...etc.`)

	CloneDeploymentRunner := new(CloneDeploymentDeploymentRunner)
	CloneDeploymentCmd := resCmd.Command("CloneDeployment", `Clones a given deployment.`)
	CloneDeploymentRunner.Flag(`deployment.description`, `The description for the cloned deployment.`).StringVar(&CloneDeploymentRunner.deploymentDescription)
	CloneDeploymentRunner.Flag(`deployment.name`, `The name for the cloned deployment.`).StringVar(&CloneDeploymentRunner.deploymentName)
	CloneDeploymentRunner.Flag(`deployment.serverTagScope`, `The routing scope for tags for servers in the cloned deployment.`).StringVar(&CloneDeploymentRunner.deploymentServerTagScope)
	CloneDeploymentRunner.Flag(`id`, ``).Required().StringVar(&CloneDeploymentRunner.id)
	registry[CloneDeploymentCmd.FullCommand()] = CloneDeploymentRunner

	CreateDeploymentRunner := new(CreateDeploymentDeploymentRunner)
	CreateDeploymentCmd := resCmd.Command("CreateDeployment", `Creates a new deployment with the given parameters.`)
	CreateDeploymentRunner.Flag(`deployment.description`, `The description of the deployment to be created.`).StringVar(&CreateDeploymentRunner.deploymentDescription)
	CreateDeploymentRunner.Flag(`deployment.name`, `The name of the deployment to be created.`).Required().StringVar(&CreateDeploymentRunner.deploymentName)
	CreateDeploymentRunner.Flag(`deployment.serverTagScope`, `The routing scope for tags for servers in the deployment.`).StringVar(&CreateDeploymentRunner.deploymentServerTagScope)
	registry[CreateDeploymentCmd.FullCommand()] = CreateDeploymentRunner

	DestroyDeploymentRunner := new(DestroyDeploymentDeploymentRunner)
	DestroyDeploymentCmd := resCmd.Command("DestroyDeployment", `Deletes a given deployment.`)
	DestroyDeploymentRunner.Flag(`id`, ``).Required().StringVar(&DestroyDeploymentRunner.id)
	registry[DestroyDeploymentCmd.FullCommand()] = DestroyDeploymentRunner

	IndexDeploymentsRunner := new(IndexDeploymentsDeploymentRunner)
	IndexDeploymentsCmd := resCmd.Command("IndexDeployments", `Lists deployments of the account.`)
	IndexDeploymentsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexDeploymentsRunner.filterPos).StringsVar(&IndexDeploymentsRunner.filter)
	IndexDeploymentsRunner.Flag(`view`, ``).StringVar(&IndexDeploymentsRunner.view)
	registry[IndexDeploymentsCmd.FullCommand()] = IndexDeploymentsRunner

	LockDeploymentRunner := new(LockDeploymentDeploymentRunner)
	LockDeploymentCmd := resCmd.Command("LockDeployment", `Locks a given deployment. Idempotent.`)
	LockDeploymentRunner.Flag(`id`, ``).Required().StringVar(&LockDeploymentRunner.id)
	registry[LockDeploymentCmd.FullCommand()] = LockDeploymentRunner

	ServersDeploymentRunner := new(ServersDeploymentDeploymentRunner)
	ServersDeploymentCmd := resCmd.Command("ServersDeployment", `Lists the servers belonging to this deployment`)
	ServersDeploymentRunner.Flag(`id`, ``).Required().StringVar(&ServersDeploymentRunner.id)
	registry[ServersDeploymentCmd.FullCommand()] = ServersDeploymentRunner

	ShowDeploymentRunner := new(ShowDeploymentDeploymentRunner)
	ShowDeploymentCmd := resCmd.Command("ShowDeployment", `Lists the attributes of a given deployment.`)
	ShowDeploymentRunner.Flag(`id`, ``).Required().StringVar(&ShowDeploymentRunner.id)
	ShowDeploymentRunner.Flag(`view`, ``).StringVar(&ShowDeploymentRunner.view)
	registry[ShowDeploymentCmd.FullCommand()] = ShowDeploymentRunner

	UnlockDeploymentRunner := new(UnlockDeploymentDeploymentRunner)
	UnlockDeploymentCmd := resCmd.Command("UnlockDeployment", `Unlocks a given deployment. Idempotent.`)
	UnlockDeploymentRunner.Flag(`id`, ``).Required().StringVar(&UnlockDeploymentRunner.id)
	registry[UnlockDeploymentCmd.FullCommand()] = UnlockDeploymentRunner

	UpdateDeploymentRunner := new(UpdateDeploymentDeploymentRunner)
	UpdateDeploymentCmd := resCmd.Command("UpdateDeployment", `Updates attributes of a given deployment.`)
	UpdateDeploymentRunner.Flag(`deployment.description`, `The updated description for the deployment.`).StringVar(&UpdateDeploymentRunner.deploymentDescription)
	UpdateDeploymentRunner.Flag(`deployment.name`, `The updated name for the deployment.`).StringVar(&UpdateDeploymentRunner.deploymentName)
	UpdateDeploymentRunner.Flag(`deployment.serverTagScope`, `The routing scope for tags for servers in the deployment.`).StringVar(&UpdateDeploymentRunner.deploymentServerTagScope)
	UpdateDeploymentRunner.Flag(`id`, ``).Required().StringVar(&UpdateDeploymentRunner.id)
	registry[UpdateDeploymentCmd.FullCommand()] = UpdateDeploymentRunner
}

/****** HealthCheck ******/

type IndexHealthCheckHealthCheckRunner struct {
}

func (r *IndexHealthCheckHealthCheckRunner) Run(c *Client) (interface{}, error) {
	return c.IndexHealthCheck()
}

// Register all HealthCheck actions
func registerHealthCheckCmds(app *kingpin.Application) {
	resCmd := app.Cmd("HealthCheck", ``)

	IndexHealthCheckRunner := new(IndexHealthCheckHealthCheckRunner)
	IndexHealthCheckCmd := resCmd.Command("IndexHealthCheck", `Check health of RightApi controllers`)
	registry[IndexHealthCheckCmd.FullCommand()] = IndexHealthCheckRunner
}

/****** IdentityProvider ******/

type IndexIdentityProvidersIdentityProviderRunner struct {
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexIdentityProvidersIdentityProviderRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexIdentityProviders(filter, r.view)
}

type ShowIdentityProviderIdentityProviderRunner struct {
	id   string
	view string
}

func (r *ShowIdentityProviderIdentityProviderRunner) Run(c *Client) (interface{}, error) {
	return c.ShowIdentityProvider(r.id, r.view)
}

// Register all IdentityProvider actions
func registerIdentityProviderCmds(app *kingpin.Application) {
	resCmd := app.Cmd("IdentityProvider", `An Identity Provider represents a SAML identity provider (IdP) that is linked to your RightScale Enterprise account, and is trusted by the RightScale dashboard to authenticate your enterprise's end users...`)

	IndexIdentityProvidersRunner := new(IndexIdentityProvidersIdentityProviderRunner)
	IndexIdentityProvidersCmd := resCmd.Command("IndexIdentityProviders", `Lists the identity providers associated with this enterprise account.`)
	IndexIdentityProvidersRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexIdentityProvidersRunner.filterPos).StringsVar(&IndexIdentityProvidersRunner.filter)
	IndexIdentityProvidersRunner.Flag(`view`, ``).StringVar(&IndexIdentityProvidersRunner.view)
	registry[IndexIdentityProvidersCmd.FullCommand()] = IndexIdentityProvidersRunner

	ShowIdentityProviderRunner := new(ShowIdentityProviderIdentityProviderRunner)
	ShowIdentityProviderCmd := resCmd.Command("ShowIdentityProvider", `Show the specified identity provider, if associated with this enterprise account.`)
	ShowIdentityProviderRunner.Flag(`id`, ``).Required().StringVar(&ShowIdentityProviderRunner.id)
	ShowIdentityProviderRunner.Flag(`view`, ``).StringVar(&ShowIdentityProviderRunner.view)
	registry[ShowIdentityProviderCmd.FullCommand()] = ShowIdentityProviderRunner
}

/****** Image ******/

type IndexImagesImageRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexImagesImageRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexImages(r.cloudId, filter, r.view)
}

type ShowImageImageRunner struct {
	cloudId string
	id      string
	view    string
}

func (r *ShowImageImageRunner) Run(c *Client) (interface{}, error) {
	return c.ShowImage(r.cloudId, r.id, r.view)
}

// Register all Image actions
func registerImageCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Image", `Images represent base VM image existing in a cloud`)

	IndexImagesRunner := new(IndexImagesImageRunner)
	IndexImagesCmd := resCmd.Command("IndexImages", `Lists all Images for the given Cloud.`)
	IndexImagesRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexImagesRunner.cloudId)
	IndexImagesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexImagesRunner.filterPos).StringsVar(&IndexImagesRunner.filter)
	IndexImagesRunner.Flag(`view`, ``).StringVar(&IndexImagesRunner.view)
	registry[IndexImagesCmd.FullCommand()] = IndexImagesRunner

	ShowImageRunner := new(ShowImageImageRunner)
	ShowImageCmd := resCmd.Command("ShowImage", `Shows information about a single Image.`)
	ShowImageRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowImageRunner.cloudId)
	ShowImageRunner.Flag(`id`, ``).Required().StringVar(&ShowImageRunner.id)
	ShowImageRunner.Flag(`view`, ``).StringVar(&ShowImageRunner.view)
	registry[ShowImageCmd.FullCommand()] = ShowImageRunner
}

/****** Input ******/

type IndexInputsInputRunner struct {
	cloudId    string
	instanceId string
	view       string
}

func (r *IndexInputsInputRunner) Run(c *Client) (interface{}, error) {
	return c.IndexInputs(r.cloudId, r.instanceId, r.view)
}

type MultiUpdateInputsInputRunner struct {
	cloudId      string
	inputsValues []string
	inputsNames  []string
	instanceId   string
}

func (r *MultiUpdateInputsInputRunner) Run(c *Client) (interface{}, error) {

	/** Handle inputs parameter **/
	var inputs map[string]string

	for i, n := range r.inputsNames {
		inputs[n] = r.inputsValues[i]
	}

	return c.MultiUpdateInputs(r.cloudId, inputs, r.instanceId)
}

// Register all Input actions
func registerInputCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Input", `Inputs help extract dynamic information, usually specified at runtime, from repeatable configuration operations that can be codified.`)

	IndexInputsRunner := new(IndexInputsInputRunner)
	IndexInputsCmd := resCmd.Command("IndexInputs", `Retrieves the full list of existing inputs of the specified resource.`)
	IndexInputsRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexInputsRunner.cloudId)
	IndexInputsRunner.Flag(`instanceId`, ``).Required().StringVar(&IndexInputsRunner.instanceId)
	IndexInputsRunner.Flag(`view`, ``).StringVar(&IndexInputsRunner.view)
	registry[IndexInputsCmd.FullCommand()] = IndexInputsRunner

	MultiUpdateInputsRunner := new(MultiUpdateInputsInputRunner)
	MultiUpdateInputsCmd := resCmd.Command("MultiUpdateInputs", `Performs a bulk update of inputs on the specified resource.`)
	MultiUpdateInputsRunner.Flag(`cloudId`, ``).Required().StringVar(&MultiUpdateInputsRunner.cloudId)
	MultiUpdateInputsRunner.FlagPattern(`inputs\.([a-z0-9_]+)`, ``).Required().Capture(&MultiUpdateInputsRunner.inputsNames).StringVar(&MultiUpdateInputsRunner.inputsValues)
	MultiUpdateInputsRunner.Flag(`instanceId`, ``).Required().StringVar(&MultiUpdateInputsRunner.instanceId)
	registry[MultiUpdateInputsCmd.FullCommand()] = MultiUpdateInputsRunner
}

/****** Instance ******/

type CreateInstanceInstanceRunner struct {
	cloudId                                                      string
	instanceAssociatePublicIpAddress                             string
	instanceCloudSpecificAttributesAutomaticInstanceStoreMapping string
	instanceCloudSpecificAttributesEbsOptimized                  string
	instanceCloudSpecificAttributesIamInstanceProfile            string
	instanceCloudSpecificAttributesRootVolumePerformance         string
	instanceCloudSpecificAttributesRootVolumeSize                string
	instanceCloudSpecificAttributesRootVolumeTypeUid             string
	instanceDatacenterHref                                       string
	instanceDeploymentHref                                       string
	instanceImageHref                                            string
	instanceInstanceTypeHref                                     string
	instanceKernelImageHref                                      string
	instanceName                                                 string
	instancePlacementGroupHref                                   string
	instanceRamdiskImageHref                                     string
	instanceItem                                                 []string
	instanceItemPos                                              []string
	instanceSshKeyHref                                           string
	instanceItem                                                 []string
	instanceItemPos                                              []string
	instanceUserData                                             string
}

func (r *CreateInstanceInstanceRunner) Run(c *Client) (interface{}, error) {

	/** Handle instance parameter **/
	var instance InstanceParam

	// Load JSON if provided
	if len(r.instanceJson) > 0 {
		if err := Json.Unmarshal(r.instanceJson, &instance); err != nil {
			return nil, fmt.Errorf("Could not load instance JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.instanceassociatePublicIpAddress) > 0 {
		instance.instance.associatePublicIpAddress = r.instanceassociatePublicIpAddress
	}

	if len(r.instancecloudSpecificAttributesautomaticInstanceStoreMapping) > 0 {
		instance.instance.cloudSpecificAttributes.automaticInstanceStoreMapping = r.instancecloudSpecificAttributesautomaticInstanceStoreMapping
	}

	if len(r.instancecloudSpecificAttributesebsOptimized) > 0 {
		instance.instance.cloudSpecificAttributes.ebsOptimized = r.instancecloudSpecificAttributesebsOptimized
	}

	if len(r.instancecloudSpecificAttributesiamInstanceProfile) > 0 {
		instance.instance.cloudSpecificAttributes.iamInstanceProfile = r.instancecloudSpecificAttributesiamInstanceProfile
	}

	if len(r.instancecloudSpecificAttributesrootVolumePerformance) > 0 {
		instance.instance.cloudSpecificAttributes.rootVolumePerformance = r.instancecloudSpecificAttributesrootVolumePerformance
	}

	if len(r.instancecloudSpecificAttributesrootVolumeSize) > 0 {
		instance.instance.cloudSpecificAttributes.rootVolumeSize = r.instancecloudSpecificAttributesrootVolumeSize
	}

	if len(r.instancecloudSpecificAttributesrootVolumeTypeUid) > 0 {
		instance.instance.cloudSpecificAttributes.rootVolumeTypeUid = r.instancecloudSpecificAttributesrootVolumeTypeUid
	}

	if len(r.instancedatacenterHref) > 0 {
		instance.instance.datacenterHref = r.instancedatacenterHref
	}

	if len(r.instancedeploymentHref) > 0 {
		instance.instance.deploymentHref = r.instancedeploymentHref
	}

	if len(r.instanceimageHref) > 0 {
		instance.instance.imageHref = r.instanceimageHref
	}

	if len(r.instanceinstanceTypeHref) > 0 {
		instance.instance.instanceTypeHref = r.instanceinstanceTypeHref
	}

	if len(r.instancekernelImageHref) > 0 {
		instance.instance.kernelImageHref = r.instancekernelImageHref
	}

	if len(r.instancename) > 0 {
		instance.instance.name = r.instancename
	}

	if len(r.instanceplacementGroupHref) > 0 {
		instance.instance.placementGroupHref = r.instanceplacementGroupHref
	}

	if len(r.instanceramdiskImageHref) > 0 {
		instance.instance.ramdiskImageHref = r.instanceramdiskImageHref
	}

	if len(r.instancesecurityGroupHrefsitem) > 0 {
		instance.instance.securityGroupHrefs.item = r.instancesecurityGroupHrefsitem
	}

	if len(r.instancesshKeyHref) > 0 {
		instance.instance.sshKeyHref = r.instancesshKeyHref
	}

	if len(r.instancesubnetHrefsitem) > 0 {
		instance.instance.subnetHrefs.item = r.instancesubnetHrefsitem
	}

	if len(r.instanceuserData) > 0 {
		instance.instance.userData = r.instanceuserData
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxinstancesecurityGroupHrefsitemPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.instancesecurityGroupHrefsitemPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for instance.securityGroupHrefs.item field of securityGroupHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of instance.securityGroupHrefs.item field of securityGroupHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxinstancesecurityGroupHrefsitemPos {
			maxinstancesecurityGroupHrefsitemPos = pos
		}
	}
	if len(r.instancesecurityGroupHrefsitem) != maxinstancesecurityGroupHrefsitemPos {
		return nil, fmt.Errof("Invalid flags for instance.securityGroupHrefs.item field of securityGroupHrefs array, %s were defined but max position is %s",
			len(r.instancesecurityGroupHrefsitem), maxinstancesecurityGroupHrefsitemPos)
	}
	if maxinstancesecurityGroupHrefsitemPos > -1 {
		instancesecurityGroupHrefs := make([]*string, maxinstancesecurityGroupHrefsitemPos+1)
		for i := 0; i < maxinstancesecurityGroupHrefsitemPos+1; i++ {
			instancesecurityGroupHrefs[i] = string{
			//TBD
			}
		}
		instance.securityGroupHrefs = instancesecurityGroupHrefs
	}

	maxinstancesubnetHrefsitemPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.instancesubnetHrefsitemPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for instance.subnetHrefs.item field of subnetHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of instance.subnetHrefs.item field of subnetHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxinstancesubnetHrefsitemPos {
			maxinstancesubnetHrefsitemPos = pos
		}
	}
	if len(r.instancesubnetHrefsitem) != maxinstancesubnetHrefsitemPos {
		return nil, fmt.Errof("Invalid flags for instance.subnetHrefs.item field of subnetHrefs array, %s were defined but max position is %s",
			len(r.instancesubnetHrefsitem), maxinstancesubnetHrefsitemPos)
	}
	if maxinstancesubnetHrefsitemPos > -1 {
		instancesubnetHrefs := make([]*string, maxinstancesubnetHrefsitemPos+1)
		for i := 0; i < maxinstancesubnetHrefsitemPos+1; i++ {
			instancesubnetHrefs[i] = string{
			//TBD
			}
		}
		instance.subnetHrefs = instancesubnetHrefs
	}

	return c.CreateInstance(r.cloudId, &instance)
}

type IndexInstancesInstanceRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexInstancesInstanceRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexInstances(r.cloudId, filter, r.view)
}

type LaunchInstanceInstanceRunner struct {
	apiBehavior  string
	cloudId      string
	id           string
	inputsValues []string
	inputsNames  []string
}

func (r *LaunchInstanceInstanceRunner) Run(c *Client) (interface{}, error) {

	/** Handle inputs parameter **/
	var inputs map[string]string

	for i, n := range r.inputsNames {
		inputs[n] = r.inputsValues[i]
	}

	return c.LaunchInstance(r.apiBehavior, r.cloudId, r.id, inputs)
}

type LockInstanceInstanceRunner struct {
	cloudId string
	id      string
}

func (r *LockInstanceInstanceRunner) Run(c *Client) (interface{}, error) {
	return c.LockInstance(r.cloudId, r.id)
}

type MultiRunExecutableInstancesInstanceRunner struct {
	cloudId         string
	filter          []string
	filterPos       []string
	ignoreLock      string
	inputsValues    []string
	inputsNames     []string
	recipeName      string
	rightScriptHref string
}

func (r *MultiRunExecutableInstancesInstanceRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle inputs parameter **/
	var inputs map[string]string

	for i, n := range r.inputsNames {
		inputs[n] = r.inputsValues[i]
	}

	return c.MultiRunExecutableInstances(r.cloudId, filter, r.ignoreLock, inputs, r.recipeName, r.rightScriptHref)
}

type MultiTerminateInstancesInstanceRunner struct {
	cloudId      string
	filter       []string
	filterPos    []string
	terminateAll string
}

func (r *MultiTerminateInstancesInstanceRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.MultiTerminateInstances(r.cloudId, filter, r.terminateAll)
}

type RebootInstanceInstanceRunner struct {
	cloudId string
	id      string
}

func (r *RebootInstanceInstanceRunner) Run(c *Client) (interface{}, error) {
	return c.RebootInstance(r.cloudId, r.id)
}

type RunExecutableInstanceInstanceRunner struct {
	cloudId         string
	id              string
	ignoreLock      string
	inputsValues    []string
	inputsNames     []string
	recipeName      string
	rightScriptHref string
}

func (r *RunExecutableInstanceInstanceRunner) Run(c *Client) (interface{}, error) {

	/** Handle inputs parameter **/
	var inputs map[string]string

	for i, n := range r.inputsNames {
		inputs[n] = r.inputsValues[i]
	}

	return c.RunExecutableInstance(r.cloudId, r.id, r.ignoreLock, inputs, r.recipeName, r.rightScriptHref)
}

type SetCustomLodgementInstanceInstanceRunner struct {
	cloudId          string
	id               string
	quantityName     []string
	quantityNamePos  []string
	quantityValue    []string
	quantityValuePos []string
	timeframe        string
}

func (r *SetCustomLodgementInstanceInstanceRunner) Run(c *Client) (interface{}, error) {

	/** Handle quantity parameter **/
	var quantity []*Quantity

	for i, v := range r.quantity {
		pos, err := strconv.Atoi(r.quantityPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for quantity array", r.quantityPos[i])
		}
		quantity[pos] = v
	}

	return c.SetCustomLodgementInstance(r.cloudId, r.id, quantity, r.timeframe)
}

type ShowInstanceInstanceRunner struct {
	cloudId string
	id      string
	view    string
}

func (r *ShowInstanceInstanceRunner) Run(c *Client) (interface{}, error) {
	return c.ShowInstance(r.cloudId, r.id, r.view)
}

type StartInstanceInstanceRunner struct {
	cloudId string
	id      string
}

func (r *StartInstanceInstanceRunner) Run(c *Client) (interface{}, error) {
	return c.StartInstance(r.cloudId, r.id)
}

type StopInstanceInstanceRunner struct {
	cloudId string
	id      string
}

func (r *StopInstanceInstanceRunner) Run(c *Client) (interface{}, error) {
	return c.StopInstance(r.cloudId, r.id)
}

type TerminateInstanceInstanceRunner struct {
	cloudId string
	id      string
}

func (r *TerminateInstanceInstanceRunner) Run(c *Client) (interface{}, error) {
	return c.TerminateInstance(r.cloudId, r.id)
}

type UnlockInstanceInstanceRunner struct {
	cloudId string
	id      string
}

func (r *UnlockInstanceInstanceRunner) Run(c *Client) (interface{}, error) {
	return c.UnlockInstance(r.cloudId, r.id)
}

type UpdateInstanceInstanceRunner struct {
	cloudId                                                      string
	id                                                           string
	instanceAssociatePublicIpAddress                             string
	instanceCloudSpecificAttributesAutomaticInstanceStoreMapping string
	instanceCloudSpecificAttributesIamInstanceProfile            string
	instanceCloudSpecificAttributesRootVolumePerformance         string
	instanceCloudSpecificAttributesRootVolumeSize                string
	instanceCloudSpecificAttributesRootVolumeTypeUid             string
	instanceDatacenterHref                                       string
	instanceDeploymentHref                                       string
	instanceImageHref                                            string
	instanceInstanceTypeHref                                     string
	instanceIpForwardingEnabled                                  string
	instanceKernelImageHref                                      string
	instanceMultiCloudImageHref                                  string
	instanceName                                                 string
	instanceRamdiskImageHref                                     string
	instanceItem                                                 []string
	instanceItemPos                                              []string
	instanceServerTemplateHref                                   string
	instanceSshKeyHref                                           string
	instanceItem                                                 []string
	instanceItemPos                                              []string
	instanceUserData                                             string
}

func (r *UpdateInstanceInstanceRunner) Run(c *Client) (interface{}, error) {

	/** Handle instance parameter **/
	var instance InstanceParam2

	// Load JSON if provided
	if len(r.instanceJson) > 0 {
		if err := Json.Unmarshal(r.instanceJson, &instance); err != nil {
			return nil, fmt.Errorf("Could not load instance JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.instanceassociatePublicIpAddress) > 0 {
		instance.instance.associatePublicIpAddress = r.instanceassociatePublicIpAddress
	}

	if len(r.instancecloudSpecificAttributesautomaticInstanceStoreMapping) > 0 {
		instance.instance.cloudSpecificAttributes.automaticInstanceStoreMapping = r.instancecloudSpecificAttributesautomaticInstanceStoreMapping
	}

	if len(r.instancecloudSpecificAttributesiamInstanceProfile) > 0 {
		instance.instance.cloudSpecificAttributes.iamInstanceProfile = r.instancecloudSpecificAttributesiamInstanceProfile
	}

	if len(r.instancecloudSpecificAttributesrootVolumePerformance) > 0 {
		instance.instance.cloudSpecificAttributes.rootVolumePerformance = r.instancecloudSpecificAttributesrootVolumePerformance
	}

	if len(r.instancecloudSpecificAttributesrootVolumeSize) > 0 {
		instance.instance.cloudSpecificAttributes.rootVolumeSize = r.instancecloudSpecificAttributesrootVolumeSize
	}

	if len(r.instancecloudSpecificAttributesrootVolumeTypeUid) > 0 {
		instance.instance.cloudSpecificAttributes.rootVolumeTypeUid = r.instancecloudSpecificAttributesrootVolumeTypeUid
	}

	if len(r.instancedatacenterHref) > 0 {
		instance.instance.datacenterHref = r.instancedatacenterHref
	}

	if len(r.instancedeploymentHref) > 0 {
		instance.instance.deploymentHref = r.instancedeploymentHref
	}

	if len(r.instanceimageHref) > 0 {
		instance.instance.imageHref = r.instanceimageHref
	}

	if len(r.instanceinstanceTypeHref) > 0 {
		instance.instance.instanceTypeHref = r.instanceinstanceTypeHref
	}

	if len(r.instanceipForwardingEnabled) > 0 {
		instance.instance.ipForwardingEnabled = r.instanceipForwardingEnabled
	}

	if len(r.instancekernelImageHref) > 0 {
		instance.instance.kernelImageHref = r.instancekernelImageHref
	}

	if len(r.instancemultiCloudImageHref) > 0 {
		instance.instance.multiCloudImageHref = r.instancemultiCloudImageHref
	}

	if len(r.instancename) > 0 {
		instance.instance.name = r.instancename
	}

	if len(r.instanceramdiskImageHref) > 0 {
		instance.instance.ramdiskImageHref = r.instanceramdiskImageHref
	}

	if len(r.instancesecurityGroupHrefsitem) > 0 {
		instance.instance.securityGroupHrefs.item = r.instancesecurityGroupHrefsitem
	}

	if len(r.instanceserverTemplateHref) > 0 {
		instance.instance.serverTemplateHref = r.instanceserverTemplateHref
	}

	if len(r.instancesshKeyHref) > 0 {
		instance.instance.sshKeyHref = r.instancesshKeyHref
	}

	if len(r.instancesubnetHrefsitem) > 0 {
		instance.instance.subnetHrefs.item = r.instancesubnetHrefsitem
	}

	if len(r.instanceuserData) > 0 {
		instance.instance.userData = r.instanceuserData
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxinstancesecurityGroupHrefsitemPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.instancesecurityGroupHrefsitemPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for instance.securityGroupHrefs.item field of securityGroupHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of instance.securityGroupHrefs.item field of securityGroupHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxinstancesecurityGroupHrefsitemPos {
			maxinstancesecurityGroupHrefsitemPos = pos
		}
	}
	if len(r.instancesecurityGroupHrefsitem) != maxinstancesecurityGroupHrefsitemPos {
		return nil, fmt.Errof("Invalid flags for instance.securityGroupHrefs.item field of securityGroupHrefs array, %s were defined but max position is %s",
			len(r.instancesecurityGroupHrefsitem), maxinstancesecurityGroupHrefsitemPos)
	}
	if maxinstancesecurityGroupHrefsitemPos > -1 {
		instancesecurityGroupHrefs := make([]*string, maxinstancesecurityGroupHrefsitemPos+1)
		for i := 0; i < maxinstancesecurityGroupHrefsitemPos+1; i++ {
			instancesecurityGroupHrefs[i] = string{
			//TBD
			}
		}
		instance.securityGroupHrefs = instancesecurityGroupHrefs
	}

	maxinstancesubnetHrefsitemPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.instancesubnetHrefsitemPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for instance.subnetHrefs.item field of subnetHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of instance.subnetHrefs.item field of subnetHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxinstancesubnetHrefsitemPos {
			maxinstancesubnetHrefsitemPos = pos
		}
	}
	if len(r.instancesubnetHrefsitem) != maxinstancesubnetHrefsitemPos {
		return nil, fmt.Errof("Invalid flags for instance.subnetHrefs.item field of subnetHrefs array, %s were defined but max position is %s",
			len(r.instancesubnetHrefsitem), maxinstancesubnetHrefsitemPos)
	}
	if maxinstancesubnetHrefsitemPos > -1 {
		instancesubnetHrefs := make([]*string, maxinstancesubnetHrefsitemPos+1)
		for i := 0; i < maxinstancesubnetHrefsitemPos+1; i++ {
			instancesubnetHrefs[i] = string{
			//TBD
			}
		}
		instance.subnetHrefs = instancesubnetHrefs
	}

	return c.UpdateInstance(r.cloudId, r.id, &instance)
}

// Register all Instance actions
func registerInstanceCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Instance", `Instances represent an entity that is runnable in the cloud.`)

	CreateInstanceRunner := new(CreateInstanceInstanceRunner)
	CreateInstanceCmd := resCmd.Command("CreateInstance", `Creates and launches a raw instance using the provided parameters.`)
	CreateInstanceRunner.Flag(`cloudId`, ``).Required().StringVar(&CreateInstanceRunner.cloudId)
	CreateInstanceRunner.Flag(`instance.associatePublicIpAddress`, `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`).StringVar(&CreateInstanceRunner.instanceAssociatePublicIpAddress)
	CreateInstanceRunner.Flag(`instance.cloudSpecificAttributes.automaticInstanceStoreMapping`, `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`).StringVar(&CreateInstanceRunner.instanceCloudSpecificAttributesAutomaticInstanceStoreMapping)
	CreateInstanceRunner.Flag(`instance.cloudSpecificAttributes.ebsOptimized`, `Whether the instance is able to connect to IOPS-enabled volumes.`).StringVar(&CreateInstanceRunner.instanceCloudSpecificAttributesEbsOptimized)
	CreateInstanceRunner.Flag(`instance.cloudSpecificAttributes.iamInstanceProfile`, `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`).StringVar(&CreateInstanceRunner.instanceCloudSpecificAttributesIamInstanceProfile)
	CreateInstanceRunner.Flag(`instance.cloudSpecificAttributes.rootVolumePerformance`, `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`).StringVar(&CreateInstanceRunner.instanceCloudSpecificAttributesRootVolumePerformance)
	CreateInstanceRunner.Flag(`instance.cloudSpecificAttributes.rootVolumeSize`, `The size for root disk. Not supported in all Clouds.`).StringVar(&CreateInstanceRunner.instanceCloudSpecificAttributesRootVolumeSize)
	CreateInstanceRunner.Flag(`instance.cloudSpecificAttributes.rootVolumeTypeUid`, `The type of root volume for instance. Only available on clouds supporting root volume type.`).StringVar(&CreateInstanceRunner.instanceCloudSpecificAttributesRootVolumeTypeUid)
	CreateInstanceRunner.Flag(`instance.datacenterHref`, `The href of the Datacenter / Zone.`).StringVar(&CreateInstanceRunner.instanceDatacenterHref)
	CreateInstanceRunner.Flag(`instance.deploymentHref`, `The href of the deployment to which the Instance will be added.`).StringVar(&CreateInstanceRunner.instanceDeploymentHref)
	CreateInstanceRunner.Flag(`instance.imageHref`, `The href of the Image to use.`).Required().StringVar(&CreateInstanceRunner.instanceImageHref)
	CreateInstanceRunner.Flag(`instance.instanceTypeHref`, `The href of the instance type.`).Required().StringVar(&CreateInstanceRunner.instanceInstanceTypeHref)
	CreateInstanceRunner.Flag(`instance.kernelImageHref`, `The href of the kernel image.`).StringVar(&CreateInstanceRunner.instanceKernelImageHref)
	CreateInstanceRunner.Flag(`instance.name`, `The name of the instance.`).Required().StringVar(&CreateInstanceRunner.instanceName)
	CreateInstanceRunner.Flag(`instance.placementGroupHref`, `The placement group to launch the instance in. Not supported by all clouds & instance types.`).StringVar(&CreateInstanceRunner.instancePlacementGroupHref)
	CreateInstanceRunner.Flag(`instance.ramdiskImageHref`, `The href of the ramdisk image.`).StringVar(&CreateInstanceRunner.instanceRamdiskImageHref)
	CreateInstanceRunner.FlagPattern(`instance\.item\.(\d+)`, `The hrefs of the security groups.`).Capture(&CreateInstanceRunner.instanceItemPos).StringsVar(&CreateInstanceRunner.instanceItem)
	CreateInstanceRunner.Flag(`instance.sshKeyHref`, `The href of the SSH key to use.`).StringVar(&CreateInstanceRunner.instanceSshKeyHref)
	CreateInstanceRunner.FlagPattern(`instance\.item\.(\d+)`, `The hrefs of the updated subnets.`).Capture(&CreateInstanceRunner.instanceItemPos).StringsVar(&CreateInstanceRunner.instanceItem)
	CreateInstanceRunner.Flag(`instance.userData`, `User data that RightScale automatically passes to your instance at boot time.`).StringVar(&CreateInstanceRunner.instanceUserData)
	registry[CreateInstanceCmd.FullCommand()] = CreateInstanceRunner

	IndexInstancesRunner := new(IndexInstancesInstanceRunner)
	IndexInstancesCmd := resCmd.Command("IndexInstances", `Lists instances of a given cloud, server array.`)
	IndexInstancesRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexInstancesRunner.cloudId)
	IndexInstancesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexInstancesRunner.filterPos).StringsVar(&IndexInstancesRunner.filter)
	IndexInstancesRunner.Flag(`view`, ``).StringVar(&IndexInstancesRunner.view)
	registry[IndexInstancesCmd.FullCommand()] = IndexInstancesRunner

	LaunchInstanceRunner := new(LaunchInstanceInstanceRunner)
	LaunchInstanceCmd := resCmd.Command("LaunchInstance", `Launches an instance using the parameters that this instance has been configured with.`)
	LaunchInstanceRunner.Flag(`apiBehavior`, `When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'`).StringVar(&LaunchInstanceRunner.apiBehavior)
	LaunchInstanceRunner.Flag(`cloudId`, ``).Required().StringVar(&LaunchInstanceRunner.cloudId)
	LaunchInstanceRunner.Flag(`id`, ``).Required().StringVar(&LaunchInstanceRunner.id)
	LaunchInstanceRunner.FlagPattern(`inputs\.([a-z0-9_]+)`, ``).Capture(&LaunchInstanceRunner.inputsNames).StringVar(&LaunchInstanceRunner.inputsValues)
	registry[LaunchInstanceCmd.FullCommand()] = LaunchInstanceRunner

	LockInstanceRunner := new(LockInstanceInstanceRunner)
	LockInstanceCmd := resCmd.Command("LockInstance", ``)
	LockInstanceRunner.Flag(`cloudId`, ``).Required().StringVar(&LockInstanceRunner.cloudId)
	LockInstanceRunner.Flag(`id`, ``).Required().StringVar(&LockInstanceRunner.id)
	registry[LockInstanceCmd.FullCommand()] = LockInstanceRunner

	MultiRunExecutableInstancesRunner := new(MultiRunExecutableInstancesInstanceRunner)
	MultiRunExecutableInstancesCmd := resCmd.Command("MultiRunExecutableInstances", `Runs a script or a recipe in the running instances.`)
	MultiRunExecutableInstancesRunner.Flag(`cloudId`, ``).Required().StringVar(&MultiRunExecutableInstancesRunner.cloudId)
	MultiRunExecutableInstancesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&MultiRunExecutableInstancesRunner.filterPos).StringsVar(&MultiRunExecutableInstancesRunner.filter)
	MultiRunExecutableInstancesRunner.Flag(`ignoreLock`, `Specifies the ability to ignore the lock(s) on the Instance(s).`).StringVar(&MultiRunExecutableInstancesRunner.ignoreLock)
	MultiRunExecutableInstancesRunner.FlagPattern(`inputs\.([a-z0-9_]+)`, ``).Capture(&MultiRunExecutableInstancesRunner.inputsNames).StringVar(&MultiRunExecutableInstancesRunner.inputsValues)
	MultiRunExecutableInstancesRunner.Flag(`recipeName`, `The name of the recipe to be run.`).StringVar(&MultiRunExecutableInstancesRunner.recipeName)
	MultiRunExecutableInstancesRunner.Flag(`rightScriptHref`, `The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.`).StringVar(&MultiRunExecutableInstancesRunner.rightScriptHref)
	registry[MultiRunExecutableInstancesCmd.FullCommand()] = MultiRunExecutableInstancesRunner

	MultiTerminateInstancesRunner := new(MultiTerminateInstancesInstanceRunner)
	MultiTerminateInstancesCmd := resCmd.Command("MultiTerminateInstances", `Terminates running instances.`)
	MultiTerminateInstancesRunner.Flag(`cloudId`, ``).Required().StringVar(&MultiTerminateInstancesRunner.cloudId)
	MultiTerminateInstancesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&MultiTerminateInstancesRunner.filterPos).StringsVar(&MultiTerminateInstancesRunner.filter)
	MultiTerminateInstancesRunner.Flag(`terminateAll`, `Specifies the ability to terminate all instances.`).StringVar(&MultiTerminateInstancesRunner.terminateAll)
	registry[MultiTerminateInstancesCmd.FullCommand()] = MultiTerminateInstancesRunner

	RebootInstanceRunner := new(RebootInstanceInstanceRunner)
	RebootInstanceCmd := resCmd.Command("RebootInstance", `Reboot a running instance.`)
	RebootInstanceRunner.Flag(`cloudId`, ``).Required().StringVar(&RebootInstanceRunner.cloudId)
	RebootInstanceRunner.Flag(`id`, ``).Required().StringVar(&RebootInstanceRunner.id)
	registry[RebootInstanceCmd.FullCommand()] = RebootInstanceRunner

	RunExecutableInstanceRunner := new(RunExecutableInstanceInstanceRunner)
	RunExecutableInstanceCmd := resCmd.Command("RunExecutableInstance", `Runs a script or a recipe in the running instance.`)
	RunExecutableInstanceRunner.Flag(`cloudId`, ``).Required().StringVar(&RunExecutableInstanceRunner.cloudId)
	RunExecutableInstanceRunner.Flag(`id`, ``).Required().StringVar(&RunExecutableInstanceRunner.id)
	RunExecutableInstanceRunner.Flag(`ignoreLock`, `Specifies the ability to ignore the lock on the Instance.`).StringVar(&RunExecutableInstanceRunner.ignoreLock)
	RunExecutableInstanceRunner.FlagPattern(`inputs\.([a-z0-9_]+)`, ``).Capture(&RunExecutableInstanceRunner.inputsNames).StringVar(&RunExecutableInstanceRunner.inputsValues)
	RunExecutableInstanceRunner.Flag(`recipeName`, `The name of the recipe to run.`).StringVar(&RunExecutableInstanceRunner.recipeName)
	RunExecutableInstanceRunner.Flag(`rightScriptHref`, `The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.`).StringVar(&RunExecutableInstanceRunner.rightScriptHref)
	registry[RunExecutableInstanceCmd.FullCommand()] = RunExecutableInstanceRunner

	SetCustomLodgementInstanceRunner := new(SetCustomLodgementInstanceInstanceRunner)
	SetCustomLodgementInstanceCmd := resCmd.Command("SetCustomLodgementInstance", `This method is deprecated.  Please use InstanceCustomLodgement.`)
	SetCustomLodgementInstanceRunner.Flag(`cloudId`, ``).Required().StringVar(&SetCustomLodgementInstanceRunner.cloudId)
	SetCustomLodgementInstanceRunner.Flag(`id`, ``).Required().StringVar(&SetCustomLodgementInstanceRunner.id)
	SetCustomLodgementInstanceRunner.FlagPattern(`quantity\.(\d+)\.name`, `The name of the quantity. A customer-specific string, e.g. "MB/s" or "GB/Month".`).Capture(&SetCustomLodgementInstanceRunner.quantityNamePos).StringsVar(&SetCustomLodgementInstanceRunner.quantityName)
	SetCustomLodgementInstanceRunner.FlagPattern(`quantity\.(\d+)\.value`, `The value of the quantity. Should be a positive integer.`).Capture(&SetCustomLodgementInstanceRunner.quantityValuePos).StringsVar(&SetCustomLodgementInstanceRunner.quantityValue)
	SetCustomLodgementInstanceRunner.Flag(`timeframe`, `The timeframe (either a month or a single day) for which the quantity value is valid (currently for the PDT timezone only).`).Required().StringVar(&SetCustomLodgementInstanceRunner.timeframe)
	registry[SetCustomLodgementInstanceCmd.FullCommand()] = SetCustomLodgementInstanceRunner

	ShowInstanceRunner := new(ShowInstanceInstanceRunner)
	ShowInstanceCmd := resCmd.Command("ShowInstance", `Shows attributes of a single instance.`)
	ShowInstanceRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowInstanceRunner.cloudId)
	ShowInstanceRunner.Flag(`id`, ``).Required().StringVar(&ShowInstanceRunner.id)
	ShowInstanceRunner.Flag(`view`, ``).StringVar(&ShowInstanceRunner.view)
	registry[ShowInstanceCmd.FullCommand()] = ShowInstanceRunner

	StartInstanceRunner := new(StartInstanceInstanceRunner)
	StartInstanceCmd := resCmd.Command("StartInstance", `Starts an instance that has been stopped, resuming it to its previously saved volume state.`)
	StartInstanceRunner.Flag(`cloudId`, ``).Required().StringVar(&StartInstanceRunner.cloudId)
	StartInstanceRunner.Flag(`id`, ``).Required().StringVar(&StartInstanceRunner.id)
	registry[StartInstanceCmd.FullCommand()] = StartInstanceRunner

	StopInstanceRunner := new(StopInstanceInstanceRunner)
	StopInstanceCmd := resCmd.Command("StopInstance", `Stores the instance's current volume state to resume later using the 'start' action.`)
	StopInstanceRunner.Flag(`cloudId`, ``).Required().StringVar(&StopInstanceRunner.cloudId)
	StopInstanceRunner.Flag(`id`, ``).Required().StringVar(&StopInstanceRunner.id)
	registry[StopInstanceCmd.FullCommand()] = StopInstanceRunner

	TerminateInstanceRunner := new(TerminateInstanceInstanceRunner)
	TerminateInstanceCmd := resCmd.Command("TerminateInstance", `Terminates a running instance.`)
	TerminateInstanceRunner.Flag(`cloudId`, ``).Required().StringVar(&TerminateInstanceRunner.cloudId)
	TerminateInstanceRunner.Flag(`id`, ``).Required().StringVar(&TerminateInstanceRunner.id)
	registry[TerminateInstanceCmd.FullCommand()] = TerminateInstanceRunner

	UnlockInstanceRunner := new(UnlockInstanceInstanceRunner)
	UnlockInstanceCmd := resCmd.Command("UnlockInstance", ``)
	UnlockInstanceRunner.Flag(`cloudId`, ``).Required().StringVar(&UnlockInstanceRunner.cloudId)
	UnlockInstanceRunner.Flag(`id`, ``).Required().StringVar(&UnlockInstanceRunner.id)
	registry[UnlockInstanceCmd.FullCommand()] = UnlockInstanceRunner

	UpdateInstanceRunner := new(UpdateInstanceInstanceRunner)
	UpdateInstanceCmd := resCmd.Command("UpdateInstance", `Updates attributes of a single instance.`)
	UpdateInstanceRunner.Flag(`cloudId`, ``).Required().StringVar(&UpdateInstanceRunner.cloudId)
	UpdateInstanceRunner.Flag(`id`, ``).Required().StringVar(&UpdateInstanceRunner.id)
	UpdateInstanceRunner.Flag(`instance.associatePublicIpAddress`, `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`).StringVar(&UpdateInstanceRunner.instanceAssociatePublicIpAddress)
	UpdateInstanceRunner.Flag(`instance.cloudSpecificAttributes.automaticInstanceStoreMapping`, `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`).StringVar(&UpdateInstanceRunner.instanceCloudSpecificAttributesAutomaticInstanceStoreMapping)
	UpdateInstanceRunner.Flag(`instance.cloudSpecificAttributes.iamInstanceProfile`, `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`).StringVar(&UpdateInstanceRunner.instanceCloudSpecificAttributesIamInstanceProfile)
	UpdateInstanceRunner.Flag(`instance.cloudSpecificAttributes.rootVolumePerformance`, `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`).StringVar(&UpdateInstanceRunner.instanceCloudSpecificAttributesRootVolumePerformance)
	UpdateInstanceRunner.Flag(`instance.cloudSpecificAttributes.rootVolumeSize`, `The size for root disk. Not supported in all Clouds.`).StringVar(&UpdateInstanceRunner.instanceCloudSpecificAttributesRootVolumeSize)
	UpdateInstanceRunner.Flag(`instance.cloudSpecificAttributes.rootVolumeTypeUid`, `The type of root volume for instance. Only available on clouds supporting root volume type.`).StringVar(&UpdateInstanceRunner.instanceCloudSpecificAttributesRootVolumeTypeUid)
	UpdateInstanceRunner.Flag(`instance.datacenterHref`, `The href of the updated Datacenter / Zone for the Instance.`).StringVar(&UpdateInstanceRunner.instanceDatacenterHref)
	UpdateInstanceRunner.Flag(`instance.deploymentHref`, `The href of the updated Deployment for the Instance. This is only supported for Instances that are not associated with a Server or ServerArray.`).StringVar(&UpdateInstanceRunner.instanceDeploymentHref)
	UpdateInstanceRunner.Flag(`instance.imageHref`, `The href of the updated Image for the Instance.`).StringVar(&UpdateInstanceRunner.instanceImageHref)
	UpdateInstanceRunner.Flag(`instance.instanceTypeHref`, `The href of the updated Instance Type.`).StringVar(&UpdateInstanceRunner.instanceInstanceTypeHref)
	UpdateInstanceRunner.Flag(`instance.ipForwardingEnabled`, `Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.`).StringVar(&UpdateInstanceRunner.instanceIpForwardingEnabled)
	UpdateInstanceRunner.Flag(`instance.kernelImageHref`, `The href of the updated kernel image for the Instance.`).StringVar(&UpdateInstanceRunner.instanceKernelImageHref)
	UpdateInstanceRunner.Flag(`instance.multiCloudImageHref`, `The href of the updated MultiCloudImage for the Instance.`).StringVar(&UpdateInstanceRunner.instanceMultiCloudImageHref)
	UpdateInstanceRunner.Flag(`instance.name`, `The updated name to give the Instance.`).StringVar(&UpdateInstanceRunner.instanceName)
	UpdateInstanceRunner.Flag(`instance.ramdiskImageHref`, `The href of the updated ramdisk image for the Instance.`).StringVar(&UpdateInstanceRunner.instanceRamdiskImageHref)
	UpdateInstanceRunner.FlagPattern(`instance\.item\.(\d+)`, `The hrefs of the updated security groups.`).Capture(&UpdateInstanceRunner.instanceItemPos).StringsVar(&UpdateInstanceRunner.instanceItem)
	UpdateInstanceRunner.Flag(`instance.serverTemplateHref`, `The href of the updated ServerTemplate for the Instance.`).StringVar(&UpdateInstanceRunner.instanceServerTemplateHref)
	UpdateInstanceRunner.Flag(`instance.sshKeyHref`, `The href of the updated SSH key for the Instance.`).StringVar(&UpdateInstanceRunner.instanceSshKeyHref)
	UpdateInstanceRunner.FlagPattern(`instance\.item\.(\d+)`, `The hrefs of the updated subnets.`).Capture(&UpdateInstanceRunner.instanceItemPos).StringsVar(&UpdateInstanceRunner.instanceItem)
	UpdateInstanceRunner.Flag(`instance.userData`, `User data that RightScale automatically passes to your instance at boot time.`).StringVar(&UpdateInstanceRunner.instanceUserData)
	registry[UpdateInstanceCmd.FullCommand()] = UpdateInstanceRunner
}

/****** InstanceCustomLodgement ******/

type CreateInstanceCustomLodgementInstanceCustomLodgementRunner struct {
	cloudId          string
	instanceId       string
	quantityName     []string
	quantityNamePos  []string
	quantityValue    []string
	quantityValuePos []string
	timeframe        string
}

func (r *CreateInstanceCustomLodgementInstanceCustomLodgementRunner) Run(c *Client) (interface{}, error) {

	/** Handle quantity parameter **/
	var quantity []*Quantity

	for i, v := range r.quantity {
		pos, err := strconv.Atoi(r.quantityPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for quantity array", r.quantityPos[i])
		}
		quantity[pos] = v
	}

	return c.CreateInstanceCustomLodgement(r.cloudId, r.instanceId, quantity, r.timeframe)
}

type DestroyInstanceCustomLodgementInstanceCustomLodgementRunner struct {
	cloudId    string
	id         string
	instanceId string
}

func (r *DestroyInstanceCustomLodgementInstanceCustomLodgementRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyInstanceCustomLodgement(r.cloudId, r.id, r.instanceId)
}

type IndexInstanceCustomLodgementsInstanceCustomLodgementRunner struct {
	cloudId    string
	instanceId string
	view       string
}

func (r *IndexInstanceCustomLodgementsInstanceCustomLodgementRunner) Run(c *Client) (interface{}, error) {
	return c.IndexInstanceCustomLodgements(r.cloudId, r.instanceId, r.view)
}

type ShowInstanceCustomLodgementInstanceCustomLodgementRunner struct {
	cloudId    string
	id         string
	instanceId string
}

func (r *ShowInstanceCustomLodgementInstanceCustomLodgementRunner) Run(c *Client) (interface{}, error) {
	return c.ShowInstanceCustomLodgement(r.cloudId, r.id, r.instanceId)
}

type UpdateInstanceCustomLodgementInstanceCustomLodgementRunner struct {
	cloudId          string
	id               string
	instanceId       string
	quantityName     []string
	quantityNamePos  []string
	quantityValue    []string
	quantityValuePos []string
}

func (r *UpdateInstanceCustomLodgementInstanceCustomLodgementRunner) Run(c *Client) (interface{}, error) {

	/** Handle quantity parameter **/
	var quantity []*Quantity2

	for i, v := range r.quantity {
		pos, err := strconv.Atoi(r.quantityPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for quantity array", r.quantityPos[i])
		}
		quantity[pos] = v
	}

	return c.UpdateInstanceCustomLodgement(r.cloudId, r.id, r.instanceId, quantity)
}

// Register all InstanceCustomLodgement actions
func registerInstanceCustomLodgementCmds(app *kingpin.Application) {
	resCmd := app.Cmd("InstanceCustomLodgement", `An InstanceCustomLodgement represents a way to create custom reports about a specific instance with a user defined quantity.  Replaces the legacy Instances#setcustomlodgement interface.`)

	CreateInstanceCustomLodgementRunner := new(CreateInstanceCustomLodgementInstanceCustomLodgementRunner)
	CreateInstanceCustomLodgementCmd := resCmd.Command("CreateInstanceCustomLodgement", `Create a lodgement with the quantity and timeframe specified.`)
	CreateInstanceCustomLodgementRunner.Flag(`cloudId`, ``).Required().StringVar(&CreateInstanceCustomLodgementRunner.cloudId)
	CreateInstanceCustomLodgementRunner.Flag(`instanceId`, ``).Required().StringVar(&CreateInstanceCustomLodgementRunner.instanceId)
	CreateInstanceCustomLodgementRunner.FlagPattern(`quantity\.(\d+)\.name`, `The name of the quantity. A customer-specific string, e.g. "MB/s" or "GB/Month".`).Capture(&CreateInstanceCustomLodgementRunner.quantityNamePos).StringsVar(&CreateInstanceCustomLodgementRunner.quantityName)
	CreateInstanceCustomLodgementRunner.FlagPattern(`quantity\.(\d+)\.value`, `The value of the quantity. Should be a positive integer.`).Capture(&CreateInstanceCustomLodgementRunner.quantityValuePos).StringsVar(&CreateInstanceCustomLodgementRunner.quantityValue)
	CreateInstanceCustomLodgementRunner.Flag(`timeframe`, `The time-frame (either a month "YYYY_MM" or a single day "YYYY_MM_DD") for which the quantity value is valid (currently for the PDT timezone only).`).Required().StringVar(&CreateInstanceCustomLodgementRunner.timeframe)
	registry[CreateInstanceCustomLodgementCmd.FullCommand()] = CreateInstanceCustomLodgementRunner

	DestroyInstanceCustomLodgementRunner := new(DestroyInstanceCustomLodgementInstanceCustomLodgementRunner)
	DestroyInstanceCustomLodgementCmd := resCmd.Command("DestroyInstanceCustomLodgement", `Destroy the specified lodgement.`)
	DestroyInstanceCustomLodgementRunner.Flag(`cloudId`, ``).Required().StringVar(&DestroyInstanceCustomLodgementRunner.cloudId)
	DestroyInstanceCustomLodgementRunner.Flag(`id`, ``).Required().StringVar(&DestroyInstanceCustomLodgementRunner.id)
	DestroyInstanceCustomLodgementRunner.Flag(`instanceId`, ``).Required().StringVar(&DestroyInstanceCustomLodgementRunner.instanceId)
	registry[DestroyInstanceCustomLodgementCmd.FullCommand()] = DestroyInstanceCustomLodgementRunner

	IndexInstanceCustomLodgementsRunner := new(IndexInstanceCustomLodgementsInstanceCustomLodgementRunner)
	IndexInstanceCustomLodgementsCmd := resCmd.Command("IndexInstanceCustomLodgements", `List InstanceCustomLodgements of a given cloud and instance.`)
	IndexInstanceCustomLodgementsRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexInstanceCustomLodgementsRunner.cloudId)
	IndexInstanceCustomLodgementsRunner.Flag(`instanceId`, ``).Required().StringVar(&IndexInstanceCustomLodgementsRunner.instanceId)
	IndexInstanceCustomLodgementsRunner.Flag(`view`, ``).StringVar(&IndexInstanceCustomLodgementsRunner.view)
	registry[IndexInstanceCustomLodgementsCmd.FullCommand()] = IndexInstanceCustomLodgementsRunner

	ShowInstanceCustomLodgementRunner := new(ShowInstanceCustomLodgementInstanceCustomLodgementRunner)
	ShowInstanceCustomLodgementCmd := resCmd.Command("ShowInstanceCustomLodgement", `Show the specified lodgement.`)
	ShowInstanceCustomLodgementRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowInstanceCustomLodgementRunner.cloudId)
	ShowInstanceCustomLodgementRunner.Flag(`id`, ``).Required().StringVar(&ShowInstanceCustomLodgementRunner.id)
	ShowInstanceCustomLodgementRunner.Flag(`instanceId`, ``).Required().StringVar(&ShowInstanceCustomLodgementRunner.instanceId)
	registry[ShowInstanceCustomLodgementCmd.FullCommand()] = ShowInstanceCustomLodgementRunner

	UpdateInstanceCustomLodgementRunner := new(UpdateInstanceCustomLodgementInstanceCustomLodgementRunner)
	UpdateInstanceCustomLodgementCmd := resCmd.Command("UpdateInstanceCustomLodgement", `Update a lodgement with the quantity specified.`)
	UpdateInstanceCustomLodgementRunner.Flag(`cloudId`, ``).Required().StringVar(&UpdateInstanceCustomLodgementRunner.cloudId)
	UpdateInstanceCustomLodgementRunner.Flag(`id`, ``).Required().StringVar(&UpdateInstanceCustomLodgementRunner.id)
	UpdateInstanceCustomLodgementRunner.Flag(`instanceId`, ``).Required().StringVar(&UpdateInstanceCustomLodgementRunner.instanceId)
	UpdateInstanceCustomLodgementRunner.FlagPattern(`quantity\.(\d+)\.name`, `The name of the quantity. A customer-specific string, e.g. "MB/s" or "GB/Month".`).Capture(&UpdateInstanceCustomLodgementRunner.quantityNamePos).StringsVar(&UpdateInstanceCustomLodgementRunner.quantityName)
	UpdateInstanceCustomLodgementRunner.FlagPattern(`quantity\.(\d+)\.value`, `The value of the quantity. Should be a positive integer.`).Capture(&UpdateInstanceCustomLodgementRunner.quantityValuePos).StringsVar(&UpdateInstanceCustomLodgementRunner.quantityValue)
	registry[UpdateInstanceCustomLodgementCmd.FullCommand()] = UpdateInstanceCustomLodgementRunner
}

/****** InstanceType ******/

type IndexInstanceTypesInstanceTypeRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexInstanceTypesInstanceTypeRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexInstanceTypes(r.cloudId, filter, r.view)
}

type ShowInstanceTypeInstanceTypeRunner struct {
	cloudId string
	id      string
	view    string
}

func (r *ShowInstanceTypeInstanceTypeRunner) Run(c *Client) (interface{}, error) {
	return c.ShowInstanceType(r.cloudId, r.id, r.view)
}

// Register all InstanceType actions
func registerInstanceTypeCmds(app *kingpin.Application) {
	resCmd := app.Cmd("InstanceType", ``)

	IndexInstanceTypesRunner := new(IndexInstanceTypesInstanceTypeRunner)
	IndexInstanceTypesCmd := resCmd.Command("IndexInstanceTypes", `Lists instance types.`)
	IndexInstanceTypesRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexInstanceTypesRunner.cloudId)
	IndexInstanceTypesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexInstanceTypesRunner.filterPos).StringsVar(&IndexInstanceTypesRunner.filter)
	IndexInstanceTypesRunner.Flag(`view`, ``).StringVar(&IndexInstanceTypesRunner.view)
	registry[IndexInstanceTypesCmd.FullCommand()] = IndexInstanceTypesRunner

	ShowInstanceTypeRunner := new(ShowInstanceTypeInstanceTypeRunner)
	ShowInstanceTypeCmd := resCmd.Command("ShowInstanceType", `Displays information about a single Instance type.`)
	ShowInstanceTypeRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowInstanceTypeRunner.cloudId)
	ShowInstanceTypeRunner.Flag(`id`, ``).Required().StringVar(&ShowInstanceTypeRunner.id)
	ShowInstanceTypeRunner.Flag(`view`, ``).StringVar(&ShowInstanceTypeRunner.view)
	registry[ShowInstanceTypeCmd.FullCommand()] = ShowInstanceTypeRunner
}

/****** IpAddress ******/

type CreateIpAddressIpAddressRunner struct {
	cloudId                 string
	ipAddressDeploymentHref string
	ipAddressDomain         string
	ipAddressName           string
	ipAddressNetworkHref    string
}

func (r *CreateIpAddressIpAddressRunner) Run(c *Client) (interface{}, error) {

	/** Handle ipAddress parameter **/
	var ipAddress IpAddressParam

	// Load JSON if provided
	if len(r.ipAddressJson) > 0 {
		if err := Json.Unmarshal(r.ipAddressJson, &ipAddress); err != nil {
			return nil, fmt.Errorf("Could not load ipAddress JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.ipAddressdeploymentHref) > 0 {
		ipAddress.ipAddress.deploymentHref = r.ipAddressdeploymentHref
	}

	if len(r.ipAddressdomain) > 0 {
		ipAddress.ipAddress.domain = r.ipAddressdomain
	}

	if len(r.ipAddressname) > 0 {
		ipAddress.ipAddress.name = r.ipAddressname
	}

	if len(r.ipAddressnetworkHref) > 0 {
		ipAddress.ipAddress.networkHref = r.ipAddressnetworkHref
	}

	return c.CreateIpAddress(r.cloudId, &ipAddress)
}

type DestroyIpAddressIpAddressRunner struct {
	cloudId string
	id      string
}

func (r *DestroyIpAddressIpAddressRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyIpAddress(r.cloudId, r.id)
}

type IndexIpAddressesIpAddressRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
}

func (r *IndexIpAddressesIpAddressRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexIpAddresses(r.cloudId, filter)
}

type ShowIpAddressIpAddressRunner struct {
	cloudId string
	id      string
}

func (r *ShowIpAddressIpAddressRunner) Run(c *Client) (interface{}, error) {
	return c.ShowIpAddress(r.cloudId, r.id)
}

type UpdateIpAddressIpAddressRunner struct {
	cloudId                 string
	id                      string
	ipAddressDeploymentHref string
	ipAddressName           string
}

func (r *UpdateIpAddressIpAddressRunner) Run(c *Client) (interface{}, error) {

	/** Handle ipAddress parameter **/
	var ipAddress IpAddressParam2

	// Load JSON if provided
	if len(r.ipAddressJson) > 0 {
		if err := Json.Unmarshal(r.ipAddressJson, &ipAddress); err != nil {
			return nil, fmt.Errorf("Could not load ipAddress JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.ipAddressdeploymentHref) > 0 {
		ipAddress.ipAddress.deploymentHref = r.ipAddressdeploymentHref
	}

	if len(r.ipAddressname) > 0 {
		ipAddress.ipAddress.name = r.ipAddressname
	}

	return c.UpdateIpAddress(r.cloudId, r.id, &ipAddress)
}

// Register all IpAddress actions
func registerIpAddressCmds(app *kingpin.Application) {
	resCmd := app.Cmd("IpAddress", `An IpAddress provides an abstraction for IPv4 addresses bindable to Instance resources running in a Cloud.`)

	CreateIpAddressRunner := new(CreateIpAddressIpAddressRunner)
	CreateIpAddressCmd := resCmd.Command("CreateIpAddress", `Creates a new IpAddress with the given parameters.`)
	CreateIpAddressRunner.Flag(`cloudId`, ``).Required().StringVar(&CreateIpAddressRunner.cloudId)
	CreateIpAddressRunner.Flag(`ipAddress.deploymentHref`, `The href of the Deployment that owns this IpAddress.`).StringVar(&CreateIpAddressRunner.ipAddressDeploymentHref)
	CreateIpAddressRunner.Flag(`ipAddress.domain`, `(Amazon Only) Pass vpc to create this IP for EC2-VPC only environments. Pass ec2_classic to create this IP for EC2-Classic environments. Defaults to ec2_classic.`).StringVar(&CreateIpAddressRunner.ipAddressDomain)
	CreateIpAddressRunner.Flag(`ipAddress.name`, `The name of the IpAddress to be created.`).Required().StringVar(&CreateIpAddressRunner.ipAddressName)
	CreateIpAddressRunner.Flag(`ipAddress.networkHref`, `(OpenStack Only) The href of the Network that the IpAddress will be associated to. This parameter is required for OpenStack with Neutron clouds.`).StringVar(&CreateIpAddressRunner.ipAddressNetworkHref)
	registry[CreateIpAddressCmd.FullCommand()] = CreateIpAddressRunner

	DestroyIpAddressRunner := new(DestroyIpAddressIpAddressRunner)
	DestroyIpAddressCmd := resCmd.Command("DestroyIpAddress", `Deletes a given IpAddress.`)
	DestroyIpAddressRunner.Flag(`cloudId`, ``).Required().StringVar(&DestroyIpAddressRunner.cloudId)
	DestroyIpAddressRunner.Flag(`id`, ``).Required().StringVar(&DestroyIpAddressRunner.id)
	registry[DestroyIpAddressCmd.FullCommand()] = DestroyIpAddressRunner

	IndexIpAddressesRunner := new(IndexIpAddressesIpAddressRunner)
	IndexIpAddressesCmd := resCmd.Command("IndexIpAddresses", `Lists the IpAddresses available to this account.`)
	IndexIpAddressesRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexIpAddressesRunner.cloudId)
	IndexIpAddressesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexIpAddressesRunner.filterPos).StringsVar(&IndexIpAddressesRunner.filter)
	registry[IndexIpAddressesCmd.FullCommand()] = IndexIpAddressesRunner

	ShowIpAddressRunner := new(ShowIpAddressIpAddressRunner)
	ShowIpAddressCmd := resCmd.Command("ShowIpAddress", `Show information about a single IpAddress.`)
	ShowIpAddressRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowIpAddressRunner.cloudId)
	ShowIpAddressRunner.Flag(`id`, ``).Required().StringVar(&ShowIpAddressRunner.id)
	registry[ShowIpAddressCmd.FullCommand()] = ShowIpAddressRunner

	UpdateIpAddressRunner := new(UpdateIpAddressIpAddressRunner)
	UpdateIpAddressCmd := resCmd.Command("UpdateIpAddress", `Updates attributes of a given IpAddress.`)
	UpdateIpAddressRunner.Flag(`cloudId`, ``).Required().StringVar(&UpdateIpAddressRunner.cloudId)
	UpdateIpAddressRunner.Flag(`id`, ``).Required().StringVar(&UpdateIpAddressRunner.id)
	UpdateIpAddressRunner.Flag(`ipAddress.deploymentHref`, `The href of the Deployment that owns this IpAddress.`).StringVar(&UpdateIpAddressRunner.ipAddressDeploymentHref)
	UpdateIpAddressRunner.Flag(`ipAddress.name`, `The updated name of the IpAddress.`).Required().StringVar(&UpdateIpAddressRunner.ipAddressName)
	registry[UpdateIpAddressCmd.FullCommand()] = UpdateIpAddressRunner
}

/****** IpAddressBinding ******/

type CreateIpAddressBindingIpAddressBindingRunner struct {
	cloudId                             string
	ipAddressBindingInstanceHref        string
	ipAddressBindingPrivatePort         string
	ipAddressBindingProtocol            string
	ipAddressBindingPublicIpAddressHref string
	ipAddressBindingPublicPort          string
	ipAddressId                         string
}

func (r *CreateIpAddressBindingIpAddressBindingRunner) Run(c *Client) (interface{}, error) {

	/** Handle ipAddressBinding parameter **/
	var ipAddressBinding IpAddressBindingParam

	// Load JSON if provided
	if len(r.ipAddressBindingJson) > 0 {
		if err := Json.Unmarshal(r.ipAddressBindingJson, &ipAddressBinding); err != nil {
			return nil, fmt.Errorf("Could not load ipAddressBinding JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.ipAddressBindinginstanceHref) > 0 {
		ipAddressBinding.ipAddressBinding.instanceHref = r.ipAddressBindinginstanceHref
	}

	if len(r.ipAddressBindingprivatePort) > 0 {
		ipAddressBinding.ipAddressBinding.privatePort = r.ipAddressBindingprivatePort
	}

	if len(r.ipAddressBindingprotocol) > 0 {
		ipAddressBinding.ipAddressBinding.protocol = r.ipAddressBindingprotocol
	}

	if len(r.ipAddressBindingpublicIpAddressHref) > 0 {
		ipAddressBinding.ipAddressBinding.publicIpAddressHref = r.ipAddressBindingpublicIpAddressHref
	}

	if len(r.ipAddressBindingpublicPort) > 0 {
		ipAddressBinding.ipAddressBinding.publicPort = r.ipAddressBindingpublicPort
	}

	return c.CreateIpAddressBinding(r.cloudId, &ipAddressBinding, r.ipAddressId)
}

type DestroyIpAddressBindingIpAddressBindingRunner struct {
	cloudId     string
	id          string
	ipAddressId string
}

func (r *DestroyIpAddressBindingIpAddressBindingRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyIpAddressBinding(r.cloudId, r.id, r.ipAddressId)
}

type IndexIpAddressBindingsIpAddressBindingRunner struct {
	cloudId     string
	filter      []string
	filterPos   []string
	ipAddressId string
}

func (r *IndexIpAddressBindingsIpAddressBindingRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexIpAddressBindings(r.cloudId, filter, r.ipAddressId)
}

type ShowIpAddressBindingIpAddressBindingRunner struct {
	cloudId     string
	id          string
	ipAddressId string
}

func (r *ShowIpAddressBindingIpAddressBindingRunner) Run(c *Client) (interface{}, error) {
	return c.ShowIpAddressBinding(r.cloudId, r.id, r.ipAddressId)
}

// Register all IpAddressBinding actions
func registerIpAddressBindingCmds(app *kingpin.Application) {
	resCmd := app.Cmd("IpAddressBinding", `An IpAddressBinding represents an abstraction for binding an IpAddress to an instance.`)

	CreateIpAddressBindingRunner := new(CreateIpAddressBindingIpAddressBindingRunner)
	CreateIpAddressBindingCmd := resCmd.Command("CreateIpAddressBinding", `Creates an ip address binding which attaches a specified IpAddress resource to a specified instance, and also allows for configuration of port forwarding rules...`)
	CreateIpAddressBindingRunner.Flag(`cloudId`, ``).Required().StringVar(&CreateIpAddressBindingRunner.cloudId)
	CreateIpAddressBindingRunner.Flag(`ipAddressBinding.instanceHref`, `The Instance to which this IpAddress should be bound.`).Required().StringVar(&CreateIpAddressBindingRunner.ipAddressBindingInstanceHref)
	CreateIpAddressBindingRunner.Flag(`ipAddressBinding.privatePort`, `Incoming network traffic will get forwarded to this port number on the specified Instance. If not specified, will use public port. Required unless public_ip_address_href is passed.`).StringVar(&CreateIpAddressBindingRunner.ipAddressBindingPrivatePort)
	CreateIpAddressBindingRunner.Flag(`ipAddressBinding.protocol`, `Transport layer protocol of traffic that may be forwarded from public port to private port on the Instance. Required unless public_ip_address_href is passed.`).StringVar(&CreateIpAddressBindingRunner.ipAddressBindingProtocol)
	CreateIpAddressBindingRunner.Flag(`ipAddressBinding.publicIpAddressHref`, `The IpAddress to bind to the specified instance. Required unless port forwarding rule params are passed.`).StringVar(&CreateIpAddressBindingRunner.ipAddressBindingPublicIpAddressHref)
	CreateIpAddressBindingRunner.Flag(`ipAddressBinding.publicPort`, `The incoming port for port forwarding. Incoming network traffic on this port will get forwarded (to the IP:Private Port of the specified Instance). Required unless public_ip_address_href is passed.`).StringVar(&CreateIpAddressBindingRunner.ipAddressBindingPublicPort)
	CreateIpAddressBindingRunner.Flag(`ipAddressId`, ``).Required().StringVar(&CreateIpAddressBindingRunner.ipAddressId)
	registry[CreateIpAddressBindingCmd.FullCommand()] = CreateIpAddressBindingRunner

	DestroyIpAddressBindingRunner := new(DestroyIpAddressBindingIpAddressBindingRunner)
	DestroyIpAddressBindingCmd := resCmd.Command("DestroyIpAddressBinding", `<no description>`)
	DestroyIpAddressBindingRunner.Flag(`cloudId`, ``).Required().StringVar(&DestroyIpAddressBindingRunner.cloudId)
	DestroyIpAddressBindingRunner.Flag(`id`, ``).Required().StringVar(&DestroyIpAddressBindingRunner.id)
	DestroyIpAddressBindingRunner.Flag(`ipAddressId`, ``).Required().StringVar(&DestroyIpAddressBindingRunner.ipAddressId)
	registry[DestroyIpAddressBindingCmd.FullCommand()] = DestroyIpAddressBindingRunner

	IndexIpAddressBindingsRunner := new(IndexIpAddressBindingsIpAddressBindingRunner)
	IndexIpAddressBindingsCmd := resCmd.Command("IndexIpAddressBindings", `Lists the ip address bindings available to this account.`)
	IndexIpAddressBindingsRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexIpAddressBindingsRunner.cloudId)
	IndexIpAddressBindingsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexIpAddressBindingsRunner.filterPos).StringsVar(&IndexIpAddressBindingsRunner.filter)
	IndexIpAddressBindingsRunner.Flag(`ipAddressId`, ``).Required().StringVar(&IndexIpAddressBindingsRunner.ipAddressId)
	registry[IndexIpAddressBindingsCmd.FullCommand()] = IndexIpAddressBindingsRunner

	ShowIpAddressBindingRunner := new(ShowIpAddressBindingIpAddressBindingRunner)
	ShowIpAddressBindingCmd := resCmd.Command("ShowIpAddressBinding", `Show information about a single ip address binding.`)
	ShowIpAddressBindingRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowIpAddressBindingRunner.cloudId)
	ShowIpAddressBindingRunner.Flag(`id`, ``).Required().StringVar(&ShowIpAddressBindingRunner.id)
	ShowIpAddressBindingRunner.Flag(`ipAddressId`, ``).Required().StringVar(&ShowIpAddressBindingRunner.ipAddressId)
	registry[ShowIpAddressBindingCmd.FullCommand()] = ShowIpAddressBindingRunner
}

/****** MonitoringMetric ******/

type DataMonitoringMetricMonitoringMetricRunner struct {
	cloudId    string
	end        string
	id         string
	instanceId string
	start      string
}

func (r *DataMonitoringMetricMonitoringMetricRunner) Run(c *Client) (interface{}, error) {
	return c.DataMonitoringMetric(r.cloudId, r.end, r.id, r.instanceId, r.start)
}

type IndexMonitoringMetricsMonitoringMetricRunner struct {
	cloudId    string
	filter     []string
	filterPos  []string
	instanceId string
	period     string
	size       string
	title      string
	tz         string
}

func (r *IndexMonitoringMetricsMonitoringMetricRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexMonitoringMetrics(r.cloudId, filter, r.instanceId, r.period, r.size, r.title, r.tz)
}

type ShowMonitoringMetricMonitoringMetricRunner struct {
	cloudId    string
	id         string
	instanceId string
	period     string
	size       string
	title      string
	tz         string
}

func (r *ShowMonitoringMetricMonitoringMetricRunner) Run(c *Client) (interface{}, error) {
	return c.ShowMonitoringMetric(r.cloudId, r.id, r.instanceId, r.period, r.size, r.title, r.tz)
}

// Register all MonitoringMetric actions
func registerMonitoringMetricCmds(app *kingpin.Application) {
	resCmd := app.Cmd("MonitoringMetric", `A monitoring metric is a stream of data that is captured in an instance. Metrics can be monitored, graphed and can be used as the basis for triggering alerts.`)

	DataMonitoringMetricRunner := new(DataMonitoringMetricMonitoringMetricRunner)
	DataMonitoringMetricCmd := resCmd.Command("DataMonitoringMetric", `Gives the raw monitoring data for a particular metric`)
	DataMonitoringMetricRunner.Flag(`cloudId`, ``).Required().StringVar(&DataMonitoringMetricRunner.cloudId)
	DataMonitoringMetricRunner.Flag(`end`, `An integer number of seconds from current time. e.g. -150 or 0 `).Required().StringVar(&DataMonitoringMetricRunner.end)
	DataMonitoringMetricRunner.Flag(`id`, ``).Required().StringVar(&DataMonitoringMetricRunner.id)
	DataMonitoringMetricRunner.Flag(`instanceId`, ``).Required().StringVar(&DataMonitoringMetricRunner.instanceId)
	DataMonitoringMetricRunner.Flag(`start`, `An integer number of seconds from current time. e.g. -300`).Required().StringVar(&DataMonitoringMetricRunner.start)
	registry[DataMonitoringMetricCmd.FullCommand()] = DataMonitoringMetricRunner

	IndexMonitoringMetricsRunner := new(IndexMonitoringMetricsMonitoringMetricRunner)
	IndexMonitoringMetricsCmd := resCmd.Command("IndexMonitoringMetrics", `Lists the monitoring metrics available for the instance and their corresponding graph hrefs.`)
	IndexMonitoringMetricsRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexMonitoringMetricsRunner.cloudId)
	IndexMonitoringMetricsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexMonitoringMetricsRunner.filterPos).StringsVar(&IndexMonitoringMetricsRunner.filter)
	IndexMonitoringMetricsRunner.Flag(`instanceId`, ``).Required().StringVar(&IndexMonitoringMetricsRunner.instanceId)
	IndexMonitoringMetricsRunner.Flag(`period`, `The time scale for which the graph is generated. Default is 'day'`).StringVar(&IndexMonitoringMetricsRunner.period)
	IndexMonitoringMetricsRunner.Flag(`size`, `The size of the graph to be generated. Default is 'small'.`).StringVar(&IndexMonitoringMetricsRunner.size)
	IndexMonitoringMetricsRunner.Flag(`title`, `The title of the graph.`).StringVar(&IndexMonitoringMetricsRunner.title)
	IndexMonitoringMetricsRunner.Flag(`tz`, `The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences. `).StringVar(&IndexMonitoringMetricsRunner.tz)
	registry[IndexMonitoringMetricsCmd.FullCommand()] = IndexMonitoringMetricsRunner

	ShowMonitoringMetricRunner := new(ShowMonitoringMetricMonitoringMetricRunner)
	ShowMonitoringMetricCmd := resCmd.Command("ShowMonitoringMetric", `Shows attributes of a single monitoring metric.`)
	ShowMonitoringMetricRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowMonitoringMetricRunner.cloudId)
	ShowMonitoringMetricRunner.Flag(`id`, ``).Required().StringVar(&ShowMonitoringMetricRunner.id)
	ShowMonitoringMetricRunner.Flag(`instanceId`, ``).Required().StringVar(&ShowMonitoringMetricRunner.instanceId)
	ShowMonitoringMetricRunner.Flag(`period`, `The time scale for which the graph is generated. Default is 'day'.`).StringVar(&ShowMonitoringMetricRunner.period)
	ShowMonitoringMetricRunner.Flag(`size`, `The size of the graph to be generated. Default is 'small'.`).StringVar(&ShowMonitoringMetricRunner.size)
	ShowMonitoringMetricRunner.Flag(`title`, `The title of the graph.`).StringVar(&ShowMonitoringMetricRunner.title)
	ShowMonitoringMetricRunner.Flag(`tz`, `The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences. `).StringVar(&ShowMonitoringMetricRunner.tz)
	registry[ShowMonitoringMetricCmd.FullCommand()] = ShowMonitoringMetricRunner
}

/****** MultiCloudImage ******/

type CloneMultiCloudImageMultiCloudImageRunner struct {
	id                         string
	multiCloudImageDescription string
	multiCloudImageName        string
}

func (r *CloneMultiCloudImageMultiCloudImageRunner) Run(c *Client) (interface{}, error) {

	/** Handle multiCloudImage parameter **/
	var multiCloudImage MultiCloudImageParam

	// Load JSON if provided
	if len(r.multiCloudImageJson) > 0 {
		if err := Json.Unmarshal(r.multiCloudImageJson, &multiCloudImage); err != nil {
			return nil, fmt.Errorf("Could not load multiCloudImage JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.multiCloudImagedescription) > 0 {
		multiCloudImage.multiCloudImage.description = r.multiCloudImagedescription
	}

	if len(r.multiCloudImagename) > 0 {
		multiCloudImage.multiCloudImage.name = r.multiCloudImagename
	}

	return c.CloneMultiCloudImage(r.id, &multiCloudImage)
}

type CommitMultiCloudImageMultiCloudImageRunner struct {
	commitMessage string
	id            string
}

func (r *CommitMultiCloudImageMultiCloudImageRunner) Run(c *Client) (interface{}, error) {
	return c.CommitMultiCloudImage(r.commitMessage, r.id)
}

type CreateMultiCloudImageMultiCloudImageRunner struct {
	multiCloudImageDescription string
	multiCloudImageName        string
	serverTemplateId           string
}

func (r *CreateMultiCloudImageMultiCloudImageRunner) Run(c *Client) (interface{}, error) {

	/** Handle multiCloudImage parameter **/
	var multiCloudImage MultiCloudImageParam2

	// Load JSON if provided
	if len(r.multiCloudImageJson) > 0 {
		if err := Json.Unmarshal(r.multiCloudImageJson, &multiCloudImage); err != nil {
			return nil, fmt.Errorf("Could not load multiCloudImage JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.multiCloudImagedescription) > 0 {
		multiCloudImage.multiCloudImage.description = r.multiCloudImagedescription
	}

	if len(r.multiCloudImagename) > 0 {
		multiCloudImage.multiCloudImage.name = r.multiCloudImagename
	}

	return c.CreateMultiCloudImage(&multiCloudImage, r.serverTemplateId)
}

type DestroyMultiCloudImageMultiCloudImageRunner struct {
	id               string
	serverTemplateId string
}

func (r *DestroyMultiCloudImageMultiCloudImageRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyMultiCloudImage(r.id, r.serverTemplateId)
}

type IndexMultiCloudImagesMultiCloudImageRunner struct {
	filter           []string
	filterPos        []string
	serverTemplateId string
}

func (r *IndexMultiCloudImagesMultiCloudImageRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexMultiCloudImages(filter, r.serverTemplateId)
}

type ShowMultiCloudImageMultiCloudImageRunner struct {
	id               string
	serverTemplateId string
}

func (r *ShowMultiCloudImageMultiCloudImageRunner) Run(c *Client) (interface{}, error) {
	return c.ShowMultiCloudImage(r.id, r.serverTemplateId)
}

type UpdateMultiCloudImageMultiCloudImageRunner struct {
	id                         string
	multiCloudImageDescription string
	multiCloudImageName        string
	serverTemplateId           string
}

func (r *UpdateMultiCloudImageMultiCloudImageRunner) Run(c *Client) (interface{}, error) {

	/** Handle multiCloudImage parameter **/
	var multiCloudImage MultiCloudImageParam

	// Load JSON if provided
	if len(r.multiCloudImageJson) > 0 {
		if err := Json.Unmarshal(r.multiCloudImageJson, &multiCloudImage); err != nil {
			return nil, fmt.Errorf("Could not load multiCloudImage JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.multiCloudImagedescription) > 0 {
		multiCloudImage.multiCloudImage.description = r.multiCloudImagedescription
	}

	if len(r.multiCloudImagename) > 0 {
		multiCloudImage.multiCloudImage.name = r.multiCloudImagename
	}

	return c.UpdateMultiCloudImage(r.id, &multiCloudImage, r.serverTemplateId)
}

// Register all MultiCloudImage actions
func registerMultiCloudImageCmds(app *kingpin.Application) {
	resCmd := app.Cmd("MultiCloudImage", `A MultiCloudImage is a RightScale component that functions as a pointer to machine images in specific clouds (e...`)

	CloneMultiCloudImageRunner := new(CloneMultiCloudImageMultiCloudImageRunner)
	CloneMultiCloudImageCmd := resCmd.Command("CloneMultiCloudImage", `Clones a given MultiCloudImage.`)
	CloneMultiCloudImageRunner.Flag(`id`, ``).Required().StringVar(&CloneMultiCloudImageRunner.id)
	CloneMultiCloudImageRunner.Flag(`multiCloudImage.description`, `The description for the cloned MultiCloudImage.`).StringVar(&CloneMultiCloudImageRunner.multiCloudImageDescription)
	CloneMultiCloudImageRunner.Flag(`multiCloudImage.name`, `The name for the cloned MultiCloudImage.`).Required().StringVar(&CloneMultiCloudImageRunner.multiCloudImageName)
	registry[CloneMultiCloudImageCmd.FullCommand()] = CloneMultiCloudImageRunner

	CommitMultiCloudImageRunner := new(CommitMultiCloudImageMultiCloudImageRunner)
	CommitMultiCloudImageCmd := resCmd.Command("CommitMultiCloudImage", `Commits a given MultiCloudImage. Only HEAD revisions can be committed.`)
	CommitMultiCloudImageRunner.Flag(`commitMessage`, `The message associated with the commit.`).Required().StringVar(&CommitMultiCloudImageRunner.commitMessage)
	CommitMultiCloudImageRunner.Flag(`id`, ``).Required().StringVar(&CommitMultiCloudImageRunner.id)
	registry[CommitMultiCloudImageCmd.FullCommand()] = CommitMultiCloudImageRunner

	CreateMultiCloudImageRunner := new(CreateMultiCloudImageMultiCloudImageRunner)
	CreateMultiCloudImageCmd := resCmd.Command("CreateMultiCloudImage", `Creates a new MultiCloudImage with the given parameters.`)
	CreateMultiCloudImageRunner.Flag(`multiCloudImage.description`, `The description of the MultiCloudImage to be created.`).StringVar(&CreateMultiCloudImageRunner.multiCloudImageDescription)
	CreateMultiCloudImageRunner.Flag(`multiCloudImage.name`, `The name of the MultiCloudImage to be created.`).Required().StringVar(&CreateMultiCloudImageRunner.multiCloudImageName)
	CreateMultiCloudImageRunner.Flag(`serverTemplateId`, ``).Required().StringVar(&CreateMultiCloudImageRunner.serverTemplateId)
	registry[CreateMultiCloudImageCmd.FullCommand()] = CreateMultiCloudImageRunner

	DestroyMultiCloudImageRunner := new(DestroyMultiCloudImageMultiCloudImageRunner)
	DestroyMultiCloudImageCmd := resCmd.Command("DestroyMultiCloudImage", `Deletes a given MultiCloudImage.`)
	DestroyMultiCloudImageRunner.Flag(`id`, ``).Required().StringVar(&DestroyMultiCloudImageRunner.id)
	DestroyMultiCloudImageRunner.Flag(`serverTemplateId`, ``).Required().StringVar(&DestroyMultiCloudImageRunner.serverTemplateId)
	registry[DestroyMultiCloudImageCmd.FullCommand()] = DestroyMultiCloudImageRunner

	IndexMultiCloudImagesRunner := new(IndexMultiCloudImagesMultiCloudImageRunner)
	IndexMultiCloudImagesCmd := resCmd.Command("IndexMultiCloudImages", `Lists the MultiCloudImages available to this account. HEAD revisions have a revision of 0.`)
	IndexMultiCloudImagesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexMultiCloudImagesRunner.filterPos).StringsVar(&IndexMultiCloudImagesRunner.filter)
	IndexMultiCloudImagesRunner.Flag(`serverTemplateId`, ``).Required().StringVar(&IndexMultiCloudImagesRunner.serverTemplateId)
	registry[IndexMultiCloudImagesCmd.FullCommand()] = IndexMultiCloudImagesRunner

	ShowMultiCloudImageRunner := new(ShowMultiCloudImageMultiCloudImageRunner)
	ShowMultiCloudImageCmd := resCmd.Command("ShowMultiCloudImage", `Show information about a single MultiCloudImage. HEAD revisions have a revision of 0.`)
	ShowMultiCloudImageRunner.Flag(`id`, ``).Required().StringVar(&ShowMultiCloudImageRunner.id)
	ShowMultiCloudImageRunner.Flag(`serverTemplateId`, ``).Required().StringVar(&ShowMultiCloudImageRunner.serverTemplateId)
	registry[ShowMultiCloudImageCmd.FullCommand()] = ShowMultiCloudImageRunner

	UpdateMultiCloudImageRunner := new(UpdateMultiCloudImageMultiCloudImageRunner)
	UpdateMultiCloudImageCmd := resCmd.Command("UpdateMultiCloudImage", `Updates attributes of a given MultiCloudImage. Only HEAD revisions can be updated (revision 0).`)
	UpdateMultiCloudImageRunner.Flag(`id`, ``).Required().StringVar(&UpdateMultiCloudImageRunner.id)
	UpdateMultiCloudImageRunner.Flag(`multiCloudImage.description`, `The updated description for the MultiCloudImage.`).StringVar(&UpdateMultiCloudImageRunner.multiCloudImageDescription)
	UpdateMultiCloudImageRunner.Flag(`multiCloudImage.name`, `The updated name for the MultiCloudImage.`).StringVar(&UpdateMultiCloudImageRunner.multiCloudImageName)
	UpdateMultiCloudImageRunner.Flag(`serverTemplateId`, ``).Required().StringVar(&UpdateMultiCloudImageRunner.serverTemplateId)
	registry[UpdateMultiCloudImageCmd.FullCommand()] = UpdateMultiCloudImageRunner
}

/****** MultiCloudImageSetting ******/

type CreateMultiCloudImageSettingMultiCloudImageSettingRunner struct {
	multiCloudImageId                      string
	multiCloudImageSettingCloudHref        string
	multiCloudImageSettingImageHref        string
	multiCloudImageSettingInstanceTypeHref string
	multiCloudImageSettingKernelImageHref  string
	multiCloudImageSettingRamdiskImageHref string
	multiCloudImageSettingUserData         string
}

func (r *CreateMultiCloudImageSettingMultiCloudImageSettingRunner) Run(c *Client) (interface{}, error) {

	/** Handle multiCloudImageSetting parameter **/
	var multiCloudImageSetting MultiCloudImageSettingParam

	// Load JSON if provided
	if len(r.multiCloudImageSettingJson) > 0 {
		if err := Json.Unmarshal(r.multiCloudImageSettingJson, &multiCloudImageSetting); err != nil {
			return nil, fmt.Errorf("Could not load multiCloudImageSetting JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.multiCloudImageSettingcloudHref) > 0 {
		multiCloudImageSetting.multiCloudImageSetting.cloudHref = r.multiCloudImageSettingcloudHref
	}

	if len(r.multiCloudImageSettingimageHref) > 0 {
		multiCloudImageSetting.multiCloudImageSetting.imageHref = r.multiCloudImageSettingimageHref
	}

	if len(r.multiCloudImageSettinginstanceTypeHref) > 0 {
		multiCloudImageSetting.multiCloudImageSetting.instanceTypeHref = r.multiCloudImageSettinginstanceTypeHref
	}

	if len(r.multiCloudImageSettingkernelImageHref) > 0 {
		multiCloudImageSetting.multiCloudImageSetting.kernelImageHref = r.multiCloudImageSettingkernelImageHref
	}

	if len(r.multiCloudImageSettingramdiskImageHref) > 0 {
		multiCloudImageSetting.multiCloudImageSetting.ramdiskImageHref = r.multiCloudImageSettingramdiskImageHref
	}

	if len(r.multiCloudImageSettinguserData) > 0 {
		multiCloudImageSetting.multiCloudImageSetting.userData = r.multiCloudImageSettinguserData
	}

	return c.CreateMultiCloudImageSetting(r.multiCloudImageId, &multiCloudImageSetting)
}

type DestroyMultiCloudImageSettingMultiCloudImageSettingRunner struct {
	id                string
	multiCloudImageId string
}

func (r *DestroyMultiCloudImageSettingMultiCloudImageSettingRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyMultiCloudImageSetting(r.id, r.multiCloudImageId)
}

type IndexMultiCloudImageSettingsMultiCloudImageSettingRunner struct {
	filter            []string
	filterPos         []string
	multiCloudImageId string
}

func (r *IndexMultiCloudImageSettingsMultiCloudImageSettingRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexMultiCloudImageSettings(filter, r.multiCloudImageId)
}

type ShowMultiCloudImageSettingMultiCloudImageSettingRunner struct {
	id                string
	multiCloudImageId string
}

func (r *ShowMultiCloudImageSettingMultiCloudImageSettingRunner) Run(c *Client) (interface{}, error) {
	return c.ShowMultiCloudImageSetting(r.id, r.multiCloudImageId)
}

type UpdateMultiCloudImageSettingMultiCloudImageSettingRunner struct {
	id                                     string
	multiCloudImageId                      string
	multiCloudImageSettingCloudHref        string
	multiCloudImageSettingImageHref        string
	multiCloudImageSettingInstanceTypeHref string
	multiCloudImageSettingKernelImageHref  string
	multiCloudImageSettingRamdiskImageHref string
	multiCloudImageSettingUserData         string
}

func (r *UpdateMultiCloudImageSettingMultiCloudImageSettingRunner) Run(c *Client) (interface{}, error) {

	/** Handle multiCloudImageSetting parameter **/
	var multiCloudImageSetting MultiCloudImageSettingParam2

	// Load JSON if provided
	if len(r.multiCloudImageSettingJson) > 0 {
		if err := Json.Unmarshal(r.multiCloudImageSettingJson, &multiCloudImageSetting); err != nil {
			return nil, fmt.Errorf("Could not load multiCloudImageSetting JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.multiCloudImageSettingcloudHref) > 0 {
		multiCloudImageSetting.multiCloudImageSetting.cloudHref = r.multiCloudImageSettingcloudHref
	}

	if len(r.multiCloudImageSettingimageHref) > 0 {
		multiCloudImageSetting.multiCloudImageSetting.imageHref = r.multiCloudImageSettingimageHref
	}

	if len(r.multiCloudImageSettinginstanceTypeHref) > 0 {
		multiCloudImageSetting.multiCloudImageSetting.instanceTypeHref = r.multiCloudImageSettinginstanceTypeHref
	}

	if len(r.multiCloudImageSettingkernelImageHref) > 0 {
		multiCloudImageSetting.multiCloudImageSetting.kernelImageHref = r.multiCloudImageSettingkernelImageHref
	}

	if len(r.multiCloudImageSettingramdiskImageHref) > 0 {
		multiCloudImageSetting.multiCloudImageSetting.ramdiskImageHref = r.multiCloudImageSettingramdiskImageHref
	}

	if len(r.multiCloudImageSettinguserData) > 0 {
		multiCloudImageSetting.multiCloudImageSetting.userData = r.multiCloudImageSettinguserData
	}

	return c.UpdateMultiCloudImageSetting(r.id, r.multiCloudImageId, &multiCloudImageSetting)
}

// Register all MultiCloudImageSetting actions
func registerMultiCloudImageSettingCmds(app *kingpin.Application) {
	resCmd := app.Cmd("MultiCloudImageSetting", `A MultiCloudImageSetting defines which settings should be used when a server is launched in a cloud...`)

	CreateMultiCloudImageSettingRunner := new(CreateMultiCloudImageSettingMultiCloudImageSettingRunner)
	CreateMultiCloudImageSettingCmd := resCmd.Command("CreateMultiCloudImageSetting", `Creates a new setting for an existing MultiCloudImage.`)
	CreateMultiCloudImageSettingRunner.Flag(`multiCloudImageId`, ``).Required().StringVar(&CreateMultiCloudImageSettingRunner.multiCloudImageId)
	CreateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.cloudHref`, `The href of the Cloud to use.`).StringVar(&CreateMultiCloudImageSettingRunner.multiCloudImageSettingCloudHref)
	CreateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.imageHref`, `The href of the Image to use. Mandatory if specifying cloud_href.`).StringVar(&CreateMultiCloudImageSettingRunner.multiCloudImageSettingImageHref)
	CreateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.instanceTypeHref`, `The href of the instance type. Mandatory if specifying cloud_href.`).StringVar(&CreateMultiCloudImageSettingRunner.multiCloudImageSettingInstanceTypeHref)
	CreateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.kernelImageHref`, `The href of the kernel image. Optional if specifying cloud_href.`).StringVar(&CreateMultiCloudImageSettingRunner.multiCloudImageSettingKernelImageHref)
	CreateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.ramdiskImageHref`, `The href of the ramdisk image. Optional if specifying cloud_href.`).StringVar(&CreateMultiCloudImageSettingRunner.multiCloudImageSettingRamdiskImageHref)
	CreateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.userData`, `User data that RightScale automatically passes to your instance at boot time.`).StringVar(&CreateMultiCloudImageSettingRunner.multiCloudImageSettingUserData)
	registry[CreateMultiCloudImageSettingCmd.FullCommand()] = CreateMultiCloudImageSettingRunner

	DestroyMultiCloudImageSettingRunner := new(DestroyMultiCloudImageSettingMultiCloudImageSettingRunner)
	DestroyMultiCloudImageSettingCmd := resCmd.Command("DestroyMultiCloudImageSetting", `Deletes a MultiCloudImage setting.`)
	DestroyMultiCloudImageSettingRunner.Flag(`id`, ``).Required().StringVar(&DestroyMultiCloudImageSettingRunner.id)
	DestroyMultiCloudImageSettingRunner.Flag(`multiCloudImageId`, ``).Required().StringVar(&DestroyMultiCloudImageSettingRunner.multiCloudImageId)
	registry[DestroyMultiCloudImageSettingCmd.FullCommand()] = DestroyMultiCloudImageSettingRunner

	IndexMultiCloudImageSettingsRunner := new(IndexMultiCloudImageSettingsMultiCloudImageSettingRunner)
	IndexMultiCloudImageSettingsCmd := resCmd.Command("IndexMultiCloudImageSettings", `Lists the MultiCloudImage settings.`)
	IndexMultiCloudImageSettingsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexMultiCloudImageSettingsRunner.filterPos).StringsVar(&IndexMultiCloudImageSettingsRunner.filter)
	IndexMultiCloudImageSettingsRunner.Flag(`multiCloudImageId`, ``).Required().StringVar(&IndexMultiCloudImageSettingsRunner.multiCloudImageId)
	registry[IndexMultiCloudImageSettingsCmd.FullCommand()] = IndexMultiCloudImageSettingsRunner

	ShowMultiCloudImageSettingRunner := new(ShowMultiCloudImageSettingMultiCloudImageSettingRunner)
	ShowMultiCloudImageSettingCmd := resCmd.Command("ShowMultiCloudImageSetting", `Show information about a single MultiCloudImage setting.`)
	ShowMultiCloudImageSettingRunner.Flag(`id`, ``).Required().StringVar(&ShowMultiCloudImageSettingRunner.id)
	ShowMultiCloudImageSettingRunner.Flag(`multiCloudImageId`, ``).Required().StringVar(&ShowMultiCloudImageSettingRunner.multiCloudImageId)
	registry[ShowMultiCloudImageSettingCmd.FullCommand()] = ShowMultiCloudImageSettingRunner

	UpdateMultiCloudImageSettingRunner := new(UpdateMultiCloudImageSettingMultiCloudImageSettingRunner)
	UpdateMultiCloudImageSettingCmd := resCmd.Command("UpdateMultiCloudImageSetting", `Updates a settings for a MultiCLoudImage.`)
	UpdateMultiCloudImageSettingRunner.Flag(`id`, ``).Required().StringVar(&UpdateMultiCloudImageSettingRunner.id)
	UpdateMultiCloudImageSettingRunner.Flag(`multiCloudImageId`, ``).Required().StringVar(&UpdateMultiCloudImageSettingRunner.multiCloudImageId)
	UpdateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.cloudHref`, `The href of the Cloud to use.`).StringVar(&UpdateMultiCloudImageSettingRunner.multiCloudImageSettingCloudHref)
	UpdateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.imageHref`, `The href of the Image to use.`).StringVar(&UpdateMultiCloudImageSettingRunner.multiCloudImageSettingImageHref)
	UpdateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.instanceTypeHref`, `The href of the instance type.`).StringVar(&UpdateMultiCloudImageSettingRunner.multiCloudImageSettingInstanceTypeHref)
	UpdateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.kernelImageHref`, `The href of the kernel image.`).StringVar(&UpdateMultiCloudImageSettingRunner.multiCloudImageSettingKernelImageHref)
	UpdateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.ramdiskImageHref`, `The href of the ramdisk image.`).StringVar(&UpdateMultiCloudImageSettingRunner.multiCloudImageSettingRamdiskImageHref)
	UpdateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.userData`, `User data that RightScale automatically passes to your instance at boot time.`).StringVar(&UpdateMultiCloudImageSettingRunner.multiCloudImageSettingUserData)
	registry[UpdateMultiCloudImageSettingCmd.FullCommand()] = UpdateMultiCloudImageSettingRunner
}

/****** Network ******/

type CreateNetworkNetworkRunner struct {
	networkCidrBlock       string
	networkCloudHref       string
	networkDescription     string
	networkInstanceTenancy string
	networkName            string
}

func (r *CreateNetworkNetworkRunner) Run(c *Client) (interface{}, error) {

	/** Handle network parameter **/
	var network NetworkParam

	// Load JSON if provided
	if len(r.networkJson) > 0 {
		if err := Json.Unmarshal(r.networkJson, &network); err != nil {
			return nil, fmt.Errorf("Could not load network JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.networkcidrBlock) > 0 {
		network.network.cidrBlock = r.networkcidrBlock
	}

	if len(r.networkcloudHref) > 0 {
		network.network.cloudHref = r.networkcloudHref
	}

	if len(r.networkdescription) > 0 {
		network.network.description = r.networkdescription
	}

	if len(r.networkinstanceTenancy) > 0 {
		network.network.instanceTenancy = r.networkinstanceTenancy
	}

	if len(r.networkname) > 0 {
		network.network.name = r.networkname
	}

	return c.CreateNetwork(&network)
}

type DestroyNetworkNetworkRunner struct {
	id string
}

func (r *DestroyNetworkNetworkRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyNetwork(r.id)
}

type IndexNetworksNetworkRunner struct {
	filter    []string
	filterPos []string
}

func (r *IndexNetworksNetworkRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexNetworks(filter)
}

type ShowNetworkNetworkRunner struct {
	id string
}

func (r *ShowNetworkNetworkRunner) Run(c *Client) (interface{}, error) {
	return c.ShowNetwork(r.id)
}

type UpdateNetworkNetworkRunner struct {
	id                    string
	networkDescription    string
	networkName           string
	networkRouteTableHref string
}

func (r *UpdateNetworkNetworkRunner) Run(c *Client) (interface{}, error) {

	/** Handle network parameter **/
	var network NetworkParam2

	// Load JSON if provided
	if len(r.networkJson) > 0 {
		if err := Json.Unmarshal(r.networkJson, &network); err != nil {
			return nil, fmt.Errorf("Could not load network JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.networkdescription) > 0 {
		network.network.description = r.networkdescription
	}

	if len(r.networkname) > 0 {
		network.network.name = r.networkname
	}

	if len(r.networkrouteTableHref) > 0 {
		network.network.routeTableHref = r.networkrouteTableHref
	}

	return c.UpdateNetwork(r.id, &network)
}

// Register all Network actions
func registerNetworkCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Network", `A Network is a logical grouping of network devices.`)

	CreateNetworkRunner := new(CreateNetworkNetworkRunner)
	CreateNetworkCmd := resCmd.Command("CreateNetwork", `Creates a new network.`)
	CreateNetworkRunner.Flag(`network.cidrBlock`, `The range of IP addresses for the Network. This parameter is required for Amazon clouds.`).StringVar(&CreateNetworkRunner.networkCidrBlock)
	CreateNetworkRunner.Flag(`network.cloudHref`, `The Cloud to create the Network in`).Required().StringVar(&CreateNetworkRunner.networkCloudHref)
	CreateNetworkRunner.Flag(`network.description`, `The description for the Network.`).StringVar(&CreateNetworkRunner.networkDescription)
	CreateNetworkRunner.Flag(`network.instanceTenancy`, `The launch policy for AWS instances in the Network. Specify 'default' to allow instances to decide their own launch policy. Specify 'dedicated' to force all instances to be launched as 'dedicated'.`).StringVar(&CreateNetworkRunner.networkInstanceTenancy)
	CreateNetworkRunner.Flag(`network.name`, `The name for the Network.`).StringVar(&CreateNetworkRunner.networkName)
	registry[CreateNetworkCmd.FullCommand()] = CreateNetworkRunner

	DestroyNetworkRunner := new(DestroyNetworkNetworkRunner)
	DestroyNetworkCmd := resCmd.Command("DestroyNetwork", `Deletes the given network(s).`)
	DestroyNetworkRunner.Flag(`id`, ``).Required().StringVar(&DestroyNetworkRunner.id)
	registry[DestroyNetworkCmd.FullCommand()] = DestroyNetworkRunner

	IndexNetworksRunner := new(IndexNetworksNetworkRunner)
	IndexNetworksCmd := resCmd.Command("IndexNetworks", `Lists networks in this account.`)
	IndexNetworksRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexNetworksRunner.filterPos).StringsVar(&IndexNetworksRunner.filter)
	registry[IndexNetworksCmd.FullCommand()] = IndexNetworksRunner

	ShowNetworkRunner := new(ShowNetworkNetworkRunner)
	ShowNetworkCmd := resCmd.Command("ShowNetwork", `Shows attributes of a single network.`)
	ShowNetworkRunner.Flag(`id`, ``).Required().StringVar(&ShowNetworkRunner.id)
	registry[ShowNetworkCmd.FullCommand()] = ShowNetworkRunner

	UpdateNetworkRunner := new(UpdateNetworkNetworkRunner)
	UpdateNetworkCmd := resCmd.Command("UpdateNetwork", `Updates the given network.`)
	UpdateNetworkRunner.Flag(`id`, ``).Required().StringVar(&UpdateNetworkRunner.id)
	UpdateNetworkRunner.Flag(`network.description`, `The updated description for the Network.`).StringVar(&UpdateNetworkRunner.networkDescription)
	UpdateNetworkRunner.Flag(`network.name`, `The updated name for the Network.`).StringVar(&UpdateNetworkRunner.networkName)
	UpdateNetworkRunner.Flag(`network.routeTableHref`, `Sets the default RouteTable for this Network.`).StringVar(&UpdateNetworkRunner.networkRouteTableHref)
	registry[UpdateNetworkCmd.FullCommand()] = UpdateNetworkRunner
}

/****** NetworkGateway ******/

type CreateNetworkGatewayNetworkGatewayRunner struct {
	networkGatewayCloudHref   string
	networkGatewayDescription string
	networkGatewayName        string
	networkGatewayType        string
}

func (r *CreateNetworkGatewayNetworkGatewayRunner) Run(c *Client) (interface{}, error) {

	/** Handle networkGateway parameter **/
	var networkGateway NetworkGatewayParam

	// Load JSON if provided
	if len(r.networkGatewayJson) > 0 {
		if err := Json.Unmarshal(r.networkGatewayJson, &networkGateway); err != nil {
			return nil, fmt.Errorf("Could not load networkGateway JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.networkGatewaycloudHref) > 0 {
		networkGateway.networkGateway.cloudHref = r.networkGatewaycloudHref
	}

	if len(r.networkGatewaydescription) > 0 {
		networkGateway.networkGateway.description = r.networkGatewaydescription
	}

	if len(r.networkGatewayname) > 0 {
		networkGateway.networkGateway.name = r.networkGatewayname
	}

	if len(r.networkGatewaytype_) > 0 {
		networkGateway.networkGateway.type_ = r.networkGatewaytype_
	}

	return c.CreateNetworkGateway(&networkGateway)
}

type DestroyNetworkGatewayNetworkGatewayRunner struct {
	id string
}

func (r *DestroyNetworkGatewayNetworkGatewayRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyNetworkGateway(r.id)
}

type IndexNetworkGatewaysNetworkGatewayRunner struct {
	filter    []string
	filterPos []string
}

func (r *IndexNetworkGatewaysNetworkGatewayRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexNetworkGateways(filter)
}

type ShowNetworkGatewayNetworkGatewayRunner struct {
	id string
}

func (r *ShowNetworkGatewayNetworkGatewayRunner) Run(c *Client) (interface{}, error) {
	return c.ShowNetworkGateway(r.id)
}

type UpdateNetworkGatewayNetworkGatewayRunner struct {
	id                        string
	networkGatewayDescription string
	networkGatewayName        string
	networkGatewayNetworkHref string
}

func (r *UpdateNetworkGatewayNetworkGatewayRunner) Run(c *Client) (interface{}, error) {

	/** Handle networkGateway parameter **/
	var networkGateway NetworkGatewayParam2

	// Load JSON if provided
	if len(r.networkGatewayJson) > 0 {
		if err := Json.Unmarshal(r.networkGatewayJson, &networkGateway); err != nil {
			return nil, fmt.Errorf("Could not load networkGateway JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.networkGatewaydescription) > 0 {
		networkGateway.networkGateway.description = r.networkGatewaydescription
	}

	if len(r.networkGatewayname) > 0 {
		networkGateway.networkGateway.name = r.networkGatewayname
	}

	if len(r.networkGatewaynetworkHref) > 0 {
		networkGateway.networkGateway.networkHref = r.networkGatewaynetworkHref
	}

	return c.UpdateNetworkGateway(r.id, &networkGateway)
}

// Register all NetworkGateway actions
func registerNetworkGatewayCmds(app *kingpin.Application) {
	resCmd := app.Cmd("NetworkGateway", `A NetworkGateway is an interface that allows traffic to be routed between networks.`)

	CreateNetworkGatewayRunner := new(CreateNetworkGatewayNetworkGatewayRunner)
	CreateNetworkGatewayCmd := resCmd.Command("CreateNetworkGateway", `Create a new NetworkGateway.`)
	CreateNetworkGatewayRunner.Flag(`networkGateway.cloudHref`, `The cloud to create the NetworkGateway in.`).Required().StringVar(&CreateNetworkGatewayRunner.networkGatewayCloudHref)
	CreateNetworkGatewayRunner.Flag(`networkGateway.description`, `The description to be set on the NetworkGateway.`).StringVar(&CreateNetworkGatewayRunner.networkGatewayDescription)
	CreateNetworkGatewayRunner.Flag(`networkGateway.name`, `The name to be set on the NetworkGateway.`).Required().StringVar(&CreateNetworkGatewayRunner.networkGatewayName)
	CreateNetworkGatewayRunner.Flag(`networkGateway.type_`, `The type of the NetworkGateway.`).Required().StringVar(&CreateNetworkGatewayRunner.networkGatewayType)
	registry[CreateNetworkGatewayCmd.FullCommand()] = CreateNetworkGatewayRunner

	DestroyNetworkGatewayRunner := new(DestroyNetworkGatewayNetworkGatewayRunner)
	DestroyNetworkGatewayCmd := resCmd.Command("DestroyNetworkGateway", `Delete an existing NetworkGateway.`)
	DestroyNetworkGatewayRunner.Flag(`id`, ``).Required().StringVar(&DestroyNetworkGatewayRunner.id)
	registry[DestroyNetworkGatewayCmd.FullCommand()] = DestroyNetworkGatewayRunner

	IndexNetworkGatewaysRunner := new(IndexNetworkGatewaysNetworkGatewayRunner)
	IndexNetworkGatewaysCmd := resCmd.Command("IndexNetworkGateways", `Lists the NetworkGateways available to this account.`)
	IndexNetworkGatewaysRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexNetworkGatewaysRunner.filterPos).StringsVar(&IndexNetworkGatewaysRunner.filter)
	registry[IndexNetworkGatewaysCmd.FullCommand()] = IndexNetworkGatewaysRunner

	ShowNetworkGatewayRunner := new(ShowNetworkGatewayNetworkGatewayRunner)
	ShowNetworkGatewayCmd := resCmd.Command("ShowNetworkGateway", `Show information about a single NetworkGateway.`)
	ShowNetworkGatewayRunner.Flag(`id`, ``).Required().StringVar(&ShowNetworkGatewayRunner.id)
	registry[ShowNetworkGatewayCmd.FullCommand()] = ShowNetworkGatewayRunner

	UpdateNetworkGatewayRunner := new(UpdateNetworkGatewayNetworkGatewayRunner)
	UpdateNetworkGatewayCmd := resCmd.Command("UpdateNetworkGateway", `Update an existing NetworkGateway.`)
	UpdateNetworkGatewayRunner.Flag(`id`, ``).Required().StringVar(&UpdateNetworkGatewayRunner.id)
	UpdateNetworkGatewayRunner.Flag(`networkGateway.description`, `The description to be set on the NetworkGateway.`).StringVar(&UpdateNetworkGatewayRunner.networkGatewayDescription)
	UpdateNetworkGatewayRunner.Flag(`networkGateway.name`, `The name to be set on the NetworkGateway.`).StringVar(&UpdateNetworkGatewayRunner.networkGatewayName)
	UpdateNetworkGatewayRunner.Flag(`networkGateway.networkHref`, `Pass a blank string to detach from the specified Network, or pass a valid Network href to attach to the specified network.`).StringVar(&UpdateNetworkGatewayRunner.networkGatewayNetworkHref)
	registry[UpdateNetworkGatewayCmd.FullCommand()] = UpdateNetworkGatewayRunner
}

/****** NetworkOptionGroup ******/

type CreateNetworkOptionGroupNetworkOptionGroupRunner struct {
	networkOptionGroupCloudHref     string
	networkOptionGroupDescription   string
	networkOptionGroupName          string
	networkOptionGroupOptionsValues []string
	networkOptionGroupOptionsNames  []string
	networkOptionGroupType          string
}

func (r *CreateNetworkOptionGroupNetworkOptionGroupRunner) Run(c *Client) (interface{}, error) {

	/** Handle networkOptionGroup parameter **/
	var networkOptionGroup NetworkOptionGroupParam

	// Load JSON if provided
	if len(r.networkOptionGroupJson) > 0 {
		if err := Json.Unmarshal(r.networkOptionGroupJson, &networkOptionGroup); err != nil {
			return nil, fmt.Errorf("Could not load networkOptionGroup JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.networkOptionGroupcloudHref) > 0 {
		networkOptionGroup.networkOptionGroup.cloudHref = r.networkOptionGroupcloudHref
	}

	if len(r.networkOptionGroupdescription) > 0 {
		networkOptionGroup.networkOptionGroup.description = r.networkOptionGroupdescription
	}

	if len(r.networkOptionGroupname) > 0 {
		networkOptionGroup.networkOptionGroup.name = r.networkOptionGroupname
	}

	if len(r.networkOptionGrouptype_) > 0 {
		networkOptionGroup.networkOptionGroup.type_ = r.networkOptionGrouptype_
	}

	// Create enumarable fields from flags
	networkOptionGroupoptions := make(map[string]string, len(r.networkOptionGroupoptionsNames))
	for i, n := range r.networkOptionGroupoptionsNames {
		networkOptionGroupoptions[n] = r.networkOptionGroupoptionsValues[i]
	}
	networkOptionGroup.networkOptionGroup.options = networkOptionGroupoptions

	return c.CreateNetworkOptionGroup(&networkOptionGroup)
}

type DestroyNetworkOptionGroupNetworkOptionGroupRunner struct {
	id string
}

func (r *DestroyNetworkOptionGroupNetworkOptionGroupRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyNetworkOptionGroup(r.id)
}

type IndexNetworkOptionGroupsNetworkOptionGroupRunner struct {
	filter    []string
	filterPos []string
}

func (r *IndexNetworkOptionGroupsNetworkOptionGroupRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexNetworkOptionGroups(filter)
}

type ShowNetworkOptionGroupNetworkOptionGroupRunner struct {
	id string
}

func (r *ShowNetworkOptionGroupNetworkOptionGroupRunner) Run(c *Client) (interface{}, error) {
	return c.ShowNetworkOptionGroup(r.id)
}

type UpdateNetworkOptionGroupNetworkOptionGroupRunner struct {
	id                            string
	networkOptionGroupDescription string
	networkOptionGroupName        string
}

func (r *UpdateNetworkOptionGroupNetworkOptionGroupRunner) Run(c *Client) (interface{}, error) {

	/** Handle networkOptionGroup parameter **/
	var networkOptionGroup NetworkOptionGroupParam2

	// Load JSON if provided
	if len(r.networkOptionGroupJson) > 0 {
		if err := Json.Unmarshal(r.networkOptionGroupJson, &networkOptionGroup); err != nil {
			return nil, fmt.Errorf("Could not load networkOptionGroup JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.networkOptionGroupdescription) > 0 {
		networkOptionGroup.networkOptionGroup.description = r.networkOptionGroupdescription
	}

	if len(r.networkOptionGroupname) > 0 {
		networkOptionGroup.networkOptionGroup.name = r.networkOptionGroupname
	}

	return c.UpdateNetworkOptionGroup(r.id, &networkOptionGroup)
}

// Register all NetworkOptionGroup actions
func registerNetworkOptionGroupCmds(app *kingpin.Application) {
	resCmd := app.Cmd("NetworkOptionGroup", `A key/value pair hash containing options for configuring a Network.`)

	CreateNetworkOptionGroupRunner := new(CreateNetworkOptionGroupNetworkOptionGroupRunner)
	CreateNetworkOptionGroupCmd := resCmd.Command("CreateNetworkOptionGroup", `Create a new NetworkOptionGroup.`)
	CreateNetworkOptionGroupRunner.Flag(`networkOptionGroup.cloudHref`, `The Cloud to create this NetworkOptionGroup in`).Required().StringVar(&CreateNetworkOptionGroupRunner.networkOptionGroupCloudHref)
	CreateNetworkOptionGroupRunner.Flag(`networkOptionGroup.description`, `Description of this NetworkOptionGroup`).StringVar(&CreateNetworkOptionGroupRunner.networkOptionGroupDescription)
	CreateNetworkOptionGroupRunner.Flag(`networkOptionGroup.name`, `Name of this NetworkOptionGroup`).StringVar(&CreateNetworkOptionGroupRunner.networkOptionGroupName)
	CreateNetworkOptionGroupRunner.FlagPattern(`networkOptionGroup\.options\.([a-z0-9_]+)`, ``).Required().Capture(&CreateNetworkOptionGroupRunner.networkOptionGroupOptionsNames).StringVar(&CreateNetworkOptionGroupRunner.networkOptionGroupOptionsValues)
	CreateNetworkOptionGroupRunner.Flag(`networkOptionGroup.type_`, `Type of this NetworkOptionGroup`).Required().StringVar(&CreateNetworkOptionGroupRunner.networkOptionGroupType)
	registry[CreateNetworkOptionGroupCmd.FullCommand()] = CreateNetworkOptionGroupRunner

	DestroyNetworkOptionGroupRunner := new(DestroyNetworkOptionGroupNetworkOptionGroupRunner)
	DestroyNetworkOptionGroupCmd := resCmd.Command("DestroyNetworkOptionGroup", `Delete an existing NetworkOptionGroup.`)
	DestroyNetworkOptionGroupRunner.Flag(`id`, ``).Required().StringVar(&DestroyNetworkOptionGroupRunner.id)
	registry[DestroyNetworkOptionGroupCmd.FullCommand()] = DestroyNetworkOptionGroupRunner

	IndexNetworkOptionGroupsRunner := new(IndexNetworkOptionGroupsNetworkOptionGroupRunner)
	IndexNetworkOptionGroupsCmd := resCmd.Command("IndexNetworkOptionGroups", `List NetworkOptionGroups available in this account.`)
	IndexNetworkOptionGroupsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexNetworkOptionGroupsRunner.filterPos).StringsVar(&IndexNetworkOptionGroupsRunner.filter)
	registry[IndexNetworkOptionGroupsCmd.FullCommand()] = IndexNetworkOptionGroupsRunner

	ShowNetworkOptionGroupRunner := new(ShowNetworkOptionGroupNetworkOptionGroupRunner)
	ShowNetworkOptionGroupCmd := resCmd.Command("ShowNetworkOptionGroup", `Show information about a single NetworkOptionGroup.`)
	ShowNetworkOptionGroupRunner.Flag(`id`, ``).Required().StringVar(&ShowNetworkOptionGroupRunner.id)
	registry[ShowNetworkOptionGroupCmd.FullCommand()] = ShowNetworkOptionGroupRunner

	UpdateNetworkOptionGroupRunner := new(UpdateNetworkOptionGroupNetworkOptionGroupRunner)
	UpdateNetworkOptionGroupCmd := resCmd.Command("UpdateNetworkOptionGroup", `Update an existing NetworkOptionGroup.`)
	UpdateNetworkOptionGroupRunner.Flag(`id`, ``).Required().StringVar(&UpdateNetworkOptionGroupRunner.id)
	UpdateNetworkOptionGroupRunner.Flag(`networkOptionGroup.description`, `Update the description`).StringVar(&UpdateNetworkOptionGroupRunner.networkOptionGroupDescription)
	UpdateNetworkOptionGroupRunner.Flag(`networkOptionGroup.name`, `Update the name`).StringVar(&UpdateNetworkOptionGroupRunner.networkOptionGroupName)
	registry[UpdateNetworkOptionGroupCmd.FullCommand()] = UpdateNetworkOptionGroupRunner
}

/****** NetworkOptionGroupAttachment ******/

type CreateNetworkOptionGroupAttachmentNetworkOptionGroupAttachmentRunner struct {
	networkOptionGroupAttachmentNetworkHref            string
	networkOptionGroupAttachmentNetworkOptionGroupHref string
}

func (r *CreateNetworkOptionGroupAttachmentNetworkOptionGroupAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle networkOptionGroupAttachment parameter **/
	var networkOptionGroupAttachment NetworkOptionGroupAttachmentParam

	// Load JSON if provided
	if len(r.networkOptionGroupAttachmentJson) > 0 {
		if err := Json.Unmarshal(r.networkOptionGroupAttachmentJson, &networkOptionGroupAttachment); err != nil {
			return nil, fmt.Errorf("Could not load networkOptionGroupAttachment JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.networkOptionGroupAttachmentnetworkHref) > 0 {
		networkOptionGroupAttachment.networkOptionGroupAttachment.networkHref = r.networkOptionGroupAttachmentnetworkHref
	}

	if len(r.networkOptionGroupAttachmentnetworkOptionGroupHref) > 0 {
		networkOptionGroupAttachment.networkOptionGroupAttachment.networkOptionGroupHref = r.networkOptionGroupAttachmentnetworkOptionGroupHref
	}

	return c.CreateNetworkOptionGroupAttachment(&networkOptionGroupAttachment)
}

type DestroyNetworkOptionGroupAttachmentNetworkOptionGroupAttachmentRunner struct {
	id string
}

func (r *DestroyNetworkOptionGroupAttachmentNetworkOptionGroupAttachmentRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyNetworkOptionGroupAttachment(r.id)
}

type IndexNetworkOptionGroupAttachmentsNetworkOptionGroupAttachmentRunner struct {
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexNetworkOptionGroupAttachmentsNetworkOptionGroupAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexNetworkOptionGroupAttachments(filter, r.view)
}

type ShowNetworkOptionGroupAttachmentNetworkOptionGroupAttachmentRunner struct {
	id   string
	view string
}

func (r *ShowNetworkOptionGroupAttachmentNetworkOptionGroupAttachmentRunner) Run(c *Client) (interface{}, error) {
	return c.ShowNetworkOptionGroupAttachment(r.id, r.view)
}

type UpdateNetworkOptionGroupAttachmentNetworkOptionGroupAttachmentRunner struct {
	id                                                 string
	networkOptionGroupAttachmentNetworkOptionGroupHref string
}

func (r *UpdateNetworkOptionGroupAttachmentNetworkOptionGroupAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle networkOptionGroupAttachment parameter **/
	var networkOptionGroupAttachment NetworkOptionGroupAttachmentParam2

	// Load JSON if provided
	if len(r.networkOptionGroupAttachmentJson) > 0 {
		if err := Json.Unmarshal(r.networkOptionGroupAttachmentJson, &networkOptionGroupAttachment); err != nil {
			return nil, fmt.Errorf("Could not load networkOptionGroupAttachment JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.networkOptionGroupAttachmentnetworkOptionGroupHref) > 0 {
		networkOptionGroupAttachment.networkOptionGroupAttachment.networkOptionGroupHref = r.networkOptionGroupAttachmentnetworkOptionGroupHref
	}

	return c.UpdateNetworkOptionGroupAttachment(r.id, &networkOptionGroupAttachment)
}

// Register all NetworkOptionGroupAttachment actions
func registerNetworkOptionGroupAttachmentCmds(app *kingpin.Application) {
	resCmd := app.Cmd("NetworkOptionGroupAttachment", `Resource for attaching NetworkOptionGroups to Networks.`)

	CreateNetworkOptionGroupAttachmentRunner := new(CreateNetworkOptionGroupAttachmentNetworkOptionGroupAttachmentRunner)
	CreateNetworkOptionGroupAttachmentCmd := resCmd.Command("CreateNetworkOptionGroupAttachment", `Create a new NetworkOptionGroupAttachment.`)
	CreateNetworkOptionGroupAttachmentRunner.Flag(`networkOptionGroupAttachment.networkHref`, `The Network to attach the specified NetworkOptionGroup to.`).Required().StringVar(&CreateNetworkOptionGroupAttachmentRunner.networkOptionGroupAttachmentNetworkHref)
	CreateNetworkOptionGroupAttachmentRunner.Flag(`networkOptionGroupAttachment.networkOptionGroupHref`, `The NetworkOptionGroup to attach to the specified resource.`).StringVar(&CreateNetworkOptionGroupAttachmentRunner.networkOptionGroupAttachmentNetworkOptionGroupHref)
	registry[CreateNetworkOptionGroupAttachmentCmd.FullCommand()] = CreateNetworkOptionGroupAttachmentRunner

	DestroyNetworkOptionGroupAttachmentRunner := new(DestroyNetworkOptionGroupAttachmentNetworkOptionGroupAttachmentRunner)
	DestroyNetworkOptionGroupAttachmentCmd := resCmd.Command("DestroyNetworkOptionGroupAttachment", `Delete an existing NetworkOptionGroupAttachment.`)
	DestroyNetworkOptionGroupAttachmentRunner.Flag(`id`, ``).Required().StringVar(&DestroyNetworkOptionGroupAttachmentRunner.id)
	registry[DestroyNetworkOptionGroupAttachmentCmd.FullCommand()] = DestroyNetworkOptionGroupAttachmentRunner

	IndexNetworkOptionGroupAttachmentsRunner := new(IndexNetworkOptionGroupAttachmentsNetworkOptionGroupAttachmentRunner)
	IndexNetworkOptionGroupAttachmentsCmd := resCmd.Command("IndexNetworkOptionGroupAttachments", `List NetworkOptionGroupAttachments in this account.`)
	IndexNetworkOptionGroupAttachmentsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexNetworkOptionGroupAttachmentsRunner.filterPos).StringsVar(&IndexNetworkOptionGroupAttachmentsRunner.filter)
	IndexNetworkOptionGroupAttachmentsRunner.Flag(`view`, ``).StringVar(&IndexNetworkOptionGroupAttachmentsRunner.view)
	registry[IndexNetworkOptionGroupAttachmentsCmd.FullCommand()] = IndexNetworkOptionGroupAttachmentsRunner

	ShowNetworkOptionGroupAttachmentRunner := new(ShowNetworkOptionGroupAttachmentNetworkOptionGroupAttachmentRunner)
	ShowNetworkOptionGroupAttachmentCmd := resCmd.Command("ShowNetworkOptionGroupAttachment", `Show information about a single NetworkOptionGroupAttachment.`)
	ShowNetworkOptionGroupAttachmentRunner.Flag(`id`, ``).Required().StringVar(&ShowNetworkOptionGroupAttachmentRunner.id)
	ShowNetworkOptionGroupAttachmentRunner.Flag(`view`, ``).StringVar(&ShowNetworkOptionGroupAttachmentRunner.view)
	registry[ShowNetworkOptionGroupAttachmentCmd.FullCommand()] = ShowNetworkOptionGroupAttachmentRunner

	UpdateNetworkOptionGroupAttachmentRunner := new(UpdateNetworkOptionGroupAttachmentNetworkOptionGroupAttachmentRunner)
	UpdateNetworkOptionGroupAttachmentCmd := resCmd.Command("UpdateNetworkOptionGroupAttachment", `Update an existing NetworkOptionGroupAttachment.`)
	UpdateNetworkOptionGroupAttachmentRunner.Flag(`id`, ``).Required().StringVar(&UpdateNetworkOptionGroupAttachmentRunner.id)
	UpdateNetworkOptionGroupAttachmentRunner.Flag(`networkOptionGroupAttachment.networkOptionGroupHref`, `The NetworkOptionGroup to attach to the specified resource.`).StringVar(&UpdateNetworkOptionGroupAttachmentRunner.networkOptionGroupAttachmentNetworkOptionGroupHref)
	registry[UpdateNetworkOptionGroupAttachmentCmd.FullCommand()] = UpdateNetworkOptionGroupAttachmentRunner
}

/****** Oauth2 ******/

type CreateOauth2Oauth2Runner struct {
	accountId        int
	clientId         string
	clientSecret     string
	grantType        string
	refreshToken     string
	rightLinkVersion string
	rsVersion        int
}

func (r *CreateOauth2Oauth2Runner) Run(c *Client) (interface{}, error) {
	return c.CreateOauth2(r.accountId, r.clientId, r.clientSecret, r.grantType, r.refreshToken, r.rightLinkVersion, r.rsVersion)
}

// Register all Oauth2 actions
func registerOauth2Cmds(app *kingpin.Application) {
	resCmd := app.Cmd("Oauth2", `Note that all API calls irrespective of the resource it is acting on, should pass a header "X-Api-Version" with the value "1...`)

	CreateOauth2Runner := new(CreateOauth2Oauth2Runner)
	CreateOauth2Cmd := resCmd.Command("CreateOauth2", `Perform an OAuth 2`)
	CreateOauth2Runner.Flag(`accountId`, `The client's account ID (only needed for instance agent clients).`).IntVar(&CreateOauth2Runner.accountId)
	CreateOauth2Runner.Flag(`clientId`, `The client ID (only needed for confidential clients).`).StringVar(&CreateOauth2Runner.clientId)
	CreateOauth2Runner.Flag(`clientSecret`, `The client secret (only needed for confidential clients).`).StringVar(&CreateOauth2Runner.clientSecret)
	CreateOauth2Runner.Flag(`grantType`, `Type of grant.`).Required().StringVar(&CreateOauth2Runner.grantType)
	CreateOauth2Runner.Flag(`refreshToken`, `The refresh token obtained from OAuth grant.`).StringVar(&CreateOauth2Runner.refreshToken)
	CreateOauth2Runner.Flag(`rightLinkVersion`, `The RightLink gem version the client conforms to (only needed for instance agent clients).`).StringVar(&CreateOauth2Runner.rightLinkVersion)
	CreateOauth2Runner.Flag(`rsVersion`, `The RightAgent protocol version the client conforms to (only needed for instance agent clients).`).IntVar(&CreateOauth2Runner.rsVersion)
	registry[CreateOauth2Cmd.FullCommand()] = CreateOauth2Runner
}

/****** Permission ******/

type CreatePermissionPermissionRunner struct {
	permissionRoleTitle string
	permissionUserHref  string
}

func (r *CreatePermissionPermissionRunner) Run(c *Client) (interface{}, error) {

	/** Handle permission parameter **/
	var permission PermissionParam

	// Load JSON if provided
	if len(r.permissionJson) > 0 {
		if err := Json.Unmarshal(r.permissionJson, &permission); err != nil {
			return nil, fmt.Errorf("Could not load permission JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.permissionroleTitle) > 0 {
		permission.permission.roleTitle = r.permissionroleTitle
	}

	if len(r.permissionuserHref) > 0 {
		permission.permission.userHref = r.permissionuserHref
	}

	return c.CreatePermission(&permission)
}

type DestroyPermissionPermissionRunner struct {
	id string
}

func (r *DestroyPermissionPermissionRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyPermission(r.id)
}

type IndexPermissionsPermissionRunner struct {
	filter    []string
	filterPos []string
}

func (r *IndexPermissionsPermissionRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexPermissions(filter)
}

type ShowPermissionPermissionRunner struct {
	id string
}

func (r *ShowPermissionPermissionRunner) Run(c *Client) (interface{}, error) {
	return c.ShowPermission(r.id)
}

// Register all Permission actions
func registerPermissionCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Permission", ``)

	CreatePermissionRunner := new(CreatePermissionPermissionRunner)
	CreatePermissionCmd := resCmd.Command("CreatePermission", `Create a permission, thereby granting some user a particular role with respect to the current account...`)
	CreatePermissionRunner.Flag(`permission.roleTitle`, ``).Required().StringVar(&CreatePermissionRunner.permissionRoleTitle)
	CreatePermissionRunner.Flag(`permission.userHref`, ``).Required().StringVar(&CreatePermissionRunner.permissionUserHref)
	registry[CreatePermissionCmd.FullCommand()] = CreatePermissionRunner

	DestroyPermissionRunner := new(DestroyPermissionPermissionRunner)
	DestroyPermissionCmd := resCmd.Command("DestroyPermission", `Destroy a permission, thereby revoking a user's role with respect to the current account...`)
	DestroyPermissionRunner.Flag(`id`, ``).Required().StringVar(&DestroyPermissionRunner.id)
	registry[DestroyPermissionCmd.FullCommand()] = DestroyPermissionRunner

	IndexPermissionsRunner := new(IndexPermissionsPermissionRunner)
	IndexPermissionsCmd := resCmd.Command("IndexPermissions", `List all permissions for all users of the current acount.`)
	IndexPermissionsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexPermissionsRunner.filterPos).StringsVar(&IndexPermissionsRunner.filter)
	registry[IndexPermissionsCmd.FullCommand()] = IndexPermissionsRunner

	ShowPermissionRunner := new(ShowPermissionPermissionRunner)
	ShowPermissionCmd := resCmd.Command("ShowPermission", `Show information about a single permission.`)
	ShowPermissionRunner.Flag(`id`, ``).Required().StringVar(&ShowPermissionRunner.id)
	registry[ShowPermissionCmd.FullCommand()] = ShowPermissionRunner
}

/****** PlacementGroup ******/

type CreatePlacementGroupPlacementGroupRunner struct {
	placementGroupCloudHref   string
	placementGroupDescription string
	placementGroupName        string
}

func (r *CreatePlacementGroupPlacementGroupRunner) Run(c *Client) (interface{}, error) {

	/** Handle placementGroup parameter **/
	var placementGroup PlacementGroupParam

	// Load JSON if provided
	if len(r.placementGroupJson) > 0 {
		if err := Json.Unmarshal(r.placementGroupJson, &placementGroup); err != nil {
			return nil, fmt.Errorf("Could not load placementGroup JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.placementGroupcloudHref) > 0 {
		placementGroup.placementGroup.cloudHref = r.placementGroupcloudHref
	}

	if len(r.placementGroupdescription) > 0 {
		placementGroup.placementGroup.description = r.placementGroupdescription
	}

	if len(r.placementGroupname) > 0 {
		placementGroup.placementGroup.name = r.placementGroupname
	}

	return c.CreatePlacementGroup(&placementGroup)
}

type DestroyPlacementGroupPlacementGroupRunner struct {
	id string
}

func (r *DestroyPlacementGroupPlacementGroupRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyPlacementGroup(r.id)
}

type IndexPlacementGroupsPlacementGroupRunner struct {
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexPlacementGroupsPlacementGroupRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexPlacementGroups(filter, r.view)
}

type ShowPlacementGroupPlacementGroupRunner struct {
	id   string
	view string
}

func (r *ShowPlacementGroupPlacementGroupRunner) Run(c *Client) (interface{}, error) {
	return c.ShowPlacementGroup(r.id, r.view)
}

// Register all PlacementGroup actions
func registerPlacementGroupCmds(app *kingpin.Application) {
	resCmd := app.Cmd("PlacementGroup", ``)

	CreatePlacementGroupRunner := new(CreatePlacementGroupPlacementGroupRunner)
	CreatePlacementGroupCmd := resCmd.Command("CreatePlacementGroup", `Creates a PlacementGroup.`)
	CreatePlacementGroupRunner.Flag(`placementGroup.cloudHref`, `The Href of the Cloud in which the PlacementGroup should be created. Note: This feature is not supported for all clouds.`).Required().StringVar(&CreatePlacementGroupRunner.placementGroupCloudHref)
	CreatePlacementGroupRunner.Flag(`placementGroup.description`, `The description of the Placement Group to be created.`).StringVar(&CreatePlacementGroupRunner.placementGroupDescription)
	CreatePlacementGroupRunner.Flag(`placementGroup.name`, `The name of the Placement Group to be created.`).Required().StringVar(&CreatePlacementGroupRunner.placementGroupName)
	registry[CreatePlacementGroupCmd.FullCommand()] = CreatePlacementGroupRunner

	DestroyPlacementGroupRunner := new(DestroyPlacementGroupPlacementGroupRunner)
	DestroyPlacementGroupCmd := resCmd.Command("DestroyPlacementGroup", `Destroys a PlacementGroup.`)
	DestroyPlacementGroupRunner.Flag(`id`, ``).Required().StringVar(&DestroyPlacementGroupRunner.id)
	registry[DestroyPlacementGroupCmd.FullCommand()] = DestroyPlacementGroupRunner

	IndexPlacementGroupsRunner := new(IndexPlacementGroupsPlacementGroupRunner)
	IndexPlacementGroupsCmd := resCmd.Command("IndexPlacementGroups", `Lists all PlacementGroups in an account.`)
	IndexPlacementGroupsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexPlacementGroupsRunner.filterPos).StringsVar(&IndexPlacementGroupsRunner.filter)
	IndexPlacementGroupsRunner.Flag(`view`, ``).StringVar(&IndexPlacementGroupsRunner.view)
	registry[IndexPlacementGroupsCmd.FullCommand()] = IndexPlacementGroupsRunner

	ShowPlacementGroupRunner := new(ShowPlacementGroupPlacementGroupRunner)
	ShowPlacementGroupCmd := resCmd.Command("ShowPlacementGroup", `Shows information about a single PlacementGroup.`)
	ShowPlacementGroupRunner.Flag(`id`, ``).Required().StringVar(&ShowPlacementGroupRunner.id)
	ShowPlacementGroupRunner.Flag(`view`, ``).StringVar(&ShowPlacementGroupRunner.view)
	registry[ShowPlacementGroupCmd.FullCommand()] = ShowPlacementGroupRunner
}

/****** Preference ******/

type DestroyPreferencePreferenceRunner struct {
	id string
}

func (r *DestroyPreferencePreferenceRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyPreference(r.id)
}

type IndexPreferencesPreferenceRunner struct {
	filter    []string
	filterPos []string
}

func (r *IndexPreferencesPreferenceRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexPreferences(filter)
}

type ShowPreferencePreferenceRunner struct {
	id string
}

func (r *ShowPreferencePreferenceRunner) Run(c *Client) (interface{}, error) {
	return c.ShowPreference(r.id)
}

type UpdatePreferencePreferenceRunner struct {
	id                 string
	preferenceContents string
}

func (r *UpdatePreferencePreferenceRunner) Run(c *Client) (interface{}, error) {

	/** Handle preference parameter **/
	var preference PreferenceParam

	// Load JSON if provided
	if len(r.preferenceJson) > 0 {
		if err := Json.Unmarshal(r.preferenceJson, &preference); err != nil {
			return nil, fmt.Errorf("Could not load preference JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.preferencecontents) > 0 {
		preference.preference.contents = r.preferencecontents
	}

	return c.UpdatePreference(r.id, &preference)
}

// Register all Preference actions
func registerPreferenceCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Preference", `A Preference is a user and account-specific setting. Preferences are used in many part of the RightScale platform and can be used for custom purposes if desired.`)

	DestroyPreferenceRunner := new(DestroyPreferencePreferenceRunner)
	DestroyPreferenceCmd := resCmd.Command("DestroyPreference", `Deletes the given preference.`)
	DestroyPreferenceRunner.Flag(`id`, ``).Required().StringVar(&DestroyPreferenceRunner.id)
	registry[DestroyPreferenceCmd.FullCommand()] = DestroyPreferenceRunner

	IndexPreferencesRunner := new(IndexPreferencesPreferenceRunner)
	IndexPreferencesCmd := resCmd.Command("IndexPreferences", `Lists all preferences.`)
	IndexPreferencesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexPreferencesRunner.filterPos).StringsVar(&IndexPreferencesRunner.filter)
	registry[IndexPreferencesCmd.FullCommand()] = IndexPreferencesRunner

	ShowPreferenceRunner := new(ShowPreferencePreferenceRunner)
	ShowPreferenceCmd := resCmd.Command("ShowPreference", `Shows a single preference.`)
	ShowPreferenceRunner.Flag(`id`, ``).Required().StringVar(&ShowPreferenceRunner.id)
	registry[ShowPreferenceCmd.FullCommand()] = ShowPreferenceRunner

	UpdatePreferenceRunner := new(UpdatePreferencePreferenceRunner)
	UpdatePreferenceCmd := resCmd.Command("UpdatePreference", `If 'id' is known, updates preference with given contents.`)
	UpdatePreferenceRunner.Flag(`id`, ``).Required().StringVar(&UpdatePreferenceRunner.id)
	UpdatePreferenceRunner.Flag(`preference.contents`, `The updated contents for the Preference.`).Required().StringVar(&UpdatePreferenceRunner.preferenceContents)
	registry[UpdatePreferenceCmd.FullCommand()] = UpdatePreferenceRunner
}

/****** Publication ******/

type ImportPublicationPublicationRunner struct {
	id string
}

func (r *ImportPublicationPublicationRunner) Run(c *Client) (interface{}, error) {
	return c.ImportPublication(r.id)
}

type IndexPublicationsPublicationRunner struct {
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexPublicationsPublicationRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexPublications(filter, r.view)
}

type ShowPublicationPublicationRunner struct {
	id   string
	view string
}

func (r *ShowPublicationPublicationRunner) Run(c *Client) (interface{}, error) {
	return c.ShowPublication(r.id, r.view)
}

// Register all Publication actions
func registerPublicationCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Publication", `A Publication is a revisioned component shared with a set of Account Groups.`)

	ImportPublicationRunner := new(ImportPublicationPublicationRunner)
	ImportPublicationCmd := resCmd.Command("ImportPublication", `Imports the given publication and its subordinates to this account.`)
	ImportPublicationRunner.Flag(`id`, ``).Required().StringVar(&ImportPublicationRunner.id)
	registry[ImportPublicationCmd.FullCommand()] = ImportPublicationRunner

	IndexPublicationsRunner := new(IndexPublicationsPublicationRunner)
	IndexPublicationsCmd := resCmd.Command("IndexPublications", `Lists the publications available to this account. Only non-HEAD revisions are possible.`)
	IndexPublicationsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexPublicationsRunner.filterPos).StringsVar(&IndexPublicationsRunner.filter)
	IndexPublicationsRunner.Flag(`view`, ``).StringVar(&IndexPublicationsRunner.view)
	registry[IndexPublicationsCmd.FullCommand()] = IndexPublicationsRunner

	ShowPublicationRunner := new(ShowPublicationPublicationRunner)
	ShowPublicationCmd := resCmd.Command("ShowPublication", `Show information about a single publication. Only non-HEAD revisions are possible.`)
	ShowPublicationRunner.Flag(`id`, ``).Required().StringVar(&ShowPublicationRunner.id)
	ShowPublicationRunner.Flag(`view`, ``).StringVar(&ShowPublicationRunner.view)
	registry[ShowPublicationCmd.FullCommand()] = ShowPublicationRunner
}

/****** PublicationLineage ******/

type ShowPublicationLineagePublicationLineageRunner struct {
	id   string
	view string
}

func (r *ShowPublicationLineagePublicationLineageRunner) Run(c *Client) (interface{}, error) {
	return c.ShowPublicationLineage(r.id, r.view)
}

// Register all PublicationLineage actions
func registerPublicationLineageCmds(app *kingpin.Application) {
	resCmd := app.Cmd("PublicationLineage", `A Publication Lineage contains lineage information for a Publication in the MultiCloudMarketplace.`)

	ShowPublicationLineageRunner := new(ShowPublicationLineagePublicationLineageRunner)
	ShowPublicationLineageCmd := resCmd.Command("ShowPublicationLineage", `Show information about a single publication lineage. Only non-HEAD revisions are possible.`)
	ShowPublicationLineageRunner.Flag(`id`, ``).Required().StringVar(&ShowPublicationLineageRunner.id)
	ShowPublicationLineageRunner.Flag(`view`, ``).StringVar(&ShowPublicationLineageRunner.view)
	registry[ShowPublicationLineageCmd.FullCommand()] = ShowPublicationLineageRunner
}

/****** RecurringVolumeAttachment ******/

type CreateRecurringVolumeAttachmentRecurringVolumeAttachmentRunner struct {
	cloudId                               string
	recurringVolumeAttachmentDevice       string
	recurringVolumeAttachmentRunnableHref string
	recurringVolumeAttachmentStorageHref  string
}

func (r *CreateRecurringVolumeAttachmentRecurringVolumeAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle recurringVolumeAttachment parameter **/
	var recurringVolumeAttachment RecurringVolumeAttachmentParam

	// Load JSON if provided
	if len(r.recurringVolumeAttachmentJson) > 0 {
		if err := Json.Unmarshal(r.recurringVolumeAttachmentJson, &recurringVolumeAttachment); err != nil {
			return nil, fmt.Errorf("Could not load recurringVolumeAttachment JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.recurringVolumeAttachmentdevice) > 0 {
		recurringVolumeAttachment.recurringVolumeAttachment.device = r.recurringVolumeAttachmentdevice
	}

	if len(r.recurringVolumeAttachmentrunnableHref) > 0 {
		recurringVolumeAttachment.recurringVolumeAttachment.runnableHref = r.recurringVolumeAttachmentrunnableHref
	}

	if len(r.recurringVolumeAttachmentstorageHref) > 0 {
		recurringVolumeAttachment.recurringVolumeAttachment.storageHref = r.recurringVolumeAttachmentstorageHref
	}

	return c.CreateRecurringVolumeAttachment(r.cloudId, &recurringVolumeAttachment)
}

type DestroyRecurringVolumeAttachmentRecurringVolumeAttachmentRunner struct {
	cloudId string
	id      string
}

func (r *DestroyRecurringVolumeAttachmentRecurringVolumeAttachmentRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyRecurringVolumeAttachment(r.cloudId, r.id)
}

type IndexRecurringVolumeAttachmentsRecurringVolumeAttachmentRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexRecurringVolumeAttachmentsRecurringVolumeAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexRecurringVolumeAttachments(r.cloudId, filter, r.view)
}

type ShowRecurringVolumeAttachmentRecurringVolumeAttachmentRunner struct {
	cloudId string
	id      string
	view    string
}

func (r *ShowRecurringVolumeAttachmentRecurringVolumeAttachmentRunner) Run(c *Client) (interface{}, error) {
	return c.ShowRecurringVolumeAttachment(r.cloudId, r.id, r.view)
}

// Register all RecurringVolumeAttachment actions
func registerRecurringVolumeAttachmentCmds(app *kingpin.Application) {
	resCmd := app.Cmd("RecurringVolumeAttachment", `A RecurringVolumeAttachment specifies a Volume/VolumeSnapshot to attach to a Server/ServerArray the next time an instance is launched.`)

	CreateRecurringVolumeAttachmentRunner := new(CreateRecurringVolumeAttachmentRecurringVolumeAttachmentRunner)
	CreateRecurringVolumeAttachmentCmd := resCmd.Command("CreateRecurringVolumeAttachment", `Creates a new recurring volume attachment.`)
	CreateRecurringVolumeAttachmentRunner.Flag(`cloudId`, ``).Required().StringVar(&CreateRecurringVolumeAttachmentRunner.cloudId)
	CreateRecurringVolumeAttachmentRunner.Flag(`recurringVolumeAttachment.device`, `The device location where the volume or volume snapshot will be mounted. Value must be of format /dev/xvd[bcefghij]. This is not reliable and will be deprecated.`).Required().StringVar(&CreateRecurringVolumeAttachmentRunner.recurringVolumeAttachmentDevice)
	CreateRecurringVolumeAttachmentRunner.Flag(`recurringVolumeAttachment.runnableHref`, `The href of the server or server array to attach to.`).Required().StringVar(&CreateRecurringVolumeAttachmentRunner.recurringVolumeAttachmentRunnableHref)
	CreateRecurringVolumeAttachmentRunner.Flag(`recurringVolumeAttachment.storageHref`, `The href of the volume or volume snapshot to be attached on launch of a next instance.`).Required().StringVar(&CreateRecurringVolumeAttachmentRunner.recurringVolumeAttachmentStorageHref)
	registry[CreateRecurringVolumeAttachmentCmd.FullCommand()] = CreateRecurringVolumeAttachmentRunner

	DestroyRecurringVolumeAttachmentRunner := new(DestroyRecurringVolumeAttachmentRecurringVolumeAttachmentRunner)
	DestroyRecurringVolumeAttachmentCmd := resCmd.Command("DestroyRecurringVolumeAttachment", `Deletes a given recurring volume attachment.`)
	DestroyRecurringVolumeAttachmentRunner.Flag(`cloudId`, ``).Required().StringVar(&DestroyRecurringVolumeAttachmentRunner.cloudId)
	DestroyRecurringVolumeAttachmentRunner.Flag(`id`, ``).Required().StringVar(&DestroyRecurringVolumeAttachmentRunner.id)
	registry[DestroyRecurringVolumeAttachmentCmd.FullCommand()] = DestroyRecurringVolumeAttachmentRunner

	IndexRecurringVolumeAttachmentsRunner := new(IndexRecurringVolumeAttachmentsRecurringVolumeAttachmentRunner)
	IndexRecurringVolumeAttachmentsCmd := resCmd.Command("IndexRecurringVolumeAttachments", `Lists all recurring volume attachments.`)
	IndexRecurringVolumeAttachmentsRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexRecurringVolumeAttachmentsRunner.cloudId)
	IndexRecurringVolumeAttachmentsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexRecurringVolumeAttachmentsRunner.filterPos).StringsVar(&IndexRecurringVolumeAttachmentsRunner.filter)
	IndexRecurringVolumeAttachmentsRunner.Flag(`view`, ``).StringVar(&IndexRecurringVolumeAttachmentsRunner.view)
	registry[IndexRecurringVolumeAttachmentsCmd.FullCommand()] = IndexRecurringVolumeAttachmentsRunner

	ShowRecurringVolumeAttachmentRunner := new(ShowRecurringVolumeAttachmentRecurringVolumeAttachmentRunner)
	ShowRecurringVolumeAttachmentCmd := resCmd.Command("ShowRecurringVolumeAttachment", `Displays information about a single recurring volume attachment.`)
	ShowRecurringVolumeAttachmentRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowRecurringVolumeAttachmentRunner.cloudId)
	ShowRecurringVolumeAttachmentRunner.Flag(`id`, ``).Required().StringVar(&ShowRecurringVolumeAttachmentRunner.id)
	ShowRecurringVolumeAttachmentRunner.Flag(`view`, ``).StringVar(&ShowRecurringVolumeAttachmentRunner.view)
	registry[ShowRecurringVolumeAttachmentCmd.FullCommand()] = ShowRecurringVolumeAttachmentRunner
}

/****** Repository ******/

type CookbookImportRepositoryRepositoryRunner struct {
	assetHrefs                []string
	assetHrefsPos             []string
	follow                    string
	id                        string
	namespace                 string
	repositoryCommitReference string
	withDependencies          string
}

func (r *CookbookImportRepositoryRepositoryRunner) Run(c *Client) (interface{}, error) {

	/** Handle assetHrefs parameter **/
	var assetHrefs []string

	for i, v := range r.assetHrefs {
		pos, err := strconv.Atoi(r.assetHrefsPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for assetHrefs array", r.assetHrefsPos[i])
		}
		assetHrefs[pos] = v
	}

	return c.CookbookImportRepository(assetHrefs, r.follow, r.id, r.namespace, r.repositoryCommitReference, r.withDependencies)
}

type CookbookImportPreviewRepositoryRepositoryRunner struct {
	assetHrefs    []string
	assetHrefsPos []string
	id            string
	namespace     string
}

func (r *CookbookImportPreviewRepositoryRepositoryRunner) Run(c *Client) (interface{}, error) {

	/** Handle assetHrefs parameter **/
	var assetHrefs []string

	for i, v := range r.assetHrefs {
		pos, err := strconv.Atoi(r.assetHrefsPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for assetHrefs array", r.assetHrefsPos[i])
		}
		assetHrefs[pos] = v
	}

	return c.CookbookImportPreviewRepository(assetHrefs, r.id, r.namespace)
}

type CreateRepositoryRepositoryRunner struct {
	repositoryAssetPathsItem      []string
	repositoryAssetPathsItemPos   []string
	repositoryAutoImport          string
	repositoryCommitReference     string
	repositoryCredentialsPassword string
	repositoryCredentialsSshKey   string
	repositoryCredentialsUsername string
	repositoryDescription         string
	repositoryName                string
	repositorySource              string
	repositorySourceType          string
}

func (r *CreateRepositoryRepositoryRunner) Run(c *Client) (interface{}, error) {

	/** Handle repository parameter **/
	var repository RepositoryParam

	// Load JSON if provided
	if len(r.repositoryJson) > 0 {
		if err := Json.Unmarshal(r.repositoryJson, &repository); err != nil {
			return nil, fmt.Errorf("Could not load repository JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.repositoryassetPathscookbooksitem) > 0 {
		repository.repository.assetPaths.cookbooks.item = r.repositoryassetPathscookbooksitem
	}

	if len(r.repositoryautoImport) > 0 {
		repository.repository.autoImport = r.repositoryautoImport
	}

	if len(r.repositorycommitReference) > 0 {
		repository.repository.commitReference = r.repositorycommitReference
	}

	if len(r.repositorycredentialspassword) > 0 {
		repository.repository.credentials.password = r.repositorycredentialspassword
	}

	if len(r.repositorycredentialssshKey) > 0 {
		repository.repository.credentials.sshKey = r.repositorycredentialssshKey
	}

	if len(r.repositorycredentialsusername) > 0 {
		repository.repository.credentials.username = r.repositorycredentialsusername
	}

	if len(r.repositorydescription) > 0 {
		repository.repository.description = r.repositorydescription
	}

	if len(r.repositoryname) > 0 {
		repository.repository.name = r.repositoryname
	}

	if len(r.repositorysource) > 0 {
		repository.repository.source = r.repositorysource
	}

	if len(r.repositorysourceType) > 0 {
		repository.repository.sourceType = r.repositorysourceType
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxrepositoryassetPathscookbooksitemPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.repositoryassetPathscookbooksitemPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for repository.assetPaths.cookbooks.item field of cookbooks array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of repository.assetPaths.cookbooks.item field of cookbooks array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxrepositoryassetPathscookbooksitemPos {
			maxrepositoryassetPathscookbooksitemPos = pos
		}
	}
	if len(r.repositoryassetPathscookbooksitem) != maxrepositoryassetPathscookbooksitemPos {
		return nil, fmt.Errof("Invalid flags for repository.assetPaths.cookbooks.item field of cookbooks array, %s were defined but max position is %s",
			len(r.repositoryassetPathscookbooksitem), maxrepositoryassetPathscookbooksitemPos)
	}
	if maxrepositoryassetPathscookbooksitemPos > -1 {
		repositoryassetPathscookbooks := make([]*string, maxrepositoryassetPathscookbooksitemPos+1)
		for i := 0; i < maxrepositoryassetPathscookbooksitemPos+1; i++ {
			repositoryassetPathscookbooks[i] = string{
			//TBD
			}
		}
		repository.assetPaths.cookbooks = repositoryassetPathscookbooks
	}

	return c.CreateRepository(&repository)
}

type DestroyRepositoryRepositoryRunner struct {
	id string
}

func (r *DestroyRepositoryRepositoryRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyRepository(r.id)
}

type IndexRepositoriesRepositoryRunner struct {
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexRepositoriesRepositoryRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexRepositories(filter, r.view)
}

type RefetchRepositoryRepositoryRunner struct {
	autoImport string
	id         string
}

func (r *RefetchRepositoryRepositoryRunner) Run(c *Client) (interface{}, error) {
	return c.RefetchRepository(r.autoImport, r.id)
}

type ResolveRepositoryRepositoryRunner struct {
	importedCookbookName    []string
	importedCookbookNamePos []string
}

func (r *ResolveRepositoryRepositoryRunner) Run(c *Client) (interface{}, error) {

	/** Handle importedCookbookName parameter **/
	var importedCookbookName []string

	for i, v := range r.importedCookbookName {
		pos, err := strconv.Atoi(r.importedCookbookNamePos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for importedCookbookName array", r.importedCookbookNamePos[i])
		}
		importedCookbookName[pos] = v
	}

	return c.ResolveRepository(importedCookbookName)
}

type ShowRepositoryRepositoryRunner struct {
	id   string
	view string
}

func (r *ShowRepositoryRepositoryRunner) Run(c *Client) (interface{}, error) {
	return c.ShowRepository(r.id, r.view)
}

type UpdateRepositoryRepositoryRunner struct {
	id                            string
	repositoryAssetPathsItem      []string
	repositoryAssetPathsItemPos   []string
	repositoryCommitReference     string
	repositoryCredentialsPassword string
	repositoryCredentialsSshKey   string
	repositoryCredentialsUsername string
	repositoryDescription         string
	repositoryName                string
	repositorySource              string
	repositorySourceType          string
}

func (r *UpdateRepositoryRepositoryRunner) Run(c *Client) (interface{}, error) {

	/** Handle repository parameter **/
	var repository RepositoryParam2

	// Load JSON if provided
	if len(r.repositoryJson) > 0 {
		if err := Json.Unmarshal(r.repositoryJson, &repository); err != nil {
			return nil, fmt.Errorf("Could not load repository JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.repositoryassetPathscookbooksitem) > 0 {
		repository.repository.assetPaths.cookbooks.item = r.repositoryassetPathscookbooksitem
	}

	if len(r.repositorycommitReference) > 0 {
		repository.repository.commitReference = r.repositorycommitReference
	}

	if len(r.repositorycredentialspassword) > 0 {
		repository.repository.credentials.password = r.repositorycredentialspassword
	}

	if len(r.repositorycredentialssshKey) > 0 {
		repository.repository.credentials.sshKey = r.repositorycredentialssshKey
	}

	if len(r.repositorycredentialsusername) > 0 {
		repository.repository.credentials.username = r.repositorycredentialsusername
	}

	if len(r.repositorydescription) > 0 {
		repository.repository.description = r.repositorydescription
	}

	if len(r.repositoryname) > 0 {
		repository.repository.name = r.repositoryname
	}

	if len(r.repositorysource) > 0 {
		repository.repository.source = r.repositorysource
	}

	if len(r.repositorysourceType) > 0 {
		repository.repository.sourceType = r.repositorysourceType
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxrepositoryassetPathscookbooksitemPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.repositoryassetPathscookbooksitemPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for repository.assetPaths.cookbooks.item field of cookbooks array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of repository.assetPaths.cookbooks.item field of cookbooks array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxrepositoryassetPathscookbooksitemPos {
			maxrepositoryassetPathscookbooksitemPos = pos
		}
	}
	if len(r.repositoryassetPathscookbooksitem) != maxrepositoryassetPathscookbooksitemPos {
		return nil, fmt.Errof("Invalid flags for repository.assetPaths.cookbooks.item field of cookbooks array, %s were defined but max position is %s",
			len(r.repositoryassetPathscookbooksitem), maxrepositoryassetPathscookbooksitemPos)
	}
	if maxrepositoryassetPathscookbooksitemPos > -1 {
		repositoryassetPathscookbooks := make([]*string, maxrepositoryassetPathscookbooksitemPos+1)
		for i := 0; i < maxrepositoryassetPathscookbooksitemPos+1; i++ {
			repositoryassetPathscookbooks[i] = string{
			//TBD
			}
		}
		repository.assetPaths.cookbooks = repositoryassetPathscookbooks
	}

	return c.UpdateRepository(r.id, &repository)
}

// Register all Repository actions
func registerRepositoryCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Repository", `A Repository is a location from which you can download and import design objects such as Chef cookbooks. Using this resource you can add and modify repository information and import assets discovered in the repository.`)

	CookbookImportRepositoryRunner := new(CookbookImportRepositoryRepositoryRunner)
	CookbookImportRepositoryCmd := resCmd.Command("CookbookImportRepository", `Performs a Cookbook import, which allows you to use the specified cookbooks in your design objects.`)
	CookbookImportRepositoryRunner.FlagPattern(`assetHrefs\.(\d+)`, `Hrefs of the assets that should be imported.`).Required().Capture(&CookbookImportRepositoryRunner.assetHrefsPos).StringsVar(&CookbookImportRepositoryRunner.assetHrefs)
	CookbookImportRepositoryRunner.Flag(`follow`, `A flag indicating whether imported cookbooks should be followed.`).StringVar(&CookbookImportRepositoryRunner.follow)
	CookbookImportRepositoryRunner.Flag(`id`, ``).Required().StringVar(&CookbookImportRepositoryRunner.id)
	CookbookImportRepositoryRunner.Flag(`namespace`, `The namespace to import into.`).StringVar(&CookbookImportRepositoryRunner.namespace)
	CookbookImportRepositoryRunner.Flag(`repositoryCommitReference`, `Optional commit reference indicating last succeeded commit. Must match the Repository's fetch_status.succeeded_commit attribute or the import will not be performed.`).StringVar(&CookbookImportRepositoryRunner.repositoryCommitReference)
	CookbookImportRepositoryRunner.Flag(`withDependencies`, `A flag indicating whether dependencies should automatically be imported.`).StringVar(&CookbookImportRepositoryRunner.withDependencies)
	registry[CookbookImportRepositoryCmd.FullCommand()] = CookbookImportRepositoryRunner

	CookbookImportPreviewRepositoryRunner := new(CookbookImportPreviewRepositoryRepositoryRunner)
	CookbookImportPreviewRepositoryCmd := resCmd.Command("CookbookImportPreviewRepository", `Retrieves a preview of the effects of a Cookbook import.`)
	CookbookImportPreviewRepositoryRunner.FlagPattern(`assetHrefs\.(\d+)`, `Hrefs of the assets that should be imported.`).Required().Capture(&CookbookImportPreviewRepositoryRunner.assetHrefsPos).StringsVar(&CookbookImportPreviewRepositoryRunner.assetHrefs)
	CookbookImportPreviewRepositoryRunner.Flag(`id`, ``).Required().StringVar(&CookbookImportPreviewRepositoryRunner.id)
	CookbookImportPreviewRepositoryRunner.Flag(`namespace`, `The namespace to import into.`).Required().StringVar(&CookbookImportPreviewRepositoryRunner.namespace)
	registry[CookbookImportPreviewRepositoryCmd.FullCommand()] = CookbookImportPreviewRepositoryRunner

	CreateRepositoryRunner := new(CreateRepositoryRepositoryRunner)
	CreateRepositoryCmd := resCmd.Command("CreateRepository", `Creates a Repository.`)
	CreateRepositoryRunner.FlagPattern(`repository\.assetPaths\.item\.(\d+)`, `The cookbook paths for the repository`).Capture(&CreateRepositoryRunner.repositoryAssetPathsItemPos).StringsVar(&CreateRepositoryRunner.repositoryAssetPathsItem)
	CreateRepositoryRunner.Flag(`repository.autoImport`, `Whether cookbooks should automatically be imported upon repository creation.`).StringVar(&CreateRepositoryRunner.repositoryAutoImport)
	CreateRepositoryRunner.Flag(`repository.commitReference`, `The revision for the repository`).StringVar(&CreateRepositoryRunner.repositoryCommitReference)
	CreateRepositoryRunner.Flag(`repository.credentials.password`, `The password, or credential, for the repository (only valid for svn or download repositories).`).StringVar(&CreateRepositoryRunner.repositoryCredentialsPassword)
	CreateRepositoryRunner.Flag(`repository.credentials.sshKey`, `The SSH key, or credential, for the repository (only valid for git repositories).`).StringVar(&CreateRepositoryRunner.repositoryCredentialsSshKey)
	CreateRepositoryRunner.Flag(`repository.credentials.username`, `The user name, or credential, for the repository (only valid for svn or download repositories).`).StringVar(&CreateRepositoryRunner.repositoryCredentialsUsername)
	CreateRepositoryRunner.Flag(`repository.description`, `The description for the repository.`).StringVar(&CreateRepositoryRunner.repositoryDescription)
	CreateRepositoryRunner.Flag(`repository.name`, `The repository name.`).Required().StringVar(&CreateRepositoryRunner.repositoryName)
	CreateRepositoryRunner.Flag(`repository.source`, `The URL for the repository.`).Required().StringVar(&CreateRepositoryRunner.repositorySource)
	CreateRepositoryRunner.Flag(`repository.sourceType`, `The source type for the repository.`).Required().StringVar(&CreateRepositoryRunner.repositorySourceType)
	registry[CreateRepositoryCmd.FullCommand()] = CreateRepositoryRunner

	DestroyRepositoryRunner := new(DestroyRepositoryRepositoryRunner)
	DestroyRepositoryCmd := resCmd.Command("DestroyRepository", `Deletes the specified Repositories.`)
	DestroyRepositoryRunner.Flag(`id`, ``).Required().StringVar(&DestroyRepositoryRunner.id)
	registry[DestroyRepositoryCmd.FullCommand()] = DestroyRepositoryRunner

	IndexRepositoriesRunner := new(IndexRepositoriesRepositoryRunner)
	IndexRepositoriesCmd := resCmd.Command("IndexRepositories", `Lists all Repositories for this Account.`)
	IndexRepositoriesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexRepositoriesRunner.filterPos).StringsVar(&IndexRepositoriesRunner.filter)
	IndexRepositoriesRunner.Flag(`view`, ``).StringVar(&IndexRepositoriesRunner.view)
	registry[IndexRepositoriesCmd.FullCommand()] = IndexRepositoriesRunner

	RefetchRepositoryRunner := new(RefetchRepositoryRepositoryRunner)
	RefetchRepositoryCmd := resCmd.Command("RefetchRepository", `Refetches all RepositoryAssets associated with the Repository.`)
	RefetchRepositoryRunner.Flag(`autoImport`, `Whether cookbooks should automatically be imported after repositories are fetched.`).StringVar(&RefetchRepositoryRunner.autoImport)
	RefetchRepositoryRunner.Flag(`id`, ``).Required().StringVar(&RefetchRepositoryRunner.id)
	registry[RefetchRepositoryCmd.FullCommand()] = RefetchRepositoryRunner

	ResolveRepositoryRunner := new(ResolveRepositoryRepositoryRunner)
	ResolveRepositoryCmd := resCmd.Command("ResolveRepository", `Show a list of repositories that have imported cookbooks with the given names.`)
	ResolveRepositoryRunner.FlagPattern(`importedCookbookName\.(\d+)`, `A list of cookbook names that were imported by the repository.`).Capture(&ResolveRepositoryRunner.importedCookbookNamePos).StringsVar(&ResolveRepositoryRunner.importedCookbookName)
	registry[ResolveRepositoryCmd.FullCommand()] = ResolveRepositoryRunner

	ShowRepositoryRunner := new(ShowRepositoryRepositoryRunner)
	ShowRepositoryCmd := resCmd.Command("ShowRepository", `Shows a specified Repository.`)
	ShowRepositoryRunner.Flag(`id`, ``).Required().StringVar(&ShowRepositoryRunner.id)
	ShowRepositoryRunner.Flag(`view`, ``).StringVar(&ShowRepositoryRunner.view)
	registry[ShowRepositoryCmd.FullCommand()] = ShowRepositoryRunner

	UpdateRepositoryRunner := new(UpdateRepositoryRepositoryRunner)
	UpdateRepositoryCmd := resCmd.Command("UpdateRepository", `Updates a specified Repository.`)
	UpdateRepositoryRunner.Flag(`id`, ``).Required().StringVar(&UpdateRepositoryRunner.id)
	UpdateRepositoryRunner.FlagPattern(`repository\.assetPaths\.item\.(\d+)`, `The updated cookbook paths for the repository`).Capture(&UpdateRepositoryRunner.repositoryAssetPathsItemPos).StringsVar(&UpdateRepositoryRunner.repositoryAssetPathsItem)
	UpdateRepositoryRunner.Flag(`repository.commitReference`, `The updated commit reference (tag, branch, revision...) for the repository`).StringVar(&UpdateRepositoryRunner.repositoryCommitReference)
	UpdateRepositoryRunner.Flag(`repository.credentials.password`, `The updated password, or credential, for the repository (only valid for svn or download repositories).`).StringVar(&UpdateRepositoryRunner.repositoryCredentialsPassword)
	UpdateRepositoryRunner.Flag(`repository.credentials.sshKey`, `The updated SSH key for the repository (only valid for git repositories).`).StringVar(&UpdateRepositoryRunner.repositoryCredentialsSshKey)
	UpdateRepositoryRunner.Flag(`repository.credentials.username`, `The updated user name, or credential, for the repository (only valid for svn or download repositories).`).StringVar(&UpdateRepositoryRunner.repositoryCredentialsUsername)
	UpdateRepositoryRunner.Flag(`repository.description`, `The updated description for the repository.`).StringVar(&UpdateRepositoryRunner.repositoryDescription)
	UpdateRepositoryRunner.Flag(`repository.name`, `The updated repository name.`).StringVar(&UpdateRepositoryRunner.repositoryName)
	UpdateRepositoryRunner.Flag(`repository.source`, `The updated URL for the repository.`).StringVar(&UpdateRepositoryRunner.repositorySource)
	UpdateRepositoryRunner.Flag(`repository.sourceType`, `The updated source type for the repository.`).StringVar(&UpdateRepositoryRunner.repositorySourceType)
	registry[UpdateRepositoryCmd.FullCommand()] = UpdateRepositoryRunner
}

/****** RepositoryAsset ******/

type IndexRepositoryAssetsRepositoryAssetRunner struct {
	repositoryId string
	view         string
}

func (r *IndexRepositoryAssetsRepositoryAssetRunner) Run(c *Client) (interface{}, error) {
	return c.IndexRepositoryAssets(r.repositoryId, r.view)
}

type ShowRepositoryAssetRepositoryAssetRunner struct {
	id           string
	repositoryId string
	view         string
}

func (r *ShowRepositoryAssetRepositoryAssetRunner) Run(c *Client) (interface{}, error) {
	return c.ShowRepositoryAsset(r.id, r.repositoryId, r.view)
}

// Register all RepositoryAsset actions
func registerRepositoryAssetCmds(app *kingpin.Application) {
	resCmd := app.Cmd("RepositoryAsset", `A RepositoryAsset represents an item discovered in a Repository`)

	IndexRepositoryAssetsRunner := new(IndexRepositoryAssetsRepositoryAssetRunner)
	IndexRepositoryAssetsCmd := resCmd.Command("IndexRepositoryAssets", `List a repository's current assets.`)
	IndexRepositoryAssetsRunner.Flag(`repositoryId`, ``).Required().StringVar(&IndexRepositoryAssetsRunner.repositoryId)
	IndexRepositoryAssetsRunner.Flag(`view`, ``).StringVar(&IndexRepositoryAssetsRunner.view)
	registry[IndexRepositoryAssetsCmd.FullCommand()] = IndexRepositoryAssetsRunner

	ShowRepositoryAssetRunner := new(ShowRepositoryAssetRepositoryAssetRunner)
	ShowRepositoryAssetCmd := resCmd.Command("ShowRepositoryAsset", `Show information about a single asset.`)
	ShowRepositoryAssetRunner.Flag(`id`, ``).Required().StringVar(&ShowRepositoryAssetRunner.id)
	ShowRepositoryAssetRunner.Flag(`repositoryId`, ``).Required().StringVar(&ShowRepositoryAssetRunner.repositoryId)
	ShowRepositoryAssetRunner.Flag(`view`, ``).StringVar(&ShowRepositoryAssetRunner.view)
	registry[ShowRepositoryAssetCmd.FullCommand()] = ShowRepositoryAssetRunner
}

/****** RightScript ******/

type CommitRightScriptRightScriptRunner struct {
	id                       string
	rightScriptCommitMessage string
}

func (r *CommitRightScriptRightScriptRunner) Run(c *Client) (interface{}, error) {

	/** Handle rightScript parameter **/
	var rightScript RightScriptParam

	// Load JSON if provided
	if len(r.rightScriptJson) > 0 {
		if err := Json.Unmarshal(r.rightScriptJson, &rightScript); err != nil {
			return nil, fmt.Errorf("Could not load rightScript JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.rightScriptcommitMessage) > 0 {
		rightScript.rightScript.commitMessage = r.rightScriptcommitMessage
	}

	return c.CommitRightScript(r.id, &rightScript)
}

type IndexRightScriptsRightScriptRunner struct {
	filter     []string
	filterPos  []string
	latestOnly string
	view       string
}

func (r *IndexRightScriptsRightScriptRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexRightScripts(filter, r.latestOnly, r.view)
}

type ShowRightScriptRightScriptRunner struct {
	id string
}

func (r *ShowRightScriptRightScriptRunner) Run(c *Client) (interface{}, error) {
	return c.ShowRightScript(r.id)
}

type ShowSourceRightScriptRightScriptRunner struct {
	id string
}

func (r *ShowSourceRightScriptRightScriptRunner) Run(c *Client) (interface{}, error) {
	return c.ShowSourceRightScript(r.id)
}

type UpdateRightScriptRightScriptRunner struct {
	id                     string
	rightScriptDescription string
	rightScriptName        string
}

func (r *UpdateRightScriptRightScriptRunner) Run(c *Client) (interface{}, error) {

	/** Handle rightScript parameter **/
	var rightScript RightScriptParam2

	// Load JSON if provided
	if len(r.rightScriptJson) > 0 {
		if err := Json.Unmarshal(r.rightScriptJson, &rightScript); err != nil {
			return nil, fmt.Errorf("Could not load rightScript JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.rightScriptdescription) > 0 {
		rightScript.rightScript.description = r.rightScriptdescription
	}

	if len(r.rightScriptname) > 0 {
		rightScript.rightScript.name = r.rightScriptname
	}

	return c.UpdateRightScript(r.id, &rightScript)
}

type UpdateSourceRightScriptRightScriptRunner struct {
	id string
}

func (r *UpdateSourceRightScriptRightScriptRunner) Run(c *Client) (interface{}, error) {
	return c.UpdateSourceRightScript(r.id)
}

// Register all RightScript actions
func registerRightScriptCmds(app *kingpin.Application) {
	resCmd := app.Cmd("RightScript", `A RightScript is an executable piece of code that can be run on a server during the boot, operational, or decommission phases...`)

	CommitRightScriptRunner := new(CommitRightScriptRightScriptRunner)
	CommitRightScriptCmd := resCmd.Command("CommitRightScript", `Commits the given RightScript. Only HEAD revisions (revision 0) can be committed.`)
	CommitRightScriptRunner.Flag(`id`, ``).Required().StringVar(&CommitRightScriptRunner.id)
	CommitRightScriptRunner.Flag(`rightScript.commitMessage`, `The message to be included with the requested commit`).Required().StringVar(&CommitRightScriptRunner.rightScriptCommitMessage)
	registry[CommitRightScriptCmd.FullCommand()] = CommitRightScriptRunner

	IndexRightScriptsRunner := new(IndexRightScriptsRightScriptRunner)
	IndexRightScriptsCmd := resCmd.Command("IndexRightScripts", `Lists RightScripts.`)
	IndexRightScriptsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexRightScriptsRunner.filterPos).StringsVar(&IndexRightScriptsRunner.filter)
	IndexRightScriptsRunner.Flag(`latestOnly`, `Whether or not to return only the latest version for each lineage.`).StringVar(&IndexRightScriptsRunner.latestOnly)
	IndexRightScriptsRunner.Flag(`view`, ``).StringVar(&IndexRightScriptsRunner.view)
	registry[IndexRightScriptsCmd.FullCommand()] = IndexRightScriptsRunner

	ShowRightScriptRunner := new(ShowRightScriptRightScriptRunner)
	ShowRightScriptCmd := resCmd.Command("ShowRightScript", `Displays information about a single RightScript.`)
	ShowRightScriptRunner.Flag(`id`, ``).Required().StringVar(&ShowRightScriptRunner.id)
	registry[ShowRightScriptCmd.FullCommand()] = ShowRightScriptRunner

	ShowSourceRightScriptRunner := new(ShowSourceRightScriptRightScriptRunner)
	ShowSourceRightScriptCmd := resCmd.Command("ShowSourceRightScript", `Returns the script source for a RightScript`)
	ShowSourceRightScriptRunner.Flag(`id`, ``).Required().StringVar(&ShowSourceRightScriptRunner.id)
	registry[ShowSourceRightScriptCmd.FullCommand()] = ShowSourceRightScriptRunner

	UpdateRightScriptRunner := new(UpdateRightScriptRightScriptRunner)
	UpdateRightScriptCmd := resCmd.Command("UpdateRightScript", `Updates RightScript name/description`)
	UpdateRightScriptRunner.Flag(`id`, ``).Required().StringVar(&UpdateRightScriptRunner.id)
	UpdateRightScriptRunner.Flag(`rightScript.description`, `The new description for the RightScript`).StringVar(&UpdateRightScriptRunner.rightScriptDescription)
	UpdateRightScriptRunner.Flag(`rightScript.name`, `The new name for the RightScript`).StringVar(&UpdateRightScriptRunner.rightScriptName)
	registry[UpdateRightScriptCmd.FullCommand()] = UpdateRightScriptRunner

	UpdateSourceRightScriptRunner := new(UpdateSourceRightScriptRightScriptRunner)
	UpdateSourceRightScriptCmd := resCmd.Command("UpdateSourceRightScript", `Updates the source of the given RightScript`)
	UpdateSourceRightScriptRunner.Flag(`id`, ``).Required().StringVar(&UpdateSourceRightScriptRunner.id)
	registry[UpdateSourceRightScriptCmd.FullCommand()] = UpdateSourceRightScriptRunner
}

/****** Route ******/

type CreateRouteRouteRunner struct {
	routeDescription          string
	routeDestinationCidrBlock string
	routeNextHopHref          string
	routeNextHopIp            string
	routeNextHopType          string
	routeRouteTableHref       string
}

func (r *CreateRouteRouteRunner) Run(c *Client) (interface{}, error) {

	/** Handle route parameter **/
	var route RouteParam

	// Load JSON if provided
	if len(r.routeJson) > 0 {
		if err := Json.Unmarshal(r.routeJson, &route); err != nil {
			return nil, fmt.Errorf("Could not load route JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.routedescription) > 0 {
		route.route.description = r.routedescription
	}

	if len(r.routedestinationCidrBlock) > 0 {
		route.route.destinationCidrBlock = r.routedestinationCidrBlock
	}

	if len(r.routenextHopHref) > 0 {
		route.route.nextHopHref = r.routenextHopHref
	}

	if len(r.routenextHopIp) > 0 {
		route.route.nextHopIp = r.routenextHopIp
	}

	if len(r.routenextHopType) > 0 {
		route.route.nextHopType = r.routenextHopType
	}

	if len(r.routerouteTableHref) > 0 {
		route.route.routeTableHref = r.routerouteTableHref
	}

	return c.CreateRoute(&route)
}

type DestroyRouteRouteRunner struct {
	id string
}

func (r *DestroyRouteRouteRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyRoute(r.id)
}

type IndexRoutesRouteRunner struct {
	filter    []string
	filterPos []string
}

func (r *IndexRoutesRouteRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexRoutes(filter)
}

type ShowRouteRouteRunner struct {
	id string
}

func (r *ShowRouteRouteRunner) Run(c *Client) (interface{}, error) {
	return c.ShowRoute(r.id)
}

type UpdateRouteRouteRunner struct {
	id                        string
	routeDescription          string
	routeDestinationCidrBlock string
	routeNextHopHref          string
	routeNextHopIp            string
	routeNextHopType          string
}

func (r *UpdateRouteRouteRunner) Run(c *Client) (interface{}, error) {

	/** Handle route parameter **/
	var route RouteParam2

	// Load JSON if provided
	if len(r.routeJson) > 0 {
		if err := Json.Unmarshal(r.routeJson, &route); err != nil {
			return nil, fmt.Errorf("Could not load route JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.routedescription) > 0 {
		route.route.description = r.routedescription
	}

	if len(r.routedestinationCidrBlock) > 0 {
		route.route.destinationCidrBlock = r.routedestinationCidrBlock
	}

	if len(r.routenextHopHref) > 0 {
		route.route.nextHopHref = r.routenextHopHref
	}

	if len(r.routenextHopIp) > 0 {
		route.route.nextHopIp = r.routenextHopIp
	}

	if len(r.routenextHopType) > 0 {
		route.route.nextHopType = r.routenextHopType
	}

	return c.UpdateRoute(r.id, &route)
}

// Register all Route actions
func registerRouteCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Route", `A Route defines how networking traffic should be routed from one destination to another...`)

	CreateRouteRunner := new(CreateRouteRouteRunner)
	CreateRouteCmd := resCmd.Command("CreateRoute", `Create a new Route.`)
	CreateRouteRunner.Flag(`route.description`, `The description to be set on the Route.`).StringVar(&CreateRouteRunner.routeDescription)
	CreateRouteRunner.Flag(`route.destinationCidrBlock`, `The destination (CIDR IP address) for the Route.`).Required().StringVar(&CreateRouteRunner.routeDestinationCidrBlock)
	CreateRouteRunner.Flag(`route.nextHopHref`, `The href of the Route's next hop.`).StringVar(&CreateRouteRunner.routeNextHopHref)
	CreateRouteRunner.Flag(`route.nextHopIp`, `The IP Address of the Route's next hop. Required if route[next_hop_type] is 'ip_string'. Not allowed otherwise.`).StringVar(&CreateRouteRunner.routeNextHopIp)
	CreateRouteRunner.Flag(`route.nextHopType`, `The Route's next hop type.`).Required().StringVar(&CreateRouteRunner.routeNextHopType)
	CreateRouteRunner.Flag(`route.routeTableHref`, `The RouteTable to create the Route in.`).Required().StringVar(&CreateRouteRunner.routeRouteTableHref)
	registry[CreateRouteCmd.FullCommand()] = CreateRouteRunner

	DestroyRouteRunner := new(DestroyRouteRouteRunner)
	DestroyRouteCmd := resCmd.Command("DestroyRoute", `Delete an existing Route.`)
	DestroyRouteRunner.Flag(`id`, ``).Required().StringVar(&DestroyRouteRunner.id)
	registry[DestroyRouteCmd.FullCommand()] = DestroyRouteRunner

	IndexRoutesRunner := new(IndexRoutesRouteRunner)
	IndexRoutesCmd := resCmd.Command("IndexRoutes", `List Routes available in this account.`)
	IndexRoutesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexRoutesRunner.filterPos).StringsVar(&IndexRoutesRunner.filter)
	registry[IndexRoutesCmd.FullCommand()] = IndexRoutesRunner

	ShowRouteRunner := new(ShowRouteRouteRunner)
	ShowRouteCmd := resCmd.Command("ShowRoute", `Show information about a single Route.`)
	ShowRouteRunner.Flag(`id`, ``).Required().StringVar(&ShowRouteRunner.id)
	registry[ShowRouteCmd.FullCommand()] = ShowRouteRunner

	UpdateRouteRunner := new(UpdateRouteRouteRunner)
	UpdateRouteCmd := resCmd.Command("UpdateRoute", `Update an existing Route.`)
	UpdateRouteRunner.Flag(`id`, ``).Required().StringVar(&UpdateRouteRunner.id)
	UpdateRouteRunner.Flag(`route.description`, `The updated description of the Route.`).StringVar(&UpdateRouteRunner.routeDescription)
	UpdateRouteRunner.Flag(`route.destinationCidrBlock`, `The updated destination (CIDR IP address) for the Route.`).StringVar(&UpdateRouteRunner.routeDestinationCidrBlock)
	UpdateRouteRunner.Flag(`route.nextHopHref`, `The updated href of the Route's next hop. Required if route[next_hop_type] is 'instance', 'network_interface', or 'network_gateway'. Not allowed otherwise.`).StringVar(&UpdateRouteRunner.routeNextHopHref)
	UpdateRouteRunner.Flag(`route.nextHopIp`, `The updated IP Address of the Route's next hop. Required if route[next_hop_type] is 'ip_string'. Not allowed otherwise.`).StringVar(&UpdateRouteRunner.routeNextHopIp)
	UpdateRouteRunner.Flag(`route.nextHopType`, `The updated Route's next hop type.`).StringVar(&UpdateRouteRunner.routeNextHopType)
	registry[UpdateRouteCmd.FullCommand()] = UpdateRouteRunner
}

/****** RouteTable ******/

type CreateRouteTableRouteTableRunner struct {
	routeTableCloudHref   string
	routeTableDescription string
	routeTableName        string
	routeTableNetworkHref string
}

func (r *CreateRouteTableRouteTableRunner) Run(c *Client) (interface{}, error) {

	/** Handle routeTable parameter **/
	var routeTable RouteTableParam

	// Load JSON if provided
	if len(r.routeTableJson) > 0 {
		if err := Json.Unmarshal(r.routeTableJson, &routeTable); err != nil {
			return nil, fmt.Errorf("Could not load routeTable JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.routeTablecloudHref) > 0 {
		routeTable.routeTable.cloudHref = r.routeTablecloudHref
	}

	if len(r.routeTabledescription) > 0 {
		routeTable.routeTable.description = r.routeTabledescription
	}

	if len(r.routeTablename) > 0 {
		routeTable.routeTable.name = r.routeTablename
	}

	if len(r.routeTablenetworkHref) > 0 {
		routeTable.routeTable.networkHref = r.routeTablenetworkHref
	}

	return c.CreateRouteTable(&routeTable)
}

type DestroyRouteTableRouteTableRunner struct {
	id string
}

func (r *DestroyRouteTableRouteTableRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyRouteTable(r.id)
}

type IndexRouteTablesRouteTableRunner struct {
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexRouteTablesRouteTableRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexRouteTables(filter, r.view)
}

type ShowRouteTableRouteTableRunner struct {
	id   string
	view string
}

func (r *ShowRouteTableRouteTableRunner) Run(c *Client) (interface{}, error) {
	return c.ShowRouteTable(r.id, r.view)
}

type UpdateRouteTableRouteTableRunner struct {
	id                    string
	routeTableDescription string
	routeTableName        string
}

func (r *UpdateRouteTableRouteTableRunner) Run(c *Client) (interface{}, error) {

	/** Handle routeTable parameter **/
	var routeTable RouteTableParam2

	// Load JSON if provided
	if len(r.routeTableJson) > 0 {
		if err := Json.Unmarshal(r.routeTableJson, &routeTable); err != nil {
			return nil, fmt.Errorf("Could not load routeTable JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.routeTabledescription) > 0 {
		routeTable.routeTable.description = r.routeTabledescription
	}

	if len(r.routeTablename) > 0 {
		routeTable.routeTable.name = r.routeTablename
	}

	return c.UpdateRouteTable(r.id, &routeTable)
}

// Register all RouteTable actions
func registerRouteTableCmds(app *kingpin.Application) {
	resCmd := app.Cmd("RouteTable", `Grouped listing of Routes`)

	CreateRouteTableRunner := new(CreateRouteTableRouteTableRunner)
	CreateRouteTableCmd := resCmd.Command("CreateRouteTable", `Create a new RouteTable.`)
	CreateRouteTableRunner.Flag(`routeTable.cloudHref`, `The cloud to create the RouteTable in.`).Required().StringVar(&CreateRouteTableRunner.routeTableCloudHref)
	CreateRouteTableRunner.Flag(`routeTable.description`, `The description to be set on the RouteTable.`).StringVar(&CreateRouteTableRunner.routeTableDescription)
	CreateRouteTableRunner.Flag(`routeTable.name`, `The name to be set on the RouteTable.`).Required().StringVar(&CreateRouteTableRunner.routeTableName)
	CreateRouteTableRunner.Flag(`routeTable.networkHref`, `The Network to create the RouteTable in.`).Required().StringVar(&CreateRouteTableRunner.routeTableNetworkHref)
	registry[CreateRouteTableCmd.FullCommand()] = CreateRouteTableRunner

	DestroyRouteTableRunner := new(DestroyRouteTableRouteTableRunner)
	DestroyRouteTableCmd := resCmd.Command("DestroyRouteTable", `Delete an existing RouteTable.`)
	DestroyRouteTableRunner.Flag(`id`, ``).Required().StringVar(&DestroyRouteTableRunner.id)
	registry[DestroyRouteTableCmd.FullCommand()] = DestroyRouteTableRunner

	IndexRouteTablesRunner := new(IndexRouteTablesRouteTableRunner)
	IndexRouteTablesCmd := resCmd.Command("IndexRouteTables", `List RouteTables available in this account.`)
	IndexRouteTablesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexRouteTablesRunner.filterPos).StringsVar(&IndexRouteTablesRunner.filter)
	IndexRouteTablesRunner.Flag(`view`, ``).StringVar(&IndexRouteTablesRunner.view)
	registry[IndexRouteTablesCmd.FullCommand()] = IndexRouteTablesRunner

	ShowRouteTableRunner := new(ShowRouteTableRouteTableRunner)
	ShowRouteTableCmd := resCmd.Command("ShowRouteTable", `Show information about a single RouteTable.`)
	ShowRouteTableRunner.Flag(`id`, ``).Required().StringVar(&ShowRouteTableRunner.id)
	ShowRouteTableRunner.Flag(`view`, ``).StringVar(&ShowRouteTableRunner.view)
	registry[ShowRouteTableCmd.FullCommand()] = ShowRouteTableRunner

	UpdateRouteTableRunner := new(UpdateRouteTableRouteTableRunner)
	UpdateRouteTableCmd := resCmd.Command("UpdateRouteTable", `Update an existing RouteTable.`)
	UpdateRouteTableRunner.Flag(`id`, ``).Required().StringVar(&UpdateRouteTableRunner.id)
	UpdateRouteTableRunner.Flag(`routeTable.description`, `The description to be set on the RouteTable.`).StringVar(&UpdateRouteTableRunner.routeTableDescription)
	UpdateRouteTableRunner.Flag(`routeTable.name`, `The name to be set on the RouteTable.`).StringVar(&UpdateRouteTableRunner.routeTableName)
	registry[UpdateRouteTableCmd.FullCommand()] = UpdateRouteTableRunner
}

/****** RunnableBinding ******/

type CreateRunnableBindingRunnableBindingRunner struct {
	runnableBindingPosition        string
	runnableBindingRecipe          string
	runnableBindingRightScriptHref string
	runnableBindingSequence        string
	serverTemplateId               string
}

func (r *CreateRunnableBindingRunnableBindingRunner) Run(c *Client) (interface{}, error) {

	/** Handle runnableBinding parameter **/
	var runnableBinding RunnableBindingParam

	// Load JSON if provided
	if len(r.runnableBindingJson) > 0 {
		if err := Json.Unmarshal(r.runnableBindingJson, &runnableBinding); err != nil {
			return nil, fmt.Errorf("Could not load runnableBinding JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.runnableBindingposition) > 0 {
		runnableBinding.runnableBinding.position = r.runnableBindingposition
	}

	if len(r.runnableBindingrecipe) > 0 {
		runnableBinding.runnableBinding.recipe = r.runnableBindingrecipe
	}

	if len(r.runnableBindingrightScriptHref) > 0 {
		runnableBinding.runnableBinding.rightScriptHref = r.runnableBindingrightScriptHref
	}

	if len(r.runnableBindingsequence) > 0 {
		runnableBinding.runnableBinding.sequence = r.runnableBindingsequence
	}

	return c.CreateRunnableBinding(&runnableBinding, r.serverTemplateId)
}

type DestroyRunnableBindingRunnableBindingRunner struct {
	id               string
	serverTemplateId string
}

func (r *DestroyRunnableBindingRunnableBindingRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyRunnableBinding(r.id, r.serverTemplateId)
}

type IndexRunnableBindingsRunnableBindingRunner struct {
	serverTemplateId string
	view             string
}

func (r *IndexRunnableBindingsRunnableBindingRunner) Run(c *Client) (interface{}, error) {
	return c.IndexRunnableBindings(r.serverTemplateId, r.view)
}

type MultiUpdateRunnableBindingsRunnableBindingRunner struct {
	runnableBindingsId                 []string
	runnableBindingsIdPos              []string
	runnableBindingsPosition           []string
	runnableBindingsPositionPos        []string
	runnableBindingsRecipe             []string
	runnableBindingsRecipePos          []string
	runnableBindingsRightScriptHref    []string
	runnableBindingsRightScriptHrefPos []string
	runnableBindingsSequence           []string
	runnableBindingsSequencePos        []string
	serverTemplateId                   string
}

func (r *MultiUpdateRunnableBindingsRunnableBindingRunner) Run(c *Client) (interface{}, error) {

	/** Handle runnableBindings parameter **/
	var runnableBindings []*RunnableBindings

	for i, v := range r.runnableBindings {
		pos, err := strconv.Atoi(r.runnableBindingsPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for runnableBindings array", r.runnableBindingsPos[i])
		}
		runnableBindings[pos] = v
	}

	return c.MultiUpdateRunnableBindings(runnableBindings, r.serverTemplateId)
}

type ShowRunnableBindingRunnableBindingRunner struct {
	id               string
	serverTemplateId string
	view             string
}

func (r *ShowRunnableBindingRunnableBindingRunner) Run(c *Client) (interface{}, error) {
	return c.ShowRunnableBinding(r.id, r.serverTemplateId, r.view)
}

// Register all RunnableBinding actions
func registerRunnableBindingCmds(app *kingpin.Application) {
	resCmd := app.Cmd("RunnableBinding", `A RunnableBinding represents an item in a runlist of a ServerTemplate`)

	CreateRunnableBindingRunner := new(CreateRunnableBindingRunnableBindingRunner)
	CreateRunnableBindingCmd := resCmd.Command("CreateRunnableBinding", `Bind an executable to the given ServerTemplate.`)
	CreateRunnableBindingRunner.Flag(`runnableBinding.position`, `The position of the executable in the execution order. If not specified, will be added to the end. If specified, will be inserted in that location and cause all others to move down.`).StringVar(&CreateRunnableBindingRunner.runnableBindingPosition)
	CreateRunnableBindingRunner.Flag(`runnableBinding.recipe`, `The Chef recipe name. Note: right_script_href cannot be specified when this param is given.`).StringVar(&CreateRunnableBindingRunner.runnableBindingRecipe)
	CreateRunnableBindingRunner.Flag(`runnableBinding.rightScriptHref`, `The RightScript href. Note: recipe cannot be specified when this param is given.`).StringVar(&CreateRunnableBindingRunner.runnableBindingRightScriptHref)
	CreateRunnableBindingRunner.Flag(`runnableBinding.sequence`, `The sequence at which this executable should be run. Default is 'operational'.`).StringVar(&CreateRunnableBindingRunner.runnableBindingSequence)
	CreateRunnableBindingRunner.Flag(`serverTemplateId`, ``).Required().StringVar(&CreateRunnableBindingRunner.serverTemplateId)
	registry[CreateRunnableBindingCmd.FullCommand()] = CreateRunnableBindingRunner

	DestroyRunnableBindingRunner := new(DestroyRunnableBindingRunnableBindingRunner)
	DestroyRunnableBindingCmd := resCmd.Command("DestroyRunnableBinding", `Unbind an executable from the given resource.`)
	DestroyRunnableBindingRunner.Flag(`id`, ``).Required().StringVar(&DestroyRunnableBindingRunner.id)
	DestroyRunnableBindingRunner.Flag(`serverTemplateId`, ``).Required().StringVar(&DestroyRunnableBindingRunner.serverTemplateId)
	registry[DestroyRunnableBindingCmd.FullCommand()] = DestroyRunnableBindingRunner

	IndexRunnableBindingsRunner := new(IndexRunnableBindingsRunnableBindingRunner)
	IndexRunnableBindingsCmd := resCmd.Command("IndexRunnableBindings", `Lists the executables bound to the ServerTemplate.`)
	IndexRunnableBindingsRunner.Flag(`serverTemplateId`, ``).Required().StringVar(&IndexRunnableBindingsRunner.serverTemplateId)
	IndexRunnableBindingsRunner.Flag(`view`, ``).StringVar(&IndexRunnableBindingsRunner.view)
	registry[IndexRunnableBindingsCmd.FullCommand()] = IndexRunnableBindingsRunner

	MultiUpdateRunnableBindingsRunner := new(MultiUpdateRunnableBindingsRunnableBindingRunner)
	MultiUpdateRunnableBindingsCmd := resCmd.Command("MultiUpdateRunnableBindings", `Update attributes for multiple bound executables.`)
	MultiUpdateRunnableBindingsRunner.FlagPattern(`runnableBindings\.(\d+)\.id`, `The ID of the RunnableBinding to update.`).Required().Capture(&MultiUpdateRunnableBindingsRunner.runnableBindingsIdPos).StringsVar(&MultiUpdateRunnableBindingsRunner.runnableBindingsId)
	MultiUpdateRunnableBindingsRunner.FlagPattern(`runnableBindings\.(\d+)\.position`, `The updated position of the RunnableBinding in the execution order. If specified, will be inserted in that location and cause all others to move down.`).Capture(&MultiUpdateRunnableBindingsRunner.runnableBindingsPositionPos).StringsVar(&MultiUpdateRunnableBindingsRunner.runnableBindingsPosition)
	MultiUpdateRunnableBindingsRunner.FlagPattern(`runnableBindings\.(\d+)\.recipe`, `The updated Chef recipe name. Note: right_script_href cannot be specified when this param is given.`).Capture(&MultiUpdateRunnableBindingsRunner.runnableBindingsRecipePos).StringsVar(&MultiUpdateRunnableBindingsRunner.runnableBindingsRecipe)
	MultiUpdateRunnableBindingsRunner.FlagPattern(`runnableBindings\.(\d+)\.rightScriptHref`, `The updated RightScript href. Note: recipe cannot be specified when this param is given.`).Capture(&MultiUpdateRunnableBindingsRunner.runnableBindingsRightScriptHrefPos).StringsVar(&MultiUpdateRunnableBindingsRunner.runnableBindingsRightScriptHref)
	MultiUpdateRunnableBindingsRunner.FlagPattern(`runnableBindings\.(\d+)\.sequence`, `The sequence at which this executable should be run.  Default is 'operational'.`).Capture(&MultiUpdateRunnableBindingsRunner.runnableBindingsSequencePos).StringsVar(&MultiUpdateRunnableBindingsRunner.runnableBindingsSequence)
	MultiUpdateRunnableBindingsRunner.Flag(`serverTemplateId`, ``).Required().StringVar(&MultiUpdateRunnableBindingsRunner.serverTemplateId)
	registry[MultiUpdateRunnableBindingsCmd.FullCommand()] = MultiUpdateRunnableBindingsRunner

	ShowRunnableBindingRunner := new(ShowRunnableBindingRunnableBindingRunner)
	ShowRunnableBindingCmd := resCmd.Command("ShowRunnableBinding", `Show information about a single executable binding.`)
	ShowRunnableBindingRunner.Flag(`id`, ``).Required().StringVar(&ShowRunnableBindingRunner.id)
	ShowRunnableBindingRunner.Flag(`serverTemplateId`, ``).Required().StringVar(&ShowRunnableBindingRunner.serverTemplateId)
	ShowRunnableBindingRunner.Flag(`view`, ``).StringVar(&ShowRunnableBindingRunner.view)
	registry[ShowRunnableBindingCmd.FullCommand()] = ShowRunnableBindingRunner
}

/****** SecurityGroup ******/

type CreateSecurityGroupSecurityGroupRunner struct {
	cloudId                  string
	securityGroupDescription string
	securityGroupName        string
	securityGroupNetworkHref string
}

func (r *CreateSecurityGroupSecurityGroupRunner) Run(c *Client) (interface{}, error) {

	/** Handle securityGroup parameter **/
	var securityGroup SecurityGroupParam

	// Load JSON if provided
	if len(r.securityGroupJson) > 0 {
		if err := Json.Unmarshal(r.securityGroupJson, &securityGroup); err != nil {
			return nil, fmt.Errorf("Could not load securityGroup JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.securityGroupdescription) > 0 {
		securityGroup.securityGroup.description = r.securityGroupdescription
	}

	if len(r.securityGroupname) > 0 {
		securityGroup.securityGroup.name = r.securityGroupname
	}

	if len(r.securityGroupnetworkHref) > 0 {
		securityGroup.securityGroup.networkHref = r.securityGroupnetworkHref
	}

	return c.CreateSecurityGroup(r.cloudId, &securityGroup)
}

type DestroySecurityGroupSecurityGroupRunner struct {
	cloudId string
	id      string
}

func (r *DestroySecurityGroupSecurityGroupRunner) Run(c *Client) (interface{}, error) {
	return c.DestroySecurityGroup(r.cloudId, r.id)
}

type IndexSecurityGroupsSecurityGroupRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexSecurityGroupsSecurityGroupRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexSecurityGroups(r.cloudId, filter, r.view)
}

type ShowSecurityGroupSecurityGroupRunner struct {
	cloudId string
	id      string
	view    string
}

func (r *ShowSecurityGroupSecurityGroupRunner) Run(c *Client) (interface{}, error) {
	return c.ShowSecurityGroup(r.cloudId, r.id, r.view)
}

// Register all SecurityGroup actions
func registerSecurityGroupCmds(app *kingpin.Application) {
	resCmd := app.Cmd("SecurityGroup", `Security Groups represent network security profiles that contain lists of firewall rules for different ports and source IP addresses, as well as trust relationships amongst different security groups...`)

	CreateSecurityGroupRunner := new(CreateSecurityGroupSecurityGroupRunner)
	CreateSecurityGroupCmd := resCmd.Command("CreateSecurityGroup", `Create a security group.`)
	CreateSecurityGroupRunner.Flag(`cloudId`, ``).Required().StringVar(&CreateSecurityGroupRunner.cloudId)
	CreateSecurityGroupRunner.Flag(`securityGroup.description`, ``).StringVar(&CreateSecurityGroupRunner.securityGroupDescription)
	CreateSecurityGroupRunner.Flag(`securityGroup.name`, ``).Required().StringVar(&CreateSecurityGroupRunner.securityGroupName)
	CreateSecurityGroupRunner.Flag(`securityGroup.networkHref`, ``).StringVar(&CreateSecurityGroupRunner.securityGroupNetworkHref)
	registry[CreateSecurityGroupCmd.FullCommand()] = CreateSecurityGroupRunner

	DestroySecurityGroupRunner := new(DestroySecurityGroupSecurityGroupRunner)
	DestroySecurityGroupCmd := resCmd.Command("DestroySecurityGroup", `Delete security group(s)`)
	DestroySecurityGroupRunner.Flag(`cloudId`, ``).Required().StringVar(&DestroySecurityGroupRunner.cloudId)
	DestroySecurityGroupRunner.Flag(`id`, ``).Required().StringVar(&DestroySecurityGroupRunner.id)
	registry[DestroySecurityGroupCmd.FullCommand()] = DestroySecurityGroupRunner

	IndexSecurityGroupsRunner := new(IndexSecurityGroupsSecurityGroupRunner)
	IndexSecurityGroupsCmd := resCmd.Command("IndexSecurityGroups", `Lists Security Groups.`)
	IndexSecurityGroupsRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexSecurityGroupsRunner.cloudId)
	IndexSecurityGroupsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexSecurityGroupsRunner.filterPos).StringsVar(&IndexSecurityGroupsRunner.filter)
	IndexSecurityGroupsRunner.Flag(`view`, ``).StringVar(&IndexSecurityGroupsRunner.view)
	registry[IndexSecurityGroupsCmd.FullCommand()] = IndexSecurityGroupsRunner

	ShowSecurityGroupRunner := new(ShowSecurityGroupSecurityGroupRunner)
	ShowSecurityGroupCmd := resCmd.Command("ShowSecurityGroup", `Displays information about a single Security Group.`)
	ShowSecurityGroupRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowSecurityGroupRunner.cloudId)
	ShowSecurityGroupRunner.Flag(`id`, ``).Required().StringVar(&ShowSecurityGroupRunner.id)
	ShowSecurityGroupRunner.Flag(`view`, ``).StringVar(&ShowSecurityGroupRunner.view)
	registry[ShowSecurityGroupCmd.FullCommand()] = ShowSecurityGroupRunner
}

/****** SecurityGroupRule ******/

type CreateSecurityGroupRuleSecurityGroupRuleRunner struct {
	securityGroupRuleCidrIps                  string
	securityGroupRuleDirection                string
	securityGroupRuleGroupName                string
	securityGroupRuleGroupOwner               string
	securityGroupRuleProtocol                 string
	securityGroupRuleProtocolDetailsEndPort   string
	securityGroupRuleProtocolDetailsIcmpCode  string
	securityGroupRuleProtocolDetailsIcmpType  string
	securityGroupRuleProtocolDetailsStartPort string
	securityGroupRuleSecurityGroupHref        string
	securityGroupRuleSourceType               string
}

func (r *CreateSecurityGroupRuleSecurityGroupRuleRunner) Run(c *Client) (interface{}, error) {

	/** Handle securityGroupRule parameter **/
	var securityGroupRule SecurityGroupRuleParam

	// Load JSON if provided
	if len(r.securityGroupRuleJson) > 0 {
		if err := Json.Unmarshal(r.securityGroupRuleJson, &securityGroupRule); err != nil {
			return nil, fmt.Errorf("Could not load securityGroupRule JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.securityGroupRulecidrIps) > 0 {
		securityGroupRule.securityGroupRule.cidrIps = r.securityGroupRulecidrIps
	}

	if len(r.securityGroupRuledirection) > 0 {
		securityGroupRule.securityGroupRule.direction = r.securityGroupRuledirection
	}

	if len(r.securityGroupRulegroupName) > 0 {
		securityGroupRule.securityGroupRule.groupName = r.securityGroupRulegroupName
	}

	if len(r.securityGroupRulegroupOwner) > 0 {
		securityGroupRule.securityGroupRule.groupOwner = r.securityGroupRulegroupOwner
	}

	if len(r.securityGroupRuleprotocol) > 0 {
		securityGroupRule.securityGroupRule.protocol = r.securityGroupRuleprotocol
	}

	if len(r.securityGroupRuleprotocolDetailsendPort) > 0 {
		securityGroupRule.securityGroupRule.protocolDetails.endPort = r.securityGroupRuleprotocolDetailsendPort
	}

	if len(r.securityGroupRuleprotocolDetailsicmpCode) > 0 {
		securityGroupRule.securityGroupRule.protocolDetails.icmpCode = r.securityGroupRuleprotocolDetailsicmpCode
	}

	if len(r.securityGroupRuleprotocolDetailsicmpType) > 0 {
		securityGroupRule.securityGroupRule.protocolDetails.icmpType = r.securityGroupRuleprotocolDetailsicmpType
	}

	if len(r.securityGroupRuleprotocolDetailsstartPort) > 0 {
		securityGroupRule.securityGroupRule.protocolDetails.startPort = r.securityGroupRuleprotocolDetailsstartPort
	}

	if len(r.securityGroupRulesecurityGroupHref) > 0 {
		securityGroupRule.securityGroupRule.securityGroupHref = r.securityGroupRulesecurityGroupHref
	}

	if len(r.securityGroupRulesourceType) > 0 {
		securityGroupRule.securityGroupRule.sourceType = r.securityGroupRulesourceType
	}

	return c.CreateSecurityGroupRule(&securityGroupRule)
}

type DestroySecurityGroupRuleSecurityGroupRuleRunner struct {
	id string
}

func (r *DestroySecurityGroupRuleSecurityGroupRuleRunner) Run(c *Client) (interface{}, error) {
	return c.DestroySecurityGroupRule(r.id)
}

type IndexSecurityGroupRulesSecurityGroupRuleRunner struct {
	view string
}

func (r *IndexSecurityGroupRulesSecurityGroupRuleRunner) Run(c *Client) (interface{}, error) {
	return c.IndexSecurityGroupRules(r.view)
}

type ShowSecurityGroupRuleSecurityGroupRuleRunner struct {
	id   string
	view string
}

func (r *ShowSecurityGroupRuleSecurityGroupRuleRunner) Run(c *Client) (interface{}, error) {
	return c.ShowSecurityGroupRule(r.id, r.view)
}

type UpdateSecurityGroupRuleSecurityGroupRuleRunner struct {
	id                           string
	securityGroupRuleDescription string
}

func (r *UpdateSecurityGroupRuleSecurityGroupRuleRunner) Run(c *Client) (interface{}, error) {

	/** Handle securityGroupRule parameter **/
	var securityGroupRule SecurityGroupRuleParam2

	// Load JSON if provided
	if len(r.securityGroupRuleJson) > 0 {
		if err := Json.Unmarshal(r.securityGroupRuleJson, &securityGroupRule); err != nil {
			return nil, fmt.Errorf("Could not load securityGroupRule JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.securityGroupRuledescription) > 0 {
		securityGroupRule.securityGroupRule.description = r.securityGroupRuledescription
	}

	return c.UpdateSecurityGroupRule(r.id, &securityGroupRule)
}

// Register all SecurityGroupRule actions
func registerSecurityGroupRuleCmds(app *kingpin.Application) {
	resCmd := app.Cmd("SecurityGroupRule", ``)

	CreateSecurityGroupRuleRunner := new(CreateSecurityGroupRuleSecurityGroupRuleRunner)
	CreateSecurityGroupRuleCmd := resCmd.Command("CreateSecurityGroupRule", `Create a security group rule for a security group.`)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.cidrIps`, `An IP address range in CIDR notation. Required if source_type is 'cidr_ips'.`).StringVar(&CreateSecurityGroupRuleRunner.securityGroupRuleCidrIps)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.direction`, `Direction of traffic.`).StringVar(&CreateSecurityGroupRuleRunner.securityGroupRuleDirection)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.groupName`, `Name of source Security Group. Required if source_type is 'group'.`).StringVar(&CreateSecurityGroupRuleRunner.securityGroupRuleGroupName)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.groupOwner`, `Owner of source Security Group. Required if source_type is 'group'.`).StringVar(&CreateSecurityGroupRuleRunner.securityGroupRuleGroupOwner)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.protocol`, `Protocol to filter on.`).Required().StringVar(&CreateSecurityGroupRuleRunner.securityGroupRuleProtocol)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.protocolDetails.endPort`, `End of port range (inclusive). Required if protocol is 'tcp' or 'udp'.`).StringVar(&CreateSecurityGroupRuleRunner.securityGroupRuleProtocolDetailsEndPort)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.protocolDetails.icmpCode`, `ICMP code. Required if protocol is 'icmp'.`).StringVar(&CreateSecurityGroupRuleRunner.securityGroupRuleProtocolDetailsIcmpCode)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.protocolDetails.icmpType`, `ICMP type. Required if protocol is 'icmp'.`).StringVar(&CreateSecurityGroupRuleRunner.securityGroupRuleProtocolDetailsIcmpType)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.protocolDetails.startPort`, `Start of port range (inclusive). Required if protocol is 'tcp' or 'udp'.`).StringVar(&CreateSecurityGroupRuleRunner.securityGroupRuleProtocolDetailsStartPort)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.securityGroupHref`, `Security Group to add rule to.`).StringVar(&CreateSecurityGroupRuleRunner.securityGroupRuleSecurityGroupHref)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.sourceType`, `Source type. May be a CIDR block or another Security Group.`).Required().StringVar(&CreateSecurityGroupRuleRunner.securityGroupRuleSourceType)
	registry[CreateSecurityGroupRuleCmd.FullCommand()] = CreateSecurityGroupRuleRunner

	DestroySecurityGroupRuleRunner := new(DestroySecurityGroupRuleSecurityGroupRuleRunner)
	DestroySecurityGroupRuleCmd := resCmd.Command("DestroySecurityGroupRule", `Delete security group rule(s)`)
	DestroySecurityGroupRuleRunner.Flag(`id`, ``).Required().StringVar(&DestroySecurityGroupRuleRunner.id)
	registry[DestroySecurityGroupRuleCmd.FullCommand()] = DestroySecurityGroupRuleRunner

	IndexSecurityGroupRulesRunner := new(IndexSecurityGroupRulesSecurityGroupRuleRunner)
	IndexSecurityGroupRulesCmd := resCmd.Command("IndexSecurityGroupRules", `Lists SecurityGroupRules.`)
	IndexSecurityGroupRulesRunner.Flag(`view`, ``).StringVar(&IndexSecurityGroupRulesRunner.view)
	registry[IndexSecurityGroupRulesCmd.FullCommand()] = IndexSecurityGroupRulesRunner

	ShowSecurityGroupRuleRunner := new(ShowSecurityGroupRuleSecurityGroupRuleRunner)
	ShowSecurityGroupRuleCmd := resCmd.Command("ShowSecurityGroupRule", `Displays information about a single SecurityGroupRule.`)
	ShowSecurityGroupRuleRunner.Flag(`id`, ``).Required().StringVar(&ShowSecurityGroupRuleRunner.id)
	ShowSecurityGroupRuleRunner.Flag(`view`, ``).StringVar(&ShowSecurityGroupRuleRunner.view)
	registry[ShowSecurityGroupRuleCmd.FullCommand()] = ShowSecurityGroupRuleRunner

	UpdateSecurityGroupRuleRunner := new(UpdateSecurityGroupRuleSecurityGroupRuleRunner)
	UpdateSecurityGroupRuleCmd := resCmd.Command("UpdateSecurityGroupRule", ``)
	UpdateSecurityGroupRuleRunner.Flag(`id`, ``).Required().StringVar(&UpdateSecurityGroupRuleRunner.id)
	UpdateSecurityGroupRuleRunner.Flag(`securityGroupRule.description`, ``).StringVar(&UpdateSecurityGroupRuleRunner.securityGroupRuleDescription)
	registry[UpdateSecurityGroupRuleCmd.FullCommand()] = UpdateSecurityGroupRuleRunner
}

/****** Server ******/

type CloneServerServerRunner struct {
	id string
}

func (r *CloneServerServerRunner) Run(c *Client) (interface{}, error) {
	return c.CloneServer(r.id)
}

type CreateServerServerRunner struct {
	serverDeploymentHref                                               string
	serverDescription                                                  string
	serverInstanceAssociatePublicIpAddress                             string
	serverInstanceCloudHref                                            string
	serverInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping string
	serverInstanceCloudSpecificAttributesIamInstanceProfile            string
	serverInstanceCloudSpecificAttributesRootVolumePerformance         string
	serverInstanceCloudSpecificAttributesRootVolumeSize                string
	serverInstanceCloudSpecificAttributesRootVolumeTypeUid             string
	serverInstanceDatacenterHref                                       string
	serverInstanceImageHref                                            string
	serverInstanceInputsValues                                         []string
	serverInstanceInputsNames                                          []string
	serverInstanceInstanceTypeHref                                     string
	serverInstanceIpForwardingEnabled                                  string
	serverInstanceKernelImageHref                                      string
	serverInstanceMultiCloudImageHref                                  string
	serverInstancePlacementGroupHref                                   string
	serverInstanceRamdiskImageHref                                     string
	serverInstanceItem                                                 []string
	serverInstanceItemPos                                              []string
	serverInstanceServerTemplateHref                                   string
	serverInstanceSshKeyHref                                           string
	serverInstanceItem                                                 []string
	serverInstanceItemPos                                              []string
	serverInstanceUserData                                             string
	serverName                                                         string
	serverOptimized                                                    string
}

func (r *CreateServerServerRunner) Run(c *Client) (interface{}, error) {

	/** Handle server parameter **/
	var server ServerParam

	// Load JSON if provided
	if len(r.serverJson) > 0 {
		if err := Json.Unmarshal(r.serverJson, &server); err != nil {
			return nil, fmt.Errorf("Could not load server JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.serverdeploymentHref) > 0 {
		server.server.deploymentHref = r.serverdeploymentHref
	}

	if len(r.serverdescription) > 0 {
		server.server.description = r.serverdescription
	}

	if len(r.serverinstanceassociatePublicIpAddress) > 0 {
		server.server.instance.associatePublicIpAddress = r.serverinstanceassociatePublicIpAddress
	}

	if len(r.serverinstancecloudHref) > 0 {
		server.server.instance.cloudHref = r.serverinstancecloudHref
	}

	if len(r.serverinstancecloudSpecificAttributesautomaticInstanceStoreMapping) > 0 {
		server.server.instance.cloudSpecificAttributes.automaticInstanceStoreMapping = r.serverinstancecloudSpecificAttributesautomaticInstanceStoreMapping
	}

	if len(r.serverinstancecloudSpecificAttributesiamInstanceProfile) > 0 {
		server.server.instance.cloudSpecificAttributes.iamInstanceProfile = r.serverinstancecloudSpecificAttributesiamInstanceProfile
	}

	if len(r.serverinstancecloudSpecificAttributesrootVolumePerformance) > 0 {
		server.server.instance.cloudSpecificAttributes.rootVolumePerformance = r.serverinstancecloudSpecificAttributesrootVolumePerformance
	}

	if len(r.serverinstancecloudSpecificAttributesrootVolumeSize) > 0 {
		server.server.instance.cloudSpecificAttributes.rootVolumeSize = r.serverinstancecloudSpecificAttributesrootVolumeSize
	}

	if len(r.serverinstancecloudSpecificAttributesrootVolumeTypeUid) > 0 {
		server.server.instance.cloudSpecificAttributes.rootVolumeTypeUid = r.serverinstancecloudSpecificAttributesrootVolumeTypeUid
	}

	if len(r.serverinstancedatacenterHref) > 0 {
		server.server.instance.datacenterHref = r.serverinstancedatacenterHref
	}

	if len(r.serverinstanceimageHref) > 0 {
		server.server.instance.imageHref = r.serverinstanceimageHref
	}

	if len(r.serverinstanceinstanceTypeHref) > 0 {
		server.server.instance.instanceTypeHref = r.serverinstanceinstanceTypeHref
	}

	if len(r.serverinstanceipForwardingEnabled) > 0 {
		server.server.instance.ipForwardingEnabled = r.serverinstanceipForwardingEnabled
	}

	if len(r.serverinstancekernelImageHref) > 0 {
		server.server.instance.kernelImageHref = r.serverinstancekernelImageHref
	}

	if len(r.serverinstancemultiCloudImageHref) > 0 {
		server.server.instance.multiCloudImageHref = r.serverinstancemultiCloudImageHref
	}

	if len(r.serverinstanceplacementGroupHref) > 0 {
		server.server.instance.placementGroupHref = r.serverinstanceplacementGroupHref
	}

	if len(r.serverinstanceramdiskImageHref) > 0 {
		server.server.instance.ramdiskImageHref = r.serverinstanceramdiskImageHref
	}

	if len(r.serverinstancesecurityGroupHrefsitem) > 0 {
		server.server.instance.securityGroupHrefs.item = r.serverinstancesecurityGroupHrefsitem
	}

	if len(r.serverinstanceserverTemplateHref) > 0 {
		server.server.instance.serverTemplateHref = r.serverinstanceserverTemplateHref
	}

	if len(r.serverinstancesshKeyHref) > 0 {
		server.server.instance.sshKeyHref = r.serverinstancesshKeyHref
	}

	if len(r.serverinstancesubnetHrefsitem) > 0 {
		server.server.instance.subnetHrefs.item = r.serverinstancesubnetHrefsitem
	}

	if len(r.serverinstanceuserData) > 0 {
		server.server.instance.userData = r.serverinstanceuserData
	}

	if len(r.servername) > 0 {
		server.server.name = r.servername
	}

	if len(r.serveroptimized) > 0 {
		server.server.optimized = r.serveroptimized
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxserverinstancesecurityGroupHrefsitemPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.serverinstancesecurityGroupHrefsitemPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for server.instance.securityGroupHrefs.item field of securityGroupHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of server.instance.securityGroupHrefs.item field of securityGroupHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxserverinstancesecurityGroupHrefsitemPos {
			maxserverinstancesecurityGroupHrefsitemPos = pos
		}
	}
	if len(r.serverinstancesecurityGroupHrefsitem) != maxserverinstancesecurityGroupHrefsitemPos {
		return nil, fmt.Errof("Invalid flags for server.instance.securityGroupHrefs.item field of securityGroupHrefs array, %s were defined but max position is %s",
			len(r.serverinstancesecurityGroupHrefsitem), maxserverinstancesecurityGroupHrefsitemPos)
	}
	if maxserverinstancesecurityGroupHrefsitemPos > -1 {
		serverinstancesecurityGroupHrefs := make([]*string, maxserverinstancesecurityGroupHrefsitemPos+1)
		for i := 0; i < maxserverinstancesecurityGroupHrefsitemPos+1; i++ {
			serverinstancesecurityGroupHrefs[i] = string{
			//TBD
			}
		}
		server.instance.securityGroupHrefs = serverinstancesecurityGroupHrefs
	}

	maxserverinstancesubnetHrefsitemPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.serverinstancesubnetHrefsitemPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for server.instance.subnetHrefs.item field of subnetHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of server.instance.subnetHrefs.item field of subnetHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxserverinstancesubnetHrefsitemPos {
			maxserverinstancesubnetHrefsitemPos = pos
		}
	}
	if len(r.serverinstancesubnetHrefsitem) != maxserverinstancesubnetHrefsitemPos {
		return nil, fmt.Errof("Invalid flags for server.instance.subnetHrefs.item field of subnetHrefs array, %s were defined but max position is %s",
			len(r.serverinstancesubnetHrefsitem), maxserverinstancesubnetHrefsitemPos)
	}
	if maxserverinstancesubnetHrefsitemPos > -1 {
		serverinstancesubnetHrefs := make([]*string, maxserverinstancesubnetHrefsitemPos+1)
		for i := 0; i < maxserverinstancesubnetHrefsitemPos+1; i++ {
			serverinstancesubnetHrefs[i] = string{
			//TBD
			}
		}
		server.instance.subnetHrefs = serverinstancesubnetHrefs
	}

	// Create enumarable fields from flags
	serverinstanceinputs := make(map[string]string, len(r.serverinstanceinputsNames))
	for i, n := range r.serverinstanceinputsNames {
		serverinstanceinputs[n] = r.serverinstanceinputsValues[i]
	}
	server.server.instance.inputs = serverinstanceinputs

	return c.CreateServer(&server)
}

type DestroyServerServerRunner struct {
	id string
}

func (r *DestroyServerServerRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyServer(r.id)
}

type IndexServersServerRunner struct {
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexServersServerRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexServers(filter, r.view)
}

type LaunchServerServerRunner struct {
	id string
}

func (r *LaunchServerServerRunner) Run(c *Client) (interface{}, error) {
	return c.LaunchServer(r.id)
}

type ShowServerServerRunner struct {
	id   string
	view string
}

func (r *ShowServerServerRunner) Run(c *Client) (interface{}, error) {
	return c.ShowServer(r.id, r.view)
}

type TerminateServerServerRunner struct {
	id string
}

func (r *TerminateServerServerRunner) Run(c *Client) (interface{}, error) {
	return c.TerminateServer(r.id)
}

type UpdateServerServerRunner struct {
	id                                  string
	serverAutomaticInstanceStoreMapping string
	serverDescription                   string
	serverName                          string
	serverOptimized                     string
	serverRootVolumeSize                string
}

func (r *UpdateServerServerRunner) Run(c *Client) (interface{}, error) {

	/** Handle server parameter **/
	var server ServerParam2

	// Load JSON if provided
	if len(r.serverJson) > 0 {
		if err := Json.Unmarshal(r.serverJson, &server); err != nil {
			return nil, fmt.Errorf("Could not load server JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.serverautomaticInstanceStoreMapping) > 0 {
		server.server.automaticInstanceStoreMapping = r.serverautomaticInstanceStoreMapping
	}

	if len(r.serverdescription) > 0 {
		server.server.description = r.serverdescription
	}

	if len(r.servername) > 0 {
		server.server.name = r.servername
	}

	if len(r.serveroptimized) > 0 {
		server.server.optimized = r.serveroptimized
	}

	if len(r.serverrootVolumeSize) > 0 {
		server.server.rootVolumeSize = r.serverrootVolumeSize
	}

	return c.UpdateServer(r.id, &server)
}

type WrapInstanceServerServerRunner struct {
	serverDeploymentHref              string
	serverDescription                 string
	serverInstanceHref                string
	serverInstanceInputsValues        []string
	serverInstanceInputsNames         []string
	serverInstanceMultiCloudImageHref string
	serverInstanceServerTemplateHref  string
	serverName                        string
}

func (r *WrapInstanceServerServerRunner) Run(c *Client) (interface{}, error) {

	/** Handle server parameter **/
	var server ServerParam2

	// Load JSON if provided
	if len(r.serverJson) > 0 {
		if err := Json.Unmarshal(r.serverJson, &server); err != nil {
			return nil, fmt.Errorf("Could not load server JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.serverdeploymentHref) > 0 {
		server.server.deploymentHref = r.serverdeploymentHref
	}

	if len(r.serverdescription) > 0 {
		server.server.description = r.serverdescription
	}

	if len(r.serverinstancehref) > 0 {
		server.server.instance.href = r.serverinstancehref
	}

	if len(r.serverinstancemultiCloudImageHref) > 0 {
		server.server.instance.multiCloudImageHref = r.serverinstancemultiCloudImageHref
	}

	if len(r.serverinstanceserverTemplateHref) > 0 {
		server.server.instance.serverTemplateHref = r.serverinstanceserverTemplateHref
	}

	if len(r.servername) > 0 {
		server.server.name = r.servername
	}

	// Create enumarable fields from flags
	serverinstanceinputs := make(map[string]string, len(r.serverinstanceinputsNames))
	for i, n := range r.serverinstanceinputsNames {
		serverinstanceinputs[n] = r.serverinstanceinputsValues[i]
	}
	server.server.instance.inputs = serverinstanceinputs

	return c.WrapInstanceServer(&server)
}

// Register all Server actions
func registerServerCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Server", `Servers represent the notion of a server/machine from the RightScale's perspective`)

	CloneServerRunner := new(CloneServerServerRunner)
	CloneServerCmd := resCmd.Command("CloneServer", `Clones a given server.`)
	CloneServerRunner.Flag(`id`, ``).Required().StringVar(&CloneServerRunner.id)
	registry[CloneServerCmd.FullCommand()] = CloneServerRunner

	CreateServerRunner := new(CreateServerServerRunner)
	CreateServerCmd := resCmd.Command("CreateServer", `Creates a new server, and configures its corresponding "next" instance with the received parameters.`)
	CreateServerRunner.Flag(`server.deploymentHref`, `The href of the deployment to which the Server will be added.`).StringVar(&CreateServerRunner.serverDeploymentHref)
	CreateServerRunner.Flag(`server.description`, `The Server description.`).StringVar(&CreateServerRunner.serverDescription)
	CreateServerRunner.Flag(`server.instance.associatePublicIpAddress`, `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`).StringVar(&CreateServerRunner.serverInstanceAssociatePublicIpAddress)
	CreateServerRunner.Flag(`server.instance.cloudHref`, `The href of the cloud that the Server should be added to.`).Required().StringVar(&CreateServerRunner.serverInstanceCloudHref)
	CreateServerRunner.Flag(`server.instance.cloudSpecificAttributes.automaticInstanceStoreMapping`, `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`).StringVar(&CreateServerRunner.serverInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping)
	CreateServerRunner.Flag(`server.instance.cloudSpecificAttributes.iamInstanceProfile`, `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`).StringVar(&CreateServerRunner.serverInstanceCloudSpecificAttributesIamInstanceProfile)
	CreateServerRunner.Flag(`server.instance.cloudSpecificAttributes.rootVolumePerformance`, `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`).StringVar(&CreateServerRunner.serverInstanceCloudSpecificAttributesRootVolumePerformance)
	CreateServerRunner.Flag(`server.instance.cloudSpecificAttributes.rootVolumeSize`, `The size for root disk. Not supported in all Clouds.`).StringVar(&CreateServerRunner.serverInstanceCloudSpecificAttributesRootVolumeSize)
	CreateServerRunner.Flag(`server.instance.cloudSpecificAttributes.rootVolumeTypeUid`, `The type of root volume for instance. Only available on clouds supporting root volume type.`).StringVar(&CreateServerRunner.serverInstanceCloudSpecificAttributesRootVolumeTypeUid)
	CreateServerRunner.Flag(`server.instance.datacenterHref`, `The href of the Datacenter / Zone.`).StringVar(&CreateServerRunner.serverInstanceDatacenterHref)
	CreateServerRunner.Flag(`server.instance.imageHref`, `The href of the Image to use.`).StringVar(&CreateServerRunner.serverInstanceImageHref)
	CreateServerRunner.FlagPattern(`server\.instance\.inputs\.([a-z0-9_]+)`, ``).Capture(&CreateServerRunner.serverInstanceInputsNames).StringVar(&CreateServerRunner.serverInstanceInputsValues)
	CreateServerRunner.Flag(`server.instance.instanceTypeHref`, `The href of the Instance Type.`).StringVar(&CreateServerRunner.serverInstanceInstanceTypeHref)
	CreateServerRunner.Flag(`server.instance.ipForwardingEnabled`, `Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.`).StringVar(&CreateServerRunner.serverInstanceIpForwardingEnabled)
	CreateServerRunner.Flag(`server.instance.kernelImageHref`, `The href of the Kernel Image.`).StringVar(&CreateServerRunner.serverInstanceKernelImageHref)
	CreateServerRunner.Flag(`server.instance.multiCloudImageHref`, `The href of the Multi Cloud Image to use.`).StringVar(&CreateServerRunner.serverInstanceMultiCloudImageHref)
	CreateServerRunner.Flag(`server.instance.placementGroupHref`, `The href of the Placement Group.`).StringVar(&CreateServerRunner.serverInstancePlacementGroupHref)
	CreateServerRunner.Flag(`server.instance.ramdiskImageHref`, `The href of the Ramdisk Image.`).StringVar(&CreateServerRunner.serverInstanceRamdiskImageHref)
	CreateServerRunner.FlagPattern(`server\.instance\.item\.(\d+)`, `The hrefs of the security groups.`).Capture(&CreateServerRunner.serverInstanceItemPos).StringsVar(&CreateServerRunner.serverInstanceItem)
	CreateServerRunner.Flag(`server.instance.serverTemplateHref`, `The href of the Server Template.`).Required().StringVar(&CreateServerRunner.serverInstanceServerTemplateHref)
	CreateServerRunner.Flag(`server.instance.sshKeyHref`, `The href of the SSH key to use.`).StringVar(&CreateServerRunner.serverInstanceSshKeyHref)
	CreateServerRunner.FlagPattern(`server\.instance\.item\.(\d+)`, `The hrefs of the updated subnets.`).Capture(&CreateServerRunner.serverInstanceItemPos).StringsVar(&CreateServerRunner.serverInstanceItem)
	CreateServerRunner.Flag(`server.instance.userData`, `User data that RightScale automatically passes to your instance at boot time.`).StringVar(&CreateServerRunner.serverInstanceUserData)
	CreateServerRunner.Flag(`server.name`, `The name of the Server.`).Required().StringVar(&CreateServerRunner.serverName)
	CreateServerRunner.Flag(`server.optimized`, `A flag indicating whether Instances of this Server should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`).StringVar(&CreateServerRunner.serverOptimized)
	registry[CreateServerCmd.FullCommand()] = CreateServerRunner

	DestroyServerRunner := new(DestroyServerServerRunner)
	DestroyServerCmd := resCmd.Command("DestroyServer", `Deletes a given server.`)
	DestroyServerRunner.Flag(`id`, ``).Required().StringVar(&DestroyServerRunner.id)
	registry[DestroyServerCmd.FullCommand()] = DestroyServerRunner

	IndexServersRunner := new(IndexServersServerRunner)
	IndexServersCmd := resCmd.Command("IndexServers", `Lists servers.`)
	IndexServersRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexServersRunner.filterPos).StringsVar(&IndexServersRunner.filter)
	IndexServersRunner.Flag(`view`, ``).StringVar(&IndexServersRunner.view)
	registry[IndexServersCmd.FullCommand()] = IndexServersRunner

	LaunchServerRunner := new(LaunchServerServerRunner)
	LaunchServerCmd := resCmd.Command("LaunchServer", `Launches the "next" instance of this server`)
	LaunchServerRunner.Flag(`id`, ``).Required().StringVar(&LaunchServerRunner.id)
	registry[LaunchServerCmd.FullCommand()] = LaunchServerRunner

	ShowServerRunner := new(ShowServerServerRunner)
	ShowServerCmd := resCmd.Command("ShowServer", `Shows the information of a single server.`)
	ShowServerRunner.Flag(`id`, ``).Required().StringVar(&ShowServerRunner.id)
	ShowServerRunner.Flag(`view`, ``).StringVar(&ShowServerRunner.view)
	registry[ShowServerCmd.FullCommand()] = ShowServerRunner

	TerminateServerRunner := new(TerminateServerServerRunner)
	TerminateServerCmd := resCmd.Command("TerminateServer", `Terminates the current instance of this server`)
	TerminateServerRunner.Flag(`id`, ``).Required().StringVar(&TerminateServerRunner.id)
	registry[TerminateServerCmd.FullCommand()] = TerminateServerRunner

	UpdateServerRunner := new(UpdateServerServerRunner)
	UpdateServerCmd := resCmd.Command("UpdateServer", `Updates attributes of a single server.`)
	UpdateServerRunner.Flag(`id`, ``).Required().StringVar(&UpdateServerRunner.id)
	UpdateServerRunner.Flag(`server.automaticInstanceStoreMapping`, `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`).StringVar(&UpdateServerRunner.serverAutomaticInstanceStoreMapping)
	UpdateServerRunner.Flag(`server.description`, `The updated description for the server.`).StringVar(&UpdateServerRunner.serverDescription)
	UpdateServerRunner.Flag(`server.name`, `The updated server name.`).StringVar(&UpdateServerRunner.serverName)
	UpdateServerRunner.Flag(`server.optimized`, `A flag indicating whether Instances of this Server should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`).StringVar(&UpdateServerRunner.serverOptimized)
	UpdateServerRunner.Flag(`server.rootVolumeSize`, `The size for root disk. Not supported in all Clouds.`).StringVar(&UpdateServerRunner.serverRootVolumeSize)
	registry[UpdateServerCmd.FullCommand()] = UpdateServerRunner

	WrapInstanceServerRunner := new(WrapInstanceServerServerRunner)
	WrapInstanceServerCmd := resCmd.Command("WrapInstanceServer", `Wrap an existing instance and set current instance for new server`)
	WrapInstanceServerRunner.Flag(`server.deploymentHref`, `The href of the deployment to which the Server will be added.`).StringVar(&WrapInstanceServerRunner.serverDeploymentHref)
	WrapInstanceServerRunner.Flag(`server.description`, `The Server description.`).StringVar(&WrapInstanceServerRunner.serverDescription)
	WrapInstanceServerRunner.Flag(`server.instance.href`, `The href of the Instance around which the server should be created.`).Required().StringVar(&WrapInstanceServerRunner.serverInstanceHref)
	WrapInstanceServerRunner.FlagPattern(`server\.instance\.inputs\.([a-z0-9_]+)`, ``).Capture(&WrapInstanceServerRunner.serverInstanceInputsNames).StringVar(&WrapInstanceServerRunner.serverInstanceInputsValues)
	WrapInstanceServerRunner.Flag(`server.instance.multiCloudImageHref`, `The href of the Multi Cloud Image to use.`).StringVar(&WrapInstanceServerRunner.serverInstanceMultiCloudImageHref)
	WrapInstanceServerRunner.Flag(`server.instance.serverTemplateHref`, `The href of the Server Template.`).Required().StringVar(&WrapInstanceServerRunner.serverInstanceServerTemplateHref)
	WrapInstanceServerRunner.Flag(`server.name`, `The name of the Server.`).StringVar(&WrapInstanceServerRunner.serverName)
	registry[WrapInstanceServerCmd.FullCommand()] = WrapInstanceServerRunner
}

/****** ServerArray ******/

type CloneServerArrayServerArrayRunner struct {
	id string
}

func (r *CloneServerArrayServerArrayRunner) Run(c *Client) (interface{}, error) {
	return c.CloneServerArray(r.id)
}

type CreateServerArrayServerArrayRunner struct {
	serverArrayArrayType                                                    string
	serverArrayItemDatacenterHref                                           []string
	serverArrayItemDatacenterHrefPos                                        []string
	serverArrayItemMax                                                      []string
	serverArrayItemMaxPos                                                   []string
	serverArrayItemWeight                                                   []string
	serverArrayItemWeightPos                                                []string
	serverArrayDeploymentHref                                               string
	serverArrayDescription                                                  string
	serverArrayElasticityParamsAlertSpecificParamsDecisionThreshold         string
	serverArrayElasticityParamsAlertSpecificParamsVotersTagPredicate        string
	serverArrayElasticityParamsBoundsMaxCount                               string
	serverArrayElasticityParamsBoundsMinCount                               string
	serverArrayElasticityParamsPacingResizeCalmTime                         string
	serverArrayElasticityParamsPacingResizeDownBy                           string
	serverArrayElasticityParamsPacingResizeUpBy                             string
	serverArrayElasticityParamsQueueSpecificParamsCollectAuditEntries       string
	serverArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm          string
	serverArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge             string
	serverArrayElasticityParamsQueueSpecificParamsItemAgeRegexp             string
	serverArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance string
	serverArrayElasticityParamsItemDay                                      []string
	serverArrayElasticityParamsItemDayPos                                   []string
	serverArrayElasticityParamsItemMaxCount                                 []string
	serverArrayElasticityParamsItemMaxCountPos                              []string
	serverArrayElasticityParamsItemMinCount                                 []string
	serverArrayElasticityParamsItemMinCountPos                              []string
	serverArrayElasticityParamsItemTime                                     []string
	serverArrayElasticityParamsItemTimePos                                  []string
	serverArrayInstanceAssociatePublicIpAddress                             string
	serverArrayInstanceCloudHref                                            string
	serverArrayInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping string
	serverArrayInstanceCloudSpecificAttributesIamInstanceProfile            string
	serverArrayInstanceCloudSpecificAttributesRootVolumePerformance         string
	serverArrayInstanceCloudSpecificAttributesRootVolumeSize                string
	serverArrayInstanceCloudSpecificAttributesRootVolumeTypeUid             string
	serverArrayInstanceDatacenterHref                                       string
	serverArrayInstanceImageHref                                            string
	serverArrayInstanceInputsValues                                         []string
	serverArrayInstanceInputsNames                                          []string
	serverArrayInstanceInstanceTypeHref                                     string
	serverArrayInstanceIpForwardingEnabled                                  string
	serverArrayInstanceKernelImageHref                                      string
	serverArrayInstanceMultiCloudImageHref                                  string
	serverArrayInstancePlacementGroupHref                                   string
	serverArrayInstanceRamdiskImageHref                                     string
	serverArrayInstanceItem                                                 []string
	serverArrayInstanceItemPos                                              []string
	serverArrayInstanceServerTemplateHref                                   string
	serverArrayInstanceSshKeyHref                                           string
	serverArrayInstanceItem                                                 []string
	serverArrayInstanceItemPos                                              []string
	serverArrayInstanceUserData                                             string
	serverArrayName                                                         string
	serverArrayOptimized                                                    string
	serverArrayState                                                        string
}

func (r *CreateServerArrayServerArrayRunner) Run(c *Client) (interface{}, error) {

	/** Handle serverArray parameter **/
	var serverArray ServerArrayParam

	// Load JSON if provided
	if len(r.serverArrayJson) > 0 {
		if err := Json.Unmarshal(r.serverArrayJson, &serverArray); err != nil {
			return nil, fmt.Errorf("Could not load serverArray JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.serverArrayarrayType) > 0 {
		serverArray.serverArray.arrayType = r.serverArrayarrayType
	}

	if len(r.serverArraydatacenterPolicyitemdatacenterHref) > 0 {
		serverArray.serverArray.datacenterPolicy.item.datacenterHref = r.serverArraydatacenterPolicyitemdatacenterHref
	}

	if len(r.serverArraydatacenterPolicyitemmax) > 0 {
		serverArray.serverArray.datacenterPolicy.item.max = r.serverArraydatacenterPolicyitemmax
	}

	if len(r.serverArraydatacenterPolicyitemweight) > 0 {
		serverArray.serverArray.datacenterPolicy.item.weight = r.serverArraydatacenterPolicyitemweight
	}

	if len(r.serverArraydeploymentHref) > 0 {
		serverArray.serverArray.deploymentHref = r.serverArraydeploymentHref
	}

	if len(r.serverArraydescription) > 0 {
		serverArray.serverArray.description = r.serverArraydescription
	}

	if len(r.serverArrayelasticityParamsalertSpecificParamsdecisionThreshold) > 0 {
		serverArray.serverArray.elasticityParams.alertSpecificParams.decisionThreshold = r.serverArrayelasticityParamsalertSpecificParamsdecisionThreshold
	}

	if len(r.serverArrayelasticityParamsalertSpecificParamsvotersTagPredicate) > 0 {
		serverArray.serverArray.elasticityParams.alertSpecificParams.votersTagPredicate = r.serverArrayelasticityParamsalertSpecificParamsvotersTagPredicate
	}

	if len(r.serverArrayelasticityParamsboundsmaxCount) > 0 {
		serverArray.serverArray.elasticityParams.bounds.maxCount = r.serverArrayelasticityParamsboundsmaxCount
	}

	if len(r.serverArrayelasticityParamsboundsminCount) > 0 {
		serverArray.serverArray.elasticityParams.bounds.minCount = r.serverArrayelasticityParamsboundsminCount
	}

	if len(r.serverArrayelasticityParamspacingresizeCalmTime) > 0 {
		serverArray.serverArray.elasticityParams.pacing.resizeCalmTime = r.serverArrayelasticityParamspacingresizeCalmTime
	}

	if len(r.serverArrayelasticityParamspacingresizeDownBy) > 0 {
		serverArray.serverArray.elasticityParams.pacing.resizeDownBy = r.serverArrayelasticityParamspacingresizeDownBy
	}

	if len(r.serverArrayelasticityParamspacingresizeUpBy) > 0 {
		serverArray.serverArray.elasticityParams.pacing.resizeUpBy = r.serverArrayelasticityParamspacingresizeUpBy
	}

	if len(r.serverArrayelasticityParamsqueueSpecificParamscollectAuditEntries) > 0 {
		serverArray.serverArray.elasticityParams.queueSpecificParams.collectAuditEntries = r.serverArrayelasticityParamsqueueSpecificParamscollectAuditEntries
	}

	if len(r.serverArrayelasticityParamsqueueSpecificParamsitemAgealgorithm) > 0 {
		serverArray.serverArray.elasticityParams.queueSpecificParams.itemAge.algorithm = r.serverArrayelasticityParamsqueueSpecificParamsitemAgealgorithm
	}

	if len(r.serverArrayelasticityParamsqueueSpecificParamsitemAgemaxAge) > 0 {
		serverArray.serverArray.elasticityParams.queueSpecificParams.itemAge.maxAge = r.serverArrayelasticityParamsqueueSpecificParamsitemAgemaxAge
	}

	if len(r.serverArrayelasticityParamsqueueSpecificParamsitemAgeregexp) > 0 {
		serverArray.serverArray.elasticityParams.queueSpecificParams.itemAge.regexp = r.serverArrayelasticityParamsqueueSpecificParamsitemAgeregexp
	}

	if len(r.serverArrayelasticityParamsqueueSpecificParamsqueueSizeitemsPerInstance) > 0 {
		serverArray.serverArray.elasticityParams.queueSpecificParams.queueSize.itemsPerInstance = r.serverArrayelasticityParamsqueueSpecificParamsqueueSizeitemsPerInstance
	}

	if len(r.serverArrayelasticityParamsscheduleitemday) > 0 {
		serverArray.serverArray.elasticityParams.schedule.item.day = r.serverArrayelasticityParamsscheduleitemday
	}

	if len(r.serverArrayelasticityParamsscheduleitemmaxCount) > 0 {
		serverArray.serverArray.elasticityParams.schedule.item.maxCount = r.serverArrayelasticityParamsscheduleitemmaxCount
	}

	if len(r.serverArrayelasticityParamsscheduleitemminCount) > 0 {
		serverArray.serverArray.elasticityParams.schedule.item.minCount = r.serverArrayelasticityParamsscheduleitemminCount
	}

	if len(r.serverArrayelasticityParamsscheduleitemtime) > 0 {
		serverArray.serverArray.elasticityParams.schedule.item.time = r.serverArrayelasticityParamsscheduleitemtime
	}

	if len(r.serverArrayinstanceassociatePublicIpAddress) > 0 {
		serverArray.serverArray.instance.associatePublicIpAddress = r.serverArrayinstanceassociatePublicIpAddress
	}

	if len(r.serverArrayinstancecloudHref) > 0 {
		serverArray.serverArray.instance.cloudHref = r.serverArrayinstancecloudHref
	}

	if len(r.serverArrayinstancecloudSpecificAttributesautomaticInstanceStoreMapping) > 0 {
		serverArray.serverArray.instance.cloudSpecificAttributes.automaticInstanceStoreMapping = r.serverArrayinstancecloudSpecificAttributesautomaticInstanceStoreMapping
	}

	if len(r.serverArrayinstancecloudSpecificAttributesiamInstanceProfile) > 0 {
		serverArray.serverArray.instance.cloudSpecificAttributes.iamInstanceProfile = r.serverArrayinstancecloudSpecificAttributesiamInstanceProfile
	}

	if len(r.serverArrayinstancecloudSpecificAttributesrootVolumePerformance) > 0 {
		serverArray.serverArray.instance.cloudSpecificAttributes.rootVolumePerformance = r.serverArrayinstancecloudSpecificAttributesrootVolumePerformance
	}

	if len(r.serverArrayinstancecloudSpecificAttributesrootVolumeSize) > 0 {
		serverArray.serverArray.instance.cloudSpecificAttributes.rootVolumeSize = r.serverArrayinstancecloudSpecificAttributesrootVolumeSize
	}

	if len(r.serverArrayinstancecloudSpecificAttributesrootVolumeTypeUid) > 0 {
		serverArray.serverArray.instance.cloudSpecificAttributes.rootVolumeTypeUid = r.serverArrayinstancecloudSpecificAttributesrootVolumeTypeUid
	}

	if len(r.serverArrayinstancedatacenterHref) > 0 {
		serverArray.serverArray.instance.datacenterHref = r.serverArrayinstancedatacenterHref
	}

	if len(r.serverArrayinstanceimageHref) > 0 {
		serverArray.serverArray.instance.imageHref = r.serverArrayinstanceimageHref
	}

	if len(r.serverArrayinstanceinstanceTypeHref) > 0 {
		serverArray.serverArray.instance.instanceTypeHref = r.serverArrayinstanceinstanceTypeHref
	}

	if len(r.serverArrayinstanceipForwardingEnabled) > 0 {
		serverArray.serverArray.instance.ipForwardingEnabled = r.serverArrayinstanceipForwardingEnabled
	}

	if len(r.serverArrayinstancekernelImageHref) > 0 {
		serverArray.serverArray.instance.kernelImageHref = r.serverArrayinstancekernelImageHref
	}

	if len(r.serverArrayinstancemultiCloudImageHref) > 0 {
		serverArray.serverArray.instance.multiCloudImageHref = r.serverArrayinstancemultiCloudImageHref
	}

	if len(r.serverArrayinstanceplacementGroupHref) > 0 {
		serverArray.serverArray.instance.placementGroupHref = r.serverArrayinstanceplacementGroupHref
	}

	if len(r.serverArrayinstanceramdiskImageHref) > 0 {
		serverArray.serverArray.instance.ramdiskImageHref = r.serverArrayinstanceramdiskImageHref
	}

	if len(r.serverArrayinstancesecurityGroupHrefsitem) > 0 {
		serverArray.serverArray.instance.securityGroupHrefs.item = r.serverArrayinstancesecurityGroupHrefsitem
	}

	if len(r.serverArrayinstanceserverTemplateHref) > 0 {
		serverArray.serverArray.instance.serverTemplateHref = r.serverArrayinstanceserverTemplateHref
	}

	if len(r.serverArrayinstancesshKeyHref) > 0 {
		serverArray.serverArray.instance.sshKeyHref = r.serverArrayinstancesshKeyHref
	}

	if len(r.serverArrayinstancesubnetHrefsitem) > 0 {
		serverArray.serverArray.instance.subnetHrefs.item = r.serverArrayinstancesubnetHrefsitem
	}

	if len(r.serverArrayinstanceuserData) > 0 {
		serverArray.serverArray.instance.userData = r.serverArrayinstanceuserData
	}

	if len(r.serverArrayname) > 0 {
		serverArray.serverArray.name = r.serverArrayname
	}

	if len(r.serverArrayoptimized) > 0 {
		serverArray.serverArray.optimized = r.serverArrayoptimized
	}

	if len(r.serverArraystate) > 0 {
		serverArray.serverArray.state = r.serverArraystate
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxserverArraydatacenterPolicyitemdatacenterHrefPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.serverArraydatacenterPolicyitemdatacenterHrefPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.datacenterPolicy.item.datacenterHref field of datacenterPolicy array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.datacenterPolicy.item.datacenterHref field of datacenterPolicy array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxserverArraydatacenterPolicyitemdatacenterHrefPos {
			maxserverArraydatacenterPolicyitemdatacenterHrefPos = pos
		}
	}
	if len(r.serverArraydatacenterPolicyitemdatacenterHref) != maxserverArraydatacenterPolicyitemdatacenterHrefPos {
		return nil, fmt.Errof("Invalid flags for serverArray.datacenterPolicy.item.datacenterHref field of datacenterPolicy array, %s were defined but max position is %s",
			len(r.serverArraydatacenterPolicyitemdatacenterHref), maxserverArraydatacenterPolicyitemdatacenterHrefPos)
	}

	maxserverArraydatacenterPolicyitemmaxPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.serverArraydatacenterPolicyitemmaxPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.datacenterPolicy.item.max field of datacenterPolicy array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.datacenterPolicy.item.max field of datacenterPolicy array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxserverArraydatacenterPolicyitemmaxPos {
			maxserverArraydatacenterPolicyitemmaxPos = pos
		}
	}
	if len(r.serverArraydatacenterPolicyitemmax) != maxserverArraydatacenterPolicyitemmaxPos {
		return nil, fmt.Errof("Invalid flags for serverArray.datacenterPolicy.item.max field of datacenterPolicy array, %s were defined but max position is %s",
			len(r.serverArraydatacenterPolicyitemmax), maxserverArraydatacenterPolicyitemmaxPos)
	}

	maxserverArraydatacenterPolicyitemweightPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.serverArraydatacenterPolicyitemweightPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.datacenterPolicy.item.weight field of datacenterPolicy array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.datacenterPolicy.item.weight field of datacenterPolicy array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxserverArraydatacenterPolicyitemweightPos {
			maxserverArraydatacenterPolicyitemweightPos = pos
		}
	}
	if len(r.serverArraydatacenterPolicyitemweight) != maxserverArraydatacenterPolicyitemweightPos {
		return nil, fmt.Errof("Invalid flags for serverArray.datacenterPolicy.item.weight field of datacenterPolicy array, %s were defined but max position is %s",
			len(r.serverArraydatacenterPolicyitemweight), maxserverArraydatacenterPolicyitemweightPos)
	}
	if maxserverArraydatacenterPolicyitemdatacenterHrefPos > -1 {
		serverArraydatacenterPolicy := make([]*DatacenterPolicy, maxserverArraydatacenterPolicyitemdatacenterHrefPos+1)
		for i := 0; i < maxserverArraydatacenterPolicyitemdatacenterHrefPos+1; i++ {
			serverArraydatacenterPolicy[i] = DatacenterPolicy{
			//TBD
			}
		}
		serverArray.datacenterPolicy = serverArraydatacenterPolicy
	}

	maxserverArrayelasticityParamsscheduleitemdayPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.serverArrayelasticityParamsscheduleitemdayPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.elasticityParams.schedule.item.day field of schedule array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.elasticityParams.schedule.item.day field of schedule array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxserverArrayelasticityParamsscheduleitemdayPos {
			maxserverArrayelasticityParamsscheduleitemdayPos = pos
		}
	}
	if len(r.serverArrayelasticityParamsscheduleitemday) != maxserverArrayelasticityParamsscheduleitemdayPos {
		return nil, fmt.Errof("Invalid flags for serverArray.elasticityParams.schedule.item.day field of schedule array, %s were defined but max position is %s",
			len(r.serverArrayelasticityParamsscheduleitemday), maxserverArrayelasticityParamsscheduleitemdayPos)
	}

	maxserverArrayelasticityParamsscheduleitemmaxCountPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.serverArrayelasticityParamsscheduleitemmaxCountPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.elasticityParams.schedule.item.maxCount field of schedule array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.elasticityParams.schedule.item.maxCount field of schedule array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxserverArrayelasticityParamsscheduleitemmaxCountPos {
			maxserverArrayelasticityParamsscheduleitemmaxCountPos = pos
		}
	}
	if len(r.serverArrayelasticityParamsscheduleitemmaxCount) != maxserverArrayelasticityParamsscheduleitemmaxCountPos {
		return nil, fmt.Errof("Invalid flags for serverArray.elasticityParams.schedule.item.maxCount field of schedule array, %s were defined but max position is %s",
			len(r.serverArrayelasticityParamsscheduleitemmaxCount), maxserverArrayelasticityParamsscheduleitemmaxCountPos)
	}

	maxserverArrayelasticityParamsscheduleitemminCountPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.serverArrayelasticityParamsscheduleitemminCountPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.elasticityParams.schedule.item.minCount field of schedule array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.elasticityParams.schedule.item.minCount field of schedule array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxserverArrayelasticityParamsscheduleitemminCountPos {
			maxserverArrayelasticityParamsscheduleitemminCountPos = pos
		}
	}
	if len(r.serverArrayelasticityParamsscheduleitemminCount) != maxserverArrayelasticityParamsscheduleitemminCountPos {
		return nil, fmt.Errof("Invalid flags for serverArray.elasticityParams.schedule.item.minCount field of schedule array, %s were defined but max position is %s",
			len(r.serverArrayelasticityParamsscheduleitemminCount), maxserverArrayelasticityParamsscheduleitemminCountPos)
	}

	maxserverArrayelasticityParamsscheduleitemtimePos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.serverArrayelasticityParamsscheduleitemtimePos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.elasticityParams.schedule.item.time field of schedule array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.elasticityParams.schedule.item.time field of schedule array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxserverArrayelasticityParamsscheduleitemtimePos {
			maxserverArrayelasticityParamsscheduleitemtimePos = pos
		}
	}
	if len(r.serverArrayelasticityParamsscheduleitemtime) != maxserverArrayelasticityParamsscheduleitemtimePos {
		return nil, fmt.Errof("Invalid flags for serverArray.elasticityParams.schedule.item.time field of schedule array, %s were defined but max position is %s",
			len(r.serverArrayelasticityParamsscheduleitemtime), maxserverArrayelasticityParamsscheduleitemtimePos)
	}
	if maxserverArrayelasticityParamsscheduleitemdayPos > -1 {
		serverArrayelasticityParamsschedule := make([]*Schedule, maxserverArrayelasticityParamsscheduleitemdayPos+1)
		for i := 0; i < maxserverArrayelasticityParamsscheduleitemdayPos+1; i++ {
			serverArrayelasticityParamsschedule[i] = Schedule{
			//TBD
			}
		}
		serverArray.elasticityParams.schedule = serverArrayelasticityParamsschedule
	}

	maxserverArrayinstancesecurityGroupHrefsitemPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.serverArrayinstancesecurityGroupHrefsitemPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.instance.securityGroupHrefs.item field of securityGroupHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.instance.securityGroupHrefs.item field of securityGroupHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxserverArrayinstancesecurityGroupHrefsitemPos {
			maxserverArrayinstancesecurityGroupHrefsitemPos = pos
		}
	}
	if len(r.serverArrayinstancesecurityGroupHrefsitem) != maxserverArrayinstancesecurityGroupHrefsitemPos {
		return nil, fmt.Errof("Invalid flags for serverArray.instance.securityGroupHrefs.item field of securityGroupHrefs array, %s were defined but max position is %s",
			len(r.serverArrayinstancesecurityGroupHrefsitem), maxserverArrayinstancesecurityGroupHrefsitemPos)
	}
	if maxserverArrayinstancesecurityGroupHrefsitemPos > -1 {
		serverArrayinstancesecurityGroupHrefs := make([]*string, maxserverArrayinstancesecurityGroupHrefsitemPos+1)
		for i := 0; i < maxserverArrayinstancesecurityGroupHrefsitemPos+1; i++ {
			serverArrayinstancesecurityGroupHrefs[i] = string{
			//TBD
			}
		}
		serverArray.instance.securityGroupHrefs = serverArrayinstancesecurityGroupHrefs
	}

	maxserverArrayinstancesubnetHrefsitemPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.serverArrayinstancesubnetHrefsitemPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.instance.subnetHrefs.item field of subnetHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.instance.subnetHrefs.item field of subnetHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxserverArrayinstancesubnetHrefsitemPos {
			maxserverArrayinstancesubnetHrefsitemPos = pos
		}
	}
	if len(r.serverArrayinstancesubnetHrefsitem) != maxserverArrayinstancesubnetHrefsitemPos {
		return nil, fmt.Errof("Invalid flags for serverArray.instance.subnetHrefs.item field of subnetHrefs array, %s were defined but max position is %s",
			len(r.serverArrayinstancesubnetHrefsitem), maxserverArrayinstancesubnetHrefsitemPos)
	}
	if maxserverArrayinstancesubnetHrefsitemPos > -1 {
		serverArrayinstancesubnetHrefs := make([]*string, maxserverArrayinstancesubnetHrefsitemPos+1)
		for i := 0; i < maxserverArrayinstancesubnetHrefsitemPos+1; i++ {
			serverArrayinstancesubnetHrefs[i] = string{
			//TBD
			}
		}
		serverArray.instance.subnetHrefs = serverArrayinstancesubnetHrefs
	}

	// Create enumarable fields from flags
	serverArrayinstanceinputs := make(map[string]string, len(r.serverArrayinstanceinputsNames))
	for i, n := range r.serverArrayinstanceinputsNames {
		serverArrayinstanceinputs[n] = r.serverArrayinstanceinputsValues[i]
	}
	serverArray.serverArray.instance.inputs = serverArrayinstanceinputs

	return c.CreateServerArray(&serverArray)
}

type CurrentInstancesServerArrayServerArrayRunner struct {
	id string
}

func (r *CurrentInstancesServerArrayServerArrayRunner) Run(c *Client) (interface{}, error) {
	return c.CurrentInstancesServerArray(r.id)
}

type DestroyServerArrayServerArrayRunner struct {
	id string
}

func (r *DestroyServerArrayServerArrayRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyServerArray(r.id)
}

type IndexServerArraysServerArrayRunner struct {
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexServerArraysServerArrayRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexServerArrays(filter, r.view)
}

type LaunchServerArrayServerArrayRunner struct {
	id string
}

func (r *LaunchServerArrayServerArrayRunner) Run(c *Client) (interface{}, error) {
	return c.LaunchServerArray(r.id)
}

type MultiRunExecutableServerArraysServerArrayRunner struct {
	id string
}

func (r *MultiRunExecutableServerArraysServerArrayRunner) Run(c *Client) (interface{}, error) {
	return c.MultiRunExecutableServerArrays(r.id)
}

type MultiTerminateServerArraysServerArrayRunner struct {
	id string
}

func (r *MultiTerminateServerArraysServerArrayRunner) Run(c *Client) (interface{}, error) {
	return c.MultiTerminateServerArrays(r.id)
}

type ShowServerArrayServerArrayRunner struct {
	id   string
	view string
}

func (r *ShowServerArrayServerArrayRunner) Run(c *Client) (interface{}, error) {
	return c.ShowServerArray(r.id, r.view)
}

type UpdateServerArrayServerArrayRunner struct {
	id                                                                      string
	serverArrayArrayType                                                    string
	serverArrayItemDatacenterHref                                           []string
	serverArrayItemDatacenterHrefPos                                        []string
	serverArrayItemMax                                                      []string
	serverArrayItemMaxPos                                                   []string
	serverArrayItemWeight                                                   []string
	serverArrayItemWeightPos                                                []string
	serverArrayDeploymentHref                                               string
	serverArrayDescription                                                  string
	serverArrayElasticityParamsAlertSpecificParamsDecisionThreshold         string
	serverArrayElasticityParamsAlertSpecificParamsVotersTagPredicate        string
	serverArrayElasticityParamsBoundsMaxCount                               string
	serverArrayElasticityParamsBoundsMinCount                               string
	serverArrayElasticityParamsPacingResizeCalmTime                         string
	serverArrayElasticityParamsPacingResizeDownBy                           string
	serverArrayElasticityParamsPacingResizeUpBy                             string
	serverArrayElasticityParamsQueueSpecificParamsCollectAuditEntries       string
	serverArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm          string
	serverArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge             string
	serverArrayElasticityParamsQueueSpecificParamsItemAgeRegexp             string
	serverArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance string
	serverArrayElasticityParamsItemDay                                      []string
	serverArrayElasticityParamsItemDayPos                                   []string
	serverArrayElasticityParamsItemMaxCount                                 []string
	serverArrayElasticityParamsItemMaxCountPos                              []string
	serverArrayElasticityParamsItemMinCount                                 []string
	serverArrayElasticityParamsItemMinCountPos                              []string
	serverArrayElasticityParamsItemTime                                     []string
	serverArrayElasticityParamsItemTimePos                                  []string
	serverArrayName                                                         string
	serverArrayOptimized                                                    string
	serverArrayState                                                        string
}

func (r *UpdateServerArrayServerArrayRunner) Run(c *Client) (interface{}, error) {

	/** Handle serverArray parameter **/
	var serverArray ServerArrayParam2

	// Load JSON if provided
	if len(r.serverArrayJson) > 0 {
		if err := Json.Unmarshal(r.serverArrayJson, &serverArray); err != nil {
			return nil, fmt.Errorf("Could not load serverArray JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.serverArrayarrayType) > 0 {
		serverArray.serverArray.arrayType = r.serverArrayarrayType
	}

	if len(r.serverArraydatacenterPolicyitemdatacenterHref) > 0 {
		serverArray.serverArray.datacenterPolicy.item.datacenterHref = r.serverArraydatacenterPolicyitemdatacenterHref
	}

	if len(r.serverArraydatacenterPolicyitemmax) > 0 {
		serverArray.serverArray.datacenterPolicy.item.max = r.serverArraydatacenterPolicyitemmax
	}

	if len(r.serverArraydatacenterPolicyitemweight) > 0 {
		serverArray.serverArray.datacenterPolicy.item.weight = r.serverArraydatacenterPolicyitemweight
	}

	if len(r.serverArraydeploymentHref) > 0 {
		serverArray.serverArray.deploymentHref = r.serverArraydeploymentHref
	}

	if len(r.serverArraydescription) > 0 {
		serverArray.serverArray.description = r.serverArraydescription
	}

	if len(r.serverArrayelasticityParamsalertSpecificParamsdecisionThreshold) > 0 {
		serverArray.serverArray.elasticityParams.alertSpecificParams.decisionThreshold = r.serverArrayelasticityParamsalertSpecificParamsdecisionThreshold
	}

	if len(r.serverArrayelasticityParamsalertSpecificParamsvotersTagPredicate) > 0 {
		serverArray.serverArray.elasticityParams.alertSpecificParams.votersTagPredicate = r.serverArrayelasticityParamsalertSpecificParamsvotersTagPredicate
	}

	if len(r.serverArrayelasticityParamsboundsmaxCount) > 0 {
		serverArray.serverArray.elasticityParams.bounds.maxCount = r.serverArrayelasticityParamsboundsmaxCount
	}

	if len(r.serverArrayelasticityParamsboundsminCount) > 0 {
		serverArray.serverArray.elasticityParams.bounds.minCount = r.serverArrayelasticityParamsboundsminCount
	}

	if len(r.serverArrayelasticityParamspacingresizeCalmTime) > 0 {
		serverArray.serverArray.elasticityParams.pacing.resizeCalmTime = r.serverArrayelasticityParamspacingresizeCalmTime
	}

	if len(r.serverArrayelasticityParamspacingresizeDownBy) > 0 {
		serverArray.serverArray.elasticityParams.pacing.resizeDownBy = r.serverArrayelasticityParamspacingresizeDownBy
	}

	if len(r.serverArrayelasticityParamspacingresizeUpBy) > 0 {
		serverArray.serverArray.elasticityParams.pacing.resizeUpBy = r.serverArrayelasticityParamspacingresizeUpBy
	}

	if len(r.serverArrayelasticityParamsqueueSpecificParamscollectAuditEntries) > 0 {
		serverArray.serverArray.elasticityParams.queueSpecificParams.collectAuditEntries = r.serverArrayelasticityParamsqueueSpecificParamscollectAuditEntries
	}

	if len(r.serverArrayelasticityParamsqueueSpecificParamsitemAgealgorithm) > 0 {
		serverArray.serverArray.elasticityParams.queueSpecificParams.itemAge.algorithm = r.serverArrayelasticityParamsqueueSpecificParamsitemAgealgorithm
	}

	if len(r.serverArrayelasticityParamsqueueSpecificParamsitemAgemaxAge) > 0 {
		serverArray.serverArray.elasticityParams.queueSpecificParams.itemAge.maxAge = r.serverArrayelasticityParamsqueueSpecificParamsitemAgemaxAge
	}

	if len(r.serverArrayelasticityParamsqueueSpecificParamsitemAgeregexp) > 0 {
		serverArray.serverArray.elasticityParams.queueSpecificParams.itemAge.regexp = r.serverArrayelasticityParamsqueueSpecificParamsitemAgeregexp
	}

	if len(r.serverArrayelasticityParamsqueueSpecificParamsqueueSizeitemsPerInstance) > 0 {
		serverArray.serverArray.elasticityParams.queueSpecificParams.queueSize.itemsPerInstance = r.serverArrayelasticityParamsqueueSpecificParamsqueueSizeitemsPerInstance
	}

	if len(r.serverArrayelasticityParamsscheduleitemday) > 0 {
		serverArray.serverArray.elasticityParams.schedule.item.day = r.serverArrayelasticityParamsscheduleitemday
	}

	if len(r.serverArrayelasticityParamsscheduleitemmaxCount) > 0 {
		serverArray.serverArray.elasticityParams.schedule.item.maxCount = r.serverArrayelasticityParamsscheduleitemmaxCount
	}

	if len(r.serverArrayelasticityParamsscheduleitemminCount) > 0 {
		serverArray.serverArray.elasticityParams.schedule.item.minCount = r.serverArrayelasticityParamsscheduleitemminCount
	}

	if len(r.serverArrayelasticityParamsscheduleitemtime) > 0 {
		serverArray.serverArray.elasticityParams.schedule.item.time = r.serverArrayelasticityParamsscheduleitemtime
	}

	if len(r.serverArrayname) > 0 {
		serverArray.serverArray.name = r.serverArrayname
	}

	if len(r.serverArrayoptimized) > 0 {
		serverArray.serverArray.optimized = r.serverArrayoptimized
	}

	if len(r.serverArraystate) > 0 {
		serverArray.serverArray.state = r.serverArraystate
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxserverArraydatacenterPolicyitemdatacenterHrefPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.serverArraydatacenterPolicyitemdatacenterHrefPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.datacenterPolicy.item.datacenterHref field of datacenterPolicy array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.datacenterPolicy.item.datacenterHref field of datacenterPolicy array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxserverArraydatacenterPolicyitemdatacenterHrefPos {
			maxserverArraydatacenterPolicyitemdatacenterHrefPos = pos
		}
	}
	if len(r.serverArraydatacenterPolicyitemdatacenterHref) != maxserverArraydatacenterPolicyitemdatacenterHrefPos {
		return nil, fmt.Errof("Invalid flags for serverArray.datacenterPolicy.item.datacenterHref field of datacenterPolicy array, %s were defined but max position is %s",
			len(r.serverArraydatacenterPolicyitemdatacenterHref), maxserverArraydatacenterPolicyitemdatacenterHrefPos)
	}

	maxserverArraydatacenterPolicyitemmaxPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.serverArraydatacenterPolicyitemmaxPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.datacenterPolicy.item.max field of datacenterPolicy array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.datacenterPolicy.item.max field of datacenterPolicy array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxserverArraydatacenterPolicyitemmaxPos {
			maxserverArraydatacenterPolicyitemmaxPos = pos
		}
	}
	if len(r.serverArraydatacenterPolicyitemmax) != maxserverArraydatacenterPolicyitemmaxPos {
		return nil, fmt.Errof("Invalid flags for serverArray.datacenterPolicy.item.max field of datacenterPolicy array, %s were defined but max position is %s",
			len(r.serverArraydatacenterPolicyitemmax), maxserverArraydatacenterPolicyitemmaxPos)
	}

	maxserverArraydatacenterPolicyitemweightPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.serverArraydatacenterPolicyitemweightPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.datacenterPolicy.item.weight field of datacenterPolicy array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.datacenterPolicy.item.weight field of datacenterPolicy array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxserverArraydatacenterPolicyitemweightPos {
			maxserverArraydatacenterPolicyitemweightPos = pos
		}
	}
	if len(r.serverArraydatacenterPolicyitemweight) != maxserverArraydatacenterPolicyitemweightPos {
		return nil, fmt.Errof("Invalid flags for serverArray.datacenterPolicy.item.weight field of datacenterPolicy array, %s were defined but max position is %s",
			len(r.serverArraydatacenterPolicyitemweight), maxserverArraydatacenterPolicyitemweightPos)
	}
	if maxserverArraydatacenterPolicyitemdatacenterHrefPos > -1 {
		serverArraydatacenterPolicy := make([]*DatacenterPolicy2, maxserverArraydatacenterPolicyitemdatacenterHrefPos+1)
		for i := 0; i < maxserverArraydatacenterPolicyitemdatacenterHrefPos+1; i++ {
			serverArraydatacenterPolicy[i] = DatacenterPolicy2{
			//TBD
			}
		}
		serverArray.datacenterPolicy = serverArraydatacenterPolicy
	}

	maxserverArrayelasticityParamsscheduleitemdayPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.serverArrayelasticityParamsscheduleitemdayPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.elasticityParams.schedule.item.day field of schedule array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.elasticityParams.schedule.item.day field of schedule array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxserverArrayelasticityParamsscheduleitemdayPos {
			maxserverArrayelasticityParamsscheduleitemdayPos = pos
		}
	}
	if len(r.serverArrayelasticityParamsscheduleitemday) != maxserverArrayelasticityParamsscheduleitemdayPos {
		return nil, fmt.Errof("Invalid flags for serverArray.elasticityParams.schedule.item.day field of schedule array, %s were defined but max position is %s",
			len(r.serverArrayelasticityParamsscheduleitemday), maxserverArrayelasticityParamsscheduleitemdayPos)
	}

	maxserverArrayelasticityParamsscheduleitemmaxCountPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.serverArrayelasticityParamsscheduleitemmaxCountPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.elasticityParams.schedule.item.maxCount field of schedule array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.elasticityParams.schedule.item.maxCount field of schedule array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxserverArrayelasticityParamsscheduleitemmaxCountPos {
			maxserverArrayelasticityParamsscheduleitemmaxCountPos = pos
		}
	}
	if len(r.serverArrayelasticityParamsscheduleitemmaxCount) != maxserverArrayelasticityParamsscheduleitemmaxCountPos {
		return nil, fmt.Errof("Invalid flags for serverArray.elasticityParams.schedule.item.maxCount field of schedule array, %s were defined but max position is %s",
			len(r.serverArrayelasticityParamsscheduleitemmaxCount), maxserverArrayelasticityParamsscheduleitemmaxCountPos)
	}

	maxserverArrayelasticityParamsscheduleitemminCountPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.serverArrayelasticityParamsscheduleitemminCountPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.elasticityParams.schedule.item.minCount field of schedule array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.elasticityParams.schedule.item.minCount field of schedule array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxserverArrayelasticityParamsscheduleitemminCountPos {
			maxserverArrayelasticityParamsscheduleitemminCountPos = pos
		}
	}
	if len(r.serverArrayelasticityParamsscheduleitemminCount) != maxserverArrayelasticityParamsscheduleitemminCountPos {
		return nil, fmt.Errof("Invalid flags for serverArray.elasticityParams.schedule.item.minCount field of schedule array, %s were defined but max position is %s",
			len(r.serverArrayelasticityParamsscheduleitemminCount), maxserverArrayelasticityParamsscheduleitemminCountPos)
	}

	maxserverArrayelasticityParamsscheduleitemtimePos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.serverArrayelasticityParamsscheduleitemtimePos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.elasticityParams.schedule.item.time field of schedule array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.elasticityParams.schedule.item.time field of schedule array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxserverArrayelasticityParamsscheduleitemtimePos {
			maxserverArrayelasticityParamsscheduleitemtimePos = pos
		}
	}
	if len(r.serverArrayelasticityParamsscheduleitemtime) != maxserverArrayelasticityParamsscheduleitemtimePos {
		return nil, fmt.Errof("Invalid flags for serverArray.elasticityParams.schedule.item.time field of schedule array, %s were defined but max position is %s",
			len(r.serverArrayelasticityParamsscheduleitemtime), maxserverArrayelasticityParamsscheduleitemtimePos)
	}
	if maxserverArrayelasticityParamsscheduleitemdayPos > -1 {
		serverArrayelasticityParamsschedule := make([]*Schedule2, maxserverArrayelasticityParamsscheduleitemdayPos+1)
		for i := 0; i < maxserverArrayelasticityParamsscheduleitemdayPos+1; i++ {
			serverArrayelasticityParamsschedule[i] = Schedule2{
			//TBD
			}
		}
		serverArray.elasticityParams.schedule = serverArrayelasticityParamsschedule
	}

	return c.UpdateServerArray(r.id, &serverArray)
}

// Register all ServerArray actions
func registerServerArrayCmds(app *kingpin.Application) {
	resCmd := app.Cmd("ServerArray", `A server array represents a logical group of instances and allows to resize(grow/shrink) that group based on certain elasticity parameters.`)

	CloneServerArrayRunner := new(CloneServerArrayServerArrayRunner)
	CloneServerArrayCmd := resCmd.Command("CloneServerArray", `Clones a given server array.`)
	CloneServerArrayRunner.Flag(`id`, ``).Required().StringVar(&CloneServerArrayRunner.id)
	registry[CloneServerArrayCmd.FullCommand()] = CloneServerArrayRunner

	CreateServerArrayRunner := new(CreateServerArrayServerArrayRunner)
	CreateServerArrayCmd := resCmd.Command("CreateServerArray", `Creates a new server array, and configures its corresponding "next" instance with the received parameters.`)
	CreateServerArrayRunner.Flag(`serverArray.arrayType`, `The array type for the Server Array.`).Required().StringVar(&CreateServerArrayRunner.serverArrayArrayType)
	CreateServerArrayRunner.FlagPattern(`serverArray\.item\.(\d+)\.datacenterHref`, `The href of the Datacenter / Zone.`).Required().Capture(&CreateServerArrayRunner.serverArrayItemDatacenterHrefPos).StringsVar(&CreateServerArrayRunner.serverArrayItemDatacenterHref)
	CreateServerArrayRunner.FlagPattern(`serverArray\.item\.(\d+)\.max`, `Max instances (0 for unlimited).`).Required().Capture(&CreateServerArrayRunner.serverArrayItemMaxPos).StringsVar(&CreateServerArrayRunner.serverArrayItemMax)
	CreateServerArrayRunner.FlagPattern(`serverArray\.item\.(\d+)\.weight`, `Instance allocation (should total 100%).`).Required().Capture(&CreateServerArrayRunner.serverArrayItemWeightPos).StringsVar(&CreateServerArrayRunner.serverArrayItemWeight)
	CreateServerArrayRunner.Flag(`serverArray.deploymentHref`, `The href of the deployment for the Server Array.`).StringVar(&CreateServerArrayRunner.serverArrayDeploymentHref)
	CreateServerArrayRunner.Flag(`serverArray.description`, `The description for the Server Array.`).StringVar(&CreateServerArrayRunner.serverArrayDescription)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.alertSpecificParams.decisionThreshold`, `The percentage of servers that must agree in order to trigger an alert before an action is taken.`).StringVar(&CreateServerArrayRunner.serverArrayElasticityParamsAlertSpecificParamsDecisionThreshold)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.alertSpecificParams.votersTagPredicate`, `The Voters Tag that RightScale will use in order to determine when to scale up/down.`).StringVar(&CreateServerArrayRunner.serverArrayElasticityParamsAlertSpecificParamsVotersTagPredicate)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.bounds.maxCount`, `The maximum number of servers that can be operational at the same time in the server array.`).StringVar(&CreateServerArrayRunner.serverArrayElasticityParamsBoundsMaxCount)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.bounds.minCount`, `The minimum number of servers that must be operational at all times in the server array.`).StringVar(&CreateServerArrayRunner.serverArrayElasticityParamsBoundsMinCount)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.pacing.resizeCalmTime`, `The time (in minutes) on how long you want to wait before you repeat another action.`).StringVar(&CreateServerArrayRunner.serverArrayElasticityParamsPacingResizeCalmTime)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.pacing.resizeDownBy`, `The number of servers to scale down by.`).StringVar(&CreateServerArrayRunner.serverArrayElasticityParamsPacingResizeDownBy)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.pacing.resizeUpBy`, `The number of servers to scale up by.`).StringVar(&CreateServerArrayRunner.serverArrayElasticityParamsPacingResizeUpBy)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.collectAuditEntries`, `The audit SQS queue that will store audit entries.`).StringVar(&CreateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsCollectAuditEntries)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.itemAge.algorithm`, `The algorithm that defines how an item's age will be determined, either by the average age or max (oldest) age.`).StringVar(&CreateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.itemAge.maxAge`, `The threshold (in seconds) before a resize action occurs on the server array.`).StringVar(&CreateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.itemAge.regexp`, `The regexp that helps the system determine an item's "age" in the queue. Example: created_at: (\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d UTC)`).StringVar(&CreateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsItemAgeRegexp)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.queueSize.itemsPerInstance`, `Defines the ratio of worker instances per items in the queue. Example: If there are 50 items in the queue and "Items per instance" is set to 10, the server array will resize to 5 worker instances (50/10).  Default = 10`).StringVar(&CreateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance)
	CreateServerArrayRunner.FlagPattern(`serverArray\.elasticityParams\.item\.(\d+)\.day`, `Specifies the day when an alert-based array resizes.`).Required().Capture(&CreateServerArrayRunner.serverArrayElasticityParamsItemDayPos).StringsVar(&CreateServerArrayRunner.serverArrayElasticityParamsItemDay)
	CreateServerArrayRunner.FlagPattern(`serverArray\.elasticityParams\.item\.(\d+)\.maxCount`, `The maximum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`).Required().Capture(&CreateServerArrayRunner.serverArrayElasticityParamsItemMaxCountPos).StringsVar(&CreateServerArrayRunner.serverArrayElasticityParamsItemMaxCount)
	CreateServerArrayRunner.FlagPattern(`serverArray\.elasticityParams\.item\.(\d+)\.minCount`, `The minimum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`).Required().Capture(&CreateServerArrayRunner.serverArrayElasticityParamsItemMinCountPos).StringsVar(&CreateServerArrayRunner.serverArrayElasticityParamsItemMinCount)
	CreateServerArrayRunner.FlagPattern(`serverArray\.elasticityParams\.item\.(\d+)\.time`, `Specifies the time when an alert-based array resizes.`).Required().Capture(&CreateServerArrayRunner.serverArrayElasticityParamsItemTimePos).StringsVar(&CreateServerArrayRunner.serverArrayElasticityParamsItemTime)
	CreateServerArrayRunner.Flag(`serverArray.instance.associatePublicIpAddress`, `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`).StringVar(&CreateServerArrayRunner.serverArrayInstanceAssociatePublicIpAddress)
	CreateServerArrayRunner.Flag(`serverArray.instance.cloudHref`, `The href of the Cloud that the array will be associated with.`).Required().StringVar(&CreateServerArrayRunner.serverArrayInstanceCloudHref)
	CreateServerArrayRunner.Flag(`serverArray.instance.cloudSpecificAttributes.automaticInstanceStoreMapping`, `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`).StringVar(&CreateServerArrayRunner.serverArrayInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping)
	CreateServerArrayRunner.Flag(`serverArray.instance.cloudSpecificAttributes.iamInstanceProfile`, `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`).StringVar(&CreateServerArrayRunner.serverArrayInstanceCloudSpecificAttributesIamInstanceProfile)
	CreateServerArrayRunner.Flag(`serverArray.instance.cloudSpecificAttributes.rootVolumePerformance`, `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`).StringVar(&CreateServerArrayRunner.serverArrayInstanceCloudSpecificAttributesRootVolumePerformance)
	CreateServerArrayRunner.Flag(`serverArray.instance.cloudSpecificAttributes.rootVolumeSize`, `The size for root disk. Not supported in all Clouds.`).StringVar(&CreateServerArrayRunner.serverArrayInstanceCloudSpecificAttributesRootVolumeSize)
	CreateServerArrayRunner.Flag(`serverArray.instance.cloudSpecificAttributes.rootVolumeTypeUid`, `The type of root volume for instance. Only available on clouds supporting root volume type.`).StringVar(&CreateServerArrayRunner.serverArrayInstanceCloudSpecificAttributesRootVolumeTypeUid)
	CreateServerArrayRunner.Flag(`serverArray.instance.datacenterHref`, `The href of the Datacenter / Zone. For multiple Datacenters, use 'datacenter_policy' instead.`).StringVar(&CreateServerArrayRunner.serverArrayInstanceDatacenterHref)
	CreateServerArrayRunner.Flag(`serverArray.instance.imageHref`, `The href of the Image to be used.`).StringVar(&CreateServerArrayRunner.serverArrayInstanceImageHref)
	CreateServerArrayRunner.FlagPattern(`serverArray\.instance\.inputs\.([a-z0-9_]+)`, ``).Capture(&CreateServerArrayRunner.serverArrayInstanceInputsNames).StringVar(&CreateServerArrayRunner.serverArrayInstanceInputsValues)
	CreateServerArrayRunner.Flag(`serverArray.instance.instanceTypeHref`, `The href of the Instance Type.`).StringVar(&CreateServerArrayRunner.serverArrayInstanceInstanceTypeHref)
	CreateServerArrayRunner.Flag(`serverArray.instance.ipForwardingEnabled`, `Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.`).StringVar(&CreateServerArrayRunner.serverArrayInstanceIpForwardingEnabled)
	CreateServerArrayRunner.Flag(`serverArray.instance.kernelImageHref`, `The href of the Kernel Image.`).StringVar(&CreateServerArrayRunner.serverArrayInstanceKernelImageHref)
	CreateServerArrayRunner.Flag(`serverArray.instance.multiCloudImageHref`, `The href of the MultiCloudImage to be used.`).StringVar(&CreateServerArrayRunner.serverArrayInstanceMultiCloudImageHref)
	CreateServerArrayRunner.Flag(`serverArray.instance.placementGroupHref`, `The href of the Placement Group.`).StringVar(&CreateServerArrayRunner.serverArrayInstancePlacementGroupHref)
	CreateServerArrayRunner.Flag(`serverArray.instance.ramdiskImageHref`, `The href of the Ramdisk Image.`).StringVar(&CreateServerArrayRunner.serverArrayInstanceRamdiskImageHref)
	CreateServerArrayRunner.FlagPattern(`serverArray\.instance\.item\.(\d+)`, `The hrefs of the Security Groups.`).Capture(&CreateServerArrayRunner.serverArrayInstanceItemPos).StringsVar(&CreateServerArrayRunner.serverArrayInstanceItem)
	CreateServerArrayRunner.Flag(`serverArray.instance.serverTemplateHref`, `The ServerTemplate that will be used to create the worker instances in the server array.`).Required().StringVar(&CreateServerArrayRunner.serverArrayInstanceServerTemplateHref)
	CreateServerArrayRunner.Flag(`serverArray.instance.sshKeyHref`, `The href of the SSH Key to be used.`).StringVar(&CreateServerArrayRunner.serverArrayInstanceSshKeyHref)
	CreateServerArrayRunner.FlagPattern(`serverArray\.instance\.item\.(\d+)`, `The hrefs of the updated Subnets.`).Capture(&CreateServerArrayRunner.serverArrayInstanceItemPos).StringsVar(&CreateServerArrayRunner.serverArrayInstanceItem)
	CreateServerArrayRunner.Flag(`serverArray.instance.userData`, `User data that RightScale automatically passes to your instance at boot time.`).StringVar(&CreateServerArrayRunner.serverArrayInstanceUserData)
	CreateServerArrayRunner.Flag(`serverArray.name`, `The name for the Server Array.`).Required().StringVar(&CreateServerArrayRunner.serverArrayName)
	CreateServerArrayRunner.Flag(`serverArray.optimized`, `A flag indicating whether Instances of this ServerArray should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`).StringVar(&CreateServerArrayRunner.serverArrayOptimized)
	CreateServerArrayRunner.Flag(`serverArray.state`, `The status of the server array. If active, the server array is enabled for scaling actions.`).Required().StringVar(&CreateServerArrayRunner.serverArrayState)
	registry[CreateServerArrayCmd.FullCommand()] = CreateServerArrayRunner

	CurrentInstancesServerArrayRunner := new(CurrentInstancesServerArrayServerArrayRunner)
	CurrentInstancesServerArrayCmd := resCmd.Command("CurrentInstancesServerArray", `List the running instances belonging to the server array. See Instances#index for details.`)
	CurrentInstancesServerArrayRunner.Flag(`id`, ``).Required().StringVar(&CurrentInstancesServerArrayRunner.id)
	registry[CurrentInstancesServerArrayCmd.FullCommand()] = CurrentInstancesServerArrayRunner

	DestroyServerArrayRunner := new(DestroyServerArrayServerArrayRunner)
	DestroyServerArrayCmd := resCmd.Command("DestroyServerArray", `Deletes a given server array.`)
	DestroyServerArrayRunner.Flag(`id`, ``).Required().StringVar(&DestroyServerArrayRunner.id)
	registry[DestroyServerArrayCmd.FullCommand()] = DestroyServerArrayRunner

	IndexServerArraysRunner := new(IndexServerArraysServerArrayRunner)
	IndexServerArraysCmd := resCmd.Command("IndexServerArrays", `Lists server arrays.`)
	IndexServerArraysRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexServerArraysRunner.filterPos).StringsVar(&IndexServerArraysRunner.filter)
	IndexServerArraysRunner.Flag(`view`, ``).StringVar(&IndexServerArraysRunner.view)
	registry[IndexServerArraysCmd.FullCommand()] = IndexServerArraysRunner

	LaunchServerArrayRunner := new(LaunchServerArrayServerArrayRunner)
	LaunchServerArrayCmd := resCmd.Command("LaunchServerArray", `Launches a new instance in the server array with the configuration defined in the 'next_instance'`)
	LaunchServerArrayRunner.Flag(`id`, ``).Required().StringVar(&LaunchServerArrayRunner.id)
	registry[LaunchServerArrayCmd.FullCommand()] = LaunchServerArrayRunner

	MultiRunExecutableServerArraysRunner := new(MultiRunExecutableServerArraysServerArrayRunner)
	MultiRunExecutableServerArraysCmd := resCmd.Command("MultiRunExecutableServerArrays", `Run an executable on all instances of this array`)
	MultiRunExecutableServerArraysRunner.Flag(`id`, ``).Required().StringVar(&MultiRunExecutableServerArraysRunner.id)
	registry[MultiRunExecutableServerArraysCmd.FullCommand()] = MultiRunExecutableServerArraysRunner

	MultiTerminateServerArraysRunner := new(MultiTerminateServerArraysServerArrayRunner)
	MultiTerminateServerArraysCmd := resCmd.Command("MultiTerminateServerArrays", `Terminate all instances of this array`)
	MultiTerminateServerArraysRunner.Flag(`id`, ``).Required().StringVar(&MultiTerminateServerArraysRunner.id)
	registry[MultiTerminateServerArraysCmd.FullCommand()] = MultiTerminateServerArraysRunner

	ShowServerArrayRunner := new(ShowServerArrayServerArrayRunner)
	ShowServerArrayCmd := resCmd.Command("ShowServerArray", `Shows the information of a single server array.`)
	ShowServerArrayRunner.Flag(`id`, ``).Required().StringVar(&ShowServerArrayRunner.id)
	ShowServerArrayRunner.Flag(`view`, ``).StringVar(&ShowServerArrayRunner.view)
	registry[ShowServerArrayCmd.FullCommand()] = ShowServerArrayRunner

	UpdateServerArrayRunner := new(UpdateServerArrayServerArrayRunner)
	UpdateServerArrayCmd := resCmd.Command("UpdateServerArray", `Updates attributes of a single server array.`)
	UpdateServerArrayRunner.Flag(`id`, ``).Required().StringVar(&UpdateServerArrayRunner.id)
	UpdateServerArrayRunner.Flag(`serverArray.arrayType`, `The updated array type for the Server Array.`).StringVar(&UpdateServerArrayRunner.serverArrayArrayType)
	UpdateServerArrayRunner.FlagPattern(`serverArray\.item\.(\d+)\.datacenterHref`, `The href of the Datacenter / Zone.`).Required().Capture(&UpdateServerArrayRunner.serverArrayItemDatacenterHrefPos).StringsVar(&UpdateServerArrayRunner.serverArrayItemDatacenterHref)
	UpdateServerArrayRunner.FlagPattern(`serverArray\.item\.(\d+)\.max`, `Max instances (0 for unlimited).`).Required().Capture(&UpdateServerArrayRunner.serverArrayItemMaxPos).StringsVar(&UpdateServerArrayRunner.serverArrayItemMax)
	UpdateServerArrayRunner.FlagPattern(`serverArray\.item\.(\d+)\.weight`, `Instance allocation (should total 100%).`).Required().Capture(&UpdateServerArrayRunner.serverArrayItemWeightPos).StringsVar(&UpdateServerArrayRunner.serverArrayItemWeight)
	UpdateServerArrayRunner.Flag(`serverArray.deploymentHref`, `The updated href of the deployment for the Server Array.`).StringVar(&UpdateServerArrayRunner.serverArrayDeploymentHref)
	UpdateServerArrayRunner.Flag(`serverArray.description`, `The updated description for the Server Array.`).StringVar(&UpdateServerArrayRunner.serverArrayDescription)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.alertSpecificParams.decisionThreshold`, `The updated percentage of servers that must agree in order to trigger an alert before an action is taken.`).StringVar(&UpdateServerArrayRunner.serverArrayElasticityParamsAlertSpecificParamsDecisionThreshold)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.alertSpecificParams.votersTagPredicate`, `The updated Voters Tag that RightScale will use in order to determine when to scale up/down.`).StringVar(&UpdateServerArrayRunner.serverArrayElasticityParamsAlertSpecificParamsVotersTagPredicate)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.bounds.maxCount`, `The updated maximum number of servers that can be operational at the same time in the server array.`).StringVar(&UpdateServerArrayRunner.serverArrayElasticityParamsBoundsMaxCount)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.bounds.minCount`, `The updated minimum number of servers that must be operational at all times in the server array.`).StringVar(&UpdateServerArrayRunner.serverArrayElasticityParamsBoundsMinCount)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.pacing.resizeCalmTime`, `The updated time (in minutes) on how long you want to wait before you repeat another action.`).StringVar(&UpdateServerArrayRunner.serverArrayElasticityParamsPacingResizeCalmTime)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.pacing.resizeDownBy`, `The updated number of servers to scale down by.`).StringVar(&UpdateServerArrayRunner.serverArrayElasticityParamsPacingResizeDownBy)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.pacing.resizeUpBy`, `The updated number of servers to scale up by.`).StringVar(&UpdateServerArrayRunner.serverArrayElasticityParamsPacingResizeUpBy)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.collectAuditEntries`, `The updated audit SQS queue that will store audit entries.`).StringVar(&UpdateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsCollectAuditEntries)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.itemAge.algorithm`, `The updated algorithm that defines how an item's age will be determined, either by the average age or max (oldest) age.`).StringVar(&UpdateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.itemAge.maxAge`, `The updated threshold (in seconds) before a resize action occurs on the server array.`).StringVar(&UpdateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.itemAge.regexp`, `The updated regexp that helps the system determine an item's "age" in the queue. Example: created_at: (\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d UTC)`).StringVar(&UpdateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsItemAgeRegexp)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.queueSize.itemsPerInstance`, `Defines the ratio of worker instances per items in the queue. Example: If there are 50 items in the queue and "Items per instance" is set to 10, the server array will resize to 5 worker instances (50/10).  Default = 10`).StringVar(&UpdateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance)
	UpdateServerArrayRunner.FlagPattern(`serverArray\.elasticityParams\.item\.(\d+)\.day`, `The updated day when an alert-based array resizes.`).Required().Capture(&UpdateServerArrayRunner.serverArrayElasticityParamsItemDayPos).StringsVar(&UpdateServerArrayRunner.serverArrayElasticityParamsItemDay)
	UpdateServerArrayRunner.FlagPattern(`serverArray\.elasticityParams\.item\.(\d+)\.maxCount`, `The updated maximum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`).Required().Capture(&UpdateServerArrayRunner.serverArrayElasticityParamsItemMaxCountPos).StringsVar(&UpdateServerArrayRunner.serverArrayElasticityParamsItemMaxCount)
	UpdateServerArrayRunner.FlagPattern(`serverArray\.elasticityParams\.item\.(\d+)\.minCount`, `The updated minimum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`).Required().Capture(&UpdateServerArrayRunner.serverArrayElasticityParamsItemMinCountPos).StringsVar(&UpdateServerArrayRunner.serverArrayElasticityParamsItemMinCount)
	UpdateServerArrayRunner.FlagPattern(`serverArray\.elasticityParams\.item\.(\d+)\.time`, `The updated time when an alert-based array resizes.`).Required().Capture(&UpdateServerArrayRunner.serverArrayElasticityParamsItemTimePos).StringsVar(&UpdateServerArrayRunner.serverArrayElasticityParamsItemTime)
	UpdateServerArrayRunner.Flag(`serverArray.name`, `The updated name for the Server Array.`).StringVar(&UpdateServerArrayRunner.serverArrayName)
	UpdateServerArrayRunner.Flag(`serverArray.optimized`, `A flag indicating whether Instances of this ServerArray should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`).StringVar(&UpdateServerArrayRunner.serverArrayOptimized)
	UpdateServerArrayRunner.Flag(`serverArray.state`, `The updated status of the server array. If active, the server array is enabled for scaling actions.`).StringVar(&UpdateServerArrayRunner.serverArrayState)
	registry[UpdateServerArrayCmd.FullCommand()] = UpdateServerArrayRunner
}

/****** ServerTemplate ******/

type CloneServerTemplateServerTemplateRunner struct {
	id                        string
	serverTemplateDescription string
	serverTemplateName        string
}

func (r *CloneServerTemplateServerTemplateRunner) Run(c *Client) (interface{}, error) {

	/** Handle serverTemplate parameter **/
	var serverTemplate ServerTemplateParam

	// Load JSON if provided
	if len(r.serverTemplateJson) > 0 {
		if err := Json.Unmarshal(r.serverTemplateJson, &serverTemplate); err != nil {
			return nil, fmt.Errorf("Could not load serverTemplate JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.serverTemplatedescription) > 0 {
		serverTemplate.serverTemplate.description = r.serverTemplatedescription
	}

	if len(r.serverTemplatename) > 0 {
		serverTemplate.serverTemplate.name = r.serverTemplatename
	}

	return c.CloneServerTemplate(r.id, &serverTemplate)
}

type CommitServerTemplateServerTemplateRunner struct {
	commitHeadDependencies string
	commitMessage          string
	freezeRepositories     string
	id                     string
}

func (r *CommitServerTemplateServerTemplateRunner) Run(c *Client) (interface{}, error) {
	return c.CommitServerTemplate(r.commitHeadDependencies, r.commitMessage, r.freezeRepositories, r.id)
}

type CreateServerTemplateServerTemplateRunner struct {
	serverTemplateDescription string
	serverTemplateName        string
}

func (r *CreateServerTemplateServerTemplateRunner) Run(c *Client) (interface{}, error) {

	/** Handle serverTemplate parameter **/
	var serverTemplate ServerTemplateParam2

	// Load JSON if provided
	if len(r.serverTemplateJson) > 0 {
		if err := Json.Unmarshal(r.serverTemplateJson, &serverTemplate); err != nil {
			return nil, fmt.Errorf("Could not load serverTemplate JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.serverTemplatedescription) > 0 {
		serverTemplate.serverTemplate.description = r.serverTemplatedescription
	}

	if len(r.serverTemplatename) > 0 {
		serverTemplate.serverTemplate.name = r.serverTemplatename
	}

	return c.CreateServerTemplate(&serverTemplate)
}

type DestroyServerTemplateServerTemplateRunner struct {
	id string
}

func (r *DestroyServerTemplateServerTemplateRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyServerTemplate(r.id)
}

type DetectChangesInHeadServerTemplateServerTemplateRunner struct {
	id string
}

func (r *DetectChangesInHeadServerTemplateServerTemplateRunner) Run(c *Client) (interface{}, error) {
	return c.DetectChangesInHeadServerTemplate(r.id)
}

type IndexServerTemplatesServerTemplateRunner struct {
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexServerTemplatesServerTemplateRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexServerTemplates(filter, r.view)
}

type PublishServerTemplateServerTemplateRunner struct {
	accountGroupHrefs    []string
	accountGroupHrefsPos []string
	allowComments        string
	categories           []string
	categoriesPos        []string
	descriptionsLong     string
	descriptionsNotes    string
	descriptionsShort    string
	emailComments        string
	id                   string
}

func (r *PublishServerTemplateServerTemplateRunner) Run(c *Client) (interface{}, error) {

	/** Handle accountGroupHrefs parameter **/
	var accountGroupHrefs []string

	for i, v := range r.accountGroupHrefs {
		pos, err := strconv.Atoi(r.accountGroupHrefsPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for accountGroupHrefs array", r.accountGroupHrefsPos[i])
		}
		accountGroupHrefs[pos] = v
	}

	/** Handle categories parameter **/
	var categories []string

	for i, v := range r.categories {
		pos, err := strconv.Atoi(r.categoriesPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for categories array", r.categoriesPos[i])
		}
		categories[pos] = v
	}

	/** Handle descriptions parameter **/
	var descriptions Descriptions

	// Load JSON if provided
	if len(r.descriptionsJson) > 0 {
		if err := Json.Unmarshal(r.descriptionsJson, &descriptions); err != nil {
			return nil, fmt.Errorf("Could not load descriptions JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.descriptionslong) > 0 {
		descriptions.descriptions.long = r.descriptionslong
	}

	if len(r.descriptionsnotes) > 0 {
		descriptions.descriptions.notes = r.descriptionsnotes
	}

	if len(r.descriptionsshort) > 0 {
		descriptions.descriptions.short = r.descriptionsshort
	}

	return c.PublishServerTemplate(accountGroupHrefs, r.allowComments, categories, &descriptions, r.emailComments, r.id)
}

type ResolveServerTemplateServerTemplateRunner struct {
	id string
}

func (r *ResolveServerTemplateServerTemplateRunner) Run(c *Client) (interface{}, error) {
	return c.ResolveServerTemplate(r.id)
}

type ShowServerTemplateServerTemplateRunner struct {
	id   string
	view string
}

func (r *ShowServerTemplateServerTemplateRunner) Run(c *Client) (interface{}, error) {
	return c.ShowServerTemplate(r.id, r.view)
}

type SwapRepositoryServerTemplateServerTemplateRunner struct {
	id                   string
	sourceRepositoryHref string
	targetRepositoryHref string
}

func (r *SwapRepositoryServerTemplateServerTemplateRunner) Run(c *Client) (interface{}, error) {
	return c.SwapRepositoryServerTemplate(r.id, r.sourceRepositoryHref, r.targetRepositoryHref)
}

type UpdateServerTemplateServerTemplateRunner struct {
	id                        string
	serverTemplateDescription string
	serverTemplateName        string
}

func (r *UpdateServerTemplateServerTemplateRunner) Run(c *Client) (interface{}, error) {

	/** Handle serverTemplate parameter **/
	var serverTemplate ServerTemplateParam

	// Load JSON if provided
	if len(r.serverTemplateJson) > 0 {
		if err := Json.Unmarshal(r.serverTemplateJson, &serverTemplate); err != nil {
			return nil, fmt.Errorf("Could not load serverTemplate JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.serverTemplatedescription) > 0 {
		serverTemplate.serverTemplate.description = r.serverTemplatedescription
	}

	if len(r.serverTemplatename) > 0 {
		serverTemplate.serverTemplate.name = r.serverTemplatename
	}

	return c.UpdateServerTemplate(r.id, &serverTemplate)
}

// Register all ServerTemplate actions
func registerServerTemplateCmds(app *kingpin.Application) {
	resCmd := app.Cmd("ServerTemplate", `ServerTemplates allow you to pre-configure servers by starting from a base image and adding scripts that run during the boot, operational, and shutdown phases...`)

	CloneServerTemplateRunner := new(CloneServerTemplateServerTemplateRunner)
	CloneServerTemplateCmd := resCmd.Command("CloneServerTemplate", `Clones a given ServerTemplate.`)
	CloneServerTemplateRunner.Flag(`id`, ``).Required().StringVar(&CloneServerTemplateRunner.id)
	CloneServerTemplateRunner.Flag(`serverTemplate.description`, `The description for the cloned ServerTemplate.`).StringVar(&CloneServerTemplateRunner.serverTemplateDescription)
	CloneServerTemplateRunner.Flag(`serverTemplate.name`, `The name for the cloned ServerTemplate.`).Required().StringVar(&CloneServerTemplateRunner.serverTemplateName)
	registry[CloneServerTemplateCmd.FullCommand()] = CloneServerTemplateRunner

	CommitServerTemplateRunner := new(CommitServerTemplateServerTemplateRunner)
	CommitServerTemplateCmd := resCmd.Command("CommitServerTemplate", `Commits a given ServerTemplate. Only HEAD revisions (revision 0) that are owned by the account can be committed.`)
	CommitServerTemplateRunner.Flag(`commitHeadDependencies`, `Commit all HEAD revisions (if any) of the associated MultiCloud Images, RightScripts and Chef repo sequences.`).Required().StringVar(&CommitServerTemplateRunner.commitHeadDependencies)
	CommitServerTemplateRunner.Flag(`commitMessage`, `The message associated with the commit.`).Required().StringVar(&CommitServerTemplateRunner.commitMessage)
	CommitServerTemplateRunner.Flag(`freezeRepositories`, `Freeze the repositories.`).Required().StringVar(&CommitServerTemplateRunner.freezeRepositories)
	CommitServerTemplateRunner.Flag(`id`, ``).Required().StringVar(&CommitServerTemplateRunner.id)
	registry[CommitServerTemplateCmd.FullCommand()] = CommitServerTemplateRunner

	CreateServerTemplateRunner := new(CreateServerTemplateServerTemplateRunner)
	CreateServerTemplateCmd := resCmd.Command("CreateServerTemplate", `Creates a new ServerTemplate with the given parameters.`)
	CreateServerTemplateRunner.Flag(`serverTemplate.description`, `The description of the ServerTemplate to be created.`).StringVar(&CreateServerTemplateRunner.serverTemplateDescription)
	CreateServerTemplateRunner.Flag(`serverTemplate.name`, `The name of the ServerTemplate to be created.`).Required().StringVar(&CreateServerTemplateRunner.serverTemplateName)
	registry[CreateServerTemplateCmd.FullCommand()] = CreateServerTemplateRunner

	DestroyServerTemplateRunner := new(DestroyServerTemplateServerTemplateRunner)
	DestroyServerTemplateCmd := resCmd.Command("DestroyServerTemplate", `Deletes a given ServerTemplate.`)
	DestroyServerTemplateRunner.Flag(`id`, ``).Required().StringVar(&DestroyServerTemplateRunner.id)
	registry[DestroyServerTemplateCmd.FullCommand()] = DestroyServerTemplateRunner

	DetectChangesInHeadServerTemplateRunner := new(DetectChangesInHeadServerTemplateServerTemplateRunner)
	DetectChangesInHeadServerTemplateCmd := resCmd.Command("DetectChangesInHeadServerTemplate", `Identifies RightScripts attached to the resource that differ from their HEAD.`)
	DetectChangesInHeadServerTemplateRunner.Flag(`id`, ``).Required().StringVar(&DetectChangesInHeadServerTemplateRunner.id)
	registry[DetectChangesInHeadServerTemplateCmd.FullCommand()] = DetectChangesInHeadServerTemplateRunner

	IndexServerTemplatesRunner := new(IndexServerTemplatesServerTemplateRunner)
	IndexServerTemplatesCmd := resCmd.Command("IndexServerTemplates", `Lists the ServerTemplates available to this account. HEAD revisions have a revision of 0.`)
	IndexServerTemplatesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexServerTemplatesRunner.filterPos).StringsVar(&IndexServerTemplatesRunner.filter)
	IndexServerTemplatesRunner.Flag(`view`, ``).StringVar(&IndexServerTemplatesRunner.view)
	registry[IndexServerTemplatesCmd.FullCommand()] = IndexServerTemplatesRunner

	PublishServerTemplateRunner := new(PublishServerTemplateServerTemplateRunner)
	PublishServerTemplateCmd := resCmd.Command("PublishServerTemplate", `Publishes a given ServerTemplate and its subordinates.`)
	PublishServerTemplateRunner.FlagPattern(`accountGroupHrefs\.(\d+)`, `List of hrefs of account groups to publish to.`).Required().Capture(&PublishServerTemplateRunner.accountGroupHrefsPos).StringsVar(&PublishServerTemplateRunner.accountGroupHrefs)
	PublishServerTemplateRunner.Flag(`allowComments`, `Allow users to leave comments on this ServerTemplate.`).StringVar(&PublishServerTemplateRunner.allowComments)
	PublishServerTemplateRunner.FlagPattern(`categories\.(\d+)`, `List of Categories.`).Capture(&PublishServerTemplateRunner.categoriesPos).StringsVar(&PublishServerTemplateRunner.categories)
	PublishServerTemplateRunner.Flag(`descriptions.long`, `Long Description.`).Required().StringVar(&PublishServerTemplateRunner.descriptionsLong)
	PublishServerTemplateRunner.Flag(`descriptions.notes`, `New Revision Notes.`).Required().StringVar(&PublishServerTemplateRunner.descriptionsNotes)
	PublishServerTemplateRunner.Flag(`descriptions.short`, `Short Description.`).Required().StringVar(&PublishServerTemplateRunner.descriptionsShort)
	PublishServerTemplateRunner.Flag(`emailComments`, `Email me when a user comments on this ServerTemplate.`).StringVar(&PublishServerTemplateRunner.emailComments)
	PublishServerTemplateRunner.Flag(`id`, ``).Required().StringVar(&PublishServerTemplateRunner.id)
	registry[PublishServerTemplateCmd.FullCommand()] = PublishServerTemplateRunner

	ResolveServerTemplateRunner := new(ResolveServerTemplateServerTemplateRunner)
	ResolveServerTemplateCmd := resCmd.Command("ResolveServerTemplate", `Enumerates all attached cookbooks, missing dependencies and bound executables.`)
	ResolveServerTemplateRunner.Flag(`id`, ``).Required().StringVar(&ResolveServerTemplateRunner.id)
	registry[ResolveServerTemplateCmd.FullCommand()] = ResolveServerTemplateRunner

	ShowServerTemplateRunner := new(ShowServerTemplateServerTemplateRunner)
	ShowServerTemplateCmd := resCmd.Command("ShowServerTemplate", `Show information about a single ServerTemplate. HEAD revisions have a revision of 0.`)
	ShowServerTemplateRunner.Flag(`id`, ``).Required().StringVar(&ShowServerTemplateRunner.id)
	ShowServerTemplateRunner.Flag(`view`, ``).StringVar(&ShowServerTemplateRunner.view)
	registry[ShowServerTemplateCmd.FullCommand()] = ShowServerTemplateRunner

	SwapRepositoryServerTemplateRunner := new(SwapRepositoryServerTemplateServerTemplateRunner)
	SwapRepositoryServerTemplateCmd := resCmd.Command("SwapRepositoryServerTemplate", `In-place replacement of attached cookbooks from a given repository.`)
	SwapRepositoryServerTemplateRunner.Flag(`id`, ``).Required().StringVar(&SwapRepositoryServerTemplateRunner.id)
	SwapRepositoryServerTemplateRunner.Flag(`sourceRepositoryHref`, `The repository whose cookbook attachments are to be replaced.`).Required().StringVar(&SwapRepositoryServerTemplateRunner.sourceRepositoryHref)
	SwapRepositoryServerTemplateRunner.Flag(`targetRepositoryHref`, `The repository whose cookbook attachments are to be utilized.`).Required().StringVar(&SwapRepositoryServerTemplateRunner.targetRepositoryHref)
	registry[SwapRepositoryServerTemplateCmd.FullCommand()] = SwapRepositoryServerTemplateRunner

	UpdateServerTemplateRunner := new(UpdateServerTemplateServerTemplateRunner)
	UpdateServerTemplateCmd := resCmd.Command("UpdateServerTemplate", `Updates attributes of a given ServerTemplate. Only HEAD revisions can be updated (revision 0).`)
	UpdateServerTemplateRunner.Flag(`id`, ``).Required().StringVar(&UpdateServerTemplateRunner.id)
	UpdateServerTemplateRunner.Flag(`serverTemplate.description`, `The updated description for the ServerTemplate.`).StringVar(&UpdateServerTemplateRunner.serverTemplateDescription)
	UpdateServerTemplateRunner.Flag(`serverTemplate.name`, `The updated name for the ServerTemplate.`).StringVar(&UpdateServerTemplateRunner.serverTemplateName)
	registry[UpdateServerTemplateCmd.FullCommand()] = UpdateServerTemplateRunner
}

/****** ServerTemplateMultiCloudImage ******/

type CreateServerTemplateMultiCloudImageServerTemplateMultiCloudImageRunner struct {
	serverTemplateMultiCloudImageMultiCloudImageHref string
	serverTemplateMultiCloudImageServerTemplateHref  string
}

func (r *CreateServerTemplateMultiCloudImageServerTemplateMultiCloudImageRunner) Run(c *Client) (interface{}, error) {

	/** Handle serverTemplateMultiCloudImage parameter **/
	var serverTemplateMultiCloudImage ServerTemplateMultiCloudImageParam

	// Load JSON if provided
	if len(r.serverTemplateMultiCloudImageJson) > 0 {
		if err := Json.Unmarshal(r.serverTemplateMultiCloudImageJson, &serverTemplateMultiCloudImage); err != nil {
			return nil, fmt.Errorf("Could not load serverTemplateMultiCloudImage JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.serverTemplateMultiCloudImagemultiCloudImageHref) > 0 {
		serverTemplateMultiCloudImage.serverTemplateMultiCloudImage.multiCloudImageHref = r.serverTemplateMultiCloudImagemultiCloudImageHref
	}

	if len(r.serverTemplateMultiCloudImageserverTemplateHref) > 0 {
		serverTemplateMultiCloudImage.serverTemplateMultiCloudImage.serverTemplateHref = r.serverTemplateMultiCloudImageserverTemplateHref
	}

	return c.CreateServerTemplateMultiCloudImage(&serverTemplateMultiCloudImage)
}

type DestroyServerTemplateMultiCloudImageServerTemplateMultiCloudImageRunner struct {
	id string
}

func (r *DestroyServerTemplateMultiCloudImageServerTemplateMultiCloudImageRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyServerTemplateMultiCloudImage(r.id)
}

type IndexServerTemplateMultiCloudImagesServerTemplateMultiCloudImageRunner struct {
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexServerTemplateMultiCloudImagesServerTemplateMultiCloudImageRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexServerTemplateMultiCloudImages(filter, r.view)
}

type MakeDefaultServerTemplateMultiCloudImageServerTemplateMultiCloudImageRunner struct {
	id string
}

func (r *MakeDefaultServerTemplateMultiCloudImageServerTemplateMultiCloudImageRunner) Run(c *Client) (interface{}, error) {
	return c.MakeDefaultServerTemplateMultiCloudImage(r.id)
}

type ShowServerTemplateMultiCloudImageServerTemplateMultiCloudImageRunner struct {
	id   string
	view string
}

func (r *ShowServerTemplateMultiCloudImageServerTemplateMultiCloudImageRunner) Run(c *Client) (interface{}, error) {
	return c.ShowServerTemplateMultiCloudImage(r.id, r.view)
}

// Register all ServerTemplateMultiCloudImage actions
func registerServerTemplateMultiCloudImageCmds(app *kingpin.Application) {
	resCmd := app.Cmd("ServerTemplateMultiCloudImage", `This resource represents links between ServerTemplates and MultiCloud Images and enables you to effectively add/delete MultiCloudImages to ServerTemplates and make them the default one...`)

	CreateServerTemplateMultiCloudImageRunner := new(CreateServerTemplateMultiCloudImageServerTemplateMultiCloudImageRunner)
	CreateServerTemplateMultiCloudImageCmd := resCmd.Command("CreateServerTemplateMultiCloudImage", `Creates a new ServerTemplateMultiCloudImage with the given parameters.`)
	CreateServerTemplateMultiCloudImageRunner.Flag(`serverTemplateMultiCloudImage.multiCloudImageHref`, `The href of the MultiCloud Image to be used.`).Required().StringVar(&CreateServerTemplateMultiCloudImageRunner.serverTemplateMultiCloudImageMultiCloudImageHref)
	CreateServerTemplateMultiCloudImageRunner.Flag(`serverTemplateMultiCloudImage.serverTemplateHref`, `The href of the ServerTemplate to be used.`).Required().StringVar(&CreateServerTemplateMultiCloudImageRunner.serverTemplateMultiCloudImageServerTemplateHref)
	registry[CreateServerTemplateMultiCloudImageCmd.FullCommand()] = CreateServerTemplateMultiCloudImageRunner

	DestroyServerTemplateMultiCloudImageRunner := new(DestroyServerTemplateMultiCloudImageServerTemplateMultiCloudImageRunner)
	DestroyServerTemplateMultiCloudImageCmd := resCmd.Command("DestroyServerTemplateMultiCloudImage", `Deletes a given ServerTemplateMultiCloudImage.`)
	DestroyServerTemplateMultiCloudImageRunner.Flag(`id`, ``).Required().StringVar(&DestroyServerTemplateMultiCloudImageRunner.id)
	registry[DestroyServerTemplateMultiCloudImageCmd.FullCommand()] = DestroyServerTemplateMultiCloudImageRunner

	IndexServerTemplateMultiCloudImagesRunner := new(IndexServerTemplateMultiCloudImagesServerTemplateMultiCloudImageRunner)
	IndexServerTemplateMultiCloudImagesCmd := resCmd.Command("IndexServerTemplateMultiCloudImages", `Lists the ServerTemplateMultiCloudImages available to this account.`)
	IndexServerTemplateMultiCloudImagesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexServerTemplateMultiCloudImagesRunner.filterPos).StringsVar(&IndexServerTemplateMultiCloudImagesRunner.filter)
	IndexServerTemplateMultiCloudImagesRunner.Flag(`view`, ``).StringVar(&IndexServerTemplateMultiCloudImagesRunner.view)
	registry[IndexServerTemplateMultiCloudImagesCmd.FullCommand()] = IndexServerTemplateMultiCloudImagesRunner

	MakeDefaultServerTemplateMultiCloudImageRunner := new(MakeDefaultServerTemplateMultiCloudImageServerTemplateMultiCloudImageRunner)
	MakeDefaultServerTemplateMultiCloudImageCmd := resCmd.Command("MakeDefaultServerTemplateMultiCloudImage", `Makes a given ServerTemplateMultiCloudImage the default for the ServerTemplate.`)
	MakeDefaultServerTemplateMultiCloudImageRunner.Flag(`id`, ``).Required().StringVar(&MakeDefaultServerTemplateMultiCloudImageRunner.id)
	registry[MakeDefaultServerTemplateMultiCloudImageCmd.FullCommand()] = MakeDefaultServerTemplateMultiCloudImageRunner

	ShowServerTemplateMultiCloudImageRunner := new(ShowServerTemplateMultiCloudImageServerTemplateMultiCloudImageRunner)
	ShowServerTemplateMultiCloudImageCmd := resCmd.Command("ShowServerTemplateMultiCloudImage", `Show information about a single ServerTemplateMultiCloudImage which represents an association between a ServerTemplate and a MultiCloudImage.`)
	ShowServerTemplateMultiCloudImageRunner.Flag(`id`, ``).Required().StringVar(&ShowServerTemplateMultiCloudImageRunner.id)
	ShowServerTemplateMultiCloudImageRunner.Flag(`view`, ``).StringVar(&ShowServerTemplateMultiCloudImageRunner.view)
	registry[ShowServerTemplateMultiCloudImageCmd.FullCommand()] = ShowServerTemplateMultiCloudImageRunner
}

/****** Session ******/

type AccountsSessionSessionRunner struct {
	email    string
	password string
	view     string
}

func (r *AccountsSessionSessionRunner) Run(c *Client) (interface{}, error) {
	return c.AccountsSession(r.email, r.password, r.view)
}

type CreateSessionSessionRunner struct {
	accountHref string
	email       string
	password    string
}

func (r *CreateSessionSessionRunner) Run(c *Client) (interface{}, error) {
	return c.CreateSession(r.accountHref, r.email, r.password)
}

type CreateInstanceSessionSessionSessionRunner struct {
	accountHref   string
	instanceToken string
}

func (r *CreateInstanceSessionSessionSessionRunner) Run(c *Client) (interface{}, error) {
	return c.CreateInstanceSessionSession(r.accountHref, r.instanceToken)
}

type IndexSessionsSessionRunner struct {
}

func (r *IndexSessionsSessionRunner) Run(c *Client) (interface{}, error) {
	return c.IndexSessions()
}

type IndexInstanceSessionSessionSessionRunner struct {
}

func (r *IndexInstanceSessionSessionSessionRunner) Run(c *Client) (interface{}, error) {
	return c.IndexInstanceSessionSession()
}

// Register all Session actions
func registerSessionCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Session", `The sessions resource is in charge of creating API sessions that are bound to a given account`)

	AccountsSessionRunner := new(AccountsSessionSessionRunner)
	AccountsSessionCmd := resCmd.Command("AccountsSession", `List all the accounts that a user has access to.`)
	AccountsSessionRunner.Flag(`email`, `The email to login with if not using existing session.`).StringVar(&AccountsSessionRunner.email)
	AccountsSessionRunner.Flag(`password`, `The corresponding password.`).StringVar(&AccountsSessionRunner.password)
	AccountsSessionRunner.Flag(`view`, `Extended view shows account permissions and products`).StringVar(&AccountsSessionRunner.view)
	registry[AccountsSessionCmd.FullCommand()] = AccountsSessionRunner

	CreateSessionRunner := new(CreateSessionSessionRunner)
	CreateSessionCmd := resCmd.Command("CreateSession", `Creates API session scoped to a given account`)
	CreateSessionRunner.Flag(`accountHref`, `The account href for which the session needs to be created.`).Required().StringVar(&CreateSessionRunner.accountHref)
	CreateSessionRunner.Flag(`email`, `The email to login with.`).Required().StringVar(&CreateSessionRunner.email)
	CreateSessionRunner.Flag(`password`, `The corresponding password.`).Required().StringVar(&CreateSessionRunner.password)
	registry[CreateSessionCmd.FullCommand()] = CreateSessionRunner

	CreateInstanceSessionSessionRunner := new(CreateInstanceSessionSessionSessionRunner)
	CreateInstanceSessionSessionCmd := resCmd.Command("CreateInstanceSessionSession", `Creates API session scoped to a given account and instance.`)
	CreateInstanceSessionSessionRunner.Flag(`accountHref`, `The account href for which the session needs to be created.`).Required().StringVar(&CreateInstanceSessionSessionRunner.accountHref)
	CreateInstanceSessionSessionRunner.Flag(`instanceToken`, `The instance token to login with.`).Required().StringVar(&CreateInstanceSessionSessionRunner.instanceToken)
	registry[CreateInstanceSessionSessionCmd.FullCommand()] = CreateInstanceSessionSessionRunner

	IndexSessionsRunner := new(IndexSessionsSessionRunner)
	IndexSessionsCmd := resCmd.Command("IndexSessions", `Returns a list of root resources so an authenticated session can use them as a starting point or a way to know what features are available within its privileges...`)
	registry[IndexSessionsCmd.FullCommand()] = IndexSessionsRunner

	IndexInstanceSessionSessionRunner := new(IndexInstanceSessionSessionSessionRunner)
	IndexInstanceSessionSessionCmd := resCmd.Command("IndexInstanceSessionSession", `Shows the full attributes of the instance (that has the token used to log-in).`)
	registry[IndexInstanceSessionSessionCmd.FullCommand()] = IndexInstanceSessionSessionRunner
}

/****** SshKey ******/

type CreateSshKeySshKeyRunner struct {
	cloudId    string
	sshKeyName string
}

func (r *CreateSshKeySshKeyRunner) Run(c *Client) (interface{}, error) {

	/** Handle sshKey parameter **/
	var sshKey SshKeyParam

	// Load JSON if provided
	if len(r.sshKeyJson) > 0 {
		if err := Json.Unmarshal(r.sshKeyJson, &sshKey); err != nil {
			return nil, fmt.Errorf("Could not load sshKey JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.sshKeyname) > 0 {
		sshKey.sshKey.name = r.sshKeyname
	}

	return c.CreateSshKey(r.cloudId, &sshKey)
}

type DestroySshKeySshKeyRunner struct {
	cloudId string
	id      string
}

func (r *DestroySshKeySshKeyRunner) Run(c *Client) (interface{}, error) {
	return c.DestroySshKey(r.cloudId, r.id)
}

type IndexSshKeysSshKeyRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexSshKeysSshKeyRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexSshKeys(r.cloudId, filter, r.view)
}

type ShowSshKeySshKeyRunner struct {
	cloudId string
	id      string
	view    string
}

func (r *ShowSshKeySshKeyRunner) Run(c *Client) (interface{}, error) {
	return c.ShowSshKey(r.cloudId, r.id, r.view)
}

// Register all SshKey actions
func registerSshKeyCmds(app *kingpin.Application) {
	resCmd := app.Cmd("SshKey", `Ssh Keys represent a created SSH Key that exists in the cloud.`)

	CreateSshKeyRunner := new(CreateSshKeySshKeyRunner)
	CreateSshKeyCmd := resCmd.Command("CreateSshKey", `Creates a new ssh key.`)
	CreateSshKeyRunner.Flag(`cloudId`, ``).Required().StringVar(&CreateSshKeyRunner.cloudId)
	CreateSshKeyRunner.Flag(`sshKey.name`, `The name for the SSH key to be created.`).Required().StringVar(&CreateSshKeyRunner.sshKeyName)
	registry[CreateSshKeyCmd.FullCommand()] = CreateSshKeyRunner

	DestroySshKeyRunner := new(DestroySshKeySshKeyRunner)
	DestroySshKeyCmd := resCmd.Command("DestroySshKey", `Deletes a given ssh key.`)
	DestroySshKeyRunner.Flag(`cloudId`, ``).Required().StringVar(&DestroySshKeyRunner.cloudId)
	DestroySshKeyRunner.Flag(`id`, ``).Required().StringVar(&DestroySshKeyRunner.id)
	registry[DestroySshKeyCmd.FullCommand()] = DestroySshKeyRunner

	IndexSshKeysRunner := new(IndexSshKeysSshKeyRunner)
	IndexSshKeysCmd := resCmd.Command("IndexSshKeys", `Lists ssh keys.`)
	IndexSshKeysRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexSshKeysRunner.cloudId)
	IndexSshKeysRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexSshKeysRunner.filterPos).StringsVar(&IndexSshKeysRunner.filter)
	IndexSshKeysRunner.Flag(`view`, ``).StringVar(&IndexSshKeysRunner.view)
	registry[IndexSshKeysCmd.FullCommand()] = IndexSshKeysRunner

	ShowSshKeyRunner := new(ShowSshKeySshKeyRunner)
	ShowSshKeyCmd := resCmd.Command("ShowSshKey", `Displays information about a single ssh key.`)
	ShowSshKeyRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowSshKeyRunner.cloudId)
	ShowSshKeyRunner.Flag(`id`, ``).Required().StringVar(&ShowSshKeyRunner.id)
	ShowSshKeyRunner.Flag(`view`, ``).StringVar(&ShowSshKeyRunner.view)
	registry[ShowSshKeyCmd.FullCommand()] = ShowSshKeyRunner
}

/****** Subnet ******/

type CreateSubnetSubnetRunner struct {
	cloudId              string
	instanceId           string
	subnetCidrBlock      string
	subnetDatacenterHref string
	subnetDescription    string
	subnetName           string
	subnetNetworkHref    string
}

func (r *CreateSubnetSubnetRunner) Run(c *Client) (interface{}, error) {

	/** Handle subnet parameter **/
	var subnet SubnetParam

	// Load JSON if provided
	if len(r.subnetJson) > 0 {
		if err := Json.Unmarshal(r.subnetJson, &subnet); err != nil {
			return nil, fmt.Errorf("Could not load subnet JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.subnetcidrBlock) > 0 {
		subnet.subnet.cidrBlock = r.subnetcidrBlock
	}

	if len(r.subnetdatacenterHref) > 0 {
		subnet.subnet.datacenterHref = r.subnetdatacenterHref
	}

	if len(r.subnetdescription) > 0 {
		subnet.subnet.description = r.subnetdescription
	}

	if len(r.subnetname) > 0 {
		subnet.subnet.name = r.subnetname
	}

	if len(r.subnetnetworkHref) > 0 {
		subnet.subnet.networkHref = r.subnetnetworkHref
	}

	return c.CreateSubnet(r.cloudId, r.instanceId, &subnet)
}

type DestroySubnetSubnetRunner struct {
	cloudId    string
	id         string
	instanceId string
}

func (r *DestroySubnetSubnetRunner) Run(c *Client) (interface{}, error) {
	return c.DestroySubnet(r.cloudId, r.id, r.instanceId)
}

type IndexSubnetsSubnetRunner struct {
	cloudId    string
	filter     []string
	filterPos  []string
	instanceId string
}

func (r *IndexSubnetsSubnetRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexSubnets(r.cloudId, filter, r.instanceId)
}

type ShowSubnetSubnetRunner struct {
	cloudId    string
	id         string
	instanceId string
}

func (r *ShowSubnetSubnetRunner) Run(c *Client) (interface{}, error) {
	return c.ShowSubnet(r.cloudId, r.id, r.instanceId)
}

type UpdateSubnetSubnetRunner struct {
	cloudId              string
	id                   string
	instanceId           string
	subnetDescription    string
	subnetName           string
	subnetRouteTableHref string
}

func (r *UpdateSubnetSubnetRunner) Run(c *Client) (interface{}, error) {

	/** Handle subnet parameter **/
	var subnet SubnetParam2

	// Load JSON if provided
	if len(r.subnetJson) > 0 {
		if err := Json.Unmarshal(r.subnetJson, &subnet); err != nil {
			return nil, fmt.Errorf("Could not load subnet JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.subnetdescription) > 0 {
		subnet.subnet.description = r.subnetdescription
	}

	if len(r.subnetname) > 0 {
		subnet.subnet.name = r.subnetname
	}

	if len(r.subnetrouteTableHref) > 0 {
		subnet.subnet.routeTableHref = r.subnetrouteTableHref
	}

	return c.UpdateSubnet(r.cloudId, r.id, r.instanceId, &subnet)
}

// Register all Subnet actions
func registerSubnetCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Subnet", `A Subnet is a logical grouping of network devices`)

	CreateSubnetRunner := new(CreateSubnetSubnetRunner)
	CreateSubnetCmd := resCmd.Command("CreateSubnet", `Creates a new subnet.`)
	CreateSubnetRunner.Flag(`cloudId`, ``).Required().StringVar(&CreateSubnetRunner.cloudId)
	CreateSubnetRunner.Flag(`instanceId`, ``).Required().StringVar(&CreateSubnetRunner.instanceId)
	CreateSubnetRunner.Flag(`subnet.cidrBlock`, `The range of IP addresses for the Subnet.`).Required().StringVar(&CreateSubnetRunner.subnetCidrBlock)
	CreateSubnetRunner.Flag(`subnet.datacenterHref`, `The associated Datacenter.`).StringVar(&CreateSubnetRunner.subnetDatacenterHref)
	CreateSubnetRunner.Flag(`subnet.description`, `The description for the Subnet.`).StringVar(&CreateSubnetRunner.subnetDescription)
	CreateSubnetRunner.Flag(`subnet.name`, `The name for the Subnet.`).StringVar(&CreateSubnetRunner.subnetName)
	CreateSubnetRunner.Flag(`subnet.networkHref`, `The associated Network.`).Required().StringVar(&CreateSubnetRunner.subnetNetworkHref)
	registry[CreateSubnetCmd.FullCommand()] = CreateSubnetRunner

	DestroySubnetRunner := new(DestroySubnetSubnetRunner)
	DestroySubnetCmd := resCmd.Command("DestroySubnet", `Deletes the given subnet(s).`)
	DestroySubnetRunner.Flag(`cloudId`, ``).Required().StringVar(&DestroySubnetRunner.cloudId)
	DestroySubnetRunner.Flag(`id`, ``).Required().StringVar(&DestroySubnetRunner.id)
	DestroySubnetRunner.Flag(`instanceId`, ``).Required().StringVar(&DestroySubnetRunner.instanceId)
	registry[DestroySubnetCmd.FullCommand()] = DestroySubnetRunner

	IndexSubnetsRunner := new(IndexSubnetsSubnetRunner)
	IndexSubnetsCmd := resCmd.Command("IndexSubnets", `Lists subnets of a given cloud.`)
	IndexSubnetsRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexSubnetsRunner.cloudId)
	IndexSubnetsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexSubnetsRunner.filterPos).StringsVar(&IndexSubnetsRunner.filter)
	IndexSubnetsRunner.Flag(`instanceId`, ``).Required().StringVar(&IndexSubnetsRunner.instanceId)
	registry[IndexSubnetsCmd.FullCommand()] = IndexSubnetsRunner

	ShowSubnetRunner := new(ShowSubnetSubnetRunner)
	ShowSubnetCmd := resCmd.Command("ShowSubnet", `Shows attributes of a single subnet.`)
	ShowSubnetRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowSubnetRunner.cloudId)
	ShowSubnetRunner.Flag(`id`, ``).Required().StringVar(&ShowSubnetRunner.id)
	ShowSubnetRunner.Flag(`instanceId`, ``).Required().StringVar(&ShowSubnetRunner.instanceId)
	registry[ShowSubnetCmd.FullCommand()] = ShowSubnetRunner

	UpdateSubnetRunner := new(UpdateSubnetSubnetRunner)
	UpdateSubnetCmd := resCmd.Command("UpdateSubnet", `Updates name and description for the current subnet.`)
	UpdateSubnetRunner.Flag(`cloudId`, ``).Required().StringVar(&UpdateSubnetRunner.cloudId)
	UpdateSubnetRunner.Flag(`id`, ``).Required().StringVar(&UpdateSubnetRunner.id)
	UpdateSubnetRunner.Flag(`instanceId`, ``).Required().StringVar(&UpdateSubnetRunner.instanceId)
	UpdateSubnetRunner.Flag(`subnet.description`, `The updated description for the Subnet.`).StringVar(&UpdateSubnetRunner.subnetDescription)
	UpdateSubnetRunner.Flag(`subnet.name`, `The updated name for the Subnet.`).StringVar(&UpdateSubnetRunner.subnetName)
	UpdateSubnetRunner.Flag(`subnet.routeTableHref`, `The RouteTable to associate/dissociate with this Subnet. Pass this parameter with an empty string to reset this Subnet to use the default RouteTable.`).StringVar(&UpdateSubnetRunner.subnetRouteTableHref)
	registry[UpdateSubnetCmd.FullCommand()] = UpdateSubnetRunner
}

/****** Tag ******/

type ByResourceTagTagRunner struct {
	resourceHrefs    []string
	resourceHrefsPos []string
}

func (r *ByResourceTagTagRunner) Run(c *Client) (interface{}, error) {

	/** Handle resourceHrefs parameter **/
	var resourceHrefs []string

	for i, v := range r.resourceHrefs {
		pos, err := strconv.Atoi(r.resourceHrefsPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for resourceHrefs array", r.resourceHrefsPos[i])
		}
		resourceHrefs[pos] = v
	}

	return c.ByResourceTag(resourceHrefs)
}

type ByTagTagTagRunner struct {
	includeTagsWithPrefix string
	matchAll              string
	resourceType          string
	tags                  []string
	tagsPos               []string
	withDeleted           string
}

func (r *ByTagTagTagRunner) Run(c *Client) (interface{}, error) {

	/** Handle tags parameter **/
	var tags []string

	for i, v := range r.tags {
		pos, err := strconv.Atoi(r.tagsPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for tags array", r.tagsPos[i])
		}
		tags[pos] = v
	}

	return c.ByTagTag(r.includeTagsWithPrefix, r.matchAll, r.resourceType, tags, r.withDeleted)
}

type MultiAddTagsTagRunner struct {
	resourceHrefs    []string
	resourceHrefsPos []string
	tags             []string
	tagsPos          []string
}

func (r *MultiAddTagsTagRunner) Run(c *Client) (interface{}, error) {

	/** Handle resourceHrefs parameter **/
	var resourceHrefs []string

	for i, v := range r.resourceHrefs {
		pos, err := strconv.Atoi(r.resourceHrefsPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for resourceHrefs array", r.resourceHrefsPos[i])
		}
		resourceHrefs[pos] = v
	}

	/** Handle tags parameter **/
	var tags []string

	for i, v := range r.tags {
		pos, err := strconv.Atoi(r.tagsPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for tags array", r.tagsPos[i])
		}
		tags[pos] = v
	}

	return c.MultiAddTags(resourceHrefs, tags)
}

type MultiDeleteTagsTagRunner struct {
	resourceHrefs    []string
	resourceHrefsPos []string
	tags             []string
	tagsPos          []string
}

func (r *MultiDeleteTagsTagRunner) Run(c *Client) (interface{}, error) {

	/** Handle resourceHrefs parameter **/
	var resourceHrefs []string

	for i, v := range r.resourceHrefs {
		pos, err := strconv.Atoi(r.resourceHrefsPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for resourceHrefs array", r.resourceHrefsPos[i])
		}
		resourceHrefs[pos] = v
	}

	/** Handle tags parameter **/
	var tags []string

	for i, v := range r.tags {
		pos, err := strconv.Atoi(r.tagsPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for tags array", r.tagsPos[i])
		}
		tags[pos] = v
	}

	return c.MultiDeleteTags(resourceHrefs, tags)
}

// Register all Tag actions
func registerTagCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Tag", `A tag or machine tag is a useful way of attaching useful metadata to an object/resource.`)

	ByResourceTagRunner := new(ByResourceTagTagRunner)
	ByResourceTagCmd := resCmd.Command("ByResourceTag", `Get tags for a list of resource hrefs.`)
	ByResourceTagRunner.FlagPattern(`resourceHrefs\.(\d+)`, `Hrefs of the resources for which tags are to be returned.`).Required().Capture(&ByResourceTagRunner.resourceHrefsPos).StringsVar(&ByResourceTagRunner.resourceHrefs)
	registry[ByResourceTagCmd.FullCommand()] = ByResourceTagRunner

	ByTagTagRunner := new(ByTagTagTagRunner)
	ByTagTagCmd := resCmd.Command("ByTagTag", `Search for resources having a list of tags in a specific resource_type.`)
	ByTagTagRunner.Flag(`includeTagsWithPrefix`, `If included, all tags matching this prefix will be returned. If not included, no tags will be returned.`).StringVar(&ByTagTagRunner.includeTagsWithPrefix)
	ByTagTagRunner.Flag(`matchAll`, `If set to 'true', resources having all the tags specified in the 'tags' parameter are returned. Otherwise, resources having any of the tags are returned.`).StringVar(&ByTagTagRunner.matchAll)
	ByTagTagRunner.Flag(`resourceType`, `Search among a single resource type.`).Required().StringVar(&ByTagTagRunner.resourceType)
	ByTagTagRunner.FlagPattern(`tags\.(\d+)`, `The tags which must be present on the resource.`).Required().Capture(&ByTagTagRunner.tagsPos).StringsVar(&ByTagTagRunner.tags)
	ByTagTagRunner.Flag(`withDeleted`, `If set to 'true', tags for deleted resources will also be returned. Default value is 'false'.`).StringVar(&ByTagTagRunner.withDeleted)
	registry[ByTagTagCmd.FullCommand()] = ByTagTagRunner

	MultiAddTagsRunner := new(MultiAddTagsTagRunner)
	MultiAddTagsCmd := resCmd.Command("MultiAddTags", `Add a list of tags to a list of hrefs. The tags must be either plain_tags or machine_tags.`)
	MultiAddTagsRunner.FlagPattern(`resourceHrefs\.(\d+)`, `Hrefs of the resources for which the tags are to be added.`).Required().Capture(&MultiAddTagsRunner.resourceHrefsPos).StringsVar(&MultiAddTagsRunner.resourceHrefs)
	MultiAddTagsRunner.FlagPattern(`tags\.(\d+)`, `Tags to be added.`).Required().Capture(&MultiAddTagsRunner.tagsPos).StringsVar(&MultiAddTagsRunner.tags)
	registry[MultiAddTagsCmd.FullCommand()] = MultiAddTagsRunner

	MultiDeleteTagsRunner := new(MultiDeleteTagsTagRunner)
	MultiDeleteTagsCmd := resCmd.Command("MultiDeleteTags", `Delete a list of tags on a list of hrefs. The tags must be either plain_tags or machine_tags.`)
	MultiDeleteTagsRunner.FlagPattern(`resourceHrefs\.(\d+)`, `Hrefs of the resources for which tags are to be deleted.`).Required().Capture(&MultiDeleteTagsRunner.resourceHrefsPos).StringsVar(&MultiDeleteTagsRunner.resourceHrefs)
	MultiDeleteTagsRunner.FlagPattern(`tags\.(\d+)`, `Tags to be deleted.`).Required().Capture(&MultiDeleteTagsRunner.tagsPos).StringsVar(&MultiDeleteTagsRunner.tags)
	registry[MultiDeleteTagsCmd.FullCommand()] = MultiDeleteTagsRunner
}

/****** Task ******/

type ShowTaskTaskRunner struct {
	cloudId    string
	id         string
	instanceId string
	view       string
}

func (r *ShowTaskTaskRunner) Run(c *Client) (interface{}, error) {
	return c.ShowTask(r.cloudId, r.id, r.instanceId, r.view)
}

// Register all Task actions
func registerTaskCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Task", `Tasks represent processes that happen (or have happened) asynchronously within the context of an Instance.`)

	ShowTaskRunner := new(ShowTaskTaskRunner)
	ShowTaskCmd := resCmd.Command("ShowTask", `Displays information of a given task within the context of an instance.`)
	ShowTaskRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowTaskRunner.cloudId)
	ShowTaskRunner.Flag(`id`, ``).Required().StringVar(&ShowTaskRunner.id)
	ShowTaskRunner.Flag(`instanceId`, ``).Required().StringVar(&ShowTaskRunner.instanceId)
	ShowTaskRunner.Flag(`view`, ``).StringVar(&ShowTaskRunner.view)
	registry[ShowTaskCmd.FullCommand()] = ShowTaskRunner
}

/****** User ******/

type CreateUserUserRunner struct {
	userCompany              string
	userEmail                string
	userFirstName            string
	userIdentityProviderHref string
	userLastName             string
	userPassword             string
	userPhone                string
	userPrincipalUid         string
	userTimezoneName         string
}

func (r *CreateUserUserRunner) Run(c *Client) (interface{}, error) {

	/** Handle user parameter **/
	var user UserParam

	// Load JSON if provided
	if len(r.userJson) > 0 {
		if err := Json.Unmarshal(r.userJson, &user); err != nil {
			return nil, fmt.Errorf("Could not load user JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.usercompany) > 0 {
		user.user.company = r.usercompany
	}

	if len(r.useremail) > 0 {
		user.user.email = r.useremail
	}

	if len(r.userfirstName) > 0 {
		user.user.firstName = r.userfirstName
	}

	if len(r.useridentityProviderHref) > 0 {
		user.user.identityProviderHref = r.useridentityProviderHref
	}

	if len(r.userlastName) > 0 {
		user.user.lastName = r.userlastName
	}

	if len(r.userpassword) > 0 {
		user.user.password = r.userpassword
	}

	if len(r.userphone) > 0 {
		user.user.phone = r.userphone
	}

	if len(r.userprincipalUid) > 0 {
		user.user.principalUid = r.userprincipalUid
	}

	if len(r.usertimezoneName) > 0 {
		user.user.timezoneName = r.usertimezoneName
	}

	return c.CreateUser(&user)
}

type IndexUsersUserRunner struct {
	filter    []string
	filterPos []string
}

func (r *IndexUsersUserRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexUsers(filter)
}

type ShowUserUserRunner struct {
	id string
}

func (r *ShowUserUserRunner) Run(c *Client) (interface{}, error) {
	return c.ShowUser(r.id)
}

type UpdateUserUserRunner struct {
	id                       string
	userCompany              string
	userCurrentEmail         string
	userCurrentPassword      string
	userFirstName            string
	userIdentityProviderHref string
	userLastName             string
	userNewEmail             string
	userNewPassword          string
	userPhone                string
	userPrincipalUid         string
	userTimezoneName         string
}

func (r *UpdateUserUserRunner) Run(c *Client) (interface{}, error) {

	/** Handle user parameter **/
	var user UserParam2

	// Load JSON if provided
	if len(r.userJson) > 0 {
		if err := Json.Unmarshal(r.userJson, &user); err != nil {
			return nil, fmt.Errorf("Could not load user JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.usercompany) > 0 {
		user.user.company = r.usercompany
	}

	if len(r.usercurrentEmail) > 0 {
		user.user.currentEmail = r.usercurrentEmail
	}

	if len(r.usercurrentPassword) > 0 {
		user.user.currentPassword = r.usercurrentPassword
	}

	if len(r.userfirstName) > 0 {
		user.user.firstName = r.userfirstName
	}

	if len(r.useridentityProviderHref) > 0 {
		user.user.identityProviderHref = r.useridentityProviderHref
	}

	if len(r.userlastName) > 0 {
		user.user.lastName = r.userlastName
	}

	if len(r.usernewEmail) > 0 {
		user.user.newEmail = r.usernewEmail
	}

	if len(r.usernewPassword) > 0 {
		user.user.newPassword = r.usernewPassword
	}

	if len(r.userphone) > 0 {
		user.user.phone = r.userphone
	}

	if len(r.userprincipalUid) > 0 {
		user.user.principalUid = r.userprincipalUid
	}

	if len(r.usertimezoneName) > 0 {
		user.user.timezoneName = r.usertimezoneName
	}

	return c.UpdateUser(r.id, &user)
}

// Register all User actions
func registerUserCmds(app *kingpin.Application) {
	resCmd := app.Cmd("User", `A User represents an individual RightScale user`)

	CreateUserRunner := new(CreateUserUserRunner)
	CreateUserCmd := resCmd.Command("CreateUser", `Create a user. If a user already exists with the same email, that user will be returned.`)
	CreateUserRunner.Flag(`user.company`, ``).Required().StringVar(&CreateUserRunner.userCompany)
	CreateUserRunner.Flag(`user.email`, ``).Required().StringVar(&CreateUserRunner.userEmail)
	CreateUserRunner.Flag(`user.firstName`, ``).Required().StringVar(&CreateUserRunner.userFirstName)
	CreateUserRunner.Flag(`user.identityProviderHref`, `The RightScale API href of the Identity Provider through which this user will login to RightScale. Required to create an SSO-authenticated user.`).StringVar(&CreateUserRunner.userIdentityProviderHref)
	CreateUserRunner.Flag(`user.lastName`, ``).Required().StringVar(&CreateUserRunner.userLastName)
	CreateUserRunner.Flag(`user.password`, `The password of this user. Required to create a password-authenticated user.`).StringVar(&CreateUserRunner.userPassword)
	CreateUserRunner.Flag(`user.phone`, ``).Required().StringVar(&CreateUserRunner.userPhone)
	CreateUserRunner.Flag(`user.principalUid`, `The principal identifier (SAML NameID or OpenID identity URL) of this user. Required to create an SSO-authenticated user.`).StringVar(&CreateUserRunner.userPrincipalUid)
	CreateUserRunner.Flag(`user.timezoneName`, `This can be in the form of country/region or timezone name. For example 'America/Los_Angeles' or 'GB' or 'UTC'. A complete list of acceptable values is available in the Settings > User Settings > Preferences page.`).StringVar(&CreateUserRunner.userTimezoneName)
	registry[CreateUserCmd.FullCommand()] = CreateUserRunner

	IndexUsersRunner := new(IndexUsersUserRunner)
	IndexUsersCmd := resCmd.Command("IndexUsers", `List the users available to the account the user is logged in to`)
	IndexUsersRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexUsersRunner.filterPos).StringsVar(&IndexUsersRunner.filter)
	registry[IndexUsersCmd.FullCommand()] = IndexUsersRunner

	ShowUserRunner := new(ShowUserUserRunner)
	ShowUserCmd := resCmd.Command("ShowUser", `Show information about a single user.`)
	ShowUserRunner.Flag(`id`, ``).Required().StringVar(&ShowUserRunner.id)
	registry[ShowUserCmd.FullCommand()] = ShowUserRunner

	UpdateUserRunner := new(UpdateUserUserRunner)
	UpdateUserCmd := resCmd.Command("UpdateUser", `Update a user's contact information, change her password, or update SSO her settings`)
	UpdateUserRunner.Flag(`id`, ``).Required().StringVar(&UpdateUserRunner.id)
	UpdateUserRunner.Flag(`user.company`, ``).StringVar(&UpdateUserRunner.userCompany)
	UpdateUserRunner.Flag(`user.currentEmail`, `The existing email of this user.`).Required().StringVar(&UpdateUserRunner.userCurrentEmail)
	UpdateUserRunner.Flag(`user.currentPassword`, `The current password for the user.`).StringVar(&UpdateUserRunner.userCurrentPassword)
	UpdateUserRunner.Flag(`user.firstName`, ``).StringVar(&UpdateUserRunner.userFirstName)
	UpdateUserRunner.Flag(`user.identityProviderHref`, `The updated RightScale API href of the associated Identity Provider.`).StringVar(&UpdateUserRunner.userIdentityProviderHref)
	UpdateUserRunner.Flag(`user.lastName`, ``).StringVar(&UpdateUserRunner.userLastName)
	UpdateUserRunner.Flag(`user.newEmail`, `The updated email of this user.`).StringVar(&UpdateUserRunner.userNewEmail)
	UpdateUserRunner.Flag(`user.newPassword`, `The new password for this user.`).StringVar(&UpdateUserRunner.userNewPassword)
	UpdateUserRunner.Flag(`user.phone`, ``).StringVar(&UpdateUserRunner.userPhone)
	UpdateUserRunner.Flag(`user.principalUid`, `The updated principal identifier (SAML NameID or OpenID identity URL) of this user.`).StringVar(&UpdateUserRunner.userPrincipalUid)
	UpdateUserRunner.Flag(`user.timezoneName`, `This can be in the form of country/region or timezone name. For example 'America/Los_Angeles' or 'GB' or 'UTC'. A complete list of acceptable values is available in the Settings > User Settings > Preferences page.`).StringVar(&UpdateUserRunner.userTimezoneName)
	registry[UpdateUserCmd.FullCommand()] = UpdateUserRunner
}

/****** UserData ******/

type ShowUserDataUserDataRunner struct {
}

func (r *ShowUserDataUserDataRunner) Run(c *Client) (interface{}, error) {
	return c.ShowUserData()
}

// Register all UserData actions
func registerUserDataCmds(app *kingpin.Application) {
	resCmd := app.Cmd("UserData", ``)

	ShowUserDataRunner := new(ShowUserDataUserDataRunner)
	ShowUserDataCmd := resCmd.Command("ShowUserData", `<no description>`)
	registry[ShowUserDataCmd.FullCommand()] = ShowUserDataRunner
}

/****** Volume ******/

type CreateVolumeVolumeRunner struct {
	cloudId                        string
	volumeDatacenterHref           string
	volumeDeploymentHref           string
	volumeDescription              string
	volumeEncrypted                string
	volumeIops                     string
	volumeName                     string
	volumeParentVolumeSnapshotHref string
	volumePlacementGroupHref       string
	volumeSize                     string
	volumeVolumeTypeHref           string
}

func (r *CreateVolumeVolumeRunner) Run(c *Client) (interface{}, error) {

	/** Handle volume parameter **/
	var volume VolumeParam

	// Load JSON if provided
	if len(r.volumeJson) > 0 {
		if err := Json.Unmarshal(r.volumeJson, &volume); err != nil {
			return nil, fmt.Errorf("Could not load volume JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.volumedatacenterHref) > 0 {
		volume.volume.datacenterHref = r.volumedatacenterHref
	}

	if len(r.volumedeploymentHref) > 0 {
		volume.volume.deploymentHref = r.volumedeploymentHref
	}

	if len(r.volumedescription) > 0 {
		volume.volume.description = r.volumedescription
	}

	if len(r.volumeencrypted) > 0 {
		volume.volume.encrypted = r.volumeencrypted
	}

	if len(r.volumeiops) > 0 {
		volume.volume.iops = r.volumeiops
	}

	if len(r.volumename) > 0 {
		volume.volume.name = r.volumename
	}

	if len(r.volumeparentVolumeSnapshotHref) > 0 {
		volume.volume.parentVolumeSnapshotHref = r.volumeparentVolumeSnapshotHref
	}

	if len(r.volumeplacementGroupHref) > 0 {
		volume.volume.placementGroupHref = r.volumeplacementGroupHref
	}

	if len(r.volumesize) > 0 {
		volume.volume.size = r.volumesize
	}

	if len(r.volumevolumeTypeHref) > 0 {
		volume.volume.volumeTypeHref = r.volumevolumeTypeHref
	}

	return c.CreateVolume(r.cloudId, &volume)
}

type DestroyVolumeVolumeRunner struct {
	cloudId string
	id      string
}

func (r *DestroyVolumeVolumeRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyVolume(r.cloudId, r.id)
}

type IndexVolumesVolumeRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexVolumesVolumeRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexVolumes(r.cloudId, filter, r.view)
}

type ShowVolumeVolumeRunner struct {
	cloudId string
	id      string
	view    string
}

func (r *ShowVolumeVolumeRunner) Run(c *Client) (interface{}, error) {
	return c.ShowVolume(r.cloudId, r.id, r.view)
}

// Register all Volume actions
func registerVolumeCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Volume", `A Volume provides a highly reliable, efficient and persistent storage solution that can be mounted to a cloud instance (in the same datacenter / zone).`)

	CreateVolumeRunner := new(CreateVolumeVolumeRunner)
	CreateVolumeCmd := resCmd.Command("CreateVolume", `Creates a new volume.`)
	CreateVolumeRunner.Flag(`cloudId`, ``).Required().StringVar(&CreateVolumeRunner.cloudId)
	CreateVolumeRunner.Flag(`volume.datacenterHref`, `The href of the Datacenter / Zone that the Volume will be in. This parameter is required for non-OpenStack clouds.`).StringVar(&CreateVolumeRunner.volumeDatacenterHref)
	CreateVolumeRunner.Flag(`volume.deploymentHref`, `The href of the Deployment that owns this Volume.`).StringVar(&CreateVolumeRunner.volumeDeploymentHref)
	CreateVolumeRunner.Flag(`volume.description`, `The description of the Volume to be created.`).StringVar(&CreateVolumeRunner.volumeDescription)
	CreateVolumeRunner.Flag(`volume.encrypted`, `A flag indicating whether Volume should be encrypted. Only available on clouds supporting volume encryption.`).StringVar(&CreateVolumeRunner.volumeEncrypted)
	CreateVolumeRunner.Flag(`volume.iops`, `The number of IOPS (I/O Operations Per Second) this Volume should support. Only available on clouds supporting performance provisioning.`).StringVar(&CreateVolumeRunner.volumeIops)
	CreateVolumeRunner.Flag(`volume.name`, `The name for the Volume to be created.`).Required().StringVar(&CreateVolumeRunner.volumeName)
	CreateVolumeRunner.Flag(`volume.parentVolumeSnapshotHref`, `The href of the snapshot from which the volume will be created.`).StringVar(&CreateVolumeRunner.volumeParentVolumeSnapshotHref)
	CreateVolumeRunner.Flag(`volume.placementGroupHref`, `The href of the Placement Group.`).StringVar(&CreateVolumeRunner.volumePlacementGroupHref)
	CreateVolumeRunner.Flag(`volume.size`, `The size of a Volume to be created in gigabytes (GB). Some Volume Types have predefined sizes and do not allow selecting a custom size on Volume creation.`).StringVar(&CreateVolumeRunner.volumeSize)
	CreateVolumeRunner.Flag(`volume.volumeTypeHref`, `The href of the volume type. A Name, Resource UID and optional Size is associated with a Volume Type.`).StringVar(&CreateVolumeRunner.volumeVolumeTypeHref)
	registry[CreateVolumeCmd.FullCommand()] = CreateVolumeRunner

	DestroyVolumeRunner := new(DestroyVolumeVolumeRunner)
	DestroyVolumeCmd := resCmd.Command("DestroyVolume", `Deletes a given volume.`)
	DestroyVolumeRunner.Flag(`cloudId`, ``).Required().StringVar(&DestroyVolumeRunner.cloudId)
	DestroyVolumeRunner.Flag(`id`, ``).Required().StringVar(&DestroyVolumeRunner.id)
	registry[DestroyVolumeCmd.FullCommand()] = DestroyVolumeRunner

	IndexVolumesRunner := new(IndexVolumesVolumeRunner)
	IndexVolumesCmd := resCmd.Command("IndexVolumes", `Lists volumes.`)
	IndexVolumesRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexVolumesRunner.cloudId)
	IndexVolumesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexVolumesRunner.filterPos).StringsVar(&IndexVolumesRunner.filter)
	IndexVolumesRunner.Flag(`view`, ``).StringVar(&IndexVolumesRunner.view)
	registry[IndexVolumesCmd.FullCommand()] = IndexVolumesRunner

	ShowVolumeRunner := new(ShowVolumeVolumeRunner)
	ShowVolumeCmd := resCmd.Command("ShowVolume", `Displays information about a single volume.`)
	ShowVolumeRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowVolumeRunner.cloudId)
	ShowVolumeRunner.Flag(`id`, ``).Required().StringVar(&ShowVolumeRunner.id)
	ShowVolumeRunner.Flag(`view`, ``).StringVar(&ShowVolumeRunner.view)
	registry[ShowVolumeCmd.FullCommand()] = ShowVolumeRunner
}

/****** VolumeAttachment ******/

type CreateVolumeAttachmentVolumeAttachmentRunner struct {
	cloudId                      string
	instanceId                   string
	volumeAttachmentDevice       string
	volumeAttachmentInstanceHref string
	volumeAttachmentVolumeHref   string
}

func (r *CreateVolumeAttachmentVolumeAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle volumeAttachment parameter **/
	var volumeAttachment VolumeAttachmentParam

	// Load JSON if provided
	if len(r.volumeAttachmentJson) > 0 {
		if err := Json.Unmarshal(r.volumeAttachmentJson, &volumeAttachment); err != nil {
			return nil, fmt.Errorf("Could not load volumeAttachment JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.volumeAttachmentdevice) > 0 {
		volumeAttachment.volumeAttachment.device = r.volumeAttachmentdevice
	}

	if len(r.volumeAttachmentinstanceHref) > 0 {
		volumeAttachment.volumeAttachment.instanceHref = r.volumeAttachmentinstanceHref
	}

	if len(r.volumeAttachmentvolumeHref) > 0 {
		volumeAttachment.volumeAttachment.volumeHref = r.volumeAttachmentvolumeHref
	}

	return c.CreateVolumeAttachment(r.cloudId, r.instanceId, &volumeAttachment)
}

type DestroyVolumeAttachmentVolumeAttachmentRunner struct {
	cloudId    string
	force      string
	id         string
	instanceId string
}

func (r *DestroyVolumeAttachmentVolumeAttachmentRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyVolumeAttachment(r.cloudId, r.force, r.id, r.instanceId)
}

type IndexVolumeAttachmentsVolumeAttachmentRunner struct {
	cloudId    string
	filter     []string
	filterPos  []string
	instanceId string
	view       string
}

func (r *IndexVolumeAttachmentsVolumeAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexVolumeAttachments(r.cloudId, filter, r.instanceId, r.view)
}

type ShowVolumeAttachmentVolumeAttachmentRunner struct {
	cloudId    string
	id         string
	instanceId string
	view       string
}

func (r *ShowVolumeAttachmentVolumeAttachmentRunner) Run(c *Client) (interface{}, error) {
	return c.ShowVolumeAttachment(r.cloudId, r.id, r.instanceId, r.view)
}

// Register all VolumeAttachment actions
func registerVolumeAttachmentCmds(app *kingpin.Application) {
	resCmd := app.Cmd("VolumeAttachment", `A VolumeAttachment represents a relationship between a volume and an instance.`)

	CreateVolumeAttachmentRunner := new(CreateVolumeAttachmentVolumeAttachmentRunner)
	CreateVolumeAttachmentCmd := resCmd.Command("CreateVolumeAttachment", `Creates a new volume attachment.`)
	CreateVolumeAttachmentRunner.Flag(`cloudId`, ``).Required().StringVar(&CreateVolumeAttachmentRunner.cloudId)
	CreateVolumeAttachmentRunner.Flag(`instanceId`, ``).Required().StringVar(&CreateVolumeAttachmentRunner.instanceId)
	CreateVolumeAttachmentRunner.Flag(`volumeAttachment.device`, `The device location where the volume will be mounted. Value must be of format /dev/xvd[bcefghij]. This is not reliable and will be deprecated.`).StringVar(&CreateVolumeAttachmentRunner.volumeAttachmentDevice)
	CreateVolumeAttachmentRunner.Flag(`volumeAttachment.instanceHref`, `The href of the instance to which the volume will be attached to.`).StringVar(&CreateVolumeAttachmentRunner.volumeAttachmentInstanceHref)
	CreateVolumeAttachmentRunner.Flag(`volumeAttachment.volumeHref`, `The href of the volume to be attached.`).StringVar(&CreateVolumeAttachmentRunner.volumeAttachmentVolumeHref)
	registry[CreateVolumeAttachmentCmd.FullCommand()] = CreateVolumeAttachmentRunner

	DestroyVolumeAttachmentRunner := new(DestroyVolumeAttachmentVolumeAttachmentRunner)
	DestroyVolumeAttachmentCmd := resCmd.Command("DestroyVolumeAttachment", `Deletes a given volume attachment.`)
	DestroyVolumeAttachmentRunner.Flag(`cloudId`, ``).Required().StringVar(&DestroyVolumeAttachmentRunner.cloudId)
	DestroyVolumeAttachmentRunner.Flag(`force`, `Specifies whether to force the detachment of a volume.`).StringVar(&DestroyVolumeAttachmentRunner.force)
	DestroyVolumeAttachmentRunner.Flag(`id`, ``).Required().StringVar(&DestroyVolumeAttachmentRunner.id)
	DestroyVolumeAttachmentRunner.Flag(`instanceId`, ``).Required().StringVar(&DestroyVolumeAttachmentRunner.instanceId)
	registry[DestroyVolumeAttachmentCmd.FullCommand()] = DestroyVolumeAttachmentRunner

	IndexVolumeAttachmentsRunner := new(IndexVolumeAttachmentsVolumeAttachmentRunner)
	IndexVolumeAttachmentsCmd := resCmd.Command("IndexVolumeAttachments", `Lists all volume attachments.`)
	IndexVolumeAttachmentsRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexVolumeAttachmentsRunner.cloudId)
	IndexVolumeAttachmentsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexVolumeAttachmentsRunner.filterPos).StringsVar(&IndexVolumeAttachmentsRunner.filter)
	IndexVolumeAttachmentsRunner.Flag(`instanceId`, ``).Required().StringVar(&IndexVolumeAttachmentsRunner.instanceId)
	IndexVolumeAttachmentsRunner.Flag(`view`, ``).StringVar(&IndexVolumeAttachmentsRunner.view)
	registry[IndexVolumeAttachmentsCmd.FullCommand()] = IndexVolumeAttachmentsRunner

	ShowVolumeAttachmentRunner := new(ShowVolumeAttachmentVolumeAttachmentRunner)
	ShowVolumeAttachmentCmd := resCmd.Command("ShowVolumeAttachment", `Displays information about a single volume attachment.`)
	ShowVolumeAttachmentRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowVolumeAttachmentRunner.cloudId)
	ShowVolumeAttachmentRunner.Flag(`id`, ``).Required().StringVar(&ShowVolumeAttachmentRunner.id)
	ShowVolumeAttachmentRunner.Flag(`instanceId`, ``).Required().StringVar(&ShowVolumeAttachmentRunner.instanceId)
	ShowVolumeAttachmentRunner.Flag(`view`, ``).StringVar(&ShowVolumeAttachmentRunner.view)
	registry[ShowVolumeAttachmentCmd.FullCommand()] = ShowVolumeAttachmentRunner
}

/****** VolumeSnapshot ******/

type CreateVolumeSnapshotVolumeSnapshotRunner struct {
	cloudId                        string
	volumeId                       string
	volumeSnapshotDeploymentHref   string
	volumeSnapshotDescription      string
	volumeSnapshotName             string
	volumeSnapshotParentVolumeHref string
}

func (r *CreateVolumeSnapshotVolumeSnapshotRunner) Run(c *Client) (interface{}, error) {

	/** Handle volumeSnapshot parameter **/
	var volumeSnapshot VolumeSnapshotParam

	// Load JSON if provided
	if len(r.volumeSnapshotJson) > 0 {
		if err := Json.Unmarshal(r.volumeSnapshotJson, &volumeSnapshot); err != nil {
			return nil, fmt.Errorf("Could not load volumeSnapshot JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.volumeSnapshotdeploymentHref) > 0 {
		volumeSnapshot.volumeSnapshot.deploymentHref = r.volumeSnapshotdeploymentHref
	}

	if len(r.volumeSnapshotdescription) > 0 {
		volumeSnapshot.volumeSnapshot.description = r.volumeSnapshotdescription
	}

	if len(r.volumeSnapshotname) > 0 {
		volumeSnapshot.volumeSnapshot.name = r.volumeSnapshotname
	}

	if len(r.volumeSnapshotparentVolumeHref) > 0 {
		volumeSnapshot.volumeSnapshot.parentVolumeHref = r.volumeSnapshotparentVolumeHref
	}

	return c.CreateVolumeSnapshot(r.cloudId, r.volumeId, &volumeSnapshot)
}

type DestroyVolumeSnapshotVolumeSnapshotRunner struct {
	cloudId  string
	id       string
	volumeId string
}

func (r *DestroyVolumeSnapshotVolumeSnapshotRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyVolumeSnapshot(r.cloudId, r.id, r.volumeId)
}

type IndexVolumeSnapshotsVolumeSnapshotRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      string
	volumeId  string
}

func (r *IndexVolumeSnapshotsVolumeSnapshotRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexVolumeSnapshots(r.cloudId, filter, r.view, r.volumeId)
}

type ShowVolumeSnapshotVolumeSnapshotRunner struct {
	cloudId  string
	id       string
	view     string
	volumeId string
}

func (r *ShowVolumeSnapshotVolumeSnapshotRunner) Run(c *Client) (interface{}, error) {
	return c.ShowVolumeSnapshot(r.cloudId, r.id, r.view, r.volumeId)
}

// Register all VolumeSnapshot actions
func registerVolumeSnapshotCmds(app *kingpin.Application) {
	resCmd := app.Cmd("VolumeSnapshot", `A VolumeSnapshot represents a Cloud storage volume at a particular point in time`)

	CreateVolumeSnapshotRunner := new(CreateVolumeSnapshotVolumeSnapshotRunner)
	CreateVolumeSnapshotCmd := resCmd.Command("CreateVolumeSnapshot", `Creates a new volume_snapshot.`)
	CreateVolumeSnapshotRunner.Flag(`cloudId`, ``).Required().StringVar(&CreateVolumeSnapshotRunner.cloudId)
	CreateVolumeSnapshotRunner.Flag(`volumeId`, ``).Required().StringVar(&CreateVolumeSnapshotRunner.volumeId)
	CreateVolumeSnapshotRunner.Flag(`volumeSnapshot.deploymentHref`, `The href of the Deployment that owns this Volume Snapshot.`).StringVar(&CreateVolumeSnapshotRunner.volumeSnapshotDeploymentHref)
	CreateVolumeSnapshotRunner.Flag(`volumeSnapshot.description`, `The description for the Volume Snapshot to be created.`).StringVar(&CreateVolumeSnapshotRunner.volumeSnapshotDescription)
	CreateVolumeSnapshotRunner.Flag(`volumeSnapshot.name`, `The name for the Volume Snapshot to be created.`).Required().StringVar(&CreateVolumeSnapshotRunner.volumeSnapshotName)
	CreateVolumeSnapshotRunner.Flag(`volumeSnapshot.parentVolumeHref`, `The href of the Volume from which the Volume Snapshot will be created.`).StringVar(&CreateVolumeSnapshotRunner.volumeSnapshotParentVolumeHref)
	registry[CreateVolumeSnapshotCmd.FullCommand()] = CreateVolumeSnapshotRunner

	DestroyVolumeSnapshotRunner := new(DestroyVolumeSnapshotVolumeSnapshotRunner)
	DestroyVolumeSnapshotCmd := resCmd.Command("DestroyVolumeSnapshot", `Deletes a given volume_snapshot.`)
	DestroyVolumeSnapshotRunner.Flag(`cloudId`, ``).Required().StringVar(&DestroyVolumeSnapshotRunner.cloudId)
	DestroyVolumeSnapshotRunner.Flag(`id`, ``).Required().StringVar(&DestroyVolumeSnapshotRunner.id)
	DestroyVolumeSnapshotRunner.Flag(`volumeId`, ``).Required().StringVar(&DestroyVolumeSnapshotRunner.volumeId)
	registry[DestroyVolumeSnapshotCmd.FullCommand()] = DestroyVolumeSnapshotRunner

	IndexVolumeSnapshotsRunner := new(IndexVolumeSnapshotsVolumeSnapshotRunner)
	IndexVolumeSnapshotsCmd := resCmd.Command("IndexVolumeSnapshots", `Lists all volume_snapshots.`)
	IndexVolumeSnapshotsRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexVolumeSnapshotsRunner.cloudId)
	IndexVolumeSnapshotsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexVolumeSnapshotsRunner.filterPos).StringsVar(&IndexVolumeSnapshotsRunner.filter)
	IndexVolumeSnapshotsRunner.Flag(`view`, ``).StringVar(&IndexVolumeSnapshotsRunner.view)
	IndexVolumeSnapshotsRunner.Flag(`volumeId`, ``).Required().StringVar(&IndexVolumeSnapshotsRunner.volumeId)
	registry[IndexVolumeSnapshotsCmd.FullCommand()] = IndexVolumeSnapshotsRunner

	ShowVolumeSnapshotRunner := new(ShowVolumeSnapshotVolumeSnapshotRunner)
	ShowVolumeSnapshotCmd := resCmd.Command("ShowVolumeSnapshot", `Displays information about a single volume_snapshot.`)
	ShowVolumeSnapshotRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowVolumeSnapshotRunner.cloudId)
	ShowVolumeSnapshotRunner.Flag(`id`, ``).Required().StringVar(&ShowVolumeSnapshotRunner.id)
	ShowVolumeSnapshotRunner.Flag(`view`, ``).StringVar(&ShowVolumeSnapshotRunner.view)
	ShowVolumeSnapshotRunner.Flag(`volumeId`, ``).Required().StringVar(&ShowVolumeSnapshotRunner.volumeId)
	registry[ShowVolumeSnapshotCmd.FullCommand()] = ShowVolumeSnapshotRunner
}

/****** VolumeType ******/

type IndexVolumeTypesVolumeTypeRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      string
}

func (r *IndexVolumeTypesVolumeTypeRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	return c.IndexVolumeTypes(r.cloudId, filter, r.view)
}

type ShowVolumeTypeVolumeTypeRunner struct {
	cloudId string
	id      string
	view    string
}

func (r *ShowVolumeTypeVolumeTypeRunner) Run(c *Client) (interface{}, error) {
	return c.ShowVolumeType(r.cloudId, r.id, r.view)
}

// Register all VolumeType actions
func registerVolumeTypeCmds(app *kingpin.Application) {
	resCmd := app.Cmd("VolumeType", `A VolumeType describes the type of volume, particularly the size.`)

	IndexVolumeTypesRunner := new(IndexVolumeTypesVolumeTypeRunner)
	IndexVolumeTypesCmd := resCmd.Command("IndexVolumeTypes", `Lists Volume Types.`)
	IndexVolumeTypesRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexVolumeTypesRunner.cloudId)
	IndexVolumeTypesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexVolumeTypesRunner.filterPos).StringsVar(&IndexVolumeTypesRunner.filter)
	IndexVolumeTypesRunner.Flag(`view`, ``).StringVar(&IndexVolumeTypesRunner.view)
	registry[IndexVolumeTypesCmd.FullCommand()] = IndexVolumeTypesRunner

	ShowVolumeTypeRunner := new(ShowVolumeTypeVolumeTypeRunner)
	ShowVolumeTypeCmd := resCmd.Command("ShowVolumeType", `Displays information about a single Volume Type.`)
	ShowVolumeTypeRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowVolumeTypeRunner.cloudId)
	ShowVolumeTypeRunner.Flag(`id`, ``).Required().StringVar(&ShowVolumeTypeRunner.id)
	ShowVolumeTypeRunner.Flag(`view`, ``).StringVar(&ShowVolumeTypeRunner.view)
	registry[ShowVolumeTypeCmd.FullCommand()] = ShowVolumeTypeRunner
}
