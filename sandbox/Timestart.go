package sandbox

var (
	Timestart = []string{`t := time.Now()
	hour := t.Hour()
	minute := t.Minute()
	day := t.Day()
	pass := strconv.Itoa(hour) + strconv.Itoa(minute) + strconv.Itoa(day)
	if len(os.Args) < 2 || os.Args[1] != pass{
		os.Exit(0)
	}
	//__SANDBOX__
`, `
	"time"
	"strconv"
	"os"
	//__IMPORT__`}
)
