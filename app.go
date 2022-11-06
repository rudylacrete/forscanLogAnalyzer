package main

import (
	"context"
	"fmt"

	"github.com/rudylacrete/forscanLogAnalyzer/lib"
	"github.com/rudylacrete/forscanLogAnalyzer/models"
	"github.com/rudylacrete/forscanLogAnalyzer/plugin"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type internalLogger struct {
	ctx context.Context
}

func (l *internalLogger) Write(p []byte) (n int, err error) {
	runtime.LogDebug(l.ctx, string(p))
	return len(p), nil
}

// App struct
type App struct {
	ctx    context.Context
	logger *internalLogger
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.logger = &internalLogger{ctx: ctx}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Test2() string {
	f, e := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	if e != nil {
		panic(e)
	}
	return f
}

func (a *App) LoadLogFile() (logs *models.ForscanLogs, err error) {
	p, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return
	}
	plugins := []models.Plugin{
		plugin.NewTurboPlugin(plugin.Bar),
	}
	logs, err = lib.ParseFile(p, a.logger)
	if err != nil {
		return nil, err
	}
	for _, p := range plugins {
		runtime.LogDebugf(a.ctx, "Applying plugin: %s\n", p.Info())
		plgErr := p.Transform(logs)
		if plgErr != nil {
			runtime.LogErrorf(a.ctx, "An error occured during plugin transform: %s\n", err)
		}
	}
	return logs, err
}
