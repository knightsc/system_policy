package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/knightsc/system_policy/sp"

	"github.com/kolide/osquery-go"
	"github.com/kolide/osquery-go/plugin/table"
)

func main() {
	flSocket := flag.String("socket", "", "")
	flag.Int("timeout", 0, "")
	flag.Int("interval", 0, "")
	flag.Bool("verbose", false, "")
	flag.Parse()

	if *flSocket == "" {
		log.Fatalln("--socket flag cannot be empty")
	}

	server, err := osquery.NewExtensionManagerServer("system_policy", *flSocket)
	if err != nil {
		log.Fatalf("Error creating osquery extension server: %s\n", err)
	}

	server.RegisterPlugin(table.NewPlugin("legacy_exec_history", execColumns(), execGenerate))
	server.RegisterPlugin(table.NewPlugin("kext_policy", kextColumns(), kextGenerate))

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func execColumns() []table.ColumnDefinition {
	return []table.ColumnDefinition{
		table.TextColumn("exec_path"),
		table.TextColumn("mmap_path"),
		table.TextColumn("signing_id"),
		table.TextColumn("team_id"),
		table.TextColumn("cd_hash"),
		table.TextColumn("responsible_path"),
		table.TextColumn("developer_name"),
		table.TextColumn("last_seen"),
	}
}

func execGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	results := make([]map[string]string, 0)

	items := sp.LegacyExecutionHistory()
	for _, item := range items {
		row := map[string]string{}
		row["exec_path"] = item.ExecPath
		row["last_seen"] = item.LastSeen.Format(time.RFC3339)
		if item.MmapPath != "" {
			row["mmap_path"] = item.MmapPath
		}
		if item.SigningID != "" {
			row["signing_id"] = item.SigningID
		}
		if item.TeamID != "" {
			row["team_id"] = item.TeamID
		}
		if item.CDHash != "" {
			row["cd_hash"] = item.CDHash
		}
		if item.ResponsiblePath != "" {
			row["responsible_path"] = item.ResponsiblePath
		}
		if item.DeveloperName != "" {
			row["developer_name"] = item.DeveloperName
		}

		results = append(results, row)
	}

	return results, nil
}

func kextColumns() []table.ColumnDefinition {
	return []table.ColumnDefinition{
		table.TextColumn("developer_name"),
		table.TextColumn("application_name"),
		table.TextColumn("application_path"),
		table.TextColumn("team_id"),
		table.TextColumn("bundle_id"),
		table.IntegerColumn("allowed"),
		table.IntegerColumn("reboot_required"),
		table.IntegerColumn("modified"),
	}
}

func kextGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
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
			row ["allowed"] = "0"
		}
		if item.RebootRequired {
			row["reboot_required"] = "1"
		} else {
			row ["reboot_required"] = "0"
		}
		if item.Modified {
			row["modified"] = "1"
		} else {
			row ["modified"] = "0"
		}

		results = append(results, row)
	}

	return results, nil
}
