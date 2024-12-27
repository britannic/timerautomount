package main

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

func main() {
	path := `Software\Microsoft\OneDrive\Accounts\Business1`
	keyName := "Timerautomount"
	value := uint64(1)

	key, err := registry.OpenKey(registry.CURRENT_USER, path, registry.ALL_ACCESS)
	if err != nil {
		// If the key doesn't exist, create it
		key, _, err = registry.CreateKey(registry.CURRENT_USER, path, registry.ALL_ACCESS)
		if err != nil {
			fmt.Printf("Error creating registry key: %v\n", err)
			return
		}
	}
	defer key.Close()

	err = key.SetQWordValue(keyName, value)
	if err != nil {
		fmt.Printf("Error setting registry value: %v\n", err)
		return
	}

	fmt.Println("Registry entry added successfully")
}
