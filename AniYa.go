package main

import (
	"AniYa/core"
	"AniYa/themes"
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/dialog"
	"github.com/flopp/go-findfont"
	"os"
	"strings"
)

const preferenceCurrentTutorial = "currentTutorial"
const Version = "1.2.0"

var topWindow fyne.Window

//设置中文
func init() {
	//os.RemoveAll(core.TEMP_DIR)
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		// 微软雅黑-常规
		//println(path)
		if strings.Contains(path, "msyh.ttc") {
			//println(path)
			os.Setenv("FYNE_FONT", path)
			println(os.Getenv("FYNE_FONT"))
			println("设置中文成功")
			break
			//兼容win7中文
		} else if strings.Contains(path, "msyh.ttf") {
			os.Setenv("FYNE_FONT", path)
			println(os.Getenv("FYNE_FONT"))
			println("设置中文成功")
			break
		}
	}
	//println("设置中文失败")
}

func main() {
	defer os.Unsetenv("FYNE_FONT")
	//退出时
	defer os.RemoveAll(core.TempDir)
	a := app.NewWithID("io.fyne.demo")
	a.SetIcon(themes.Resource2Png)
	//a.SetIcon(theme.FyneLogo())
	//a := app.New() //新建一个应用
	w := a.NewWindow("AniYa") //新建一个窗口
	settingsItem := fyne.NewMenuItem("Settings", func() {
		w := a.NewWindow("Fyne Settings")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	})
	version := fyne.NewMenuItem("VERSION", func() {
		v1 := dialog.NewInformation("Version", Version, w)
		v1.Show()
	})
	author := fyne.NewMenuItem("Author", func() {
		//author := a.NewWindow("piiperxyz")
		v2 := dialog.NewInformation("Author", "piiperxyz\nhttps://github.com/piiperxyz/AniYa", w)
		v2.Show()
	})
	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu("FILE", settingsItem),
		fyne.NewMenu("ABOUT", version, author),
	)
	tmp := themes.BypassAV(w)

	w.SetContent(tmp)
	w.SetMainMenu(mainMenu)
	w.SetMaster()
	w.Resize(fyne.NewSize(800, 700))
	w.ShowAndRun() //显示窗口并运行，后续的窗口只能用show
	w.Show()       //显示窗口并运行，后续的窗口只能用show
}
