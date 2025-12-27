package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kkdai/youtube/v2"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	var ytUrl string

	if len(os.Args) < 2 {
		fmt.Print("Enter YouTube URL: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read input:", err)
			os.Exit(1)
		}

		ytUrl = strings.TrimSpace(input)

		if ytUrl == "" {
			fmt.Println("URL cannot be empty")
			os.Exit(1)
		}

	}else {
		ytUrl = os.Args[1]
	}

	client := youtube.Client{}

	video, err := client.GetVideoContext(context.Background(), ytUrl)
	if err != nil {
		fmt.Printf("Error fetching video info: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Title: %s\n", video.Title)
	fmt.Printf("Duration: %s\n", video.Duration)

	formats := video.Formats.WithAudioChannels()
	formats.Sort()
	bestAudio := formats[0]

	fmt.Printf("Downloading best audio format: %s (itag: %d)\n", bestAudio.MimeType, bestAudio.ItagNo)

	// Download the audio stream
	stream, _, err := client.GetStreamContext(context.Background(), video, &bestAudio)
	if err != nil {
		fmt.Printf("Error getting stream: %v\n", err)
		os.Exit(1)
	}
	defer stream.Close()

	tempFile := filepath.Join(os.TempDir(), "yt_audio_"+video.ID)
	outFile, err := os.Create(tempFile)
	if err != nil {
		fmt.Printf("Error creating temp file: %v\n", err)
		os.Exit(1)
	}
	defer os.Remove(tempFile)

	_, err = io.Copy(outFile, stream)
	outFile.Close()
	if err != nil {
		fmt.Printf("Error downloading audio: %v\n", err)
		os.Exit(1)
	}

	safeTitle := strings.Map(func(r rune) rune {
		if strings.ContainsRune(`<>:"/\|?*`, r) {
			return '_'
		}
		return r
	}, video.Title)


	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Failed to get home directory:", err)
		os.Exit(1)
	}

	outputDir := filepath.Join(homeDir, "Music")
	mp3Filename := safeTitle + ".mp3"

	outputPath := filepath.Join(outputDir, mp3Filename)

	// Convert to MP3 using ffmpeg
	fmt.Println("Converting to MP3...")
	cmd := exec.Command("ffmpeg", "-i", tempFile, "-q:a", "0", "-map", "a", outputPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error converting to MP3 (is ffmpeg installed?): %v\n", err)
		os.Exit(1)
	}

	fmt.Printf(
		"\033[1;32m✔ Done\033[0m %s\n",
		outputPath,
	)
}