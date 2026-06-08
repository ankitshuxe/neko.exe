package cmd

import (
	"github.com/ankitshuxe/neko.exe/storage"
	"fmt"
	"os"
	"strconv"
	"time"

	"embed"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gen2brain/beeep"
	"github.com/google/uuid"
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/gopxl/beep/v2/wav"
	"github.com/spf13/cobra"
)

var (
	monoColor = lipgloss.Color("#E4D8C8")
	monoStyle = lipgloss.NewStyle().Foreground(monoColor)
)

type timerModel struct {
	session        storage.Session
	totalSeconds   int
	remaining      int
	status         string
	animationFrame int
	sessionType    string
	equippedToy    string
	equippedEnv    string
	equippedBreed  string
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second/2, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m timerModel) Init() tea.Cmd {
	return tickCmd()
}

func (m timerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.status = "startled"
			
			history, err := storage.LoadHistory()
			if err == nil {
				elapsedSeconds := m.totalSeconds - m.remaining
				if m.sessionType != "break" && history.FishCoins > 0 && elapsedSeconds > 600 {
					history.FishCoins--
				}
				for i := range history.Sessions {
					if history.Sessions[i].ID == m.session.ID {
						history.Sessions[i].Status = "startled"
						_ = storage.SaveHistory(history)
						break
					}
				}
			}

			return m, tea.Batch(tea.ClearScreen, tea.Quit)
		case "p":
			switch m.status {
			case "running":
				m.status = "paused"
			case "paused":
				m.status = "running"
				m.session.StartTimeUnix = time.Now().Unix() - int64(m.totalSeconds - m.remaining)
			}
			return m, nil
		}
	case tickMsg:
		if m.status == "paused" {
			return m, tickCmd()
		}
		if m.status != "running" {
			return m, nil
		}

		currentUnix := time.Now().Unix()
		targetUnix := m.session.StartTimeUnix + int64(m.totalSeconds)
		m.remaining = int(targetUnix - currentUnix)

		m.animationFrame++

		if m.remaining <= 0 {
			m.remaining = 0
			m.status = "rested"
			return m, tea.Batch(tea.ClearScreen, tea.Quit)
		}

		return m, tickCmd()
	}
	return m, nil
}

func (m timerModel) View() string {
	var statusText string
	switch m.status {
	case "running":
		if m.sessionType == "break" {
			statusText = " [ STATUS // SLEEPING ]"
		} else {
			statusText = " [ STATUS // PLAYING  ]"
		}
	case "paused":
		statusText = " [ STATUS // PAUSED   ]"
	case "startled":
		statusText = " [ STATUS // CANCELED ]"
	case "rested":
		statusText = " [ STATUS // RESTED   ]"
	default:
		statusText = " [ STATUS // UNKNOWN  ]"
	}

	currentArt := composeArt(m)

	mins := m.remaining / 60
	secs := m.remaining % 60
	timeString := fmt.Sprintf(" [ TIMER  // %02d:%02d ]", mins, secs)

	artRender := lipgloss.NewStyle().Align(lipgloss.Center).Width(26).Render(monoStyle.Render(currentArt))
	divider := monoStyle.Render("──────────────────────────")

	timeRender := lipgloss.NewStyle().Align(lipgloss.Center).Width(26).Render(monoStyle.Render(timeString))
	statusRender := lipgloss.NewStyle().Align(lipgloss.Center).Width(26).Render(monoStyle.Render(statusText))

	uiBox := lipgloss.JoinVertical(lipgloss.Left,
		artRender,
		divider,
		timeRender,
		statusRender,
		divider,
	)

	return lipgloss.NewStyle().
		Width(80).Align(lipgloss.Center).
		Padding(1).
		Render(uiBox)
}

var sessionTag string

