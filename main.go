package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"github.com/fsnotify/fsnotify"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error fetching home directory:", err)
	}

	screenshotDir := filepath.Join(homeDir, "Desktop")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					handleNewFile(event.Name)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()

	err = watcher.Add(screenshotDir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Watching for new screenshots in:", screenshotDir)
	select {} //keeps program running without consuming CPU
}

func handleNewFile(filePath string) {
	time.Sleep(1 * time.Second) //let's make sure the file fully saves

	fileName := filepath.Base(filePath)

	if !strings.HasPrefix(fileName, "Screenshot") || !strings.Contains(fileName, " at ") {
		return
	}

	newName := sanitizeFileName(fileName)
	newPath := filepath.Join(filepath.Dir(filePath), newName)

	if err := os.Rename(filePath, newPath); err != nil {
		log.Println("Error renaming file:", err)
	} else {
		log.Printf("Renamed: %s -> %s\n", fileName, newName)
	}
}

func sanitizeFileName(fileName string) string {
	return replaceUnicodeSpaces(fileName, '_')
}

func replaceUnicodeSpaces(str string, replacement rune) string {
	var result strings.Builder
	for _, r := range str {
		if unicode.IsSpace(r) { // This catches all Unicode space characters which for some reason Mac uses mixed spaces
			result.WriteRune(replacement)
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}
