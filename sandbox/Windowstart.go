package sandbox

var (
	Windowstart = []string{
		`	var kernel = syscall.NewLazyDLL("Kernel32.dll")
	GetTickCount := kernel.NewProc("GetTickCount")
	r, _, _ := GetTickCount.Call()
	if r == 0 {
		os.Exit(0)
	}
	ms := time.Duration(r * 1000 * 1000)
	tm := time.Duration(30 * time.Minute)
	if ms < tm {
		os.Exit(0)
	}
	//__SANDBOX__
`, `"os"
	"syscall"
	"time"
	//__IMPORT__`}
)
