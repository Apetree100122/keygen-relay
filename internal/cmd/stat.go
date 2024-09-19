package cmd

import (
	"fmt"
	"github.com/keygen-sh/keygen-relay/internal/ui"
	"strconv"
	"time"

	"database/sql"
	"github.com/charmbracelet/bubbles/table"
	"github.com/keygen-sh/keygen-relay/internal/licenses"
	"github.com/spf13/cobra"
)

func formatTime(t sql.NullString) string {
	if t.Valid {
		parsedTime, err := time.Parse(time.RFC3339, t.String)
		if err == nil {
			return parsedTime.Format("2006-01-02 15:04:05")
		}
	}
	return "-"
}

func StatCmd(manager licenses.Manager, tableRenderer ui.TableRenderer) *cobra.Command {
	var licenseID string

	cmd := &cobra.Command{
		Use:   "stat",
		Short: "Print stats for a license in the local relay server's pool",
		RunE: func(cmd *cobra.Command, args []string) error {
			license, err := manager.GetLicenseByID(cmd.Context(), licenseID)
			if err != nil {
				return err
			}

			columns := []table.Column{
				{Title: "ID", Width: 36},
				{Title: "Key", Width: 36},
				{Title: "Claims", Width: 8},
				{Title: "NodeID", Width: 8},
				{Title: "Last Claimed At", Width: 20},
				{Title: "Last Released At", Width: 20},
			}

			claimsStr := fmt.Sprintf("%d", license.Claims)

			var nodeIDStr string
			if license.NodeID.Valid {
				nodeIDStr = strconv.FormatInt(license.NodeID.Int64, 10)
			} else {
				nodeIDStr = "-"
			}

			lastClaimedAtStr := formatTime(license.LastClaimedAt)
			lastReleasedAtStr := formatTime(license.LastReleasedAt)

			tableRows := []table.Row{
				{license.ID, license.Key, claimsStr, nodeIDStr, lastClaimedAtStr, lastReleasedAtStr},
			}

			if err := tableRenderer.Render(tableRows, columns); err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Error rendering table: %v", err)
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&licenseID, "id", "", "License ID to print stats for")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}