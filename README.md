# HLS Video/Audio Uploader & Player

## Overview
This project allows users to upload video/audio files, convert them to HLS format using **FFmpeg**, and stream them through an HTML5 video player.

## Features
- Upload video/audio files via a web form.
- Convert uploaded files to HLS format using **FFmpeg**.
- Store converted files in the **hls/** directory.
- Automatically update the video player to play the latest uploaded video.
- Serve HLS video files over HTTP.

## Requirements
- Go (Golang)
- FFmpeg (Must be installed and accessible from the command line)
- A web browser that supports HLS (or HLS.js for broader support)

## Installation
1. **Clone the repository**
   ```sh
   git clone https://github.com/your-username/audio-video-go.git
   cd audio-video-go
   ```

2. **Install FFmpeg**
   - Windows: Download from [ffmpeg.org](https://ffmpeg.org/) and add it to your system PATH.
   - macOS: Install via Homebrew:
     ```sh
     brew install ffmpeg
     ```
   - Linux: Install via package manager:
     ```sh
     sudo apt update && sudo apt install ffmpeg
     ```

3. **Run the server**
   ```sh
   go run main.go
   ```
   The server will start on **http://localhost:8000**

## Usage
1. **Upload a file**
   - Open [http://localhost:8000](http://localhost:8000) in a browser.
   - Select a video/audio file and click **Upload & Convert**.
   - The server will process the file and convert it to HLS format.

2. **Stream the latest uploaded file**
   - The webpage will automatically play the latest uploaded video.
   - The HLS stream will be served from `/hls/{filename}/index.m3u8`.

## Project Structure
```
/audio-video-go
│-- main.go           # Entry point for the server
│-- handler.go        # Handles file uploads and HLS conversion
│-- templates/
│   └── index.html    # Frontend for uploading and playing videos
     └── styles.css
│-- assets/           # Stores uploaded files
│-- hls/              # Stores converted HLS video segments
│-- README.md         # Project documentation
```

## API Endpoints
| Endpoint           | Method | Description                    |
|-------------------|--------|--------------------------------|
| `/`               | GET    | Renders the upload & playback page |
| `/upload`         | POST   | Uploads and converts a file to HLS |
| `/hls/{filename}/index.m3u8` | GET | Serves the HLS playlist for playback |
| `/latest-video`   | GET    | Returns the latest uploaded video filename |

## Dependencies
- **net/http** (Go’s built-in HTTP server)
- **html/template** (For rendering the HTML page)
- **os/exec** (For executing FFmpeg commands)

## Troubleshooting
### Video is not playing
- Ensure **FFmpeg** is installed and accessible from the command line.
- Check if the **hls/** directory contains the converted `.m3u8` and `.ts` files.
- Make sure your browser supports HLS (use **HLS.js** for better support).

### HTTP 405 Error on Upload
- Ensure the form action is correctly pointing to `http://localhost:8000/upload`.
- Check if the **Go server is running** (`go run main.go`).

## Future Enhancements
- Support for multiple video formats.
- Display a list of previously uploaded videos.
- Add progress bars for file uploads and conversions.

## License
This project is open-source and available under the **MIT License**.

---

Happy Coding!

