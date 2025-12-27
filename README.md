# yt-mp3

A simple Go CLI tool to download a video’s audio and convert it to an MP3 file for personal listening.

This tool runs locally, prompts for input when needed, and saves the MP3 file to your Music directory.

---

## Features

- Run as a global command (`yt-mp3`)
- Prompt for URL if not provided as an argument
- Converts audio to MP3 using `ffmpeg`
- Automatically saves output to `~/Music`

---

## Requirements

- Go 1.20+
- `ffmpeg` installed and available in `PATH`

### Install ffmpeg

**Fedora**

```bash
sudo dnf install ffmpeg
```

### Install using Go (recommended)

```bash
go install github.com/vinh0809it/yt-mp3-downloader@latest
```

Make sure Go’s bin directory is in your `PATH`:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
source ~/.bashrc
```

### Usage

Run the command:

```bash
yt-mp3
```

You will be prompted to enter a YouTube URL:

```bash
Enter YouTube URL: https://www.youtube.com/watch?v=<youtube-url>
```
