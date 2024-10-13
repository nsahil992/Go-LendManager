package main

import (
	"fmt"
)

func main() {
	// The map stores friends and the items lent to them
	lent := map[string][]string{
		"Courage":  {},
		"Ben":      {"Watch"},
		"Mr.Bean":  {"Teddy"},
	}

	for {
		fmt.Println("What do you want to do? (takeback/give/newfriend/quit)")
		var userAction string
		fmt.Scan(&userAction) // Read user input for the action

		if userAction == "quit" {
			break // Exit the loop if user types "quit"
		}

		switch userAction {
		case "takeback":
			// List all friends
			fmt.Println("These are your friends:")
			for friend := range lent {
				fmt.Println(friend)
			}

			// Ask for the friend's name
			fmt.Println("Which friend did you lend to?")
			var friendName string
			fmt.Scan(&friendName)

			if items, ok := lent[friendName]; ok {
				if len(items) == 0 {
					fmt.Printf("You haven't given anything to %s.\n", friendName)
					continue
				}

				// Display the items lent to the friend
				fmt.Printf("This is what you gave to %s:\n", friendName)
				for _, item := range items {
					fmt.Println(item)
				}

				// Ask which item to take back
				fmt.Printf("What did you take back from %s?\n", friendName)
				var itemName string
				fmt.Scan(&itemName)

				// Search for the item in the list
				itemIndex := -1
				for i, item := range lent[friendName] {
					if item == itemName {
						itemIndex = i
						break
					}
				}

				if itemIndex == -1 {
					fmt.Println("Sorry, I didn't find that item.")
				} else {
					// Remove the item from the list
					lent[friendName] = append(lent[friendName][:itemIndex], lent[friendName][itemIndex+1:]...)
					fmt.Printf("Alright, I'll remember that you took %s from %s.\n", itemName, friendName)
				}
			} else {
				fmt.Println("Sorry, I didn't find that friend.")
			}

		case "give":
			// List all friends
			fmt.Println("These are your friends:")
			for friend := range lent {
				fmt.Println(friend)
			}

			// Ask for the friend's name
			fmt.Println("Which friend do you want to lend to?")
			var friendName string
			fmt.Scan(&friendName)

			if _, ok := lent[friendName]; ok {
				// Ask for the item to lend
				fmt.Printf("What do you want to lend to %s?\n", friendName)
				var itemName string
				fmt.Scan(&itemName)

				// Add the item to the friend's list
				lent[friendName] = append(lent[friendName], itemName)
				fmt.Printf("Got it! You lent %s to %s.\n", itemName, friendName)
			} else {
				fmt.Println("Sorry, I didn't find that friend.")
			}

		case "newfriend":
			// Add a new friend
			fmt.Println("Who is your new friend?")
			var friendName string
			fmt.Scan(&friendName)
			lent[friendName] = []string{}
			fmt.Printf("Added %s as a new friend.\n", friendName)

		default:
			fmt.Println("Sorry, I didn't understand that. (Valid choices: give/takeback/newfriend/quit)")
		}
	}
	fmt.Println("Goodbye!")
}
