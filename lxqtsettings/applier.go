package lxqtsettings

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type Applier struct {
	configDir string
}

func NewApplier() *Applier {
	return &Applier{
		configDir: filepath.Join(os.Getenv("HOME"), ".config", "lxqt"),
	}
}

func (a *Applier) Apply() {
	a.applyStyle()
	a.applyFont()
	a.applyIconTheme()
	a.applyPalette()
	a.applyCursor()
}

func (a *Applier) applyStyle() {
	content, err := ioutil.ReadFile(filepath.Join(a.configDir, "session.conf"))
	if err == nil && strings.Contains(string(content), "style=Fusion") {
		widgets.QApplication_SetStyle2("Fusion")
	}
}

func (a *Applier) applyFont() {
	font := gui.NewQFont2("Sans Serif", 10, int(gui.QFont__Normal), false)
	widgets.QApplication_SetFont(font, "")
}

func (a *Applier) applyIconTheme() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	//core.QCoreApplication_SetAttribute(core.Qt__AAUseHighDpiPixmaps, true)
	gui.QIcon_SetThemeName("Papirus")
}

func (a *Applier) applyPalette() {
	pal := gui.NewQPalette()
	pal.SetColor2(gui.QPalette__Window, gui.NewQColor3(245, 245, 245, 1))
	widgets.QApplication_SetPalette(pal, "")
}

func (a *Applier) applyCursor() {
	cursor := gui.NewQCursor2(core.Qt__ArrowCursor)
	gui.QGuiApplication_SetOverrideCursor(cursor)
}

//func (a *Applier) Watch() {
//	log.Println("Watching LXQt settings for changes... (not implemented: requires inotify or polling)")
//}
