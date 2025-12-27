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
