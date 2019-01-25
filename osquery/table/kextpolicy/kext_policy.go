package kextpolicy

import (
	"context"

	"github.com/knightsc/system_policy/sp"
	"github.com/kolide/osquery-go/plugin/table"
)

func TablePlugin() *table.Plugin {
	columns := []table.ColumnDefinition{
		table.TextColumn("developer_name"),
		table.TextColumn("application_name"),
		table.TextColumn("application_path"),
		table.TextColumn("team_id"),
		table.TextColumn("bundle_id"),
		table.IntegerColumn("allowed"),
		table.IntegerColumn("reboot_required"),
		table.IntegerColumn("modified"),
	}
	return table.NewPlugin("kext_policy", columns, generate)
}

func generate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	results := make([]map[string]string, 0)

	items := sp.CurrentKernelExtensionPolicy()
	for _, item := range items {
		row := map[string]string{}
		row["developer_name"] = item.DeveloperName
		if item.ApplicationName != "" {
			row["application_name"] = item.ApplicationName
		}
		if item.ApplicationPath != "" {
			row["application_path"] = item.ApplicationPath
		}
		row["team_id"] = item.TeamID
		row["bundle_id"] = item.BundleID
		if item.Allowed {
			row["allowed"] = "1"
		} else {
			row["allowed"] = "0"
		}
		if item.RebootRequired {
			row["reboot_required"] = "1"
		} else {
			row["reboot_required"] = "0"
		}
		if item.Modified {
			row["modified"] = "1"
		} else {
			row["modified"] = "0"
		}

		results = append(results, row)
	}

	return results, nil
}
