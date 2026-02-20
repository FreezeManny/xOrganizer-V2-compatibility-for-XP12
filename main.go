package main

import (
	"os"
	"os/exec"
	"strings"
)

const (
	xp12Path      = "../X-Plane.exe"
	xp12GsFolder  = "SCENERY_PACK *GLOBAL_AIRPORTS*"
	xoDummyFolder = "SCENERY_PACK Custom Scenery/###GLOBAL SCENERY####/"
	dummyPath     = "Custom Scenery/###GLOBAL SCENERY####"
)

func main() {
	os.MkdirAll(dummyPath, 0755)

	data, _ := os.ReadFile("scenery_packs.ini")
	content := strings.ReplaceAll(string(data), xoDummyFolder, xp12GsFolder)
	os.WriteFile("scenery_packs.ini", []byte(content), 0644)
	exec.Command(xp12Path).Start()
}
