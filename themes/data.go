package themes

import (
"fyne.io/fyne/v2"
)

// Tutorial defines the data structure for a tutorial
type Tutorial struct {
	Title string
	View         func(w fyne.Window) fyne.CanvasObject
}

var (
	// Tutorials defines the metadata for each tutorial
	Tutorials = map[string]Tutorial{
		"welcome": {"Welcome", welcomeScreen},
		"Bypass": {"Bypass",
			BypassAV,
		},
		//"BypassAV":{"BypassAV",BypassAV},
		//"WebShell":{"WebShell",Bypassws},
	}

	// TutorialIndex  defines how our tutorials should be laid out in the index tree
	TutorialIndex = map[string][]string{
		"":            {"welcome", "Bypass"},
		//"Bypass": {"BypassAV", "WebShell"},
	}
)
