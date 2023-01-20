package main

import (
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	log "github.com/sirupsen/logrus"
)

var l *log.Logger

func init() {
	l = log.New()
	l.SetFormatter(&log.TextFormatter{
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
	})
}

func main() {
	// Create astilectron app
	a, err := astilectron.New(l, astilectron.Options{
		AppName:           "Go Electron App",
		BaseDirectoryPath: ".go-electron-data",
		VersionElectron:   "17.0.0",
	})
	if err != nil {
		l.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	}
	defer a.Close()

	// Handle signals
	a.HandleSignals()

	// Start
	if err = a.Start(); err != nil {
		l.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
	}

	// New window
	var w *astilectron.Window
	if w, err = a.NewWindow("web/index.html", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(700),
		Width:  astikit.IntPtr(700),
	}); err != nil {
		l.Fatal(fmt.Errorf("main: new window failed: %w", err))
	}

	// Create windows
	if err = w.Create(); err != nil {
		l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}

	// Blocking pattern
	a.Wait()
}
