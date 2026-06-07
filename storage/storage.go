package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Session represents a single nekoing session.
type Session struct {
	ID              string `json:"id"`
	Date            string `json:"date"`
	StartTimeUnix   int64  `json:"start_time_unix"`
	DurationMinutes int    `json:"duration_minutes"`
	Recipe          string `json:"recipe"`
	Tag             string `json:"tag,omitempty"`
	Status          string `json:"status"` // "rested", "startled", "in_progress"
}

// History holds the entire state of the user's sessions and achievements.
type History struct {
	TotalFocusMinutes int       `json:"total_focus_minutes"`
	FishCoins         int       `json:"fish_coins"`
	EquippedToy       string    `json:"equipped_toy"`
	EquippedEnv       string    `json:"equipped_env"`
	EquippedBreed     string    `json:"equipped_breed"`
	PurchasedItems    []string  `json:"purchased_items"`
	UnlockedRecipes   []string  `json:"unlocked_recipes"`
	Sessions          []Session `json:"sessions"`
}

func getHistoryFilePath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(configDir, "neko")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", err
	}
	return filepath.Join(dir, "history.json"), nil
}

// LoadHistory reads the history from the user's config directory.
func LoadHistory() (*History, error) {
	path, err := getHistoryFilePath()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &History{
				TotalFocusMinutes: 0,
				FishCoins:         0,
				EquippedToy:       "toy_yarn_ball",
				EquippedEnv:       "env_floor",
				EquippedBreed:     "breed_tuxedo",
				PurchasedItems:    []string{"toy_yarn_ball", "env_floor", "breed_tuxedo"},
				UnlockedRecipes:   []string{},
				Sessions:          []Session{},
			}, nil
		}
		return nil, err
	}
	defer f.Close()

	var history History
	if err := json.NewDecoder(f).Decode(&history); err != nil {
		return nil, err
	}

	if history.PurchasedItems == nil {
		history.PurchasedItems = []string{"toy_yarn_ball", "env_floor", "breed_tuxedo"}
		
		// Migrate old unlocks
		for _, r := range history.UnlockedRecipes {
			if r == "french_press" {
				history.PurchasedItems = append(history.PurchasedItems, "toy_cardboard_box")
			}
			if r == "matcha" {
				history.PurchasedItems = append(history.PurchasedItems, "toy_catnip_mouse")
			}
		}
	}
	if history.EquippedToy == "" {
		history.EquippedToy = "toy_yarn_ball"
	}
	if history.EquippedEnv == "" {
		history.EquippedEnv = "env_floor"
	}
	if history.EquippedBreed == "" {
		history.EquippedBreed = "breed_tuxedo"
	}
	if history.UnlockedRecipes == nil {
		history.UnlockedRecipes = []string{}
	}
	if history.Sessions == nil {
		history.Sessions = []Session{}
	}

	return &history, nil
}

// SaveHistory writes the history back to the user's config directory.
func SaveHistory(h *History) error {
	path, err := getHistoryFilePath()
	if err != nil {
		return err
	}

	if len(h.Sessions) > 1000 {
		h.Sessions = h.Sessions[len(h.Sessions)-1000:]
	}

	data, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return err
	}

	tempFile := path + ".tmp"
	err = func() error {
		f, err := os.OpenFile(tempFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := f.Write(data); err != nil {
			return err
		}
		return nil
	}()
	if err != nil {
		return err
	}

	return os.Rename(tempFile, path)
}
