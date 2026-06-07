package cmd

import (
	"neko/storage"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

type ShopItem struct {
	ID   string
	Name string
	Cost int
	Type string
}

var shopItems = []ShopItem{
	{"toy_yarn_ball", "Yarn Ball", 0, "toy"},
	{"toy_cardboard_box", "Cardboard Box", 100, "toy"},
	{"toy_catnip_mouse", "Catnip Mouse", 200, "toy"},
	{"toy_feather_wand", "Feather Wand", 300, "toy"},
	{"toy_laser_pointer", "Laser Pointer", 500, "toy"},
	{"toy_crinkle_tunnel", "Crinkle Tunnel", 800, "toy"},
	{"toy_scratch_post", "Scratching Post", 1000, "toy"},
	{"toy_catnip_kicker", "Catnip Kicker", 1200, "toy"},
	{"toy_puzzle_feeder", "Puzzle Feeder", 1500, "toy"},
	{"toy_auto_mouse", "Automatic Mouse", 2000, "toy"},

	{"env_floor", "Floor", 0, "env"},
	{"env_cozy_blanket", "Cozy Blanket", 200, "env"},
	{"env_window_sill", "Window Sill", 500, "env"},
	{"env_laundry_basket", "Laundry Basket", 800, "env"},
	{"env_cardboard_castle", "Cardboard Castle", 1200, "env"},
	{"env_cat_tree", "Cat Tree", 1800, "env"},
	{"env_sunbeam", "Sunbeam", 2500, "env"},
	{"env_office_chair", "Office Chair", 3000, "env"},
	{"env_keyboard", "Keyboard", 4000, "env"},
	{"env_greenhouse", "Greenhouse", 5000, "env"},

	{"breed_tuxedo", "Tuxedo", 0, "breed"},
	{"breed_tabby", "Tabby", 300, "breed"},
	{"breed_calico", "Calico", 600, "breed"},
	{"breed_siamese", "Siamese", 1000, "breed"},
	{"breed_maine_coon", "Maine Coon", 1500, "breed"},
	{"breed_russian_blue", "Russian Blue", 2200, "breed"},
	{"breed_sphynx", "Sphynx", 3000, "breed"},
	{"breed_bengal", "Bengal", 4000, "breed"},
	{"breed_persian", "Persian", 5000, "breed"},
	{"breed_galactic", "Galactic Cat", 10000, "breed"},
}

var shopCmd = &cobra.Command{
	Use:   "shop",
	Short: "View the shop and your Fish Coins",
	Run: func(cmd *cobra.Command, args []string) {
		setupTerminal()

		history, err := storage.LoadHistory()
		if err != nil {
			fmt.Println("Error loading history:", err)
			os.Exit(1)
		}

		fmt.Println()
		titleStyle := lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("#FFD700")).Bold(true)
		fmt.Println(titleStyle.Render(fmt.Sprintf("=== NEKO SHOP (Fish Coins: %d) ===", history.FishCoins)))
		fmt.Println("Use 'neko shop buy [item_id]' to purchase!")
		fmt.Println()

		purchased := make(map[string]bool)
		for _, p := range history.PurchasedItems {
			purchased[p] = true
		}

		printCategory := func(catType, title string) {
			fmt.Println(lipgloss.NewStyle().Bold(true).Underline(true).Render(title))
			for _, item := range shopItems {
				if item.Type == catType {
					status := fmt.Sprintf("[%4d FC]", item.Cost)
					if purchased[item.ID] {
						status = "[  OWNED ]"
						if history.EquippedToy == item.ID || history.EquippedEnv == item.ID || history.EquippedBreed == item.ID {
							status = "[EQUIPPED]"
						}
					}
					fmt.Printf("%-12s %-25s %s\n", status, item.Name, item.ID)
				}
			}
			fmt.Println()
		}

		printCategory("toy", "TOYS")
		printCategory("env", "ENVIRONMENTS")
		printCategory("breed", "BREEDS")
	},
}

var buyCmd = &cobra.Command{
	Use:   "buy [item_id]",
	Short: "Buy an item from the shop using Fish Coins",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		setupTerminal()
		itemID := args[0]

		history, err := storage.LoadHistory()
		if err != nil {
			fmt.Println("Error loading history:", err)
			os.Exit(1)
		}

		var targetItem *ShopItem
		for _, item := range shopItems {
			if item.ID == itemID {
				targetItem = &item
				break
			}
		}

		if targetItem == nil {
			fmt.Println("Item not found!")
			os.Exit(1)
		}

		for _, p := range history.PurchasedItems {
			if p == itemID {
				fmt.Println("You already own this item!")
				os.Exit(1)
			}
		}

		if history.FishCoins < targetItem.Cost {
			fmt.Printf("Not enough Fish Coins! You need %d more.\n", targetItem.Cost-history.FishCoins)
			os.Exit(1)
		}

		history.FishCoins -= targetItem.Cost
		history.PurchasedItems = append(history.PurchasedItems, itemID)
		_ = storage.SaveHistory(history)

		successStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))
		fmt.Println(successStyle.Render(fmt.Sprintf("Successfully purchased %s! (-%d Fish Coins)", targetItem.Name, targetItem.Cost)))
	},
}

var equipCmd = &cobra.Command{
	Use:   "equip [item_id]",
	Short: "Equip an item you own",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		setupTerminal()
		itemID := args[0]

		history, err := storage.LoadHistory()
		if err != nil {
			fmt.Println("Error loading history:", err)
			os.Exit(1)
		}

		var targetItem *ShopItem
		for _, item := range shopItems {
			if item.ID == itemID {
				targetItem = &item
				break
			}
		}

		if targetItem == nil {
			fmt.Println("Item not found in shop!")
			os.Exit(1)
		}

		owned := false
		for _, p := range history.PurchasedItems {
			if p == itemID {
				owned = true
				break
			}
		}

		if !owned {
			fmt.Println("You do not own this item! Buy it first with 'neko shop buy'.")
			os.Exit(1)
		}

		switch targetItem.Type {
		case "toy":
			history.EquippedToy = itemID
		case "env":
			history.EquippedEnv = itemID
		case "breed":
			history.EquippedBreed = itemID
		}

		_ = storage.SaveHistory(history)
		successStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))
		fmt.Println(successStyle.Render(fmt.Sprintf("Successfully equipped %s!", targetItem.Name)))
	},
}

func init() {
	shopCmd.AddCommand(buyCmd)
	shopCmd.AddCommand(equipCmd)
	rootCmd.AddCommand(shopCmd)
}
