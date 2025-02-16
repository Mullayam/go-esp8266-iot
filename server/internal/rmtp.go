package routes

import (
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
	"sync"

	"github.com/torresjeff/rtmp"
	"go.uber.org/zap"
)

type Camera struct {
	IP        string `json:"ip"`
	StreamKey string `json:"stream_key"`
}

var (
	mutex   sync.Mutex
	streams = make(map[string]Camera)
)

func StartRTMPServer() {
	listener, err := net.Listen("tcp", ":1935")
	if err != nil {
		log.Fatalf("Failed to start RTMP server: %v", err)
	}

	fmt.Println("RTMP server started on rtmp://192.168.29.57:1935/live")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept RTMP connection:", err)
			continue
		}

		go handleRTMPConnection(conn)
	}
}
func handleRTMPConnection(conn net.Conn) {
	defer conn.Close()
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	server := &rtmp.Server{
		Logger:      logger,
		Broadcaster: rtmp.NewBroadcaster(rtmp.NewInMemoryContext()),
	}

	err := server.Listen()
	if err != nil {
		log.Println("RTMP connection error:", err)
	}

	io.Copy(io.Discard, conn)
}

func startRTMPStream(streamKey string) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := streams[streamKey]; exists {
		log.Println("Stream already exists:", streamKey)
		return
	}
	hlsOutput := fmt.Sprintf("hls/%s.m3u8", streamKey)
	// cmd := exec.Command("ffmpeg",
	// 	"-i", fmt.Sprintf("rtmp://localhost:1935/live/%s", streamKey),
	// 	"-c:v", "libx264", "-preset", "ultrafast",
	// 	"-b:v", "500k", "-f", "hls",
	// 	"-hls_time", "2", "-hls_list_size", "5",
	// 	"-hls_flags", "delete_segments",
	// 	hlsOutput,
	// )
	cmd := exec.Command("ffmpeg",
		"-i", fmt.Sprintf("rtmp://localhost:1935/live/%s", streamKey), // Input from RTMP
		"-c:v", "libx264", // Use H.264 for better compression
		"-preset", "fast", // Balances quality and CPU usage (Options: ultrafast, superfast, fast, medium, slow, veryslow)
		"-b:v", "2500k", // Video bitrate (2.5 Mbps)
		"-maxrate", "3000k", // Max bitrate to prevent spikes
		"-bufsize", "6000k", // Buffer size
		"-r", "30", // Frame rate (FPS)
		"-s", "1280x720", // Resolution (720p)
		"-pix_fmt", "yuv420p", // Pixel format for compatibility
		"-g", "50", // GOP (Group of Pictures) for better compression
		"-keyint_min", "50", // Min keyframe interval
		"-sc_threshold", "0", // Disable scene change detection
		"-c:a", "aac", // Audio codec
		"-b:a", "128k", // Audio bitrate
		"-ac", "2", // Stereo audio
		"-f", "hls", // Output format
		"-hls_time", "2", // Segment duration (2s per .ts file)
		"-hls_list_size", "10", // Keep last 10 segments in the playlist
		"-hls_flags", "delete_segments", // Delete old segments to save space
		hlsOutput,
	)
	// 	Reduce Latency for Real-Time Streaming
	// Set -tune zerolatency
	// Lower GOP size (-g 30)
	// Reduce buffer (-bufsize 2000k)
	// Shorten HLS segment (-hls_time 1)
	// Example:

	// go
	// Copy
	// Edit
	// "-tune", "zerolatency",
	// "-g", "30",
	// "-bufsize", "2000k",
	// "-hls_time", "1",
	if err := cmd.Start(); err != nil {
		log.Println("FFmpeg error:", err)
	}
}
