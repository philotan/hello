package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	var (
		cmd      *exec.Cmd
		url      string = "https://s6.iimage.su/s/01/gnowGtxxaBSS5fd0QWzp9GzhbZQr7ErJbHAwtGrnI.png"
		filename string = "24c1e8ff4eeb05cda79a6c96d1181f1b.png"
	)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: %v", resp.Status)
	}
	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Image download")

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", filename)
	case "darwin":
		cmd = exec.Command("open", filename)
	case "linux":
		cmd = exec.Command("xdg-open", filename)
	default:
		fmt.Printf("Unknow OS: %v\n", runtime.GOOS)
	}

	cmd.Start()
}