func composeArt(m timerModel) string {
	breedFaces := map[string]string{
		"breed_tuxedo":       "(=^･ω･^=)",
		"breed_tabby":        "(=^◕ω◕^=)",
		"breed_calico":       "(=^•ω•^=)",
		"breed_siamese":      "(=^☼ω☼^=)",
		"breed_maine_coon":   "(=^ﻌ^=)∫",
		"breed_russian_blue": "[=^･ω･^=]",
		"breed_sphynx":       "(=^⚆ω⚆^=)",
		"breed_bengal":       "(=^✧ω✧^=)",
		"breed_persian":      "(ﾐΦ ﻌ Φﾐ)",
		"breed_galactic":     "(=^★ω★^=)",
	}

	face := breedFaces[m.equippedBreed]
	if face == "" {
		face = "(=^･ω･^=)"
	}

	if m.status == "startled" {
		face = "(=>.<=)"
	} else if m.status == "paused" {
		face = "(=O.O=)"
	} else if m.status == "rested" || m.sessionType == "break" {
		face = "(=^ω^=)"
		if m.equippedBreed == "breed_russian_blue" { face = "[=^ω^=]" }
		if m.equippedBreed == "breed_persian" { face = "(ﾐΦ ω Φﾐ)" }
	}

	envLines := []string{
		"          ",
		"          ",
		"          ",
	}

	switch m.equippedEnv {
	case "env_window_sill":
		envLines[0] = "  |￣￣￣￣|"
		envLines[1] = "  |        |"
		envLines[2] = "  |＿＿＿＿|"
	case "env_cardboard_castle":
		envLines[0] = "   /\\_/\\  "
		envLines[1] = "  |     | "
		envLines[2] = "  |_____| "
	case "env_cozy_blanket":
		envLines[2] = "  ~~~~~~~~"
	case "env_cat_tree":
		envLines[0] = "      === "
		envLines[1] = "      |   "
		envLines[2] = "     ===  "
	}

	toyArt := "  "
	if m.sessionType != "break" && m.status == "running" {
		frame := m.animationFrame % 4
		switch m.equippedToy {
		case "toy_yarn_ball":
			toys := []string{" O", "o ", " O", "  o"}
			toyArt = toys[frame]
		case "toy_cardboard_box":
			toys := []string{"[ ]", "[ ]", "[ ]", "[ ]"}
			toyArt = toys[frame]
		case "toy_catnip_mouse":
			toys := []string{"~(:3 )", "~(:3)", "~(:3 )", "~(:3)"}
			toyArt = toys[frame]
		case "toy_feather_wand":
			toys := []string{"\\", "/", "\\", "-"}
			toyArt = toys[frame]
		case "toy_laser_pointer":
			toys := []string{".", " .", "  .", " ."}
			toyArt = toys[frame]
		default:
			toys := []string{" *", "* ", " *", "  *"}
			toyArt = toys[frame]
		}
	} else if m.sessionType == "break" && m.status == "running" {
		frame := m.animationFrame % 3
		switch frame {
		case 0:
			envLines[1] = "       z  "
		case 1:
			envLines[1] = "       Z  "
		default:
			envLines[0] = "        Z "
		}
	}

	return fmt.Sprintf("%s\n%s\n%s\n    %s%s", envLines[0], envLines[1], envLines[2], face, toyArt)
}

var sessionSound string

//go:embed sounds/*.wav
var embeddedSounds embed.FS

var startCmd = &cobra.Command{
	Use:   "start [minutes]",
	Short: "Start a new session (timer). Defaults to 25 minutes.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		setupTerminal()

		minutes := 25
		if len(args) > 0 {
			parsed, err := strconv.Atoi(args[0])
			if err != nil || parsed <= 0 {
				os.Exit(1)
			}
			minutes = parsed
		}

		history, err := storage.LoadHistory()
		if err != nil {
			os.Exit(1)
		}

		if len(history.Sessions) > 0 {
			lastIdx := len(history.Sessions) - 1
			if history.Sessions[lastIdx].Status == "in_progress" {
				history.Sessions[lastIdx].Status = "startled"
				_ = storage.SaveHistory(history)
			}
		}

		now := time.Now()
		session := storage.Session{
			ID:              uuid.New().String(),
			Date:            now.Format("2006-01-02"),
			StartTimeUnix:   now.Unix(),
			DurationMinutes: minutes,
			Recipe:          "default",
			Tag:             sessionTag,
			Status:          "in_progress",
		}

		history.Sessions = append(history.Sessions, session)
		if err := storage.SaveHistory(history); err != nil {
			os.Exit(1)
		}

		m := timerModel{
			session:        session,
			totalSeconds:   minutes * 60,
			remaining:      minutes * 60,
			status:         "running",
			animationFrame: 0,
			sessionType:    "work",
			equippedToy:    history.EquippedToy,
			equippedEnv:    history.EquippedEnv,
			equippedBreed:  history.EquippedBreed,
		}

		if sessionSound != "" {
			path := "sounds/" + sessionSound + ".wav"
			f, err := embeddedSounds.Open(path)
			if err == nil {
				defer f.Close()
				streamer, format, err := wav.Decode(f)
				if err == nil {
					defer streamer.Close()
					_ = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
					loop := beep.Loop(-1, streamer)
					speaker.Play(loop)
				} else {
					fmt.Printf("Error decoding sound: %v\n", err)
				}
			} else {
				fmt.Printf("Unknown sound '%s'. Try: purr, rain, lofi.\n", sessionSound)
			}
		}

		p := tea.NewProgram(m)
		finalModel, err := p.Run()
		if err != nil {
			fmt.Printf("Error running timer: %v\n", err)
			os.Exit(1)
		}

		oldHistory := history
		history, _ = storage.LoadHistory()

		fm := finalModel.(timerModel)
		if fm.status == "startled" {
			redStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#F05522"))
			fmt.Println(redStyle.Render("\n    (=>.<=) Session canceled."))

			penaltyMsg := ""
			if oldHistory.FishCoins > history.FishCoins {
				penaltyMsg = " (-1 Fish Coin Penalty)"
			}

			fmt.Println(redStyle.Render("    You startled the cat!" + penaltyMsg))

			os.Exit(1)
		} else if fm.status == "rested" {
			_ = beeep.Notify("Meow!", "Session completed!", "")
			_ = beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
			coinsEarned := minutes / 25
			for i := range history.Sessions {
				if history.Sessions[i].ID == session.ID {
					history.Sessions[i].Status = "rested"
					history.TotalFocusMinutes += minutes
					history.FishCoins += coinsEarned
					_ = storage.SaveHistory(history)
					break
				}
			}
			greenStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#F05522"))
			fmt.Println(greenStyle.Render(fmt.Sprintf("\n    (=^ω^=) Session completed! +%d Fish Coins", coinsEarned)))
		}
	},
}



