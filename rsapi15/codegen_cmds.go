//************************************************************************//
//                RightScale API 1.5 command line client
//
// Generated Feb 12, 2015 at 10:40am (PST)
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

type ShowAccountRunner struct {
	id string
}

func (r *ShowAccountRunner) Run(c *Client) (interface{}, error) {
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

type IndexAccountGroupsRunner struct {
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexAccountGroupsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexAccountGroups(options)
}

type ShowAccountGroupRunner struct {
	id   string
	view *string
}

func (r *ShowAccountGroupRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowAccountGroup(r.id, options)
}

// Register all AccountGroup actions
func registerAccountGroupCmds(app *kingpin.Application) {
	resCmd := app.Cmd("AccountGroup", `An Account Group specifies which RightScale accounts will have access to import a shared RightScale component (e.g. ServerTemplate, RightScript, etc.) from the MultiCloud Marketplace.`)

	IndexAccountGroupsRunner := new(IndexAccountGroupsAccountGroupRunner)
	IndexAccountGroupsCmd := resCmd.Command("IndexAccountGroups", `Lists the AccountGroups owned by this Account.`)
	IndexAccountGroupsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexAccountGroupsRunner.filterPos).StringsVar(IndexAccountGroupsRunner.filter)
	IndexAccountGroupsRunner.Flag(`view`, ``).StringVar(IndexAccountGroupsRunner.view)
	registry[IndexAccountGroupsCmd.FullCommand()] = IndexAccountGroupsRunner

	ShowAccountGroupRunner := new(ShowAccountGroupAccountGroupRunner)
	ShowAccountGroupCmd := resCmd.Command("ShowAccountGroup", `Show information about a single AccountGroup.`)
	ShowAccountGroupRunner.Flag(`id`, ``).Required().StringVar(&ShowAccountGroupRunner.id)
	ShowAccountGroupRunner.Flag(`view`, ``).StringVar(ShowAccountGroupRunner.view)
	registry[ShowAccountGroupCmd.FullCommand()] = ShowAccountGroupRunner
}

/****** Alert ******/

type DisableAlertRunner struct {
	cloudId    string
	id         string
	instanceId string
}

func (r *DisableAlertRunner) Run(c *Client) (interface{}, error) {
	return c.DisableAlert(r.cloudId, r.id, r.instanceId)
}

type EnableAlertRunner struct {
	cloudId    string
	id         string
	instanceId string
}

func (r *EnableAlertRunner) Run(c *Client) (interface{}, error) {
	return c.EnableAlert(r.cloudId, r.id, r.instanceId)
}

type IndexAlertsRunner struct {
	cloudId    string
	filter     []string
	filterPos  []string
	instanceId string
	view       *string
}

func (r *IndexAlertsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexAlerts(r.cloudId, r.instanceId, options)
}

type QuenchAlertRunner struct {
	cloudId    string
	duration   string
	id         string
	instanceId string
}

func (r *QuenchAlertRunner) Run(c *Client) (interface{}, error) {
	return c.QuenchAlert(r.cloudId, r.duration, r.id, r.instanceId)
}

type ShowAlertRunner struct {
	cloudId    string
	id         string
	instanceId string
	view       *string
}

func (r *ShowAlertRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowAlert(r.cloudId, r.id, r.instanceId, options)
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
	IndexAlertsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexAlertsRunner.filterPos).StringsVar(IndexAlertsRunner.filter)
	IndexAlertsRunner.Flag(`instanceId`, ``).Required().StringVar(&IndexAlertsRunner.instanceId)
	IndexAlertsRunner.Flag(`view`, ``).StringVar(IndexAlertsRunner.view)
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
	ShowAlertRunner.Flag(`view`, ``).StringVar(ShowAlertRunner.view)
	registry[ShowAlertCmd.FullCommand()] = ShowAlertRunner
}

/****** AlertSpec ******/

type CreateAlertSpecRunner struct {
	alertSpecCondition      string
	alertSpecDescription    *string
	alertSpecDuration       string
	alertSpecEscalationName *string
	alertSpecFile           string
	alertSpecName           string
	alertSpecSubjectHref    *string
	alertSpecThreshold      string
	alertSpecVariable       string
	alertSpecVoteTag        *string
	alertSpecVoteType       *string
	serverId                string
}

func (r *CreateAlertSpecRunner) Run(c *Client) (interface{}, error) {

	/** Handle alertSpec parameter **/
	var alertSpec AlertSpecParam
	// Load JSON if provided
	if len(r.alertSpecJson) > 0 {
		if err := Json.Unmarshal(r.alertSpecJson, &alertSpec); err != nil {
			return nil, fmt.Errorf("Could not load alertSpec JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.AlertSpecCondition) > 0 {
		alertSpec.condition = r.AlertSpecCondition
	}

	if r.AlertSpecDescription != nil {
		alertSpec.description = r.AlertSpecDescription
	}

	if len(r.AlertSpecDuration) > 0 {
		alertSpec.duration = r.AlertSpecDuration
	}

	if r.AlertSpecEscalationName != nil {
		alertSpec.escalationName = r.AlertSpecEscalationName
	}

	if len(r.AlertSpecFile) > 0 {
		alertSpec.file = r.AlertSpecFile
	}

	if len(r.AlertSpecName) > 0 {
		alertSpec.name = r.AlertSpecName
	}

	if r.AlertSpecSubjectHref != nil {
		alertSpec.subjectHref = r.AlertSpecSubjectHref
	}

	if len(r.AlertSpecThreshold) > 0 {
		alertSpec.threshold = r.AlertSpecThreshold
	}

	if len(r.AlertSpecVariable) > 0 {
		alertSpec.variable = r.AlertSpecVariable
	}

	if r.AlertSpecVoteTag != nil {
		alertSpec.voteTag = r.AlertSpecVoteTag
	}

	if r.AlertSpecVoteType != nil {
		alertSpec.voteType = r.AlertSpecVoteType
	}

	return c.CreateAlertSpec(&alertSpec, r.serverId)
}

type DestroyAlertSpecRunner struct {
	id       string
	serverId string
}

func (r *DestroyAlertSpecRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyAlertSpec(r.id, r.serverId)
}

type IndexAlertSpecsRunner struct {
	filter        []string
	filterPos     []string
	serverId      string
	view          *string
	withInherited *string
}

func (r *IndexAlertSpecsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}
	if r.withInherited != nil {
		options["with_inherited"] = *r.withInherited
	}

	return c.IndexAlertSpecs(r.serverId, options)
}

type ShowAlertSpecRunner struct {
	id       string
	serverId string
	view     *string
}

func (r *ShowAlertSpecRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowAlertSpec(r.id, r.serverId, options)
}

type UpdateAlertSpecRunner struct {
	alertSpecCondition      *string
	alertSpecDescription    *string
	alertSpecDuration       *string
	alertSpecEscalationName *string
	alertSpecFile           *string
	alertSpecName           *string
	alertSpecThreshold      *string
	alertSpecVariable       *string
	alertSpecVoteTag        *string
	alertSpecVoteType       *string
	id                      string
	serverId                string
}

func (r *UpdateAlertSpecRunner) Run(c *Client) (interface{}, error) {

	/** Handle alertSpec parameter **/
	var alertSpec AlertSpecParam2
	// Load JSON if provided
	if len(r.alertSpecJson) > 0 {
		if err := Json.Unmarshal(r.alertSpecJson, &alertSpec); err != nil {
			return nil, fmt.Errorf("Could not load alertSpec JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.AlertSpecCondition != nil {
		alertSpec.condition = r.AlertSpecCondition
	}

	if r.AlertSpecDescription != nil {
		alertSpec.description = r.AlertSpecDescription
	}

	if r.AlertSpecDuration != nil {
		alertSpec.duration = r.AlertSpecDuration
	}

	if r.AlertSpecEscalationName != nil {
		alertSpec.escalationName = r.AlertSpecEscalationName
	}

	if r.AlertSpecFile != nil {
		alertSpec.file = r.AlertSpecFile
	}

	if r.AlertSpecName != nil {
		alertSpec.name = r.AlertSpecName
	}

	if r.AlertSpecThreshold != nil {
		alertSpec.threshold = r.AlertSpecThreshold
	}

	if r.AlertSpecVariable != nil {
		alertSpec.variable = r.AlertSpecVariable
	}

	if r.AlertSpecVoteTag != nil {
		alertSpec.voteTag = r.AlertSpecVoteTag
	}

	if r.AlertSpecVoteType != nil {
		alertSpec.voteType = r.AlertSpecVoteType
	}

	return c.UpdateAlertSpec(&alertSpec, r.id, r.serverId)
}

// Register all AlertSpec actions
func registerAlertSpecCmds(app *kingpin.Application) {
	resCmd := app.Cmd("AlertSpec", `An AlertSpec defines the conditions under which an Alert is triggered and escalated.`)

	CreateAlertSpecRunner := new(CreateAlertSpecAlertSpecRunner)
	CreateAlertSpecCmd := resCmd.Command("CreateAlertSpec", `Creates a new AlertSpec with the given parameters.`)
	CreateAlertSpecRunner.Flag(`alertSpec.condition`, `The condition (operator) in the condition sentence.`).Required().StringVar(&CreateAlertSpecRunner.alertSpecCondition)
	CreateAlertSpecRunner.Flag(`alertSpec.description`, `The description of the AlertSpec.`).StringVar(CreateAlertSpecRunner.alertSpecDescription)
	CreateAlertSpecRunner.Flag(`alertSpec.duration`, `The duration in minutes of the condition sentence.`).Required().StringVar(&CreateAlertSpecRunner.alertSpecDuration)
	CreateAlertSpecRunner.Flag(`alertSpec.escalationName`, `Escalate to the named alert escalation when the alert is triggered. Must either escalate or vote.`).StringVar(CreateAlertSpecRunner.alertSpecEscalationName)
	CreateAlertSpecRunner.Flag(`alertSpec.file`, `The RRD path/file_name of the condition sentence.`).Required().StringVar(&CreateAlertSpecRunner.alertSpecFile)
	CreateAlertSpecRunner.Flag(`alertSpec.name`, `The name of the AlertSpec.`).Required().StringVar(&CreateAlertSpecRunner.alertSpecName)
	CreateAlertSpecRunner.Flag(`alertSpec.subjectHref`, `The href of the resource that this AlertSpec should be associated with. The subject can be a ServerTemplate, Server, ServerArray, or Instance.`).StringVar(CreateAlertSpecRunner.alertSpecSubjectHref)
	CreateAlertSpecRunner.Flag(`alertSpec.threshold`, `The threshold of the condition sentence.`).Required().StringVar(&CreateAlertSpecRunner.alertSpecThreshold)
	CreateAlertSpecRunner.Flag(`alertSpec.variable`, `The RRD variable of the condition sentence.`).Required().StringVar(&CreateAlertSpecRunner.alertSpecVariable)
	CreateAlertSpecRunner.Flag(`alertSpec.voteTag`, `Should correspond to a vote tag on a ServerArray if vote to grow or shrink.`).StringVar(CreateAlertSpecRunner.alertSpecVoteTag)
	CreateAlertSpecRunner.Flag(`alertSpec.voteType`, `Vote to grow or shrink a ServerArray when the alert is triggered. Must either escalate or vote.`).StringVar(CreateAlertSpecRunner.alertSpecVoteType)
	CreateAlertSpecRunner.Flag(`serverId`, ``).Required().StringVar(&CreateAlertSpecRunner.serverId)
	registry[CreateAlertSpecCmd.FullCommand()] = CreateAlertSpecRunner

	DestroyAlertSpecRunner := new(DestroyAlertSpecAlertSpecRunner)
	DestroyAlertSpecCmd := resCmd.Command("DestroyAlertSpec", `Deletes a given AlertSpec.`)
	DestroyAlertSpecRunner.Flag(`id`, ``).Required().StringVar(&DestroyAlertSpecRunner.id)
	DestroyAlertSpecRunner.Flag(`serverId`, ``).Required().StringVar(&DestroyAlertSpecRunner.serverId)
	registry[DestroyAlertSpecCmd.FullCommand()] = DestroyAlertSpecRunner

	IndexAlertSpecsRunner := new(IndexAlertSpecsAlertSpecRunner)
	IndexAlertSpecsCmd := resCmd.Command("IndexAlertSpecs", `<no description> -- Optional parameters: 	filter 	view 	with_inherited: Flag indicating whether or not to include AlertSpecs from the ServerTemplate in the index...`)
	IndexAlertSpecsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexAlertSpecsRunner.filterPos).StringsVar(IndexAlertSpecsRunner.filter)
	IndexAlertSpecsRunner.Flag(`serverId`, ``).Required().StringVar(&IndexAlertSpecsRunner.serverId)
	IndexAlertSpecsRunner.Flag(`view`, ``).StringVar(IndexAlertSpecsRunner.view)
	IndexAlertSpecsRunner.Flag(`withInherited`, `Flag indicating whether or not to include AlertSpecs from the ServerTemplate in the index.`).StringVar(IndexAlertSpecsRunner.withInherited)
	registry[IndexAlertSpecsCmd.FullCommand()] = IndexAlertSpecsRunner

	ShowAlertSpecRunner := new(ShowAlertSpecAlertSpecRunner)
	ShowAlertSpecCmd := resCmd.Command("ShowAlertSpec", `<no description> -- Optional parameters: 	view...`)
	ShowAlertSpecRunner.Flag(`id`, ``).Required().StringVar(&ShowAlertSpecRunner.id)
	ShowAlertSpecRunner.Flag(`serverId`, ``).Required().StringVar(&ShowAlertSpecRunner.serverId)
	ShowAlertSpecRunner.Flag(`view`, ``).StringVar(ShowAlertSpecRunner.view)
	registry[ShowAlertSpecCmd.FullCommand()] = ShowAlertSpecRunner

	UpdateAlertSpecRunner := new(UpdateAlertSpecAlertSpecRunner)
	UpdateAlertSpecCmd := resCmd.Command("UpdateAlertSpec", `Updates an AlertSpec with the given parameters.`)
	UpdateAlertSpecRunner.Flag(`alertSpec.condition`, `The condition (operator) in the condition sentence.`).StringVar(UpdateAlertSpecRunner.alertSpecCondition)
	UpdateAlertSpecRunner.Flag(`alertSpec.description`, `The description of the AlertSpec.`).StringVar(UpdateAlertSpecRunner.alertSpecDescription)
	UpdateAlertSpecRunner.Flag(`alertSpec.duration`, `The duration in minutes of the condition sentence.`).StringVar(UpdateAlertSpecRunner.alertSpecDuration)
	UpdateAlertSpecRunner.Flag(`alertSpec.escalationName`, `Escalate to the named alert escalation when the alert is triggered.`).StringVar(UpdateAlertSpecRunner.alertSpecEscalationName)
	UpdateAlertSpecRunner.Flag(`alertSpec.file`, `The RRD path/file_name of the condition sentence.`).StringVar(UpdateAlertSpecRunner.alertSpecFile)
	UpdateAlertSpecRunner.Flag(`alertSpec.name`, `The name of the AlertSpec.`).StringVar(UpdateAlertSpecRunner.alertSpecName)
	UpdateAlertSpecRunner.Flag(`alertSpec.threshold`, `The threshold of the condition sentence.`).StringVar(UpdateAlertSpecRunner.alertSpecThreshold)
	UpdateAlertSpecRunner.Flag(`alertSpec.variable`, `The RRD variable of the condition sentence.`).StringVar(UpdateAlertSpecRunner.alertSpecVariable)
	UpdateAlertSpecRunner.Flag(`alertSpec.voteTag`, `Should correspond to a vote tag on a ServerArray if vote to grow or shrink.`).StringVar(UpdateAlertSpecRunner.alertSpecVoteTag)
	UpdateAlertSpecRunner.Flag(`alertSpec.voteType`, `Vote to grow or shrink a ServerArray when the alert is triggered.`).StringVar(UpdateAlertSpecRunner.alertSpecVoteType)
	UpdateAlertSpecRunner.Flag(`id`, ``).Required().StringVar(&UpdateAlertSpecRunner.id)
	UpdateAlertSpecRunner.Flag(`serverId`, ``).Required().StringVar(&UpdateAlertSpecRunner.serverId)
	registry[UpdateAlertSpecCmd.FullCommand()] = UpdateAlertSpecRunner
}

/****** AuditEntry ******/

type AppendAuditEntryRunner struct {
	detail  *string
	id      string
	notify  *string
	offset  *int
	summary *string
}

func (r *AppendAuditEntryRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.detail != nil {
		options["detail"] = *r.detail
	}
	if r.notify != nil {
		options["notify"] = *r.notify
	}
	if r.offset != nil {
		options["offset"] = *r.offset
	}
	if r.summary != nil {
		options["summary"] = *r.summary
	}

	return c.AppendAuditEntry(r.id, options)
}

type CreateAuditEntryRunner struct {
	auditEntryAuditeeHref string
	auditEntryDetail      *string
	auditEntrySummary     string
	notify                *string
	userEmail             *string
}

func (r *CreateAuditEntryRunner) Run(c *Client) (interface{}, error) {

	/** Handle auditEntry parameter **/
	var auditEntry AuditEntryParam
	// Load JSON if provided
	if len(r.auditEntryJson) > 0 {
		if err := Json.Unmarshal(r.auditEntryJson, &auditEntry); err != nil {
			return nil, fmt.Errorf("Could not load auditEntry JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.AuditEntryAuditeeHref) > 0 {
		auditEntry.auditeeHref = r.AuditEntryAuditeeHref
	}

	if r.AuditEntryDetail != nil {
		auditEntry.detail = r.AuditEntryDetail
	}

	if len(r.AuditEntrySummary) > 0 {
		auditEntry.summary = r.AuditEntrySummary
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.notify != nil {
		options["notify"] = *r.notify
	}
	if r.userEmail != nil {
		options["user_email"] = *r.userEmail
	}

	return c.CreateAuditEntry(&auditEntry, options)
}

type DetailAuditEntryRunner struct {
	id string
}

func (r *DetailAuditEntryRunner) Run(c *Client) (interface{}, error) {
	return c.DetailAuditEntry(r.id)
}

type IndexAuditEntriesRunner struct {
	endDate   string
	filter    []string
	filterPos []string
	limit     string
	startDate string
	view      *string
}

func (r *IndexAuditEntriesRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexAuditEntries(r.endDate, r.limit, r.startDate, options)
}

type ShowAuditEntryRunner struct {
	id   string
	view *string
}

func (r *ShowAuditEntryRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowAuditEntry(r.id, options)
}

type UpdateAuditEntryRunner struct {
	auditEntryOffset  *int
	auditEntrySummary string
	id                string
	notify            *string
}

func (r *UpdateAuditEntryRunner) Run(c *Client) (interface{}, error) {

	/** Handle auditEntry parameter **/
	var auditEntry AuditEntryParam2
	// Load JSON if provided
	if len(r.auditEntryJson) > 0 {
		if err := Json.Unmarshal(r.auditEntryJson, &auditEntry); err != nil {
			return nil, fmt.Errorf("Could not load auditEntry JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.AuditEntryOffset != nil {
		auditEntry.offset = r.AuditEntryOffset
	}

	if len(r.AuditEntrySummary) > 0 {
		auditEntry.summary = r.AuditEntrySummary
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.notify != nil {
		options["notify"] = *r.notify
	}

	return c.UpdateAuditEntry(&auditEntry, r.id, options)
}

// Register all AuditEntry actions
func registerAuditEntryCmds(app *kingpin.Application) {
	resCmd := app.Cmd("AuditEntry", `An Audit Entry can be used to track various activities of a resource.`)

	AppendAuditEntryRunner := new(AppendAuditEntryAuditEntryRunner)
	AppendAuditEntryCmd := resCmd.Command("AppendAuditEntry", `Updates the summary and appends more details to a given AuditEntry`)
	AppendAuditEntryRunner.Flag(`detail`, `The details to be appended to the audit entry record.`).StringVar(AppendAuditEntryRunner.detail)
	AppendAuditEntryRunner.Flag(`id`, ``).Required().StringVar(&AppendAuditEntryRunner.id)
	AppendAuditEntryRunner.Flag(`notify`, `The event notification category. Defaults to 'None'.`).StringVar(AppendAuditEntryRunner.notify)
	AppendAuditEntryRunner.Flag(`offset`, `The offset where the new details should be appended to in the audit entry's existing details section. Also used in ordering of summary updates. Defaults to end.`).IntVar(AppendAuditEntryRunner.offset)
	AppendAuditEntryRunner.Flag(`summary`, `The updated summary for the audit entry, maximum length is 255 characters.`).StringVar(AppendAuditEntryRunner.summary)
	registry[AppendAuditEntryCmd.FullCommand()] = AppendAuditEntryRunner

	CreateAuditEntryRunner := new(CreateAuditEntryAuditEntryRunner)
	CreateAuditEntryCmd := resCmd.Command("CreateAuditEntry", `Creates a new AuditEntry with the given parameters.`)
	CreateAuditEntryRunner.Flag(`auditEntry.auditeeHref`, `The href of the resource that this audit entry should be associated with (e.g. an instance's href).`).Required().StringVar(&CreateAuditEntryRunner.auditEntryAuditeeHref)
	CreateAuditEntryRunner.Flag(`auditEntry.detail`, `The initial details of the audit entry to be created.`).StringVar(CreateAuditEntryRunner.auditEntryDetail)
	CreateAuditEntryRunner.Flag(`auditEntry.summary`, `The summary of the audit entry to be created, maximum length is 255 characters.`).Required().StringVar(&CreateAuditEntryRunner.auditEntrySummary)
	CreateAuditEntryRunner.Flag(`notify`, `The event notification category. Defaults to 'None'.`).StringVar(CreateAuditEntryRunner.notify)
	CreateAuditEntryRunner.Flag(`userEmail`, `The email of the user (who created/triggered the audit entry). Only usable with instance role.`).StringVar(CreateAuditEntryRunner.userEmail)
	registry[CreateAuditEntryCmd.FullCommand()] = CreateAuditEntryRunner

	DetailAuditEntryRunner := new(DetailAuditEntryAuditEntryRunner)
	DetailAuditEntryCmd := resCmd.Command("DetailAuditEntry", `shows the details of a given AuditEntry.`)
	DetailAuditEntryRunner.Flag(`id`, ``).Required().StringVar(&DetailAuditEntryRunner.id)
	registry[DetailAuditEntryCmd.FullCommand()] = DetailAuditEntryRunner

	IndexAuditEntriesRunner := new(IndexAuditEntriesAuditEntryRunner)
	IndexAuditEntriesCmd := resCmd.Command("IndexAuditEntries", `Lists AuditEntries of the account`)
	IndexAuditEntriesRunner.Flag(`endDate`, `The end date for retrieving audit entries (the format must be the same as start date). The time period between start and end date must be less than 3 months (93 days).`).Required().StringVar(&IndexAuditEntriesRunner.endDate)
	IndexAuditEntriesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexAuditEntriesRunner.filterPos).StringsVar(IndexAuditEntriesRunner.filter)
	IndexAuditEntriesRunner.Flag(`limit`, `Limit the audit entries to this number. The limit should >= 1 and <= 1000`).Required().StringVar(&IndexAuditEntriesRunner.limit)
	IndexAuditEntriesRunner.Flag(`startDate`, `The start date for retrieving audit entries, the format must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g., 2011/06/25 00:00:00 +0000`).Required().StringVar(&IndexAuditEntriesRunner.startDate)
	IndexAuditEntriesRunner.Flag(`view`, ``).StringVar(IndexAuditEntriesRunner.view)
	registry[IndexAuditEntriesCmd.FullCommand()] = IndexAuditEntriesRunner

	ShowAuditEntryRunner := new(ShowAuditEntryAuditEntryRunner)
	ShowAuditEntryCmd := resCmd.Command("ShowAuditEntry", `Lists the attributes of a given audit entry.`)
	ShowAuditEntryRunner.Flag(`id`, ``).Required().StringVar(&ShowAuditEntryRunner.id)
	ShowAuditEntryRunner.Flag(`view`, ``).StringVar(ShowAuditEntryRunner.view)
	registry[ShowAuditEntryCmd.FullCommand()] = ShowAuditEntryRunner

	UpdateAuditEntryRunner := new(UpdateAuditEntryAuditEntryRunner)
	UpdateAuditEntryCmd := resCmd.Command("UpdateAuditEntry", `Updates the summary of a given AuditEntry.`)
	UpdateAuditEntryRunner.Flag(`auditEntry.offset`, `The offset where the next details will be appended. Used in ordering of summary updates.`).IntVar(UpdateAuditEntryRunner.auditEntryOffset)
	UpdateAuditEntryRunner.Flag(`auditEntry.summary`, `The updated summary for the audit entry, maximum length is 255 characters.`).Required().StringVar(&UpdateAuditEntryRunner.auditEntrySummary)
	UpdateAuditEntryRunner.Flag(`id`, ``).Required().StringVar(&UpdateAuditEntryRunner.id)
	UpdateAuditEntryRunner.Flag(`notify`, `The event notification category. Defaults to 'None'.`).StringVar(UpdateAuditEntryRunner.notify)
	registry[UpdateAuditEntryCmd.FullCommand()] = UpdateAuditEntryRunner
}

/****** Backup ******/

type CleanupBackupRunner struct {
	cloudHref *string
	dailies   *string
	keepLast  string
	lineage   string
	monthlies *string
	weeklies  *string
	yearlies  *string
}

func (r *CleanupBackupRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.cloudHref != nil {
		options["cloud_href"] = *r.cloudHref
	}
	if r.dailies != nil {
		options["dailies"] = *r.dailies
	}
	if r.monthlies != nil {
		options["monthlies"] = *r.monthlies
	}
	if r.weeklies != nil {
		options["weeklies"] = *r.weeklies
	}
	if r.yearlies != nil {
		options["yearlies"] = *r.yearlies
	}

	return c.CleanupBackup(r.keepLast, r.lineage, options)
}

type CreateBackupRunner struct {
	backupDescription              *string
	backupFromMaster               *string
	backupLineage                  string
	backupName                     string
	backupVolumeAttachmentHrefs    []string
	backupVolumeAttachmentHrefsPos []string
}

func (r *CreateBackupRunner) Run(c *Client) (interface{}, error) {

	/** Handle backup parameter **/
	var backup BackupParam
	// Load JSON if provided
	if len(r.backupJson) > 0 {
		if err := Json.Unmarshal(r.backupJson, &backup); err != nil {
			return nil, fmt.Errorf("Could not load backup JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.BackupDescription != nil {
		backup.description = r.BackupDescription
	}

	if r.BackupFromMaster != nil {
		backup.fromMaster = r.BackupFromMaster
	}

	if len(r.BackupLineage) > 0 {
		backup.lineage = r.BackupLineage
	}

	if len(r.BackupName) > 0 {
		backup.name = r.BackupName
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxBackupVolumeAttachmentHrefsPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.BackupVolumeAttachmentHrefsPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for backup.volumeAttachmentHrefs field of volumeAttachmentHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of backup.volumeAttachmentHrefs field of volumeAttachmentHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxBackupVolumeAttachmentHrefsPos {
			maxBackupVolumeAttachmentHrefsPos = pos
		}
	}
	if len(r.BackupVolumeAttachmentHrefs) != maxBackupVolumeAttachmentHrefsPos {
		return nil, fmt.Errof("Invalid flags for backup.volumeAttachmentHrefs field of volumeAttachmentHrefs array, %s were defined but max position is %s",
			len(r.BackupVolumeAttachmentHrefs), maxBackupVolumeAttachmentHrefsPos)
	}
	if maxBackupVolumeAttachmentHrefsPos > -1 {
		backupVolumeAttachmentHrefs := make([]*string, maxBackupVolumeAttachmentHrefsPos+1)
		for i := 0; i < maxBackupVolumeAttachmentHrefsPos+1; i++ {
			backupVolumeAttachmentHrefs[i] = r.backupVolumeAttachmentHrefs[r.backupVolumeAttachmentHrefsPos[i]]
		}
		backup.volumeAttachmentHrefs = backupVolumeAttachmentHrefs
	}

	return c.CreateBackup(&backup)
}

type DestroyBackupRunner struct {
	id string
}

func (r *DestroyBackupRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyBackup(r.id)
}

type IndexBackupsRunner struct {
	filter    []string
	filterPos []string
	lineage   string
}

func (r *IndexBackupsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter

	return c.IndexBackups(r.lineage, options)
}

type RestoreBackupRunner struct {
	backupDescription    *string
	backupIops           *string
	backupName           *string
	backupSize           *string
	backupVolumeTypeHref *string
	id                   string
	instanceHref         string
}

func (r *RestoreBackupRunner) Run(c *Client) (interface{}, error) {

	/** Handle backup parameter **/
	var backup BackupParam2
	// Load JSON if provided
	if len(r.backupJson) > 0 {
		if err := Json.Unmarshal(r.backupJson, &backup); err != nil {
			return nil, fmt.Errorf("Could not load backup JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.BackupDescription != nil {
		backup.description = r.BackupDescription
	}

	if r.BackupIops != nil {
		backup.iops = r.BackupIops
	}

	if r.BackupName != nil {
		backup.name = r.BackupName
	}

	if r.BackupSize != nil {
		backup.size = r.BackupSize
	}

	if r.BackupVolumeTypeHref != nil {
		backup.volumeTypeHref = r.BackupVolumeTypeHref
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["backup"] = backup

	return c.RestoreBackup(r.id, r.instanceHref, options)
}

type ShowBackupRunner struct {
	id string
}

func (r *ShowBackupRunner) Run(c *Client) (interface{}, error) {
	return c.ShowBackup(r.id)
}

type UpdateBackupRunner struct {
	backupCommitted string
	id              string
}

func (r *UpdateBackupRunner) Run(c *Client) (interface{}, error) {

	/** Handle backup parameter **/
	var backup BackupParam2
	// Load JSON if provided
	if len(r.backupJson) > 0 {
		if err := Json.Unmarshal(r.backupJson, &backup); err != nil {
			return nil, fmt.Errorf("Could not load backup JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.BackupCommitted) > 0 {
		backup.committed = r.BackupCommitted
	}

	return c.UpdateBackup(&backup, r.id)
}

// Register all Backup actions
func registerBackupCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Backup", ``)

	CleanupBackupRunner := new(CleanupBackupBackupRunner)
	CleanupBackupCmd := resCmd.Command("CleanupBackup", `Deletes old backups that meet the given criteria`)
	CleanupBackupRunner.Flag(`cloudHref`, `Backups belonging to only this cloud are considered for cleanup. Otherwise, all backups in the account with the same lineage will be considered.`).StringVar(CleanupBackupRunner.cloudHref)
	CleanupBackupRunner.Flag(`dailies`, `The number of daily backups(the latest one in each day) that should be kept.`).StringVar(CleanupBackupRunner.dailies)
	CleanupBackupRunner.Flag(`keepLast`, `The number of backups that should be kept.`).Required().StringVar(&CleanupBackupRunner.keepLast)
	CleanupBackupRunner.Flag(`lineage`, `The lineage of the backups that are to be cleaned-up.`).Required().StringVar(&CleanupBackupRunner.lineage)
	CleanupBackupRunner.Flag(`monthlies`, `The number of monthly backups(the latest one in each month) that should be kept.`).StringVar(CleanupBackupRunner.monthlies)
	CleanupBackupRunner.Flag(`weeklies`, `The number of weekly backups(the latest one in each week) that should be kept.`).StringVar(CleanupBackupRunner.weeklies)
	CleanupBackupRunner.Flag(`yearlies`, `The number of yearly backups(the latest one in each year) that should be kept.`).StringVar(CleanupBackupRunner.yearlies)
	registry[CleanupBackupCmd.FullCommand()] = CleanupBackupRunner

	CreateBackupRunner := new(CreateBackupBackupRunner)
	CreateBackupCmd := resCmd.Command("CreateBackup", `Takes in an array of volumeattachmenthrefs and takes a snapshot of each.`)
	CreateBackupRunner.Flag(`backup.description`, `The description to be set on each of the volume snapshots`).StringVar(CreateBackupRunner.backupDescription)
	CreateBackupRunner.Flag(`backup.fromMaster`, `Setting this to 'true' will create a tag 'rs_backup:from_master=true' on the snapshots so that one can filter them later.`).StringVar(CreateBackupRunner.backupFromMaster)
	CreateBackupRunner.Flag(`backup.lineage`, `A unique value to create backups belonging to a particular system.`).Required().StringVar(&CreateBackupRunner.backupLineage)
	CreateBackupRunner.Flag(`backup.name`, `The name to be set on each of the volume snapshots.`).Required().StringVar(&CreateBackupRunner.backupName)
	CreateBackupRunner.FlagPattern(`backup\.volumeAttachmentHrefs\.(\d+)`, `List of volume attachment hrefs that are to be backed-up.`).Required().Capture(&CreateBackupRunner.backupVolumeAttachmentHrefsPos).StringsVar(&CreateBackupRunner.backupVolumeAttachmentHrefs)
	registry[CreateBackupCmd.FullCommand()] = CreateBackupRunner

	DestroyBackupRunner := new(DestroyBackupBackupRunner)
	DestroyBackupCmd := resCmd.Command("DestroyBackup", `Deletes a given backup by deleting all of its snapshots, this call will succeed even if the backup has not completed.`)
	DestroyBackupRunner.Flag(`id`, ``).Required().StringVar(&DestroyBackupRunner.id)
	registry[DestroyBackupCmd.FullCommand()] = DestroyBackupRunner

	IndexBackupsRunner := new(IndexBackupsBackupRunner)
	IndexBackupsCmd := resCmd.Command("IndexBackups", `Lists all of the backups with the given lineage tag`)
	IndexBackupsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexBackupsRunner.filterPos).StringsVar(IndexBackupsRunner.filter)
	IndexBackupsRunner.Flag(`lineage`, `Backups belonging to this lineage.`).Required().StringVar(&IndexBackupsRunner.lineage)
	registry[IndexBackupsCmd.FullCommand()] = IndexBackupsRunner

	RestoreBackupRunner := new(RestoreBackupBackupRunner)
	RestoreBackupCmd := resCmd.Command("RestoreBackup", `Restores the given Backup.`)
	RestoreBackupRunner.Flag(`backup.description`, `Each volume is created with this description instead of the volume snapshot's description`).StringVar(RestoreBackupRunner.backupDescription)
	RestoreBackupRunner.Flag(`backup.iops`, `The number of IOPS (I/O Operations Per Second) each volume should support. Only available on clouds supporting performance provisioning.`).StringVar(RestoreBackupRunner.backupIops)
	RestoreBackupRunner.Flag(`backup.name`, `Each volume is created with this name instead of the volume snapshot's name`).StringVar(RestoreBackupRunner.backupName)
	RestoreBackupRunner.Flag(`backup.size`, `Each volume is created with this size in gigabytes (GB) instead of the volume snapshot's size (must be equal or larger). Some volume types have predefined sizes and do not allow selecting a custom size on volume creation.`).StringVar(RestoreBackupRunner.backupSize)
	RestoreBackupRunner.Flag(`backup.volumeTypeHref`, `The href of the volume type. Each volume is created with this volume type instead of the default volume type for the cloud. A Name, Resource UID and optional Size is associated with a volume type.`).StringVar(RestoreBackupRunner.backupVolumeTypeHref)
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

type CreateChildAccountRunner struct {
	childAccountClusterHref *string
	childAccountName        string
}

func (r *CreateChildAccountRunner) Run(c *Client) (interface{}, error) {

	/** Handle childAccount parameter **/
	var childAccount ChildAccountParam
	// Load JSON if provided
	if len(r.childAccountJson) > 0 {
		if err := Json.Unmarshal(r.childAccountJson, &childAccount); err != nil {
			return nil, fmt.Errorf("Could not load childAccount JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.ChildAccountClusterHref != nil {
		childAccount.clusterHref = r.ChildAccountClusterHref
	}

	if len(r.ChildAccountName) > 0 {
		childAccount.name = r.ChildAccountName
	}

	return c.CreateChildAccount(&childAccount)
}

type IndexChildAccountsRunner struct {
	filter    []string
	filterPos []string
}

func (r *IndexChildAccountsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter

	return c.IndexChildAccounts(options)
}

type UpdateChildAccountRunner struct {
	childAccountName *string
	id               string
}

func (r *UpdateChildAccountRunner) Run(c *Client) (interface{}, error) {

	/** Handle childAccount parameter **/
	var childAccount ChildAccountParam2
	// Load JSON if provided
	if len(r.childAccountJson) > 0 {
		if err := Json.Unmarshal(r.childAccountJson, &childAccount); err != nil {
			return nil, fmt.Errorf("Could not load childAccount JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.ChildAccountName != nil {
		childAccount.name = r.ChildAccountName
	}

	return c.UpdateChildAccount(&childAccount, r.id)
}

// Register all ChildAccount actions
func registerChildAccountCmds(app *kingpin.Application) {
	resCmd := app.Cmd("ChildAccount", ``)

	CreateChildAccountRunner := new(CreateChildAccountChildAccountRunner)
	CreateChildAccountCmd := resCmd.Command("CreateChildAccount", `Create an enterprise ChildAccount for this Account`)
	CreateChildAccountRunner.Flag(`childAccount.clusterHref`, `The href of the cluster in which to create the account. If not specified, will default to the cluster of the parent account.`).StringVar(CreateChildAccountRunner.childAccountClusterHref)
	CreateChildAccountRunner.Flag(`childAccount.name`, ``).Required().StringVar(&CreateChildAccountRunner.childAccountName)
	registry[CreateChildAccountCmd.FullCommand()] = CreateChildAccountRunner

	IndexChildAccountsRunner := new(IndexChildAccountsChildAccountRunner)
	IndexChildAccountsCmd := resCmd.Command("IndexChildAccounts", `Lists the enterprise ChildAccounts available for this Account.`)
	IndexChildAccountsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexChildAccountsRunner.filterPos).StringsVar(IndexChildAccountsRunner.filter)
	registry[IndexChildAccountsCmd.FullCommand()] = IndexChildAccountsRunner

	UpdateChildAccountRunner := new(UpdateChildAccountChildAccountRunner)
	UpdateChildAccountCmd := resCmd.Command("UpdateChildAccount", `Update an enterprise ChildAccount for this Account.`)
	UpdateChildAccountRunner.Flag(`childAccount.name`, `The updated name for the account.`).StringVar(UpdateChildAccountRunner.childAccountName)
	UpdateChildAccountRunner.Flag(`id`, ``).Required().StringVar(&UpdateChildAccountRunner.id)
	registry[UpdateChildAccountCmd.FullCommand()] = UpdateChildAccountRunner
}

/****** Cloud ******/

type IndexCloudsRunner struct {
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexCloudsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexClouds(options)
}

type ShowCloudRunner struct {
	id   string
	view *string
}

func (r *ShowCloudRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowCloud(r.id, options)
}

// Register all Cloud actions
func registerCloudCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Cloud", `Represents a Cloud (within the context of the account in the session).`)

	IndexCloudsRunner := new(IndexCloudsCloudRunner)
	IndexCloudsCmd := resCmd.Command("IndexClouds", `Lists the clouds available to this account.`)
	IndexCloudsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexCloudsRunner.filterPos).StringsVar(IndexCloudsRunner.filter)
	IndexCloudsRunner.Flag(`view`, ``).StringVar(IndexCloudsRunner.view)
	registry[IndexCloudsCmd.FullCommand()] = IndexCloudsRunner

	ShowCloudRunner := new(ShowCloudCloudRunner)
	ShowCloudCmd := resCmd.Command("ShowCloud", `Show information about a single cloud`)
	ShowCloudRunner.Flag(`id`, ``).Required().StringVar(&ShowCloudRunner.id)
	ShowCloudRunner.Flag(`view`, ``).StringVar(ShowCloudRunner.view)
	registry[ShowCloudCmd.FullCommand()] = ShowCloudRunner
}

/****** CloudAccount ******/

type CreateCloudAccountRunner struct {
	cloudAccountCloudHref   *string
	cloudAccountCredsValues []string
	cloudAccountCredsNames  []string
	cloudAccountToken       *string
}

func (r *CreateCloudAccountRunner) Run(c *Client) (interface{}, error) {

	/** Handle cloudAccount parameter **/
	var cloudAccount CloudAccountParam
	// Load JSON if provided
	if len(r.cloudAccountJson) > 0 {
		if err := Json.Unmarshal(r.cloudAccountJson, &cloudAccount); err != nil {
			return nil, fmt.Errorf("Could not load cloudAccount JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.CloudAccountCloudHref != nil {
		cloudAccount.cloudHref = r.CloudAccountCloudHref
	}

	if r.CloudAccountToken != nil {
		cloudAccount.token = r.CloudAccountToken
	}

	// Create enumarable fields from flags
	cloudAccountCreds := make(map[string]string, len(r.CloudAccountCredsNames))
	for i, n := range r.CloudAccountCredsNames {
		cloudAccountCreds[n] = r.CloudAccountCredsValues[i]
	}
	cloudAccount.cloudAccount.creds = cloudAccountCreds

	return c.CreateCloudAccount(&cloudAccount)
}

type DestroyCloudAccountRunner struct {
	id string
}

func (r *DestroyCloudAccountRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyCloudAccount(r.id)
}

type IndexCloudAccountsRunner struct {
}

func (r *IndexCloudAccountsRunner) Run(c *Client) (interface{}, error) {
	return c.IndexCloudAccounts()
}

type ShowCloudAccountRunner struct {
	id string
}

func (r *ShowCloudAccountRunner) Run(c *Client) (interface{}, error) {
	return c.ShowCloudAccount(r.id)
}

// Register all CloudAccount actions
func registerCloudAccountCmds(app *kingpin.Application) {
	resCmd := app.Cmd("CloudAccount", `Represents a Cloud Account (An association between the account and a cloud).`)

	CreateCloudAccountRunner := new(CreateCloudAccountCloudAccountRunner)
	CreateCloudAccountCmd := resCmd.Command("CreateCloudAccount", `Create a CloudAccount by passing in the respective credentials for each cloud.`)
	CreateCloudAccountRunner.Flag(`cloudAccount.cloudHref`, `The href of the cloud if it is known. For valid values see support portal link above.`).StringVar(CreateCloudAccountRunner.cloudAccountCloudHref)
	CreateCloudAccountRunner.FlagPattern(`cloudAccount\.creds\.([a-z0-9_]+)`, ``).Required().Capture(&CreateCloudAccountRunner.cloudAccountCredsNames).StringVar(&CreateCloudAccountRunner.cloudAccountCredsValues)
	CreateCloudAccountRunner.Flag(`cloudAccount.token`, `The cloud token to identify a private cloud`).StringVar(CreateCloudAccountRunner.cloudAccountToken)
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

type DestroyCookbookRunner struct {
	id string
}

func (r *DestroyCookbookRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyCookbook(r.id)
}

type FollowCookbookRunner struct {
	id    string
	value string
}

func (r *FollowCookbookRunner) Run(c *Client) (interface{}, error) {
	return c.FollowCookbook(r.id, r.value)
}

type FreezeCookbookRunner struct {
	id    string
	value string
}

func (r *FreezeCookbookRunner) Run(c *Client) (interface{}, error) {
	return c.FreezeCookbook(r.id, r.value)
}

type IndexCookbooksRunner struct {
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexCookbooksRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexCookbooks(options)
}

type ObsoleteCookbookRunner struct {
	id    string
	value string
}

func (r *ObsoleteCookbookRunner) Run(c *Client) (interface{}, error) {
	return c.ObsoleteCookbook(r.id, r.value)
}

type ShowCookbookRunner struct {
	id   string
	view *string
}

func (r *ShowCookbookRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowCookbook(r.id, options)
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
	IndexCookbooksRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexCookbooksRunner.filterPos).StringsVar(IndexCookbooksRunner.filter)
	IndexCookbooksRunner.Flag(`view`, ``).StringVar(IndexCookbooksRunner.view)
	registry[IndexCookbooksCmd.FullCommand()] = IndexCookbooksRunner

	ObsoleteCookbookRunner := new(ObsoleteCookbookCookbookRunner)
	ObsoleteCookbookCmd := resCmd.Command("ObsoleteCookbook", `Marks a Cookbook as obsolete (or un-obsolete).`)
	ObsoleteCookbookRunner.Flag(`id`, ``).Required().StringVar(&ObsoleteCookbookRunner.id)
	ObsoleteCookbookRunner.Flag(`value`, `Indicates if this action should obsolete (true) or un-obsolete (false) a Cookbook.`).Required().StringVar(&ObsoleteCookbookRunner.value)
	registry[ObsoleteCookbookCmd.FullCommand()] = ObsoleteCookbookRunner

	ShowCookbookRunner := new(ShowCookbookCookbookRunner)
	ShowCookbookCmd := resCmd.Command("ShowCookbook", `Show information about a single Cookbook.`)
	ShowCookbookRunner.Flag(`id`, ``).Required().StringVar(&ShowCookbookRunner.id)
	ShowCookbookRunner.Flag(`view`, ``).StringVar(ShowCookbookRunner.view)
	registry[ShowCookbookCmd.FullCommand()] = ShowCookbookRunner
}

/****** CookbookAttachment ******/

type CreateCookbookAttachmentRunner struct {
	cookbookAttachmentCookbookHref       *string
	cookbookAttachmentServerTemplateHref *string
	cookbookId                           string
}

func (r *CreateCookbookAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle cookbookAttachment parameter **/
	var cookbookAttachment CookbookAttachmentParam
	// Load JSON if provided
	if len(r.cookbookAttachmentJson) > 0 {
		if err := Json.Unmarshal(r.cookbookAttachmentJson, &cookbookAttachment); err != nil {
			return nil, fmt.Errorf("Could not load cookbookAttachment JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.CookbookAttachmentCookbookHref != nil {
		cookbookAttachment.cookbookHref = r.CookbookAttachmentCookbookHref
	}

	if r.CookbookAttachmentServerTemplateHref != nil {
		cookbookAttachment.serverTemplateHref = r.CookbookAttachmentServerTemplateHref
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["cookbook_attachment"] = cookbookAttachment

	return c.CreateCookbookAttachment(r.cookbookId, options)
}

type DestroyCookbookAttachmentRunner struct {
	cookbookId string
	id         string
}

func (r *DestroyCookbookAttachmentRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyCookbookAttachment(r.cookbookId, r.id)
}

type IndexCookbookAttachmentsRunner struct {
	cookbookId string
	view       *string
}

func (r *IndexCookbookAttachmentsRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexCookbookAttachments(r.cookbookId, options)
}

type MultiAttachCookbookAttachmentsRunner struct {
	cookbookAttachmentsCookbookHrefs      []string
	cookbookAttachmentsCookbookHrefsPos   []string
	cookbookAttachmentsServerTemplateHref *string
	serverTemplateId                      string
}

func (r *MultiAttachCookbookAttachmentsRunner) Run(c *Client) (interface{}, error) {

	/** Handle cookbookAttachments parameter **/
	var cookbookAttachments CookbookAttachments
	// Load JSON if provided
	if len(r.cookbookAttachmentsJson) > 0 {
		if err := Json.Unmarshal(r.cookbookAttachmentsJson, &cookbookAttachments); err != nil {
			return nil, fmt.Errorf("Could not load cookbookAttachments JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.CookbookAttachmentsServerTemplateHref != nil {
		cookbookAttachments.serverTemplateHref = r.CookbookAttachmentsServerTemplateHref
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxCookbookAttachmentsCookbookHrefsPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.CookbookAttachmentsCookbookHrefsPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for cookbookAttachments.cookbookHrefs field of cookbookHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of cookbookAttachments.cookbookHrefs field of cookbookHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxCookbookAttachmentsCookbookHrefsPos {
			maxCookbookAttachmentsCookbookHrefsPos = pos
		}
	}
	if len(r.CookbookAttachmentsCookbookHrefs) != maxCookbookAttachmentsCookbookHrefsPos {
		return nil, fmt.Errof("Invalid flags for cookbookAttachments.cookbookHrefs field of cookbookHrefs array, %s were defined but max position is %s",
			len(r.CookbookAttachmentsCookbookHrefs), maxCookbookAttachmentsCookbookHrefsPos)
	}
	if maxCookbookAttachmentsCookbookHrefsPos > -1 {
		cookbookAttachmentsCookbookHrefs := make([]*string, maxCookbookAttachmentsCookbookHrefsPos+1)
		for i := 0; i < maxCookbookAttachmentsCookbookHrefsPos+1; i++ {
			cookbookAttachmentsCookbookHrefs[i] = r.cookbookAttachmentsCookbookHrefs[r.cookbookAttachmentsCookbookHrefsPos[i]]
		}
		cookbookAttachments.cookbookHrefs = cookbookAttachmentsCookbookHrefs
	}

	return c.MultiAttachCookbookAttachments(&cookbookAttachments, r.serverTemplateId)
}

type MultiDetachCookbookAttachmentsRunner struct {
	cookbookAttachmentsCookbookAttachmentHrefs    []string
	cookbookAttachmentsCookbookAttachmentHrefsPos []string
	serverTemplateId                              string
}

func (r *MultiDetachCookbookAttachmentsRunner) Run(c *Client) (interface{}, error) {

	/** Handle cookbookAttachments parameter **/
	var cookbookAttachments CookbookAttachments2
	// Load JSON if provided
	if len(r.cookbookAttachmentsJson) > 0 {
		if err := Json.Unmarshal(r.cookbookAttachmentsJson, &cookbookAttachments); err != nil {
			return nil, fmt.Errorf("Could not load cookbookAttachments JSON: %s", err.Error())
		}
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxCookbookAttachmentsCookbookAttachmentHrefsPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.CookbookAttachmentsCookbookAttachmentHrefsPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for cookbookAttachments.cookbookAttachmentHrefs field of cookbookAttachmentHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of cookbookAttachments.cookbookAttachmentHrefs field of cookbookAttachmentHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxCookbookAttachmentsCookbookAttachmentHrefsPos {
			maxCookbookAttachmentsCookbookAttachmentHrefsPos = pos
		}
	}
	if len(r.CookbookAttachmentsCookbookAttachmentHrefs) != maxCookbookAttachmentsCookbookAttachmentHrefsPos {
		return nil, fmt.Errof("Invalid flags for cookbookAttachments.cookbookAttachmentHrefs field of cookbookAttachmentHrefs array, %s were defined but max position is %s",
			len(r.CookbookAttachmentsCookbookAttachmentHrefs), maxCookbookAttachmentsCookbookAttachmentHrefsPos)
	}
	if maxCookbookAttachmentsCookbookAttachmentHrefsPos > -1 {
		cookbookAttachmentsCookbookAttachmentHrefs := make([]*string, maxCookbookAttachmentsCookbookAttachmentHrefsPos+1)
		for i := 0; i < maxCookbookAttachmentsCookbookAttachmentHrefsPos+1; i++ {
			cookbookAttachmentsCookbookAttachmentHrefs[i] = r.cookbookAttachmentsCookbookAttachmentHrefs[r.cookbookAttachmentsCookbookAttachmentHrefsPos[i]]
		}
		cookbookAttachments.cookbookAttachmentHrefs = cookbookAttachmentsCookbookAttachmentHrefs
	}

	return c.MultiDetachCookbookAttachments(&cookbookAttachments, r.serverTemplateId)
}

type ShowCookbookAttachmentRunner struct {
	cookbookId string
	id         string
	view       *string
}

func (r *ShowCookbookAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowCookbookAttachment(r.cookbookId, r.id, options)
}

// Register all CookbookAttachment actions
func registerCookbookAttachmentCmds(app *kingpin.Application) {
	resCmd := app.Cmd("CookbookAttachment", `Cookbook Attachment is used to associate a particular cookbook with a ServerTemplate. A Cookbook Attachment must be in place before a recipe can be bound to a runlist using RunnableBinding.`)

	CreateCookbookAttachmentRunner := new(CreateCookbookAttachmentCookbookAttachmentRunner)
	CreateCookbookAttachmentCmd := resCmd.Command("CreateCookbookAttachment", `Attach a cookbook to a given resource.`)
	CreateCookbookAttachmentRunner.Flag(`cookbookAttachment.cookbookHref`, `The href of the cookbook to attach.`).StringVar(CreateCookbookAttachmentRunner.cookbookAttachmentCookbookHref)
	CreateCookbookAttachmentRunner.Flag(`cookbookAttachment.serverTemplateHref`, `The href of the server template to attach the cookbook to.`).StringVar(CreateCookbookAttachmentRunner.cookbookAttachmentServerTemplateHref)
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
	IndexCookbookAttachmentsRunner.Flag(`view`, ``).StringVar(IndexCookbookAttachmentsRunner.view)
	registry[IndexCookbookAttachmentsCmd.FullCommand()] = IndexCookbookAttachmentsRunner

	MultiAttachCookbookAttachmentsRunner := new(MultiAttachCookbookAttachmentsCookbookAttachmentRunner)
	MultiAttachCookbookAttachmentsCmd := resCmd.Command("MultiAttachCookbookAttachments", `Attach multiple cookbooks to a given resource.`)
	MultiAttachCookbookAttachmentsRunner.FlagPattern(`cookbookAttachments\.cookbookHrefs\.(\d+)`, `The hrefs of the cookbooks to be attached`).Capture(&MultiAttachCookbookAttachmentsRunner.cookbookAttachmentsCookbookHrefsPos).StringsVar(MultiAttachCookbookAttachmentsRunner.cookbookAttachmentsCookbookHrefs)
	MultiAttachCookbookAttachmentsRunner.Flag(`cookbookAttachments.serverTemplateHref`, `The href of the server template to attach the cookbooks to.`).StringVar(MultiAttachCookbookAttachmentsRunner.cookbookAttachmentsServerTemplateHref)
	MultiAttachCookbookAttachmentsRunner.Flag(`serverTemplateId`, ``).Required().StringVar(&MultiAttachCookbookAttachmentsRunner.serverTemplateId)
	registry[MultiAttachCookbookAttachmentsCmd.FullCommand()] = MultiAttachCookbookAttachmentsRunner

	MultiDetachCookbookAttachmentsRunner := new(MultiDetachCookbookAttachmentsCookbookAttachmentRunner)
	MultiDetachCookbookAttachmentsCmd := resCmd.Command("MultiDetachCookbookAttachments", `Detach multiple cookbooks from a given resource.`)
	MultiDetachCookbookAttachmentsRunner.FlagPattern(`cookbookAttachments\.cookbookAttachmentHrefs\.(\d+)`, `The hrefs of the cookbook attachments to be detached`).Capture(&MultiDetachCookbookAttachmentsRunner.cookbookAttachmentsCookbookAttachmentHrefsPos).StringsVar(MultiDetachCookbookAttachmentsRunner.cookbookAttachmentsCookbookAttachmentHrefs)
	MultiDetachCookbookAttachmentsRunner.Flag(`serverTemplateId`, ``).Required().StringVar(&MultiDetachCookbookAttachmentsRunner.serverTemplateId)
	registry[MultiDetachCookbookAttachmentsCmd.FullCommand()] = MultiDetachCookbookAttachmentsRunner

	ShowCookbookAttachmentRunner := new(ShowCookbookAttachmentCookbookAttachmentRunner)
	ShowCookbookAttachmentCmd := resCmd.Command("ShowCookbookAttachment", `Displays information about a single cookbook attachment to a ServerTemplate.`)
	ShowCookbookAttachmentRunner.Flag(`cookbookId`, ``).Required().StringVar(&ShowCookbookAttachmentRunner.cookbookId)
	ShowCookbookAttachmentRunner.Flag(`id`, ``).Required().StringVar(&ShowCookbookAttachmentRunner.id)
	ShowCookbookAttachmentRunner.Flag(`view`, ``).StringVar(ShowCookbookAttachmentRunner.view)
	registry[ShowCookbookAttachmentCmd.FullCommand()] = ShowCookbookAttachmentRunner
}

/****** Credential ******/

type CreateCredentialRunner struct {
	credentialDescription *string
	credentialName        string
	credentialValue       string
}

func (r *CreateCredentialRunner) Run(c *Client) (interface{}, error) {

	/** Handle credential parameter **/
	var credential CredentialParam
	// Load JSON if provided
	if len(r.credentialJson) > 0 {
		if err := Json.Unmarshal(r.credentialJson, &credential); err != nil {
			return nil, fmt.Errorf("Could not load credential JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.CredentialDescription != nil {
		credential.description = r.CredentialDescription
	}

	if len(r.CredentialName) > 0 {
		credential.name = r.CredentialName
	}

	if len(r.CredentialValue) > 0 {
		credential.value = r.CredentialValue
	}

	return c.CreateCredential(&credential)
}

type DestroyCredentialRunner struct {
	id string
}

func (r *DestroyCredentialRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyCredential(r.id)
}

type IndexCredentialsRunner struct {
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexCredentialsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexCredentials(options)
}

type ShowCredentialRunner struct {
	id   string
	view *string
}

func (r *ShowCredentialRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowCredential(r.id, options)
}

type UpdateCredentialRunner struct {
	credentialDescription *string
	credentialName        *string
	credentialValue       *string
	id                    string
}

func (r *UpdateCredentialRunner) Run(c *Client) (interface{}, error) {

	/** Handle credential parameter **/
	var credential CredentialParam2
	// Load JSON if provided
	if len(r.credentialJson) > 0 {
		if err := Json.Unmarshal(r.credentialJson, &credential); err != nil {
			return nil, fmt.Errorf("Could not load credential JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.CredentialDescription != nil {
		credential.description = r.CredentialDescription
	}

	if r.CredentialName != nil {
		credential.name = r.CredentialName
	}

	if r.CredentialValue != nil {
		credential.value = r.CredentialValue
	}

	return c.UpdateCredential(&credential, r.id)
}

// Register all Credential actions
func registerCredentialCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Credential", `A Credential provides an abstraction for sensitive textual information, such as passphrases or cloud credentials, whose value should be encrypted when stored in the database and not generally shown in the UI or through the API...`)

	CreateCredentialRunner := new(CreateCredentialCredentialRunner)
	CreateCredentialCmd := resCmd.Command("CreateCredential", `Creates a new Credential with the given parameters.`)
	CreateCredentialRunner.Flag(`credential.description`, `The description of the Credential to be created.`).StringVar(CreateCredentialRunner.credentialDescription)
	CreateCredentialRunner.Flag(`credential.name`, `The name of the Credential to be created.`).Required().StringVar(&CreateCredentialRunner.credentialName)
	CreateCredentialRunner.Flag(`credential.value`, `The value of the Credential to be created.`).Required().StringVar(&CreateCredentialRunner.credentialValue)
	registry[CreateCredentialCmd.FullCommand()] = CreateCredentialRunner

	DestroyCredentialRunner := new(DestroyCredentialCredentialRunner)
	DestroyCredentialCmd := resCmd.Command("DestroyCredential", `Deletes a Credential.`)
	DestroyCredentialRunner.Flag(`id`, ``).Required().StringVar(&DestroyCredentialRunner.id)
	registry[DestroyCredentialCmd.FullCommand()] = DestroyCredentialRunner

	IndexCredentialsRunner := new(IndexCredentialsCredentialRunner)
	IndexCredentialsCmd := resCmd.Command("IndexCredentials", `Lists the Credentials available to this account.`)
	IndexCredentialsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexCredentialsRunner.filterPos).StringsVar(IndexCredentialsRunner.filter)
	IndexCredentialsRunner.Flag(`view`, ``).StringVar(IndexCredentialsRunner.view)
	registry[IndexCredentialsCmd.FullCommand()] = IndexCredentialsRunner

	ShowCredentialRunner := new(ShowCredentialCredentialRunner)
	ShowCredentialCmd := resCmd.Command("ShowCredential", `Show information about a single Credential. NOTE: Credential values may be updated through the API, but values cannot be retrieved via the API.`)
	ShowCredentialRunner.Flag(`id`, ``).Required().StringVar(&ShowCredentialRunner.id)
	ShowCredentialRunner.Flag(`view`, ``).StringVar(ShowCredentialRunner.view)
	registry[ShowCredentialCmd.FullCommand()] = ShowCredentialRunner

	UpdateCredentialRunner := new(UpdateCredentialCredentialRunner)
	UpdateCredentialCmd := resCmd.Command("UpdateCredential", `Updates attributes of a Credential.`)
	UpdateCredentialRunner.Flag(`credential.description`, `The updated description of the Credential.`).StringVar(UpdateCredentialRunner.credentialDescription)
	UpdateCredentialRunner.Flag(`credential.name`, `The updated name of the Credential.`).StringVar(UpdateCredentialRunner.credentialName)
	UpdateCredentialRunner.Flag(`credential.value`, `The updated value of the Credential.`).StringVar(UpdateCredentialRunner.credentialValue)
	UpdateCredentialRunner.Flag(`id`, ``).Required().StringVar(&UpdateCredentialRunner.id)
	registry[UpdateCredentialCmd.FullCommand()] = UpdateCredentialRunner
}

/****** Datacenter ******/

type IndexDatacentersRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexDatacentersRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexDatacenters(r.cloudId, options)
}

type ShowDatacenterRunner struct {
	cloudId string
	id      string
	view    *string
}

func (r *ShowDatacenterRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowDatacenter(r.cloudId, r.id, options)
}

// Register all Datacenter actions
func registerDatacenterCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Datacenter", `Datacenters represent isolated facilities within a cloud`)

	IndexDatacentersRunner := new(IndexDatacentersDatacenterRunner)
	IndexDatacentersCmd := resCmd.Command("IndexDatacenters", `Lists all Datacenters for a particular cloud.`)
	IndexDatacentersRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexDatacentersRunner.cloudId)
	IndexDatacentersRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexDatacentersRunner.filterPos).StringsVar(IndexDatacentersRunner.filter)
	IndexDatacentersRunner.Flag(`view`, ``).StringVar(IndexDatacentersRunner.view)
	registry[IndexDatacentersCmd.FullCommand()] = IndexDatacentersRunner

	ShowDatacenterRunner := new(ShowDatacenterDatacenterRunner)
	ShowDatacenterCmd := resCmd.Command("ShowDatacenter", `Displays information about a single Datacenter.`)
	ShowDatacenterRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowDatacenterRunner.cloudId)
	ShowDatacenterRunner.Flag(`id`, ``).Required().StringVar(&ShowDatacenterRunner.id)
	ShowDatacenterRunner.Flag(`view`, ``).StringVar(ShowDatacenterRunner.view)
	registry[ShowDatacenterCmd.FullCommand()] = ShowDatacenterRunner
}

/****** Deployment ******/

type CloneDeploymentRunner struct {
	deploymentDescription    *string
	deploymentName           *string
	deploymentServerTagScope *string
	id                       string
}

func (r *CloneDeploymentRunner) Run(c *Client) (interface{}, error) {

	/** Handle deployment parameter **/
	var deployment DeploymentParam
	// Load JSON if provided
	if len(r.deploymentJson) > 0 {
		if err := Json.Unmarshal(r.deploymentJson, &deployment); err != nil {
			return nil, fmt.Errorf("Could not load deployment JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.DeploymentDescription != nil {
		deployment.description = r.DeploymentDescription
	}

	if r.DeploymentName != nil {
		deployment.name = r.DeploymentName
	}

	if r.DeploymentServerTagScope != nil {
		deployment.serverTagScope = r.DeploymentServerTagScope
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["deployment"] = deployment

	return c.CloneDeployment(r.id, options)
}

type CreateDeploymentRunner struct {
	deploymentDescription    *string
	deploymentName           string
	deploymentServerTagScope *string
}

func (r *CreateDeploymentRunner) Run(c *Client) (interface{}, error) {

	/** Handle deployment parameter **/
	var deployment DeploymentParam2
	// Load JSON if provided
	if len(r.deploymentJson) > 0 {
		if err := Json.Unmarshal(r.deploymentJson, &deployment); err != nil {
			return nil, fmt.Errorf("Could not load deployment JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.DeploymentDescription != nil {
		deployment.description = r.DeploymentDescription
	}

	if len(r.DeploymentName) > 0 {
		deployment.name = r.DeploymentName
	}

	if r.DeploymentServerTagScope != nil {
		deployment.serverTagScope = r.DeploymentServerTagScope
	}

	return c.CreateDeployment(&deployment)
}

type DestroyDeploymentRunner struct {
	id string
}

func (r *DestroyDeploymentRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyDeployment(r.id)
}

type IndexDeploymentsRunner struct {
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexDeploymentsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexDeployments(options)
}

type LockDeploymentRunner struct {
	id string
}

func (r *LockDeploymentRunner) Run(c *Client) (interface{}, error) {
	return c.LockDeployment(r.id)
}

type ServersDeploymentRunner struct {
	id string
}

func (r *ServersDeploymentRunner) Run(c *Client) (interface{}, error) {
	return c.ServersDeployment(r.id)
}

type ShowDeploymentRunner struct {
	id   string
	view *string
}

func (r *ShowDeploymentRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowDeployment(r.id, options)
}

type UnlockDeploymentRunner struct {
	id string
}

func (r *UnlockDeploymentRunner) Run(c *Client) (interface{}, error) {
	return c.UnlockDeployment(r.id)
}

type UpdateDeploymentRunner struct {
	deploymentDescription    *string
	deploymentName           *string
	deploymentServerTagScope *string
	id                       string
}

func (r *UpdateDeploymentRunner) Run(c *Client) (interface{}, error) {

	/** Handle deployment parameter **/
	var deployment DeploymentParam
	// Load JSON if provided
	if len(r.deploymentJson) > 0 {
		if err := Json.Unmarshal(r.deploymentJson, &deployment); err != nil {
			return nil, fmt.Errorf("Could not load deployment JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.DeploymentDescription != nil {
		deployment.description = r.DeploymentDescription
	}

	if r.DeploymentName != nil {
		deployment.name = r.DeploymentName
	}

	if r.DeploymentServerTagScope != nil {
		deployment.serverTagScope = r.DeploymentServerTagScope
	}

	return c.UpdateDeployment(&deployment, r.id)
}

// Register all Deployment actions
func registerDeploymentCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Deployment", `Deployments represent logical groupings of related assets such as servers, server arrays, default configuration settings...etc.`)

	CloneDeploymentRunner := new(CloneDeploymentDeploymentRunner)
	CloneDeploymentCmd := resCmd.Command("CloneDeployment", `Clones a given deployment.`)
	CloneDeploymentRunner.Flag(`deployment.description`, `The description for the cloned deployment.`).StringVar(CloneDeploymentRunner.deploymentDescription)
	CloneDeploymentRunner.Flag(`deployment.name`, `The name for the cloned deployment.`).StringVar(CloneDeploymentRunner.deploymentName)
	CloneDeploymentRunner.Flag(`deployment.serverTagScope`, `The routing scope for tags for servers in the cloned deployment.`).StringVar(CloneDeploymentRunner.deploymentServerTagScope)
	CloneDeploymentRunner.Flag(`id`, ``).Required().StringVar(&CloneDeploymentRunner.id)
	registry[CloneDeploymentCmd.FullCommand()] = CloneDeploymentRunner

	CreateDeploymentRunner := new(CreateDeploymentDeploymentRunner)
	CreateDeploymentCmd := resCmd.Command("CreateDeployment", `Creates a new deployment with the given parameters.`)
	CreateDeploymentRunner.Flag(`deployment.description`, `The description of the deployment to be created.`).StringVar(CreateDeploymentRunner.deploymentDescription)
	CreateDeploymentRunner.Flag(`deployment.name`, `The name of the deployment to be created.`).Required().StringVar(&CreateDeploymentRunner.deploymentName)
	CreateDeploymentRunner.Flag(`deployment.serverTagScope`, `The routing scope for tags for servers in the deployment.`).StringVar(CreateDeploymentRunner.deploymentServerTagScope)
	registry[CreateDeploymentCmd.FullCommand()] = CreateDeploymentRunner

	DestroyDeploymentRunner := new(DestroyDeploymentDeploymentRunner)
	DestroyDeploymentCmd := resCmd.Command("DestroyDeployment", `Deletes a given deployment.`)
	DestroyDeploymentRunner.Flag(`id`, ``).Required().StringVar(&DestroyDeploymentRunner.id)
	registry[DestroyDeploymentCmd.FullCommand()] = DestroyDeploymentRunner

	IndexDeploymentsRunner := new(IndexDeploymentsDeploymentRunner)
	IndexDeploymentsCmd := resCmd.Command("IndexDeployments", `Lists deployments of the account.`)
	IndexDeploymentsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexDeploymentsRunner.filterPos).StringsVar(IndexDeploymentsRunner.filter)
	IndexDeploymentsRunner.Flag(`view`, ``).StringVar(IndexDeploymentsRunner.view)
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
	ShowDeploymentRunner.Flag(`view`, ``).StringVar(ShowDeploymentRunner.view)
	registry[ShowDeploymentCmd.FullCommand()] = ShowDeploymentRunner

	UnlockDeploymentRunner := new(UnlockDeploymentDeploymentRunner)
	UnlockDeploymentCmd := resCmd.Command("UnlockDeployment", `Unlocks a given deployment. Idempotent.`)
	UnlockDeploymentRunner.Flag(`id`, ``).Required().StringVar(&UnlockDeploymentRunner.id)
	registry[UnlockDeploymentCmd.FullCommand()] = UnlockDeploymentRunner

	UpdateDeploymentRunner := new(UpdateDeploymentDeploymentRunner)
	UpdateDeploymentCmd := resCmd.Command("UpdateDeployment", `Updates attributes of a given deployment.`)
	UpdateDeploymentRunner.Flag(`deployment.description`, `The updated description for the deployment.`).StringVar(UpdateDeploymentRunner.deploymentDescription)
	UpdateDeploymentRunner.Flag(`deployment.name`, `The updated name for the deployment.`).StringVar(UpdateDeploymentRunner.deploymentName)
	UpdateDeploymentRunner.Flag(`deployment.serverTagScope`, `The routing scope for tags for servers in the deployment.`).StringVar(UpdateDeploymentRunner.deploymentServerTagScope)
	UpdateDeploymentRunner.Flag(`id`, ``).Required().StringVar(&UpdateDeploymentRunner.id)
	registry[UpdateDeploymentCmd.FullCommand()] = UpdateDeploymentRunner
}

/****** HealthCheck ******/

type IndexHealthCheckRunner struct {
}

func (r *IndexHealthCheckRunner) Run(c *Client) (interface{}, error) {
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

type IndexIdentityProvidersRunner struct {
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexIdentityProvidersRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexIdentityProviders(options)
}

type ShowIdentityProviderRunner struct {
	id   string
	view *string
}

func (r *ShowIdentityProviderRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowIdentityProvider(r.id, options)
}

// Register all IdentityProvider actions
func registerIdentityProviderCmds(app *kingpin.Application) {
	resCmd := app.Cmd("IdentityProvider", `An Identity Provider represents a SAML identity provider (IdP) that is linked to your RightScale Enterprise account, and is trusted by the RightScale dashboard to authenticate your enterprise's end users...`)

	IndexIdentityProvidersRunner := new(IndexIdentityProvidersIdentityProviderRunner)
	IndexIdentityProvidersCmd := resCmd.Command("IndexIdentityProviders", `Lists the identity providers associated with this enterprise account.`)
	IndexIdentityProvidersRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexIdentityProvidersRunner.filterPos).StringsVar(IndexIdentityProvidersRunner.filter)
	IndexIdentityProvidersRunner.Flag(`view`, ``).StringVar(IndexIdentityProvidersRunner.view)
	registry[IndexIdentityProvidersCmd.FullCommand()] = IndexIdentityProvidersRunner

	ShowIdentityProviderRunner := new(ShowIdentityProviderIdentityProviderRunner)
	ShowIdentityProviderCmd := resCmd.Command("ShowIdentityProvider", `Show the specified identity provider, if associated with this enterprise account.`)
	ShowIdentityProviderRunner.Flag(`id`, ``).Required().StringVar(&ShowIdentityProviderRunner.id)
	ShowIdentityProviderRunner.Flag(`view`, ``).StringVar(ShowIdentityProviderRunner.view)
	registry[ShowIdentityProviderCmd.FullCommand()] = ShowIdentityProviderRunner
}

/****** Image ******/

type IndexImagesRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexImagesRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexImages(r.cloudId, options)
}

type ShowImageRunner struct {
	cloudId string
	id      string
	view    *string
}

func (r *ShowImageRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowImage(r.cloudId, r.id, options)
}

// Register all Image actions
func registerImageCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Image", `Images represent base VM image existing in a cloud`)

	IndexImagesRunner := new(IndexImagesImageRunner)
	IndexImagesCmd := resCmd.Command("IndexImages", `Lists all Images for the given Cloud.`)
	IndexImagesRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexImagesRunner.cloudId)
	IndexImagesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexImagesRunner.filterPos).StringsVar(IndexImagesRunner.filter)
	IndexImagesRunner.Flag(`view`, ``).StringVar(IndexImagesRunner.view)
	registry[IndexImagesCmd.FullCommand()] = IndexImagesRunner

	ShowImageRunner := new(ShowImageImageRunner)
	ShowImageCmd := resCmd.Command("ShowImage", `Shows information about a single Image.`)
	ShowImageRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowImageRunner.cloudId)
	ShowImageRunner.Flag(`id`, ``).Required().StringVar(&ShowImageRunner.id)
	ShowImageRunner.Flag(`view`, ``).StringVar(ShowImageRunner.view)
	registry[ShowImageCmd.FullCommand()] = ShowImageRunner
}

/****** Input ******/

type IndexInputsRunner struct {
	cloudId    string
	instanceId string
	view       *string
}

func (r *IndexInputsRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexInputs(r.cloudId, r.instanceId, options)
}

type MultiUpdateInputsRunner struct {
	cloudId      string
	inputsValues []string
	inputsNames  []string
	instanceId   string
}

func (r *MultiUpdateInputsRunner) Run(c *Client) (interface{}, error) {

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
	IndexInputsRunner.Flag(`view`, ``).StringVar(IndexInputsRunner.view)
	registry[IndexInputsCmd.FullCommand()] = IndexInputsRunner

	MultiUpdateInputsRunner := new(MultiUpdateInputsInputRunner)
	MultiUpdateInputsCmd := resCmd.Command("MultiUpdateInputs", `Performs a bulk update of inputs on the specified resource.`)
	MultiUpdateInputsRunner.Flag(`cloudId`, ``).Required().StringVar(&MultiUpdateInputsRunner.cloudId)
	MultiUpdateInputsRunner.FlagPattern(`inputs\.([a-z0-9_]+)`, ``).Required().Capture(&MultiUpdateInputsRunner.inputsNames).StringVar(&MultiUpdateInputsRunner.inputsValues)
	MultiUpdateInputsRunner.Flag(`instanceId`, ``).Required().StringVar(&MultiUpdateInputsRunner.instanceId)
	registry[MultiUpdateInputsCmd.FullCommand()] = MultiUpdateInputsRunner
}

/****** Instance ******/

type CreateInstanceRunner struct {
	cloudId                                                      string
	instanceAssociatePublicIpAddress                             *string
	instanceCloudSpecificAttributesAutomaticInstanceStoreMapping *string
	instanceCloudSpecificAttributesEbsOptimized                  *string
	instanceCloudSpecificAttributesIamInstanceProfile            *string
	instanceCloudSpecificAttributesRootVolumePerformance         *string
	instanceCloudSpecificAttributesRootVolumeSize                *string
	instanceCloudSpecificAttributesRootVolumeTypeUid             *string
	instanceDatacenterHref                                       *string
	instanceDeploymentHref                                       *string
	instanceImageHref                                            string
	instanceInstanceTypeHref                                     string
	instanceKernelImageHref                                      *string
	instanceName                                                 string
	instancePlacementGroupHref                                   *string
	instanceRamdiskImageHref                                     *string
	instanceSecurityGroupHrefs                                   []string
	instanceSecurityGroupHrefsPos                                []string
	instanceSshKeyHref                                           *string
	instanceSubnetHrefs                                          []string
	instanceSubnetHrefsPos                                       []string
	instanceUserData                                             *string
}

func (r *CreateInstanceRunner) Run(c *Client) (interface{}, error) {

	/** Handle instance parameter **/
	var instance InstanceParam
	// Load JSON if provided
	if len(r.instanceJson) > 0 {
		if err := Json.Unmarshal(r.instanceJson, &instance); err != nil {
			return nil, fmt.Errorf("Could not load instance JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.InstanceAssociatePublicIpAddress != nil {
		instance.associatePublicIpAddress = r.InstanceAssociatePublicIpAddress
	}

	if r.InstanceCloudSpecificAttributesAutomaticInstanceStoreMapping != nil {
		instance.cloudSpecificAttributes.automaticInstanceStoreMapping = r.InstanceCloudSpecificAttributesAutomaticInstanceStoreMapping
	}

	if r.InstanceCloudSpecificAttributesEbsOptimized != nil {
		instance.cloudSpecificAttributes.ebsOptimized = r.InstanceCloudSpecificAttributesEbsOptimized
	}

	if r.InstanceCloudSpecificAttributesIamInstanceProfile != nil {
		instance.cloudSpecificAttributes.iamInstanceProfile = r.InstanceCloudSpecificAttributesIamInstanceProfile
	}

	if r.InstanceCloudSpecificAttributesRootVolumePerformance != nil {
		instance.cloudSpecificAttributes.rootVolumePerformance = r.InstanceCloudSpecificAttributesRootVolumePerformance
	}

	if r.InstanceCloudSpecificAttributesRootVolumeSize != nil {
		instance.cloudSpecificAttributes.rootVolumeSize = r.InstanceCloudSpecificAttributesRootVolumeSize
	}

	if r.InstanceCloudSpecificAttributesRootVolumeTypeUid != nil {
		instance.cloudSpecificAttributes.rootVolumeTypeUid = r.InstanceCloudSpecificAttributesRootVolumeTypeUid
	}

	if r.InstanceDatacenterHref != nil {
		instance.datacenterHref = r.InstanceDatacenterHref
	}

	if r.InstanceDeploymentHref != nil {
		instance.deploymentHref = r.InstanceDeploymentHref
	}

	if len(r.InstanceImageHref) > 0 {
		instance.imageHref = r.InstanceImageHref
	}

	if len(r.InstanceInstanceTypeHref) > 0 {
		instance.instanceTypeHref = r.InstanceInstanceTypeHref
	}

	if r.InstanceKernelImageHref != nil {
		instance.kernelImageHref = r.InstanceKernelImageHref
	}

	if len(r.InstanceName) > 0 {
		instance.name = r.InstanceName
	}

	if r.InstancePlacementGroupHref != nil {
		instance.placementGroupHref = r.InstancePlacementGroupHref
	}

	if r.InstanceRamdiskImageHref != nil {
		instance.ramdiskImageHref = r.InstanceRamdiskImageHref
	}

	if r.InstanceSshKeyHref != nil {
		instance.sshKeyHref = r.InstanceSshKeyHref
	}

	if r.InstanceUserData != nil {
		instance.userData = r.InstanceUserData
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxInstanceSecurityGroupHrefsPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.InstanceSecurityGroupHrefsPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for instance.securityGroupHrefs field of securityGroupHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of instance.securityGroupHrefs field of securityGroupHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxInstanceSecurityGroupHrefsPos {
			maxInstanceSecurityGroupHrefsPos = pos
		}
	}
	if len(r.InstanceSecurityGroupHrefs) != maxInstanceSecurityGroupHrefsPos {
		return nil, fmt.Errof("Invalid flags for instance.securityGroupHrefs field of securityGroupHrefs array, %s were defined but max position is %s",
			len(r.InstanceSecurityGroupHrefs), maxInstanceSecurityGroupHrefsPos)
	}
	if maxInstanceSecurityGroupHrefsPos > -1 {
		instanceSecurityGroupHrefs := make([]*string, maxInstanceSecurityGroupHrefsPos+1)
		for i := 0; i < maxInstanceSecurityGroupHrefsPos+1; i++ {
			instanceSecurityGroupHrefs[i] = r.instanceSecurityGroupHrefs[r.instanceSecurityGroupHrefsPos[i]]
		}
		instance.securityGroupHrefs = instanceSecurityGroupHrefs
	}

	maxInstanceSubnetHrefsPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.InstanceSubnetHrefsPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for instance.subnetHrefs field of subnetHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of instance.subnetHrefs field of subnetHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxInstanceSubnetHrefsPos {
			maxInstanceSubnetHrefsPos = pos
		}
	}
	if len(r.InstanceSubnetHrefs) != maxInstanceSubnetHrefsPos {
		return nil, fmt.Errof("Invalid flags for instance.subnetHrefs field of subnetHrefs array, %s were defined but max position is %s",
			len(r.InstanceSubnetHrefs), maxInstanceSubnetHrefsPos)
	}
	if maxInstanceSubnetHrefsPos > -1 {
		instanceSubnetHrefs := make([]*string, maxInstanceSubnetHrefsPos+1)
		for i := 0; i < maxInstanceSubnetHrefsPos+1; i++ {
			instanceSubnetHrefs[i] = r.instanceSubnetHrefs[r.instanceSubnetHrefsPos[i]]
		}
		instance.subnetHrefs = instanceSubnetHrefs
	}

	return c.CreateInstance(r.cloudId, &instance)
}

type IndexInstancesRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexInstancesRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexInstances(r.cloudId, options)
}

type LaunchInstanceRunner struct {
	apiBehavior  *string
	cloudId      string
	id           string
	inputsValues []string
	inputsNames  []string
}

func (r *LaunchInstanceRunner) Run(c *Client) (interface{}, error) {

	/** Handle inputs parameter **/
	var inputs map[string]string

	for i, n := range r.inputsNames {
		inputs[n] = r.inputsValues[i]
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.apiBehavior != nil {
		options["api_behavior"] = *r.apiBehavior
	}
	options["inputs"] = inputs

	return c.LaunchInstance(r.cloudId, r.id, options)
}

type LockInstanceRunner struct {
	cloudId string
	id      string
}

func (r *LockInstanceRunner) Run(c *Client) (interface{}, error) {
	return c.LockInstance(r.cloudId, r.id)
}

type MultiRunExecutableInstancesRunner struct {
	cloudId         string
	filter          []string
	filterPos       []string
	ignoreLock      *string
	inputsValues    []string
	inputsNames     []string
	recipeName      *string
	rightScriptHref *string
}

func (r *MultiRunExecutableInstancesRunner) Run(c *Client) (interface{}, error) {

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

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.ignoreLock != nil {
		options["ignore_lock"] = *r.ignoreLock
	}
	options["inputs"] = inputs
	if r.recipeName != nil {
		options["recipe_name"] = *r.recipeName
	}
	if r.rightScriptHref != nil {
		options["right_script_href"] = *r.rightScriptHref
	}

	return c.MultiRunExecutableInstances(r.cloudId, options)
}

type MultiTerminateInstancesRunner struct {
	cloudId      string
	filter       []string
	filterPos    []string
	terminateAll *string
}

func (r *MultiTerminateInstancesRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.terminateAll != nil {
		options["terminate_all"] = *r.terminateAll
	}

	return c.MultiTerminateInstances(r.cloudId, options)
}

type RebootInstanceRunner struct {
	cloudId string
	id      string
}

func (r *RebootInstanceRunner) Run(c *Client) (interface{}, error) {
	return c.RebootInstance(r.cloudId, r.id)
}

type RunExecutableInstanceRunner struct {
	cloudId         string
	id              string
	ignoreLock      *string
	inputsValues    []string
	inputsNames     []string
	recipeName      *string
	rightScriptHref *string
}

func (r *RunExecutableInstanceRunner) Run(c *Client) (interface{}, error) {

	/** Handle inputs parameter **/
	var inputs map[string]string

	for i, n := range r.inputsNames {
		inputs[n] = r.inputsValues[i]
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.ignoreLock != nil {
		options["ignore_lock"] = *r.ignoreLock
	}
	options["inputs"] = inputs
	if r.recipeName != nil {
		options["recipe_name"] = *r.recipeName
	}
	if r.rightScriptHref != nil {
		options["right_script_href"] = *r.rightScriptHref
	}

	return c.RunExecutableInstance(r.cloudId, r.id, options)
}

type SetCustomLodgementInstanceRunner struct {
	cloudId          string
	id               string
	quantityName     []string
	quantityNamePos  []string
	quantityValue    []string
	quantityValuePos []string
	timeframe        string
}

func (r *SetCustomLodgementInstanceRunner) Run(c *Client) (interface{}, error) {

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

type ShowInstanceRunner struct {
	cloudId string
	id      string
	view    *string
}

func (r *ShowInstanceRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowInstance(r.cloudId, r.id, options)
}

type StartInstanceRunner struct {
	cloudId string
	id      string
}

func (r *StartInstanceRunner) Run(c *Client) (interface{}, error) {
	return c.StartInstance(r.cloudId, r.id)
}

type StopInstanceRunner struct {
	cloudId string
	id      string
}

func (r *StopInstanceRunner) Run(c *Client) (interface{}, error) {
	return c.StopInstance(r.cloudId, r.id)
}

type TerminateInstanceRunner struct {
	cloudId string
	id      string
}

func (r *TerminateInstanceRunner) Run(c *Client) (interface{}, error) {
	return c.TerminateInstance(r.cloudId, r.id)
}

type UnlockInstanceRunner struct {
	cloudId string
	id      string
}

func (r *UnlockInstanceRunner) Run(c *Client) (interface{}, error) {
	return c.UnlockInstance(r.cloudId, r.id)
}

type UpdateInstanceRunner struct {
	cloudId                                                      string
	id                                                           string
	instanceAssociatePublicIpAddress                             *string
	instanceCloudSpecificAttributesAutomaticInstanceStoreMapping *string
	instanceCloudSpecificAttributesIamInstanceProfile            *string
	instanceCloudSpecificAttributesRootVolumePerformance         *string
	instanceCloudSpecificAttributesRootVolumeSize                *string
	instanceCloudSpecificAttributesRootVolumeTypeUid             *string
	instanceDatacenterHref                                       *string
	instanceDeploymentHref                                       *string
	instanceImageHref                                            *string
	instanceInstanceTypeHref                                     *string
	instanceIpForwardingEnabled                                  *string
	instanceKernelImageHref                                      *string
	instanceMultiCloudImageHref                                  *string
	instanceName                                                 *string
	instanceRamdiskImageHref                                     *string
	instanceSecurityGroupHrefs                                   []string
	instanceSecurityGroupHrefsPos                                []string
	instanceServerTemplateHref                                   *string
	instanceSshKeyHref                                           *string
	instanceSubnetHrefs                                          []string
	instanceSubnetHrefsPos                                       []string
	instanceUserData                                             *string
}

func (r *UpdateInstanceRunner) Run(c *Client) (interface{}, error) {

	/** Handle instance parameter **/
	var instance InstanceParam2
	// Load JSON if provided
	if len(r.instanceJson) > 0 {
		if err := Json.Unmarshal(r.instanceJson, &instance); err != nil {
			return nil, fmt.Errorf("Could not load instance JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.InstanceAssociatePublicIpAddress != nil {
		instance.associatePublicIpAddress = r.InstanceAssociatePublicIpAddress
	}

	if r.InstanceCloudSpecificAttributesAutomaticInstanceStoreMapping != nil {
		instance.cloudSpecificAttributes.automaticInstanceStoreMapping = r.InstanceCloudSpecificAttributesAutomaticInstanceStoreMapping
	}

	if r.InstanceCloudSpecificAttributesIamInstanceProfile != nil {
		instance.cloudSpecificAttributes.iamInstanceProfile = r.InstanceCloudSpecificAttributesIamInstanceProfile
	}

	if r.InstanceCloudSpecificAttributesRootVolumePerformance != nil {
		instance.cloudSpecificAttributes.rootVolumePerformance = r.InstanceCloudSpecificAttributesRootVolumePerformance
	}

	if r.InstanceCloudSpecificAttributesRootVolumeSize != nil {
		instance.cloudSpecificAttributes.rootVolumeSize = r.InstanceCloudSpecificAttributesRootVolumeSize
	}

	if r.InstanceCloudSpecificAttributesRootVolumeTypeUid != nil {
		instance.cloudSpecificAttributes.rootVolumeTypeUid = r.InstanceCloudSpecificAttributesRootVolumeTypeUid
	}

	if r.InstanceDatacenterHref != nil {
		instance.datacenterHref = r.InstanceDatacenterHref
	}

	if r.InstanceDeploymentHref != nil {
		instance.deploymentHref = r.InstanceDeploymentHref
	}

	if r.InstanceImageHref != nil {
		instance.imageHref = r.InstanceImageHref
	}

	if r.InstanceInstanceTypeHref != nil {
		instance.instanceTypeHref = r.InstanceInstanceTypeHref
	}

	if r.InstanceIpForwardingEnabled != nil {
		instance.ipForwardingEnabled = r.InstanceIpForwardingEnabled
	}

	if r.InstanceKernelImageHref != nil {
		instance.kernelImageHref = r.InstanceKernelImageHref
	}

	if r.InstanceMultiCloudImageHref != nil {
		instance.multiCloudImageHref = r.InstanceMultiCloudImageHref
	}

	if r.InstanceName != nil {
		instance.name = r.InstanceName
	}

	if r.InstanceRamdiskImageHref != nil {
		instance.ramdiskImageHref = r.InstanceRamdiskImageHref
	}

	if r.InstanceServerTemplateHref != nil {
		instance.serverTemplateHref = r.InstanceServerTemplateHref
	}

	if r.InstanceSshKeyHref != nil {
		instance.sshKeyHref = r.InstanceSshKeyHref
	}

	if r.InstanceUserData != nil {
		instance.userData = r.InstanceUserData
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxInstanceSecurityGroupHrefsPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.InstanceSecurityGroupHrefsPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for instance.securityGroupHrefs field of securityGroupHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of instance.securityGroupHrefs field of securityGroupHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxInstanceSecurityGroupHrefsPos {
			maxInstanceSecurityGroupHrefsPos = pos
		}
	}
	if len(r.InstanceSecurityGroupHrefs) != maxInstanceSecurityGroupHrefsPos {
		return nil, fmt.Errof("Invalid flags for instance.securityGroupHrefs field of securityGroupHrefs array, %s were defined but max position is %s",
			len(r.InstanceSecurityGroupHrefs), maxInstanceSecurityGroupHrefsPos)
	}
	if maxInstanceSecurityGroupHrefsPos > -1 {
		instanceSecurityGroupHrefs := make([]*string, maxInstanceSecurityGroupHrefsPos+1)
		for i := 0; i < maxInstanceSecurityGroupHrefsPos+1; i++ {
			instanceSecurityGroupHrefs[i] = r.instanceSecurityGroupHrefs[r.instanceSecurityGroupHrefsPos[i]]
		}
		instance.securityGroupHrefs = instanceSecurityGroupHrefs
	}

	maxInstanceSubnetHrefsPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.InstanceSubnetHrefsPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for instance.subnetHrefs field of subnetHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of instance.subnetHrefs field of subnetHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxInstanceSubnetHrefsPos {
			maxInstanceSubnetHrefsPos = pos
		}
	}
	if len(r.InstanceSubnetHrefs) != maxInstanceSubnetHrefsPos {
		return nil, fmt.Errof("Invalid flags for instance.subnetHrefs field of subnetHrefs array, %s were defined but max position is %s",
			len(r.InstanceSubnetHrefs), maxInstanceSubnetHrefsPos)
	}
	if maxInstanceSubnetHrefsPos > -1 {
		instanceSubnetHrefs := make([]*string, maxInstanceSubnetHrefsPos+1)
		for i := 0; i < maxInstanceSubnetHrefsPos+1; i++ {
			instanceSubnetHrefs[i] = r.instanceSubnetHrefs[r.instanceSubnetHrefsPos[i]]
		}
		instance.subnetHrefs = instanceSubnetHrefs
	}

	return c.UpdateInstance(r.cloudId, r.id, &instance)
}

// Register all Instance actions
func registerInstanceCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Instance", `Instances represent an entity that is runnable in the cloud.`)

	CreateInstanceRunner := new(CreateInstanceInstanceRunner)
	CreateInstanceCmd := resCmd.Command("CreateInstance", `Creates and launches a raw instance using the provided parameters.`)
	CreateInstanceRunner.Flag(`cloudId`, ``).Required().StringVar(&CreateInstanceRunner.cloudId)
	CreateInstanceRunner.Flag(`instance.associatePublicIpAddress`, `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`).StringVar(CreateInstanceRunner.instanceAssociatePublicIpAddress)
	CreateInstanceRunner.Flag(`instance.cloudSpecificAttributes.automaticInstanceStoreMapping`, `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`).StringVar(CreateInstanceRunner.instanceCloudSpecificAttributesAutomaticInstanceStoreMapping)
	CreateInstanceRunner.Flag(`instance.cloudSpecificAttributes.ebsOptimized`, `Whether the instance is able to connect to IOPS-enabled volumes.`).StringVar(CreateInstanceRunner.instanceCloudSpecificAttributesEbsOptimized)
	CreateInstanceRunner.Flag(`instance.cloudSpecificAttributes.iamInstanceProfile`, `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`).StringVar(CreateInstanceRunner.instanceCloudSpecificAttributesIamInstanceProfile)
	CreateInstanceRunner.Flag(`instance.cloudSpecificAttributes.rootVolumePerformance`, `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`).StringVar(CreateInstanceRunner.instanceCloudSpecificAttributesRootVolumePerformance)
	CreateInstanceRunner.Flag(`instance.cloudSpecificAttributes.rootVolumeSize`, `The size for root disk. Not supported in all Clouds.`).StringVar(CreateInstanceRunner.instanceCloudSpecificAttributesRootVolumeSize)
	CreateInstanceRunner.Flag(`instance.cloudSpecificAttributes.rootVolumeTypeUid`, `The type of root volume for instance. Only available on clouds supporting root volume type.`).StringVar(CreateInstanceRunner.instanceCloudSpecificAttributesRootVolumeTypeUid)
	CreateInstanceRunner.Flag(`instance.datacenterHref`, `The href of the Datacenter / Zone.`).StringVar(CreateInstanceRunner.instanceDatacenterHref)
	CreateInstanceRunner.Flag(`instance.deploymentHref`, `The href of the deployment to which the Instance will be added.`).StringVar(CreateInstanceRunner.instanceDeploymentHref)
	CreateInstanceRunner.Flag(`instance.imageHref`, `The href of the Image to use.`).Required().StringVar(&CreateInstanceRunner.instanceImageHref)
	CreateInstanceRunner.Flag(`instance.instanceTypeHref`, `The href of the instance type.`).Required().StringVar(&CreateInstanceRunner.instanceInstanceTypeHref)
	CreateInstanceRunner.Flag(`instance.kernelImageHref`, `The href of the kernel image.`).StringVar(CreateInstanceRunner.instanceKernelImageHref)
	CreateInstanceRunner.Flag(`instance.name`, `The name of the instance.`).Required().StringVar(&CreateInstanceRunner.instanceName)
	CreateInstanceRunner.Flag(`instance.placementGroupHref`, `The placement group to launch the instance in. Not supported by all clouds & instance types.`).StringVar(CreateInstanceRunner.instancePlacementGroupHref)
	CreateInstanceRunner.Flag(`instance.ramdiskImageHref`, `The href of the ramdisk image.`).StringVar(CreateInstanceRunner.instanceRamdiskImageHref)
	CreateInstanceRunner.FlagPattern(`instance\.securityGroupHrefs\.(\d+)`, `The hrefs of the security groups.`).Capture(&CreateInstanceRunner.instanceSecurityGroupHrefsPos).StringsVar(CreateInstanceRunner.instanceSecurityGroupHrefs)
	CreateInstanceRunner.Flag(`instance.sshKeyHref`, `The href of the SSH key to use.`).StringVar(CreateInstanceRunner.instanceSshKeyHref)
	CreateInstanceRunner.FlagPattern(`instance\.subnetHrefs\.(\d+)`, `The hrefs of the updated subnets.`).Capture(&CreateInstanceRunner.instanceSubnetHrefsPos).StringsVar(CreateInstanceRunner.instanceSubnetHrefs)
	CreateInstanceRunner.Flag(`instance.userData`, `User data that RightScale automatically passes to your instance at boot time.`).StringVar(CreateInstanceRunner.instanceUserData)
	registry[CreateInstanceCmd.FullCommand()] = CreateInstanceRunner

	IndexInstancesRunner := new(IndexInstancesInstanceRunner)
	IndexInstancesCmd := resCmd.Command("IndexInstances", `Lists instances of a given cloud, server array.`)
	IndexInstancesRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexInstancesRunner.cloudId)
	IndexInstancesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexInstancesRunner.filterPos).StringsVar(IndexInstancesRunner.filter)
	IndexInstancesRunner.Flag(`view`, ``).StringVar(IndexInstancesRunner.view)
	registry[IndexInstancesCmd.FullCommand()] = IndexInstancesRunner

	LaunchInstanceRunner := new(LaunchInstanceInstanceRunner)
	LaunchInstanceCmd := resCmd.Command("LaunchInstance", `Launches an instance using the parameters that this instance has been configured with.`)
	LaunchInstanceRunner.Flag(`apiBehavior`, `When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'`).StringVar(LaunchInstanceRunner.apiBehavior)
	LaunchInstanceRunner.Flag(`cloudId`, ``).Required().StringVar(&LaunchInstanceRunner.cloudId)
	LaunchInstanceRunner.Flag(`id`, ``).Required().StringVar(&LaunchInstanceRunner.id)
	LaunchInstanceRunner.FlagPattern(`inputs\.([a-z0-9_]+)`, ``).Capture(&LaunchInstanceRunner.inputsNames).StringVar(LaunchInstanceRunner.inputsValues)
	registry[LaunchInstanceCmd.FullCommand()] = LaunchInstanceRunner

	LockInstanceRunner := new(LockInstanceInstanceRunner)
	LockInstanceCmd := resCmd.Command("LockInstance", ``)
	LockInstanceRunner.Flag(`cloudId`, ``).Required().StringVar(&LockInstanceRunner.cloudId)
	LockInstanceRunner.Flag(`id`, ``).Required().StringVar(&LockInstanceRunner.id)
	registry[LockInstanceCmd.FullCommand()] = LockInstanceRunner

	MultiRunExecutableInstancesRunner := new(MultiRunExecutableInstancesInstanceRunner)
	MultiRunExecutableInstancesCmd := resCmd.Command("MultiRunExecutableInstances", `Runs a script or a recipe in the running instances.`)
	MultiRunExecutableInstancesRunner.Flag(`cloudId`, ``).Required().StringVar(&MultiRunExecutableInstancesRunner.cloudId)
	MultiRunExecutableInstancesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&MultiRunExecutableInstancesRunner.filterPos).StringsVar(MultiRunExecutableInstancesRunner.filter)
	MultiRunExecutableInstancesRunner.Flag(`ignoreLock`, `Specifies the ability to ignore the lock(s) on the Instance(s).`).StringVar(MultiRunExecutableInstancesRunner.ignoreLock)
	MultiRunExecutableInstancesRunner.FlagPattern(`inputs\.([a-z0-9_]+)`, ``).Capture(&MultiRunExecutableInstancesRunner.inputsNames).StringVar(MultiRunExecutableInstancesRunner.inputsValues)
	MultiRunExecutableInstancesRunner.Flag(`recipeName`, `The name of the recipe to be run.`).StringVar(MultiRunExecutableInstancesRunner.recipeName)
	MultiRunExecutableInstancesRunner.Flag(`rightScriptHref`, `The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.`).StringVar(MultiRunExecutableInstancesRunner.rightScriptHref)
	registry[MultiRunExecutableInstancesCmd.FullCommand()] = MultiRunExecutableInstancesRunner

	MultiTerminateInstancesRunner := new(MultiTerminateInstancesInstanceRunner)
	MultiTerminateInstancesCmd := resCmd.Command("MultiTerminateInstances", `Terminates running instances.`)
	MultiTerminateInstancesRunner.Flag(`cloudId`, ``).Required().StringVar(&MultiTerminateInstancesRunner.cloudId)
	MultiTerminateInstancesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&MultiTerminateInstancesRunner.filterPos).StringsVar(MultiTerminateInstancesRunner.filter)
	MultiTerminateInstancesRunner.Flag(`terminateAll`, `Specifies the ability to terminate all instances.`).StringVar(MultiTerminateInstancesRunner.terminateAll)
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
	RunExecutableInstanceRunner.Flag(`ignoreLock`, `Specifies the ability to ignore the lock on the Instance.`).StringVar(RunExecutableInstanceRunner.ignoreLock)
	RunExecutableInstanceRunner.FlagPattern(`inputs\.([a-z0-9_]+)`, ``).Capture(&RunExecutableInstanceRunner.inputsNames).StringVar(RunExecutableInstanceRunner.inputsValues)
	RunExecutableInstanceRunner.Flag(`recipeName`, `The name of the recipe to run.`).StringVar(RunExecutableInstanceRunner.recipeName)
	RunExecutableInstanceRunner.Flag(`rightScriptHref`, `The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.`).StringVar(RunExecutableInstanceRunner.rightScriptHref)
	registry[RunExecutableInstanceCmd.FullCommand()] = RunExecutableInstanceRunner

	SetCustomLodgementInstanceRunner := new(SetCustomLodgementInstanceInstanceRunner)
	SetCustomLodgementInstanceCmd := resCmd.Command("SetCustomLodgementInstance", `This method is deprecated.  Please use InstanceCustomLodgement.`)
	SetCustomLodgementInstanceRunner.Flag(`cloudId`, ``).Required().StringVar(&SetCustomLodgementInstanceRunner.cloudId)
	SetCustomLodgementInstanceRunner.Flag(`id`, ``).Required().StringVar(&SetCustomLodgementInstanceRunner.id)
	SetCustomLodgementInstanceRunner.FlagPattern(`quantity\.(\d+)\.name`, `The name of the quantity. A customer-specific string, e.g. "MB/s" or "GB/Month".`).Capture(&SetCustomLodgementInstanceRunner.quantityNamePos).StringsVar(SetCustomLodgementInstanceRunner.quantityName)
	SetCustomLodgementInstanceRunner.FlagPattern(`quantity\.(\d+)\.value`, `The value of the quantity. Should be a positive integer.`).Capture(&SetCustomLodgementInstanceRunner.quantityValuePos).StringsVar(SetCustomLodgementInstanceRunner.quantityValue)
	SetCustomLodgementInstanceRunner.Flag(`timeframe`, `The timeframe (either a month or a single day) for which the quantity value is valid (currently for the PDT timezone only).`).Required().StringVar(&SetCustomLodgementInstanceRunner.timeframe)
	registry[SetCustomLodgementInstanceCmd.FullCommand()] = SetCustomLodgementInstanceRunner

	ShowInstanceRunner := new(ShowInstanceInstanceRunner)
	ShowInstanceCmd := resCmd.Command("ShowInstance", `Shows attributes of a single instance.`)
	ShowInstanceRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowInstanceRunner.cloudId)
	ShowInstanceRunner.Flag(`id`, ``).Required().StringVar(&ShowInstanceRunner.id)
	ShowInstanceRunner.Flag(`view`, ``).StringVar(ShowInstanceRunner.view)
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
	UpdateInstanceRunner.Flag(`instance.associatePublicIpAddress`, `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`).StringVar(UpdateInstanceRunner.instanceAssociatePublicIpAddress)
	UpdateInstanceRunner.Flag(`instance.cloudSpecificAttributes.automaticInstanceStoreMapping`, `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`).StringVar(UpdateInstanceRunner.instanceCloudSpecificAttributesAutomaticInstanceStoreMapping)
	UpdateInstanceRunner.Flag(`instance.cloudSpecificAttributes.iamInstanceProfile`, `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`).StringVar(UpdateInstanceRunner.instanceCloudSpecificAttributesIamInstanceProfile)
	UpdateInstanceRunner.Flag(`instance.cloudSpecificAttributes.rootVolumePerformance`, `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`).StringVar(UpdateInstanceRunner.instanceCloudSpecificAttributesRootVolumePerformance)
	UpdateInstanceRunner.Flag(`instance.cloudSpecificAttributes.rootVolumeSize`, `The size for root disk. Not supported in all Clouds.`).StringVar(UpdateInstanceRunner.instanceCloudSpecificAttributesRootVolumeSize)
	UpdateInstanceRunner.Flag(`instance.cloudSpecificAttributes.rootVolumeTypeUid`, `The type of root volume for instance. Only available on clouds supporting root volume type.`).StringVar(UpdateInstanceRunner.instanceCloudSpecificAttributesRootVolumeTypeUid)
	UpdateInstanceRunner.Flag(`instance.datacenterHref`, `The href of the updated Datacenter / Zone for the Instance.`).StringVar(UpdateInstanceRunner.instanceDatacenterHref)
	UpdateInstanceRunner.Flag(`instance.deploymentHref`, `The href of the updated Deployment for the Instance. This is only supported for Instances that are not associated with a Server or ServerArray.`).StringVar(UpdateInstanceRunner.instanceDeploymentHref)
	UpdateInstanceRunner.Flag(`instance.imageHref`, `The href of the updated Image for the Instance.`).StringVar(UpdateInstanceRunner.instanceImageHref)
	UpdateInstanceRunner.Flag(`instance.instanceTypeHref`, `The href of the updated Instance Type.`).StringVar(UpdateInstanceRunner.instanceInstanceTypeHref)
	UpdateInstanceRunner.Flag(`instance.ipForwardingEnabled`, `Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.`).StringVar(UpdateInstanceRunner.instanceIpForwardingEnabled)
	UpdateInstanceRunner.Flag(`instance.kernelImageHref`, `The href of the updated kernel image for the Instance.`).StringVar(UpdateInstanceRunner.instanceKernelImageHref)
	UpdateInstanceRunner.Flag(`instance.multiCloudImageHref`, `The href of the updated MultiCloudImage for the Instance.`).StringVar(UpdateInstanceRunner.instanceMultiCloudImageHref)
	UpdateInstanceRunner.Flag(`instance.name`, `The updated name to give the Instance.`).StringVar(UpdateInstanceRunner.instanceName)
	UpdateInstanceRunner.Flag(`instance.ramdiskImageHref`, `The href of the updated ramdisk image for the Instance.`).StringVar(UpdateInstanceRunner.instanceRamdiskImageHref)
	UpdateInstanceRunner.FlagPattern(`instance\.securityGroupHrefs\.(\d+)`, `The hrefs of the updated security groups.`).Capture(&UpdateInstanceRunner.instanceSecurityGroupHrefsPos).StringsVar(UpdateInstanceRunner.instanceSecurityGroupHrefs)
	UpdateInstanceRunner.Flag(`instance.serverTemplateHref`, `The href of the updated ServerTemplate for the Instance.`).StringVar(UpdateInstanceRunner.instanceServerTemplateHref)
	UpdateInstanceRunner.Flag(`instance.sshKeyHref`, `The href of the updated SSH key for the Instance.`).StringVar(UpdateInstanceRunner.instanceSshKeyHref)
	UpdateInstanceRunner.FlagPattern(`instance\.subnetHrefs\.(\d+)`, `The hrefs of the updated subnets.`).Capture(&UpdateInstanceRunner.instanceSubnetHrefsPos).StringsVar(UpdateInstanceRunner.instanceSubnetHrefs)
	UpdateInstanceRunner.Flag(`instance.userData`, `User data that RightScale automatically passes to your instance at boot time.`).StringVar(UpdateInstanceRunner.instanceUserData)
	registry[UpdateInstanceCmd.FullCommand()] = UpdateInstanceRunner
}

/****** InstanceCustomLodgement ******/

type CreateInstanceCustomLodgementRunner struct {
	cloudId          string
	instanceId       string
	quantityName     []string
	quantityNamePos  []string
	quantityValue    []string
	quantityValuePos []string
	timeframe        string
}

func (r *CreateInstanceCustomLodgementRunner) Run(c *Client) (interface{}, error) {

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

type DestroyInstanceCustomLodgementRunner struct {
	cloudId    string
	id         string
	instanceId string
}

func (r *DestroyInstanceCustomLodgementRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyInstanceCustomLodgement(r.cloudId, r.id, r.instanceId)
}

type IndexInstanceCustomLodgementsRunner struct {
	cloudId    string
	instanceId string
	view       *string
}

func (r *IndexInstanceCustomLodgementsRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexInstanceCustomLodgements(r.cloudId, r.instanceId, options)
}

type ShowInstanceCustomLodgementRunner struct {
	cloudId    string
	id         string
	instanceId string
}

func (r *ShowInstanceCustomLodgementRunner) Run(c *Client) (interface{}, error) {
	return c.ShowInstanceCustomLodgement(r.cloudId, r.id, r.instanceId)
}

type UpdateInstanceCustomLodgementRunner struct {
	cloudId          string
	id               string
	instanceId       string
	quantityName     []string
	quantityNamePos  []string
	quantityValue    []string
	quantityValuePos []string
}

func (r *UpdateInstanceCustomLodgementRunner) Run(c *Client) (interface{}, error) {

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
	CreateInstanceCustomLodgementRunner.FlagPattern(`quantity\.(\d+)\.name`, `The name of the quantity. A customer-specific string, e.g. "MB/s" or "GB/Month".`).Capture(&CreateInstanceCustomLodgementRunner.quantityNamePos).StringsVar(CreateInstanceCustomLodgementRunner.quantityName)
	CreateInstanceCustomLodgementRunner.FlagPattern(`quantity\.(\d+)\.value`, `The value of the quantity. Should be a positive integer.`).Capture(&CreateInstanceCustomLodgementRunner.quantityValuePos).StringsVar(CreateInstanceCustomLodgementRunner.quantityValue)
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
	IndexInstanceCustomLodgementsRunner.Flag(`view`, ``).StringVar(IndexInstanceCustomLodgementsRunner.view)
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
	UpdateInstanceCustomLodgementRunner.FlagPattern(`quantity\.(\d+)\.name`, `The name of the quantity. A customer-specific string, e.g. "MB/s" or "GB/Month".`).Capture(&UpdateInstanceCustomLodgementRunner.quantityNamePos).StringsVar(UpdateInstanceCustomLodgementRunner.quantityName)
	UpdateInstanceCustomLodgementRunner.FlagPattern(`quantity\.(\d+)\.value`, `The value of the quantity. Should be a positive integer.`).Capture(&UpdateInstanceCustomLodgementRunner.quantityValuePos).StringsVar(UpdateInstanceCustomLodgementRunner.quantityValue)
	registry[UpdateInstanceCustomLodgementCmd.FullCommand()] = UpdateInstanceCustomLodgementRunner
}

/****** InstanceType ******/

type IndexInstanceTypesRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexInstanceTypesRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexInstanceTypes(r.cloudId, options)
}

type ShowInstanceTypeRunner struct {
	cloudId string
	id      string
	view    *string
}

func (r *ShowInstanceTypeRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowInstanceType(r.cloudId, r.id, options)
}

// Register all InstanceType actions
func registerInstanceTypeCmds(app *kingpin.Application) {
	resCmd := app.Cmd("InstanceType", ``)

	IndexInstanceTypesRunner := new(IndexInstanceTypesInstanceTypeRunner)
	IndexInstanceTypesCmd := resCmd.Command("IndexInstanceTypes", `Lists instance types.`)
	IndexInstanceTypesRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexInstanceTypesRunner.cloudId)
	IndexInstanceTypesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexInstanceTypesRunner.filterPos).StringsVar(IndexInstanceTypesRunner.filter)
	IndexInstanceTypesRunner.Flag(`view`, ``).StringVar(IndexInstanceTypesRunner.view)
	registry[IndexInstanceTypesCmd.FullCommand()] = IndexInstanceTypesRunner

	ShowInstanceTypeRunner := new(ShowInstanceTypeInstanceTypeRunner)
	ShowInstanceTypeCmd := resCmd.Command("ShowInstanceType", `Displays information about a single Instance type.`)
	ShowInstanceTypeRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowInstanceTypeRunner.cloudId)
	ShowInstanceTypeRunner.Flag(`id`, ``).Required().StringVar(&ShowInstanceTypeRunner.id)
	ShowInstanceTypeRunner.Flag(`view`, ``).StringVar(ShowInstanceTypeRunner.view)
	registry[ShowInstanceTypeCmd.FullCommand()] = ShowInstanceTypeRunner
}

/****** IpAddress ******/

type CreateIpAddressRunner struct {
	cloudId                 string
	ipAddressDeploymentHref *string
	ipAddressDomain         *string
	ipAddressName           string
	ipAddressNetworkHref    *string
}

func (r *CreateIpAddressRunner) Run(c *Client) (interface{}, error) {

	/** Handle ipAddress parameter **/
	var ipAddress IpAddressParam
	// Load JSON if provided
	if len(r.ipAddressJson) > 0 {
		if err := Json.Unmarshal(r.ipAddressJson, &ipAddress); err != nil {
			return nil, fmt.Errorf("Could not load ipAddress JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.IpAddressDeploymentHref != nil {
		ipAddress.deploymentHref = r.IpAddressDeploymentHref
	}

	if r.IpAddressDomain != nil {
		ipAddress.domain = r.IpAddressDomain
	}

	if len(r.IpAddressName) > 0 {
		ipAddress.name = r.IpAddressName
	}

	if r.IpAddressNetworkHref != nil {
		ipAddress.networkHref = r.IpAddressNetworkHref
	}

	return c.CreateIpAddress(r.cloudId, &ipAddress)
}

type DestroyIpAddressRunner struct {
	cloudId string
	id      string
}

func (r *DestroyIpAddressRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyIpAddress(r.cloudId, r.id)
}

type IndexIpAddressesRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
}

func (r *IndexIpAddressesRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter

	return c.IndexIpAddresses(r.cloudId, options)
}

type ShowIpAddressRunner struct {
	cloudId string
	id      string
}

func (r *ShowIpAddressRunner) Run(c *Client) (interface{}, error) {
	return c.ShowIpAddress(r.cloudId, r.id)
}

type UpdateIpAddressRunner struct {
	cloudId                 string
	id                      string
	ipAddressDeploymentHref *string
	ipAddressName           string
}

func (r *UpdateIpAddressRunner) Run(c *Client) (interface{}, error) {

	/** Handle ipAddress parameter **/
	var ipAddress IpAddressParam2
	// Load JSON if provided
	if len(r.ipAddressJson) > 0 {
		if err := Json.Unmarshal(r.ipAddressJson, &ipAddress); err != nil {
			return nil, fmt.Errorf("Could not load ipAddress JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.IpAddressDeploymentHref != nil {
		ipAddress.deploymentHref = r.IpAddressDeploymentHref
	}

	if len(r.IpAddressName) > 0 {
		ipAddress.name = r.IpAddressName
	}

	return c.UpdateIpAddress(r.cloudId, r.id, &ipAddress)
}

// Register all IpAddress actions
func registerIpAddressCmds(app *kingpin.Application) {
	resCmd := app.Cmd("IpAddress", `An IpAddress provides an abstraction for IPv4 addresses bindable to Instance resources running in a Cloud.`)

	CreateIpAddressRunner := new(CreateIpAddressIpAddressRunner)
	CreateIpAddressCmd := resCmd.Command("CreateIpAddress", `Creates a new IpAddress with the given parameters.`)
	CreateIpAddressRunner.Flag(`cloudId`, ``).Required().StringVar(&CreateIpAddressRunner.cloudId)
	CreateIpAddressRunner.Flag(`ipAddress.deploymentHref`, `The href of the Deployment that owns this IpAddress.`).StringVar(CreateIpAddressRunner.ipAddressDeploymentHref)
	CreateIpAddressRunner.Flag(`ipAddress.domain`, `(Amazon Only) Pass vpc to create this IP for EC2-VPC only environments. Pass ec2_classic to create this IP for EC2-Classic environments. Defaults to ec2_classic.`).StringVar(CreateIpAddressRunner.ipAddressDomain)
	CreateIpAddressRunner.Flag(`ipAddress.name`, `The name of the IpAddress to be created.`).Required().StringVar(&CreateIpAddressRunner.ipAddressName)
	CreateIpAddressRunner.Flag(`ipAddress.networkHref`, `(OpenStack Only) The href of the Network that the IpAddress will be associated to. This parameter is required for OpenStack with Neutron clouds.`).StringVar(CreateIpAddressRunner.ipAddressNetworkHref)
	registry[CreateIpAddressCmd.FullCommand()] = CreateIpAddressRunner

	DestroyIpAddressRunner := new(DestroyIpAddressIpAddressRunner)
	DestroyIpAddressCmd := resCmd.Command("DestroyIpAddress", `Deletes a given IpAddress.`)
	DestroyIpAddressRunner.Flag(`cloudId`, ``).Required().StringVar(&DestroyIpAddressRunner.cloudId)
	DestroyIpAddressRunner.Flag(`id`, ``).Required().StringVar(&DestroyIpAddressRunner.id)
	registry[DestroyIpAddressCmd.FullCommand()] = DestroyIpAddressRunner

	IndexIpAddressesRunner := new(IndexIpAddressesIpAddressRunner)
	IndexIpAddressesCmd := resCmd.Command("IndexIpAddresses", `Lists the IpAddresses available to this account.`)
	IndexIpAddressesRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexIpAddressesRunner.cloudId)
	IndexIpAddressesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexIpAddressesRunner.filterPos).StringsVar(IndexIpAddressesRunner.filter)
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
	UpdateIpAddressRunner.Flag(`ipAddress.deploymentHref`, `The href of the Deployment that owns this IpAddress.`).StringVar(UpdateIpAddressRunner.ipAddressDeploymentHref)
	UpdateIpAddressRunner.Flag(`ipAddress.name`, `The updated name of the IpAddress.`).Required().StringVar(&UpdateIpAddressRunner.ipAddressName)
	registry[UpdateIpAddressCmd.FullCommand()] = UpdateIpAddressRunner
}

/****** IpAddressBinding ******/

type CreateIpAddressBindingRunner struct {
	cloudId                             string
	ipAddressBindingInstanceHref        string
	ipAddressBindingPrivatePort         *string
	ipAddressBindingProtocol            *string
	ipAddressBindingPublicIpAddressHref *string
	ipAddressBindingPublicPort          *string
	ipAddressId                         string
}

func (r *CreateIpAddressBindingRunner) Run(c *Client) (interface{}, error) {

	/** Handle ipAddressBinding parameter **/
	var ipAddressBinding IpAddressBindingParam
	// Load JSON if provided
	if len(r.ipAddressBindingJson) > 0 {
		if err := Json.Unmarshal(r.ipAddressBindingJson, &ipAddressBinding); err != nil {
			return nil, fmt.Errorf("Could not load ipAddressBinding JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.IpAddressBindingInstanceHref) > 0 {
		ipAddressBinding.instanceHref = r.IpAddressBindingInstanceHref
	}

	if r.IpAddressBindingPrivatePort != nil {
		ipAddressBinding.privatePort = r.IpAddressBindingPrivatePort
	}

	if r.IpAddressBindingProtocol != nil {
		ipAddressBinding.protocol = r.IpAddressBindingProtocol
	}

	if r.IpAddressBindingPublicIpAddressHref != nil {
		ipAddressBinding.publicIpAddressHref = r.IpAddressBindingPublicIpAddressHref
	}

	if r.IpAddressBindingPublicPort != nil {
		ipAddressBinding.publicPort = r.IpAddressBindingPublicPort
	}

	return c.CreateIpAddressBinding(r.cloudId, &ipAddressBinding, r.ipAddressId)
}

type DestroyIpAddressBindingRunner struct {
	cloudId     string
	id          string
	ipAddressId string
}

func (r *DestroyIpAddressBindingRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyIpAddressBinding(r.cloudId, r.id, r.ipAddressId)
}

type IndexIpAddressBindingsRunner struct {
	cloudId     string
	filter      []string
	filterPos   []string
	ipAddressId string
}

func (r *IndexIpAddressBindingsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter

	return c.IndexIpAddressBindings(r.cloudId, r.ipAddressId, options)
}

type ShowIpAddressBindingRunner struct {
	cloudId     string
	id          string
	ipAddressId string
}

func (r *ShowIpAddressBindingRunner) Run(c *Client) (interface{}, error) {
	return c.ShowIpAddressBinding(r.cloudId, r.id, r.ipAddressId)
}

// Register all IpAddressBinding actions
func registerIpAddressBindingCmds(app *kingpin.Application) {
	resCmd := app.Cmd("IpAddressBinding", `An IpAddressBinding represents an abstraction for binding an IpAddress to an instance.`)

	CreateIpAddressBindingRunner := new(CreateIpAddressBindingIpAddressBindingRunner)
	CreateIpAddressBindingCmd := resCmd.Command("CreateIpAddressBinding", `Creates an ip address binding which attaches a specified IpAddress resource to a specified instance, and also allows for configuration of port forwarding rules...`)
	CreateIpAddressBindingRunner.Flag(`cloudId`, ``).Required().StringVar(&CreateIpAddressBindingRunner.cloudId)
	CreateIpAddressBindingRunner.Flag(`ipAddressBinding.instanceHref`, `The Instance to which this IpAddress should be bound.`).Required().StringVar(&CreateIpAddressBindingRunner.ipAddressBindingInstanceHref)
	CreateIpAddressBindingRunner.Flag(`ipAddressBinding.privatePort`, `Incoming network traffic will get forwarded to this port number on the specified Instance. If not specified, will use public port. Required unless public_ip_address_href is passed.`).StringVar(CreateIpAddressBindingRunner.ipAddressBindingPrivatePort)
	CreateIpAddressBindingRunner.Flag(`ipAddressBinding.protocol`, `Transport layer protocol of traffic that may be forwarded from public port to private port on the Instance. Required unless public_ip_address_href is passed.`).StringVar(CreateIpAddressBindingRunner.ipAddressBindingProtocol)
	CreateIpAddressBindingRunner.Flag(`ipAddressBinding.publicIpAddressHref`, `The IpAddress to bind to the specified instance. Required unless port forwarding rule params are passed.`).StringVar(CreateIpAddressBindingRunner.ipAddressBindingPublicIpAddressHref)
	CreateIpAddressBindingRunner.Flag(`ipAddressBinding.publicPort`, `The incoming port for port forwarding. Incoming network traffic on this port will get forwarded (to the IP:Private Port of the specified Instance). Required unless public_ip_address_href is passed.`).StringVar(CreateIpAddressBindingRunner.ipAddressBindingPublicPort)
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
	IndexIpAddressBindingsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexIpAddressBindingsRunner.filterPos).StringsVar(IndexIpAddressBindingsRunner.filter)
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

type DataMonitoringMetricRunner struct {
	cloudId    string
	end        string
	id         string
	instanceId string
	start      string
}

func (r *DataMonitoringMetricRunner) Run(c *Client) (interface{}, error) {
	return c.DataMonitoringMetric(r.cloudId, r.end, r.id, r.instanceId, r.start)
}

type IndexMonitoringMetricsRunner struct {
	cloudId    string
	filter     []string
	filterPos  []string
	instanceId string
	period     *string
	size       *string
	title      *string
	tz         *string
}

func (r *IndexMonitoringMetricsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.period != nil {
		options["period"] = *r.period
	}
	if r.size != nil {
		options["size"] = *r.size
	}
	if r.title != nil {
		options["title"] = *r.title
	}
	if r.tz != nil {
		options["tz"] = *r.tz
	}

	return c.IndexMonitoringMetrics(r.cloudId, r.instanceId, options)
}

type ShowMonitoringMetricRunner struct {
	cloudId    string
	id         string
	instanceId string
	period     *string
	size       *string
	title      *string
	tz         *string
}

func (r *ShowMonitoringMetricRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.period != nil {
		options["period"] = *r.period
	}
	if r.size != nil {
		options["size"] = *r.size
	}
	if r.title != nil {
		options["title"] = *r.title
	}
	if r.tz != nil {
		options["tz"] = *r.tz
	}

	return c.ShowMonitoringMetric(r.cloudId, r.id, r.instanceId, options)
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
	IndexMonitoringMetricsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexMonitoringMetricsRunner.filterPos).StringsVar(IndexMonitoringMetricsRunner.filter)
	IndexMonitoringMetricsRunner.Flag(`instanceId`, ``).Required().StringVar(&IndexMonitoringMetricsRunner.instanceId)
	IndexMonitoringMetricsRunner.Flag(`period`, `The time scale for which the graph is generated. Default is 'day'`).StringVar(IndexMonitoringMetricsRunner.period)
	IndexMonitoringMetricsRunner.Flag(`size`, `The size of the graph to be generated. Default is 'small'.`).StringVar(IndexMonitoringMetricsRunner.size)
	IndexMonitoringMetricsRunner.Flag(`title`, `The title of the graph.`).StringVar(IndexMonitoringMetricsRunner.title)
	IndexMonitoringMetricsRunner.Flag(`tz`, `The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences. `).StringVar(IndexMonitoringMetricsRunner.tz)
	registry[IndexMonitoringMetricsCmd.FullCommand()] = IndexMonitoringMetricsRunner

	ShowMonitoringMetricRunner := new(ShowMonitoringMetricMonitoringMetricRunner)
	ShowMonitoringMetricCmd := resCmd.Command("ShowMonitoringMetric", `Shows attributes of a single monitoring metric.`)
	ShowMonitoringMetricRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowMonitoringMetricRunner.cloudId)
	ShowMonitoringMetricRunner.Flag(`id`, ``).Required().StringVar(&ShowMonitoringMetricRunner.id)
	ShowMonitoringMetricRunner.Flag(`instanceId`, ``).Required().StringVar(&ShowMonitoringMetricRunner.instanceId)
	ShowMonitoringMetricRunner.Flag(`period`, `The time scale for which the graph is generated. Default is 'day'.`).StringVar(ShowMonitoringMetricRunner.period)
	ShowMonitoringMetricRunner.Flag(`size`, `The size of the graph to be generated. Default is 'small'.`).StringVar(ShowMonitoringMetricRunner.size)
	ShowMonitoringMetricRunner.Flag(`title`, `The title of the graph.`).StringVar(ShowMonitoringMetricRunner.title)
	ShowMonitoringMetricRunner.Flag(`tz`, `The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences. `).StringVar(ShowMonitoringMetricRunner.tz)
	registry[ShowMonitoringMetricCmd.FullCommand()] = ShowMonitoringMetricRunner
}

/****** MultiCloudImage ******/

type CloneMultiCloudImageRunner struct {
	id                         string
	multiCloudImageDescription *string
	multiCloudImageName        string
}

func (r *CloneMultiCloudImageRunner) Run(c *Client) (interface{}, error) {

	/** Handle multiCloudImage parameter **/
	var multiCloudImage MultiCloudImageParam
	// Load JSON if provided
	if len(r.multiCloudImageJson) > 0 {
		if err := Json.Unmarshal(r.multiCloudImageJson, &multiCloudImage); err != nil {
			return nil, fmt.Errorf("Could not load multiCloudImage JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.MultiCloudImageDescription != nil {
		multiCloudImage.description = r.MultiCloudImageDescription
	}

	if len(r.MultiCloudImageName) > 0 {
		multiCloudImage.name = r.MultiCloudImageName
	}

	return c.CloneMultiCloudImage(r.id, &multiCloudImage)
}

type CommitMultiCloudImageRunner struct {
	commitMessage string
	id            string
}

func (r *CommitMultiCloudImageRunner) Run(c *Client) (interface{}, error) {
	return c.CommitMultiCloudImage(r.commitMessage, r.id)
}

type CreateMultiCloudImageRunner struct {
	multiCloudImageDescription *string
	multiCloudImageName        string
	serverTemplateId           string
}

func (r *CreateMultiCloudImageRunner) Run(c *Client) (interface{}, error) {

	/** Handle multiCloudImage parameter **/
	var multiCloudImage MultiCloudImageParam2
	// Load JSON if provided
	if len(r.multiCloudImageJson) > 0 {
		if err := Json.Unmarshal(r.multiCloudImageJson, &multiCloudImage); err != nil {
			return nil, fmt.Errorf("Could not load multiCloudImage JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.MultiCloudImageDescription != nil {
		multiCloudImage.description = r.MultiCloudImageDescription
	}

	if len(r.MultiCloudImageName) > 0 {
		multiCloudImage.name = r.MultiCloudImageName
	}

	return c.CreateMultiCloudImage(&multiCloudImage, r.serverTemplateId)
}

type DestroyMultiCloudImageRunner struct {
	id               string
	serverTemplateId string
}

func (r *DestroyMultiCloudImageRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyMultiCloudImage(r.id, r.serverTemplateId)
}

type IndexMultiCloudImagesRunner struct {
	filter           []string
	filterPos        []string
	serverTemplateId string
}

func (r *IndexMultiCloudImagesRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter

	return c.IndexMultiCloudImages(r.serverTemplateId, options)
}

type ShowMultiCloudImageRunner struct {
	id               string
	serverTemplateId string
}

func (r *ShowMultiCloudImageRunner) Run(c *Client) (interface{}, error) {
	return c.ShowMultiCloudImage(r.id, r.serverTemplateId)
}

type UpdateMultiCloudImageRunner struct {
	id                         string
	multiCloudImageDescription *string
	multiCloudImageName        *string
	serverTemplateId           string
}

func (r *UpdateMultiCloudImageRunner) Run(c *Client) (interface{}, error) {

	/** Handle multiCloudImage parameter **/
	var multiCloudImage MultiCloudImageParam
	// Load JSON if provided
	if len(r.multiCloudImageJson) > 0 {
		if err := Json.Unmarshal(r.multiCloudImageJson, &multiCloudImage); err != nil {
			return nil, fmt.Errorf("Could not load multiCloudImage JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.MultiCloudImageDescription != nil {
		multiCloudImage.description = r.MultiCloudImageDescription
	}

	if r.MultiCloudImageName != nil {
		multiCloudImage.name = r.MultiCloudImageName
	}

	return c.UpdateMultiCloudImage(r.id, &multiCloudImage, r.serverTemplateId)
}

// Register all MultiCloudImage actions
func registerMultiCloudImageCmds(app *kingpin.Application) {
	resCmd := app.Cmd("MultiCloudImage", `A MultiCloudImage is a RightScale component that functions as a pointer to machine images in specific clouds (e...`)

	CloneMultiCloudImageRunner := new(CloneMultiCloudImageMultiCloudImageRunner)
	CloneMultiCloudImageCmd := resCmd.Command("CloneMultiCloudImage", `Clones a given MultiCloudImage.`)
	CloneMultiCloudImageRunner.Flag(`id`, ``).Required().StringVar(&CloneMultiCloudImageRunner.id)
	CloneMultiCloudImageRunner.Flag(`multiCloudImage.description`, `The description for the cloned MultiCloudImage.`).StringVar(CloneMultiCloudImageRunner.multiCloudImageDescription)
	CloneMultiCloudImageRunner.Flag(`multiCloudImage.name`, `The name for the cloned MultiCloudImage.`).Required().StringVar(&CloneMultiCloudImageRunner.multiCloudImageName)
	registry[CloneMultiCloudImageCmd.FullCommand()] = CloneMultiCloudImageRunner

	CommitMultiCloudImageRunner := new(CommitMultiCloudImageMultiCloudImageRunner)
	CommitMultiCloudImageCmd := resCmd.Command("CommitMultiCloudImage", `Commits a given MultiCloudImage. Only HEAD revisions can be committed.`)
	CommitMultiCloudImageRunner.Flag(`commitMessage`, `The message associated with the commit.`).Required().StringVar(&CommitMultiCloudImageRunner.commitMessage)
	CommitMultiCloudImageRunner.Flag(`id`, ``).Required().StringVar(&CommitMultiCloudImageRunner.id)
	registry[CommitMultiCloudImageCmd.FullCommand()] = CommitMultiCloudImageRunner

	CreateMultiCloudImageRunner := new(CreateMultiCloudImageMultiCloudImageRunner)
	CreateMultiCloudImageCmd := resCmd.Command("CreateMultiCloudImage", `Creates a new MultiCloudImage with the given parameters.`)
	CreateMultiCloudImageRunner.Flag(`multiCloudImage.description`, `The description of the MultiCloudImage to be created.`).StringVar(CreateMultiCloudImageRunner.multiCloudImageDescription)
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
	IndexMultiCloudImagesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexMultiCloudImagesRunner.filterPos).StringsVar(IndexMultiCloudImagesRunner.filter)
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
	UpdateMultiCloudImageRunner.Flag(`multiCloudImage.description`, `The updated description for the MultiCloudImage.`).StringVar(UpdateMultiCloudImageRunner.multiCloudImageDescription)
	UpdateMultiCloudImageRunner.Flag(`multiCloudImage.name`, `The updated name for the MultiCloudImage.`).StringVar(UpdateMultiCloudImageRunner.multiCloudImageName)
	UpdateMultiCloudImageRunner.Flag(`serverTemplateId`, ``).Required().StringVar(&UpdateMultiCloudImageRunner.serverTemplateId)
	registry[UpdateMultiCloudImageCmd.FullCommand()] = UpdateMultiCloudImageRunner
}

/****** MultiCloudImageSetting ******/

type CreateMultiCloudImageSettingRunner struct {
	multiCloudImageId                      string
	multiCloudImageSettingCloudHref        *string
	multiCloudImageSettingImageHref        *string
	multiCloudImageSettingInstanceTypeHref *string
	multiCloudImageSettingKernelImageHref  *string
	multiCloudImageSettingRamdiskImageHref *string
	multiCloudImageSettingUserData         *string
}

func (r *CreateMultiCloudImageSettingRunner) Run(c *Client) (interface{}, error) {

	/** Handle multiCloudImageSetting parameter **/
	var multiCloudImageSetting MultiCloudImageSettingParam
	// Load JSON if provided
	if len(r.multiCloudImageSettingJson) > 0 {
		if err := Json.Unmarshal(r.multiCloudImageSettingJson, &multiCloudImageSetting); err != nil {
			return nil, fmt.Errorf("Could not load multiCloudImageSetting JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.MultiCloudImageSettingCloudHref != nil {
		multiCloudImageSetting.cloudHref = r.MultiCloudImageSettingCloudHref
	}

	if r.MultiCloudImageSettingImageHref != nil {
		multiCloudImageSetting.imageHref = r.MultiCloudImageSettingImageHref
	}

	if r.MultiCloudImageSettingInstanceTypeHref != nil {
		multiCloudImageSetting.instanceTypeHref = r.MultiCloudImageSettingInstanceTypeHref
	}

	if r.MultiCloudImageSettingKernelImageHref != nil {
		multiCloudImageSetting.kernelImageHref = r.MultiCloudImageSettingKernelImageHref
	}

	if r.MultiCloudImageSettingRamdiskImageHref != nil {
		multiCloudImageSetting.ramdiskImageHref = r.MultiCloudImageSettingRamdiskImageHref
	}

	if r.MultiCloudImageSettingUserData != nil {
		multiCloudImageSetting.userData = r.MultiCloudImageSettingUserData
	}

	return c.CreateMultiCloudImageSetting(r.multiCloudImageId, &multiCloudImageSetting)
}

type DestroyMultiCloudImageSettingRunner struct {
	id                string
	multiCloudImageId string
}

func (r *DestroyMultiCloudImageSettingRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyMultiCloudImageSetting(r.id, r.multiCloudImageId)
}

type IndexMultiCloudImageSettingsRunner struct {
	filter            []string
	filterPos         []string
	multiCloudImageId string
}

func (r *IndexMultiCloudImageSettingsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter

	return c.IndexMultiCloudImageSettings(r.multiCloudImageId, options)
}

type ShowMultiCloudImageSettingRunner struct {
	id                string
	multiCloudImageId string
}

func (r *ShowMultiCloudImageSettingRunner) Run(c *Client) (interface{}, error) {
	return c.ShowMultiCloudImageSetting(r.id, r.multiCloudImageId)
}

type UpdateMultiCloudImageSettingRunner struct {
	id                                     string
	multiCloudImageId                      string
	multiCloudImageSettingCloudHref        *string
	multiCloudImageSettingImageHref        *string
	multiCloudImageSettingInstanceTypeHref *string
	multiCloudImageSettingKernelImageHref  *string
	multiCloudImageSettingRamdiskImageHref *string
	multiCloudImageSettingUserData         *string
}

func (r *UpdateMultiCloudImageSettingRunner) Run(c *Client) (interface{}, error) {

	/** Handle multiCloudImageSetting parameter **/
	var multiCloudImageSetting MultiCloudImageSettingParam2
	// Load JSON if provided
	if len(r.multiCloudImageSettingJson) > 0 {
		if err := Json.Unmarshal(r.multiCloudImageSettingJson, &multiCloudImageSetting); err != nil {
			return nil, fmt.Errorf("Could not load multiCloudImageSetting JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.MultiCloudImageSettingCloudHref != nil {
		multiCloudImageSetting.cloudHref = r.MultiCloudImageSettingCloudHref
	}

	if r.MultiCloudImageSettingImageHref != nil {
		multiCloudImageSetting.imageHref = r.MultiCloudImageSettingImageHref
	}

	if r.MultiCloudImageSettingInstanceTypeHref != nil {
		multiCloudImageSetting.instanceTypeHref = r.MultiCloudImageSettingInstanceTypeHref
	}

	if r.MultiCloudImageSettingKernelImageHref != nil {
		multiCloudImageSetting.kernelImageHref = r.MultiCloudImageSettingKernelImageHref
	}

	if r.MultiCloudImageSettingRamdiskImageHref != nil {
		multiCloudImageSetting.ramdiskImageHref = r.MultiCloudImageSettingRamdiskImageHref
	}

	if r.MultiCloudImageSettingUserData != nil {
		multiCloudImageSetting.userData = r.MultiCloudImageSettingUserData
	}

	return c.UpdateMultiCloudImageSetting(r.id, r.multiCloudImageId, &multiCloudImageSetting)
}

// Register all MultiCloudImageSetting actions
func registerMultiCloudImageSettingCmds(app *kingpin.Application) {
	resCmd := app.Cmd("MultiCloudImageSetting", `A MultiCloudImageSetting defines which settings should be used when a server is launched in a cloud...`)

	CreateMultiCloudImageSettingRunner := new(CreateMultiCloudImageSettingMultiCloudImageSettingRunner)
	CreateMultiCloudImageSettingCmd := resCmd.Command("CreateMultiCloudImageSetting", `Creates a new setting for an existing MultiCloudImage.`)
	CreateMultiCloudImageSettingRunner.Flag(`multiCloudImageId`, ``).Required().StringVar(&CreateMultiCloudImageSettingRunner.multiCloudImageId)
	CreateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.cloudHref`, `The href of the Cloud to use.`).StringVar(CreateMultiCloudImageSettingRunner.multiCloudImageSettingCloudHref)
	CreateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.imageHref`, `The href of the Image to use. Mandatory if specifying cloud_href.`).StringVar(CreateMultiCloudImageSettingRunner.multiCloudImageSettingImageHref)
	CreateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.instanceTypeHref`, `The href of the instance type. Mandatory if specifying cloud_href.`).StringVar(CreateMultiCloudImageSettingRunner.multiCloudImageSettingInstanceTypeHref)
	CreateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.kernelImageHref`, `The href of the kernel image. Optional if specifying cloud_href.`).StringVar(CreateMultiCloudImageSettingRunner.multiCloudImageSettingKernelImageHref)
	CreateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.ramdiskImageHref`, `The href of the ramdisk image. Optional if specifying cloud_href.`).StringVar(CreateMultiCloudImageSettingRunner.multiCloudImageSettingRamdiskImageHref)
	CreateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.userData`, `User data that RightScale automatically passes to your instance at boot time.`).StringVar(CreateMultiCloudImageSettingRunner.multiCloudImageSettingUserData)
	registry[CreateMultiCloudImageSettingCmd.FullCommand()] = CreateMultiCloudImageSettingRunner

	DestroyMultiCloudImageSettingRunner := new(DestroyMultiCloudImageSettingMultiCloudImageSettingRunner)
	DestroyMultiCloudImageSettingCmd := resCmd.Command("DestroyMultiCloudImageSetting", `Deletes a MultiCloudImage setting.`)
	DestroyMultiCloudImageSettingRunner.Flag(`id`, ``).Required().StringVar(&DestroyMultiCloudImageSettingRunner.id)
	DestroyMultiCloudImageSettingRunner.Flag(`multiCloudImageId`, ``).Required().StringVar(&DestroyMultiCloudImageSettingRunner.multiCloudImageId)
	registry[DestroyMultiCloudImageSettingCmd.FullCommand()] = DestroyMultiCloudImageSettingRunner

	IndexMultiCloudImageSettingsRunner := new(IndexMultiCloudImageSettingsMultiCloudImageSettingRunner)
	IndexMultiCloudImageSettingsCmd := resCmd.Command("IndexMultiCloudImageSettings", `Lists the MultiCloudImage settings.`)
	IndexMultiCloudImageSettingsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexMultiCloudImageSettingsRunner.filterPos).StringsVar(IndexMultiCloudImageSettingsRunner.filter)
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
	UpdateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.cloudHref`, `The href of the Cloud to use.`).StringVar(UpdateMultiCloudImageSettingRunner.multiCloudImageSettingCloudHref)
	UpdateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.imageHref`, `The href of the Image to use.`).StringVar(UpdateMultiCloudImageSettingRunner.multiCloudImageSettingImageHref)
	UpdateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.instanceTypeHref`, `The href of the instance type.`).StringVar(UpdateMultiCloudImageSettingRunner.multiCloudImageSettingInstanceTypeHref)
	UpdateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.kernelImageHref`, `The href of the kernel image.`).StringVar(UpdateMultiCloudImageSettingRunner.multiCloudImageSettingKernelImageHref)
	UpdateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.ramdiskImageHref`, `The href of the ramdisk image.`).StringVar(UpdateMultiCloudImageSettingRunner.multiCloudImageSettingRamdiskImageHref)
	UpdateMultiCloudImageSettingRunner.Flag(`multiCloudImageSetting.userData`, `User data that RightScale automatically passes to your instance at boot time.`).StringVar(UpdateMultiCloudImageSettingRunner.multiCloudImageSettingUserData)
	registry[UpdateMultiCloudImageSettingCmd.FullCommand()] = UpdateMultiCloudImageSettingRunner
}

/****** Network ******/

type CreateNetworkRunner struct {
	networkCidrBlock       *string
	networkCloudHref       string
	networkDescription     *string
	networkInstanceTenancy *string
	networkName            *string
}

func (r *CreateNetworkRunner) Run(c *Client) (interface{}, error) {

	/** Handle network parameter **/
	var network NetworkParam
	// Load JSON if provided
	if len(r.networkJson) > 0 {
		if err := Json.Unmarshal(r.networkJson, &network); err != nil {
			return nil, fmt.Errorf("Could not load network JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.NetworkCidrBlock != nil {
		network.cidrBlock = r.NetworkCidrBlock
	}

	if len(r.NetworkCloudHref) > 0 {
		network.cloudHref = r.NetworkCloudHref
	}

	if r.NetworkDescription != nil {
		network.description = r.NetworkDescription
	}

	if r.NetworkInstanceTenancy != nil {
		network.instanceTenancy = r.NetworkInstanceTenancy
	}

	if r.NetworkName != nil {
		network.name = r.NetworkName
	}

	return c.CreateNetwork(&network)
}

type DestroyNetworkRunner struct {
	id string
}

func (r *DestroyNetworkRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyNetwork(r.id)
}

type IndexNetworksRunner struct {
	filter    []string
	filterPos []string
}

func (r *IndexNetworksRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter

	return c.IndexNetworks(options)
}

type ShowNetworkRunner struct {
	id string
}

func (r *ShowNetworkRunner) Run(c *Client) (interface{}, error) {
	return c.ShowNetwork(r.id)
}

type UpdateNetworkRunner struct {
	id                    string
	networkDescription    *string
	networkName           *string
	networkRouteTableHref *string
}

func (r *UpdateNetworkRunner) Run(c *Client) (interface{}, error) {

	/** Handle network parameter **/
	var network NetworkParam2
	// Load JSON if provided
	if len(r.networkJson) > 0 {
		if err := Json.Unmarshal(r.networkJson, &network); err != nil {
			return nil, fmt.Errorf("Could not load network JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.NetworkDescription != nil {
		network.description = r.NetworkDescription
	}

	if r.NetworkName != nil {
		network.name = r.NetworkName
	}

	if r.NetworkRouteTableHref != nil {
		network.routeTableHref = r.NetworkRouteTableHref
	}

	return c.UpdateNetwork(r.id, &network)
}

// Register all Network actions
func registerNetworkCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Network", `A Network is a logical grouping of network devices.`)

	CreateNetworkRunner := new(CreateNetworkNetworkRunner)
	CreateNetworkCmd := resCmd.Command("CreateNetwork", `Creates a new network.`)
	CreateNetworkRunner.Flag(`network.cidrBlock`, `The range of IP addresses for the Network. This parameter is required for Amazon clouds.`).StringVar(CreateNetworkRunner.networkCidrBlock)
	CreateNetworkRunner.Flag(`network.cloudHref`, `The Cloud to create the Network in`).Required().StringVar(&CreateNetworkRunner.networkCloudHref)
	CreateNetworkRunner.Flag(`network.description`, `The description for the Network.`).StringVar(CreateNetworkRunner.networkDescription)
	CreateNetworkRunner.Flag(`network.instanceTenancy`, `The launch policy for AWS instances in the Network. Specify 'default' to allow instances to decide their own launch policy. Specify 'dedicated' to force all instances to be launched as 'dedicated'.`).StringVar(CreateNetworkRunner.networkInstanceTenancy)
	CreateNetworkRunner.Flag(`network.name`, `The name for the Network.`).StringVar(CreateNetworkRunner.networkName)
	registry[CreateNetworkCmd.FullCommand()] = CreateNetworkRunner

	DestroyNetworkRunner := new(DestroyNetworkNetworkRunner)
	DestroyNetworkCmd := resCmd.Command("DestroyNetwork", `Deletes the given network(s).`)
	DestroyNetworkRunner.Flag(`id`, ``).Required().StringVar(&DestroyNetworkRunner.id)
	registry[DestroyNetworkCmd.FullCommand()] = DestroyNetworkRunner

	IndexNetworksRunner := new(IndexNetworksNetworkRunner)
	IndexNetworksCmd := resCmd.Command("IndexNetworks", `Lists networks in this account.`)
	IndexNetworksRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexNetworksRunner.filterPos).StringsVar(IndexNetworksRunner.filter)
	registry[IndexNetworksCmd.FullCommand()] = IndexNetworksRunner

	ShowNetworkRunner := new(ShowNetworkNetworkRunner)
	ShowNetworkCmd := resCmd.Command("ShowNetwork", `Shows attributes of a single network.`)
	ShowNetworkRunner.Flag(`id`, ``).Required().StringVar(&ShowNetworkRunner.id)
	registry[ShowNetworkCmd.FullCommand()] = ShowNetworkRunner

	UpdateNetworkRunner := new(UpdateNetworkNetworkRunner)
	UpdateNetworkCmd := resCmd.Command("UpdateNetwork", `Updates the given network.`)
	UpdateNetworkRunner.Flag(`id`, ``).Required().StringVar(&UpdateNetworkRunner.id)
	UpdateNetworkRunner.Flag(`network.description`, `The updated description for the Network.`).StringVar(UpdateNetworkRunner.networkDescription)
	UpdateNetworkRunner.Flag(`network.name`, `The updated name for the Network.`).StringVar(UpdateNetworkRunner.networkName)
	UpdateNetworkRunner.Flag(`network.routeTableHref`, `Sets the default RouteTable for this Network.`).StringVar(UpdateNetworkRunner.networkRouteTableHref)
	registry[UpdateNetworkCmd.FullCommand()] = UpdateNetworkRunner
}

/****** NetworkGateway ******/

type CreateNetworkGatewayRunner struct {
	networkGatewayCloudHref   string
	networkGatewayDescription *string
	networkGatewayName        string
	networkGatewayType        string
}

func (r *CreateNetworkGatewayRunner) Run(c *Client) (interface{}, error) {

	/** Handle networkGateway parameter **/
	var networkGateway NetworkGatewayParam
	// Load JSON if provided
	if len(r.networkGatewayJson) > 0 {
		if err := Json.Unmarshal(r.networkGatewayJson, &networkGateway); err != nil {
			return nil, fmt.Errorf("Could not load networkGateway JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.NetworkGatewayCloudHref) > 0 {
		networkGateway.cloudHref = r.NetworkGatewayCloudHref
	}

	if r.NetworkGatewayDescription != nil {
		networkGateway.description = r.NetworkGatewayDescription
	}

	if len(r.NetworkGatewayName) > 0 {
		networkGateway.name = r.NetworkGatewayName
	}

	if len(r.NetworkGatewayType_) > 0 {
		networkGateway.type_ = r.NetworkGatewayType_
	}

	return c.CreateNetworkGateway(&networkGateway)
}

type DestroyNetworkGatewayRunner struct {
	id string
}

func (r *DestroyNetworkGatewayRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyNetworkGateway(r.id)
}

type IndexNetworkGatewaysRunner struct {
	filter    []string
	filterPos []string
}

func (r *IndexNetworkGatewaysRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter

	return c.IndexNetworkGateways(options)
}

type ShowNetworkGatewayRunner struct {
	id string
}

func (r *ShowNetworkGatewayRunner) Run(c *Client) (interface{}, error) {
	return c.ShowNetworkGateway(r.id)
}

type UpdateNetworkGatewayRunner struct {
	id                        string
	networkGatewayDescription *string
	networkGatewayName        *string
	networkGatewayNetworkHref *string
}

func (r *UpdateNetworkGatewayRunner) Run(c *Client) (interface{}, error) {

	/** Handle networkGateway parameter **/
	var networkGateway NetworkGatewayParam2
	// Load JSON if provided
	if len(r.networkGatewayJson) > 0 {
		if err := Json.Unmarshal(r.networkGatewayJson, &networkGateway); err != nil {
			return nil, fmt.Errorf("Could not load networkGateway JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.NetworkGatewayDescription != nil {
		networkGateway.description = r.NetworkGatewayDescription
	}

	if r.NetworkGatewayName != nil {
		networkGateway.name = r.NetworkGatewayName
	}

	if r.NetworkGatewayNetworkHref != nil {
		networkGateway.networkHref = r.NetworkGatewayNetworkHref
	}

	return c.UpdateNetworkGateway(r.id, &networkGateway)
}

// Register all NetworkGateway actions
func registerNetworkGatewayCmds(app *kingpin.Application) {
	resCmd := app.Cmd("NetworkGateway", `A NetworkGateway is an interface that allows traffic to be routed between networks.`)

	CreateNetworkGatewayRunner := new(CreateNetworkGatewayNetworkGatewayRunner)
	CreateNetworkGatewayCmd := resCmd.Command("CreateNetworkGateway", `Create a new NetworkGateway.`)
	CreateNetworkGatewayRunner.Flag(`networkGateway.cloudHref`, `The cloud to create the NetworkGateway in.`).Required().StringVar(&CreateNetworkGatewayRunner.networkGatewayCloudHref)
	CreateNetworkGatewayRunner.Flag(`networkGateway.description`, `The description to be set on the NetworkGateway.`).StringVar(CreateNetworkGatewayRunner.networkGatewayDescription)
	CreateNetworkGatewayRunner.Flag(`networkGateway.name`, `The name to be set on the NetworkGateway.`).Required().StringVar(&CreateNetworkGatewayRunner.networkGatewayName)
	CreateNetworkGatewayRunner.Flag(`networkGateway.type_`, `The type of the NetworkGateway.`).Required().StringVar(&CreateNetworkGatewayRunner.networkGatewayType)
	registry[CreateNetworkGatewayCmd.FullCommand()] = CreateNetworkGatewayRunner

	DestroyNetworkGatewayRunner := new(DestroyNetworkGatewayNetworkGatewayRunner)
	DestroyNetworkGatewayCmd := resCmd.Command("DestroyNetworkGateway", `Delete an existing NetworkGateway.`)
	DestroyNetworkGatewayRunner.Flag(`id`, ``).Required().StringVar(&DestroyNetworkGatewayRunner.id)
	registry[DestroyNetworkGatewayCmd.FullCommand()] = DestroyNetworkGatewayRunner

	IndexNetworkGatewaysRunner := new(IndexNetworkGatewaysNetworkGatewayRunner)
	IndexNetworkGatewaysCmd := resCmd.Command("IndexNetworkGateways", `Lists the NetworkGateways available to this account.`)
	IndexNetworkGatewaysRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexNetworkGatewaysRunner.filterPos).StringsVar(IndexNetworkGatewaysRunner.filter)
	registry[IndexNetworkGatewaysCmd.FullCommand()] = IndexNetworkGatewaysRunner

	ShowNetworkGatewayRunner := new(ShowNetworkGatewayNetworkGatewayRunner)
	ShowNetworkGatewayCmd := resCmd.Command("ShowNetworkGateway", `Show information about a single NetworkGateway.`)
	ShowNetworkGatewayRunner.Flag(`id`, ``).Required().StringVar(&ShowNetworkGatewayRunner.id)
	registry[ShowNetworkGatewayCmd.FullCommand()] = ShowNetworkGatewayRunner

	UpdateNetworkGatewayRunner := new(UpdateNetworkGatewayNetworkGatewayRunner)
	UpdateNetworkGatewayCmd := resCmd.Command("UpdateNetworkGateway", `Update an existing NetworkGateway.`)
	UpdateNetworkGatewayRunner.Flag(`id`, ``).Required().StringVar(&UpdateNetworkGatewayRunner.id)
	UpdateNetworkGatewayRunner.Flag(`networkGateway.description`, `The description to be set on the NetworkGateway.`).StringVar(UpdateNetworkGatewayRunner.networkGatewayDescription)
	UpdateNetworkGatewayRunner.Flag(`networkGateway.name`, `The name to be set on the NetworkGateway.`).StringVar(UpdateNetworkGatewayRunner.networkGatewayName)
	UpdateNetworkGatewayRunner.Flag(`networkGateway.networkHref`, `Pass a blank string to detach from the specified Network, or pass a valid Network href to attach to the specified network.`).StringVar(UpdateNetworkGatewayRunner.networkGatewayNetworkHref)
	registry[UpdateNetworkGatewayCmd.FullCommand()] = UpdateNetworkGatewayRunner
}

/****** NetworkOptionGroup ******/

type CreateNetworkOptionGroupRunner struct {
	networkOptionGroupCloudHref     string
	networkOptionGroupDescription   *string
	networkOptionGroupName          *string
	networkOptionGroupOptionsValues []string
	networkOptionGroupOptionsNames  []string
	networkOptionGroupType          string
}

func (r *CreateNetworkOptionGroupRunner) Run(c *Client) (interface{}, error) {

	/** Handle networkOptionGroup parameter **/
	var networkOptionGroup NetworkOptionGroupParam
	// Load JSON if provided
	if len(r.networkOptionGroupJson) > 0 {
		if err := Json.Unmarshal(r.networkOptionGroupJson, &networkOptionGroup); err != nil {
			return nil, fmt.Errorf("Could not load networkOptionGroup JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.NetworkOptionGroupCloudHref) > 0 {
		networkOptionGroup.cloudHref = r.NetworkOptionGroupCloudHref
	}

	if r.NetworkOptionGroupDescription != nil {
		networkOptionGroup.description = r.NetworkOptionGroupDescription
	}

	if r.NetworkOptionGroupName != nil {
		networkOptionGroup.name = r.NetworkOptionGroupName
	}

	if len(r.NetworkOptionGroupType_) > 0 {
		networkOptionGroup.type_ = r.NetworkOptionGroupType_
	}

	// Create enumarable fields from flags
	networkOptionGroupOptions := make(map[string]string, len(r.NetworkOptionGroupOptionsNames))
	for i, n := range r.NetworkOptionGroupOptionsNames {
		networkOptionGroupOptions[n] = r.NetworkOptionGroupOptionsValues[i]
	}
	networkOptionGroup.networkOptionGroup.options = networkOptionGroupOptions

	return c.CreateNetworkOptionGroup(&networkOptionGroup)
}

type DestroyNetworkOptionGroupRunner struct {
	id string
}

func (r *DestroyNetworkOptionGroupRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyNetworkOptionGroup(r.id)
}

type IndexNetworkOptionGroupsRunner struct {
	filter    []string
	filterPos []string
}

func (r *IndexNetworkOptionGroupsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter

	return c.IndexNetworkOptionGroups(options)
}

type ShowNetworkOptionGroupRunner struct {
	id string
}

func (r *ShowNetworkOptionGroupRunner) Run(c *Client) (interface{}, error) {
	return c.ShowNetworkOptionGroup(r.id)
}

type UpdateNetworkOptionGroupRunner struct {
	id                            string
	networkOptionGroupDescription *string
	networkOptionGroupName        *string
}

func (r *UpdateNetworkOptionGroupRunner) Run(c *Client) (interface{}, error) {

	/** Handle networkOptionGroup parameter **/
	var networkOptionGroup NetworkOptionGroupParam2
	// Load JSON if provided
	if len(r.networkOptionGroupJson) > 0 {
		if err := Json.Unmarshal(r.networkOptionGroupJson, &networkOptionGroup); err != nil {
			return nil, fmt.Errorf("Could not load networkOptionGroup JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.NetworkOptionGroupDescription != nil {
		networkOptionGroup.description = r.NetworkOptionGroupDescription
	}

	if r.NetworkOptionGroupName != nil {
		networkOptionGroup.name = r.NetworkOptionGroupName
	}

	return c.UpdateNetworkOptionGroup(r.id, &networkOptionGroup)
}

// Register all NetworkOptionGroup actions
func registerNetworkOptionGroupCmds(app *kingpin.Application) {
	resCmd := app.Cmd("NetworkOptionGroup", `A key/value pair hash containing options for configuring a Network.`)

	CreateNetworkOptionGroupRunner := new(CreateNetworkOptionGroupNetworkOptionGroupRunner)
	CreateNetworkOptionGroupCmd := resCmd.Command("CreateNetworkOptionGroup", `Create a new NetworkOptionGroup.`)
	CreateNetworkOptionGroupRunner.Flag(`networkOptionGroup.cloudHref`, `The Cloud to create this NetworkOptionGroup in`).Required().StringVar(&CreateNetworkOptionGroupRunner.networkOptionGroupCloudHref)
	CreateNetworkOptionGroupRunner.Flag(`networkOptionGroup.description`, `Description of this NetworkOptionGroup`).StringVar(CreateNetworkOptionGroupRunner.networkOptionGroupDescription)
	CreateNetworkOptionGroupRunner.Flag(`networkOptionGroup.name`, `Name of this NetworkOptionGroup`).StringVar(CreateNetworkOptionGroupRunner.networkOptionGroupName)
	CreateNetworkOptionGroupRunner.FlagPattern(`networkOptionGroup\.options\.([a-z0-9_]+)`, ``).Required().Capture(&CreateNetworkOptionGroupRunner.networkOptionGroupOptionsNames).StringVar(&CreateNetworkOptionGroupRunner.networkOptionGroupOptionsValues)
	CreateNetworkOptionGroupRunner.Flag(`networkOptionGroup.type_`, `Type of this NetworkOptionGroup`).Required().StringVar(&CreateNetworkOptionGroupRunner.networkOptionGroupType)
	registry[CreateNetworkOptionGroupCmd.FullCommand()] = CreateNetworkOptionGroupRunner

	DestroyNetworkOptionGroupRunner := new(DestroyNetworkOptionGroupNetworkOptionGroupRunner)
	DestroyNetworkOptionGroupCmd := resCmd.Command("DestroyNetworkOptionGroup", `Delete an existing NetworkOptionGroup.`)
	DestroyNetworkOptionGroupRunner.Flag(`id`, ``).Required().StringVar(&DestroyNetworkOptionGroupRunner.id)
	registry[DestroyNetworkOptionGroupCmd.FullCommand()] = DestroyNetworkOptionGroupRunner

	IndexNetworkOptionGroupsRunner := new(IndexNetworkOptionGroupsNetworkOptionGroupRunner)
	IndexNetworkOptionGroupsCmd := resCmd.Command("IndexNetworkOptionGroups", `List NetworkOptionGroups available in this account.`)
	IndexNetworkOptionGroupsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexNetworkOptionGroupsRunner.filterPos).StringsVar(IndexNetworkOptionGroupsRunner.filter)
	registry[IndexNetworkOptionGroupsCmd.FullCommand()] = IndexNetworkOptionGroupsRunner

	ShowNetworkOptionGroupRunner := new(ShowNetworkOptionGroupNetworkOptionGroupRunner)
	ShowNetworkOptionGroupCmd := resCmd.Command("ShowNetworkOptionGroup", `Show information about a single NetworkOptionGroup.`)
	ShowNetworkOptionGroupRunner.Flag(`id`, ``).Required().StringVar(&ShowNetworkOptionGroupRunner.id)
	registry[ShowNetworkOptionGroupCmd.FullCommand()] = ShowNetworkOptionGroupRunner

	UpdateNetworkOptionGroupRunner := new(UpdateNetworkOptionGroupNetworkOptionGroupRunner)
	UpdateNetworkOptionGroupCmd := resCmd.Command("UpdateNetworkOptionGroup", `Update an existing NetworkOptionGroup.`)
	UpdateNetworkOptionGroupRunner.Flag(`id`, ``).Required().StringVar(&UpdateNetworkOptionGroupRunner.id)
	UpdateNetworkOptionGroupRunner.Flag(`networkOptionGroup.description`, `Update the description`).StringVar(UpdateNetworkOptionGroupRunner.networkOptionGroupDescription)
	UpdateNetworkOptionGroupRunner.Flag(`networkOptionGroup.name`, `Update the name`).StringVar(UpdateNetworkOptionGroupRunner.networkOptionGroupName)
	registry[UpdateNetworkOptionGroupCmd.FullCommand()] = UpdateNetworkOptionGroupRunner
}

/****** NetworkOptionGroupAttachment ******/

type CreateNetworkOptionGroupAttachmentRunner struct {
	networkOptionGroupAttachmentNetworkHref            string
	networkOptionGroupAttachmentNetworkOptionGroupHref *string
}

func (r *CreateNetworkOptionGroupAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle networkOptionGroupAttachment parameter **/
	var networkOptionGroupAttachment NetworkOptionGroupAttachmentParam
	// Load JSON if provided
	if len(r.networkOptionGroupAttachmentJson) > 0 {
		if err := Json.Unmarshal(r.networkOptionGroupAttachmentJson, &networkOptionGroupAttachment); err != nil {
			return nil, fmt.Errorf("Could not load networkOptionGroupAttachment JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.NetworkOptionGroupAttachmentNetworkHref) > 0 {
		networkOptionGroupAttachment.networkHref = r.NetworkOptionGroupAttachmentNetworkHref
	}

	if r.NetworkOptionGroupAttachmentNetworkOptionGroupHref != nil {
		networkOptionGroupAttachment.networkOptionGroupHref = r.NetworkOptionGroupAttachmentNetworkOptionGroupHref
	}

	return c.CreateNetworkOptionGroupAttachment(&networkOptionGroupAttachment)
}

type DestroyNetworkOptionGroupAttachmentRunner struct {
	id string
}

func (r *DestroyNetworkOptionGroupAttachmentRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyNetworkOptionGroupAttachment(r.id)
}

type IndexNetworkOptionGroupAttachmentsRunner struct {
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexNetworkOptionGroupAttachmentsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexNetworkOptionGroupAttachments(options)
}

type ShowNetworkOptionGroupAttachmentRunner struct {
	id   string
	view *string
}

func (r *ShowNetworkOptionGroupAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowNetworkOptionGroupAttachment(r.id, options)
}

type UpdateNetworkOptionGroupAttachmentRunner struct {
	id                                                 string
	networkOptionGroupAttachmentNetworkOptionGroupHref *string
}

func (r *UpdateNetworkOptionGroupAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle networkOptionGroupAttachment parameter **/
	var networkOptionGroupAttachment NetworkOptionGroupAttachmentParam2
	// Load JSON if provided
	if len(r.networkOptionGroupAttachmentJson) > 0 {
		if err := Json.Unmarshal(r.networkOptionGroupAttachmentJson, &networkOptionGroupAttachment); err != nil {
			return nil, fmt.Errorf("Could not load networkOptionGroupAttachment JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.NetworkOptionGroupAttachmentNetworkOptionGroupHref != nil {
		networkOptionGroupAttachment.networkOptionGroupHref = r.NetworkOptionGroupAttachmentNetworkOptionGroupHref
	}

	return c.UpdateNetworkOptionGroupAttachment(r.id, &networkOptionGroupAttachment)
}

// Register all NetworkOptionGroupAttachment actions
func registerNetworkOptionGroupAttachmentCmds(app *kingpin.Application) {
	resCmd := app.Cmd("NetworkOptionGroupAttachment", `Resource for attaching NetworkOptionGroups to Networks.`)

	CreateNetworkOptionGroupAttachmentRunner := new(CreateNetworkOptionGroupAttachmentNetworkOptionGroupAttachmentRunner)
	CreateNetworkOptionGroupAttachmentCmd := resCmd.Command("CreateNetworkOptionGroupAttachment", `Create a new NetworkOptionGroupAttachment.`)
	CreateNetworkOptionGroupAttachmentRunner.Flag(`networkOptionGroupAttachment.networkHref`, `The Network to attach the specified NetworkOptionGroup to.`).Required().StringVar(&CreateNetworkOptionGroupAttachmentRunner.networkOptionGroupAttachmentNetworkHref)
	CreateNetworkOptionGroupAttachmentRunner.Flag(`networkOptionGroupAttachment.networkOptionGroupHref`, `The NetworkOptionGroup to attach to the specified resource.`).StringVar(CreateNetworkOptionGroupAttachmentRunner.networkOptionGroupAttachmentNetworkOptionGroupHref)
	registry[CreateNetworkOptionGroupAttachmentCmd.FullCommand()] = CreateNetworkOptionGroupAttachmentRunner

	DestroyNetworkOptionGroupAttachmentRunner := new(DestroyNetworkOptionGroupAttachmentNetworkOptionGroupAttachmentRunner)
	DestroyNetworkOptionGroupAttachmentCmd := resCmd.Command("DestroyNetworkOptionGroupAttachment", `Delete an existing NetworkOptionGroupAttachment.`)
	DestroyNetworkOptionGroupAttachmentRunner.Flag(`id`, ``).Required().StringVar(&DestroyNetworkOptionGroupAttachmentRunner.id)
	registry[DestroyNetworkOptionGroupAttachmentCmd.FullCommand()] = DestroyNetworkOptionGroupAttachmentRunner

	IndexNetworkOptionGroupAttachmentsRunner := new(IndexNetworkOptionGroupAttachmentsNetworkOptionGroupAttachmentRunner)
	IndexNetworkOptionGroupAttachmentsCmd := resCmd.Command("IndexNetworkOptionGroupAttachments", `List NetworkOptionGroupAttachments in this account.`)
	IndexNetworkOptionGroupAttachmentsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexNetworkOptionGroupAttachmentsRunner.filterPos).StringsVar(IndexNetworkOptionGroupAttachmentsRunner.filter)
	IndexNetworkOptionGroupAttachmentsRunner.Flag(`view`, ``).StringVar(IndexNetworkOptionGroupAttachmentsRunner.view)
	registry[IndexNetworkOptionGroupAttachmentsCmd.FullCommand()] = IndexNetworkOptionGroupAttachmentsRunner

	ShowNetworkOptionGroupAttachmentRunner := new(ShowNetworkOptionGroupAttachmentNetworkOptionGroupAttachmentRunner)
	ShowNetworkOptionGroupAttachmentCmd := resCmd.Command("ShowNetworkOptionGroupAttachment", `Show information about a single NetworkOptionGroupAttachment.`)
	ShowNetworkOptionGroupAttachmentRunner.Flag(`id`, ``).Required().StringVar(&ShowNetworkOptionGroupAttachmentRunner.id)
	ShowNetworkOptionGroupAttachmentRunner.Flag(`view`, ``).StringVar(ShowNetworkOptionGroupAttachmentRunner.view)
	registry[ShowNetworkOptionGroupAttachmentCmd.FullCommand()] = ShowNetworkOptionGroupAttachmentRunner

	UpdateNetworkOptionGroupAttachmentRunner := new(UpdateNetworkOptionGroupAttachmentNetworkOptionGroupAttachmentRunner)
	UpdateNetworkOptionGroupAttachmentCmd := resCmd.Command("UpdateNetworkOptionGroupAttachment", `Update an existing NetworkOptionGroupAttachment.`)
	UpdateNetworkOptionGroupAttachmentRunner.Flag(`id`, ``).Required().StringVar(&UpdateNetworkOptionGroupAttachmentRunner.id)
	UpdateNetworkOptionGroupAttachmentRunner.Flag(`networkOptionGroupAttachment.networkOptionGroupHref`, `The NetworkOptionGroup to attach to the specified resource.`).StringVar(UpdateNetworkOptionGroupAttachmentRunner.networkOptionGroupAttachmentNetworkOptionGroupHref)
	registry[UpdateNetworkOptionGroupAttachmentCmd.FullCommand()] = UpdateNetworkOptionGroupAttachmentRunner
}

/****** Oauth2 ******/

type CreateOauth2Runner struct {
	accountId        *int
	clientId         *string
	clientSecret     *string
	grantType        string
	refreshToken     *string
	rightLinkVersion *string
	rsVersion        *int
}

func (r *CreateOauth2Runner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.accountId != nil {
		options["account_id"] = *r.accountId
	}
	if r.clientId != nil {
		options["client_id"] = *r.clientId
	}
	if r.clientSecret != nil {
		options["client_secret"] = *r.clientSecret
	}
	if r.refreshToken != nil {
		options["refresh_token"] = *r.refreshToken
	}
	if r.rightLinkVersion != nil {
		options["right_link_version"] = *r.rightLinkVersion
	}
	if r.rsVersion != nil {
		options["r_s_version"] = *r.rsVersion
	}

	return c.CreateOauth2(r.grantType, options)
}

// Register all Oauth2 actions
func registerOauth2Cmds(app *kingpin.Application) {
	resCmd := app.Cmd("Oauth2", `Note that all API calls irrespective of the resource it is acting on, should pass a header "X-Api-Version" with the value "1...`)

	CreateOauth2Runner := new(CreateOauth2Oauth2Runner)
	CreateOauth2Cmd := resCmd.Command("CreateOauth2", `Perform an OAuth 2`)
	CreateOauth2Runner.Flag(`accountId`, `The client's account ID (only needed for instance agent clients).`).IntVar(CreateOauth2Runner.accountId)
	CreateOauth2Runner.Flag(`clientId`, `The client ID (only needed for confidential clients).`).StringVar(CreateOauth2Runner.clientId)
	CreateOauth2Runner.Flag(`clientSecret`, `The client secret (only needed for confidential clients).`).StringVar(CreateOauth2Runner.clientSecret)
	CreateOauth2Runner.Flag(`grantType`, `Type of grant.`).Required().StringVar(&CreateOauth2Runner.grantType)
	CreateOauth2Runner.Flag(`refreshToken`, `The refresh token obtained from OAuth grant.`).StringVar(CreateOauth2Runner.refreshToken)
	CreateOauth2Runner.Flag(`rightLinkVersion`, `The RightLink gem version the client conforms to (only needed for instance agent clients).`).StringVar(CreateOauth2Runner.rightLinkVersion)
	CreateOauth2Runner.Flag(`rsVersion`, `The RightAgent protocol version the client conforms to (only needed for instance agent clients).`).IntVar(CreateOauth2Runner.rsVersion)
	registry[CreateOauth2Cmd.FullCommand()] = CreateOauth2Runner
}

/****** Permission ******/

type CreatePermissionRunner struct {
	permissionRoleTitle string
	permissionUserHref  string
}

func (r *CreatePermissionRunner) Run(c *Client) (interface{}, error) {

	/** Handle permission parameter **/
	var permission PermissionParam
	// Load JSON if provided
	if len(r.permissionJson) > 0 {
		if err := Json.Unmarshal(r.permissionJson, &permission); err != nil {
			return nil, fmt.Errorf("Could not load permission JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.PermissionRoleTitle) > 0 {
		permission.roleTitle = r.PermissionRoleTitle
	}

	if len(r.PermissionUserHref) > 0 {
		permission.userHref = r.PermissionUserHref
	}

	return c.CreatePermission(&permission)
}

type DestroyPermissionRunner struct {
	id string
}

func (r *DestroyPermissionRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyPermission(r.id)
}

type IndexPermissionsRunner struct {
	filter    []string
	filterPos []string
}

func (r *IndexPermissionsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter

	return c.IndexPermissions(options)
}

type ShowPermissionRunner struct {
	id string
}

func (r *ShowPermissionRunner) Run(c *Client) (interface{}, error) {
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
	IndexPermissionsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexPermissionsRunner.filterPos).StringsVar(IndexPermissionsRunner.filter)
	registry[IndexPermissionsCmd.FullCommand()] = IndexPermissionsRunner

	ShowPermissionRunner := new(ShowPermissionPermissionRunner)
	ShowPermissionCmd := resCmd.Command("ShowPermission", `Show information about a single permission.`)
	ShowPermissionRunner.Flag(`id`, ``).Required().StringVar(&ShowPermissionRunner.id)
	registry[ShowPermissionCmd.FullCommand()] = ShowPermissionRunner
}

/****** PlacementGroup ******/

type CreatePlacementGroupRunner struct {
	placementGroupCloudHref   string
	placementGroupDescription *string
	placementGroupName        string
}

func (r *CreatePlacementGroupRunner) Run(c *Client) (interface{}, error) {

	/** Handle placementGroup parameter **/
	var placementGroup PlacementGroupParam
	// Load JSON if provided
	if len(r.placementGroupJson) > 0 {
		if err := Json.Unmarshal(r.placementGroupJson, &placementGroup); err != nil {
			return nil, fmt.Errorf("Could not load placementGroup JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.PlacementGroupCloudHref) > 0 {
		placementGroup.cloudHref = r.PlacementGroupCloudHref
	}

	if r.PlacementGroupDescription != nil {
		placementGroup.description = r.PlacementGroupDescription
	}

	if len(r.PlacementGroupName) > 0 {
		placementGroup.name = r.PlacementGroupName
	}

	return c.CreatePlacementGroup(&placementGroup)
}

type DestroyPlacementGroupRunner struct {
	id string
}

func (r *DestroyPlacementGroupRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyPlacementGroup(r.id)
}

type IndexPlacementGroupsRunner struct {
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexPlacementGroupsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexPlacementGroups(options)
}

type ShowPlacementGroupRunner struct {
	id   string
	view *string
}

func (r *ShowPlacementGroupRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowPlacementGroup(r.id, options)
}

// Register all PlacementGroup actions
func registerPlacementGroupCmds(app *kingpin.Application) {
	resCmd := app.Cmd("PlacementGroup", ``)

	CreatePlacementGroupRunner := new(CreatePlacementGroupPlacementGroupRunner)
	CreatePlacementGroupCmd := resCmd.Command("CreatePlacementGroup", `Creates a PlacementGroup.`)
	CreatePlacementGroupRunner.Flag(`placementGroup.cloudHref`, `The Href of the Cloud in which the PlacementGroup should be created. Note: This feature is not supported for all clouds.`).Required().StringVar(&CreatePlacementGroupRunner.placementGroupCloudHref)
	CreatePlacementGroupRunner.Flag(`placementGroup.description`, `The description of the Placement Group to be created.`).StringVar(CreatePlacementGroupRunner.placementGroupDescription)
	CreatePlacementGroupRunner.Flag(`placementGroup.name`, `The name of the Placement Group to be created.`).Required().StringVar(&CreatePlacementGroupRunner.placementGroupName)
	registry[CreatePlacementGroupCmd.FullCommand()] = CreatePlacementGroupRunner

	DestroyPlacementGroupRunner := new(DestroyPlacementGroupPlacementGroupRunner)
	DestroyPlacementGroupCmd := resCmd.Command("DestroyPlacementGroup", `Destroys a PlacementGroup.`)
	DestroyPlacementGroupRunner.Flag(`id`, ``).Required().StringVar(&DestroyPlacementGroupRunner.id)
	registry[DestroyPlacementGroupCmd.FullCommand()] = DestroyPlacementGroupRunner

	IndexPlacementGroupsRunner := new(IndexPlacementGroupsPlacementGroupRunner)
	IndexPlacementGroupsCmd := resCmd.Command("IndexPlacementGroups", `Lists all PlacementGroups in an account.`)
	IndexPlacementGroupsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexPlacementGroupsRunner.filterPos).StringsVar(IndexPlacementGroupsRunner.filter)
	IndexPlacementGroupsRunner.Flag(`view`, ``).StringVar(IndexPlacementGroupsRunner.view)
	registry[IndexPlacementGroupsCmd.FullCommand()] = IndexPlacementGroupsRunner

	ShowPlacementGroupRunner := new(ShowPlacementGroupPlacementGroupRunner)
	ShowPlacementGroupCmd := resCmd.Command("ShowPlacementGroup", `Shows information about a single PlacementGroup.`)
	ShowPlacementGroupRunner.Flag(`id`, ``).Required().StringVar(&ShowPlacementGroupRunner.id)
	ShowPlacementGroupRunner.Flag(`view`, ``).StringVar(ShowPlacementGroupRunner.view)
	registry[ShowPlacementGroupCmd.FullCommand()] = ShowPlacementGroupRunner
}

/****** Preference ******/

type DestroyPreferenceRunner struct {
	id string
}

func (r *DestroyPreferenceRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyPreference(r.id)
}

type IndexPreferencesRunner struct {
	filter    []string
	filterPos []string
}

func (r *IndexPreferencesRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter

	return c.IndexPreferences(options)
}

type ShowPreferenceRunner struct {
	id string
}

func (r *ShowPreferenceRunner) Run(c *Client) (interface{}, error) {
	return c.ShowPreference(r.id)
}

type UpdatePreferenceRunner struct {
	id                 string
	preferenceContents string
}

func (r *UpdatePreferenceRunner) Run(c *Client) (interface{}, error) {

	/** Handle preference parameter **/
	var preference PreferenceParam
	// Load JSON if provided
	if len(r.preferenceJson) > 0 {
		if err := Json.Unmarshal(r.preferenceJson, &preference); err != nil {
			return nil, fmt.Errorf("Could not load preference JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.PreferenceContents) > 0 {
		preference.contents = r.PreferenceContents
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
	IndexPreferencesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexPreferencesRunner.filterPos).StringsVar(IndexPreferencesRunner.filter)
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

type ImportPublicationRunner struct {
	id string
}

func (r *ImportPublicationRunner) Run(c *Client) (interface{}, error) {
	return c.ImportPublication(r.id)
}

type IndexPublicationsRunner struct {
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexPublicationsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexPublications(options)
}

type ShowPublicationRunner struct {
	id   string
	view *string
}

func (r *ShowPublicationRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowPublication(r.id, options)
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
	IndexPublicationsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexPublicationsRunner.filterPos).StringsVar(IndexPublicationsRunner.filter)
	IndexPublicationsRunner.Flag(`view`, ``).StringVar(IndexPublicationsRunner.view)
	registry[IndexPublicationsCmd.FullCommand()] = IndexPublicationsRunner

	ShowPublicationRunner := new(ShowPublicationPublicationRunner)
	ShowPublicationCmd := resCmd.Command("ShowPublication", `Show information about a single publication. Only non-HEAD revisions are possible.`)
	ShowPublicationRunner.Flag(`id`, ``).Required().StringVar(&ShowPublicationRunner.id)
	ShowPublicationRunner.Flag(`view`, ``).StringVar(ShowPublicationRunner.view)
	registry[ShowPublicationCmd.FullCommand()] = ShowPublicationRunner
}

/****** PublicationLineage ******/

type ShowPublicationLineageRunner struct {
	id   string
	view *string
}

func (r *ShowPublicationLineageRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowPublicationLineage(r.id, options)
}

// Register all PublicationLineage actions
func registerPublicationLineageCmds(app *kingpin.Application) {
	resCmd := app.Cmd("PublicationLineage", `A Publication Lineage contains lineage information for a Publication in the MultiCloudMarketplace.`)

	ShowPublicationLineageRunner := new(ShowPublicationLineagePublicationLineageRunner)
	ShowPublicationLineageCmd := resCmd.Command("ShowPublicationLineage", `Show information about a single publication lineage. Only non-HEAD revisions are possible.`)
	ShowPublicationLineageRunner.Flag(`id`, ``).Required().StringVar(&ShowPublicationLineageRunner.id)
	ShowPublicationLineageRunner.Flag(`view`, ``).StringVar(ShowPublicationLineageRunner.view)
	registry[ShowPublicationLineageCmd.FullCommand()] = ShowPublicationLineageRunner
}

/****** RecurringVolumeAttachment ******/

type CreateRecurringVolumeAttachmentRunner struct {
	cloudId                               string
	recurringVolumeAttachmentDevice       string
	recurringVolumeAttachmentRunnableHref string
	recurringVolumeAttachmentStorageHref  string
}

func (r *CreateRecurringVolumeAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle recurringVolumeAttachment parameter **/
	var recurringVolumeAttachment RecurringVolumeAttachmentParam
	// Load JSON if provided
	if len(r.recurringVolumeAttachmentJson) > 0 {
		if err := Json.Unmarshal(r.recurringVolumeAttachmentJson, &recurringVolumeAttachment); err != nil {
			return nil, fmt.Errorf("Could not load recurringVolumeAttachment JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.RecurringVolumeAttachmentDevice) > 0 {
		recurringVolumeAttachment.device = r.RecurringVolumeAttachmentDevice
	}

	if len(r.RecurringVolumeAttachmentRunnableHref) > 0 {
		recurringVolumeAttachment.runnableHref = r.RecurringVolumeAttachmentRunnableHref
	}

	if len(r.RecurringVolumeAttachmentStorageHref) > 0 {
		recurringVolumeAttachment.storageHref = r.RecurringVolumeAttachmentStorageHref
	}

	return c.CreateRecurringVolumeAttachment(r.cloudId, &recurringVolumeAttachment)
}

type DestroyRecurringVolumeAttachmentRunner struct {
	cloudId string
	id      string
}

func (r *DestroyRecurringVolumeAttachmentRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyRecurringVolumeAttachment(r.cloudId, r.id)
}

type IndexRecurringVolumeAttachmentsRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexRecurringVolumeAttachmentsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexRecurringVolumeAttachments(r.cloudId, options)
}

type ShowRecurringVolumeAttachmentRunner struct {
	cloudId string
	id      string
	view    *string
}

func (r *ShowRecurringVolumeAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowRecurringVolumeAttachment(r.cloudId, r.id, options)
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
	IndexRecurringVolumeAttachmentsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexRecurringVolumeAttachmentsRunner.filterPos).StringsVar(IndexRecurringVolumeAttachmentsRunner.filter)
	IndexRecurringVolumeAttachmentsRunner.Flag(`view`, ``).StringVar(IndexRecurringVolumeAttachmentsRunner.view)
	registry[IndexRecurringVolumeAttachmentsCmd.FullCommand()] = IndexRecurringVolumeAttachmentsRunner

	ShowRecurringVolumeAttachmentRunner := new(ShowRecurringVolumeAttachmentRecurringVolumeAttachmentRunner)
	ShowRecurringVolumeAttachmentCmd := resCmd.Command("ShowRecurringVolumeAttachment", `Displays information about a single recurring volume attachment.`)
	ShowRecurringVolumeAttachmentRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowRecurringVolumeAttachmentRunner.cloudId)
	ShowRecurringVolumeAttachmentRunner.Flag(`id`, ``).Required().StringVar(&ShowRecurringVolumeAttachmentRunner.id)
	ShowRecurringVolumeAttachmentRunner.Flag(`view`, ``).StringVar(ShowRecurringVolumeAttachmentRunner.view)
	registry[ShowRecurringVolumeAttachmentCmd.FullCommand()] = ShowRecurringVolumeAttachmentRunner
}

/****** Repository ******/

type CookbookImportRepositoryRunner struct {
	assetHrefs                []string
	assetHrefsPos             []string
	follow                    *string
	id                        string
	namespace                 *string
	repositoryCommitReference *string
	withDependencies          *string
}

func (r *CookbookImportRepositoryRunner) Run(c *Client) (interface{}, error) {

	/** Handle assetHrefs parameter **/
	var assetHrefs []string

	for i, v := range r.assetHrefs {
		pos, err := strconv.Atoi(r.assetHrefsPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for assetHrefs array", r.assetHrefsPos[i])
		}
		assetHrefs[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.follow != nil {
		options["follow"] = *r.follow
	}
	if r.namespace != nil {
		options["namespace"] = *r.namespace
	}
	if r.repositoryCommitReference != nil {
		options["repository_commit_reference"] = *r.repositoryCommitReference
	}
	if r.withDependencies != nil {
		options["with_dependencies"] = *r.withDependencies
	}

	return c.CookbookImportRepository(assetHrefs, r.id, options)
}

type CookbookImportPreviewRepositoryRunner struct {
	assetHrefs    []string
	assetHrefsPos []string
	id            string
	namespace     string
}

func (r *CookbookImportPreviewRepositoryRunner) Run(c *Client) (interface{}, error) {

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

type CreateRepositoryRunner struct {
	repositoryAssetPathsCookbooks    []string
	repositoryAssetPathsCookbooksPos []string
	repositoryAutoImport             *string
	repositoryCommitReference        *string
	repositoryCredentialsPassword    *string
	repositoryCredentialsSshKey      *string
	repositoryCredentialsUsername    *string
	repositoryDescription            *string
	repositoryName                   string
	repositorySource                 string
	repositorySourceType             string
}

func (r *CreateRepositoryRunner) Run(c *Client) (interface{}, error) {

	/** Handle repository parameter **/
	var repository RepositoryParam
	// Load JSON if provided
	if len(r.repositoryJson) > 0 {
		if err := Json.Unmarshal(r.repositoryJson, &repository); err != nil {
			return nil, fmt.Errorf("Could not load repository JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.RepositoryAutoImport != nil {
		repository.autoImport = r.RepositoryAutoImport
	}

	if r.RepositoryCommitReference != nil {
		repository.commitReference = r.RepositoryCommitReference
	}

	if r.RepositoryCredentialsPassword != nil {
		repository.credentials.password = r.RepositoryCredentialsPassword
	}

	if r.RepositoryCredentialsSshKey != nil {
		repository.credentials.sshKey = r.RepositoryCredentialsSshKey
	}

	if r.RepositoryCredentialsUsername != nil {
		repository.credentials.username = r.RepositoryCredentialsUsername
	}

	if r.RepositoryDescription != nil {
		repository.description = r.RepositoryDescription
	}

	if len(r.RepositoryName) > 0 {
		repository.name = r.RepositoryName
	}

	if len(r.RepositorySource) > 0 {
		repository.source = r.RepositorySource
	}

	if len(r.RepositorySourceType) > 0 {
		repository.sourceType = r.RepositorySourceType
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxRepositoryAssetPathsCookbooksPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.RepositoryAssetPathsCookbooksPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for repository.assetPaths.cookbooks field of cookbooks array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of repository.assetPaths.cookbooks field of cookbooks array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxRepositoryAssetPathsCookbooksPos {
			maxRepositoryAssetPathsCookbooksPos = pos
		}
	}
	if len(r.RepositoryAssetPathsCookbooks) != maxRepositoryAssetPathsCookbooksPos {
		return nil, fmt.Errof("Invalid flags for repository.assetPaths.cookbooks field of cookbooks array, %s were defined but max position is %s",
			len(r.RepositoryAssetPathsCookbooks), maxRepositoryAssetPathsCookbooksPos)
	}
	if maxRepositoryAssetPathsCookbooksPos > -1 {
		repositoryAssetPathsCookbooks := make([]*string, maxRepositoryAssetPathsCookbooksPos+1)
		for i := 0; i < maxRepositoryAssetPathsCookbooksPos+1; i++ {
			repositoryAssetPathsCookbooks[i] = r.repositoryAssetPathsCookbooks[r.repositoryAssetPathsCookbooksPos[i]]
		}
		repository.assetPaths.cookbooks = repositoryAssetPathsCookbooks
	}

	return c.CreateRepository(&repository)
}

type DestroyRepositoryRunner struct {
	id string
}

func (r *DestroyRepositoryRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyRepository(r.id)
}

type IndexRepositoriesRunner struct {
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexRepositoriesRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexRepositories(options)
}

type RefetchRepositoryRunner struct {
	autoImport *string
	id         string
}

func (r *RefetchRepositoryRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.autoImport != nil {
		options["auto_import"] = *r.autoImport
	}

	return c.RefetchRepository(r.id, options)
}

type ResolveRepositoryRunner struct {
	importedCookbookName    []string
	importedCookbookNamePos []string
}

func (r *ResolveRepositoryRunner) Run(c *Client) (interface{}, error) {

	/** Handle importedCookbookName parameter **/
	var importedCookbookName []string

	for i, v := range r.importedCookbookName {
		pos, err := strconv.Atoi(r.importedCookbookNamePos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for importedCookbookName array", r.importedCookbookNamePos[i])
		}
		importedCookbookName[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["imported_cookbook_name"] = importedCookbookName

	return c.ResolveRepository(options)
}

type ShowRepositoryRunner struct {
	id   string
	view *string
}

func (r *ShowRepositoryRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowRepository(r.id, options)
}

type UpdateRepositoryRunner struct {
	id                               string
	repositoryAssetPathsCookbooks    []string
	repositoryAssetPathsCookbooksPos []string
	repositoryCommitReference        *string
	repositoryCredentialsPassword    *string
	repositoryCredentialsSshKey      *string
	repositoryCredentialsUsername    *string
	repositoryDescription            *string
	repositoryName                   *string
	repositorySource                 *string
	repositorySourceType             *string
}

func (r *UpdateRepositoryRunner) Run(c *Client) (interface{}, error) {

	/** Handle repository parameter **/
	var repository RepositoryParam2
	// Load JSON if provided
	if len(r.repositoryJson) > 0 {
		if err := Json.Unmarshal(r.repositoryJson, &repository); err != nil {
			return nil, fmt.Errorf("Could not load repository JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.RepositoryCommitReference != nil {
		repository.commitReference = r.RepositoryCommitReference
	}

	if r.RepositoryCredentialsPassword != nil {
		repository.credentials.password = r.RepositoryCredentialsPassword
	}

	if r.RepositoryCredentialsSshKey != nil {
		repository.credentials.sshKey = r.RepositoryCredentialsSshKey
	}

	if r.RepositoryCredentialsUsername != nil {
		repository.credentials.username = r.RepositoryCredentialsUsername
	}

	if r.RepositoryDescription != nil {
		repository.description = r.RepositoryDescription
	}

	if r.RepositoryName != nil {
		repository.name = r.RepositoryName
	}

	if r.RepositorySource != nil {
		repository.source = r.RepositorySource
	}

	if r.RepositorySourceType != nil {
		repository.sourceType = r.RepositorySourceType
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxRepositoryAssetPathsCookbooksPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.RepositoryAssetPathsCookbooksPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for repository.assetPaths.cookbooks field of cookbooks array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of repository.assetPaths.cookbooks field of cookbooks array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxRepositoryAssetPathsCookbooksPos {
			maxRepositoryAssetPathsCookbooksPos = pos
		}
	}
	if len(r.RepositoryAssetPathsCookbooks) != maxRepositoryAssetPathsCookbooksPos {
		return nil, fmt.Errof("Invalid flags for repository.assetPaths.cookbooks field of cookbooks array, %s were defined but max position is %s",
			len(r.RepositoryAssetPathsCookbooks), maxRepositoryAssetPathsCookbooksPos)
	}
	if maxRepositoryAssetPathsCookbooksPos > -1 {
		repositoryAssetPathsCookbooks := make([]*string, maxRepositoryAssetPathsCookbooksPos+1)
		for i := 0; i < maxRepositoryAssetPathsCookbooksPos+1; i++ {
			repositoryAssetPathsCookbooks[i] = r.repositoryAssetPathsCookbooks[r.repositoryAssetPathsCookbooksPos[i]]
		}
		repository.assetPaths.cookbooks = repositoryAssetPathsCookbooks
	}

	return c.UpdateRepository(r.id, &repository)
}

// Register all Repository actions
func registerRepositoryCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Repository", `A Repository is a location from which you can download and import design objects such as Chef cookbooks. Using this resource you can add and modify repository information and import assets discovered in the repository.`)

	CookbookImportRepositoryRunner := new(CookbookImportRepositoryRepositoryRunner)
	CookbookImportRepositoryCmd := resCmd.Command("CookbookImportRepository", `Performs a Cookbook import, which allows you to use the specified cookbooks in your design objects.`)
	CookbookImportRepositoryRunner.FlagPattern(`assetHrefs\.(\d+)`, `Hrefs of the assets that should be imported.`).Required().Capture(&CookbookImportRepositoryRunner.assetHrefsPos).StringsVar(&CookbookImportRepositoryRunner.assetHrefs)
	CookbookImportRepositoryRunner.Flag(`follow`, `A flag indicating whether imported cookbooks should be followed.`).StringVar(CookbookImportRepositoryRunner.follow)
	CookbookImportRepositoryRunner.Flag(`id`, ``).Required().StringVar(&CookbookImportRepositoryRunner.id)
	CookbookImportRepositoryRunner.Flag(`namespace`, `The namespace to import into.`).StringVar(CookbookImportRepositoryRunner.namespace)
	CookbookImportRepositoryRunner.Flag(`repositoryCommitReference`, `Optional commit reference indicating last succeeded commit. Must match the Repository's fetch_status.succeeded_commit attribute or the import will not be performed.`).StringVar(CookbookImportRepositoryRunner.repositoryCommitReference)
	CookbookImportRepositoryRunner.Flag(`withDependencies`, `A flag indicating whether dependencies should automatically be imported.`).StringVar(CookbookImportRepositoryRunner.withDependencies)
	registry[CookbookImportRepositoryCmd.FullCommand()] = CookbookImportRepositoryRunner

	CookbookImportPreviewRepositoryRunner := new(CookbookImportPreviewRepositoryRepositoryRunner)
	CookbookImportPreviewRepositoryCmd := resCmd.Command("CookbookImportPreviewRepository", `Retrieves a preview of the effects of a Cookbook import.`)
	CookbookImportPreviewRepositoryRunner.FlagPattern(`assetHrefs\.(\d+)`, `Hrefs of the assets that should be imported.`).Required().Capture(&CookbookImportPreviewRepositoryRunner.assetHrefsPos).StringsVar(&CookbookImportPreviewRepositoryRunner.assetHrefs)
	CookbookImportPreviewRepositoryRunner.Flag(`id`, ``).Required().StringVar(&CookbookImportPreviewRepositoryRunner.id)
	CookbookImportPreviewRepositoryRunner.Flag(`namespace`, `The namespace to import into.`).Required().StringVar(&CookbookImportPreviewRepositoryRunner.namespace)
	registry[CookbookImportPreviewRepositoryCmd.FullCommand()] = CookbookImportPreviewRepositoryRunner

	CreateRepositoryRunner := new(CreateRepositoryRepositoryRunner)
	CreateRepositoryCmd := resCmd.Command("CreateRepository", `Creates a Repository.`)
	CreateRepositoryRunner.FlagPattern(`repository\.assetPaths\.cookbooks\.(\d+)`, `The cookbook paths for the repository`).Capture(&CreateRepositoryRunner.repositoryAssetPathsCookbooksPos).StringsVar(CreateRepositoryRunner.repositoryAssetPathsCookbooks)
	CreateRepositoryRunner.Flag(`repository.autoImport`, `Whether cookbooks should automatically be imported upon repository creation.`).StringVar(CreateRepositoryRunner.repositoryAutoImport)
	CreateRepositoryRunner.Flag(`repository.commitReference`, `The revision for the repository`).StringVar(CreateRepositoryRunner.repositoryCommitReference)
	CreateRepositoryRunner.Flag(`repository.credentials.password`, `The password, or credential, for the repository (only valid for svn or download repositories).`).StringVar(CreateRepositoryRunner.repositoryCredentialsPassword)
	CreateRepositoryRunner.Flag(`repository.credentials.sshKey`, `The SSH key, or credential, for the repository (only valid for git repositories).`).StringVar(CreateRepositoryRunner.repositoryCredentialsSshKey)
	CreateRepositoryRunner.Flag(`repository.credentials.username`, `The user name, or credential, for the repository (only valid for svn or download repositories).`).StringVar(CreateRepositoryRunner.repositoryCredentialsUsername)
	CreateRepositoryRunner.Flag(`repository.description`, `The description for the repository.`).StringVar(CreateRepositoryRunner.repositoryDescription)
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
	IndexRepositoriesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexRepositoriesRunner.filterPos).StringsVar(IndexRepositoriesRunner.filter)
	IndexRepositoriesRunner.Flag(`view`, ``).StringVar(IndexRepositoriesRunner.view)
	registry[IndexRepositoriesCmd.FullCommand()] = IndexRepositoriesRunner

	RefetchRepositoryRunner := new(RefetchRepositoryRepositoryRunner)
	RefetchRepositoryCmd := resCmd.Command("RefetchRepository", `Refetches all RepositoryAssets associated with the Repository.`)
	RefetchRepositoryRunner.Flag(`autoImport`, `Whether cookbooks should automatically be imported after repositories are fetched.`).StringVar(RefetchRepositoryRunner.autoImport)
	RefetchRepositoryRunner.Flag(`id`, ``).Required().StringVar(&RefetchRepositoryRunner.id)
	registry[RefetchRepositoryCmd.FullCommand()] = RefetchRepositoryRunner

	ResolveRepositoryRunner := new(ResolveRepositoryRepositoryRunner)
	ResolveRepositoryCmd := resCmd.Command("ResolveRepository", `Show a list of repositories that have imported cookbooks with the given names.`)
	ResolveRepositoryRunner.FlagPattern(`importedCookbookName\.(\d+)`, `A list of cookbook names that were imported by the repository.`).Capture(&ResolveRepositoryRunner.importedCookbookNamePos).StringsVar(ResolveRepositoryRunner.importedCookbookName)
	registry[ResolveRepositoryCmd.FullCommand()] = ResolveRepositoryRunner

	ShowRepositoryRunner := new(ShowRepositoryRepositoryRunner)
	ShowRepositoryCmd := resCmd.Command("ShowRepository", `Shows a specified Repository.`)
	ShowRepositoryRunner.Flag(`id`, ``).Required().StringVar(&ShowRepositoryRunner.id)
	ShowRepositoryRunner.Flag(`view`, ``).StringVar(ShowRepositoryRunner.view)
	registry[ShowRepositoryCmd.FullCommand()] = ShowRepositoryRunner

	UpdateRepositoryRunner := new(UpdateRepositoryRepositoryRunner)
	UpdateRepositoryCmd := resCmd.Command("UpdateRepository", `Updates a specified Repository.`)
	UpdateRepositoryRunner.Flag(`id`, ``).Required().StringVar(&UpdateRepositoryRunner.id)
	UpdateRepositoryRunner.FlagPattern(`repository\.assetPaths\.cookbooks\.(\d+)`, `The updated cookbook paths for the repository`).Capture(&UpdateRepositoryRunner.repositoryAssetPathsCookbooksPos).StringsVar(UpdateRepositoryRunner.repositoryAssetPathsCookbooks)
	UpdateRepositoryRunner.Flag(`repository.commitReference`, `The updated commit reference (tag, branch, revision...) for the repository`).StringVar(UpdateRepositoryRunner.repositoryCommitReference)
	UpdateRepositoryRunner.Flag(`repository.credentials.password`, `The updated password, or credential, for the repository (only valid for svn or download repositories).`).StringVar(UpdateRepositoryRunner.repositoryCredentialsPassword)
	UpdateRepositoryRunner.Flag(`repository.credentials.sshKey`, `The updated SSH key for the repository (only valid for git repositories).`).StringVar(UpdateRepositoryRunner.repositoryCredentialsSshKey)
	UpdateRepositoryRunner.Flag(`repository.credentials.username`, `The updated user name, or credential, for the repository (only valid for svn or download repositories).`).StringVar(UpdateRepositoryRunner.repositoryCredentialsUsername)
	UpdateRepositoryRunner.Flag(`repository.description`, `The updated description for the repository.`).StringVar(UpdateRepositoryRunner.repositoryDescription)
	UpdateRepositoryRunner.Flag(`repository.name`, `The updated repository name.`).StringVar(UpdateRepositoryRunner.repositoryName)
	UpdateRepositoryRunner.Flag(`repository.source`, `The updated URL for the repository.`).StringVar(UpdateRepositoryRunner.repositorySource)
	UpdateRepositoryRunner.Flag(`repository.sourceType`, `The updated source type for the repository.`).StringVar(UpdateRepositoryRunner.repositorySourceType)
	registry[UpdateRepositoryCmd.FullCommand()] = UpdateRepositoryRunner
}

/****** RepositoryAsset ******/

type IndexRepositoryAssetsRunner struct {
	repositoryId string
	view         *string
}

func (r *IndexRepositoryAssetsRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexRepositoryAssets(r.repositoryId, options)
}

type ShowRepositoryAssetRunner struct {
	id           string
	repositoryId string
	view         *string
}

func (r *ShowRepositoryAssetRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowRepositoryAsset(r.id, r.repositoryId, options)
}

// Register all RepositoryAsset actions
func registerRepositoryAssetCmds(app *kingpin.Application) {
	resCmd := app.Cmd("RepositoryAsset", `A RepositoryAsset represents an item discovered in a Repository`)

	IndexRepositoryAssetsRunner := new(IndexRepositoryAssetsRepositoryAssetRunner)
	IndexRepositoryAssetsCmd := resCmd.Command("IndexRepositoryAssets", `List a repository's current assets.`)
	IndexRepositoryAssetsRunner.Flag(`repositoryId`, ``).Required().StringVar(&IndexRepositoryAssetsRunner.repositoryId)
	IndexRepositoryAssetsRunner.Flag(`view`, ``).StringVar(IndexRepositoryAssetsRunner.view)
	registry[IndexRepositoryAssetsCmd.FullCommand()] = IndexRepositoryAssetsRunner

	ShowRepositoryAssetRunner := new(ShowRepositoryAssetRepositoryAssetRunner)
	ShowRepositoryAssetCmd := resCmd.Command("ShowRepositoryAsset", `Show information about a single asset.`)
	ShowRepositoryAssetRunner.Flag(`id`, ``).Required().StringVar(&ShowRepositoryAssetRunner.id)
	ShowRepositoryAssetRunner.Flag(`repositoryId`, ``).Required().StringVar(&ShowRepositoryAssetRunner.repositoryId)
	ShowRepositoryAssetRunner.Flag(`view`, ``).StringVar(ShowRepositoryAssetRunner.view)
	registry[ShowRepositoryAssetCmd.FullCommand()] = ShowRepositoryAssetRunner
}

/****** RightScript ******/

type CommitRightScriptRunner struct {
	id                       string
	rightScriptCommitMessage string
}

func (r *CommitRightScriptRunner) Run(c *Client) (interface{}, error) {

	/** Handle rightScript parameter **/
	var rightScript RightScriptParam
	// Load JSON if provided
	if len(r.rightScriptJson) > 0 {
		if err := Json.Unmarshal(r.rightScriptJson, &rightScript); err != nil {
			return nil, fmt.Errorf("Could not load rightScript JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.RightScriptCommitMessage) > 0 {
		rightScript.commitMessage = r.RightScriptCommitMessage
	}

	return c.CommitRightScript(r.id, &rightScript)
}

type IndexRightScriptsRunner struct {
	filter     []string
	filterPos  []string
	latestOnly *string
	view       *string
}

func (r *IndexRightScriptsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.latestOnly != nil {
		options["latest_only"] = *r.latestOnly
	}
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexRightScripts(options)
}

type ShowRightScriptRunner struct {
	id string
}

func (r *ShowRightScriptRunner) Run(c *Client) (interface{}, error) {
	return c.ShowRightScript(r.id)
}

type ShowSourceRightScriptRunner struct {
	id string
}

func (r *ShowSourceRightScriptRunner) Run(c *Client) (interface{}, error) {
	return c.ShowSourceRightScript(r.id)
}

type UpdateRightScriptRunner struct {
	id                     string
	rightScriptDescription *string
	rightScriptName        *string
}

func (r *UpdateRightScriptRunner) Run(c *Client) (interface{}, error) {

	/** Handle rightScript parameter **/
	var rightScript RightScriptParam2
	// Load JSON if provided
	if len(r.rightScriptJson) > 0 {
		if err := Json.Unmarshal(r.rightScriptJson, &rightScript); err != nil {
			return nil, fmt.Errorf("Could not load rightScript JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.RightScriptDescription != nil {
		rightScript.description = r.RightScriptDescription
	}

	if r.RightScriptName != nil {
		rightScript.name = r.RightScriptName
	}

	return c.UpdateRightScript(r.id, &rightScript)
}

type UpdateSourceRightScriptRunner struct {
	id string
}

func (r *UpdateSourceRightScriptRunner) Run(c *Client) (interface{}, error) {
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
	IndexRightScriptsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexRightScriptsRunner.filterPos).StringsVar(IndexRightScriptsRunner.filter)
	IndexRightScriptsRunner.Flag(`latestOnly`, `Whether or not to return only the latest version for each lineage.`).StringVar(IndexRightScriptsRunner.latestOnly)
	IndexRightScriptsRunner.Flag(`view`, ``).StringVar(IndexRightScriptsRunner.view)
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
	UpdateRightScriptRunner.Flag(`rightScript.description`, `The new description for the RightScript`).StringVar(UpdateRightScriptRunner.rightScriptDescription)
	UpdateRightScriptRunner.Flag(`rightScript.name`, `The new name for the RightScript`).StringVar(UpdateRightScriptRunner.rightScriptName)
	registry[UpdateRightScriptCmd.FullCommand()] = UpdateRightScriptRunner

	UpdateSourceRightScriptRunner := new(UpdateSourceRightScriptRightScriptRunner)
	UpdateSourceRightScriptCmd := resCmd.Command("UpdateSourceRightScript", `Updates the source of the given RightScript`)
	UpdateSourceRightScriptRunner.Flag(`id`, ``).Required().StringVar(&UpdateSourceRightScriptRunner.id)
	registry[UpdateSourceRightScriptCmd.FullCommand()] = UpdateSourceRightScriptRunner
}

/****** Route ******/

type CreateRouteRunner struct {
	routeDescription          *string
	routeDestinationCidrBlock string
	routeNextHopHref          *string
	routeNextHopIp            *string
	routeNextHopType          string
	routeRouteTableHref       string
}

func (r *CreateRouteRunner) Run(c *Client) (interface{}, error) {

	/** Handle route parameter **/
	var route RouteParam
	// Load JSON if provided
	if len(r.routeJson) > 0 {
		if err := Json.Unmarshal(r.routeJson, &route); err != nil {
			return nil, fmt.Errorf("Could not load route JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.RouteDescription != nil {
		route.description = r.RouteDescription
	}

	if len(r.RouteDestinationCidrBlock) > 0 {
		route.destinationCidrBlock = r.RouteDestinationCidrBlock
	}

	if r.RouteNextHopHref != nil {
		route.nextHopHref = r.RouteNextHopHref
	}

	if r.RouteNextHopIp != nil {
		route.nextHopIp = r.RouteNextHopIp
	}

	if len(r.RouteNextHopType) > 0 {
		route.nextHopType = r.RouteNextHopType
	}

	if len(r.RouteRouteTableHref) > 0 {
		route.routeTableHref = r.RouteRouteTableHref
	}

	return c.CreateRoute(&route)
}

type DestroyRouteRunner struct {
	id string
}

func (r *DestroyRouteRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyRoute(r.id)
}

type IndexRoutesRunner struct {
	filter    []string
	filterPos []string
}

func (r *IndexRoutesRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter

	return c.IndexRoutes(options)
}

type ShowRouteRunner struct {
	id string
}

func (r *ShowRouteRunner) Run(c *Client) (interface{}, error) {
	return c.ShowRoute(r.id)
}

type UpdateRouteRunner struct {
	id                        string
	routeDescription          *string
	routeDestinationCidrBlock *string
	routeNextHopHref          *string
	routeNextHopIp            *string
	routeNextHopType          *string
}

func (r *UpdateRouteRunner) Run(c *Client) (interface{}, error) {

	/** Handle route parameter **/
	var route RouteParam2
	// Load JSON if provided
	if len(r.routeJson) > 0 {
		if err := Json.Unmarshal(r.routeJson, &route); err != nil {
			return nil, fmt.Errorf("Could not load route JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.RouteDescription != nil {
		route.description = r.RouteDescription
	}

	if r.RouteDestinationCidrBlock != nil {
		route.destinationCidrBlock = r.RouteDestinationCidrBlock
	}

	if r.RouteNextHopHref != nil {
		route.nextHopHref = r.RouteNextHopHref
	}

	if r.RouteNextHopIp != nil {
		route.nextHopIp = r.RouteNextHopIp
	}

	if r.RouteNextHopType != nil {
		route.nextHopType = r.RouteNextHopType
	}

	return c.UpdateRoute(r.id, &route)
}

// Register all Route actions
func registerRouteCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Route", `A Route defines how networking traffic should be routed from one destination to another...`)

	CreateRouteRunner := new(CreateRouteRouteRunner)
	CreateRouteCmd := resCmd.Command("CreateRoute", `Create a new Route.`)
	CreateRouteRunner.Flag(`route.description`, `The description to be set on the Route.`).StringVar(CreateRouteRunner.routeDescription)
	CreateRouteRunner.Flag(`route.destinationCidrBlock`, `The destination (CIDR IP address) for the Route.`).Required().StringVar(&CreateRouteRunner.routeDestinationCidrBlock)
	CreateRouteRunner.Flag(`route.nextHopHref`, `The href of the Route's next hop.`).StringVar(CreateRouteRunner.routeNextHopHref)
	CreateRouteRunner.Flag(`route.nextHopIp`, `The IP Address of the Route's next hop. Required if route[next_hop_type] is 'ip_string'. Not allowed otherwise.`).StringVar(CreateRouteRunner.routeNextHopIp)
	CreateRouteRunner.Flag(`route.nextHopType`, `The Route's next hop type.`).Required().StringVar(&CreateRouteRunner.routeNextHopType)
	CreateRouteRunner.Flag(`route.routeTableHref`, `The RouteTable to create the Route in.`).Required().StringVar(&CreateRouteRunner.routeRouteTableHref)
	registry[CreateRouteCmd.FullCommand()] = CreateRouteRunner

	DestroyRouteRunner := new(DestroyRouteRouteRunner)
	DestroyRouteCmd := resCmd.Command("DestroyRoute", `Delete an existing Route.`)
	DestroyRouteRunner.Flag(`id`, ``).Required().StringVar(&DestroyRouteRunner.id)
	registry[DestroyRouteCmd.FullCommand()] = DestroyRouteRunner

	IndexRoutesRunner := new(IndexRoutesRouteRunner)
	IndexRoutesCmd := resCmd.Command("IndexRoutes", `List Routes available in this account.`)
	IndexRoutesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexRoutesRunner.filterPos).StringsVar(IndexRoutesRunner.filter)
	registry[IndexRoutesCmd.FullCommand()] = IndexRoutesRunner

	ShowRouteRunner := new(ShowRouteRouteRunner)
	ShowRouteCmd := resCmd.Command("ShowRoute", `Show information about a single Route.`)
	ShowRouteRunner.Flag(`id`, ``).Required().StringVar(&ShowRouteRunner.id)
	registry[ShowRouteCmd.FullCommand()] = ShowRouteRunner

	UpdateRouteRunner := new(UpdateRouteRouteRunner)
	UpdateRouteCmd := resCmd.Command("UpdateRoute", `Update an existing Route.`)
	UpdateRouteRunner.Flag(`id`, ``).Required().StringVar(&UpdateRouteRunner.id)
	UpdateRouteRunner.Flag(`route.description`, `The updated description of the Route.`).StringVar(UpdateRouteRunner.routeDescription)
	UpdateRouteRunner.Flag(`route.destinationCidrBlock`, `The updated destination (CIDR IP address) for the Route.`).StringVar(UpdateRouteRunner.routeDestinationCidrBlock)
	UpdateRouteRunner.Flag(`route.nextHopHref`, `The updated href of the Route's next hop. Required if route[next_hop_type] is 'instance', 'network_interface', or 'network_gateway'. Not allowed otherwise.`).StringVar(UpdateRouteRunner.routeNextHopHref)
	UpdateRouteRunner.Flag(`route.nextHopIp`, `The updated IP Address of the Route's next hop. Required if route[next_hop_type] is 'ip_string'. Not allowed otherwise.`).StringVar(UpdateRouteRunner.routeNextHopIp)
	UpdateRouteRunner.Flag(`route.nextHopType`, `The updated Route's next hop type.`).StringVar(UpdateRouteRunner.routeNextHopType)
	registry[UpdateRouteCmd.FullCommand()] = UpdateRouteRunner
}

/****** RouteTable ******/

type CreateRouteTableRunner struct {
	routeTableCloudHref   string
	routeTableDescription *string
	routeTableName        string
	routeTableNetworkHref string
}

func (r *CreateRouteTableRunner) Run(c *Client) (interface{}, error) {

	/** Handle routeTable parameter **/
	var routeTable RouteTableParam
	// Load JSON if provided
	if len(r.routeTableJson) > 0 {
		if err := Json.Unmarshal(r.routeTableJson, &routeTable); err != nil {
			return nil, fmt.Errorf("Could not load routeTable JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.RouteTableCloudHref) > 0 {
		routeTable.cloudHref = r.RouteTableCloudHref
	}

	if r.RouteTableDescription != nil {
		routeTable.description = r.RouteTableDescription
	}

	if len(r.RouteTableName) > 0 {
		routeTable.name = r.RouteTableName
	}

	if len(r.RouteTableNetworkHref) > 0 {
		routeTable.networkHref = r.RouteTableNetworkHref
	}

	return c.CreateRouteTable(&routeTable)
}

type DestroyRouteTableRunner struct {
	id string
}

func (r *DestroyRouteTableRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyRouteTable(r.id)
}

type IndexRouteTablesRunner struct {
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexRouteTablesRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexRouteTables(options)
}

type ShowRouteTableRunner struct {
	id   string
	view *string
}

func (r *ShowRouteTableRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowRouteTable(r.id, options)
}

type UpdateRouteTableRunner struct {
	id                    string
	routeTableDescription *string
	routeTableName        *string
}

func (r *UpdateRouteTableRunner) Run(c *Client) (interface{}, error) {

	/** Handle routeTable parameter **/
	var routeTable RouteTableParam2
	// Load JSON if provided
	if len(r.routeTableJson) > 0 {
		if err := Json.Unmarshal(r.routeTableJson, &routeTable); err != nil {
			return nil, fmt.Errorf("Could not load routeTable JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.RouteTableDescription != nil {
		routeTable.description = r.RouteTableDescription
	}

	if r.RouteTableName != nil {
		routeTable.name = r.RouteTableName
	}

	return c.UpdateRouteTable(r.id, &routeTable)
}

// Register all RouteTable actions
func registerRouteTableCmds(app *kingpin.Application) {
	resCmd := app.Cmd("RouteTable", `Grouped listing of Routes`)

	CreateRouteTableRunner := new(CreateRouteTableRouteTableRunner)
	CreateRouteTableCmd := resCmd.Command("CreateRouteTable", `Create a new RouteTable.`)
	CreateRouteTableRunner.Flag(`routeTable.cloudHref`, `The cloud to create the RouteTable in.`).Required().StringVar(&CreateRouteTableRunner.routeTableCloudHref)
	CreateRouteTableRunner.Flag(`routeTable.description`, `The description to be set on the RouteTable.`).StringVar(CreateRouteTableRunner.routeTableDescription)
	CreateRouteTableRunner.Flag(`routeTable.name`, `The name to be set on the RouteTable.`).Required().StringVar(&CreateRouteTableRunner.routeTableName)
	CreateRouteTableRunner.Flag(`routeTable.networkHref`, `The Network to create the RouteTable in.`).Required().StringVar(&CreateRouteTableRunner.routeTableNetworkHref)
	registry[CreateRouteTableCmd.FullCommand()] = CreateRouteTableRunner

	DestroyRouteTableRunner := new(DestroyRouteTableRouteTableRunner)
	DestroyRouteTableCmd := resCmd.Command("DestroyRouteTable", `Delete an existing RouteTable.`)
	DestroyRouteTableRunner.Flag(`id`, ``).Required().StringVar(&DestroyRouteTableRunner.id)
	registry[DestroyRouteTableCmd.FullCommand()] = DestroyRouteTableRunner

	IndexRouteTablesRunner := new(IndexRouteTablesRouteTableRunner)
	IndexRouteTablesCmd := resCmd.Command("IndexRouteTables", `List RouteTables available in this account.`)
	IndexRouteTablesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexRouteTablesRunner.filterPos).StringsVar(IndexRouteTablesRunner.filter)
	IndexRouteTablesRunner.Flag(`view`, ``).StringVar(IndexRouteTablesRunner.view)
	registry[IndexRouteTablesCmd.FullCommand()] = IndexRouteTablesRunner

	ShowRouteTableRunner := new(ShowRouteTableRouteTableRunner)
	ShowRouteTableCmd := resCmd.Command("ShowRouteTable", `Show information about a single RouteTable.`)
	ShowRouteTableRunner.Flag(`id`, ``).Required().StringVar(&ShowRouteTableRunner.id)
	ShowRouteTableRunner.Flag(`view`, ``).StringVar(ShowRouteTableRunner.view)
	registry[ShowRouteTableCmd.FullCommand()] = ShowRouteTableRunner

	UpdateRouteTableRunner := new(UpdateRouteTableRouteTableRunner)
	UpdateRouteTableCmd := resCmd.Command("UpdateRouteTable", `Update an existing RouteTable.`)
	UpdateRouteTableRunner.Flag(`id`, ``).Required().StringVar(&UpdateRouteTableRunner.id)
	UpdateRouteTableRunner.Flag(`routeTable.description`, `The description to be set on the RouteTable.`).StringVar(UpdateRouteTableRunner.routeTableDescription)
	UpdateRouteTableRunner.Flag(`routeTable.name`, `The name to be set on the RouteTable.`).StringVar(UpdateRouteTableRunner.routeTableName)
	registry[UpdateRouteTableCmd.FullCommand()] = UpdateRouteTableRunner
}

/****** RunnableBinding ******/

type CreateRunnableBindingRunner struct {
	runnableBindingPosition        *string
	runnableBindingRecipe          *string
	runnableBindingRightScriptHref *string
	runnableBindingSequence        *string
	serverTemplateId               string
}

func (r *CreateRunnableBindingRunner) Run(c *Client) (interface{}, error) {

	/** Handle runnableBinding parameter **/
	var runnableBinding RunnableBindingParam
	// Load JSON if provided
	if len(r.runnableBindingJson) > 0 {
		if err := Json.Unmarshal(r.runnableBindingJson, &runnableBinding); err != nil {
			return nil, fmt.Errorf("Could not load runnableBinding JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.RunnableBindingPosition != nil {
		runnableBinding.position = r.RunnableBindingPosition
	}

	if r.RunnableBindingRecipe != nil {
		runnableBinding.recipe = r.RunnableBindingRecipe
	}

	if r.RunnableBindingRightScriptHref != nil {
		runnableBinding.rightScriptHref = r.RunnableBindingRightScriptHref
	}

	if r.RunnableBindingSequence != nil {
		runnableBinding.sequence = r.RunnableBindingSequence
	}

	return c.CreateRunnableBinding(&runnableBinding, r.serverTemplateId)
}

type DestroyRunnableBindingRunner struct {
	id               string
	serverTemplateId string
}

func (r *DestroyRunnableBindingRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyRunnableBinding(r.id, r.serverTemplateId)
}

type IndexRunnableBindingsRunner struct {
	serverTemplateId string
	view             *string
}

func (r *IndexRunnableBindingsRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexRunnableBindings(r.serverTemplateId, options)
}

type MultiUpdateRunnableBindingsRunner struct {
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

func (r *MultiUpdateRunnableBindingsRunner) Run(c *Client) (interface{}, error) {

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

type ShowRunnableBindingRunner struct {
	id               string
	serverTemplateId string
	view             *string
}

func (r *ShowRunnableBindingRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowRunnableBinding(r.id, r.serverTemplateId, options)
}

// Register all RunnableBinding actions
func registerRunnableBindingCmds(app *kingpin.Application) {
	resCmd := app.Cmd("RunnableBinding", `A RunnableBinding represents an item in a runlist of a ServerTemplate`)

	CreateRunnableBindingRunner := new(CreateRunnableBindingRunnableBindingRunner)
	CreateRunnableBindingCmd := resCmd.Command("CreateRunnableBinding", `Bind an executable to the given ServerTemplate.`)
	CreateRunnableBindingRunner.Flag(`runnableBinding.position`, `The position of the executable in the execution order. If not specified, will be added to the end. If specified, will be inserted in that location and cause all others to move down.`).StringVar(CreateRunnableBindingRunner.runnableBindingPosition)
	CreateRunnableBindingRunner.Flag(`runnableBinding.recipe`, `The Chef recipe name. Note: right_script_href cannot be specified when this param is given.`).StringVar(CreateRunnableBindingRunner.runnableBindingRecipe)
	CreateRunnableBindingRunner.Flag(`runnableBinding.rightScriptHref`, `The RightScript href. Note: recipe cannot be specified when this param is given.`).StringVar(CreateRunnableBindingRunner.runnableBindingRightScriptHref)
	CreateRunnableBindingRunner.Flag(`runnableBinding.sequence`, `The sequence at which this executable should be run. Default is 'operational'.`).StringVar(CreateRunnableBindingRunner.runnableBindingSequence)
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
	IndexRunnableBindingsRunner.Flag(`view`, ``).StringVar(IndexRunnableBindingsRunner.view)
	registry[IndexRunnableBindingsCmd.FullCommand()] = IndexRunnableBindingsRunner

	MultiUpdateRunnableBindingsRunner := new(MultiUpdateRunnableBindingsRunnableBindingRunner)
	MultiUpdateRunnableBindingsCmd := resCmd.Command("MultiUpdateRunnableBindings", `Update attributes for multiple bound executables.`)
	MultiUpdateRunnableBindingsRunner.FlagPattern(`runnableBindings\.(\d+)\.id`, `The ID of the RunnableBinding to update.`).Required().Capture(&MultiUpdateRunnableBindingsRunner.runnableBindingsIdPos).StringsVar(&MultiUpdateRunnableBindingsRunner.runnableBindingsId)
	MultiUpdateRunnableBindingsRunner.FlagPattern(`runnableBindings\.(\d+)\.position`, `The updated position of the RunnableBinding in the execution order. If specified, will be inserted in that location and cause all others to move down.`).Capture(&MultiUpdateRunnableBindingsRunner.runnableBindingsPositionPos).StringsVar(MultiUpdateRunnableBindingsRunner.runnableBindingsPosition)
	MultiUpdateRunnableBindingsRunner.FlagPattern(`runnableBindings\.(\d+)\.recipe`, `The updated Chef recipe name. Note: right_script_href cannot be specified when this param is given.`).Capture(&MultiUpdateRunnableBindingsRunner.runnableBindingsRecipePos).StringsVar(MultiUpdateRunnableBindingsRunner.runnableBindingsRecipe)
	MultiUpdateRunnableBindingsRunner.FlagPattern(`runnableBindings\.(\d+)\.rightScriptHref`, `The updated RightScript href. Note: recipe cannot be specified when this param is given.`).Capture(&MultiUpdateRunnableBindingsRunner.runnableBindingsRightScriptHrefPos).StringsVar(MultiUpdateRunnableBindingsRunner.runnableBindingsRightScriptHref)
	MultiUpdateRunnableBindingsRunner.FlagPattern(`runnableBindings\.(\d+)\.sequence`, `The sequence at which this executable should be run.  Default is 'operational'.`).Capture(&MultiUpdateRunnableBindingsRunner.runnableBindingsSequencePos).StringsVar(MultiUpdateRunnableBindingsRunner.runnableBindingsSequence)
	MultiUpdateRunnableBindingsRunner.Flag(`serverTemplateId`, ``).Required().StringVar(&MultiUpdateRunnableBindingsRunner.serverTemplateId)
	registry[MultiUpdateRunnableBindingsCmd.FullCommand()] = MultiUpdateRunnableBindingsRunner

	ShowRunnableBindingRunner := new(ShowRunnableBindingRunnableBindingRunner)
	ShowRunnableBindingCmd := resCmd.Command("ShowRunnableBinding", `Show information about a single executable binding.`)
	ShowRunnableBindingRunner.Flag(`id`, ``).Required().StringVar(&ShowRunnableBindingRunner.id)
	ShowRunnableBindingRunner.Flag(`serverTemplateId`, ``).Required().StringVar(&ShowRunnableBindingRunner.serverTemplateId)
	ShowRunnableBindingRunner.Flag(`view`, ``).StringVar(ShowRunnableBindingRunner.view)
	registry[ShowRunnableBindingCmd.FullCommand()] = ShowRunnableBindingRunner
}

/****** SecurityGroup ******/

type CreateSecurityGroupRunner struct {
	cloudId                  string
	securityGroupDescription *string
	securityGroupName        string
	securityGroupNetworkHref *string
}

func (r *CreateSecurityGroupRunner) Run(c *Client) (interface{}, error) {

	/** Handle securityGroup parameter **/
	var securityGroup SecurityGroupParam
	// Load JSON if provided
	if len(r.securityGroupJson) > 0 {
		if err := Json.Unmarshal(r.securityGroupJson, &securityGroup); err != nil {
			return nil, fmt.Errorf("Could not load securityGroup JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.SecurityGroupDescription != nil {
		securityGroup.description = r.SecurityGroupDescription
	}

	if len(r.SecurityGroupName) > 0 {
		securityGroup.name = r.SecurityGroupName
	}

	if r.SecurityGroupNetworkHref != nil {
		securityGroup.networkHref = r.SecurityGroupNetworkHref
	}

	return c.CreateSecurityGroup(r.cloudId, &securityGroup)
}

type DestroySecurityGroupRunner struct {
	cloudId string
	id      string
}

func (r *DestroySecurityGroupRunner) Run(c *Client) (interface{}, error) {
	return c.DestroySecurityGroup(r.cloudId, r.id)
}

type IndexSecurityGroupsRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexSecurityGroupsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexSecurityGroups(r.cloudId, options)
}

type ShowSecurityGroupRunner struct {
	cloudId string
	id      string
	view    *string
}

func (r *ShowSecurityGroupRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowSecurityGroup(r.cloudId, r.id, options)
}

// Register all SecurityGroup actions
func registerSecurityGroupCmds(app *kingpin.Application) {
	resCmd := app.Cmd("SecurityGroup", `Security Groups represent network security profiles that contain lists of firewall rules for different ports and source IP addresses, as well as trust relationships amongst different security groups...`)

	CreateSecurityGroupRunner := new(CreateSecurityGroupSecurityGroupRunner)
	CreateSecurityGroupCmd := resCmd.Command("CreateSecurityGroup", `Create a security group.`)
	CreateSecurityGroupRunner.Flag(`cloudId`, ``).Required().StringVar(&CreateSecurityGroupRunner.cloudId)
	CreateSecurityGroupRunner.Flag(`securityGroup.description`, ``).StringVar(CreateSecurityGroupRunner.securityGroupDescription)
	CreateSecurityGroupRunner.Flag(`securityGroup.name`, ``).Required().StringVar(&CreateSecurityGroupRunner.securityGroupName)
	CreateSecurityGroupRunner.Flag(`securityGroup.networkHref`, ``).StringVar(CreateSecurityGroupRunner.securityGroupNetworkHref)
	registry[CreateSecurityGroupCmd.FullCommand()] = CreateSecurityGroupRunner

	DestroySecurityGroupRunner := new(DestroySecurityGroupSecurityGroupRunner)
	DestroySecurityGroupCmd := resCmd.Command("DestroySecurityGroup", `Delete security group(s)`)
	DestroySecurityGroupRunner.Flag(`cloudId`, ``).Required().StringVar(&DestroySecurityGroupRunner.cloudId)
	DestroySecurityGroupRunner.Flag(`id`, ``).Required().StringVar(&DestroySecurityGroupRunner.id)
	registry[DestroySecurityGroupCmd.FullCommand()] = DestroySecurityGroupRunner

	IndexSecurityGroupsRunner := new(IndexSecurityGroupsSecurityGroupRunner)
	IndexSecurityGroupsCmd := resCmd.Command("IndexSecurityGroups", `Lists Security Groups.`)
	IndexSecurityGroupsRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexSecurityGroupsRunner.cloudId)
	IndexSecurityGroupsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexSecurityGroupsRunner.filterPos).StringsVar(IndexSecurityGroupsRunner.filter)
	IndexSecurityGroupsRunner.Flag(`view`, ``).StringVar(IndexSecurityGroupsRunner.view)
	registry[IndexSecurityGroupsCmd.FullCommand()] = IndexSecurityGroupsRunner

	ShowSecurityGroupRunner := new(ShowSecurityGroupSecurityGroupRunner)
	ShowSecurityGroupCmd := resCmd.Command("ShowSecurityGroup", `Displays information about a single Security Group.`)
	ShowSecurityGroupRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowSecurityGroupRunner.cloudId)
	ShowSecurityGroupRunner.Flag(`id`, ``).Required().StringVar(&ShowSecurityGroupRunner.id)
	ShowSecurityGroupRunner.Flag(`view`, ``).StringVar(ShowSecurityGroupRunner.view)
	registry[ShowSecurityGroupCmd.FullCommand()] = ShowSecurityGroupRunner
}

/****** SecurityGroupRule ******/

type CreateSecurityGroupRuleRunner struct {
	securityGroupRuleCidrIps                  *string
	securityGroupRuleDirection                *string
	securityGroupRuleGroupName                *string
	securityGroupRuleGroupOwner               *string
	securityGroupRuleProtocol                 string
	securityGroupRuleProtocolDetailsEndPort   *string
	securityGroupRuleProtocolDetailsIcmpCode  *string
	securityGroupRuleProtocolDetailsIcmpType  *string
	securityGroupRuleProtocolDetailsStartPort *string
	securityGroupRuleSecurityGroupHref        *string
	securityGroupRuleSourceType               string
}

func (r *CreateSecurityGroupRuleRunner) Run(c *Client) (interface{}, error) {

	/** Handle securityGroupRule parameter **/
	var securityGroupRule SecurityGroupRuleParam
	// Load JSON if provided
	if len(r.securityGroupRuleJson) > 0 {
		if err := Json.Unmarshal(r.securityGroupRuleJson, &securityGroupRule); err != nil {
			return nil, fmt.Errorf("Could not load securityGroupRule JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.SecurityGroupRuleCidrIps != nil {
		securityGroupRule.cidrIps = r.SecurityGroupRuleCidrIps
	}

	if r.SecurityGroupRuleDirection != nil {
		securityGroupRule.direction = r.SecurityGroupRuleDirection
	}

	if r.SecurityGroupRuleGroupName != nil {
		securityGroupRule.groupName = r.SecurityGroupRuleGroupName
	}

	if r.SecurityGroupRuleGroupOwner != nil {
		securityGroupRule.groupOwner = r.SecurityGroupRuleGroupOwner
	}

	if len(r.SecurityGroupRuleProtocol) > 0 {
		securityGroupRule.protocol = r.SecurityGroupRuleProtocol
	}

	if r.SecurityGroupRuleProtocolDetailsEndPort != nil {
		securityGroupRule.protocolDetails.endPort = r.SecurityGroupRuleProtocolDetailsEndPort
	}

	if r.SecurityGroupRuleProtocolDetailsIcmpCode != nil {
		securityGroupRule.protocolDetails.icmpCode = r.SecurityGroupRuleProtocolDetailsIcmpCode
	}

	if r.SecurityGroupRuleProtocolDetailsIcmpType != nil {
		securityGroupRule.protocolDetails.icmpType = r.SecurityGroupRuleProtocolDetailsIcmpType
	}

	if r.SecurityGroupRuleProtocolDetailsStartPort != nil {
		securityGroupRule.protocolDetails.startPort = r.SecurityGroupRuleProtocolDetailsStartPort
	}

	if r.SecurityGroupRuleSecurityGroupHref != nil {
		securityGroupRule.securityGroupHref = r.SecurityGroupRuleSecurityGroupHref
	}

	if len(r.SecurityGroupRuleSourceType) > 0 {
		securityGroupRule.sourceType = r.SecurityGroupRuleSourceType
	}

	return c.CreateSecurityGroupRule(&securityGroupRule)
}

type DestroySecurityGroupRuleRunner struct {
	id string
}

func (r *DestroySecurityGroupRuleRunner) Run(c *Client) (interface{}, error) {
	return c.DestroySecurityGroupRule(r.id)
}

type IndexSecurityGroupRulesRunner struct {
	view *string
}

func (r *IndexSecurityGroupRulesRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexSecurityGroupRules(options)
}

type ShowSecurityGroupRuleRunner struct {
	id   string
	view *string
}

func (r *ShowSecurityGroupRuleRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowSecurityGroupRule(r.id, options)
}

type UpdateSecurityGroupRuleRunner struct {
	id                           string
	securityGroupRuleDescription *string
}

func (r *UpdateSecurityGroupRuleRunner) Run(c *Client) (interface{}, error) {

	/** Handle securityGroupRule parameter **/
	var securityGroupRule SecurityGroupRuleParam2
	// Load JSON if provided
	if len(r.securityGroupRuleJson) > 0 {
		if err := Json.Unmarshal(r.securityGroupRuleJson, &securityGroupRule); err != nil {
			return nil, fmt.Errorf("Could not load securityGroupRule JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.SecurityGroupRuleDescription != nil {
		securityGroupRule.description = r.SecurityGroupRuleDescription
	}

	return c.UpdateSecurityGroupRule(r.id, &securityGroupRule)
}

// Register all SecurityGroupRule actions
func registerSecurityGroupRuleCmds(app *kingpin.Application) {
	resCmd := app.Cmd("SecurityGroupRule", ``)

	CreateSecurityGroupRuleRunner := new(CreateSecurityGroupRuleSecurityGroupRuleRunner)
	CreateSecurityGroupRuleCmd := resCmd.Command("CreateSecurityGroupRule", `Create a security group rule for a security group.`)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.cidrIps`, `An IP address range in CIDR notation. Required if source_type is 'cidr_ips'.`).StringVar(CreateSecurityGroupRuleRunner.securityGroupRuleCidrIps)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.direction`, `Direction of traffic.`).StringVar(CreateSecurityGroupRuleRunner.securityGroupRuleDirection)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.groupName`, `Name of source Security Group. Required if source_type is 'group'.`).StringVar(CreateSecurityGroupRuleRunner.securityGroupRuleGroupName)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.groupOwner`, `Owner of source Security Group. Required if source_type is 'group'.`).StringVar(CreateSecurityGroupRuleRunner.securityGroupRuleGroupOwner)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.protocol`, `Protocol to filter on.`).Required().StringVar(&CreateSecurityGroupRuleRunner.securityGroupRuleProtocol)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.protocolDetails.endPort`, `End of port range (inclusive). Required if protocol is 'tcp' or 'udp'.`).StringVar(CreateSecurityGroupRuleRunner.securityGroupRuleProtocolDetailsEndPort)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.protocolDetails.icmpCode`, `ICMP code. Required if protocol is 'icmp'.`).StringVar(CreateSecurityGroupRuleRunner.securityGroupRuleProtocolDetailsIcmpCode)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.protocolDetails.icmpType`, `ICMP type. Required if protocol is 'icmp'.`).StringVar(CreateSecurityGroupRuleRunner.securityGroupRuleProtocolDetailsIcmpType)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.protocolDetails.startPort`, `Start of port range (inclusive). Required if protocol is 'tcp' or 'udp'.`).StringVar(CreateSecurityGroupRuleRunner.securityGroupRuleProtocolDetailsStartPort)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.securityGroupHref`, `Security Group to add rule to.`).StringVar(CreateSecurityGroupRuleRunner.securityGroupRuleSecurityGroupHref)
	CreateSecurityGroupRuleRunner.Flag(`securityGroupRule.sourceType`, `Source type. May be a CIDR block or another Security Group.`).Required().StringVar(&CreateSecurityGroupRuleRunner.securityGroupRuleSourceType)
	registry[CreateSecurityGroupRuleCmd.FullCommand()] = CreateSecurityGroupRuleRunner

	DestroySecurityGroupRuleRunner := new(DestroySecurityGroupRuleSecurityGroupRuleRunner)
	DestroySecurityGroupRuleCmd := resCmd.Command("DestroySecurityGroupRule", `Delete security group rule(s)`)
	DestroySecurityGroupRuleRunner.Flag(`id`, ``).Required().StringVar(&DestroySecurityGroupRuleRunner.id)
	registry[DestroySecurityGroupRuleCmd.FullCommand()] = DestroySecurityGroupRuleRunner

	IndexSecurityGroupRulesRunner := new(IndexSecurityGroupRulesSecurityGroupRuleRunner)
	IndexSecurityGroupRulesCmd := resCmd.Command("IndexSecurityGroupRules", `Lists SecurityGroupRules.`)
	IndexSecurityGroupRulesRunner.Flag(`view`, ``).StringVar(IndexSecurityGroupRulesRunner.view)
	registry[IndexSecurityGroupRulesCmd.FullCommand()] = IndexSecurityGroupRulesRunner

	ShowSecurityGroupRuleRunner := new(ShowSecurityGroupRuleSecurityGroupRuleRunner)
	ShowSecurityGroupRuleCmd := resCmd.Command("ShowSecurityGroupRule", `Displays information about a single SecurityGroupRule.`)
	ShowSecurityGroupRuleRunner.Flag(`id`, ``).Required().StringVar(&ShowSecurityGroupRuleRunner.id)
	ShowSecurityGroupRuleRunner.Flag(`view`, ``).StringVar(ShowSecurityGroupRuleRunner.view)
	registry[ShowSecurityGroupRuleCmd.FullCommand()] = ShowSecurityGroupRuleRunner

	UpdateSecurityGroupRuleRunner := new(UpdateSecurityGroupRuleSecurityGroupRuleRunner)
	UpdateSecurityGroupRuleCmd := resCmd.Command("UpdateSecurityGroupRule", ``)
	UpdateSecurityGroupRuleRunner.Flag(`id`, ``).Required().StringVar(&UpdateSecurityGroupRuleRunner.id)
	UpdateSecurityGroupRuleRunner.Flag(`securityGroupRule.description`, ``).StringVar(UpdateSecurityGroupRuleRunner.securityGroupRuleDescription)
	registry[UpdateSecurityGroupRuleCmd.FullCommand()] = UpdateSecurityGroupRuleRunner
}

/****** Server ******/

type CloneServerRunner struct {
	id string
}

func (r *CloneServerRunner) Run(c *Client) (interface{}, error) {
	return c.CloneServer(r.id)
}

type CreateServerRunner struct {
	serverDeploymentHref                                               *string
	serverDescription                                                  *string
	serverInstanceAssociatePublicIpAddress                             *string
	serverInstanceCloudHref                                            string
	serverInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping *string
	serverInstanceCloudSpecificAttributesIamInstanceProfile            *string
	serverInstanceCloudSpecificAttributesRootVolumePerformance         *string
	serverInstanceCloudSpecificAttributesRootVolumeSize                *string
	serverInstanceCloudSpecificAttributesRootVolumeTypeUid             *string
	serverInstanceDatacenterHref                                       *string
	serverInstanceImageHref                                            *string
	serverInstanceInputsValues                                         []string
	serverInstanceInputsNames                                          []string
	serverInstanceInstanceTypeHref                                     *string
	serverInstanceIpForwardingEnabled                                  *string
	serverInstanceKernelImageHref                                      *string
	serverInstanceMultiCloudImageHref                                  *string
	serverInstancePlacementGroupHref                                   *string
	serverInstanceRamdiskImageHref                                     *string
	serverInstanceSecurityGroupHrefs                                   []string
	serverInstanceSecurityGroupHrefsPos                                []string
	serverInstanceServerTemplateHref                                   string
	serverInstanceSshKeyHref                                           *string
	serverInstanceSubnetHrefs                                          []string
	serverInstanceSubnetHrefsPos                                       []string
	serverInstanceUserData                                             *string
	serverName                                                         string
	serverOptimized                                                    *string
}

func (r *CreateServerRunner) Run(c *Client) (interface{}, error) {

	/** Handle server parameter **/
	var server ServerParam
	// Load JSON if provided
	if len(r.serverJson) > 0 {
		if err := Json.Unmarshal(r.serverJson, &server); err != nil {
			return nil, fmt.Errorf("Could not load server JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.ServerDeploymentHref != nil {
		server.deploymentHref = r.ServerDeploymentHref
	}

	if r.ServerDescription != nil {
		server.description = r.ServerDescription
	}

	if r.ServerInstanceAssociatePublicIpAddress != nil {
		server.instance.associatePublicIpAddress = r.ServerInstanceAssociatePublicIpAddress
	}

	if len(r.ServerInstanceCloudHref) > 0 {
		server.instance.cloudHref = r.ServerInstanceCloudHref
	}

	if r.ServerInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping != nil {
		server.instance.cloudSpecificAttributes.automaticInstanceStoreMapping = r.ServerInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping
	}

	if r.ServerInstanceCloudSpecificAttributesIamInstanceProfile != nil {
		server.instance.cloudSpecificAttributes.iamInstanceProfile = r.ServerInstanceCloudSpecificAttributesIamInstanceProfile
	}

	if r.ServerInstanceCloudSpecificAttributesRootVolumePerformance != nil {
		server.instance.cloudSpecificAttributes.rootVolumePerformance = r.ServerInstanceCloudSpecificAttributesRootVolumePerformance
	}

	if r.ServerInstanceCloudSpecificAttributesRootVolumeSize != nil {
		server.instance.cloudSpecificAttributes.rootVolumeSize = r.ServerInstanceCloudSpecificAttributesRootVolumeSize
	}

	if r.ServerInstanceCloudSpecificAttributesRootVolumeTypeUid != nil {
		server.instance.cloudSpecificAttributes.rootVolumeTypeUid = r.ServerInstanceCloudSpecificAttributesRootVolumeTypeUid
	}

	if r.ServerInstanceDatacenterHref != nil {
		server.instance.datacenterHref = r.ServerInstanceDatacenterHref
	}

	if r.ServerInstanceImageHref != nil {
		server.instance.imageHref = r.ServerInstanceImageHref
	}

	if r.ServerInstanceInstanceTypeHref != nil {
		server.instance.instanceTypeHref = r.ServerInstanceInstanceTypeHref
	}

	if r.ServerInstanceIpForwardingEnabled != nil {
		server.instance.ipForwardingEnabled = r.ServerInstanceIpForwardingEnabled
	}

	if r.ServerInstanceKernelImageHref != nil {
		server.instance.kernelImageHref = r.ServerInstanceKernelImageHref
	}

	if r.ServerInstanceMultiCloudImageHref != nil {
		server.instance.multiCloudImageHref = r.ServerInstanceMultiCloudImageHref
	}

	if r.ServerInstancePlacementGroupHref != nil {
		server.instance.placementGroupHref = r.ServerInstancePlacementGroupHref
	}

	if r.ServerInstanceRamdiskImageHref != nil {
		server.instance.ramdiskImageHref = r.ServerInstanceRamdiskImageHref
	}

	if len(r.ServerInstanceServerTemplateHref) > 0 {
		server.instance.serverTemplateHref = r.ServerInstanceServerTemplateHref
	}

	if r.ServerInstanceSshKeyHref != nil {
		server.instance.sshKeyHref = r.ServerInstanceSshKeyHref
	}

	if r.ServerInstanceUserData != nil {
		server.instance.userData = r.ServerInstanceUserData
	}

	if len(r.ServerName) > 0 {
		server.name = r.ServerName
	}

	if r.ServerOptimized != nil {
		server.optimized = r.ServerOptimized
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxServerInstanceSecurityGroupHrefsPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.ServerInstanceSecurityGroupHrefsPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for server.instance.securityGroupHrefs field of securityGroupHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of server.instance.securityGroupHrefs field of securityGroupHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerInstanceSecurityGroupHrefsPos {
			maxServerInstanceSecurityGroupHrefsPos = pos
		}
	}
	if len(r.ServerInstanceSecurityGroupHrefs) != maxServerInstanceSecurityGroupHrefsPos {
		return nil, fmt.Errof("Invalid flags for server.instance.securityGroupHrefs field of securityGroupHrefs array, %s were defined but max position is %s",
			len(r.ServerInstanceSecurityGroupHrefs), maxServerInstanceSecurityGroupHrefsPos)
	}
	if maxServerInstanceSecurityGroupHrefsPos > -1 {
		serverInstanceSecurityGroupHrefs := make([]*string, maxServerInstanceSecurityGroupHrefsPos+1)
		for i := 0; i < maxServerInstanceSecurityGroupHrefsPos+1; i++ {
			serverInstanceSecurityGroupHrefs[i] = r.serverInstanceSecurityGroupHrefs[r.serverInstanceSecurityGroupHrefsPos[i]]
		}
		server.instance.securityGroupHrefs = serverInstanceSecurityGroupHrefs
	}

	maxServerInstanceSubnetHrefsPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.ServerInstanceSubnetHrefsPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for server.instance.subnetHrefs field of subnetHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of server.instance.subnetHrefs field of subnetHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerInstanceSubnetHrefsPos {
			maxServerInstanceSubnetHrefsPos = pos
		}
	}
	if len(r.ServerInstanceSubnetHrefs) != maxServerInstanceSubnetHrefsPos {
		return nil, fmt.Errof("Invalid flags for server.instance.subnetHrefs field of subnetHrefs array, %s were defined but max position is %s",
			len(r.ServerInstanceSubnetHrefs), maxServerInstanceSubnetHrefsPos)
	}
	if maxServerInstanceSubnetHrefsPos > -1 {
		serverInstanceSubnetHrefs := make([]*string, maxServerInstanceSubnetHrefsPos+1)
		for i := 0; i < maxServerInstanceSubnetHrefsPos+1; i++ {
			serverInstanceSubnetHrefs[i] = r.serverInstanceSubnetHrefs[r.serverInstanceSubnetHrefsPos[i]]
		}
		server.instance.subnetHrefs = serverInstanceSubnetHrefs
	}

	// Create enumarable fields from flags
	serverInstanceInputs := make(map[string]string, len(r.ServerInstanceInputsNames))
	for i, n := range r.ServerInstanceInputsNames {
		serverInstanceInputs[n] = r.ServerInstanceInputsValues[i]
	}
	server.server.instance.inputs = serverInstanceInputs

	return c.CreateServer(&server)
}

type DestroyServerRunner struct {
	id string
}

func (r *DestroyServerRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyServer(r.id)
}

type IndexServersRunner struct {
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexServersRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexServers(options)
}

type LaunchServerRunner struct {
	id string
}

func (r *LaunchServerRunner) Run(c *Client) (interface{}, error) {
	return c.LaunchServer(r.id)
}

type ShowServerRunner struct {
	id   string
	view *string
}

func (r *ShowServerRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowServer(r.id, options)
}

type TerminateServerRunner struct {
	id string
}

func (r *TerminateServerRunner) Run(c *Client) (interface{}, error) {
	return c.TerminateServer(r.id)
}

type UpdateServerRunner struct {
	id                                  string
	serverAutomaticInstanceStoreMapping *string
	serverDescription                   *string
	serverName                          *string
	serverOptimized                     *string
	serverRootVolumeSize                *string
}

func (r *UpdateServerRunner) Run(c *Client) (interface{}, error) {

	/** Handle server parameter **/
	var server ServerParam2
	// Load JSON if provided
	if len(r.serverJson) > 0 {
		if err := Json.Unmarshal(r.serverJson, &server); err != nil {
			return nil, fmt.Errorf("Could not load server JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.ServerAutomaticInstanceStoreMapping != nil {
		server.automaticInstanceStoreMapping = r.ServerAutomaticInstanceStoreMapping
	}

	if r.ServerDescription != nil {
		server.description = r.ServerDescription
	}

	if r.ServerName != nil {
		server.name = r.ServerName
	}

	if r.ServerOptimized != nil {
		server.optimized = r.ServerOptimized
	}

	if r.ServerRootVolumeSize != nil {
		server.rootVolumeSize = r.ServerRootVolumeSize
	}

	return c.UpdateServer(r.id, &server)
}

type WrapInstanceServerRunner struct {
	serverDeploymentHref              *string
	serverDescription                 *string
	serverInstanceHref                string
	serverInstanceInputsValues        []string
	serverInstanceInputsNames         []string
	serverInstanceMultiCloudImageHref *string
	serverInstanceServerTemplateHref  string
	serverName                        *string
}

func (r *WrapInstanceServerRunner) Run(c *Client) (interface{}, error) {

	/** Handle server parameter **/
	var server ServerParam2
	// Load JSON if provided
	if len(r.serverJson) > 0 {
		if err := Json.Unmarshal(r.serverJson, &server); err != nil {
			return nil, fmt.Errorf("Could not load server JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.ServerDeploymentHref != nil {
		server.deploymentHref = r.ServerDeploymentHref
	}

	if r.ServerDescription != nil {
		server.description = r.ServerDescription
	}

	if len(r.ServerInstanceHref) > 0 {
		server.instance.href = r.ServerInstanceHref
	}

	if r.ServerInstanceMultiCloudImageHref != nil {
		server.instance.multiCloudImageHref = r.ServerInstanceMultiCloudImageHref
	}

	if len(r.ServerInstanceServerTemplateHref) > 0 {
		server.instance.serverTemplateHref = r.ServerInstanceServerTemplateHref
	}

	if r.ServerName != nil {
		server.name = r.ServerName
	}

	// Create enumarable fields from flags
	serverInstanceInputs := make(map[string]string, len(r.ServerInstanceInputsNames))
	for i, n := range r.ServerInstanceInputsNames {
		serverInstanceInputs[n] = r.ServerInstanceInputsValues[i]
	}
	server.server.instance.inputs = serverInstanceInputs

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
	CreateServerRunner.Flag(`server.deploymentHref`, `The href of the deployment to which the Server will be added.`).StringVar(CreateServerRunner.serverDeploymentHref)
	CreateServerRunner.Flag(`server.description`, `The Server description.`).StringVar(CreateServerRunner.serverDescription)
	CreateServerRunner.Flag(`server.instance.associatePublicIpAddress`, `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`).StringVar(CreateServerRunner.serverInstanceAssociatePublicIpAddress)
	CreateServerRunner.Flag(`server.instance.cloudHref`, `The href of the cloud that the Server should be added to.`).Required().StringVar(&CreateServerRunner.serverInstanceCloudHref)
	CreateServerRunner.Flag(`server.instance.cloudSpecificAttributes.automaticInstanceStoreMapping`, `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`).StringVar(CreateServerRunner.serverInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping)
	CreateServerRunner.Flag(`server.instance.cloudSpecificAttributes.iamInstanceProfile`, `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`).StringVar(CreateServerRunner.serverInstanceCloudSpecificAttributesIamInstanceProfile)
	CreateServerRunner.Flag(`server.instance.cloudSpecificAttributes.rootVolumePerformance`, `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`).StringVar(CreateServerRunner.serverInstanceCloudSpecificAttributesRootVolumePerformance)
	CreateServerRunner.Flag(`server.instance.cloudSpecificAttributes.rootVolumeSize`, `The size for root disk. Not supported in all Clouds.`).StringVar(CreateServerRunner.serverInstanceCloudSpecificAttributesRootVolumeSize)
	CreateServerRunner.Flag(`server.instance.cloudSpecificAttributes.rootVolumeTypeUid`, `The type of root volume for instance. Only available on clouds supporting root volume type.`).StringVar(CreateServerRunner.serverInstanceCloudSpecificAttributesRootVolumeTypeUid)
	CreateServerRunner.Flag(`server.instance.datacenterHref`, `The href of the Datacenter / Zone.`).StringVar(CreateServerRunner.serverInstanceDatacenterHref)
	CreateServerRunner.Flag(`server.instance.imageHref`, `The href of the Image to use.`).StringVar(CreateServerRunner.serverInstanceImageHref)
	CreateServerRunner.FlagPattern(`server\.instance\.inputs\.([a-z0-9_]+)`, ``).Capture(&CreateServerRunner.serverInstanceInputsNames).StringVar(CreateServerRunner.serverInstanceInputsValues)
	CreateServerRunner.Flag(`server.instance.instanceTypeHref`, `The href of the Instance Type.`).StringVar(CreateServerRunner.serverInstanceInstanceTypeHref)
	CreateServerRunner.Flag(`server.instance.ipForwardingEnabled`, `Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.`).StringVar(CreateServerRunner.serverInstanceIpForwardingEnabled)
	CreateServerRunner.Flag(`server.instance.kernelImageHref`, `The href of the Kernel Image.`).StringVar(CreateServerRunner.serverInstanceKernelImageHref)
	CreateServerRunner.Flag(`server.instance.multiCloudImageHref`, `The href of the Multi Cloud Image to use.`).StringVar(CreateServerRunner.serverInstanceMultiCloudImageHref)
	CreateServerRunner.Flag(`server.instance.placementGroupHref`, `The href of the Placement Group.`).StringVar(CreateServerRunner.serverInstancePlacementGroupHref)
	CreateServerRunner.Flag(`server.instance.ramdiskImageHref`, `The href of the Ramdisk Image.`).StringVar(CreateServerRunner.serverInstanceRamdiskImageHref)
	CreateServerRunner.FlagPattern(`server\.instance\.securityGroupHrefs\.(\d+)`, `The hrefs of the security groups.`).Capture(&CreateServerRunner.serverInstanceSecurityGroupHrefsPos).StringsVar(CreateServerRunner.serverInstanceSecurityGroupHrefs)
	CreateServerRunner.Flag(`server.instance.serverTemplateHref`, `The href of the Server Template.`).Required().StringVar(&CreateServerRunner.serverInstanceServerTemplateHref)
	CreateServerRunner.Flag(`server.instance.sshKeyHref`, `The href of the SSH key to use.`).StringVar(CreateServerRunner.serverInstanceSshKeyHref)
	CreateServerRunner.FlagPattern(`server\.instance\.subnetHrefs\.(\d+)`, `The hrefs of the updated subnets.`).Capture(&CreateServerRunner.serverInstanceSubnetHrefsPos).StringsVar(CreateServerRunner.serverInstanceSubnetHrefs)
	CreateServerRunner.Flag(`server.instance.userData`, `User data that RightScale automatically passes to your instance at boot time.`).StringVar(CreateServerRunner.serverInstanceUserData)
	CreateServerRunner.Flag(`server.name`, `The name of the Server.`).Required().StringVar(&CreateServerRunner.serverName)
	CreateServerRunner.Flag(`server.optimized`, `A flag indicating whether Instances of this Server should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`).StringVar(CreateServerRunner.serverOptimized)
	registry[CreateServerCmd.FullCommand()] = CreateServerRunner

	DestroyServerRunner := new(DestroyServerServerRunner)
	DestroyServerCmd := resCmd.Command("DestroyServer", `Deletes a given server.`)
	DestroyServerRunner.Flag(`id`, ``).Required().StringVar(&DestroyServerRunner.id)
	registry[DestroyServerCmd.FullCommand()] = DestroyServerRunner

	IndexServersRunner := new(IndexServersServerRunner)
	IndexServersCmd := resCmd.Command("IndexServers", `Lists servers.`)
	IndexServersRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexServersRunner.filterPos).StringsVar(IndexServersRunner.filter)
	IndexServersRunner.Flag(`view`, ``).StringVar(IndexServersRunner.view)
	registry[IndexServersCmd.FullCommand()] = IndexServersRunner

	LaunchServerRunner := new(LaunchServerServerRunner)
	LaunchServerCmd := resCmd.Command("LaunchServer", `Launches the "next" instance of this server`)
	LaunchServerRunner.Flag(`id`, ``).Required().StringVar(&LaunchServerRunner.id)
	registry[LaunchServerCmd.FullCommand()] = LaunchServerRunner

	ShowServerRunner := new(ShowServerServerRunner)
	ShowServerCmd := resCmd.Command("ShowServer", `Shows the information of a single server.`)
	ShowServerRunner.Flag(`id`, ``).Required().StringVar(&ShowServerRunner.id)
	ShowServerRunner.Flag(`view`, ``).StringVar(ShowServerRunner.view)
	registry[ShowServerCmd.FullCommand()] = ShowServerRunner

	TerminateServerRunner := new(TerminateServerServerRunner)
	TerminateServerCmd := resCmd.Command("TerminateServer", `Terminates the current instance of this server`)
	TerminateServerRunner.Flag(`id`, ``).Required().StringVar(&TerminateServerRunner.id)
	registry[TerminateServerCmd.FullCommand()] = TerminateServerRunner

	UpdateServerRunner := new(UpdateServerServerRunner)
	UpdateServerCmd := resCmd.Command("UpdateServer", `Updates attributes of a single server.`)
	UpdateServerRunner.Flag(`id`, ``).Required().StringVar(&UpdateServerRunner.id)
	UpdateServerRunner.Flag(`server.automaticInstanceStoreMapping`, `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`).StringVar(UpdateServerRunner.serverAutomaticInstanceStoreMapping)
	UpdateServerRunner.Flag(`server.description`, `The updated description for the server.`).StringVar(UpdateServerRunner.serverDescription)
	UpdateServerRunner.Flag(`server.name`, `The updated server name.`).StringVar(UpdateServerRunner.serverName)
	UpdateServerRunner.Flag(`server.optimized`, `A flag indicating whether Instances of this Server should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`).StringVar(UpdateServerRunner.serverOptimized)
	UpdateServerRunner.Flag(`server.rootVolumeSize`, `The size for root disk. Not supported in all Clouds.`).StringVar(UpdateServerRunner.serverRootVolumeSize)
	registry[UpdateServerCmd.FullCommand()] = UpdateServerRunner

	WrapInstanceServerRunner := new(WrapInstanceServerServerRunner)
	WrapInstanceServerCmd := resCmd.Command("WrapInstanceServer", `Wrap an existing instance and set current instance for new server`)
	WrapInstanceServerRunner.Flag(`server.deploymentHref`, `The href of the deployment to which the Server will be added.`).StringVar(WrapInstanceServerRunner.serverDeploymentHref)
	WrapInstanceServerRunner.Flag(`server.description`, `The Server description.`).StringVar(WrapInstanceServerRunner.serverDescription)
	WrapInstanceServerRunner.Flag(`server.instance.href`, `The href of the Instance around which the server should be created.`).Required().StringVar(&WrapInstanceServerRunner.serverInstanceHref)
	WrapInstanceServerRunner.FlagPattern(`server\.instance\.inputs\.([a-z0-9_]+)`, ``).Capture(&WrapInstanceServerRunner.serverInstanceInputsNames).StringVar(WrapInstanceServerRunner.serverInstanceInputsValues)
	WrapInstanceServerRunner.Flag(`server.instance.multiCloudImageHref`, `The href of the Multi Cloud Image to use.`).StringVar(WrapInstanceServerRunner.serverInstanceMultiCloudImageHref)
	WrapInstanceServerRunner.Flag(`server.instance.serverTemplateHref`, `The href of the Server Template.`).Required().StringVar(&WrapInstanceServerRunner.serverInstanceServerTemplateHref)
	WrapInstanceServerRunner.Flag(`server.name`, `The name of the Server.`).StringVar(WrapInstanceServerRunner.serverName)
	registry[WrapInstanceServerCmd.FullCommand()] = WrapInstanceServerRunner
}

/****** ServerArray ******/

type CloneServerArrayRunner struct {
	id string
}

func (r *CloneServerArrayRunner) Run(c *Client) (interface{}, error) {
	return c.CloneServerArray(r.id)
}

type CreateServerArrayRunner struct {
	serverArrayArrayType                                                    string
	serverArrayDatacenterPolicyDatacenterHref                               []string
	serverArrayDatacenterPolicyDatacenterHrefPos                            []string
	serverArrayDatacenterPolicyMax                                          []string
	serverArrayDatacenterPolicyMaxPos                                       []string
	serverArrayDatacenterPolicyWeight                                       []string
	serverArrayDatacenterPolicyWeightPos                                    []string
	serverArrayDeploymentHref                                               *string
	serverArrayDescription                                                  *string
	serverArrayElasticityParamsAlertSpecificParamsDecisionThreshold         *string
	serverArrayElasticityParamsAlertSpecificParamsVotersTagPredicate        *string
	serverArrayElasticityParamsBoundsMaxCount                               *string
	serverArrayElasticityParamsBoundsMinCount                               *string
	serverArrayElasticityParamsPacingResizeCalmTime                         *string
	serverArrayElasticityParamsPacingResizeDownBy                           *string
	serverArrayElasticityParamsPacingResizeUpBy                             *string
	serverArrayElasticityParamsQueueSpecificParamsCollectAuditEntries       *string
	serverArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm          *string
	serverArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge             *string
	serverArrayElasticityParamsQueueSpecificParamsItemAgeRegexp             *string
	serverArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance *string
	serverArrayElasticityParamsScheduleDay                                  []string
	serverArrayElasticityParamsScheduleDayPos                               []string
	serverArrayElasticityParamsScheduleMaxCount                             []string
	serverArrayElasticityParamsScheduleMaxCountPos                          []string
	serverArrayElasticityParamsScheduleMinCount                             []string
	serverArrayElasticityParamsScheduleMinCountPos                          []string
	serverArrayElasticityParamsScheduleTime                                 []string
	serverArrayElasticityParamsScheduleTimePos                              []string
	serverArrayInstanceAssociatePublicIpAddress                             *string
	serverArrayInstanceCloudHref                                            string
	serverArrayInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping *string
	serverArrayInstanceCloudSpecificAttributesIamInstanceProfile            *string
	serverArrayInstanceCloudSpecificAttributesRootVolumePerformance         *string
	serverArrayInstanceCloudSpecificAttributesRootVolumeSize                *string
	serverArrayInstanceCloudSpecificAttributesRootVolumeTypeUid             *string
	serverArrayInstanceDatacenterHref                                       *string
	serverArrayInstanceImageHref                                            *string
	serverArrayInstanceInputsValues                                         []string
	serverArrayInstanceInputsNames                                          []string
	serverArrayInstanceInstanceTypeHref                                     *string
	serverArrayInstanceIpForwardingEnabled                                  *string
	serverArrayInstanceKernelImageHref                                      *string
	serverArrayInstanceMultiCloudImageHref                                  *string
	serverArrayInstancePlacementGroupHref                                   *string
	serverArrayInstanceRamdiskImageHref                                     *string
	serverArrayInstanceSecurityGroupHrefs                                   []string
	serverArrayInstanceSecurityGroupHrefsPos                                []string
	serverArrayInstanceServerTemplateHref                                   string
	serverArrayInstanceSshKeyHref                                           *string
	serverArrayInstanceSubnetHrefs                                          []string
	serverArrayInstanceSubnetHrefsPos                                       []string
	serverArrayInstanceUserData                                             *string
	serverArrayName                                                         string
	serverArrayOptimized                                                    *string
	serverArrayState                                                        string
}

func (r *CreateServerArrayRunner) Run(c *Client) (interface{}, error) {

	/** Handle serverArray parameter **/
	var serverArray ServerArrayParam
	// Load JSON if provided
	if len(r.serverArrayJson) > 0 {
		if err := Json.Unmarshal(r.serverArrayJson, &serverArray); err != nil {
			return nil, fmt.Errorf("Could not load serverArray JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.ServerArrayArrayType) > 0 {
		serverArray.arrayType = r.ServerArrayArrayType
	}

	if r.ServerArrayDeploymentHref != nil {
		serverArray.deploymentHref = r.ServerArrayDeploymentHref
	}

	if r.ServerArrayDescription != nil {
		serverArray.description = r.ServerArrayDescription
	}

	if r.ServerArrayElasticityParamsAlertSpecificParamsDecisionThreshold != nil {
		serverArray.elasticityParams.alertSpecificParams.decisionThreshold = r.ServerArrayElasticityParamsAlertSpecificParamsDecisionThreshold
	}

	if r.ServerArrayElasticityParamsAlertSpecificParamsVotersTagPredicate != nil {
		serverArray.elasticityParams.alertSpecificParams.votersTagPredicate = r.ServerArrayElasticityParamsAlertSpecificParamsVotersTagPredicate
	}

	if r.ServerArrayElasticityParamsBoundsMaxCount != nil {
		serverArray.elasticityParams.bounds.maxCount = r.ServerArrayElasticityParamsBoundsMaxCount
	}

	if r.ServerArrayElasticityParamsBoundsMinCount != nil {
		serverArray.elasticityParams.bounds.minCount = r.ServerArrayElasticityParamsBoundsMinCount
	}

	if r.ServerArrayElasticityParamsPacingResizeCalmTime != nil {
		serverArray.elasticityParams.pacing.resizeCalmTime = r.ServerArrayElasticityParamsPacingResizeCalmTime
	}

	if r.ServerArrayElasticityParamsPacingResizeDownBy != nil {
		serverArray.elasticityParams.pacing.resizeDownBy = r.ServerArrayElasticityParamsPacingResizeDownBy
	}

	if r.ServerArrayElasticityParamsPacingResizeUpBy != nil {
		serverArray.elasticityParams.pacing.resizeUpBy = r.ServerArrayElasticityParamsPacingResizeUpBy
	}

	if r.ServerArrayElasticityParamsQueueSpecificParamsCollectAuditEntries != nil {
		serverArray.elasticityParams.queueSpecificParams.collectAuditEntries = r.ServerArrayElasticityParamsQueueSpecificParamsCollectAuditEntries
	}

	if r.ServerArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm != nil {
		serverArray.elasticityParams.queueSpecificParams.itemAge.algorithm = r.ServerArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm
	}

	if r.ServerArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge != nil {
		serverArray.elasticityParams.queueSpecificParams.itemAge.maxAge = r.ServerArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge
	}

	if r.ServerArrayElasticityParamsQueueSpecificParamsItemAgeRegexp != nil {
		serverArray.elasticityParams.queueSpecificParams.itemAge.regexp = r.ServerArrayElasticityParamsQueueSpecificParamsItemAgeRegexp
	}

	if r.ServerArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance != nil {
		serverArray.elasticityParams.queueSpecificParams.queueSize.itemsPerInstance = r.ServerArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance
	}

	if r.ServerArrayInstanceAssociatePublicIpAddress != nil {
		serverArray.instance.associatePublicIpAddress = r.ServerArrayInstanceAssociatePublicIpAddress
	}

	if len(r.ServerArrayInstanceCloudHref) > 0 {
		serverArray.instance.cloudHref = r.ServerArrayInstanceCloudHref
	}

	if r.ServerArrayInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping != nil {
		serverArray.instance.cloudSpecificAttributes.automaticInstanceStoreMapping = r.ServerArrayInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping
	}

	if r.ServerArrayInstanceCloudSpecificAttributesIamInstanceProfile != nil {
		serverArray.instance.cloudSpecificAttributes.iamInstanceProfile = r.ServerArrayInstanceCloudSpecificAttributesIamInstanceProfile
	}

	if r.ServerArrayInstanceCloudSpecificAttributesRootVolumePerformance != nil {
		serverArray.instance.cloudSpecificAttributes.rootVolumePerformance = r.ServerArrayInstanceCloudSpecificAttributesRootVolumePerformance
	}

	if r.ServerArrayInstanceCloudSpecificAttributesRootVolumeSize != nil {
		serverArray.instance.cloudSpecificAttributes.rootVolumeSize = r.ServerArrayInstanceCloudSpecificAttributesRootVolumeSize
	}

	if r.ServerArrayInstanceCloudSpecificAttributesRootVolumeTypeUid != nil {
		serverArray.instance.cloudSpecificAttributes.rootVolumeTypeUid = r.ServerArrayInstanceCloudSpecificAttributesRootVolumeTypeUid
	}

	if r.ServerArrayInstanceDatacenterHref != nil {
		serverArray.instance.datacenterHref = r.ServerArrayInstanceDatacenterHref
	}

	if r.ServerArrayInstanceImageHref != nil {
		serverArray.instance.imageHref = r.ServerArrayInstanceImageHref
	}

	if r.ServerArrayInstanceInstanceTypeHref != nil {
		serverArray.instance.instanceTypeHref = r.ServerArrayInstanceInstanceTypeHref
	}

	if r.ServerArrayInstanceIpForwardingEnabled != nil {
		serverArray.instance.ipForwardingEnabled = r.ServerArrayInstanceIpForwardingEnabled
	}

	if r.ServerArrayInstanceKernelImageHref != nil {
		serverArray.instance.kernelImageHref = r.ServerArrayInstanceKernelImageHref
	}

	if r.ServerArrayInstanceMultiCloudImageHref != nil {
		serverArray.instance.multiCloudImageHref = r.ServerArrayInstanceMultiCloudImageHref
	}

	if r.ServerArrayInstancePlacementGroupHref != nil {
		serverArray.instance.placementGroupHref = r.ServerArrayInstancePlacementGroupHref
	}

	if r.ServerArrayInstanceRamdiskImageHref != nil {
		serverArray.instance.ramdiskImageHref = r.ServerArrayInstanceRamdiskImageHref
	}

	if len(r.ServerArrayInstanceServerTemplateHref) > 0 {
		serverArray.instance.serverTemplateHref = r.ServerArrayInstanceServerTemplateHref
	}

	if r.ServerArrayInstanceSshKeyHref != nil {
		serverArray.instance.sshKeyHref = r.ServerArrayInstanceSshKeyHref
	}

	if r.ServerArrayInstanceUserData != nil {
		serverArray.instance.userData = r.ServerArrayInstanceUserData
	}

	if len(r.ServerArrayName) > 0 {
		serverArray.name = r.ServerArrayName
	}

	if r.ServerArrayOptimized != nil {
		serverArray.optimized = r.ServerArrayOptimized
	}

	if len(r.ServerArrayState) > 0 {
		serverArray.state = r.ServerArrayState
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxServerArrayDatacenterPolicyDatacenterHrefPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.ServerArrayDatacenterPolicyDatacenterHrefPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.datacenterPolicy.datacenterHref field of datacenterPolicy array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.datacenterPolicy.datacenterHref field of datacenterPolicy array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerArrayDatacenterPolicyDatacenterHrefPos {
			maxServerArrayDatacenterPolicyDatacenterHrefPos = pos
		}
	}
	if len(r.ServerArrayDatacenterPolicyDatacenterHref) != maxServerArrayDatacenterPolicyDatacenterHrefPos {
		return nil, fmt.Errof("Invalid flags for serverArray.datacenterPolicy.datacenterHref field of datacenterPolicy array, %s were defined but max position is %s",
			len(r.ServerArrayDatacenterPolicyDatacenterHref), maxServerArrayDatacenterPolicyDatacenterHrefPos)
	}

	maxServerArrayDatacenterPolicyMaxPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.ServerArrayDatacenterPolicyMaxPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.datacenterPolicy.max field of datacenterPolicy array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.datacenterPolicy.max field of datacenterPolicy array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerArrayDatacenterPolicyMaxPos {
			maxServerArrayDatacenterPolicyMaxPos = pos
		}
	}
	if len(r.ServerArrayDatacenterPolicyMax) != maxServerArrayDatacenterPolicyMaxPos {
		return nil, fmt.Errof("Invalid flags for serverArray.datacenterPolicy.max field of datacenterPolicy array, %s were defined but max position is %s",
			len(r.ServerArrayDatacenterPolicyMax), maxServerArrayDatacenterPolicyMaxPos)
	}

	maxServerArrayDatacenterPolicyWeightPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.ServerArrayDatacenterPolicyWeightPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.datacenterPolicy.weight field of datacenterPolicy array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.datacenterPolicy.weight field of datacenterPolicy array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerArrayDatacenterPolicyWeightPos {
			maxServerArrayDatacenterPolicyWeightPos = pos
		}
	}
	if len(r.ServerArrayDatacenterPolicyWeight) != maxServerArrayDatacenterPolicyWeightPos {
		return nil, fmt.Errof("Invalid flags for serverArray.datacenterPolicy.weight field of datacenterPolicy array, %s were defined but max position is %s",
			len(r.ServerArrayDatacenterPolicyWeight), maxServerArrayDatacenterPolicyWeightPos)
	}
	if maxServerArrayDatacenterPolicyDatacenterHrefPos > -1 {
		serverArrayDatacenterPolicy := make([]*DatacenterPolicy, maxServerArrayDatacenterPolicyDatacenterHrefPos+1)
		for i := 0; i < maxServerArrayDatacenterPolicyDatacenterHrefPos+1; i++ {
			serverArrayDatacenterPolicy[i] = *DatacenterPolicy{
			//TBD
			}
		}
		serverArray.datacenterPolicy = serverArrayDatacenterPolicy
	}

	maxServerArrayElasticityParamsScheduleDayPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.ServerArrayElasticityParamsScheduleDayPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.elasticityParams.schedule.day field of schedule array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.elasticityParams.schedule.day field of schedule array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerArrayElasticityParamsScheduleDayPos {
			maxServerArrayElasticityParamsScheduleDayPos = pos
		}
	}
	if len(r.ServerArrayElasticityParamsScheduleDay) != maxServerArrayElasticityParamsScheduleDayPos {
		return nil, fmt.Errof("Invalid flags for serverArray.elasticityParams.schedule.day field of schedule array, %s were defined but max position is %s",
			len(r.ServerArrayElasticityParamsScheduleDay), maxServerArrayElasticityParamsScheduleDayPos)
	}

	maxServerArrayElasticityParamsScheduleMaxCountPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.ServerArrayElasticityParamsScheduleMaxCountPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.elasticityParams.schedule.maxCount field of schedule array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.elasticityParams.schedule.maxCount field of schedule array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerArrayElasticityParamsScheduleMaxCountPos {
			maxServerArrayElasticityParamsScheduleMaxCountPos = pos
		}
	}
	if len(r.ServerArrayElasticityParamsScheduleMaxCount) != maxServerArrayElasticityParamsScheduleMaxCountPos {
		return nil, fmt.Errof("Invalid flags for serverArray.elasticityParams.schedule.maxCount field of schedule array, %s were defined but max position is %s",
			len(r.ServerArrayElasticityParamsScheduleMaxCount), maxServerArrayElasticityParamsScheduleMaxCountPos)
	}

	maxServerArrayElasticityParamsScheduleMinCountPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.ServerArrayElasticityParamsScheduleMinCountPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.elasticityParams.schedule.minCount field of schedule array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.elasticityParams.schedule.minCount field of schedule array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerArrayElasticityParamsScheduleMinCountPos {
			maxServerArrayElasticityParamsScheduleMinCountPos = pos
		}
	}
	if len(r.ServerArrayElasticityParamsScheduleMinCount) != maxServerArrayElasticityParamsScheduleMinCountPos {
		return nil, fmt.Errof("Invalid flags for serverArray.elasticityParams.schedule.minCount field of schedule array, %s were defined but max position is %s",
			len(r.ServerArrayElasticityParamsScheduleMinCount), maxServerArrayElasticityParamsScheduleMinCountPos)
	}

	maxServerArrayElasticityParamsScheduleTimePos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.ServerArrayElasticityParamsScheduleTimePos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.elasticityParams.schedule.time field of schedule array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.elasticityParams.schedule.time field of schedule array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerArrayElasticityParamsScheduleTimePos {
			maxServerArrayElasticityParamsScheduleTimePos = pos
		}
	}
	if len(r.ServerArrayElasticityParamsScheduleTime) != maxServerArrayElasticityParamsScheduleTimePos {
		return nil, fmt.Errof("Invalid flags for serverArray.elasticityParams.schedule.time field of schedule array, %s were defined but max position is %s",
			len(r.ServerArrayElasticityParamsScheduleTime), maxServerArrayElasticityParamsScheduleTimePos)
	}
	if maxServerArrayElasticityParamsScheduleDayPos > -1 {
		serverArrayElasticityParamsSchedule := make([]*Schedule, maxServerArrayElasticityParamsScheduleDayPos+1)
		for i := 0; i < maxServerArrayElasticityParamsScheduleDayPos+1; i++ {
			serverArrayElasticityParamsSchedule[i] = *Schedule{
			//TBD
			}
		}
		serverArray.elasticityParams.schedule = serverArrayElasticityParamsSchedule
	}

	maxServerArrayInstanceSecurityGroupHrefsPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.ServerArrayInstanceSecurityGroupHrefsPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.instance.securityGroupHrefs field of securityGroupHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.instance.securityGroupHrefs field of securityGroupHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerArrayInstanceSecurityGroupHrefsPos {
			maxServerArrayInstanceSecurityGroupHrefsPos = pos
		}
	}
	if len(r.ServerArrayInstanceSecurityGroupHrefs) != maxServerArrayInstanceSecurityGroupHrefsPos {
		return nil, fmt.Errof("Invalid flags for serverArray.instance.securityGroupHrefs field of securityGroupHrefs array, %s were defined but max position is %s",
			len(r.ServerArrayInstanceSecurityGroupHrefs), maxServerArrayInstanceSecurityGroupHrefsPos)
	}
	if maxServerArrayInstanceSecurityGroupHrefsPos > -1 {
		serverArrayInstanceSecurityGroupHrefs := make([]*string, maxServerArrayInstanceSecurityGroupHrefsPos+1)
		for i := 0; i < maxServerArrayInstanceSecurityGroupHrefsPos+1; i++ {
			serverArrayInstanceSecurityGroupHrefs[i] = r.serverArrayInstanceSecurityGroupHrefs[r.serverArrayInstanceSecurityGroupHrefsPos[i]]
		}
		serverArray.instance.securityGroupHrefs = serverArrayInstanceSecurityGroupHrefs
	}

	maxServerArrayInstanceSubnetHrefsPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.ServerArrayInstanceSubnetHrefsPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.instance.subnetHrefs field of subnetHrefs array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.instance.subnetHrefs field of subnetHrefs array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerArrayInstanceSubnetHrefsPos {
			maxServerArrayInstanceSubnetHrefsPos = pos
		}
	}
	if len(r.ServerArrayInstanceSubnetHrefs) != maxServerArrayInstanceSubnetHrefsPos {
		return nil, fmt.Errof("Invalid flags for serverArray.instance.subnetHrefs field of subnetHrefs array, %s were defined but max position is %s",
			len(r.ServerArrayInstanceSubnetHrefs), maxServerArrayInstanceSubnetHrefsPos)
	}
	if maxServerArrayInstanceSubnetHrefsPos > -1 {
		serverArrayInstanceSubnetHrefs := make([]*string, maxServerArrayInstanceSubnetHrefsPos+1)
		for i := 0; i < maxServerArrayInstanceSubnetHrefsPos+1; i++ {
			serverArrayInstanceSubnetHrefs[i] = r.serverArrayInstanceSubnetHrefs[r.serverArrayInstanceSubnetHrefsPos[i]]
		}
		serverArray.instance.subnetHrefs = serverArrayInstanceSubnetHrefs
	}

	// Create enumarable fields from flags
	serverArrayInstanceInputs := make(map[string]string, len(r.ServerArrayInstanceInputsNames))
	for i, n := range r.ServerArrayInstanceInputsNames {
		serverArrayInstanceInputs[n] = r.ServerArrayInstanceInputsValues[i]
	}
	serverArray.serverArray.instance.inputs = serverArrayInstanceInputs

	return c.CreateServerArray(&serverArray)
}

type CurrentInstancesServerArrayRunner struct {
	id string
}

func (r *CurrentInstancesServerArrayRunner) Run(c *Client) (interface{}, error) {
	return c.CurrentInstancesServerArray(r.id)
}

type DestroyServerArrayRunner struct {
	id string
}

func (r *DestroyServerArrayRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyServerArray(r.id)
}

type IndexServerArraysRunner struct {
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexServerArraysRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexServerArrays(options)
}

type LaunchServerArrayRunner struct {
	id string
}

func (r *LaunchServerArrayRunner) Run(c *Client) (interface{}, error) {
	return c.LaunchServerArray(r.id)
}

type MultiRunExecutableServerArraysRunner struct {
	id string
}

func (r *MultiRunExecutableServerArraysRunner) Run(c *Client) (interface{}, error) {
	return c.MultiRunExecutableServerArrays(r.id)
}

type MultiTerminateServerArraysRunner struct {
	id string
}

func (r *MultiTerminateServerArraysRunner) Run(c *Client) (interface{}, error) {
	return c.MultiTerminateServerArrays(r.id)
}

type ShowServerArrayRunner struct {
	id   string
	view *string
}

func (r *ShowServerArrayRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowServerArray(r.id, options)
}

type UpdateServerArrayRunner struct {
	id                                                                      string
	serverArrayArrayType                                                    *string
	serverArrayDatacenterPolicyDatacenterHref                               []string
	serverArrayDatacenterPolicyDatacenterHrefPos                            []string
	serverArrayDatacenterPolicyMax                                          []string
	serverArrayDatacenterPolicyMaxPos                                       []string
	serverArrayDatacenterPolicyWeight                                       []string
	serverArrayDatacenterPolicyWeightPos                                    []string
	serverArrayDeploymentHref                                               *string
	serverArrayDescription                                                  *string
	serverArrayElasticityParamsAlertSpecificParamsDecisionThreshold         *string
	serverArrayElasticityParamsAlertSpecificParamsVotersTagPredicate        *string
	serverArrayElasticityParamsBoundsMaxCount                               *string
	serverArrayElasticityParamsBoundsMinCount                               *string
	serverArrayElasticityParamsPacingResizeCalmTime                         *string
	serverArrayElasticityParamsPacingResizeDownBy                           *string
	serverArrayElasticityParamsPacingResizeUpBy                             *string
	serverArrayElasticityParamsQueueSpecificParamsCollectAuditEntries       *string
	serverArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm          *string
	serverArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge             *string
	serverArrayElasticityParamsQueueSpecificParamsItemAgeRegexp             *string
	serverArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance *string
	serverArrayElasticityParamsScheduleDay                                  []string
	serverArrayElasticityParamsScheduleDayPos                               []string
	serverArrayElasticityParamsScheduleMaxCount                             []string
	serverArrayElasticityParamsScheduleMaxCountPos                          []string
	serverArrayElasticityParamsScheduleMinCount                             []string
	serverArrayElasticityParamsScheduleMinCountPos                          []string
	serverArrayElasticityParamsScheduleTime                                 []string
	serverArrayElasticityParamsScheduleTimePos                              []string
	serverArrayName                                                         *string
	serverArrayOptimized                                                    *string
	serverArrayState                                                        *string
}

func (r *UpdateServerArrayRunner) Run(c *Client) (interface{}, error) {

	/** Handle serverArray parameter **/
	var serverArray ServerArrayParam2
	// Load JSON if provided
	if len(r.serverArrayJson) > 0 {
		if err := Json.Unmarshal(r.serverArrayJson, &serverArray); err != nil {
			return nil, fmt.Errorf("Could not load serverArray JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.ServerArrayArrayType != nil {
		serverArray.arrayType = r.ServerArrayArrayType
	}

	if r.ServerArrayDeploymentHref != nil {
		serverArray.deploymentHref = r.ServerArrayDeploymentHref
	}

	if r.ServerArrayDescription != nil {
		serverArray.description = r.ServerArrayDescription
	}

	if r.ServerArrayElasticityParamsAlertSpecificParamsDecisionThreshold != nil {
		serverArray.elasticityParams.alertSpecificParams.decisionThreshold = r.ServerArrayElasticityParamsAlertSpecificParamsDecisionThreshold
	}

	if r.ServerArrayElasticityParamsAlertSpecificParamsVotersTagPredicate != nil {
		serverArray.elasticityParams.alertSpecificParams.votersTagPredicate = r.ServerArrayElasticityParamsAlertSpecificParamsVotersTagPredicate
	}

	if r.ServerArrayElasticityParamsBoundsMaxCount != nil {
		serverArray.elasticityParams.bounds.maxCount = r.ServerArrayElasticityParamsBoundsMaxCount
	}

	if r.ServerArrayElasticityParamsBoundsMinCount != nil {
		serverArray.elasticityParams.bounds.minCount = r.ServerArrayElasticityParamsBoundsMinCount
	}

	if r.ServerArrayElasticityParamsPacingResizeCalmTime != nil {
		serverArray.elasticityParams.pacing.resizeCalmTime = r.ServerArrayElasticityParamsPacingResizeCalmTime
	}

	if r.ServerArrayElasticityParamsPacingResizeDownBy != nil {
		serverArray.elasticityParams.pacing.resizeDownBy = r.ServerArrayElasticityParamsPacingResizeDownBy
	}

	if r.ServerArrayElasticityParamsPacingResizeUpBy != nil {
		serverArray.elasticityParams.pacing.resizeUpBy = r.ServerArrayElasticityParamsPacingResizeUpBy
	}

	if r.ServerArrayElasticityParamsQueueSpecificParamsCollectAuditEntries != nil {
		serverArray.elasticityParams.queueSpecificParams.collectAuditEntries = r.ServerArrayElasticityParamsQueueSpecificParamsCollectAuditEntries
	}

	if r.ServerArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm != nil {
		serverArray.elasticityParams.queueSpecificParams.itemAge.algorithm = r.ServerArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm
	}

	if r.ServerArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge != nil {
		serverArray.elasticityParams.queueSpecificParams.itemAge.maxAge = r.ServerArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge
	}

	if r.ServerArrayElasticityParamsQueueSpecificParamsItemAgeRegexp != nil {
		serverArray.elasticityParams.queueSpecificParams.itemAge.regexp = r.ServerArrayElasticityParamsQueueSpecificParamsItemAgeRegexp
	}

	if r.ServerArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance != nil {
		serverArray.elasticityParams.queueSpecificParams.queueSize.itemsPerInstance = r.ServerArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance
	}

	if r.ServerArrayName != nil {
		serverArray.name = r.ServerArrayName
	}

	if r.ServerArrayOptimized != nil {
		serverArray.optimized = r.ServerArrayOptimized
	}

	if r.ServerArrayState != nil {
		serverArray.state = r.ServerArrayState
	}

	// Create array fields from flags
	var seenPos map[int]bool
	maxServerArrayDatacenterPolicyDatacenterHrefPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.ServerArrayDatacenterPolicyDatacenterHrefPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.datacenterPolicy.datacenterHref field of datacenterPolicy array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.datacenterPolicy.datacenterHref field of datacenterPolicy array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerArrayDatacenterPolicyDatacenterHrefPos {
			maxServerArrayDatacenterPolicyDatacenterHrefPos = pos
		}
	}
	if len(r.ServerArrayDatacenterPolicyDatacenterHref) != maxServerArrayDatacenterPolicyDatacenterHrefPos {
		return nil, fmt.Errof("Invalid flags for serverArray.datacenterPolicy.datacenterHref field of datacenterPolicy array, %s were defined but max position is %s",
			len(r.ServerArrayDatacenterPolicyDatacenterHref), maxServerArrayDatacenterPolicyDatacenterHrefPos)
	}

	maxServerArrayDatacenterPolicyMaxPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.ServerArrayDatacenterPolicyMaxPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.datacenterPolicy.max field of datacenterPolicy array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.datacenterPolicy.max field of datacenterPolicy array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerArrayDatacenterPolicyMaxPos {
			maxServerArrayDatacenterPolicyMaxPos = pos
		}
	}
	if len(r.ServerArrayDatacenterPolicyMax) != maxServerArrayDatacenterPolicyMaxPos {
		return nil, fmt.Errof("Invalid flags for serverArray.datacenterPolicy.max field of datacenterPolicy array, %s were defined but max position is %s",
			len(r.ServerArrayDatacenterPolicyMax), maxServerArrayDatacenterPolicyMaxPos)
	}

	maxServerArrayDatacenterPolicyWeightPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.ServerArrayDatacenterPolicyWeightPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.datacenterPolicy.weight field of datacenterPolicy array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.datacenterPolicy.weight field of datacenterPolicy array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerArrayDatacenterPolicyWeightPos {
			maxServerArrayDatacenterPolicyWeightPos = pos
		}
	}
	if len(r.ServerArrayDatacenterPolicyWeight) != maxServerArrayDatacenterPolicyWeightPos {
		return nil, fmt.Errof("Invalid flags for serverArray.datacenterPolicy.weight field of datacenterPolicy array, %s were defined but max position is %s",
			len(r.ServerArrayDatacenterPolicyWeight), maxServerArrayDatacenterPolicyWeightPos)
	}
	if maxServerArrayDatacenterPolicyDatacenterHrefPos > -1 {
		serverArrayDatacenterPolicy := make([]*DatacenterPolicy2, maxServerArrayDatacenterPolicyDatacenterHrefPos+1)
		for i := 0; i < maxServerArrayDatacenterPolicyDatacenterHrefPos+1; i++ {
			serverArrayDatacenterPolicy[i] = *DatacenterPolicy2{
			//TBD
			}
		}
		serverArray.datacenterPolicy = serverArrayDatacenterPolicy
	}

	maxServerArrayElasticityParamsScheduleDayPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.ServerArrayElasticityParamsScheduleDayPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.elasticityParams.schedule.day field of schedule array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.elasticityParams.schedule.day field of schedule array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerArrayElasticityParamsScheduleDayPos {
			maxServerArrayElasticityParamsScheduleDayPos = pos
		}
	}
	if len(r.ServerArrayElasticityParamsScheduleDay) != maxServerArrayElasticityParamsScheduleDayPos {
		return nil, fmt.Errof("Invalid flags for serverArray.elasticityParams.schedule.day field of schedule array, %s were defined but max position is %s",
			len(r.ServerArrayElasticityParamsScheduleDay), maxServerArrayElasticityParamsScheduleDayPos)
	}

	maxServerArrayElasticityParamsScheduleMaxCountPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.ServerArrayElasticityParamsScheduleMaxCountPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.elasticityParams.schedule.maxCount field of schedule array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.elasticityParams.schedule.maxCount field of schedule array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerArrayElasticityParamsScheduleMaxCountPos {
			maxServerArrayElasticityParamsScheduleMaxCountPos = pos
		}
	}
	if len(r.ServerArrayElasticityParamsScheduleMaxCount) != maxServerArrayElasticityParamsScheduleMaxCountPos {
		return nil, fmt.Errof("Invalid flags for serverArray.elasticityParams.schedule.maxCount field of schedule array, %s were defined but max position is %s",
			len(r.ServerArrayElasticityParamsScheduleMaxCount), maxServerArrayElasticityParamsScheduleMaxCountPos)
	}

	maxServerArrayElasticityParamsScheduleMinCountPos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.ServerArrayElasticityParamsScheduleMinCountPos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.elasticityParams.schedule.minCount field of schedule array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.elasticityParams.schedule.minCount field of schedule array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerArrayElasticityParamsScheduleMinCountPos {
			maxServerArrayElasticityParamsScheduleMinCountPos = pos
		}
	}
	if len(r.ServerArrayElasticityParamsScheduleMinCount) != maxServerArrayElasticityParamsScheduleMinCountPos {
		return nil, fmt.Errof("Invalid flags for serverArray.elasticityParams.schedule.minCount field of schedule array, %s were defined but max position is %s",
			len(r.ServerArrayElasticityParamsScheduleMinCount), maxServerArrayElasticityParamsScheduleMinCountPos)
	}

	maxServerArrayElasticityParamsScheduleTimePos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.ServerArrayElasticityParamsScheduleTimePos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for serverArray.elasticityParams.schedule.time field of schedule array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of serverArray.elasticityParams.schedule.time field of schedule array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > maxServerArrayElasticityParamsScheduleTimePos {
			maxServerArrayElasticityParamsScheduleTimePos = pos
		}
	}
	if len(r.ServerArrayElasticityParamsScheduleTime) != maxServerArrayElasticityParamsScheduleTimePos {
		return nil, fmt.Errof("Invalid flags for serverArray.elasticityParams.schedule.time field of schedule array, %s were defined but max position is %s",
			len(r.ServerArrayElasticityParamsScheduleTime), maxServerArrayElasticityParamsScheduleTimePos)
	}
	if maxServerArrayElasticityParamsScheduleDayPos > -1 {
		serverArrayElasticityParamsSchedule := make([]*Schedule2, maxServerArrayElasticityParamsScheduleDayPos+1)
		for i := 0; i < maxServerArrayElasticityParamsScheduleDayPos+1; i++ {
			serverArrayElasticityParamsSchedule[i] = *Schedule2{
			//TBD
			}
		}
		serverArray.elasticityParams.schedule = serverArrayElasticityParamsSchedule
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
	CreateServerArrayRunner.FlagPattern(`serverArray\.datacenterPolicy\.(\d+)\.datacenterHref`, `The href of the Datacenter / Zone.`).Required().Capture(&CreateServerArrayRunner.serverArrayDatacenterPolicyDatacenterHrefPos).StringsVar(&CreateServerArrayRunner.serverArrayDatacenterPolicyDatacenterHref)
	CreateServerArrayRunner.FlagPattern(`serverArray\.datacenterPolicy\.(\d+)\.max`, `Max instances (0 for unlimited).`).Required().Capture(&CreateServerArrayRunner.serverArrayDatacenterPolicyMaxPos).StringsVar(&CreateServerArrayRunner.serverArrayDatacenterPolicyMax)
	CreateServerArrayRunner.FlagPattern(`serverArray\.datacenterPolicy\.(\d+)\.weight`, `Instance allocation (should total 100%).`).Required().Capture(&CreateServerArrayRunner.serverArrayDatacenterPolicyWeightPos).StringsVar(&CreateServerArrayRunner.serverArrayDatacenterPolicyWeight)
	CreateServerArrayRunner.Flag(`serverArray.deploymentHref`, `The href of the deployment for the Server Array.`).StringVar(CreateServerArrayRunner.serverArrayDeploymentHref)
	CreateServerArrayRunner.Flag(`serverArray.description`, `The description for the Server Array.`).StringVar(CreateServerArrayRunner.serverArrayDescription)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.alertSpecificParams.decisionThreshold`, `The percentage of servers that must agree in order to trigger an alert before an action is taken.`).StringVar(CreateServerArrayRunner.serverArrayElasticityParamsAlertSpecificParamsDecisionThreshold)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.alertSpecificParams.votersTagPredicate`, `The Voters Tag that RightScale will use in order to determine when to scale up/down.`).StringVar(CreateServerArrayRunner.serverArrayElasticityParamsAlertSpecificParamsVotersTagPredicate)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.bounds.maxCount`, `The maximum number of servers that can be operational at the same time in the server array.`).StringVar(CreateServerArrayRunner.serverArrayElasticityParamsBoundsMaxCount)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.bounds.minCount`, `The minimum number of servers that must be operational at all times in the server array.`).StringVar(CreateServerArrayRunner.serverArrayElasticityParamsBoundsMinCount)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.pacing.resizeCalmTime`, `The time (in minutes) on how long you want to wait before you repeat another action.`).StringVar(CreateServerArrayRunner.serverArrayElasticityParamsPacingResizeCalmTime)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.pacing.resizeDownBy`, `The number of servers to scale down by.`).StringVar(CreateServerArrayRunner.serverArrayElasticityParamsPacingResizeDownBy)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.pacing.resizeUpBy`, `The number of servers to scale up by.`).StringVar(CreateServerArrayRunner.serverArrayElasticityParamsPacingResizeUpBy)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.collectAuditEntries`, `The audit SQS queue that will store audit entries.`).StringVar(CreateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsCollectAuditEntries)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.itemAge.algorithm`, `The algorithm that defines how an item's age will be determined, either by the average age or max (oldest) age.`).StringVar(CreateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.itemAge.maxAge`, `The threshold (in seconds) before a resize action occurs on the server array.`).StringVar(CreateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.itemAge.regexp`, `The regexp that helps the system determine an item's "age" in the queue. Example: created_at: (\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d UTC)`).StringVar(CreateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsItemAgeRegexp)
	CreateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.queueSize.itemsPerInstance`, `Defines the ratio of worker instances per items in the queue. Example: If there are 50 items in the queue and "Items per instance" is set to 10, the server array will resize to 5 worker instances (50/10).  Default = 10`).StringVar(CreateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance)
	CreateServerArrayRunner.FlagPattern(`serverArray\.elasticityParams\.schedule\.(\d+)\.day`, `Specifies the day when an alert-based array resizes.`).Required().Capture(&CreateServerArrayRunner.serverArrayElasticityParamsScheduleDayPos).StringsVar(&CreateServerArrayRunner.serverArrayElasticityParamsScheduleDay)
	CreateServerArrayRunner.FlagPattern(`serverArray\.elasticityParams\.schedule\.(\d+)\.maxCount`, `The maximum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`).Required().Capture(&CreateServerArrayRunner.serverArrayElasticityParamsScheduleMaxCountPos).StringsVar(&CreateServerArrayRunner.serverArrayElasticityParamsScheduleMaxCount)
	CreateServerArrayRunner.FlagPattern(`serverArray\.elasticityParams\.schedule\.(\d+)\.minCount`, `The minimum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`).Required().Capture(&CreateServerArrayRunner.serverArrayElasticityParamsScheduleMinCountPos).StringsVar(&CreateServerArrayRunner.serverArrayElasticityParamsScheduleMinCount)
	CreateServerArrayRunner.FlagPattern(`serverArray\.elasticityParams\.schedule\.(\d+)\.time`, `Specifies the time when an alert-based array resizes.`).Required().Capture(&CreateServerArrayRunner.serverArrayElasticityParamsScheduleTimePos).StringsVar(&CreateServerArrayRunner.serverArrayElasticityParamsScheduleTime)
	CreateServerArrayRunner.Flag(`serverArray.instance.associatePublicIpAddress`, `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`).StringVar(CreateServerArrayRunner.serverArrayInstanceAssociatePublicIpAddress)
	CreateServerArrayRunner.Flag(`serverArray.instance.cloudHref`, `The href of the Cloud that the array will be associated with.`).Required().StringVar(&CreateServerArrayRunner.serverArrayInstanceCloudHref)
	CreateServerArrayRunner.Flag(`serverArray.instance.cloudSpecificAttributes.automaticInstanceStoreMapping`, `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`).StringVar(CreateServerArrayRunner.serverArrayInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping)
	CreateServerArrayRunner.Flag(`serverArray.instance.cloudSpecificAttributes.iamInstanceProfile`, `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`).StringVar(CreateServerArrayRunner.serverArrayInstanceCloudSpecificAttributesIamInstanceProfile)
	CreateServerArrayRunner.Flag(`serverArray.instance.cloudSpecificAttributes.rootVolumePerformance`, `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`).StringVar(CreateServerArrayRunner.serverArrayInstanceCloudSpecificAttributesRootVolumePerformance)
	CreateServerArrayRunner.Flag(`serverArray.instance.cloudSpecificAttributes.rootVolumeSize`, `The size for root disk. Not supported in all Clouds.`).StringVar(CreateServerArrayRunner.serverArrayInstanceCloudSpecificAttributesRootVolumeSize)
	CreateServerArrayRunner.Flag(`serverArray.instance.cloudSpecificAttributes.rootVolumeTypeUid`, `The type of root volume for instance. Only available on clouds supporting root volume type.`).StringVar(CreateServerArrayRunner.serverArrayInstanceCloudSpecificAttributesRootVolumeTypeUid)
	CreateServerArrayRunner.Flag(`serverArray.instance.datacenterHref`, `The href of the Datacenter / Zone. For multiple Datacenters, use 'datacenter_policy' instead.`).StringVar(CreateServerArrayRunner.serverArrayInstanceDatacenterHref)
	CreateServerArrayRunner.Flag(`serverArray.instance.imageHref`, `The href of the Image to be used.`).StringVar(CreateServerArrayRunner.serverArrayInstanceImageHref)
	CreateServerArrayRunner.FlagPattern(`serverArray\.instance\.inputs\.([a-z0-9_]+)`, ``).Capture(&CreateServerArrayRunner.serverArrayInstanceInputsNames).StringVar(CreateServerArrayRunner.serverArrayInstanceInputsValues)
	CreateServerArrayRunner.Flag(`serverArray.instance.instanceTypeHref`, `The href of the Instance Type.`).StringVar(CreateServerArrayRunner.serverArrayInstanceInstanceTypeHref)
	CreateServerArrayRunner.Flag(`serverArray.instance.ipForwardingEnabled`, `Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.`).StringVar(CreateServerArrayRunner.serverArrayInstanceIpForwardingEnabled)
	CreateServerArrayRunner.Flag(`serverArray.instance.kernelImageHref`, `The href of the Kernel Image.`).StringVar(CreateServerArrayRunner.serverArrayInstanceKernelImageHref)
	CreateServerArrayRunner.Flag(`serverArray.instance.multiCloudImageHref`, `The href of the MultiCloudImage to be used.`).StringVar(CreateServerArrayRunner.serverArrayInstanceMultiCloudImageHref)
	CreateServerArrayRunner.Flag(`serverArray.instance.placementGroupHref`, `The href of the Placement Group.`).StringVar(CreateServerArrayRunner.serverArrayInstancePlacementGroupHref)
	CreateServerArrayRunner.Flag(`serverArray.instance.ramdiskImageHref`, `The href of the Ramdisk Image.`).StringVar(CreateServerArrayRunner.serverArrayInstanceRamdiskImageHref)
	CreateServerArrayRunner.FlagPattern(`serverArray\.instance\.securityGroupHrefs\.(\d+)`, `The hrefs of the Security Groups.`).Capture(&CreateServerArrayRunner.serverArrayInstanceSecurityGroupHrefsPos).StringsVar(CreateServerArrayRunner.serverArrayInstanceSecurityGroupHrefs)
	CreateServerArrayRunner.Flag(`serverArray.instance.serverTemplateHref`, `The ServerTemplate that will be used to create the worker instances in the server array.`).Required().StringVar(&CreateServerArrayRunner.serverArrayInstanceServerTemplateHref)
	CreateServerArrayRunner.Flag(`serverArray.instance.sshKeyHref`, `The href of the SSH Key to be used.`).StringVar(CreateServerArrayRunner.serverArrayInstanceSshKeyHref)
	CreateServerArrayRunner.FlagPattern(`serverArray\.instance\.subnetHrefs\.(\d+)`, `The hrefs of the updated Subnets.`).Capture(&CreateServerArrayRunner.serverArrayInstanceSubnetHrefsPos).StringsVar(CreateServerArrayRunner.serverArrayInstanceSubnetHrefs)
	CreateServerArrayRunner.Flag(`serverArray.instance.userData`, `User data that RightScale automatically passes to your instance at boot time.`).StringVar(CreateServerArrayRunner.serverArrayInstanceUserData)
	CreateServerArrayRunner.Flag(`serverArray.name`, `The name for the Server Array.`).Required().StringVar(&CreateServerArrayRunner.serverArrayName)
	CreateServerArrayRunner.Flag(`serverArray.optimized`, `A flag indicating whether Instances of this ServerArray should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`).StringVar(CreateServerArrayRunner.serverArrayOptimized)
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
	IndexServerArraysRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexServerArraysRunner.filterPos).StringsVar(IndexServerArraysRunner.filter)
	IndexServerArraysRunner.Flag(`view`, ``).StringVar(IndexServerArraysRunner.view)
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
	ShowServerArrayRunner.Flag(`view`, ``).StringVar(ShowServerArrayRunner.view)
	registry[ShowServerArrayCmd.FullCommand()] = ShowServerArrayRunner

	UpdateServerArrayRunner := new(UpdateServerArrayServerArrayRunner)
	UpdateServerArrayCmd := resCmd.Command("UpdateServerArray", `Updates attributes of a single server array.`)
	UpdateServerArrayRunner.Flag(`id`, ``).Required().StringVar(&UpdateServerArrayRunner.id)
	UpdateServerArrayRunner.Flag(`serverArray.arrayType`, `The updated array type for the Server Array.`).StringVar(UpdateServerArrayRunner.serverArrayArrayType)
	UpdateServerArrayRunner.FlagPattern(`serverArray\.datacenterPolicy\.(\d+)\.datacenterHref`, `The href of the Datacenter / Zone.`).Required().Capture(&UpdateServerArrayRunner.serverArrayDatacenterPolicyDatacenterHrefPos).StringsVar(&UpdateServerArrayRunner.serverArrayDatacenterPolicyDatacenterHref)
	UpdateServerArrayRunner.FlagPattern(`serverArray\.datacenterPolicy\.(\d+)\.max`, `Max instances (0 for unlimited).`).Required().Capture(&UpdateServerArrayRunner.serverArrayDatacenterPolicyMaxPos).StringsVar(&UpdateServerArrayRunner.serverArrayDatacenterPolicyMax)
	UpdateServerArrayRunner.FlagPattern(`serverArray\.datacenterPolicy\.(\d+)\.weight`, `Instance allocation (should total 100%).`).Required().Capture(&UpdateServerArrayRunner.serverArrayDatacenterPolicyWeightPos).StringsVar(&UpdateServerArrayRunner.serverArrayDatacenterPolicyWeight)
	UpdateServerArrayRunner.Flag(`serverArray.deploymentHref`, `The updated href of the deployment for the Server Array.`).StringVar(UpdateServerArrayRunner.serverArrayDeploymentHref)
	UpdateServerArrayRunner.Flag(`serverArray.description`, `The updated description for the Server Array.`).StringVar(UpdateServerArrayRunner.serverArrayDescription)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.alertSpecificParams.decisionThreshold`, `The updated percentage of servers that must agree in order to trigger an alert before an action is taken.`).StringVar(UpdateServerArrayRunner.serverArrayElasticityParamsAlertSpecificParamsDecisionThreshold)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.alertSpecificParams.votersTagPredicate`, `The updated Voters Tag that RightScale will use in order to determine when to scale up/down.`).StringVar(UpdateServerArrayRunner.serverArrayElasticityParamsAlertSpecificParamsVotersTagPredicate)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.bounds.maxCount`, `The updated maximum number of servers that can be operational at the same time in the server array.`).StringVar(UpdateServerArrayRunner.serverArrayElasticityParamsBoundsMaxCount)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.bounds.minCount`, `The updated minimum number of servers that must be operational at all times in the server array.`).StringVar(UpdateServerArrayRunner.serverArrayElasticityParamsBoundsMinCount)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.pacing.resizeCalmTime`, `The updated time (in minutes) on how long you want to wait before you repeat another action.`).StringVar(UpdateServerArrayRunner.serverArrayElasticityParamsPacingResizeCalmTime)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.pacing.resizeDownBy`, `The updated number of servers to scale down by.`).StringVar(UpdateServerArrayRunner.serverArrayElasticityParamsPacingResizeDownBy)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.pacing.resizeUpBy`, `The updated number of servers to scale up by.`).StringVar(UpdateServerArrayRunner.serverArrayElasticityParamsPacingResizeUpBy)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.collectAuditEntries`, `The updated audit SQS queue that will store audit entries.`).StringVar(UpdateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsCollectAuditEntries)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.itemAge.algorithm`, `The updated algorithm that defines how an item's age will be determined, either by the average age or max (oldest) age.`).StringVar(UpdateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.itemAge.maxAge`, `The updated threshold (in seconds) before a resize action occurs on the server array.`).StringVar(UpdateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.itemAge.regexp`, `The updated regexp that helps the system determine an item's "age" in the queue. Example: created_at: (\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d UTC)`).StringVar(UpdateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsItemAgeRegexp)
	UpdateServerArrayRunner.Flag(`serverArray.elasticityParams.queueSpecificParams.queueSize.itemsPerInstance`, `Defines the ratio of worker instances per items in the queue. Example: If there are 50 items in the queue and "Items per instance" is set to 10, the server array will resize to 5 worker instances (50/10).  Default = 10`).StringVar(UpdateServerArrayRunner.serverArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance)
	UpdateServerArrayRunner.FlagPattern(`serverArray\.elasticityParams\.schedule\.(\d+)\.day`, `The updated day when an alert-based array resizes.`).Required().Capture(&UpdateServerArrayRunner.serverArrayElasticityParamsScheduleDayPos).StringsVar(&UpdateServerArrayRunner.serverArrayElasticityParamsScheduleDay)
	UpdateServerArrayRunner.FlagPattern(`serverArray\.elasticityParams\.schedule\.(\d+)\.maxCount`, `The updated maximum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`).Required().Capture(&UpdateServerArrayRunner.serverArrayElasticityParamsScheduleMaxCountPos).StringsVar(&UpdateServerArrayRunner.serverArrayElasticityParamsScheduleMaxCount)
	UpdateServerArrayRunner.FlagPattern(`serverArray\.elasticityParams\.schedule\.(\d+)\.minCount`, `The updated minimum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`).Required().Capture(&UpdateServerArrayRunner.serverArrayElasticityParamsScheduleMinCountPos).StringsVar(&UpdateServerArrayRunner.serverArrayElasticityParamsScheduleMinCount)
	UpdateServerArrayRunner.FlagPattern(`serverArray\.elasticityParams\.schedule\.(\d+)\.time`, `The updated time when an alert-based array resizes.`).Required().Capture(&UpdateServerArrayRunner.serverArrayElasticityParamsScheduleTimePos).StringsVar(&UpdateServerArrayRunner.serverArrayElasticityParamsScheduleTime)
	UpdateServerArrayRunner.Flag(`serverArray.name`, `The updated name for the Server Array.`).StringVar(UpdateServerArrayRunner.serverArrayName)
	UpdateServerArrayRunner.Flag(`serverArray.optimized`, `A flag indicating whether Instances of this ServerArray should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`).StringVar(UpdateServerArrayRunner.serverArrayOptimized)
	UpdateServerArrayRunner.Flag(`serverArray.state`, `The updated status of the server array. If active, the server array is enabled for scaling actions.`).StringVar(UpdateServerArrayRunner.serverArrayState)
	registry[UpdateServerArrayCmd.FullCommand()] = UpdateServerArrayRunner
}

/****** ServerTemplate ******/

type CloneServerTemplateRunner struct {
	id                        string
	serverTemplateDescription *string
	serverTemplateName        string
}

func (r *CloneServerTemplateRunner) Run(c *Client) (interface{}, error) {

	/** Handle serverTemplate parameter **/
	var serverTemplate ServerTemplateParam
	// Load JSON if provided
	if len(r.serverTemplateJson) > 0 {
		if err := Json.Unmarshal(r.serverTemplateJson, &serverTemplate); err != nil {
			return nil, fmt.Errorf("Could not load serverTemplate JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.ServerTemplateDescription != nil {
		serverTemplate.description = r.ServerTemplateDescription
	}

	if len(r.ServerTemplateName) > 0 {
		serverTemplate.name = r.ServerTemplateName
	}

	return c.CloneServerTemplate(r.id, &serverTemplate)
}

type CommitServerTemplateRunner struct {
	commitHeadDependencies string
	commitMessage          string
	freezeRepositories     string
	id                     string
}

func (r *CommitServerTemplateRunner) Run(c *Client) (interface{}, error) {
	return c.CommitServerTemplate(r.commitHeadDependencies, r.commitMessage, r.freezeRepositories, r.id)
}

type CreateServerTemplateRunner struct {
	serverTemplateDescription *string
	serverTemplateName        string
}

func (r *CreateServerTemplateRunner) Run(c *Client) (interface{}, error) {

	/** Handle serverTemplate parameter **/
	var serverTemplate ServerTemplateParam2
	// Load JSON if provided
	if len(r.serverTemplateJson) > 0 {
		if err := Json.Unmarshal(r.serverTemplateJson, &serverTemplate); err != nil {
			return nil, fmt.Errorf("Could not load serverTemplate JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.ServerTemplateDescription != nil {
		serverTemplate.description = r.ServerTemplateDescription
	}

	if len(r.ServerTemplateName) > 0 {
		serverTemplate.name = r.ServerTemplateName
	}

	return c.CreateServerTemplate(&serverTemplate)
}

type DestroyServerTemplateRunner struct {
	id string
}

func (r *DestroyServerTemplateRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyServerTemplate(r.id)
}

type DetectChangesInHeadServerTemplateRunner struct {
	id string
}

func (r *DetectChangesInHeadServerTemplateRunner) Run(c *Client) (interface{}, error) {
	return c.DetectChangesInHeadServerTemplate(r.id)
}

type IndexServerTemplatesRunner struct {
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexServerTemplatesRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexServerTemplates(options)
}

type PublishServerTemplateRunner struct {
	accountGroupHrefs    []string
	accountGroupHrefsPos []string
	allowComments        *string
	categories           []string
	categoriesPos        []string
	descriptionsLong     string
	descriptionsNotes    string
	descriptionsShort    string
	emailComments        *string
	id                   string
}

func (r *PublishServerTemplateRunner) Run(c *Client) (interface{}, error) {

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
	if len(r.DescriptionsLong) > 0 {
		descriptions.long = r.DescriptionsLong
	}

	if len(r.DescriptionsNotes) > 0 {
		descriptions.notes = r.DescriptionsNotes
	}

	if len(r.DescriptionsShort) > 0 {
		descriptions.short = r.DescriptionsShort
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.allowComments != nil {
		options["allow_comments"] = *r.allowComments
	}
	options["categories"] = categories
	if r.emailComments != nil {
		options["email_comments"] = *r.emailComments
	}

	return c.PublishServerTemplate(accountGroupHrefs, &descriptions, r.id, options)
}

type ResolveServerTemplateRunner struct {
	id string
}

func (r *ResolveServerTemplateRunner) Run(c *Client) (interface{}, error) {
	return c.ResolveServerTemplate(r.id)
}

type ShowServerTemplateRunner struct {
	id   string
	view *string
}

func (r *ShowServerTemplateRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowServerTemplate(r.id, options)
}

type SwapRepositoryServerTemplateRunner struct {
	id                   string
	sourceRepositoryHref string
	targetRepositoryHref string
}

func (r *SwapRepositoryServerTemplateRunner) Run(c *Client) (interface{}, error) {
	return c.SwapRepositoryServerTemplate(r.id, r.sourceRepositoryHref, r.targetRepositoryHref)
}

type UpdateServerTemplateRunner struct {
	id                        string
	serverTemplateDescription *string
	serverTemplateName        *string
}

func (r *UpdateServerTemplateRunner) Run(c *Client) (interface{}, error) {

	/** Handle serverTemplate parameter **/
	var serverTemplate ServerTemplateParam
	// Load JSON if provided
	if len(r.serverTemplateJson) > 0 {
		if err := Json.Unmarshal(r.serverTemplateJson, &serverTemplate); err != nil {
			return nil, fmt.Errorf("Could not load serverTemplate JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.ServerTemplateDescription != nil {
		serverTemplate.description = r.ServerTemplateDescription
	}

	if r.ServerTemplateName != nil {
		serverTemplate.name = r.ServerTemplateName
	}

	return c.UpdateServerTemplate(r.id, &serverTemplate)
}

// Register all ServerTemplate actions
func registerServerTemplateCmds(app *kingpin.Application) {
	resCmd := app.Cmd("ServerTemplate", `ServerTemplates allow you to pre-configure servers by starting from a base image and adding scripts that run during the boot, operational, and shutdown phases...`)

	CloneServerTemplateRunner := new(CloneServerTemplateServerTemplateRunner)
	CloneServerTemplateCmd := resCmd.Command("CloneServerTemplate", `Clones a given ServerTemplate.`)
	CloneServerTemplateRunner.Flag(`id`, ``).Required().StringVar(&CloneServerTemplateRunner.id)
	CloneServerTemplateRunner.Flag(`serverTemplate.description`, `The description for the cloned ServerTemplate.`).StringVar(CloneServerTemplateRunner.serverTemplateDescription)
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
	CreateServerTemplateRunner.Flag(`serverTemplate.description`, `The description of the ServerTemplate to be created.`).StringVar(CreateServerTemplateRunner.serverTemplateDescription)
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
	IndexServerTemplatesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexServerTemplatesRunner.filterPos).StringsVar(IndexServerTemplatesRunner.filter)
	IndexServerTemplatesRunner.Flag(`view`, ``).StringVar(IndexServerTemplatesRunner.view)
	registry[IndexServerTemplatesCmd.FullCommand()] = IndexServerTemplatesRunner

	PublishServerTemplateRunner := new(PublishServerTemplateServerTemplateRunner)
	PublishServerTemplateCmd := resCmd.Command("PublishServerTemplate", `Publishes a given ServerTemplate and its subordinates.`)
	PublishServerTemplateRunner.FlagPattern(`accountGroupHrefs\.(\d+)`, `List of hrefs of account groups to publish to.`).Required().Capture(&PublishServerTemplateRunner.accountGroupHrefsPos).StringsVar(&PublishServerTemplateRunner.accountGroupHrefs)
	PublishServerTemplateRunner.Flag(`allowComments`, `Allow users to leave comments on this ServerTemplate.`).StringVar(PublishServerTemplateRunner.allowComments)
	PublishServerTemplateRunner.FlagPattern(`categories\.(\d+)`, `List of Categories.`).Capture(&PublishServerTemplateRunner.categoriesPos).StringsVar(PublishServerTemplateRunner.categories)
	PublishServerTemplateRunner.Flag(`descriptions.long`, `Long Description.`).Required().StringVar(&PublishServerTemplateRunner.descriptionsLong)
	PublishServerTemplateRunner.Flag(`descriptions.notes`, `New Revision Notes.`).Required().StringVar(&PublishServerTemplateRunner.descriptionsNotes)
	PublishServerTemplateRunner.Flag(`descriptions.short`, `Short Description.`).Required().StringVar(&PublishServerTemplateRunner.descriptionsShort)
	PublishServerTemplateRunner.Flag(`emailComments`, `Email me when a user comments on this ServerTemplate.`).StringVar(PublishServerTemplateRunner.emailComments)
	PublishServerTemplateRunner.Flag(`id`, ``).Required().StringVar(&PublishServerTemplateRunner.id)
	registry[PublishServerTemplateCmd.FullCommand()] = PublishServerTemplateRunner

	ResolveServerTemplateRunner := new(ResolveServerTemplateServerTemplateRunner)
	ResolveServerTemplateCmd := resCmd.Command("ResolveServerTemplate", `Enumerates all attached cookbooks, missing dependencies and bound executables.`)
	ResolveServerTemplateRunner.Flag(`id`, ``).Required().StringVar(&ResolveServerTemplateRunner.id)
	registry[ResolveServerTemplateCmd.FullCommand()] = ResolveServerTemplateRunner

	ShowServerTemplateRunner := new(ShowServerTemplateServerTemplateRunner)
	ShowServerTemplateCmd := resCmd.Command("ShowServerTemplate", `Show information about a single ServerTemplate. HEAD revisions have a revision of 0.`)
	ShowServerTemplateRunner.Flag(`id`, ``).Required().StringVar(&ShowServerTemplateRunner.id)
	ShowServerTemplateRunner.Flag(`view`, ``).StringVar(ShowServerTemplateRunner.view)
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
	UpdateServerTemplateRunner.Flag(`serverTemplate.description`, `The updated description for the ServerTemplate.`).StringVar(UpdateServerTemplateRunner.serverTemplateDescription)
	UpdateServerTemplateRunner.Flag(`serverTemplate.name`, `The updated name for the ServerTemplate.`).StringVar(UpdateServerTemplateRunner.serverTemplateName)
	registry[UpdateServerTemplateCmd.FullCommand()] = UpdateServerTemplateRunner
}

/****** ServerTemplateMultiCloudImage ******/

type CreateServerTemplateMultiCloudImageRunner struct {
	serverTemplateMultiCloudImageMultiCloudImageHref string
	serverTemplateMultiCloudImageServerTemplateHref  string
}

func (r *CreateServerTemplateMultiCloudImageRunner) Run(c *Client) (interface{}, error) {

	/** Handle serverTemplateMultiCloudImage parameter **/
	var serverTemplateMultiCloudImage ServerTemplateMultiCloudImageParam
	// Load JSON if provided
	if len(r.serverTemplateMultiCloudImageJson) > 0 {
		if err := Json.Unmarshal(r.serverTemplateMultiCloudImageJson, &serverTemplateMultiCloudImage); err != nil {
			return nil, fmt.Errorf("Could not load serverTemplateMultiCloudImage JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.ServerTemplateMultiCloudImageMultiCloudImageHref) > 0 {
		serverTemplateMultiCloudImage.multiCloudImageHref = r.ServerTemplateMultiCloudImageMultiCloudImageHref
	}

	if len(r.ServerTemplateMultiCloudImageServerTemplateHref) > 0 {
		serverTemplateMultiCloudImage.serverTemplateHref = r.ServerTemplateMultiCloudImageServerTemplateHref
	}

	return c.CreateServerTemplateMultiCloudImage(&serverTemplateMultiCloudImage)
}

type DestroyServerTemplateMultiCloudImageRunner struct {
	id string
}

func (r *DestroyServerTemplateMultiCloudImageRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyServerTemplateMultiCloudImage(r.id)
}

type IndexServerTemplateMultiCloudImagesRunner struct {
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexServerTemplateMultiCloudImagesRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexServerTemplateMultiCloudImages(options)
}

type MakeDefaultServerTemplateMultiCloudImageRunner struct {
	id string
}

func (r *MakeDefaultServerTemplateMultiCloudImageRunner) Run(c *Client) (interface{}, error) {
	return c.MakeDefaultServerTemplateMultiCloudImage(r.id)
}

type ShowServerTemplateMultiCloudImageRunner struct {
	id   string
	view *string
}

func (r *ShowServerTemplateMultiCloudImageRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowServerTemplateMultiCloudImage(r.id, options)
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
	IndexServerTemplateMultiCloudImagesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexServerTemplateMultiCloudImagesRunner.filterPos).StringsVar(IndexServerTemplateMultiCloudImagesRunner.filter)
	IndexServerTemplateMultiCloudImagesRunner.Flag(`view`, ``).StringVar(IndexServerTemplateMultiCloudImagesRunner.view)
	registry[IndexServerTemplateMultiCloudImagesCmd.FullCommand()] = IndexServerTemplateMultiCloudImagesRunner

	MakeDefaultServerTemplateMultiCloudImageRunner := new(MakeDefaultServerTemplateMultiCloudImageServerTemplateMultiCloudImageRunner)
	MakeDefaultServerTemplateMultiCloudImageCmd := resCmd.Command("MakeDefaultServerTemplateMultiCloudImage", `Makes a given ServerTemplateMultiCloudImage the default for the ServerTemplate.`)
	MakeDefaultServerTemplateMultiCloudImageRunner.Flag(`id`, ``).Required().StringVar(&MakeDefaultServerTemplateMultiCloudImageRunner.id)
	registry[MakeDefaultServerTemplateMultiCloudImageCmd.FullCommand()] = MakeDefaultServerTemplateMultiCloudImageRunner

	ShowServerTemplateMultiCloudImageRunner := new(ShowServerTemplateMultiCloudImageServerTemplateMultiCloudImageRunner)
	ShowServerTemplateMultiCloudImageCmd := resCmd.Command("ShowServerTemplateMultiCloudImage", `Show information about a single ServerTemplateMultiCloudImage which represents an association between a ServerTemplate and a MultiCloudImage.`)
	ShowServerTemplateMultiCloudImageRunner.Flag(`id`, ``).Required().StringVar(&ShowServerTemplateMultiCloudImageRunner.id)
	ShowServerTemplateMultiCloudImageRunner.Flag(`view`, ``).StringVar(ShowServerTemplateMultiCloudImageRunner.view)
	registry[ShowServerTemplateMultiCloudImageCmd.FullCommand()] = ShowServerTemplateMultiCloudImageRunner
}

/****** Session ******/

type AccountsSessionRunner struct {
	email    *string
	password *string
	view     *string
}

func (r *AccountsSessionRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.email != nil {
		options["email"] = *r.email
	}
	if r.password != nil {
		options["password"] = *r.password
	}
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.AccountsSession(options)
}

type CreateSessionRunner struct {
	accountHref string
	email       string
	password    string
}

func (r *CreateSessionRunner) Run(c *Client) (interface{}, error) {
	return c.CreateSession(r.accountHref, r.email, r.password)
}

type CreateInstanceSessionSessionRunner struct {
	accountHref   string
	instanceToken string
}

func (r *CreateInstanceSessionSessionRunner) Run(c *Client) (interface{}, error) {
	return c.CreateInstanceSessionSession(r.accountHref, r.instanceToken)
}

type IndexSessionsRunner struct {
}

func (r *IndexSessionsRunner) Run(c *Client) (interface{}, error) {
	return c.IndexSessions()
}

type IndexInstanceSessionSessionRunner struct {
}

func (r *IndexInstanceSessionSessionRunner) Run(c *Client) (interface{}, error) {
	return c.IndexInstanceSessionSession()
}

// Register all Session actions
func registerSessionCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Session", `The sessions resource is in charge of creating API sessions that are bound to a given account`)

	AccountsSessionRunner := new(AccountsSessionSessionRunner)
	AccountsSessionCmd := resCmd.Command("AccountsSession", `List all the accounts that a user has access to.`)
	AccountsSessionRunner.Flag(`email`, `The email to login with if not using existing session.`).StringVar(AccountsSessionRunner.email)
	AccountsSessionRunner.Flag(`password`, `The corresponding password.`).StringVar(AccountsSessionRunner.password)
	AccountsSessionRunner.Flag(`view`, `Extended view shows account permissions and products`).StringVar(AccountsSessionRunner.view)
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

type CreateSshKeyRunner struct {
	cloudId    string
	sshKeyName string
}

func (r *CreateSshKeyRunner) Run(c *Client) (interface{}, error) {

	/** Handle sshKey parameter **/
	var sshKey SshKeyParam
	// Load JSON if provided
	if len(r.sshKeyJson) > 0 {
		if err := Json.Unmarshal(r.sshKeyJson, &sshKey); err != nil {
			return nil, fmt.Errorf("Could not load sshKey JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.SshKeyName) > 0 {
		sshKey.name = r.SshKeyName
	}

	return c.CreateSshKey(r.cloudId, &sshKey)
}

type DestroySshKeyRunner struct {
	cloudId string
	id      string
}

func (r *DestroySshKeyRunner) Run(c *Client) (interface{}, error) {
	return c.DestroySshKey(r.cloudId, r.id)
}

type IndexSshKeysRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexSshKeysRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexSshKeys(r.cloudId, options)
}

type ShowSshKeyRunner struct {
	cloudId string
	id      string
	view    *string
}

func (r *ShowSshKeyRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowSshKey(r.cloudId, r.id, options)
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
	IndexSshKeysRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexSshKeysRunner.filterPos).StringsVar(IndexSshKeysRunner.filter)
	IndexSshKeysRunner.Flag(`view`, ``).StringVar(IndexSshKeysRunner.view)
	registry[IndexSshKeysCmd.FullCommand()] = IndexSshKeysRunner

	ShowSshKeyRunner := new(ShowSshKeySshKeyRunner)
	ShowSshKeyCmd := resCmd.Command("ShowSshKey", `Displays information about a single ssh key.`)
	ShowSshKeyRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowSshKeyRunner.cloudId)
	ShowSshKeyRunner.Flag(`id`, ``).Required().StringVar(&ShowSshKeyRunner.id)
	ShowSshKeyRunner.Flag(`view`, ``).StringVar(ShowSshKeyRunner.view)
	registry[ShowSshKeyCmd.FullCommand()] = ShowSshKeyRunner
}

/****** Subnet ******/

type CreateSubnetRunner struct {
	cloudId              string
	instanceId           string
	subnetCidrBlock      string
	subnetDatacenterHref *string
	subnetDescription    *string
	subnetName           *string
	subnetNetworkHref    string
}

func (r *CreateSubnetRunner) Run(c *Client) (interface{}, error) {

	/** Handle subnet parameter **/
	var subnet SubnetParam
	// Load JSON if provided
	if len(r.subnetJson) > 0 {
		if err := Json.Unmarshal(r.subnetJson, &subnet); err != nil {
			return nil, fmt.Errorf("Could not load subnet JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.SubnetCidrBlock) > 0 {
		subnet.cidrBlock = r.SubnetCidrBlock
	}

	if r.SubnetDatacenterHref != nil {
		subnet.datacenterHref = r.SubnetDatacenterHref
	}

	if r.SubnetDescription != nil {
		subnet.description = r.SubnetDescription
	}

	if r.SubnetName != nil {
		subnet.name = r.SubnetName
	}

	if len(r.SubnetNetworkHref) > 0 {
		subnet.networkHref = r.SubnetNetworkHref
	}

	return c.CreateSubnet(r.cloudId, r.instanceId, &subnet)
}

type DestroySubnetRunner struct {
	cloudId    string
	id         string
	instanceId string
}

func (r *DestroySubnetRunner) Run(c *Client) (interface{}, error) {
	return c.DestroySubnet(r.cloudId, r.id, r.instanceId)
}

type IndexSubnetsRunner struct {
	cloudId    string
	filter     []string
	filterPos  []string
	instanceId string
}

func (r *IndexSubnetsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter

	return c.IndexSubnets(r.cloudId, r.instanceId, options)
}

type ShowSubnetRunner struct {
	cloudId    string
	id         string
	instanceId string
}

func (r *ShowSubnetRunner) Run(c *Client) (interface{}, error) {
	return c.ShowSubnet(r.cloudId, r.id, r.instanceId)
}

type UpdateSubnetRunner struct {
	cloudId              string
	id                   string
	instanceId           string
	subnetDescription    *string
	subnetName           *string
	subnetRouteTableHref *string
}

func (r *UpdateSubnetRunner) Run(c *Client) (interface{}, error) {

	/** Handle subnet parameter **/
	var subnet SubnetParam2
	// Load JSON if provided
	if len(r.subnetJson) > 0 {
		if err := Json.Unmarshal(r.subnetJson, &subnet); err != nil {
			return nil, fmt.Errorf("Could not load subnet JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.SubnetDescription != nil {
		subnet.description = r.SubnetDescription
	}

	if r.SubnetName != nil {
		subnet.name = r.SubnetName
	}

	if r.SubnetRouteTableHref != nil {
		subnet.routeTableHref = r.SubnetRouteTableHref
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
	CreateSubnetRunner.Flag(`subnet.datacenterHref`, `The associated Datacenter.`).StringVar(CreateSubnetRunner.subnetDatacenterHref)
	CreateSubnetRunner.Flag(`subnet.description`, `The description for the Subnet.`).StringVar(CreateSubnetRunner.subnetDescription)
	CreateSubnetRunner.Flag(`subnet.name`, `The name for the Subnet.`).StringVar(CreateSubnetRunner.subnetName)
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
	IndexSubnetsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexSubnetsRunner.filterPos).StringsVar(IndexSubnetsRunner.filter)
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
	UpdateSubnetRunner.Flag(`subnet.description`, `The updated description for the Subnet.`).StringVar(UpdateSubnetRunner.subnetDescription)
	UpdateSubnetRunner.Flag(`subnet.name`, `The updated name for the Subnet.`).StringVar(UpdateSubnetRunner.subnetName)
	UpdateSubnetRunner.Flag(`subnet.routeTableHref`, `The RouteTable to associate/dissociate with this Subnet. Pass this parameter with an empty string to reset this Subnet to use the default RouteTable.`).StringVar(UpdateSubnetRunner.subnetRouteTableHref)
	registry[UpdateSubnetCmd.FullCommand()] = UpdateSubnetRunner
}

/****** Tag ******/

type ByResourceTagRunner struct {
	resourceHrefs    []string
	resourceHrefsPos []string
}

func (r *ByResourceTagRunner) Run(c *Client) (interface{}, error) {

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

type ByTagTagRunner struct {
	includeTagsWithPrefix *string
	matchAll              *string
	resourceType          string
	tags                  []string
	tagsPos               []string
	withDeleted           *string
}

func (r *ByTagTagRunner) Run(c *Client) (interface{}, error) {

	/** Handle tags parameter **/
	var tags []string

	for i, v := range r.tags {
		pos, err := strconv.Atoi(r.tagsPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for tags array", r.tagsPos[i])
		}
		tags[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.includeTagsWithPrefix != nil {
		options["include_tags_with_prefix"] = *r.includeTagsWithPrefix
	}
	if r.matchAll != nil {
		options["match_all"] = *r.matchAll
	}
	if r.withDeleted != nil {
		options["with_deleted"] = *r.withDeleted
	}

	return c.ByTagTag(r.resourceType, tags, options)
}

type MultiAddTagsRunner struct {
	resourceHrefs    []string
	resourceHrefsPos []string
	tags             []string
	tagsPos          []string
}

func (r *MultiAddTagsRunner) Run(c *Client) (interface{}, error) {

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

type MultiDeleteTagsRunner struct {
	resourceHrefs    []string
	resourceHrefsPos []string
	tags             []string
	tagsPos          []string
}

func (r *MultiDeleteTagsRunner) Run(c *Client) (interface{}, error) {

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
	ByTagTagRunner.Flag(`includeTagsWithPrefix`, `If included, all tags matching this prefix will be returned. If not included, no tags will be returned.`).StringVar(ByTagTagRunner.includeTagsWithPrefix)
	ByTagTagRunner.Flag(`matchAll`, `If set to 'true', resources having all the tags specified in the 'tags' parameter are returned. Otherwise, resources having any of the tags are returned.`).StringVar(ByTagTagRunner.matchAll)
	ByTagTagRunner.Flag(`resourceType`, `Search among a single resource type.`).Required().StringVar(&ByTagTagRunner.resourceType)
	ByTagTagRunner.FlagPattern(`tags\.(\d+)`, `The tags which must be present on the resource.`).Required().Capture(&ByTagTagRunner.tagsPos).StringsVar(&ByTagTagRunner.tags)
	ByTagTagRunner.Flag(`withDeleted`, `If set to 'true', tags for deleted resources will also be returned. Default value is 'false'.`).StringVar(ByTagTagRunner.withDeleted)
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

type ShowTaskRunner struct {
	cloudId    string
	id         string
	instanceId string
	view       *string
}

func (r *ShowTaskRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowTask(r.cloudId, r.id, r.instanceId, options)
}

// Register all Task actions
func registerTaskCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Task", `Tasks represent processes that happen (or have happened) asynchronously within the context of an Instance.`)

	ShowTaskRunner := new(ShowTaskTaskRunner)
	ShowTaskCmd := resCmd.Command("ShowTask", `Displays information of a given task within the context of an instance.`)
	ShowTaskRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowTaskRunner.cloudId)
	ShowTaskRunner.Flag(`id`, ``).Required().StringVar(&ShowTaskRunner.id)
	ShowTaskRunner.Flag(`instanceId`, ``).Required().StringVar(&ShowTaskRunner.instanceId)
	ShowTaskRunner.Flag(`view`, ``).StringVar(ShowTaskRunner.view)
	registry[ShowTaskCmd.FullCommand()] = ShowTaskRunner
}

/****** User ******/

type CreateUserRunner struct {
	userCompany              string
	userEmail                string
	userFirstName            string
	userIdentityProviderHref *string
	userLastName             string
	userPassword             *string
	userPhone                string
	userPrincipalUid         *string
	userTimezoneName         *string
}

func (r *CreateUserRunner) Run(c *Client) (interface{}, error) {

	/** Handle user parameter **/
	var user UserParam
	// Load JSON if provided
	if len(r.userJson) > 0 {
		if err := Json.Unmarshal(r.userJson, &user); err != nil {
			return nil, fmt.Errorf("Could not load user JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if len(r.UserCompany) > 0 {
		user.company = r.UserCompany
	}

	if len(r.UserEmail) > 0 {
		user.email = r.UserEmail
	}

	if len(r.UserFirstName) > 0 {
		user.firstName = r.UserFirstName
	}

	if r.UserIdentityProviderHref != nil {
		user.identityProviderHref = r.UserIdentityProviderHref
	}

	if len(r.UserLastName) > 0 {
		user.lastName = r.UserLastName
	}

	if r.UserPassword != nil {
		user.password = r.UserPassword
	}

	if len(r.UserPhone) > 0 {
		user.phone = r.UserPhone
	}

	if r.UserPrincipalUid != nil {
		user.principalUid = r.UserPrincipalUid
	}

	if r.UserTimezoneName != nil {
		user.timezoneName = r.UserTimezoneName
	}

	return c.CreateUser(&user)
}

type IndexUsersRunner struct {
	filter    []string
	filterPos []string
}

func (r *IndexUsersRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter

	return c.IndexUsers(options)
}

type ShowUserRunner struct {
	id string
}

func (r *ShowUserRunner) Run(c *Client) (interface{}, error) {
	return c.ShowUser(r.id)
}

type UpdateUserRunner struct {
	id                       string
	userCompany              *string
	userCurrentEmail         string
	userCurrentPassword      *string
	userFirstName            *string
	userIdentityProviderHref *string
	userLastName             *string
	userNewEmail             *string
	userNewPassword          *string
	userPhone                *string
	userPrincipalUid         *string
	userTimezoneName         *string
}

func (r *UpdateUserRunner) Run(c *Client) (interface{}, error) {

	/** Handle user parameter **/
	var user UserParam2
	// Load JSON if provided
	if len(r.userJson) > 0 {
		if err := Json.Unmarshal(r.userJson, &user); err != nil {
			return nil, fmt.Errorf("Could not load user JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.UserCompany != nil {
		user.company = r.UserCompany
	}

	if len(r.UserCurrentEmail) > 0 {
		user.currentEmail = r.UserCurrentEmail
	}

	if r.UserCurrentPassword != nil {
		user.currentPassword = r.UserCurrentPassword
	}

	if r.UserFirstName != nil {
		user.firstName = r.UserFirstName
	}

	if r.UserIdentityProviderHref != nil {
		user.identityProviderHref = r.UserIdentityProviderHref
	}

	if r.UserLastName != nil {
		user.lastName = r.UserLastName
	}

	if r.UserNewEmail != nil {
		user.newEmail = r.UserNewEmail
	}

	if r.UserNewPassword != nil {
		user.newPassword = r.UserNewPassword
	}

	if r.UserPhone != nil {
		user.phone = r.UserPhone
	}

	if r.UserPrincipalUid != nil {
		user.principalUid = r.UserPrincipalUid
	}

	if r.UserTimezoneName != nil {
		user.timezoneName = r.UserTimezoneName
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
	CreateUserRunner.Flag(`user.identityProviderHref`, `The RightScale API href of the Identity Provider through which this user will login to RightScale. Required to create an SSO-authenticated user.`).StringVar(CreateUserRunner.userIdentityProviderHref)
	CreateUserRunner.Flag(`user.lastName`, ``).Required().StringVar(&CreateUserRunner.userLastName)
	CreateUserRunner.Flag(`user.password`, `The password of this user. Required to create a password-authenticated user.`).StringVar(CreateUserRunner.userPassword)
	CreateUserRunner.Flag(`user.phone`, ``).Required().StringVar(&CreateUserRunner.userPhone)
	CreateUserRunner.Flag(`user.principalUid`, `The principal identifier (SAML NameID or OpenID identity URL) of this user. Required to create an SSO-authenticated user.`).StringVar(CreateUserRunner.userPrincipalUid)
	CreateUserRunner.Flag(`user.timezoneName`, `This can be in the form of country/region or timezone name. For example 'America/Los_Angeles' or 'GB' or 'UTC'. A complete list of acceptable values is available in the Settings > User Settings > Preferences page.`).StringVar(CreateUserRunner.userTimezoneName)
	registry[CreateUserCmd.FullCommand()] = CreateUserRunner

	IndexUsersRunner := new(IndexUsersUserRunner)
	IndexUsersCmd := resCmd.Command("IndexUsers", `List the users available to the account the user is logged in to`)
	IndexUsersRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexUsersRunner.filterPos).StringsVar(IndexUsersRunner.filter)
	registry[IndexUsersCmd.FullCommand()] = IndexUsersRunner

	ShowUserRunner := new(ShowUserUserRunner)
	ShowUserCmd := resCmd.Command("ShowUser", `Show information about a single user.`)
	ShowUserRunner.Flag(`id`, ``).Required().StringVar(&ShowUserRunner.id)
	registry[ShowUserCmd.FullCommand()] = ShowUserRunner

	UpdateUserRunner := new(UpdateUserUserRunner)
	UpdateUserCmd := resCmd.Command("UpdateUser", `Update a user's contact information, change her password, or update SSO her settings`)
	UpdateUserRunner.Flag(`id`, ``).Required().StringVar(&UpdateUserRunner.id)
	UpdateUserRunner.Flag(`user.company`, ``).StringVar(UpdateUserRunner.userCompany)
	UpdateUserRunner.Flag(`user.currentEmail`, `The existing email of this user.`).Required().StringVar(&UpdateUserRunner.userCurrentEmail)
	UpdateUserRunner.Flag(`user.currentPassword`, `The current password for the user.`).StringVar(UpdateUserRunner.userCurrentPassword)
	UpdateUserRunner.Flag(`user.firstName`, ``).StringVar(UpdateUserRunner.userFirstName)
	UpdateUserRunner.Flag(`user.identityProviderHref`, `The updated RightScale API href of the associated Identity Provider.`).StringVar(UpdateUserRunner.userIdentityProviderHref)
	UpdateUserRunner.Flag(`user.lastName`, ``).StringVar(UpdateUserRunner.userLastName)
	UpdateUserRunner.Flag(`user.newEmail`, `The updated email of this user.`).StringVar(UpdateUserRunner.userNewEmail)
	UpdateUserRunner.Flag(`user.newPassword`, `The new password for this user.`).StringVar(UpdateUserRunner.userNewPassword)
	UpdateUserRunner.Flag(`user.phone`, ``).StringVar(UpdateUserRunner.userPhone)
	UpdateUserRunner.Flag(`user.principalUid`, `The updated principal identifier (SAML NameID or OpenID identity URL) of this user.`).StringVar(UpdateUserRunner.userPrincipalUid)
	UpdateUserRunner.Flag(`user.timezoneName`, `This can be in the form of country/region or timezone name. For example 'America/Los_Angeles' or 'GB' or 'UTC'. A complete list of acceptable values is available in the Settings > User Settings > Preferences page.`).StringVar(UpdateUserRunner.userTimezoneName)
	registry[UpdateUserCmd.FullCommand()] = UpdateUserRunner
}

/****** UserData ******/

type ShowUserDataRunner struct {
}

func (r *ShowUserDataRunner) Run(c *Client) (interface{}, error) {
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

type CreateVolumeRunner struct {
	cloudId                        string
	volumeDatacenterHref           *string
	volumeDeploymentHref           *string
	volumeDescription              *string
	volumeEncrypted                *string
	volumeIops                     *string
	volumeName                     string
	volumeParentVolumeSnapshotHref *string
	volumePlacementGroupHref       *string
	volumeSize                     *string
	volumeVolumeTypeHref           *string
}

func (r *CreateVolumeRunner) Run(c *Client) (interface{}, error) {

	/** Handle volume parameter **/
	var volume VolumeParam
	// Load JSON if provided
	if len(r.volumeJson) > 0 {
		if err := Json.Unmarshal(r.volumeJson, &volume); err != nil {
			return nil, fmt.Errorf("Could not load volume JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.VolumeDatacenterHref != nil {
		volume.datacenterHref = r.VolumeDatacenterHref
	}

	if r.VolumeDeploymentHref != nil {
		volume.deploymentHref = r.VolumeDeploymentHref
	}

	if r.VolumeDescription != nil {
		volume.description = r.VolumeDescription
	}

	if r.VolumeEncrypted != nil {
		volume.encrypted = r.VolumeEncrypted
	}

	if r.VolumeIops != nil {
		volume.iops = r.VolumeIops
	}

	if len(r.VolumeName) > 0 {
		volume.name = r.VolumeName
	}

	if r.VolumeParentVolumeSnapshotHref != nil {
		volume.parentVolumeSnapshotHref = r.VolumeParentVolumeSnapshotHref
	}

	if r.VolumePlacementGroupHref != nil {
		volume.placementGroupHref = r.VolumePlacementGroupHref
	}

	if r.VolumeSize != nil {
		volume.size = r.VolumeSize
	}

	if r.VolumeVolumeTypeHref != nil {
		volume.volumeTypeHref = r.VolumeVolumeTypeHref
	}

	return c.CreateVolume(r.cloudId, &volume)
}

type DestroyVolumeRunner struct {
	cloudId string
	id      string
}

func (r *DestroyVolumeRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyVolume(r.cloudId, r.id)
}

type IndexVolumesRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexVolumesRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexVolumes(r.cloudId, options)
}

type ShowVolumeRunner struct {
	cloudId string
	id      string
	view    *string
}

func (r *ShowVolumeRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowVolume(r.cloudId, r.id, options)
}

// Register all Volume actions
func registerVolumeCmds(app *kingpin.Application) {
	resCmd := app.Cmd("Volume", `A Volume provides a highly reliable, efficient and persistent storage solution that can be mounted to a cloud instance (in the same datacenter / zone).`)

	CreateVolumeRunner := new(CreateVolumeVolumeRunner)
	CreateVolumeCmd := resCmd.Command("CreateVolume", `Creates a new volume.`)
	CreateVolumeRunner.Flag(`cloudId`, ``).Required().StringVar(&CreateVolumeRunner.cloudId)
	CreateVolumeRunner.Flag(`volume.datacenterHref`, `The href of the Datacenter / Zone that the Volume will be in. This parameter is required for non-OpenStack clouds.`).StringVar(CreateVolumeRunner.volumeDatacenterHref)
	CreateVolumeRunner.Flag(`volume.deploymentHref`, `The href of the Deployment that owns this Volume.`).StringVar(CreateVolumeRunner.volumeDeploymentHref)
	CreateVolumeRunner.Flag(`volume.description`, `The description of the Volume to be created.`).StringVar(CreateVolumeRunner.volumeDescription)
	CreateVolumeRunner.Flag(`volume.encrypted`, `A flag indicating whether Volume should be encrypted. Only available on clouds supporting volume encryption.`).StringVar(CreateVolumeRunner.volumeEncrypted)
	CreateVolumeRunner.Flag(`volume.iops`, `The number of IOPS (I/O Operations Per Second) this Volume should support. Only available on clouds supporting performance provisioning.`).StringVar(CreateVolumeRunner.volumeIops)
	CreateVolumeRunner.Flag(`volume.name`, `The name for the Volume to be created.`).Required().StringVar(&CreateVolumeRunner.volumeName)
	CreateVolumeRunner.Flag(`volume.parentVolumeSnapshotHref`, `The href of the snapshot from which the volume will be created.`).StringVar(CreateVolumeRunner.volumeParentVolumeSnapshotHref)
	CreateVolumeRunner.Flag(`volume.placementGroupHref`, `The href of the Placement Group.`).StringVar(CreateVolumeRunner.volumePlacementGroupHref)
	CreateVolumeRunner.Flag(`volume.size`, `The size of a Volume to be created in gigabytes (GB). Some Volume Types have predefined sizes and do not allow selecting a custom size on Volume creation.`).StringVar(CreateVolumeRunner.volumeSize)
	CreateVolumeRunner.Flag(`volume.volumeTypeHref`, `The href of the volume type. A Name, Resource UID and optional Size is associated with a Volume Type.`).StringVar(CreateVolumeRunner.volumeVolumeTypeHref)
	registry[CreateVolumeCmd.FullCommand()] = CreateVolumeRunner

	DestroyVolumeRunner := new(DestroyVolumeVolumeRunner)
	DestroyVolumeCmd := resCmd.Command("DestroyVolume", `Deletes a given volume.`)
	DestroyVolumeRunner.Flag(`cloudId`, ``).Required().StringVar(&DestroyVolumeRunner.cloudId)
	DestroyVolumeRunner.Flag(`id`, ``).Required().StringVar(&DestroyVolumeRunner.id)
	registry[DestroyVolumeCmd.FullCommand()] = DestroyVolumeRunner

	IndexVolumesRunner := new(IndexVolumesVolumeRunner)
	IndexVolumesCmd := resCmd.Command("IndexVolumes", `Lists volumes.`)
	IndexVolumesRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexVolumesRunner.cloudId)
	IndexVolumesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexVolumesRunner.filterPos).StringsVar(IndexVolumesRunner.filter)
	IndexVolumesRunner.Flag(`view`, ``).StringVar(IndexVolumesRunner.view)
	registry[IndexVolumesCmd.FullCommand()] = IndexVolumesRunner

	ShowVolumeRunner := new(ShowVolumeVolumeRunner)
	ShowVolumeCmd := resCmd.Command("ShowVolume", `Displays information about a single volume.`)
	ShowVolumeRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowVolumeRunner.cloudId)
	ShowVolumeRunner.Flag(`id`, ``).Required().StringVar(&ShowVolumeRunner.id)
	ShowVolumeRunner.Flag(`view`, ``).StringVar(ShowVolumeRunner.view)
	registry[ShowVolumeCmd.FullCommand()] = ShowVolumeRunner
}

/****** VolumeAttachment ******/

type CreateVolumeAttachmentRunner struct {
	cloudId                      string
	instanceId                   string
	volumeAttachmentDevice       *string
	volumeAttachmentInstanceHref *string
	volumeAttachmentVolumeHref   *string
}

func (r *CreateVolumeAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle volumeAttachment parameter **/
	var volumeAttachment VolumeAttachmentParam
	// Load JSON if provided
	if len(r.volumeAttachmentJson) > 0 {
		if err := Json.Unmarshal(r.volumeAttachmentJson, &volumeAttachment); err != nil {
			return nil, fmt.Errorf("Could not load volumeAttachment JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.VolumeAttachmentDevice != nil {
		volumeAttachment.device = r.VolumeAttachmentDevice
	}

	if r.VolumeAttachmentInstanceHref != nil {
		volumeAttachment.instanceHref = r.VolumeAttachmentInstanceHref
	}

	if r.VolumeAttachmentVolumeHref != nil {
		volumeAttachment.volumeHref = r.VolumeAttachmentVolumeHref
	}

	return c.CreateVolumeAttachment(r.cloudId, r.instanceId, &volumeAttachment)
}

type DestroyVolumeAttachmentRunner struct {
	cloudId    string
	force      *string
	id         string
	instanceId string
}

func (r *DestroyVolumeAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.force != nil {
		options["force"] = *r.force
	}

	return c.DestroyVolumeAttachment(r.cloudId, r.id, r.instanceId, options)
}

type IndexVolumeAttachmentsRunner struct {
	cloudId    string
	filter     []string
	filterPos  []string
	instanceId string
	view       *string
}

func (r *IndexVolumeAttachmentsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexVolumeAttachments(r.cloudId, r.instanceId, options)
}

type ShowVolumeAttachmentRunner struct {
	cloudId    string
	id         string
	instanceId string
	view       *string
}

func (r *ShowVolumeAttachmentRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowVolumeAttachment(r.cloudId, r.id, r.instanceId, options)
}

// Register all VolumeAttachment actions
func registerVolumeAttachmentCmds(app *kingpin.Application) {
	resCmd := app.Cmd("VolumeAttachment", `A VolumeAttachment represents a relationship between a volume and an instance.`)

	CreateVolumeAttachmentRunner := new(CreateVolumeAttachmentVolumeAttachmentRunner)
	CreateVolumeAttachmentCmd := resCmd.Command("CreateVolumeAttachment", `Creates a new volume attachment.`)
	CreateVolumeAttachmentRunner.Flag(`cloudId`, ``).Required().StringVar(&CreateVolumeAttachmentRunner.cloudId)
	CreateVolumeAttachmentRunner.Flag(`instanceId`, ``).Required().StringVar(&CreateVolumeAttachmentRunner.instanceId)
	CreateVolumeAttachmentRunner.Flag(`volumeAttachment.device`, `The device location where the volume will be mounted. Value must be of format /dev/xvd[bcefghij]. This is not reliable and will be deprecated.`).StringVar(CreateVolumeAttachmentRunner.volumeAttachmentDevice)
	CreateVolumeAttachmentRunner.Flag(`volumeAttachment.instanceHref`, `The href of the instance to which the volume will be attached to.`).StringVar(CreateVolumeAttachmentRunner.volumeAttachmentInstanceHref)
	CreateVolumeAttachmentRunner.Flag(`volumeAttachment.volumeHref`, `The href of the volume to be attached.`).StringVar(CreateVolumeAttachmentRunner.volumeAttachmentVolumeHref)
	registry[CreateVolumeAttachmentCmd.FullCommand()] = CreateVolumeAttachmentRunner

	DestroyVolumeAttachmentRunner := new(DestroyVolumeAttachmentVolumeAttachmentRunner)
	DestroyVolumeAttachmentCmd := resCmd.Command("DestroyVolumeAttachment", `Deletes a given volume attachment.`)
	DestroyVolumeAttachmentRunner.Flag(`cloudId`, ``).Required().StringVar(&DestroyVolumeAttachmentRunner.cloudId)
	DestroyVolumeAttachmentRunner.Flag(`force`, `Specifies whether to force the detachment of a volume.`).StringVar(DestroyVolumeAttachmentRunner.force)
	DestroyVolumeAttachmentRunner.Flag(`id`, ``).Required().StringVar(&DestroyVolumeAttachmentRunner.id)
	DestroyVolumeAttachmentRunner.Flag(`instanceId`, ``).Required().StringVar(&DestroyVolumeAttachmentRunner.instanceId)
	registry[DestroyVolumeAttachmentCmd.FullCommand()] = DestroyVolumeAttachmentRunner

	IndexVolumeAttachmentsRunner := new(IndexVolumeAttachmentsVolumeAttachmentRunner)
	IndexVolumeAttachmentsCmd := resCmd.Command("IndexVolumeAttachments", `Lists all volume attachments.`)
	IndexVolumeAttachmentsRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexVolumeAttachmentsRunner.cloudId)
	IndexVolumeAttachmentsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexVolumeAttachmentsRunner.filterPos).StringsVar(IndexVolumeAttachmentsRunner.filter)
	IndexVolumeAttachmentsRunner.Flag(`instanceId`, ``).Required().StringVar(&IndexVolumeAttachmentsRunner.instanceId)
	IndexVolumeAttachmentsRunner.Flag(`view`, ``).StringVar(IndexVolumeAttachmentsRunner.view)
	registry[IndexVolumeAttachmentsCmd.FullCommand()] = IndexVolumeAttachmentsRunner

	ShowVolumeAttachmentRunner := new(ShowVolumeAttachmentVolumeAttachmentRunner)
	ShowVolumeAttachmentCmd := resCmd.Command("ShowVolumeAttachment", `Displays information about a single volume attachment.`)
	ShowVolumeAttachmentRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowVolumeAttachmentRunner.cloudId)
	ShowVolumeAttachmentRunner.Flag(`id`, ``).Required().StringVar(&ShowVolumeAttachmentRunner.id)
	ShowVolumeAttachmentRunner.Flag(`instanceId`, ``).Required().StringVar(&ShowVolumeAttachmentRunner.instanceId)
	ShowVolumeAttachmentRunner.Flag(`view`, ``).StringVar(ShowVolumeAttachmentRunner.view)
	registry[ShowVolumeAttachmentCmd.FullCommand()] = ShowVolumeAttachmentRunner
}

/****** VolumeSnapshot ******/

type CreateVolumeSnapshotRunner struct {
	cloudId                        string
	volumeId                       string
	volumeSnapshotDeploymentHref   *string
	volumeSnapshotDescription      *string
	volumeSnapshotName             string
	volumeSnapshotParentVolumeHref *string
}

func (r *CreateVolumeSnapshotRunner) Run(c *Client) (interface{}, error) {

	/** Handle volumeSnapshot parameter **/
	var volumeSnapshot VolumeSnapshotParam
	// Load JSON if provided
	if len(r.volumeSnapshotJson) > 0 {
		if err := Json.Unmarshal(r.volumeSnapshotJson, &volumeSnapshot); err != nil {
			return nil, fmt.Errorf("Could not load volumeSnapshot JSON: %s", err.Error())
		}
	}

	// Override with any provided basic field
	if r.VolumeSnapshotDeploymentHref != nil {
		volumeSnapshot.deploymentHref = r.VolumeSnapshotDeploymentHref
	}

	if r.VolumeSnapshotDescription != nil {
		volumeSnapshot.description = r.VolumeSnapshotDescription
	}

	if len(r.VolumeSnapshotName) > 0 {
		volumeSnapshot.name = r.VolumeSnapshotName
	}

	if r.VolumeSnapshotParentVolumeHref != nil {
		volumeSnapshot.parentVolumeHref = r.VolumeSnapshotParentVolumeHref
	}

	return c.CreateVolumeSnapshot(r.cloudId, r.volumeId, &volumeSnapshot)
}

type DestroyVolumeSnapshotRunner struct {
	cloudId  string
	id       string
	volumeId string
}

func (r *DestroyVolumeSnapshotRunner) Run(c *Client) (interface{}, error) {
	return c.DestroyVolumeSnapshot(r.cloudId, r.id, r.volumeId)
}

type IndexVolumeSnapshotsRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      *string
	volumeId  string
}

func (r *IndexVolumeSnapshotsRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexVolumeSnapshots(r.cloudId, r.volumeId, options)
}

type ShowVolumeSnapshotRunner struct {
	cloudId  string
	id       string
	view     *string
	volumeId string
}

func (r *ShowVolumeSnapshotRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowVolumeSnapshot(r.cloudId, r.id, r.volumeId, options)
}

// Register all VolumeSnapshot actions
func registerVolumeSnapshotCmds(app *kingpin.Application) {
	resCmd := app.Cmd("VolumeSnapshot", `A VolumeSnapshot represents a Cloud storage volume at a particular point in time`)

	CreateVolumeSnapshotRunner := new(CreateVolumeSnapshotVolumeSnapshotRunner)
	CreateVolumeSnapshotCmd := resCmd.Command("CreateVolumeSnapshot", `Creates a new volume_snapshot.`)
	CreateVolumeSnapshotRunner.Flag(`cloudId`, ``).Required().StringVar(&CreateVolumeSnapshotRunner.cloudId)
	CreateVolumeSnapshotRunner.Flag(`volumeId`, ``).Required().StringVar(&CreateVolumeSnapshotRunner.volumeId)
	CreateVolumeSnapshotRunner.Flag(`volumeSnapshot.deploymentHref`, `The href of the Deployment that owns this Volume Snapshot.`).StringVar(CreateVolumeSnapshotRunner.volumeSnapshotDeploymentHref)
	CreateVolumeSnapshotRunner.Flag(`volumeSnapshot.description`, `The description for the Volume Snapshot to be created.`).StringVar(CreateVolumeSnapshotRunner.volumeSnapshotDescription)
	CreateVolumeSnapshotRunner.Flag(`volumeSnapshot.name`, `The name for the Volume Snapshot to be created.`).Required().StringVar(&CreateVolumeSnapshotRunner.volumeSnapshotName)
	CreateVolumeSnapshotRunner.Flag(`volumeSnapshot.parentVolumeHref`, `The href of the Volume from which the Volume Snapshot will be created.`).StringVar(CreateVolumeSnapshotRunner.volumeSnapshotParentVolumeHref)
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
	IndexVolumeSnapshotsRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexVolumeSnapshotsRunner.filterPos).StringsVar(IndexVolumeSnapshotsRunner.filter)
	IndexVolumeSnapshotsRunner.Flag(`view`, ``).StringVar(IndexVolumeSnapshotsRunner.view)
	IndexVolumeSnapshotsRunner.Flag(`volumeId`, ``).Required().StringVar(&IndexVolumeSnapshotsRunner.volumeId)
	registry[IndexVolumeSnapshotsCmd.FullCommand()] = IndexVolumeSnapshotsRunner

	ShowVolumeSnapshotRunner := new(ShowVolumeSnapshotVolumeSnapshotRunner)
	ShowVolumeSnapshotCmd := resCmd.Command("ShowVolumeSnapshot", `Displays information about a single volume_snapshot.`)
	ShowVolumeSnapshotRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowVolumeSnapshotRunner.cloudId)
	ShowVolumeSnapshotRunner.Flag(`id`, ``).Required().StringVar(&ShowVolumeSnapshotRunner.id)
	ShowVolumeSnapshotRunner.Flag(`view`, ``).StringVar(ShowVolumeSnapshotRunner.view)
	ShowVolumeSnapshotRunner.Flag(`volumeId`, ``).Required().StringVar(&ShowVolumeSnapshotRunner.volumeId)
	registry[ShowVolumeSnapshotCmd.FullCommand()] = ShowVolumeSnapshotRunner
}

/****** VolumeType ******/

type IndexVolumeTypesRunner struct {
	cloudId   string
	filter    []string
	filterPos []string
	view      *string
}

func (r *IndexVolumeTypesRunner) Run(c *Client) (interface{}, error) {

	/** Handle filter parameter **/
	var filter []string

	for i, v := range r.filter {
		pos, err := strconv.Atoi(r.filterPos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for filter array", r.filterPos[i])
		}
		filter[pos] = v
	}

	/** Handle optional parameters **/
	options := make(ApiParams)
	options["filter"] = filter
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.IndexVolumeTypes(r.cloudId, options)
}

type ShowVolumeTypeRunner struct {
	cloudId string
	id      string
	view    *string
}

func (r *ShowVolumeTypeRunner) Run(c *Client) (interface{}, error) {

	/** Handle optional parameters **/
	options := make(ApiParams)
	if r.view != nil {
		options["view"] = *r.view
	}

	return c.ShowVolumeType(r.cloudId, r.id, options)
}

// Register all VolumeType actions
func registerVolumeTypeCmds(app *kingpin.Application) {
	resCmd := app.Cmd("VolumeType", `A VolumeType describes the type of volume, particularly the size.`)

	IndexVolumeTypesRunner := new(IndexVolumeTypesVolumeTypeRunner)
	IndexVolumeTypesCmd := resCmd.Command("IndexVolumeTypes", `Lists Volume Types.`)
	IndexVolumeTypesRunner.Flag(`cloudId`, ``).Required().StringVar(&IndexVolumeTypesRunner.cloudId)
	IndexVolumeTypesRunner.FlagPattern(`filter\.(\d+)`, ``).Capture(&IndexVolumeTypesRunner.filterPos).StringsVar(IndexVolumeTypesRunner.filter)
	IndexVolumeTypesRunner.Flag(`view`, ``).StringVar(IndexVolumeTypesRunner.view)
	registry[IndexVolumeTypesCmd.FullCommand()] = IndexVolumeTypesRunner

	ShowVolumeTypeRunner := new(ShowVolumeTypeVolumeTypeRunner)
	ShowVolumeTypeCmd := resCmd.Command("ShowVolumeType", `Displays information about a single Volume Type.`)
	ShowVolumeTypeRunner.Flag(`cloudId`, ``).Required().StringVar(&ShowVolumeTypeRunner.cloudId)
	ShowVolumeTypeRunner.Flag(`id`, ``).Required().StringVar(&ShowVolumeTypeRunner.id)
	ShowVolumeTypeRunner.Flag(`view`, ``).StringVar(ShowVolumeTypeRunner.view)
	registry[ShowVolumeTypeCmd.FullCommand()] = ShowVolumeTypeRunner
}
