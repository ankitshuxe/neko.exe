package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "neko",
	Short: "A cozy, terminal-based Pomodoro timer",
	Run: func(cmd *cobra.Command, args []string) {
		titleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#F05522")).Bold(true)
		cmdStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#F05522")).Bold(true)
		descStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#E4D8C8"))

		asciiArt := `
    /\_/\
   ( o.o )  Welcome to Neko.exe!
    > ^ <   Your strictly CLI-based Pomodoro companion.
`
		fmt.Println(titleStyle.Render(asciiArt))
		
		fmt.Println(cmdStyle.Render("  neko start [minutes]") + descStyle.Render("  - Start a focus session (default 25m). Earn 1 Fish Coin per minute!"))
		fmt.Println(descStyle.Render("                           Flags: -t <tag> (e.g. coding), -s <sound> (purr, rain, lofi)"))
		fmt.Println(cmdStyle.Render("  neko break [minutes]") + descStyle.Render("  - Start a break. Your cat goes to sleep."))
		fmt.Println(cmdStyle.Render("  neko shop") + descStyle.Render("             - View the shop, your Fish Coins, and buy new toys/environments/breeds."))
		fmt.Println(cmdStyle.Render("  neko shop buy <id>") + descStyle.Render("    - Buy an item using its ID."))
		fmt.Println(cmdStyle.Render("  neko shop equip <id>") + descStyle.Render("  - Equip an item you own to see it during your sessions."))
		fmt.Println(cmdStyle.Render("  neko diary") + descStyle.Render("            - View today's focus sessions grouped by tags."))
		fmt.Println(cmdStyle.Render("  neko stats") + descStyle.Render("            - View your 30-day focus analytics and weekly grid."))
		fmt.Println()
		fmt.Println(descStyle.Render("  Tip: During a session, press 'p' to pause/resume or 'q' to quit."))
		fmt.Println()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
