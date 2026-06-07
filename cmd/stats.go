package cmd

import (
	"fmt"
	"neko/storage"
	"os"
	"sort"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "View your focus stats and weekly playtime",
	Run: func(cmd *cobra.Command, args []string) {
		setupTerminal()

		history, err := storage.LoadHistory()
		if err != nil {
			fmt.Println("Error loading history:", err)
			os.Exit(1)
		}

		now := time.Now()
		thirtyDaysAgo := now.AddDate(0, 0, -30)
		tagMinutes := make(map[string]int)
		totalMinutes30 := 0

		for _, s := range history.Sessions {
			if s.Status == "rested" && s.Recipe != "break" {
				t := time.Unix(s.StartTimeUnix, 0)
				if t.After(thirtyDaysAgo) {
					tagName := s.Tag
					if tagName == "" {
						tagName = "untagged"
					}
					tagMinutes[tagName] += s.DurationMinutes
					totalMinutes30 += s.DurationMinutes
				}
			}
		}

		titleStyle := lipgloss.NewStyle().Padding(1, 0, 1, 0).Foreground(lipgloss.Color("#FFD700")).Bold(true)
		fmt.Println(titleStyle.Render("=== 30-DAY FOCUS STATS ==="))

		if totalMinutes30 == 0 {
			fmt.Println("  No focus sessions in the last 30 days.")
		} else {
			type TagStat struct {
				Name string
				Mins int
			}
			var tagStats []TagStat
			for tag, mins := range tagMinutes {
				tagStats = append(tagStats, TagStat{Name: tag, Mins: mins})
			}

			sort.Slice(tagStats, func(i, j int) bool {
				return tagStats[i].Mins > tagStats[j].Mins
			})

			for _, ts := range tagStats {
				tag := ts.Name
				mins := ts.Mins
				pct := float64(mins) / float64(totalMinutes30) * 100
				barLen := int(pct / 2) // 50 chars max width
				if barLen < 1 && pct > 0 {
					barLen = 1
				}

				bar := ""
				for i := 0; i < barLen; i++ {
					bar += "█"
				}

				color := "#00FF00" // Default green
				if tag == "untagged" {
					color = "#888888"
				} else if len(tag) > 0 {
					colors := []string{"#FF5733", "#33FF57", "#3357FF", "#F033FF", "#33FFF0", "#FFC300"}
					color = colors[len(tag)%len(colors)]
				}

				barStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(color))

				fmt.Printf(" %-12s | %s %.1f%% (%d min)\n", tag, barStyle.Render(bar), pct, mins)
			}
		}

		// Weekly shelf
		weekday := int(now.Weekday())
		if weekday == 0 {
			weekday = 7 // Make Sunday 7 instead of 0
		}
		startOfWeek := now.AddDate(0, 0, -weekday+1)
		startOfWeek = time.Date(startOfWeek.Year(), startOfWeek.Month(), startOfWeek.Day(), 0, 0, 0, 0, now.Location())

		endOfWeek := startOfWeek.AddDate(0, 0, 7)

		sessionsByDay := make(map[int][]storage.Session)
		for _, s := range history.Sessions {
			t := time.Unix(s.StartTimeUnix, 0)
			if t.After(startOfWeek) && t.Before(endOfWeek) {
				dayIdx := int(t.Weekday()) - 1
				if dayIdx == -1 {
					dayIdx = 6 // Sunday
				}
				sessionsByDay[dayIdx] = append(sessionsByDay[dayIdx], s)
			}
		}

		days := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

		var rows [][]string
		maxRows := 0
		for i := 0; i < 7; i++ {
			if len(sessionsByDay[i]) > maxRows {
				maxRows = len(sessionsByDay[i])
			}
		}
		if maxRows == 0 {
			maxRows = 1
		}

		for r := 0; r < maxRows; r++ {
			row := make([]string, 7)
			for d := 0; d < 7; d++ {
				if r < len(sessionsByDay[d]) {
					s := sessionsByDay[d][r]
					if s.Status == "startled" {
						row[d] = " ! "
					} else {
						row[d] = "^ω^"
					}
				} else {
					row[d] = "---"
				}
			}
			rows = append(rows, row)
		}

		re := lipgloss.NewRenderer(os.Stdout)
		baseStyle := re.NewStyle().Padding(0, 1).Align(lipgloss.Center)
		headerStyle := baseStyle

		t := table.New().
			Border(lipgloss.HiddenBorder()).
			Headers(days...).
			Rows(rows...).
			StyleFunc(func(row, col int) lipgloss.Style {
				if row == table.HeaderRow {
					return headerStyle
				}
				if col >= 0 && col < 7 {
					if len(rows) > 0 && row > 0 {
						val := rows[row-1][col]
						if val == "---" {
							return baseStyle
						}
					}
				}
				return baseStyle
			})

		fmt.Println(titleStyle.Render("=== YOUR WEEKLY PLAYTIME ==="))
		fmt.Println(t.Render())

	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
