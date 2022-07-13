package sandbox

var (
	Ramcheck = []string{
		`	var mod = syscall.NewLazyDLL("kernel32.dll")
	var proc = mod.NewProc("GetPhysicallyInstalledSystemMemory")
	var mem uint64
	proc.Call(uintptr(unsafe.Pointer(&mem)))
	mem = mem / 1048576
	if mem < 4 {
		os.Exit(0)
	}
	//__SANDBOX__
`, `
	"os"
	"syscall"
	"unsafe"
	//__IMPORT__
`}
)
