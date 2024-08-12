package main

import (
	"fmt"
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/ledongthuc/pdf"
	"io"
	"os"
	"strings"
)

func main() {
	pdf.DebugOn = true
	// Initialize PDF reader
	_, reader, err := pdf.Open("/home/quynguyen3/Desktop/OReilly.Concurrency.in.Go.2017.8.pdf")
	if err != nil {
		fmt.Println("Error reading PDF file:", err)
		return
	}

	// Initialize text-to-speech engine
	speech := htgotts.Speech{Folder: "audio", Language: "en"}

	var text string
	// Iterate over each page in the PDF
	totalPage := reader.NumPage()
	for i := 1; i <= totalPage; i++ {
		page := reader.Page(i)
		if page.V.IsNull() {
			continue
		}
		pageText, err := page.GetPlainText(nil)
		if err != nil {
			fmt.Println("Error extracting text from page:", err)
			continue
		}
		text += pageText
	}

	// Save the extracted text to an audio file
	err = speech.Speak(text)
	if err != nil {
		fmt.Println("Error generating speech:", err)
		return
	}

	// Save text to file (if needed)
	err = saveToFile("story.txt", text)
	if err != nil {
		fmt.Println("Error saving text to file:", err)
		return
	}

	fmt.Println("Text-to-speech conversion completed and saved to 'story.mp3'")
}

// saveToFile saves the text to a file
func saveToFile(filename, text string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, strings.NewReader(text))
	return err
}
