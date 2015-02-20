//************************************************************************//
//                     rsc - RightScale API 1.5 command line tool
//
// Generated Feb 19, 2015 at 11:24pm (PST)
// Command:
// $ api15gen -metadata=../../rsapi15 -output=../../rsapi15
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package rsapi15

import "regexp"

// API 1.5 resource commands
// Each command contains sub-commands for all resource actions
var commands = map[string]*ResourceCmd{
	"Account": &ResourceCmd{
		Name:        "Account",
		Description: ``,
		HrefRegexp:  regexp.MustCompile(`/api/accounts/[[:alnum:]]`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single Account.`,
				Flags:       []*ActionFlag{},
			},
		},
	},
	"AccountGroup": &ResourceCmd{
		Name:        "AccountGroup",
		Description: `An Account Group specifies which RightScale accounts will have access to import a shared RightScale component (e.g. ServerTemplate, RightScript, etc.) from the MultiCloud Marketplace.`,
		HrefRegexp:  regexp.MustCompile(`(/api/account_groups/[[:alnum:]]|/api/account_groups)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "index",
				Description: `Lists the AccountGroups owned by this Account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single AccountGroup.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"Alert": &ResourceCmd{
		Name:        "Alert",
		Description: `An Alert represents an AlertSpec bound to a running Instance.`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/instances/[[:alnum:]]/alerts/[[:alnum:]]|/api/servers/[[:alnum:]]/alerts/[[:alnum:]]|/api/server_arrays/[[:alnum:]]/alerts/[[:alnum:]]|/api/deployments/[[:alnum:]]/alerts/[[:alnum:]]|/api/alerts/[[:alnum:]]|/api/clouds/[[:alnum:]]/instances/[[:alnum:]]/alerts|/api/servers/[[:alnum:]]/alerts|/api/server_arrays/[[:alnum:]]/alerts|/api/deployments/[[:alnum:]]/alerts|/api/alerts)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "disable",
				Description: `Disables the Alert indefinitely. Idempotent.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "enable",
				Description: `Enables the Alert indefinitely. Idempotent.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists all Alerts.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "quench",
				Description: `Suppresses the Alert from being triggered for a given time period. Idempotent.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "duration",
						Description: `The time period in seconds to suppress Alert from being triggered.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Shows the attributes of a specified Alert.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"AlertSpec": &ResourceCmd{
		Name:        "AlertSpec",
		Description: `An AlertSpec defines the conditions under which an Alert is triggered and escalated.`,
		HrefRegexp:  regexp.MustCompile(`(/api/servers/[[:alnum:]]/alert_specs/[[:alnum:]]|/api/server_arrays/[[:alnum:]]/alert_specs/[[:alnum:]]|/api/server_templates/[[:alnum:]]/alert_specs/[[:alnum:]]|/api/alert_specs/[[:alnum:]]|/api/servers/[[:alnum:]]/alert_specs|/api/server_arrays/[[:alnum:]]/alert_specs|/api/server_templates/[[:alnum:]]/alert_specs|/api/alert_specs)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Creates a new AlertSpec with the given parameters.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "alert_spec[escalation_name]",
						Description: `Escalate to the named alert escalation when the alert is triggered. Must either escalate or vote.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "alert_spec[subject_href]",
						Description: `The href of the resource that this AlertSpec should be associated with. The subject can be a ServerTemplate, Server, ServerArray, or Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "alert_spec[description]",
						Description: `The description of the AlertSpec.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "alert_spec[threshold]",
						Description: `The threshold of the condition sentence.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "alert_spec[condition]",
						Description: `The condition (operator) in the condition sentence.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{">", ">=", "<", "<=", "==", "!="},
					},
					&ActionFlag{
						Name:        "alert_spec[vote_type]",
						Description: `Vote to grow or shrink a ServerArray when the alert is triggered. Must either escalate or vote.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"grow", "shrink"},
					},
					&ActionFlag{
						Name:        "alert_spec[variable]",
						Description: `The RRD variable of the condition sentence.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "alert_spec[duration]",
						Description: `The duration in minutes of the condition sentence.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "alert_spec[vote_tag]",
						Description: `Should correspond to a vote tag on a ServerArray if vote to grow or shrink.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "alert_spec[name]",
						Description: `The name of the AlertSpec.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "alert_spec[file]",
						Description: `The RRD path/file_name of the condition sentence.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes a given AlertSpec.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `No description provided for index.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "with_inherited",
						Description: `Flag indicating whether or not to include AlertSpecs from the ServerTemplate in the index.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `No description provided for show.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Updates an AlertSpec with the given parameters.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "alert_spec[escalation_name]",
						Description: `Escalate to the named alert escalation when the alert is triggered.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "alert_spec[description]",
						Description: `The description of the AlertSpec.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "alert_spec[threshold]",
						Description: `The threshold of the condition sentence.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "alert_spec[vote_type]",
						Description: `Vote to grow or shrink a ServerArray when the alert is triggered.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"grow", "shrink"},
					},
					&ActionFlag{
						Name:        "alert_spec[condition]",
						Description: `The condition (operator) in the condition sentence.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{">", ">=", "<", "<=", "==", "!="},
					},
					&ActionFlag{
						Name:        "alert_spec[variable]",
						Description: `The RRD variable of the condition sentence.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "alert_spec[duration]",
						Description: `The duration in minutes of the condition sentence.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "alert_spec[vote_tag]",
						Description: `Should correspond to a vote tag on a ServerArray if vote to grow or shrink.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "alert_spec[file]",
						Description: `The RRD path/file_name of the condition sentence.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "alert_spec[name]",
						Description: `The name of the AlertSpec.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"AuditEntry": &ResourceCmd{
		Name:        "AuditEntry",
		Description: `An Audit Entry can be used to track various activities of a resource.`,
		HrefRegexp:  regexp.MustCompile(`(/api/audit_entries/[[:alnum:]]|/api/audit_entries)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "append",
				Description: `Updates the summary and appends more details to a given AuditEntry`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "summary",
						Description: `The updated summary for the audit entry, maximum length is 255 characters.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "detail",
						Description: `The details to be appended to the audit entry record.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "notify",
						Description: `The event notification category. Defaults to 'None'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "offset",
						Description: `The offset where the new details should be appended to in the audit entry's existing details section. Also used in ordering of summary updates. Defaults to end.`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "create",
				Description: `Creates a new AuditEntry with the given parameters.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "audit_entry[auditee_href]",
						Description: `The href of the resource that this audit entry should be associated with (e.g. an instance's href).`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "audit_entry[summary]",
						Description: `The summary of the audit entry to be created, maximum length is 255 characters.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "audit_entry[detail]",
						Description: `The initial details of the audit entry to be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "user_email",
						Description: `The email of the user (who created/triggered the audit entry). Only usable with instance role.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "notify",
						Description: `The event notification category. Defaults to 'None'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "detail",
				Description: `shows the details of a given AuditEntry.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists AuditEntries of the account`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "start_date",
						Description: `The start date for retrieving audit entries, the format must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g., 2011/06/25 00:00:00 +0000`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "end_date",
						Description: `The end date for retrieving audit entries (the format must be the same as start date). The time period between start and end date must be less than 3 months (93 days).`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "limit",
						Description: `Limit the audit entries to this number. The limit should >= 1 and <= 1000`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Lists the attributes of a given audit entry.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Updates the summary of a given AuditEntry.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "audit_entry[summary]",
						Description: `The updated summary for the audit entry, maximum length is 255 characters.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "audit_entry[offset]",
						Description: `The offset where the next details will be appended. Used in ordering of summary updates.`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "notify",
						Description: `The event notification category. Defaults to 'None'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"Backup": &ResourceCmd{
		Name:        "Backup",
		Description: ``,
		HrefRegexp:  regexp.MustCompile(`(/api/backups/[[:alnum:]]|/api/backups)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "cleanup",
				Description: `Deletes old backups that meet the given criteria`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "cloud_href",
						Description: `Backups belonging to only this cloud are considered for cleanup. Otherwise, all backups in the account with the same lineage will be considered.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "keep_last",
						Description: `The number of backups that should be kept.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "monthlies",
						Description: `The number of monthly backups(the latest one in each month) that should be kept.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "weeklies",
						Description: `The number of weekly backups(the latest one in each week) that should be kept.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "yearlies",
						Description: `The number of yearly backups(the latest one in each year) that should be kept.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "dailies",
						Description: `The number of daily backups(the latest one in each day) that should be kept.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "lineage",
						Description: `The lineage of the backups that are to be cleaned-up.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "create",
				Description: `Takes in an array of volumeattachmenthrefs and takes a snapshot of each.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "backup[volume_attachment_hrefs][]",
						Description: `List of volume attachment hrefs that are to be backed-up.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "backup[description]",
						Description: `The description to be set on each of the volume snapshots`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "backup[from_master]",
						Description: `Setting this to 'true' will create a tag 'rs_backup:from_master=true' on the snapshots so that one can filter them later.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "backup[lineage]",
						Description: `A unique value to create backups belonging to a particular system.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "backup[name]",
						Description: `The name to be set on each of the volume snapshots.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes a given backup by deleting all of its snapshots, this call will succeed even if the backup has not completed.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists all of the backups with the given lineage tag`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "lineage",
						Description: `Backups belonging to this lineage.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "restore",
				Description: `Restores the given Backup.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "backup[volume_type_href]",
						Description: `The href of the volume type. Each volume is created with this volume type instead of the default volume type for the cloud. A Name, Resource UID and optional Size is associated with a volume type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "backup[description]",
						Description: `Each volume is created with this description instead of the volume snapshot's description`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance_href",
						Description: `The instance href that the backup will be restored to.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "backup[iops]",
						Description: `The number of IOPS (I/O Operations Per Second) each volume should support. Only available on clouds supporting performance provisioning.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "backup[name]",
						Description: `Each volume is created with this name instead of the volume snapshot's name`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "backup[size]",
						Description: `Each volume is created with this size in gigabytes (GB) instead of the volume snapshot's size (must be equal or larger). Some volume types have predefined sizes and do not allow selecting a custom size on volume creation.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Lists the attributes of a given backup`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Updates the committed tag for all of the VolumeSnapshots in the given Backup to the given value.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "backup[committed]",
						Description: `Setting this to 'true' will update the 'rs_backup:committed=false' tag to 'rs_backup:committed=true' on all the snapshots.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},
		},
	},
	"ChildAccount": &ResourceCmd{
		Name:        "ChildAccount",
		Description: ``,
		HrefRegexp:  regexp.MustCompile(`/api/child_accounts`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Create an enterprise ChildAccount for this Account`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "child_account[cluster_href]",
						Description: `The href of the cluster in which to create the account. If not specified, will default to the cluster of the parent account.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "child_account[name]",
						Description: ``,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists the enterprise ChildAccounts available for this Account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Update an enterprise ChildAccount for this Account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "child_account[name]",
						Description: `The updated name for the account.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"Cloud": &ResourceCmd{
		Name:        "Cloud",
		Description: `Represents a Cloud (within the context of the account in the session).`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]|/api/clouds)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "index",
				Description: `Lists the clouds available to this account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single cloud`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},
		},
	},
	"CloudAccount": &ResourceCmd{
		Name:        "CloudAccount",
		Description: `Represents a Cloud Account (An association between the account and a cloud).`,
		HrefRegexp:  regexp.MustCompile(`(/api/cloud_accounts/[[:alnum:]]|/api/cloud_accounts)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Create a CloudAccount by passing in the respective credentials for each cloud.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "cloud_account[cloud_href]",
						Description: `The href of the cloud if it is known. For valid values see support portal link above.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "cloud_account[creds]",
						Description: ``,
						Type:        "map",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "cloud_account[token]",
						Description: `The cloud token to identify a private cloud`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Delete a CloudAccount.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists the CloudAccounts (non-aws) available to this Account.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "show",
				Description: ``,
				Flags:       []*ActionFlag{},
			},
		},
	},
	"Cookbook": &ResourceCmd{
		Name:        "Cookbook",
		Description: `Represents a given instance of a single cookbook.`,
		HrefRegexp:  regexp.MustCompile(`(/api/cookbooks/[[:alnum:]]|/api/cookbooks)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "destroy",
				Description: `Destroys a Cookbook. Only available for cookbooks that have no Cookbook Attachments.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "follow",
				Description: `Follows (or unfollows) a Cookbook. Only available for cookbooks that are in the Alternate namespace.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "value",
						Description: `Indicates if this action should follow (true) or unfollow (false) a Cookbook.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},

			&ActionCmd{
				Name:        "freeze",
				Description: `Freezes (or unfreezes) a Cookbook. Only available for cookbooks that are in the Primary namespace.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "value",
						Description: `Indicates if this action should freeze (true) or unfreeze (false) a Cookbook.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists the Cookbooks available to this account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended", "extended_designer"},
					},
				},
			},

			&ActionCmd{
				Name:        "obsolete",
				Description: `Marks a Cookbook as obsolete (or un-obsolete).`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "value",
						Description: `Indicates if this action should obsolete (true) or un-obsolete (false) a Cookbook.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single Cookbook.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended", "extended_designer"},
					},
				},
			},
		},
	},
	"CookbookAttachment": &ResourceCmd{
		Name:        "CookbookAttachment",
		Description: `Cookbook Attachment is used to associate a particular cookbook with a ServerTemplate. A Cookbook Attachment must be in place before a recipe can be bound to a runlist using RunnableBinding.`,
		HrefRegexp:  regexp.MustCompile(`(/api/cookbooks/[[:alnum:]]/cookbook_attachments/[[:alnum:]]|/api/server_templates/[[:alnum:]]/cookbook_attachments/[[:alnum:]]|/api/cookbook_attachments/[[:alnum:]]|/api/cookbooks/[[:alnum:]]/cookbook_attachments|/api/server_templates/[[:alnum:]]/cookbook_attachments|/api/cookbook_attachments)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Attach a cookbook to a given resource.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "cookbook_attachment[server_template_href]",
						Description: `The href of the server template to attach the cookbook to.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "cookbook_attachment[cookbook_href]",
						Description: `The href of the cookbook to attach.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Detach a cookbook from a given resource.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists Cookbook Attachments.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "multi_attach",
				Description: `Attach multiple cookbooks to a given resource.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "cookbook_attachments[server_template_href]",
						Description: `The href of the server template to attach the cookbooks to.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "cookbook_attachments[cookbook_hrefs][]",
						Description: `The hrefs of the cookbooks to be attached`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "multi_detach",
				Description: `Detach multiple cookbooks from a given resource.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "cookbook_attachments[cookbook_attachment_hrefs][]",
						Description: `The hrefs of the cookbook attachments to be detached`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Displays information about a single cookbook attachment to a ServerTemplate.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"Credential": &ResourceCmd{
		Name:        "Credential",
		Description: `A Credential provides an abstraction for sensitive textual information, such as passphrases or cloud credentials, whose value should be encrypted when stored in the database and not generally shown in the UI or through the API...`,
		HrefRegexp:  regexp.MustCompile(`(/api/credentials/[[:alnum:]]|/api/credentials)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Creates a new Credential with the given parameters.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "credential[description]",
						Description: `The description of the Credential to be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "credential[value]",
						Description: `The value of the Credential to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "credential[name]",
						Description: `The name of the Credential to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes a Credential.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists the Credentials available to this account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "sensitive"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single Credential. NOTE: Credential values may be updated through the API, but values cannot be retrieved via the API.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "sensitive"},
					},
				},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Updates attributes of a Credential.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "credential[description]",
						Description: `The updated description of the Credential.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "credential[value]",
						Description: `The updated value of the Credential.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "credential[name]",
						Description: `The updated name of the Credential.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"Datacenter": &ResourceCmd{
		Name:        "Datacenter",
		Description: `Datacenters represent isolated facilities within a cloud`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/datacenters/[[:alnum:]]|/api/clouds/[[:alnum:]]/datacenters)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "index",
				Description: `Lists all Datacenters for a particular cloud.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Displays information about a single Datacenter.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"Deployment": &ResourceCmd{
		Name:        "Deployment",
		Description: `Deployments represent logical groupings of related assets such as servers, server arrays, default configuration settings...etc.`,
		HrefRegexp:  regexp.MustCompile(`(/api/deployments/[[:alnum:]]|/api/deployments)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "clone",
				Description: `Clones a given deployment.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "deployment[server_tag_scope]",
						Description: `The routing scope for tags for servers in the cloned deployment.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"deployment", "account"},
					},
					&ActionFlag{
						Name:        "deployment[description]",
						Description: `The description for the cloned deployment.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "deployment[name]",
						Description: `The name for the cloned deployment.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "create",
				Description: `Creates a new deployment with the given parameters.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "deployment[server_tag_scope]",
						Description: `The routing scope for tags for servers in the deployment.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"deployment", "account"},
					},
					&ActionFlag{
						Name:        "deployment[description]",
						Description: `The description of the deployment to be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "deployment[name]",
						Description: `The name of the deployment to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes a given deployment.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists deployments of the account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs", "inputs_2_0"},
					},
				},
			},

			&ActionCmd{
				Name:        "lock",
				Description: `Locks a given deployment. Idempotent.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "servers",
				Description: `Lists the servers belonging to this deployment`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Lists the attributes of a given deployment.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs", "inputs_2_0"},
					},
				},
			},

			&ActionCmd{
				Name:        "unlock",
				Description: `Unlocks a given deployment. Idempotent.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Updates attributes of a given deployment.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "deployment[server_tag_scope]",
						Description: `The routing scope for tags for servers in the deployment.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"deployment", "account"},
					},
					&ActionFlag{
						Name:        "deployment[description]",
						Description: `The updated description for the deployment.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "deployment[name]",
						Description: `The updated name for the deployment.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"HealthCheck": &ResourceCmd{
		Name:        "HealthCheck",
		Description: ``,
		HrefRegexp:  regexp.MustCompile(`/api/health-check/`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "index",
				Description: `Check health of RightApi controllers`,
				Flags:       []*ActionFlag{},
			},
		},
	},
	"IdentityProvider": &ResourceCmd{
		Name:        "IdentityProvider",
		Description: `An Identity Provider represents a SAML identity provider (IdP) that is linked to your RightScale Enterprise account, and is trusted by the RightScale dashboard to authenticate your enterprise's end users...`,
		HrefRegexp:  regexp.MustCompile(`(/api/identity_providers/[[:alnum:]]|/api/identity_providers)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "index",
				Description: `Lists the identity providers associated with this enterprise account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show the specified identity provider, if associated with this enterprise account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"Image": &ResourceCmd{
		Name:        "Image",
		Description: `Images represent base VM image existing in a cloud`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/images/[[:alnum:]]|/api/clouds/[[:alnum:]]/images)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "index",
				Description: `Lists all Images for the given Cloud.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Shows information about a single Image.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"Input": &ResourceCmd{
		Name:        "Input",
		Description: `Inputs help extract dynamic information, usually specified at runtime, from repeatable configuration operations that can be codified.`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/instances/[[:alnum:]]/inputs|/api/deployments/[[:alnum:]]/inputs|/api/server_templates/[[:alnum:]]/inputs)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "index",
				Description: `Retrieves the full list of existing inputs of the specified resource.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs_2_0"},
					},
				},
			},

			&ActionCmd{
				Name:        "multi_update",
				Description: `Performs a bulk update of inputs on the specified resource.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "inputs[][value]",
						Description: `The value to be updated with. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "inputs",
						Description: ``,
						Type:        "map",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "inputs[][name]",
						Description: `The name of the input to be updated.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"Instance": &ResourceCmd{
		Name:        "Instance",
		Description: `Instances represent an entity that is runnable in the cloud.`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/instances/[[:alnum:]]|/api/clouds/[[:alnum:]]/instances|/api/clouds/[[:alnum:]]/instances|/api/server_arrays/[[:alnum:]]/current_instances)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Creates and launches a raw instance using the provided parameters.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "instance[cloud_specific_attributes][automatic_instance_store_mapping]",
						Description: `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "instance[cloud_specific_attributes][root_volume_performance]",
						Description: `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[cloud_specific_attributes][root_volume_type_uid]",
						Description: `The type of root volume for instance. Only available on clouds supporting root volume type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[cloud_specific_attributes][iam_instance_profile]",
						Description: `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "instance[cloud_specific_attributes][root_volume_size]",
						Description: `The size for root disk. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[cloud_specific_attributes][ebs_optimized]",
						Description: `Whether the instance is able to connect to IOPS-enabled volumes.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "instance[associate_public_ip_address]",
						Description: `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "instance[placement_group_href]",
						Description: `The placement group to launch the instance in. Not supported by all clouds & instance types.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[security_group_hrefs][]",
						Description: `The hrefs of the security groups.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[instance_type_href]",
						Description: `The href of the instance type.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[ramdisk_image_href]",
						Description: `The href of the ramdisk image.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[kernel_image_href]",
						Description: `The href of the kernel image.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[datacenter_href]",
						Description: `The href of the Datacenter / Zone.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[deployment_href]",
						Description: `The href of the deployment to which the Instance will be added.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[subnet_hrefs][]",
						Description: `The hrefs of the updated subnets.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[ssh_key_href]",
						Description: `The href of the SSH key to use.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[image_href]",
						Description: `The href of the Image to use.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[name]",
						Description: `The name of the instance.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists instances of a given cloud, server array.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended", "full", "full_inputs_2_0", "tiny"},
					},
				},
			},

			&ActionCmd{
				Name:        "launch",
				Description: `Launches an instance using the parameters that this instance has been configured with.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "inputs[][value]",
						Description: `The value of that input. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "inputs",
						Description: ``,
						Type:        "map",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "inputs[][name]",
						Description: `The input name. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "api_behavior",
						Description: `When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"async", "sync"},
					},
				},
			},

			&ActionCmd{
				Name:        "lock",
				Description: ``,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "multi_run_executable",
				Description: `Runs a script or a recipe in the running instances.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "right_script_href",
						Description: `The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "inputs[][value]",
						Description: `The value of these inputs. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "inputs",
						Description: ``,
						Type:        "map",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "inputs[][name]",
						Description: `The name of inputs needed. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "ignore_lock",
						Description: `Specifies the ability to ignore the lock(s) on the Instance(s).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "recipe_name",
						Description: `The name of the recipe to be run.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "multi_terminate",
				Description: `Terminates running instances.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "terminate_all",
						Description: `Specifies the ability to terminate all instances.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "reboot",
				Description: `Reboot a running instance.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "run_executable",
				Description: `Runs a script or a recipe in the running instance.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "right_script_href",
						Description: `The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "inputs[][value]",
						Description: `The value of these inputs. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "inputs",
						Description: ``,
						Type:        "map",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "inputs[][name]",
						Description: `The name of inputs needed. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "ignore_lock",
						Description: `Specifies the ability to ignore the lock on the Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "recipe_name",
						Description: `The name of the recipe to run.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "set_custom_lodgement",
				Description: `This method is deprecated.  Please use InstanceCustomLodgement.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "quantity[][value]",
						Description: `The value of the quantity. Should be a positive integer.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "quantity[][name]",
						Description: `The name of the quantity. A customer-specific string, e.g. "MB/s" or "GB/Month".`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "timeframe",
						Description: `The timeframe (either a month or a single day) for which the quantity value is valid (currently for the PDT timezone only).`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Shows attributes of a single instance.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended", "full", "full_inputs_2_0", "tiny"},
					},
				},
			},

			&ActionCmd{
				Name:        "start",
				Description: `Starts an instance that has been stopped, resuming it to its previously saved volume state.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "stop",
				Description: `Stores the instance's current volume state to resume later using the 'start' action.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "terminate",
				Description: `Terminates a running instance.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "unlock",
				Description: ``,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Updates attributes of a single instance.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "instance[cloud_specific_attributes][automatic_instance_store_mapping]",
						Description: `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "instance[cloud_specific_attributes][root_volume_performance]",
						Description: `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[cloud_specific_attributes][root_volume_type_uid]",
						Description: `The type of root volume for instance. Only available on clouds supporting root volume type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[cloud_specific_attributes][iam_instance_profile]",
						Description: `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "instance[cloud_specific_attributes][root_volume_size]",
						Description: `The size for root disk. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[associate_public_ip_address]",
						Description: `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "instance[multi_cloud_image_href]",
						Description: `The href of the updated MultiCloudImage for the Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[ip_forwarding_enabled]",
						Description: `Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "instance[server_template_href]",
						Description: `The href of the updated ServerTemplate for the Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[security_group_hrefs][]",
						Description: `The hrefs of the updated security groups.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[ramdisk_image_href]",
						Description: `The href of the updated ramdisk image for the Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[instance_type_href]",
						Description: `The href of the updated Instance Type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[kernel_image_href]",
						Description: `The href of the updated kernel image for the Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[deployment_href]",
						Description: `The href of the updated Deployment for the Instance. This is only supported for Instances that are not associated with a Server or ServerArray.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[datacenter_href]",
						Description: `The href of the updated Datacenter / Zone for the Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[subnet_hrefs][]",
						Description: `The hrefs of the updated subnets.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[ssh_key_href]",
						Description: `The href of the updated SSH key for the Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[image_href]",
						Description: `The href of the updated Image for the Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "instance[name]",
						Description: `The updated name to give the Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"InstanceCustomLodgement": &ResourceCmd{
		Name:        "InstanceCustomLodgement",
		Description: `An InstanceCustomLodgement represents a way to create custom reports about a specific instance with a user defined quantity.  Replaces the legacy Instances#setcustomlodgement interface.`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/instances/[[:alnum:]]/instance_custom_lodgements/[[:alnum:]]|/api/clouds/[[:alnum:]]/instances/[[:alnum:]]/instance_custom_lodgements)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Create a lodgement with the quantity and timeframe specified.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "quantity[][value]",
						Description: `The value of the quantity. Should be a positive integer.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "quantity[][name]",
						Description: `The name of the quantity. A customer-specific string, e.g. "MB/s" or "GB/Month".`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "timeframe",
						Description: `The time-frame (either a month "YYYY_MM" or a single day "YYYY_MM_DD") for which the quantity value is valid (currently for the PDT timezone only).`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Destroy the specified lodgement.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `List InstanceCustomLodgements of a given cloud and instance.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show the specified lodgement.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Update a lodgement with the quantity specified.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "quantity[][value]",
						Description: `The value of the quantity. Should be a positive integer.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "quantity[][name]",
						Description: `The name of the quantity. A customer-specific string, e.g. "MB/s" or "GB/Month".`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"InstanceType": &ResourceCmd{
		Name:        "InstanceType",
		Description: ``,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/instance_types/[[:alnum:]]|/api/clouds/[[:alnum:]]/instance_types)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "index",
				Description: `Lists instance types.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Displays information about a single Instance type.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"IpAddress": &ResourceCmd{
		Name:        "IpAddress",
		Description: `An IpAddress provides an abstraction for IPv4 addresses bindable to Instance resources running in a Cloud.`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/ip_addresses/[[:alnum:]]|/api/clouds/[[:alnum:]]/ip_addresses)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Creates a new IpAddress with the given parameters.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "ip_address[deployment_href]",
						Description: `The href of the Deployment that owns this IpAddress.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "ip_address[network_href]",
						Description: `(OpenStack Only) The href of the Network that the IpAddress will be associated to. This parameter is required for OpenStack with Neutron clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "ip_address[domain]",
						Description: `(Amazon Only) Pass vpc to create this IP for EC2-VPC only environments. Pass ec2_classic to create this IP for EC2-Classic environments. Defaults to ec2_classic.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"ec2_classic", "vpc"},
					},
					&ActionFlag{
						Name:        "ip_address[name]",
						Description: `The name of the IpAddress to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes a given IpAddress.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists the IpAddresses available to this account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single IpAddress.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Updates attributes of a given IpAddress.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "ip_address[deployment_href]",
						Description: `The href of the Deployment that owns this IpAddress.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "ip_address[name]",
						Description: `The updated name of the IpAddress.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"IpAddressBinding": &ResourceCmd{
		Name:        "IpAddressBinding",
		Description: `An IpAddressBinding represents an abstraction for binding an IpAddress to an instance.`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/ip_addresses/[[:alnum:]]/ip_address_bindings/[[:alnum:]]|/api/clouds/[[:alnum:]]/ip_address_bindings/[[:alnum:]]|/api/clouds/[[:alnum:]]/ip_addresses/[[:alnum:]]/ip_address_bindings|/api/clouds/[[:alnum:]]/ip_address_bindings)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Creates an ip address binding which attaches a specified IpAddress resource to a specified instance, and also allows for configuration of port forwarding rules...`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "ip_address_binding[public_ip_address_href]",
						Description: `The IpAddress to bind to the specified instance. Required unless port forwarding rule params are passed.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "ip_address_binding[instance_href]",
						Description: `The Instance to which this IpAddress should be bound.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "ip_address_binding[private_port]",
						Description: `Incoming network traffic will get forwarded to this port number on the specified Instance. If not specified, will use public port. Required unless public_ip_address_href is passed.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "ip_address_binding[public_port]",
						Description: `The incoming port for port forwarding. Incoming network traffic on this port will get forwarded (to the IP:Private Port of the specified Instance). Required unless public_ip_address_href is passed.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "ip_address_binding[protocol]",
						Description: `Transport layer protocol of traffic that may be forwarded from public port to private port on the Instance. Required unless public_ip_address_href is passed.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"UDP", "TCP"},
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `No description provided for destroy.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists the ip address bindings available to this account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single ip address binding.`,
				Flags:       []*ActionFlag{},
			},
		},
	},
	"MonitoringMetric": &ResourceCmd{
		Name:        "MonitoringMetric",
		Description: `A monitoring metric is a stream of data that is captured in an instance. Metrics can be monitored, graphed and can be used as the basis for triggering alerts.`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/instances/[[:alnum:]]/monitoring_metrics/[[:alnum:]]|/api/clouds/[[:alnum:]]/instances/[[:alnum:]]/monitoring_metrics)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "data",
				Description: `Gives the raw monitoring data for a particular metric`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "start",
						Description: `An integer number of seconds from current time. e.g. -300`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "end",
						Description: `An integer number of seconds from current time. e.g. -150 or 0 `,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists the monitoring metrics available for the instance and their corresponding graph hrefs.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "period",
						Description: `The time scale for which the graph is generated. Default is 'day'`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"now", "day", "yday", "week", "lweek", "month", "quarter", "year"},
					},
					&ActionFlag{
						Name:        "title",
						Description: `The title of the graph.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "size",
						Description: `The size of the graph to be generated. Default is 'small'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"thumb", "tiny", "small", "large", "xlarge"},
					},
					&ActionFlag{
						Name:        "tz",
						Description: `The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences. `,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Shows attributes of a single monitoring metric.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "period",
						Description: `The time scale for which the graph is generated. Default is 'day'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"now", "day", "yday", "week", "lweek", "month", "quarter", "year"},
					},
					&ActionFlag{
						Name:        "title",
						Description: `The title of the graph.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "size",
						Description: `The size of the graph to be generated. Default is 'small'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"thumb", "tiny", "small", "large", "xlarge"},
					},
					&ActionFlag{
						Name:        "tz",
						Description: `The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences. `,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"MultiCloudImage": &ResourceCmd{
		Name:        "MultiCloudImage",
		Description: `A MultiCloudImage is a RightScale component that functions as a pointer to machine images in specific clouds (e...`,
		HrefRegexp:  regexp.MustCompile(`(/api/server_templates/[[:alnum:]]/multi_cloud_images/[[:alnum:]]|/api/multi_cloud_images/[[:alnum:]]|/api/server_templates/[[:alnum:]]/multi_cloud_images|/api/multi_cloud_images)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "clone",
				Description: `Clones a given MultiCloudImage.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "multi_cloud_image[description]",
						Description: `The description for the cloned MultiCloudImage.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "multi_cloud_image[name]",
						Description: `The name for the cloned MultiCloudImage.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "commit",
				Description: `Commits a given MultiCloudImage. Only HEAD revisions can be committed.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "commit_message",
						Description: `The message associated with the commit.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "create",
				Description: `Creates a new MultiCloudImage with the given parameters.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "multi_cloud_image[description]",
						Description: `The description of the MultiCloudImage to be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "multi_cloud_image[name]",
						Description: `The name of the MultiCloudImage to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes a given MultiCloudImage.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists the MultiCloudImages available to this account. HEAD revisions have a revision of 0.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single MultiCloudImage. HEAD revisions have a revision of 0.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Updates attributes of a given MultiCloudImage. Only HEAD revisions can be updated (revision 0).`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "multi_cloud_image[description]",
						Description: `The updated description for the MultiCloudImage.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "multi_cloud_image[name]",
						Description: `The updated name for the MultiCloudImage.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"MultiCloudImageSetting": &ResourceCmd{
		Name:        "MultiCloudImageSetting",
		Description: `A MultiCloudImageSetting defines which settings should be used when a server is launched in a cloud...`,
		HrefRegexp:  regexp.MustCompile(`(/api/multi_cloud_images/[[:alnum:]]/settings/[[:alnum:]]|/api/multi_cloud_images/[[:alnum:]]/settings)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Creates a new setting for an existing MultiCloudImage.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "multi_cloud_image_setting[instance_type_href]",
						Description: `The href of the instance type. Mandatory if specifying cloud_href.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "multi_cloud_image_setting[ramdisk_image_href]",
						Description: `The href of the ramdisk image. Optional if specifying cloud_href.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "multi_cloud_image_setting[kernel_image_href]",
						Description: `The href of the kernel image. Optional if specifying cloud_href.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "multi_cloud_image_setting[cloud_href]",
						Description: `The href of the Cloud to use.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "multi_cloud_image_setting[image_href]",
						Description: `The href of the Image to use. Mandatory if specifying cloud_href.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "multi_cloud_image_setting[user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes a MultiCloudImage setting.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists the MultiCloudImage settings.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single MultiCloudImage setting.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Updates a settings for a MultiCLoudImage.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "multi_cloud_image_setting[instance_type_href]",
						Description: `The href of the instance type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "multi_cloud_image_setting[ramdisk_image_href]",
						Description: `The href of the ramdisk image.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "multi_cloud_image_setting[kernel_image_href]",
						Description: `The href of the kernel image.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "multi_cloud_image_setting[cloud_href]",
						Description: `The href of the Cloud to use.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "multi_cloud_image_setting[image_href]",
						Description: `The href of the Image to use.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "multi_cloud_image_setting[user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"Network": &ResourceCmd{
		Name:        "Network",
		Description: `A Network is a logical grouping of network devices.`,
		HrefRegexp:  regexp.MustCompile(`(/api/networks/[[:alnum:]]|/api/networks)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Creates a new network.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "network[instance_tenancy]",
						Description: `The launch policy for AWS instances in the Network. Specify 'default' to allow instances to decide their own launch policy. Specify 'dedicated' to force all instances to be launched as 'dedicated'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "network[description]",
						Description: `The description for the Network.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "network[cidr_block]",
						Description: `The range of IP addresses for the Network. This parameter is required for Amazon clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "network[cloud_href]",
						Description: `The Cloud to create the Network in`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "network[name]",
						Description: `The name for the Network.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes the given network(s).`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists networks in this account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Shows attributes of a single network.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Updates the given network.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "network[route_table_href]",
						Description: `Sets the default RouteTable for this Network.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "network[description]",
						Description: `The updated description for the Network.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "network[name]",
						Description: `The updated name for the Network.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"NetworkGateway": &ResourceCmd{
		Name:        "NetworkGateway",
		Description: `A NetworkGateway is an interface that allows traffic to be routed between networks.`,
		HrefRegexp:  regexp.MustCompile(`(/api/network_gateways/[[:alnum:]]|/api/network_gateways)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Create a new NetworkGateway.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "network_gateway[description]",
						Description: `The description to be set on the NetworkGateway.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "network_gateway[cloud_href]",
						Description: `The cloud to create the NetworkGateway in.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "network_gateway[name]",
						Description: `The name to be set on the NetworkGateway.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "network_gateway[type]",
						Description: `The type of the NetworkGateway.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"vpn", "internet"},
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Delete an existing NetworkGateway.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists the NetworkGateways available to this account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single NetworkGateway.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Update an existing NetworkGateway.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "network_gateway[network_href]",
						Description: `Pass a blank string to detach from the specified Network, or pass a valid Network href to attach to the specified network.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "network_gateway[description]",
						Description: `The description to be set on the NetworkGateway.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "network_gateway[name]",
						Description: `The name to be set on the NetworkGateway.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"NetworkOptionGroup": &ResourceCmd{
		Name:        "NetworkOptionGroup",
		Description: `A key/value pair hash containing options for configuring a Network.`,
		HrefRegexp:  regexp.MustCompile(`(/api/network_option_groups/[[:alnum:]]|/api/network_option_groups)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Create a new NetworkOptionGroup.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "network_option_group[description]",
						Description: `Description of this NetworkOptionGroup`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "network_option_group[cloud_href]",
						Description: `The Cloud to create this NetworkOptionGroup in`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "network_option_group[options]",
						Description: ``,
						Type:        "map",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "network_option_group[name]",
						Description: `Name of this NetworkOptionGroup`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "network_option_group[type]",
						Description: `Type of this NetworkOptionGroup`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Delete an existing NetworkOptionGroup.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `List NetworkOptionGroups available in this account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single NetworkOptionGroup.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Update an existing NetworkOptionGroup.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "network_option_group[description]",
						Description: `Update the description`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "network_option_group[options]",
						Description: ``,
						Type:        "map",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "network_option_group[name]",
						Description: `Update the name`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"NetworkOptionGroupAttachment": &ResourceCmd{
		Name:        "NetworkOptionGroupAttachment",
		Description: `Resource for attaching NetworkOptionGroups to Networks.`,
		HrefRegexp:  regexp.MustCompile(`(/api/network_option_group_attachments/[[:alnum:]]|/api/network_option_group_attachments)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Create a new NetworkOptionGroupAttachment.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "network_option_group_attachment[network_option_group_href]",
						Description: `The NetworkOptionGroup to attach to the specified resource.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "network_option_group_attachment[network_href]",
						Description: `The Network to attach the specified NetworkOptionGroup to.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Delete an existing NetworkOptionGroupAttachment.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `List NetworkOptionGroupAttachments in this account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single NetworkOptionGroupAttachment.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Update an existing NetworkOptionGroupAttachment.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "network_option_group_attachment[network_option_group_href]",
						Description: `The NetworkOptionGroup to attach to the specified resource.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"Oauth2": &ResourceCmd{
		Name:        "Oauth2",
		Description: `Note that all API calls irrespective of the resource it is acting on, should pass a header "X-Api-Version" with the value "1...`,
		HrefRegexp:  regexp.MustCompile(`/api/oauth2`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Perform an OAuth 2`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "right_link_version",
						Description: `The RightLink gem version the client conforms to (only needed for instance agent clients).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "client_secret",
						Description: `The client secret (only needed for confidential clients).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "refresh_token",
						Description: `The refresh token obtained from OAuth grant.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "r_s_version",
						Description: `The RightAgent protocol version the client conforms to (only needed for instance agent clients).`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "account_id",
						Description: `The client's account ID (only needed for instance agent clients).`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "grant_type",
						Description: `Type of grant.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"refresh_token"},
					},
					&ActionFlag{
						Name:        "client_id",
						Description: `The client ID (only needed for confidential clients).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"Permission": &ResourceCmd{
		Name:        "Permission",
		Description: ``,
		HrefRegexp:  regexp.MustCompile(`(/api/permissions/[[:alnum:]]|/api/permissions)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Create a permission, thereby granting some user a particular role with respect to the current account...`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "permission[role_title]",
						Description: ``,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "permission[user_href]",
						Description: ``,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Destroy a permission, thereby revoking a user's role with respect to the current account...`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `List all permissions for all users of the current acount.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single permission.`,
				Flags:       []*ActionFlag{},
			},
		},
	},
	"PlacementGroup": &ResourceCmd{
		Name:        "PlacementGroup",
		Description: ``,
		HrefRegexp:  regexp.MustCompile(`(/api/placement_groups/[[:alnum:]]|/api/placement_groups)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Creates a PlacementGroup.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "placement_group[description]",
						Description: `The description of the Placement Group to be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "placement_group[cloud_href]",
						Description: `The Href of the Cloud in which the PlacementGroup should be created. Note: This feature is not supported for all clouds.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "placement_group[name]",
						Description: `The name of the Placement Group to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Destroys a PlacementGroup.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists all PlacementGroups in an account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Shows information about a single PlacementGroup.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"Preference": &ResourceCmd{
		Name:        "Preference",
		Description: `A Preference is a user and account-specific setting. Preferences are used in many part of the RightScale platform and can be used for custom purposes if desired.`,
		HrefRegexp:  regexp.MustCompile(`(/api/preferences/[[:alnum:]]|/api/preferences)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes the given preference.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists all preferences.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Shows a single preference.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "update",
				Description: `If 'id' is known, updates preference with given contents.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "preference[contents]",
						Description: `The updated contents for the Preference.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"Publication": &ResourceCmd{
		Name:        "Publication",
		Description: `A Publication is a revisioned component shared with a set of Account Groups.`,
		HrefRegexp:  regexp.MustCompile(`(/api/publications/[[:alnum:]]|/api/publications)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "import",
				Description: `Imports the given publication and its subordinates to this account.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists the publications available to this account. Only non-HEAD revisions are possible.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single publication. Only non-HEAD revisions are possible.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"PublicationLineage": &ResourceCmd{
		Name:        "PublicationLineage",
		Description: `A Publication Lineage contains lineage information for a Publication in the MultiCloudMarketplace.`,
		HrefRegexp:  regexp.MustCompile(`/api/publication_lineages/[[:alnum:]]`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single publication lineage. Only non-HEAD revisions are possible.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"RecurringVolumeAttachment": &ResourceCmd{
		Name:        "RecurringVolumeAttachment",
		Description: `A RecurringVolumeAttachment specifies a Volume/VolumeSnapshot to attach to a Server/ServerArray the next time an instance is launched.`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/recurring_volume_attachments/[[:alnum:]]|/api/clouds/[[:alnum:]]/volumes/[[:alnum:]]/recurring_volume_attachments/[[:alnum:]]|/api/clouds/[[:alnum:]]/volume_snapshots/[[:alnum:]]/recurring_volume_attachments/[[:alnum:]]|/api/clouds/[[:alnum:]]/recurring_volume_attachments|/api/clouds/[[:alnum:]]/volumes/[[:alnum:]]/recurring_volume_attachments|/api/clouds/[[:alnum:]]/volume_snapshots/[[:alnum:]]/recurring_volume_attachments)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Creates a new recurring volume attachment.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "recurring_volume_attachment[runnable_href]",
						Description: `The href of the server or server array to attach to.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "recurring_volume_attachment[storage_href]",
						Description: `The href of the volume or volume snapshot to be attached on launch of a next instance.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "recurring_volume_attachment[device]",
						Description: `The device location where the volume or volume snapshot will be mounted. Value must be of format /dev/xvd[bcefghij]. This is not reliable and will be deprecated.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes a given recurring volume attachment.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists all recurring volume attachments.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Displays information about a single recurring volume attachment.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"Repository": &ResourceCmd{
		Name:        "Repository",
		Description: `A Repository is a location from which you can download and import design objects such as Chef cookbooks. Using this resource you can add and modify repository information and import assets discovered in the repository.`,
		HrefRegexp:  regexp.MustCompile(`(/api/repositories/[[:alnum:]]|/api/repositories)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "cookbook_import",
				Description: `Performs a Cookbook import, which allows you to use the specified cookbooks in your design objects.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "repository_commit_reference",
						Description: `Optional commit reference indicating last succeeded commit. Must match the Repository's fetch_status.succeeded_commit attribute or the import will not be performed.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "with_dependencies",
						Description: `A flag indicating whether dependencies should automatically be imported.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "asset_hrefs[]",
						Description: `Hrefs of the assets that should be imported.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "namespace",
						Description: `The namespace to import into.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"primary", "alternate"},
					},
					&ActionFlag{
						Name:        "follow",
						Description: `A flag indicating whether imported cookbooks should be followed.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},

			&ActionCmd{
				Name:        "cookbook_import_preview",
				Description: `Retrieves a preview of the effects of a Cookbook import.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "asset_hrefs[]",
						Description: `Hrefs of the assets that should be imported.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "namespace",
						Description: `The namespace to import into.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"primary", "alternate"},
					},
				},
			},

			&ActionCmd{
				Name:        "create",
				Description: `Creates a Repository.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "repository[asset_paths][cookbooks][]",
						Description: `The cookbook paths for the repository`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "repository[credentials][password]",
						Description: `The password, or credential, for the repository (only valid for svn or download repositories).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "repository[credentials][username]",
						Description: `The user name, or credential, for the repository (only valid for svn or download repositories).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "repository[credentials][ssh_key]",
						Description: `The SSH key, or credential, for the repository (only valid for git repositories).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "repository[commit_reference]",
						Description: `The revision for the repository`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "repository[source_type]",
						Description: `The source type for the repository.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"git", "svn", "download"},
					},
					&ActionFlag{
						Name:        "repository[auto_import]",
						Description: `Whether cookbooks should automatically be imported upon repository creation.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "repository[description]",
						Description: `The description for the repository.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "repository[source]",
						Description: `The URL for the repository.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "repository[name]",
						Description: `The repository name.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes the specified Repositories.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists all Repositories for this Account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&ActionCmd{
				Name:        "refetch",
				Description: `Refetches all RepositoryAssets associated with the Repository.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "auto_import",
						Description: `Whether cookbooks should automatically be imported after repositories are fetched.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},

			&ActionCmd{
				Name:        "resolve",
				Description: `Show a list of repositories that have imported cookbooks with the given names.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "imported_cookbook_name[]",
						Description: `A list of cookbook names that were imported by the repository.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Shows a specified Repository.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Updates a specified Repository.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "repository[asset_paths][cookbooks][]",
						Description: `The updated cookbook paths for the repository`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "repository[credentials][password]",
						Description: `The updated password, or credential, for the repository (only valid for svn or download repositories).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "repository[credentials][username]",
						Description: `The updated user name, or credential, for the repository (only valid for svn or download repositories).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "repository[credentials][ssh_key]",
						Description: `The updated SSH key for the repository (only valid for git repositories).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "repository[commit_reference]",
						Description: `The updated commit reference (tag, branch, revision...) for the repository`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "repository[source_type]",
						Description: `The updated source type for the repository.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"git", "svn", "download"},
					},
					&ActionFlag{
						Name:        "repository[description]",
						Description: `The updated description for the repository.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "repository[source]",
						Description: `The updated URL for the repository.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "repository[name]",
						Description: `The updated repository name.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"RepositoryAsset": &ResourceCmd{
		Name:        "RepositoryAsset",
		Description: `A RepositoryAsset represents an item discovered in a Repository`,
		HrefRegexp:  regexp.MustCompile(`(/api/repositories/[[:alnum:]]/repository_assets/[[:alnum:]]|/api/repositories/[[:alnum:]]/repository_assets)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "index",
				Description: `List a repository's current assets.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single asset.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"RightScript": &ResourceCmd{
		Name:        "RightScript",
		Description: `A RightScript is an executable piece of code that can be run on a server during the boot, operational, or decommission phases...`,
		HrefRegexp:  regexp.MustCompile(`(/api/right_scripts/[[:alnum:]]|/api/right_scripts)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "commit",
				Description: `Commits the given RightScript. Only HEAD revisions (revision 0) can be committed.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "right_script[commit_message]",
						Description: `The message to be included with the requested commit`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists RightScripts.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "latest_only",
						Description: `Whether or not to return only the latest version for each lineage.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Displays information about a single RightScript.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "show_source",
				Description: `Returns the script source for a RightScript`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Updates RightScript name/description`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "right_script[description]",
						Description: `The new description for the RightScript`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "right_script[name]",
						Description: `The new name for the RightScript`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "update_source",
				Description: `Updates the source of the given RightScript`,
				Flags:       []*ActionFlag{},
			},
		},
	},
	"Route": &ResourceCmd{
		Name:        "Route",
		Description: `A Route defines how networking traffic should be routed from one destination to another...`,
		HrefRegexp:  regexp.MustCompile(`(/api/routes/[[:alnum:]]|/api/route_tables/[[:alnum:]]/routes/[[:alnum:]]|/api/routes|/api/route_tables/[[:alnum:]]/routes)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Create a new Route.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "route[destination_cidr_block]",
						Description: `The destination (CIDR IP address) for the Route.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "route[route_table_href]",
						Description: `The RouteTable to create the Route in.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "route[next_hop_href]",
						Description: `The href of the Route's next hop.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "route[next_hop_type]",
						Description: `The Route's next hop type.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"instance", "network_interface", "network_gateway", "ip_string"},
					},
					&ActionFlag{
						Name:        "route[description]",
						Description: `The description to be set on the Route.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "route[next_hop_ip]",
						Description: `The IP Address of the Route's next hop. Required if route[next_hop_type] is 'ip_string'. Not allowed otherwise.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Delete an existing Route.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `List Routes available in this account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single Route.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Update an existing Route.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "route[destination_cidr_block]",
						Description: `The updated destination (CIDR IP address) for the Route.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "route[next_hop_href]",
						Description: `The updated href of the Route's next hop. Required if route[next_hop_type] is 'instance', 'network_interface', or 'network_gateway'. Not allowed otherwise.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "route[next_hop_type]",
						Description: `The updated Route's next hop type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"instance", "network_interface", "network_gateway", "ip_string"},
					},
					&ActionFlag{
						Name:        "route[description]",
						Description: `The updated description of the Route.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "route[next_hop_ip]",
						Description: `The updated IP Address of the Route's next hop. Required if route[next_hop_type] is 'ip_string'. Not allowed otherwise.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"RouteTable": &ResourceCmd{
		Name:        "RouteTable",
		Description: `Grouped listing of Routes`,
		HrefRegexp:  regexp.MustCompile(`(/api/route_tables/[[:alnum:]]|/api/route_tables)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Create a new RouteTable.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "route_table[network_href]",
						Description: `The Network to create the RouteTable in.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "route_table[description]",
						Description: `The description to be set on the RouteTable.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "route_table[cloud_href]",
						Description: `The cloud to create the RouteTable in.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "route_table[name]",
						Description: `The name to be set on the RouteTable.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Delete an existing RouteTable.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `List RouteTables available in this account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single RouteTable.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Update an existing RouteTable.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "route_table[description]",
						Description: `The description to be set on the RouteTable.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "route_table[name]",
						Description: `The name to be set on the RouteTable.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"RunnableBinding": &ResourceCmd{
		Name:        "RunnableBinding",
		Description: `A RunnableBinding represents an item in a runlist of a ServerTemplate`,
		HrefRegexp:  regexp.MustCompile(`(/api/server_templates/[[:alnum:]]/runnable_bindings/[[:alnum:]]|/api/server_templates/[[:alnum:]]/runnable_bindings)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Bind an executable to the given ServerTemplate.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "runnable_binding[right_script_href]",
						Description: `The RightScript href. Note: recipe cannot be specified when this param is given.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "runnable_binding[position]",
						Description: `The position of the executable in the execution order. If not specified, will be added to the end. If specified, will be inserted in that location and cause all others to move down.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "runnable_binding[sequence]",
						Description: `The sequence at which this executable should be run. Default is 'operational'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"boot", "decommission", "operational"},
					},
					&ActionFlag{
						Name:        "runnable_binding[recipe]",
						Description: `The Chef recipe name. Note: right_script_href cannot be specified when this param is given.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Unbind an executable from the given resource.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists the executables bound to the ServerTemplate.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "multi_update",
				Description: `Update attributes for multiple bound executables.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "runnable_bindings[][right_script_href]",
						Description: `The updated RightScript href. Note: recipe cannot be specified when this param is given.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "runnable_bindings[][position]",
						Description: `The updated position of the RunnableBinding in the execution order. If specified, will be inserted in that location and cause all others to move down.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "runnable_bindings[][sequence]",
						Description: `The sequence at which this executable should be run.  Default is 'operational'.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"boot", "decommission", "operational"},
					},
					&ActionFlag{
						Name:        "runnable_bindings[][recipe]",
						Description: `The updated Chef recipe name. Note: right_script_href cannot be specified when this param is given.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "runnable_bindings[][id]",
						Description: `The ID of the RunnableBinding to update.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single executable binding.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"SecurityGroup": &ResourceCmd{
		Name:        "SecurityGroup",
		Description: `Security Groups represent network security profiles that contain lists of firewall rules for different ports and source IP addresses, as well as trust relationships amongst different security groups...`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/security_groups/[[:alnum:]]|/api/clouds/[[:alnum:]]/security_groups)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Create a security group.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "security_group[network_href]",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "security_group[description]",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "security_group[name]",
						Description: ``,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Delete security group(s)`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists Security Groups.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "tiny"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Displays information about a single Security Group.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "tiny"},
					},
				},
			},
		},
	},
	"SecurityGroupRule": &ResourceCmd{
		Name:        "SecurityGroupRule",
		Description: ``,
		HrefRegexp:  regexp.MustCompile(`(/api/security_group_rules/[[:alnum:]]|/api/clouds/[[:alnum:]]/security_groups/[[:alnum:]]/security_group_rules/[[:alnum:]]|/api/security_group_rules|/api/clouds/[[:alnum:]]/security_groups/[[:alnum:]]/security_group_rules)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Create a security group rule for a security group.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "security_group_rule[protocol_details][start_port]",
						Description: `Start of port range (inclusive). Required if protocol is 'tcp' or 'udp'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "security_group_rule[protocol_details][icmp_type]",
						Description: `ICMP type. Required if protocol is 'icmp'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "security_group_rule[protocol_details][icmp_code]",
						Description: `ICMP code. Required if protocol is 'icmp'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "security_group_rule[protocol_details][end_port]",
						Description: `End of port range (inclusive). Required if protocol is 'tcp' or 'udp'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "security_group_rule[security_group_href]",
						Description: `Security Group to add rule to.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "security_group_rule[source_type]",
						Description: `Source type. May be a CIDR block or another Security Group.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"cidr_ips", "group"},
					},
					&ActionFlag{
						Name:        "security_group_rule[group_owner]",
						Description: `Owner of source Security Group. Required if source_type is 'group'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "security_group_rule[group_name]",
						Description: `Name of source Security Group. Required if source_type is 'group'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "security_group_rule[direction]",
						Description: `Direction of traffic.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"ingress", "egress"},
					},
					&ActionFlag{
						Name:        "security_group_rule[protocol]",
						Description: `Protocol to filter on.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"tcp", "udp", "icmp", "all"},
					},
					&ActionFlag{
						Name:        "security_group_rule[cidr_ips]",
						Description: `An IP address range in CIDR notation. Required if source_type is 'cidr_ips'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Delete security group rule(s)`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists SecurityGroupRules.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Displays information about a single SecurityGroupRule.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "update",
				Description: ``,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "security_group_rule[description]",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"Server": &ResourceCmd{
		Name:        "Server",
		Description: `Servers represent the notion of a server/machine from the RightScale's perspective`,
		HrefRegexp:  regexp.MustCompile(`(/api/servers/[[:alnum:]]|/api/deployments/[[:alnum:]]/servers/[[:alnum:]]|/api/servers|/api/deployments/[[:alnum:]]/servers)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "clone",
				Description: `Clones a given server.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "create",
				Description: `Creates a new server, and configures its corresponding "next" instance with the received parameters.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "server[instance][cloud_specific_attributes][automatic_instance_store_mapping]",
						Description: `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "server[instance][cloud_specific_attributes][root_volume_performance]",
						Description: `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][cloud_specific_attributes][root_volume_type_uid]",
						Description: `The type of root volume for instance. Only available on clouds supporting root volume type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][cloud_specific_attributes][iam_instance_profile]",
						Description: `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "server[instance][cloud_specific_attributes][root_volume_size]",
						Description: `The size for root disk. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][associate_public_ip_address]",
						Description: `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "server[instance][multi_cloud_image_href]",
						Description: `The href of the Multi Cloud Image to use.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][ip_forwarding_enabled]",
						Description: `Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "server[instance][placement_group_href]",
						Description: `The href of the Placement Group.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][security_group_hrefs][]",
						Description: `The hrefs of the security groups.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][server_template_href]",
						Description: `The href of the Server Template.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][instance_type_href]",
						Description: `The href of the Instance Type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][ramdisk_image_href]",
						Description: `The href of the Ramdisk Image.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][kernel_image_href]",
						Description: `The href of the Kernel Image.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][datacenter_href]",
						Description: `The href of the Datacenter / Zone.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][inputs][][value]",
						Description: `The value of that Input. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][inputs]",
						Description: ``,
						Type:        "map",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][inputs][][name]",
						Description: `The Input name. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][ssh_key_href]",
						Description: `The href of the SSH key to use.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][subnet_hrefs][]",
						Description: `The hrefs of the updated subnets.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][cloud_href]",
						Description: `The href of the cloud that the Server should be added to.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][image_href]",
						Description: `The href of the Image to use.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[deployment_href]",
						Description: `The href of the deployment to which the Server will be added.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[description]",
						Description: `The Server description.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[optimized]",
						Description: `A flag indicating whether Instances of this Server should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "server[name]",
						Description: `The name of the Server.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes a given server.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists servers.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "instance_detail"},
					},
				},
			},

			&ActionCmd{
				Name:        "launch",
				Description: `Launches the "next" instance of this server`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Shows the information of a single server.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "instance_detail"},
					},
				},
			},

			&ActionCmd{
				Name:        "terminate",
				Description: `Terminates the current instance of this server`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Updates attributes of a single server.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "server[automatic_instance_store_mapping]",
						Description: `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "server[root_volume_size]",
						Description: `The size for root disk. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[description]",
						Description: `The updated description for the server.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[optimized]",
						Description: `A flag indicating whether Instances of this Server should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "server[name]",
						Description: `The updated server name.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "wrap_instance",
				Description: `Wrap an existing instance and set current instance for new server`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "server[instance][multi_cloud_image_href]",
						Description: `The href of the Multi Cloud Image to use.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][server_template_href]",
						Description: `The href of the Server Template.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][inputs]",
						Description: ``,
						Type:        "map",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[deployment_href]",
						Description: `The href of the deployment to which the Server will be added.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[instance][href]",
						Description: `The href of the Instance around which the server should be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[description]",
						Description: `The Server description.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server[name]",
						Description: `The name of the Server.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"ServerArray": &ResourceCmd{
		Name:        "ServerArray",
		Description: `A server array represents a logical group of instances and allows to resize(grow/shrink) that group based on certain elasticity parameters.`,
		HrefRegexp:  regexp.MustCompile(`(/api/server_arrays/[[:alnum:]]|/api/deployments/[[:alnum:]]/server_arrays/[[:alnum:]]|/api/server_arrays|/api/deployments/[[:alnum:]]/server_arrays)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "clone",
				Description: `Clones a given server array.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "create",
				Description: `Creates a new server array, and configures its corresponding "next" instance with the received parameters.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "server_array[elasticity_params][queue_specific_params][queue_size][items_per_instance]",
						Description: `Defines the ratio of worker instances per items in the queue. Example: If there are 50 items in the queue and "Items per instance" is set to 10, the server array will resize to 5 worker instances (50/10).  Default = 10`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][cloud_specific_attributes][automatic_instance_store_mapping]",
						Description: `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][queue_specific_params][collect_audit_entries]",
						Description: `The audit SQS queue that will store audit entries.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][alert_specific_params][voters_tag_predicate]",
						Description: `The Voters Tag that RightScale will use in order to determine when to scale up/down.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][queue_specific_params][item_age][algorithm]",
						Description: `The algorithm that defines how an item's age will be determined, either by the average age or max (oldest) age.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"max_10", "avg_10"},
					},
					&ActionFlag{
						Name:        "server_array[instance][cloud_specific_attributes][root_volume_performance]",
						Description: `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][alert_specific_params][decision_threshold]",
						Description: `The percentage of servers that must agree in order to trigger an alert before an action is taken.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][queue_specific_params][item_age][max_age]",
						Description: `The threshold (in seconds) before a resize action occurs on the server array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][queue_specific_params][item_age][regexp]",
						Description: `The regexp that helps the system determine an item's "age" in the queue. Example: created_at: (\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d UTC)`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][cloud_specific_attributes][root_volume_type_uid]",
						Description: `The type of root volume for instance. Only available on clouds supporting root volume type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][cloud_specific_attributes][iam_instance_profile]",
						Description: `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "server_array[instance][cloud_specific_attributes][root_volume_size]",
						Description: `The size for root disk. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][pacing][resize_calm_time]",
						Description: `The time (in minutes) on how long you want to wait before you repeat another action.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][pacing][resize_down_by]",
						Description: `The number of servers to scale down by.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][schedule][][max_count]",
						Description: `The maximum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][schedule][][min_count]",
						Description: `The minimum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][pacing][resize_up_by]",
						Description: `The number of servers to scale up by.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][associate_public_ip_address]",
						Description: `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][bounds][max_count]",
						Description: `The maximum number of servers that can be operational at the same time in the server array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][bounds][min_count]",
						Description: `The minimum number of servers that must be operational at all times in the server array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[datacenter_policy][][datacenter_href]",
						Description: `The href of the Datacenter / Zone.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][schedule][][time]",
						Description: `Specifies the time when an alert-based array resizes.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][schedule][][day]",
						Description: `Specifies the day when an alert-based array resizes.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
					},
					&ActionFlag{
						Name:        "server_array[instance][multi_cloud_image_href]",
						Description: `The href of the MultiCloudImage to be used.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][ip_forwarding_enabled]",
						Description: `Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "server_array[instance][security_group_hrefs][]",
						Description: `The hrefs of the Security Groups.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][server_template_href]",
						Description: `The ServerTemplate that will be used to create the worker instances in the server array.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][placement_group_href]",
						Description: `The href of the Placement Group.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][ramdisk_image_href]",
						Description: `The href of the Ramdisk Image.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][instance_type_href]",
						Description: `The href of the Instance Type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[datacenter_policy][][weight]",
						Description: `Instance allocation (should total 100%).`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][kernel_image_href]",
						Description: `The href of the Kernel Image.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][inputs][][value]",
						Description: `The value of that Input. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][inputs]",
						Description: ``,
						Type:        "map",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][datacenter_href]",
						Description: `The href of the Datacenter / Zone. For multiple Datacenters, use 'datacenter_policy' instead.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[datacenter_policy][][max]",
						Description: `Max instances (0 for unlimited).`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][inputs][][name]",
						Description: `The Input name.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][ssh_key_href]",
						Description: `The href of the SSH Key to be used.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][subnet_hrefs][]",
						Description: `The hrefs of the updated Subnets.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][cloud_href]",
						Description: `The href of the Cloud that the array will be associated with.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][image_href]",
						Description: `The href of the Image to be used.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[instance][user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[deployment_href]",
						Description: `The href of the deployment for the Server Array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[description]",
						Description: `The description for the Server Array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[array_type]",
						Description: `The array type for the Server Array.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"alert", "queue"},
					},
					&ActionFlag{
						Name:        "server_array[optimized]",
						Description: `A flag indicating whether Instances of this ServerArray should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "server_array[state]",
						Description: `The status of the server array. If active, the server array is enabled for scaling actions.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"enabled", "disabled"},
					},
					&ActionFlag{
						Name:        "server_array[name]",
						Description: `The name for the Server Array.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "current_instances",
				Description: `List the running instances belonging to the server array. See Instances#index for details.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes a given server array.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists server arrays.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "instance_detail"},
					},
				},
			},

			&ActionCmd{
				Name:        "launch",
				Description: `Launches a new instance in the server array with the configuration defined in the 'next_instance'`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "multi_run_executable",
				Description: `Run an executable on all instances of this array`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "multi_terminate",
				Description: `Terminate all instances of this array`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Shows the information of a single server array.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "instance_detail"},
					},
				},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Updates attributes of a single server array.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "server_array[elasticity_params][queue_specific_params][queue_size][items_per_instance]",
						Description: `Defines the ratio of worker instances per items in the queue. Example: If there are 50 items in the queue and "Items per instance" is set to 10, the server array will resize to 5 worker instances (50/10).  Default = 10`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][queue_specific_params][collect_audit_entries]",
						Description: `The updated audit SQS queue that will store audit entries.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][alert_specific_params][voters_tag_predicate]",
						Description: `The updated Voters Tag that RightScale will use in order to determine when to scale up/down.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][queue_specific_params][item_age][algorithm]",
						Description: `The updated algorithm that defines how an item's age will be determined, either by the average age or max (oldest) age.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"max_10", "avg_10"},
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][alert_specific_params][decision_threshold]",
						Description: `The updated percentage of servers that must agree in order to trigger an alert before an action is taken.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][queue_specific_params][item_age][max_age]",
						Description: `The updated threshold (in seconds) before a resize action occurs on the server array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][queue_specific_params][item_age][regexp]",
						Description: `The updated regexp that helps the system determine an item's "age" in the queue. Example: created_at: (\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d UTC)`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][pacing][resize_calm_time]",
						Description: `The updated time (in minutes) on how long you want to wait before you repeat another action.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][pacing][resize_down_by]",
						Description: `The updated number of servers to scale down by.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][schedule][][min_count]",
						Description: `The updated minimum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][schedule][][max_count]",
						Description: `The updated maximum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][pacing][resize_up_by]",
						Description: `The updated number of servers to scale up by.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][bounds][min_count]",
						Description: `The updated minimum number of servers that must be operational at all times in the server array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][bounds][max_count]",
						Description: `The updated maximum number of servers that can be operational at the same time in the server array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[datacenter_policy][][datacenter_href]",
						Description: `The href of the Datacenter / Zone.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][schedule][][time]",
						Description: `The updated time when an alert-based array resizes.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[elasticity_params][schedule][][day]",
						Description: `The updated day when an alert-based array resizes.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
					},
					&ActionFlag{
						Name:        "server_array[datacenter_policy][][weight]",
						Description: `Instance allocation (should total 100%).`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[datacenter_policy][][max]",
						Description: `Max instances (0 for unlimited).`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[deployment_href]",
						Description: `The updated href of the deployment for the Server Array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[description]",
						Description: `The updated description for the Server Array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_array[array_type]",
						Description: `The updated array type for the Server Array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"alert", "queue"},
					},
					&ActionFlag{
						Name:        "server_array[optimized]",
						Description: `A flag indicating whether Instances of this ServerArray should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "server_array[state]",
						Description: `The updated status of the server array. If active, the server array is enabled for scaling actions.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"enabled", "disabled"},
					},
					&ActionFlag{
						Name:        "server_array[name]",
						Description: `The updated name for the Server Array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"ServerTemplate": &ResourceCmd{
		Name:        "ServerTemplate",
		Description: `ServerTemplates allow you to pre-configure servers by starting from a base image and adding scripts that run during the boot, operational, and shutdown phases...`,
		HrefRegexp:  regexp.MustCompile(`(/api/server_templates/[[:alnum:]]|/api/server_templates)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "clone",
				Description: `Clones a given ServerTemplate.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "server_template[description]",
						Description: `The description for the cloned ServerTemplate.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_template[name]",
						Description: `The name for the cloned ServerTemplate.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "commit",
				Description: `Commits a given ServerTemplate. Only HEAD revisions (revision 0) that are owned by the account can be committed.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "commit_head_dependencies",
						Description: `Commit all HEAD revisions (if any) of the associated MultiCloud Images, RightScripts and Chef repo sequences.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "freeze_repositories",
						Description: `Freeze the repositories.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "commit_message",
						Description: `The message associated with the commit.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "create",
				Description: `Creates a new ServerTemplate with the given parameters.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "server_template[description]",
						Description: `The description of the ServerTemplate to be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "server_template[name]",
						Description: `The name of the ServerTemplate to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes a given ServerTemplate.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "detect_changes_in_head",
				Description: `Identifies RightScripts attached to the resource that differ from their HEAD.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists the ServerTemplates available to this account. HEAD revisions have a revision of 0.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs", "inputs_2_0"},
					},
				},
			},

			&ActionCmd{
				Name:        "publish",
				Description: `Publishes a given ServerTemplate and its subordinates.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "account_group_hrefs[]",
						Description: `List of hrefs of account groups to publish to.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "descriptions[short]",
						Description: `Short Description.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "descriptions[notes]",
						Description: `New Revision Notes.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "descriptions[long]",
						Description: `Long Description.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "allow_comments",
						Description: `Allow users to leave comments on this ServerTemplate.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "email_comments",
						Description: `Email me when a user comments on this ServerTemplate.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "categories[]",
						Description: `List of Categories.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "resolve",
				Description: `Enumerates all attached cookbooks, missing dependencies and bound executables.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single ServerTemplate. HEAD revisions have a revision of 0.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs", "inputs_2_0"},
					},
				},
			},

			&ActionCmd{
				Name:        "swap_repository",
				Description: `In-place replacement of attached cookbooks from a given repository.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "source_repository_href",
						Description: `The repository whose cookbook attachments are to be replaced.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "target_repository_href",
						Description: `The repository whose cookbook attachments are to be utilized.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Updates attributes of a given ServerTemplate. Only HEAD revisions can be updated (revision 0).`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "server_template[description]",
						Description: `The updated description for the ServerTemplate.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "server_template[name]",
						Description: `The updated name for the ServerTemplate.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"ServerTemplateMultiCloudImage": &ResourceCmd{
		Name:        "ServerTemplateMultiCloudImage",
		Description: `This resource represents links between ServerTemplates and MultiCloud Images and enables you to effectively add/delete MultiCloudImages to ServerTemplates and make them the default one...`,
		HrefRegexp:  regexp.MustCompile(`(/api/server_template_multi_cloud_images/[[:alnum:]]|/api/server_template_multi_cloud_images)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Creates a new ServerTemplateMultiCloudImage with the given parameters.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "server_template_multi_cloud_image[multi_cloud_image_href]",
						Description: `The href of the MultiCloud Image to be used.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "server_template_multi_cloud_image[server_template_href]",
						Description: `The href of the ServerTemplate to be used.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes a given ServerTemplateMultiCloudImage.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists the ServerTemplateMultiCloudImages available to this account.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "make_default",
				Description: `Makes a given ServerTemplateMultiCloudImage the default for the ServerTemplate.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single ServerTemplateMultiCloudImage which represents an association between a ServerTemplate and a MultiCloudImage.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"Session": &ResourceCmd{
		Name:        "Session",
		Description: `The sessions resource is in charge of creating API sessions that are bound to a given account`,
		HrefRegexp:  regexp.MustCompile(`/api/sessions`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "accounts",
				Description: `List all the accounts that a user has access to.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "password",
						Description: `The corresponding password.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "email",
						Description: `The email to login with if not using existing session.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: `Extended view shows account permissions and products`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Returns a list of root resources so an authenticated session can use them as a starting point or a way to know what features are available within its privileges...`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index_instance_session",
				Description: `Shows the full attributes of the instance (that has the token used to log-in).`,
				Flags:       []*ActionFlag{},
			},
		},
	},
	"SshKey": &ResourceCmd{
		Name:        "SshKey",
		Description: `Ssh Keys represent a created SSH Key that exists in the cloud.`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/ssh_keys/[[:alnum:]]|/api/clouds/[[:alnum:]]/ssh_keys)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Creates a new ssh key.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "ssh_key[name]",
						Description: `The name for the SSH key to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes a given ssh key.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists ssh keys.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "sensitive"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Displays information about a single ssh key.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "sensitive"},
					},
				},
			},
		},
	},
	"Subnet": &ResourceCmd{
		Name:        "Subnet",
		Description: `A Subnet is a logical grouping of network devices`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/instances/[[:alnum:]]/subnets/[[:alnum:]]|/api/clouds/[[:alnum:]]/subnets/[[:alnum:]]|/api/clouds/[[:alnum:]]/instances/[[:alnum:]]/subnets|/api/clouds/[[:alnum:]]/subnets)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Creates a new subnet.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "subnet[datacenter_href]",
						Description: `The associated Datacenter.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "subnet[network_href]",
						Description: `The associated Network.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "subnet[description]",
						Description: `The description for the Subnet.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "subnet[cidr_block]",
						Description: `The range of IP addresses for the Subnet.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "subnet[name]",
						Description: `The name for the Subnet.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes the given subnet(s).`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists subnets of a given cloud.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Shows attributes of a single subnet.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Updates name and description for the current subnet.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "subnet[route_table_href]",
						Description: `The RouteTable to associate/dissociate with this Subnet. Pass this parameter with an empty string to reset this Subnet to use the default RouteTable.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "subnet[description]",
						Description: `The updated description for the Subnet.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&ActionFlag{
						Name:        "subnet[name]",
						Description: `The updated name for the Subnet.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"Tag": &ResourceCmd{
		Name:        "Tag",
		Description: `A tag or machine tag is a useful way of attaching useful metadata to an object/resource.`,
		HrefRegexp:  regexp.MustCompile(`/api/tags`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "by_resource",
				Description: `Get tags for a list of resource hrefs.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "resource_hrefs[]",
						Description: `Hrefs of the resources for which tags are to be returned.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "by_tag",
				Description: `Search for resources having a list of tags in a specific resource_type.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "include_tags_with_prefix",
						Description: `If included, all tags matching this prefix will be returned. If not included, no tags will be returned.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "resource_type",
						Description: `Search among a single resource type.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"servers", "instances", "volumes", "volume_snapshots", "deployments", "server_templates", "multi_cloud_images", "images", "server_arrays", "accounts"},
					},
					&ActionFlag{
						Name:        "with_deleted",
						Description: `If set to 'true', tags for deleted resources will also be returned. Default value is 'false'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "match_all",
						Description: `If set to 'true', resources having all the tags specified in the 'tags' parameter are returned. Otherwise, resources having any of the tags are returned.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "tags[]",
						Description: `The tags which must be present on the resource.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "multi_add",
				Description: `Add a list of tags to a list of hrefs. The tags must be either plain_tags or machine_tags.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "resource_hrefs[]",
						Description: `Hrefs of the resources for which the tags are to be added.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "tags[]",
						Description: `Tags to be added.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "multi_delete",
				Description: `Delete a list of tags on a list of hrefs. The tags must be either plain_tags or machine_tags.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "resource_hrefs[]",
						Description: `Hrefs of the resources for which tags are to be deleted.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "tags[]",
						Description: `Tags to be deleted.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"Task": &ResourceCmd{
		Name:        "Task",
		Description: `Tasks represent processes that happen (or have happened) asynchronously within the context of an Instance.`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/instances/[[:alnum:]]/live/tasks/[[:alnum:]]|/api/server_arrays/[[:alnum:]]/live/tasks/[[:alnum:]])`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "show",
				Description: `Displays information of a given task within the context of an instance.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},
		},
	},
	"User": &ResourceCmd{
		Name:        "User",
		Description: `A User represents an individual RightScale user`,
		HrefRegexp:  regexp.MustCompile(`(/api/users/[[:alnum:]]|/api/users)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Create a user. If a user already exists with the same email, that user will be returned.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "user[identity_provider_href]",
						Description: `The RightScale API href of the Identity Provider through which this user will login to RightScale. Required to create an SSO-authenticated user.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "user[principal_uid]",
						Description: `The principal identifier (SAML NameID or OpenID identity URL) of this user. Required to create an SSO-authenticated user.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "user[timezone_name]",
						Description: `This can be in the form of country/region or timezone name. For example 'America/Los_Angeles' or 'GB' or 'UTC'. A complete list of acceptable values is available in the Settings > User Settings > Preferences page.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "user[first_name]",
						Description: ``,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "user[last_name]",
						Description: ``,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "user[password]",
						Description: `The password of this user. Required to create a password-authenticated user.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "user[company]",
						Description: ``,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "user[email]",
						Description: ``,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "user[phone]",
						Description: ``,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "index",
				Description: `List the users available to the account the user is logged in to`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Show information about a single user.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "update",
				Description: `Update a user's contact information, change her password, or update SSO her settings`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "user[identity_provider_href]",
						Description: `The updated RightScale API href of the associated Identity Provider.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "user[current_password]",
						Description: `The current password for the user.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "user[principal_uid]",
						Description: `The updated principal identifier (SAML NameID or OpenID identity URL) of this user.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "user[current_email]",
						Description: `The existing email of this user.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "user[timezone_name]",
						Description: `This can be in the form of country/region or timezone name. For example 'America/Los_Angeles' or 'GB' or 'UTC'. A complete list of acceptable values is available in the Settings > User Settings > Preferences page.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "user[new_password]",
						Description: `The new password for this user.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "user[first_name]",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "user[last_name]",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "user[new_email]",
						Description: `The updated email of this user.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "user[company]",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "user[phone]",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"UserData": &ResourceCmd{
		Name:        "UserData",
		Description: ``,
		HrefRegexp:  regexp.MustCompile(`/api/user_data/`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "show",
				Description: `No description provided for show.`,
				Flags:       []*ActionFlag{},
			},
		},
	},
	"Volume": &ResourceCmd{
		Name:        "Volume",
		Description: `A Volume provides a highly reliable, efficient and persistent storage solution that can be mounted to a cloud instance (in the same datacenter / zone).`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/volumes/[[:alnum:]]|/api/clouds/[[:alnum:]]/volumes)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Creates a new volume.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "volume[parent_volume_snapshot_href]",
						Description: `The href of the snapshot from which the volume will be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "volume[placement_group_href]",
						Description: `The href of the Placement Group.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "volume[volume_type_href]",
						Description: `The href of the volume type. A Name, Resource UID and optional Size is associated with a Volume Type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "volume[deployment_href]",
						Description: `The href of the Deployment that owns this Volume.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "volume[datacenter_href]",
						Description: `The href of the Datacenter / Zone that the Volume will be in. This parameter is required for non-OpenStack clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "volume[description]",
						Description: `The description of the Volume to be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "volume[encrypted]",
						Description: `A flag indicating whether Volume should be encrypted. Only available on clouds supporting volume encryption.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&ActionFlag{
						Name:        "volume[iops]",
						Description: `The number of IOPS (I/O Operations Per Second) this Volume should support. Only available on clouds supporting performance provisioning.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "volume[name]",
						Description: `The name for the Volume to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "volume[size]",
						Description: `The size of a Volume to be created in gigabytes (GB). Some Volume Types have predefined sizes and do not allow selecting a custom size on Volume creation.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes a given volume.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists volumes.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Displays information about a single volume.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},
		},
	},
	"VolumeAttachment": &ResourceCmd{
		Name:        "VolumeAttachment",
		Description: `A VolumeAttachment represents a relationship between a volume and an instance.`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/instances/[[:alnum:]]/volume_attachments/[[:alnum:]]|/api/clouds/[[:alnum:]]/volume_attachments/[[:alnum:]]|/api/clouds/[[:alnum:]]/volumes/[[:alnum:]]/volume_attachments|/api/clouds/[[:alnum:]]/volumes/[[:alnum:]]/volume_attachment|/api/clouds/[[:alnum:]]/instances/[[:alnum:]]/volume_attachments|/api/clouds/[[:alnum:]]/volume_attachments)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Creates a new volume attachment.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "volume_attachment[instance_href]",
						Description: `The href of the instance to which the volume will be attached to.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "volume_attachment[volume_href]",
						Description: `The href of the volume to be attached.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "volume_attachment[device]",
						Description: `The device location where the volume will be mounted. Value must be of format /dev/xvd[bcefghij]. This is not reliable and will be deprecated.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes a given volume attachment.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "force",
						Description: `Specifies whether to force the detachment of a volume.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists all volume attachments.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Displays information about a single volume attachment.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"VolumeSnapshot": &ResourceCmd{
		Name:        "VolumeSnapshot",
		Description: `A VolumeSnapshot represents a Cloud storage volume at a particular point in time`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/volumes/[[:alnum:]]/volume_snapshots/[[:alnum:]]|/api/clouds/[[:alnum:]]/volume_snapshots/[[:alnum:]]|/api/clouds/[[:alnum:]]/volumes/[[:alnum:]]/volume_snapshots|/api/clouds/[[:alnum:]]/volume_snapshots)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "create",
				Description: `Creates a new volume_snapshot.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "volume_snapshot[parent_volume_href]",
						Description: `The href of the Volume from which the Volume Snapshot will be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "volume_snapshot[deployment_href]",
						Description: `The href of the Deployment that owns this Volume Snapshot.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "volume_snapshot[description]",
						Description: `The description for the Volume Snapshot to be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "volume_snapshot[name]",
						Description: `The name for the Volume Snapshot to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&ActionCmd{
				Name:        "destroy",
				Description: `Deletes a given volume_snapshot.`,
				Flags:       []*ActionFlag{},
			},

			&ActionCmd{
				Name:        "index",
				Description: `Lists all volume_snapshots.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Displays information about a single volume_snapshot.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"VolumeType": &ResourceCmd{
		Name:        "VolumeType",
		Description: `A VolumeType describes the type of volume, particularly the size.`,
		HrefRegexp:  regexp.MustCompile(`(/api/clouds/[[:alnum:]]/volume_types/[[:alnum:]]|/api/clouds/[[:alnum:]]/volume_types)`),
		Actions: []*ActionCmd{
			&ActionCmd{
				Name:        "index",
				Description: `Lists Volume Types.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&ActionCmd{
				Name:        "show",
				Description: `Displays information about a single Volume Type.`,
				Flags: []*ActionFlag{
					&ActionFlag{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
}

// Action info struct
type ActionInfo struct {
	Name        string // Action name
	Description string // Action description
}

// Map resource names to action info
var resourceActions = map[string][]ActionInfo{
	"Account": []ActionInfo{
		ActionInfo{"show", `Show information about a single Account.`},
	},
	"AccountGroup": []ActionInfo{
		ActionInfo{"index", `Lists the AccountGroups owned by this Account.`},
		ActionInfo{"show", `Show information about a single AccountGroup.`},
	},
	"Alert": []ActionInfo{
		ActionInfo{"disable", `Disables the Alert indefinitely. Idempotent.`},
		ActionInfo{"enable", `Enables the Alert indefinitely. Idempotent.`},
		ActionInfo{"index", `Lists all Alerts.`},
		ActionInfo{"quench", `Suppresses the Alert from being triggered for a given time period. Idempotent.`},
		ActionInfo{"show", `Shows the attributes of a specified Alert.`},
	},
	"AlertSpec": []ActionInfo{
		ActionInfo{"create", `Creates a new AlertSpec with the given parameters.`},
		ActionInfo{"destroy", `Deletes a given AlertSpec.`},
		ActionInfo{"index", `No description provided for index.`},
		ActionInfo{"show", `No description provided for show.`},
		ActionInfo{"update", `Updates an AlertSpec with the given parameters.`},
	},
	"AuditEntry": []ActionInfo{
		ActionInfo{"append", `Updates the summary and appends more details to a given AuditEntry`},
		ActionInfo{"create", `Creates a new AuditEntry with the given parameters.`},
		ActionInfo{"detail", `shows the details of a given AuditEntry.`},
		ActionInfo{"index", `Lists AuditEntries of the account`},
		ActionInfo{"show", `Lists the attributes of a given audit entry.`},
		ActionInfo{"update", `Updates the summary of a given AuditEntry.`},
	},
	"Backup": []ActionInfo{
		ActionInfo{"cleanup", `Deletes old backups that meet the given criteria`},
		ActionInfo{"create", `Takes in an array of volumeattachmenthrefs and takes a snapshot of each.`},
		ActionInfo{"destroy", `Deletes a given backup by deleting all of its snapshots, this call will succeed even if the backup has not completed.`},
		ActionInfo{"index", `Lists all of the backups with the given lineage tag`},
		ActionInfo{"restore", `Restores the given Backup.`},
		ActionInfo{"show", `Lists the attributes of a given backup`},
		ActionInfo{"update", `Updates the committed tag for all of the VolumeSnapshots in the given Backup to the given value.`},
	},
	"ChildAccount": []ActionInfo{
		ActionInfo{"create", `Create an enterprise ChildAccount for this Account`},
		ActionInfo{"index", `Lists the enterprise ChildAccounts available for this Account.`},
		ActionInfo{"update", `Update an enterprise ChildAccount for this Account.`},
	},
	"Cloud": []ActionInfo{
		ActionInfo{"index", `Lists the clouds available to this account.`},
		ActionInfo{"show", `Show information about a single cloud`},
	},
	"CloudAccount": []ActionInfo{
		ActionInfo{"create", `Create a CloudAccount by passing in the respective credentials for each cloud.`},
		ActionInfo{"destroy", `Delete a CloudAccount.`},
		ActionInfo{"index", `Lists the CloudAccounts (non-aws) available to this Account.`},
		ActionInfo{"show", ``},
	},
	"Cookbook": []ActionInfo{
		ActionInfo{"destroy", `Destroys a Cookbook. Only available for cookbooks that have no Cookbook Attachments.`},
		ActionInfo{"follow", `Follows (or unfollows) a Cookbook. Only available for cookbooks that are in the Alternate namespace.`},
		ActionInfo{"freeze", `Freezes (or unfreezes) a Cookbook. Only available for cookbooks that are in the Primary namespace.`},
		ActionInfo{"index", `Lists the Cookbooks available to this account.`},
		ActionInfo{"obsolete", `Marks a Cookbook as obsolete (or un-obsolete).`},
		ActionInfo{"show", `Show information about a single Cookbook.`},
	},
	"CookbookAttachment": []ActionInfo{
		ActionInfo{"create", `Attach a cookbook to a given resource.`},
		ActionInfo{"destroy", `Detach a cookbook from a given resource.`},
		ActionInfo{"index", `Lists Cookbook Attachments.`},
		ActionInfo{"multi_attach", `Attach multiple cookbooks to a given resource.`},
		ActionInfo{"multi_detach", `Detach multiple cookbooks from a given resource.`},
		ActionInfo{"show", `Displays information about a single cookbook attachment to a ServerTemplate.`},
	},
	"Credential": []ActionInfo{
		ActionInfo{"create", `Creates a new Credential with the given parameters.`},
		ActionInfo{"destroy", `Deletes a Credential.`},
		ActionInfo{"index", `Lists the Credentials available to this account.`},
		ActionInfo{"show", `Show information about a single Credential. NOTE: Credential values may be updated through the API, but values cannot be retrieved via the API.`},
		ActionInfo{"update", `Updates attributes of a Credential.`},
	},
	"Datacenter": []ActionInfo{
		ActionInfo{"index", `Lists all Datacenters for a particular cloud.`},
		ActionInfo{"show", `Displays information about a single Datacenter.`},
	},
	"Deployment": []ActionInfo{
		ActionInfo{"clone", `Clones a given deployment.`},
		ActionInfo{"create", `Creates a new deployment with the given parameters.`},
		ActionInfo{"destroy", `Deletes a given deployment.`},
		ActionInfo{"index", `Lists deployments of the account.`},
		ActionInfo{"lock", `Locks a given deployment. Idempotent.`},
		ActionInfo{"servers", `Lists the servers belonging to this deployment`},
		ActionInfo{"show", `Lists the attributes of a given deployment.`},
		ActionInfo{"unlock", `Unlocks a given deployment. Idempotent.`},
		ActionInfo{"update", `Updates attributes of a given deployment.`},
	},
	"HealthCheck": []ActionInfo{
		ActionInfo{"index", `Check health of RightApi controllers`},
	},
	"IdentityProvider": []ActionInfo{
		ActionInfo{"index", `Lists the identity providers associated with this enterprise account.`},
		ActionInfo{"show", `Show the specified identity provider, if associated with this enterprise account.`},
	},
	"Image": []ActionInfo{
		ActionInfo{"index", `Lists all Images for the given Cloud.`},
		ActionInfo{"show", `Shows information about a single Image.`},
	},
	"Input": []ActionInfo{
		ActionInfo{"index", `Retrieves the full list of existing inputs of the specified resource.`},
		ActionInfo{"multi_update", `Performs a bulk update of inputs on the specified resource.`},
	},
	"Instance": []ActionInfo{
		ActionInfo{"create", `Creates and launches a raw instance using the provided parameters.`},
		ActionInfo{"index", `Lists instances of a given cloud, server array.`},
		ActionInfo{"launch", `Launches an instance using the parameters that this instance has been configured with.`},
		ActionInfo{"lock", ``},
		ActionInfo{"multi_run_executable", `Runs a script or a recipe in the running instances.`},
		ActionInfo{"multi_terminate", `Terminates running instances.`},
		ActionInfo{"reboot", `Reboot a running instance.`},
		ActionInfo{"run_executable", `Runs a script or a recipe in the running instance.`},
		ActionInfo{"set_custom_lodgement", `This method is deprecated.  Please use InstanceCustomLodgement.`},
		ActionInfo{"show", `Shows attributes of a single instance.`},
		ActionInfo{"start", `Starts an instance that has been stopped, resuming it to its previously saved volume state.`},
		ActionInfo{"stop", `Stores the instance's current volume state to resume later using the 'start' action.`},
		ActionInfo{"terminate", `Terminates a running instance.`},
		ActionInfo{"unlock", ``},
		ActionInfo{"update", `Updates attributes of a single instance.`},
	},
	"InstanceCustomLodgement": []ActionInfo{
		ActionInfo{"create", `Create a lodgement with the quantity and timeframe specified.`},
		ActionInfo{"destroy", `Destroy the specified lodgement.`},
		ActionInfo{"index", `List InstanceCustomLodgements of a given cloud and instance.`},
		ActionInfo{"show", `Show the specified lodgement.`},
		ActionInfo{"update", `Update a lodgement with the quantity specified.`},
	},
	"InstanceType": []ActionInfo{
		ActionInfo{"index", `Lists instance types.`},
		ActionInfo{"show", `Displays information about a single Instance type.`},
	},
	"IpAddress": []ActionInfo{
		ActionInfo{"create", `Creates a new IpAddress with the given parameters.`},
		ActionInfo{"destroy", `Deletes a given IpAddress.`},
		ActionInfo{"index", `Lists the IpAddresses available to this account.`},
		ActionInfo{"show", `Show information about a single IpAddress.`},
		ActionInfo{"update", `Updates attributes of a given IpAddress.`},
	},
	"IpAddressBinding": []ActionInfo{
		ActionInfo{"create", `Creates an ip address binding which attaches a specified IpAddress resource to a specified instance, and also allows for configuration of port forwarding rules...`},
		ActionInfo{"destroy", `No description provided for destroy.`},
		ActionInfo{"index", `Lists the ip address bindings available to this account.`},
		ActionInfo{"show", `Show information about a single ip address binding.`},
	},
	"MonitoringMetric": []ActionInfo{
		ActionInfo{"data", `Gives the raw monitoring data for a particular metric`},
		ActionInfo{"index", `Lists the monitoring metrics available for the instance and their corresponding graph hrefs.`},
		ActionInfo{"show", `Shows attributes of a single monitoring metric.`},
	},
	"MultiCloudImage": []ActionInfo{
		ActionInfo{"clone", `Clones a given MultiCloudImage.`},
		ActionInfo{"commit", `Commits a given MultiCloudImage. Only HEAD revisions can be committed.`},
		ActionInfo{"create", `Creates a new MultiCloudImage with the given parameters.`},
		ActionInfo{"destroy", `Deletes a given MultiCloudImage.`},
		ActionInfo{"index", `Lists the MultiCloudImages available to this account. HEAD revisions have a revision of 0.`},
		ActionInfo{"show", `Show information about a single MultiCloudImage. HEAD revisions have a revision of 0.`},
		ActionInfo{"update", `Updates attributes of a given MultiCloudImage. Only HEAD revisions can be updated (revision 0).`},
	},
	"MultiCloudImageSetting": []ActionInfo{
		ActionInfo{"create", `Creates a new setting for an existing MultiCloudImage.`},
		ActionInfo{"destroy", `Deletes a MultiCloudImage setting.`},
		ActionInfo{"index", `Lists the MultiCloudImage settings.`},
		ActionInfo{"show", `Show information about a single MultiCloudImage setting.`},
		ActionInfo{"update", `Updates a settings for a MultiCLoudImage.`},
	},
	"Network": []ActionInfo{
		ActionInfo{"create", `Creates a new network.`},
		ActionInfo{"destroy", `Deletes the given network(s).`},
		ActionInfo{"index", `Lists networks in this account.`},
		ActionInfo{"show", `Shows attributes of a single network.`},
		ActionInfo{"update", `Updates the given network.`},
	},
	"NetworkGateway": []ActionInfo{
		ActionInfo{"create", `Create a new NetworkGateway.`},
		ActionInfo{"destroy", `Delete an existing NetworkGateway.`},
		ActionInfo{"index", `Lists the NetworkGateways available to this account.`},
		ActionInfo{"show", `Show information about a single NetworkGateway.`},
		ActionInfo{"update", `Update an existing NetworkGateway.`},
	},
	"NetworkOptionGroup": []ActionInfo{
		ActionInfo{"create", `Create a new NetworkOptionGroup.`},
		ActionInfo{"destroy", `Delete an existing NetworkOptionGroup.`},
		ActionInfo{"index", `List NetworkOptionGroups available in this account.`},
		ActionInfo{"show", `Show information about a single NetworkOptionGroup.`},
		ActionInfo{"update", `Update an existing NetworkOptionGroup.`},
	},
	"NetworkOptionGroupAttachment": []ActionInfo{
		ActionInfo{"create", `Create a new NetworkOptionGroupAttachment.`},
		ActionInfo{"destroy", `Delete an existing NetworkOptionGroupAttachment.`},
		ActionInfo{"index", `List NetworkOptionGroupAttachments in this account.`},
		ActionInfo{"show", `Show information about a single NetworkOptionGroupAttachment.`},
		ActionInfo{"update", `Update an existing NetworkOptionGroupAttachment.`},
	},
	"Oauth2": []ActionInfo{
		ActionInfo{"create", `Perform an OAuth 2`},
	},
	"Permission": []ActionInfo{
		ActionInfo{"create", `Create a permission, thereby granting some user a particular role with respect to the current account...`},
		ActionInfo{"destroy", `Destroy a permission, thereby revoking a user's role with respect to the current account...`},
		ActionInfo{"index", `List all permissions for all users of the current acount.`},
		ActionInfo{"show", `Show information about a single permission.`},
	},
	"PlacementGroup": []ActionInfo{
		ActionInfo{"create", `Creates a PlacementGroup.`},
		ActionInfo{"destroy", `Destroys a PlacementGroup.`},
		ActionInfo{"index", `Lists all PlacementGroups in an account.`},
		ActionInfo{"show", `Shows information about a single PlacementGroup.`},
	},
	"Preference": []ActionInfo{
		ActionInfo{"destroy", `Deletes the given preference.`},
		ActionInfo{"index", `Lists all preferences.`},
		ActionInfo{"show", `Shows a single preference.`},
		ActionInfo{"update", `If 'id' is known, updates preference with given contents.`},
	},
	"Publication": []ActionInfo{
		ActionInfo{"import", `Imports the given publication and its subordinates to this account.`},
		ActionInfo{"index", `Lists the publications available to this account. Only non-HEAD revisions are possible.`},
		ActionInfo{"show", `Show information about a single publication. Only non-HEAD revisions are possible.`},
	},
	"PublicationLineage": []ActionInfo{
		ActionInfo{"show", `Show information about a single publication lineage. Only non-HEAD revisions are possible.`},
	},
	"RecurringVolumeAttachment": []ActionInfo{
		ActionInfo{"create", `Creates a new recurring volume attachment.`},
		ActionInfo{"destroy", `Deletes a given recurring volume attachment.`},
		ActionInfo{"index", `Lists all recurring volume attachments.`},
		ActionInfo{"show", `Displays information about a single recurring volume attachment.`},
	},
	"Repository": []ActionInfo{
		ActionInfo{"cookbook_import", `Performs a Cookbook import, which allows you to use the specified cookbooks in your design objects.`},
		ActionInfo{"cookbook_import_preview", `Retrieves a preview of the effects of a Cookbook import.`},
		ActionInfo{"create", `Creates a Repository.`},
		ActionInfo{"destroy", `Deletes the specified Repositories.`},
		ActionInfo{"index", `Lists all Repositories for this Account.`},
		ActionInfo{"refetch", `Refetches all RepositoryAssets associated with the Repository.`},
		ActionInfo{"resolve", `Show a list of repositories that have imported cookbooks with the given names.`},
		ActionInfo{"show", `Shows a specified Repository.`},
		ActionInfo{"update", `Updates a specified Repository.`},
	},
	"RepositoryAsset": []ActionInfo{
		ActionInfo{"index", `List a repository's current assets.`},
		ActionInfo{"show", `Show information about a single asset.`},
	},
	"RightScript": []ActionInfo{
		ActionInfo{"commit", `Commits the given RightScript. Only HEAD revisions (revision 0) can be committed.`},
		ActionInfo{"index", `Lists RightScripts.`},
		ActionInfo{"show", `Displays information about a single RightScript.`},
		ActionInfo{"show_source", `Returns the script source for a RightScript`},
		ActionInfo{"update", `Updates RightScript name/description`},
		ActionInfo{"update_source", `Updates the source of the given RightScript`},
	},
	"Route": []ActionInfo{
		ActionInfo{"create", `Create a new Route.`},
		ActionInfo{"destroy", `Delete an existing Route.`},
		ActionInfo{"index", `List Routes available in this account.`},
		ActionInfo{"show", `Show information about a single Route.`},
		ActionInfo{"update", `Update an existing Route.`},
	},
	"RouteTable": []ActionInfo{
		ActionInfo{"create", `Create a new RouteTable.`},
		ActionInfo{"destroy", `Delete an existing RouteTable.`},
		ActionInfo{"index", `List RouteTables available in this account.`},
		ActionInfo{"show", `Show information about a single RouteTable.`},
		ActionInfo{"update", `Update an existing RouteTable.`},
	},
	"RunnableBinding": []ActionInfo{
		ActionInfo{"create", `Bind an executable to the given ServerTemplate.`},
		ActionInfo{"destroy", `Unbind an executable from the given resource.`},
		ActionInfo{"index", `Lists the executables bound to the ServerTemplate.`},
		ActionInfo{"multi_update", `Update attributes for multiple bound executables.`},
		ActionInfo{"show", `Show information about a single executable binding.`},
	},
	"SecurityGroup": []ActionInfo{
		ActionInfo{"create", `Create a security group.`},
		ActionInfo{"destroy", `Delete security group(s)`},
		ActionInfo{"index", `Lists Security Groups.`},
		ActionInfo{"show", `Displays information about a single Security Group.`},
	},
	"SecurityGroupRule": []ActionInfo{
		ActionInfo{"create", `Create a security group rule for a security group.`},
		ActionInfo{"destroy", `Delete security group rule(s)`},
		ActionInfo{"index", `Lists SecurityGroupRules.`},
		ActionInfo{"show", `Displays information about a single SecurityGroupRule.`},
		ActionInfo{"update", ``},
	},
	"Server": []ActionInfo{
		ActionInfo{"clone", `Clones a given server.`},
		ActionInfo{"create", `Creates a new server, and configures its corresponding "next" instance with the received parameters.`},
		ActionInfo{"destroy", `Deletes a given server.`},
		ActionInfo{"index", `Lists servers.`},
		ActionInfo{"launch", `Launches the "next" instance of this server`},
		ActionInfo{"show", `Shows the information of a single server.`},
		ActionInfo{"terminate", `Terminates the current instance of this server`},
		ActionInfo{"update", `Updates attributes of a single server.`},
		ActionInfo{"wrap_instance", `Wrap an existing instance and set current instance for new server`},
	},
	"ServerArray": []ActionInfo{
		ActionInfo{"clone", `Clones a given server array.`},
		ActionInfo{"create", `Creates a new server array, and configures its corresponding "next" instance with the received parameters.`},
		ActionInfo{"current_instances", `List the running instances belonging to the server array. See Instances#index for details.`},
		ActionInfo{"destroy", `Deletes a given server array.`},
		ActionInfo{"index", `Lists server arrays.`},
		ActionInfo{"launch", `Launches a new instance in the server array with the configuration defined in the 'next_instance'`},
		ActionInfo{"multi_run_executable", `Run an executable on all instances of this array`},
		ActionInfo{"multi_terminate", `Terminate all instances of this array`},
		ActionInfo{"show", `Shows the information of a single server array.`},
		ActionInfo{"update", `Updates attributes of a single server array.`},
	},
	"ServerTemplate": []ActionInfo{
		ActionInfo{"clone", `Clones a given ServerTemplate.`},
		ActionInfo{"commit", `Commits a given ServerTemplate. Only HEAD revisions (revision 0) that are owned by the account can be committed.`},
		ActionInfo{"create", `Creates a new ServerTemplate with the given parameters.`},
		ActionInfo{"destroy", `Deletes a given ServerTemplate.`},
		ActionInfo{"detect_changes_in_head", `Identifies RightScripts attached to the resource that differ from their HEAD.`},
		ActionInfo{"index", `Lists the ServerTemplates available to this account. HEAD revisions have a revision of 0.`},
		ActionInfo{"publish", `Publishes a given ServerTemplate and its subordinates.`},
		ActionInfo{"resolve", `Enumerates all attached cookbooks, missing dependencies and bound executables.`},
		ActionInfo{"show", `Show information about a single ServerTemplate. HEAD revisions have a revision of 0.`},
		ActionInfo{"swap_repository", `In-place replacement of attached cookbooks from a given repository.`},
		ActionInfo{"update", `Updates attributes of a given ServerTemplate. Only HEAD revisions can be updated (revision 0).`},
	},
	"ServerTemplateMultiCloudImage": []ActionInfo{
		ActionInfo{"create", `Creates a new ServerTemplateMultiCloudImage with the given parameters.`},
		ActionInfo{"destroy", `Deletes a given ServerTemplateMultiCloudImage.`},
		ActionInfo{"index", `Lists the ServerTemplateMultiCloudImages available to this account.`},
		ActionInfo{"make_default", `Makes a given ServerTemplateMultiCloudImage the default for the ServerTemplate.`},
		ActionInfo{"show", `Show information about a single ServerTemplateMultiCloudImage which represents an association between a ServerTemplate and a MultiCloudImage.`},
	},
	"Session": []ActionInfo{
		ActionInfo{"accounts", `List all the accounts that a user has access to.`},
		ActionInfo{"index", `Returns a list of root resources so an authenticated session can use them as a starting point or a way to know what features are available within its privileges...`},
		ActionInfo{"index_instance_session", `Shows the full attributes of the instance (that has the token used to log-in).`},
	},
	"SshKey": []ActionInfo{
		ActionInfo{"create", `Creates a new ssh key.`},
		ActionInfo{"destroy", `Deletes a given ssh key.`},
		ActionInfo{"index", `Lists ssh keys.`},
		ActionInfo{"show", `Displays information about a single ssh key.`},
	},
	"Subnet": []ActionInfo{
		ActionInfo{"create", `Creates a new subnet.`},
		ActionInfo{"destroy", `Deletes the given subnet(s).`},
		ActionInfo{"index", `Lists subnets of a given cloud.`},
		ActionInfo{"show", `Shows attributes of a single subnet.`},
		ActionInfo{"update", `Updates name and description for the current subnet.`},
	},
	"Tag": []ActionInfo{
		ActionInfo{"by_resource", `Get tags for a list of resource hrefs.`},
		ActionInfo{"by_tag", `Search for resources having a list of tags in a specific resource_type.`},
		ActionInfo{"multi_add", `Add a list of tags to a list of hrefs. The tags must be either plain_tags or machine_tags.`},
		ActionInfo{"multi_delete", `Delete a list of tags on a list of hrefs. The tags must be either plain_tags or machine_tags.`},
	},
	"Task": []ActionInfo{
		ActionInfo{"show", `Displays information of a given task within the context of an instance.`},
	},
	"User": []ActionInfo{
		ActionInfo{"create", `Create a user. If a user already exists with the same email, that user will be returned.`},
		ActionInfo{"index", `List the users available to the account the user is logged in to`},
		ActionInfo{"show", `Show information about a single user.`},
		ActionInfo{"update", `Update a user's contact information, change her password, or update SSO her settings`},
	},
	"UserData": []ActionInfo{
		ActionInfo{"show", `No description provided for show.`},
	},
	"Volume": []ActionInfo{
		ActionInfo{"create", `Creates a new volume.`},
		ActionInfo{"destroy", `Deletes a given volume.`},
		ActionInfo{"index", `Lists volumes.`},
		ActionInfo{"show", `Displays information about a single volume.`},
	},
	"VolumeAttachment": []ActionInfo{
		ActionInfo{"create", `Creates a new volume attachment.`},
		ActionInfo{"destroy", `Deletes a given volume attachment.`},
		ActionInfo{"index", `Lists all volume attachments.`},
		ActionInfo{"show", `Displays information about a single volume attachment.`},
	},
	"VolumeSnapshot": []ActionInfo{
		ActionInfo{"create", `Creates a new volume_snapshot.`},
		ActionInfo{"destroy", `Deletes a given volume_snapshot.`},
		ActionInfo{"index", `Lists all volume_snapshots.`},
		ActionInfo{"show", `Displays information about a single volume_snapshot.`},
	},
	"VolumeType": []ActionInfo{
		ActionInfo{"index", `Lists Volume Types.`},
		ActionInfo{"show", `Displays information about a single Volume Type.`},
	},
}
