package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/VadimOnix/mbox-splitter/src/shared/utils"
)

const maxEmailsPerFile = 1000 // Number of emails per file
const bufferSize = 4096       // Buffer size for writing

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path to mbox file>")
		return
	}

	mboxFile := os.Args[1]

	// Open the source mbox file
	file, err := os.Open(mboxFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var emailCount int
	var fileCount int
	var firstEmailDate string
	var wg sync.WaitGroup
	linesChannel := make(chan []string, 10) // Channel for asynchronous writing

	// Goroutine for asynchronous writing
	go func() {
		for lines := range linesChannel {
			wg.Add(1)
			go func(lines []string, count int, date string) {
				defer wg.Done()

				// Format the file name based on the date of the first email
				outputFileName := fmt.Sprintf("mbox_part_%d_%s.mbox", count, date) // Use fmt.Sprintf for cleaner formatting
				outputFile, err := os.Create(outputFileName)
				if err != nil {
					fmt.Printf("Error creating file %s: %v\n", outputFileName, err)
					return
				}
				defer outputFile.Close()

				// Buffered writing to file
				writer := bufio.NewWriterSize(outputFile, bufferSize)
				for _, line := range lines {
					_, err := writer.WriteString(line + "\n")
					if err != nil {
						fmt.Printf("Error writing to file: %v\n", err)
						return
					}
				}

				writer.Flush() // Flush the buffer
				fmt.Printf("File %s written successfully.\n", outputFileName)
			}(lines, fileCount, firstEmailDate)
			fileCount++ // Increment fileCount *after* sending to goroutine
		}
	}()

	var currentLines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			// Check for EOF more idiomatically
			if err != io.EOF {
				fmt.Printf("Error reading file: %v\n", err)
			}
			break // Exit loop on EOF or other error
		}

		// If a new email starts
		if strings.HasPrefix(line, "From ") {
			// If the email limit per file is reached, send the lines for writing
			if emailCount >= maxEmailsPerFile {
				linesChannel <- currentLines // Send lines for writing
				currentLines = []string{}    // Clear the current line buffer. Important to re-initialize, not just set to nil
				emailCount = 0

			}

			// If this is the first email in a new file, extract its date
			if emailCount == 0 {
				firstEmailDate = utils.ExtractDateFromLine(line)
				if firstEmailDate == "" {
					firstEmailDate = time.Now().Format("2006_01_02") // If the date is not found, use the current date
				}
			}

			emailCount++
		}

		currentLines = append(currentLines, strings.TrimRight(line, "\n"))
	}

	// Send remaining lines for writing
	if len(currentLines) > 0 {
		linesChannel <- currentLines
	}

	close(linesChannel) // Close the channel after sending all data

	wg.Wait() // Wait for all write operations to complete
	fmt.Println("Finished!")
}
