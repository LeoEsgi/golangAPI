package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type Process struct {
	Pid string
	Cwd string
	Exe string
}

func main() {
	fmt.Println(pid())
}

func pid() string {
	var process []Process
	f, err := os.Open("/proc")
	if err != nil {
		return "Error 500 , you dont have directory /proc in your linux architecture"
	}

	// Récupère les sous dossier du dossier ./proc
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return "Error 500, permission denied, cant open /proc"
	}

	for _, file := range fileInfo {

		if _, err := strconv.Atoi(file.Name()); err == nil {
			var p Process
			p.Pid = file.Name()
			p.Cwd = inspectPidCwd(file.Name())
			p.Exe = inspectPidExe(file.Name())
			process = append(process, p)

		}
	}

	return display(process)

}

func inspectPidCwd(id string) string {

	file, err := os.Readlink("/proc/" + id + "/cwd")

	if err == nil {
		path, _ := filepath.Abs(file)
		return path + "/proc/" + id + "/cwd"
	} else {
		return "no cwd"
	}

}
func inspectPidExe(id string) string {

	file, err := os.Readlink("/proc/" + id + "/exe")
	if err == nil {
		path, _ := filepath.Abs(file)
		return path + "/proc/" + id + "/exe\n"
	} else {
		return "no exe\n"
	}
}

func display(process []Process) string {
	var show string
	for i := 0; i < len(process); i++ {
		pid := "PID : " + process[i].Pid
		cwd := "CWD : " + process[i].Cwd
		exe := "EXE : " + process[i].Exe
		show = show + pid + " |  " + cwd + "  | " + exe + "\n"

	}
	return show

}
