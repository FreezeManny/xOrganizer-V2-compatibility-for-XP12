package main

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
	"unsafe"
)

const (
	xp12Path      = "../X-Plane.exe"
	xp12GsFolder  = "SCENERY_PACK *GLOBAL_AIRPORTS*"
	xoDummyFolder = "SCENERY_PACK Custom Scenery/###GLOBAL SCENERY####/"
	dummyPath     = "Custom Scenery/###GLOBAL SCENERY####"
)

var (
	user32      = syscall.NewLazyDLL("user32.dll")
	messageBoxW = user32.NewProc("MessageBoxW")
)

func fatal(msg string) {
	title, _ := syscall.UTF16PtrFromString("xOV2 Error")
	text, _ := syscall.UTF16PtrFromString(msg)
	messageBoxW.Call(0, uintptr(unsafe.Pointer(text)), uintptr(unsafe.Pointer(title)), 0x10)
	os.Exit(1)
}

func main() {
    if err := os.MkdirAll(dummyPath, 0755); err != nil {
        fatal("Failed to create dummy dir:\n" + err.Error())
    }

    data, err := os.ReadFile("scenery_packs.ini")
    if err != nil {
        fatal("Failed to read scenery_packs.ini:\n" + err.Error())
    }

    original := string(data)
    content := strings.ReplaceAll(original, xoDummyFolder, xp12GsFolder)
    if content == original {
        fatal("Dummy folder entry not found in scenery_packs.ini.\nHas xOrganizer V2 been run yet?")
    }

    if err := os.WriteFile("scenery_packs.ini", []byte(content), 0644); err != nil {
        fatal("Failed to write scenery_packs.ini:\n" + err.Error())
    }

    if err := exec.Command(xp12Path).Start(); err != nil {
        fatal("Failed to launch X-Plane:\n" + err.Error())
    }
}
