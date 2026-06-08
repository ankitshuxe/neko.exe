package cmd

import (
	"github.com/ankitshuxe/neko.exe/storage"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var diaryCmd = &cobra.Command{
	Use:   "diary",
	Short: "View your daily purr time and session history",
	Run: func(cmd *cobra.Command, args []string) {
		setupTerminal()

		history, err := storage.LoadHistory()
		if err != nil {
			fmt.Println("Error loading history:", err)
			os.Exit(1)
		}

		now := time.Now()
		todayDateStr := now.Format("2006-01-02")

		var successCount int
		var spilledCount int
		var todayMinutes int

		recipeCounts := make(map[string]int)

		for _, s := range history.Sessions {
			if s.Date == todayDateStr {
				switch s.Status {
				case "rested":
					successCount++
					if s.Recipe != "break" {
						todayMinutes += s.DurationMinutes
					}
					key := s.Recipe + "|" + s.Tag
					recipeCounts[key]++
				case "startled":
					spilledCount++
				}
			}
		}

		var sb strings.Builder
		sb.WriteString("        (=^ω^=) NEKO DIARY (=^ω^=)       \n")
		sb.WriteString("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
		sb.WriteString(fmt.Sprintf(" Date: %s (Today)\n", todayDateStr))
		sb.WriteString("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")

		if successCount == 0 && spilledCount == 0 {
			sb.WriteString("\n       No naps taken today yet.      \n\n")
		} else {
			for key, count := range recipeCounts {
				parts := strings.Split(key, "|")
				recipe := parts[0]
				tag := parts[1]

				name := recipe
				if name == "default" {
					name = "Yarn Ball"
				}
				if name == "break" {
					name = "Naps"
				}
				if name == "french_press" {
					name = "Cardboard Box"
				}
				if name == "matcha" {
					name = "Catnip Mouse"
				}

				if tag != "" {
					name = name + " (" + tag + ")"
				}

				var rMins int
				for _, s := range history.Sessions {
					if s.Date == todayDateStr && s.Status == "rested" && s.Recipe == recipe && s.Tag == tag {
						rMins += s.DurationMinutes
					}
				}

				line := fmt.Sprintf("%dx %s ", count, name)
				tail := fmt.Sprintf(" %d min", rMins)
				dots := 41 - len(line) - len(tail)
				if dots < 1 {
					dots = 1
				}

				sb.WriteString(fmt.Sprintf("%s%s%s\n", line, strings.Repeat(".", dots), tail))
			}

			if spilledCount > 0 {
				line := fmt.Sprintf("%dx Startled! ", spilledCount)
				tail := " -"
				dots := 41 - len(line) - len(tail)
				if dots < 1 {
					dots = 1
				}
				sb.WriteString(fmt.Sprintf("%s%s%s\n", line, strings.Repeat(".", dots), tail))
			}
		}

		sb.WriteString("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
		sb.WriteString(fmt.Sprintf(" TODAY'S PURR TIME:    %d min\n", todayMinutes))
		sb.WriteString(fmt.Sprintf(" TOTAL PURR TIME:      %d min\n", history.TotalFocusMinutes))
		sb.WriteString("=========================================\n")
		sb.WriteString("            Stay cozy, meow!             \n")

		diaryStyle := lipgloss.NewStyle().
			Padding(1, 2).
			Margin(1, 0)

		fmt.Println(diaryStyle.Render(sb.String()))
	},
}

func init() {
	rootCmd.AddCommand(diaryCmd)
}
