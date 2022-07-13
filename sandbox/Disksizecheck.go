package sandbox

var (
	Disksizecheck = []string{`
var (
	kernel322 = syscall.NewLazyDLL("kernel32.dll")
)
	minDiskSizeGB := float32(60)
var (
	getDiskFreeSpaceEx = kernel322.NewProc("GetDiskFreeSpaceExW")
	lpFreeBytesAvailable, lpTotalNumberOfBytes, lpTotalNumberOfFreeBytes int64
)

getDiskFreeSpaceEx.Call(
uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("C:"))),
uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)))

diskSizeGB := float32(lpTotalNumberOfBytes) / 1073741824
//fmt.Println(diskSizeGB)
if diskSizeGB < minDiskSizeGB {
	os.Exit(0)
}
	//__SANDBOX__
`, `
	"syscall"
	"unsafe"
	"os"
	//__IMPORT__`}
)
