package sandbox

var (
	Cpunumber = []string{
		`	a := runtime.NumCPU()
	if a < 4 {
		os.Exit(0)
	}
	//__SANDBOX__
`, `
	"os"
	"runtime"
	//__IMPORT__`}
)
