package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/knightsc/legacy_exec_history/sp"

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

	server, err := osquery.NewExtensionManagerServer("legacy_exec_history", *flSocket)
	if err != nil {
		log.Fatalf("Error creating osquery extension server: %s\n", err)
	}

	server.RegisterPlugin(table.NewPlugin("legacy_exec_history", columns(), generate))

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func columns() []table.ColumnDefinition {
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

func generate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
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