var breakCmd = &cobra.Command{
	Use:   "break [minutes]",
	Short: "Start a break (catnap). Defaults to 5 minutes.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		setupTerminal()

		minutes := 5
		if len(args) > 0 {
			parsed, err := strconv.Atoi(args[0])
			if err != nil || parsed <= 0 {
				os.Exit(1)
			}
			minutes = parsed
		}

		history, err := storage.LoadHistory()
		if err != nil {
			os.Exit(1)
		}

		if len(history.Sessions) > 0 {
			lastIdx := len(history.Sessions) - 1
			if history.Sessions[lastIdx].Status == "in_progress" {
				history.Sessions[lastIdx].Status = "startled"
				_ = storage.SaveHistory(history)
			}
		}

		now := time.Now()
		session := storage.Session{
			ID:              uuid.New().String(),
			Date:            now.Format("2006-01-02"),
			StartTimeUnix:   now.Unix(),
			DurationMinutes: minutes,
			Recipe:          "break",
			Status:          "in_progress",
		}

		history.Sessions = append(history.Sessions, session)
		if err := storage.SaveHistory(history); err != nil {
			os.Exit(1)
		}

		m := timerModel{
			session:        session,
			totalSeconds:   minutes * 60,
			remaining:      minutes * 60,
			status:         "running",
			animationFrame: 0,
			sessionType:    "break",
			equippedToy:    history.EquippedToy,
			equippedEnv:    history.EquippedEnv,
			equippedBreed:  history.EquippedBreed,
		}

		p := tea.NewProgram(m)
		finalModel, err := p.Run()
		if err != nil {
			fmt.Printf("Error running timer: %v\n", err)
			os.Exit(1)
		}

		history, _ = storage.LoadHistory()

		fm := finalModel.(timerModel)
		if fm.status == "startled" {
			redStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#F05522"))
			fmt.Println(redStyle.Render("\n    (=>.<=) Break canceled."))
			os.Exit(1)
		} else if fm.status == "rested" {
			_ = beeep.Notify("Meow!", "Break completed! Time to play!", "")
			_ = beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
			
			for i := range history.Sessions {
				if history.Sessions[i].ID == session.ID {
					history.Sessions[i].Status = "rested"
					_ = storage.SaveHistory(history)
					break
				}
			}
			
			greenStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#F05522"))
			fmt.Println(greenStyle.Render("\n    (=^ω^=) Break completed!"))
		}
	},
}

func init() {
	startCmd.Flags().StringVarP(&sessionTag, "tag", "t", "", "Tag for the session (e.g. coding)")
	startCmd.Flags().StringVarP(&sessionSound, "sound", "s", "", "Background sound to play (purr, rain, lofi)")
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(breakCmd)
}
