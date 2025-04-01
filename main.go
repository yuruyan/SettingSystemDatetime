package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/beevik/ntp"
)

func SetWindowsSystemTime(t time.Time) {
	dateCmd := exec.Command("cmd", "/C", "echo", t.Format("2006-01-02"), "|", "date")
	timeCmd := exec.Command("cmd", "/C", "echo", t.Format("15:04:05"), "|", "time")
	dateCmd.Stdout = os.Stdout
	timeCmd.Stdout = os.Stdout
	dateCmd.Stderr = os.Stderr
	timeCmd.Stderr = os.Stderr
	err := dateCmd.Run()
	if err != nil {
		panic(fmt.Errorf("setting date failed: %v", err))
	}
	err = timeCmd.Run()
	if err != nil {
		panic(fmt.Errorf("setting time failed: %v", err))
	}
}

func SetLinuxSystemTime(t time.Time) {
	cmd := exec.Command("date", "-s", t.Format("20060102 15:04:05"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(fmt.Errorf("setting system time failed: %v", err.Error()))
	}
}

func SetSystemTime(t time.Time) {
	switch runtime.GOOS {
	case "windows":
		SetWindowsSystemTime(t)
	case "linux":
		SetLinuxSystemTime(t)
	default:
		panic("unsupported operating system")
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Program terminated: %v\n", err)
		}
	}()

	server := "ntp.aliyun.com"
	if len(os.Args) > 1 {
		server = os.Args[1]
	}
	fmt.Printf("Using NTP Server %s\n", server)
	stime, err := ntp.Time(server)
	if err != nil {
		log.Fatalf("getting ntp from server %s failed: %v", server, err)
	}
	fmt.Printf("NTP Server time: %s\n", stime.Format(time.RFC3339))
	SetSystemTime(stime)
	fmt.Printf("system time has been set to %s\n", stime.Format(time.RFC3339))
}
