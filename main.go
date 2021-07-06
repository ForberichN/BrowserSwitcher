package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func main() {
	if os.Args[1] == "install" {
		installProtocoll()
		return
	}

	parts := strings.SplitN(os.Args[1], ":", 3)

	if !validateURL(parts[2]) {
		fmt.Println("URL Failed security Check")
		return
	}

	targetBrowser := strings.ToLower(parts[1])

	if targetBrowser == "chrome" {
		openChrome(parts[2])
	}

	if targetBrowser == "firefox" {
		openFirefox(parts[2])
	}

	if targetBrowser == "edge" {
		openEdge(parts[2])
	}

	if targetBrowser == "opera" {
		openOpera(parts[2])
	}
	fmt.Println("Process Complete")
}

func installProtocoll() {
	//Registering Protocoll
	{
		newProtocollKey, _, err := registry.CreateKey(registry.CLASSES_ROOT, `browser`, registry.WRITE)
		if err != nil {
			fmt.Println("Failed to Create Regirsty Key: " + err.Error())
			return
		}
		defer newProtocollKey.Close()
		newProtocollKey.SetStringValue("URL Protocol", "")
	}

	//Linking Executable to Protocoll
	{
		newProtocollKey, _, err := registry.CreateKey(registry.CLASSES_ROOT, `browser\shell\open\command`, registry.WRITE)
		if err != nil {
			fmt.Println("Failed to Create Regirsty Key: " + err.Error())
			return
		}
		defer newProtocollKey.Close()
		newProtocollKey.SetStringValue("", os.Args[0]+" %1")
	}
}

func validateURL(url string) bool {
	success, err := regexp.MatchString(`(https?:\/\/)`, url)
	if err != nil {
		fmt.Println("Regex Match resulted in an error: " + err.Error())
		return false
	}
	return success
}

func openBrowser(executable string, url string) {
	cmd := exec.Command(strings.Replace(executable, "\"", "", -1), url)
	fmt.Println(cmd.Path)
	for _, arg := range cmd.Args {
		fmt.Println(arg)
	}
	err := cmd.Start()
	if err != nil {
		fmt.Println("Failed to start process: " + err.Error())
	}
}

func getBrowserPath(regKeyPath string) (string, error) {
	regKey, err := registry.OpenKey(registry.LOCAL_MACHINE, regKeyPath, registry.QUERY_VALUE)
	if err != nil {
		fmt.Println("Failed to open chrome Registry Key!: " + err.Error())
		return "", err
	}
	defer regKey.Close()
	browserExecutable, _, err := regKey.GetStringValue("")

	if err != nil {
		fmt.Println("Failed to read chrome Registry Key!: " + err.Error())
		return "", err
	}

	return browserExecutable, nil
}

func openFirefox(url string) {
	firefoxExecutable, err := getBrowserPath(`SOFTWARE\Microsoft\Windows\CurrentVersion\App Paths\firefox.exe`)

	if err != nil {
		fmt.Println("Could not get Firefox Browser Path: " + err.Error())
		return
	}

	openBrowser(firefoxExecutable, url)
}

func openChrome(url string) {
	chromeExecutable, err := getBrowserPath(`SOFTWARE\Microsoft\Windows\CurrentVersion\App Paths\chrome.exe`)

	if err != nil {
		fmt.Println("Could not get Chrome Browser Path: " + err.Error())
		return
	}

	openBrowser(chromeExecutable, url)
}

func openOpera(url string) {
	operaExecutable, err := getBrowserPath(`SOFTWARE\Microsoft\Windows\CurrentVersion\App Paths\opera.exe`)

	if err != nil {
		fmt.Println("Could not get Opera Browser Path: " + err.Error())
		return
	}
	openBrowser(operaExecutable, "--remote "+url)
}

func openEdge(url string) {
	edgeExecutable, err := getBrowserPath(`SOFTWARE\Microsoft\Windows\CurrentVersion\App Paths\msedge.exe`)

	if err != nil {
		fmt.Println("Could not get Edge Browser Path: " + err.Error())
		return
	}

	openBrowser(edgeExecutable, url)
}
