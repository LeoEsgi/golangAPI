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
	(pid())
}

func pid() ([]Process, error) {
	var process []Process
	f, err := os.Open("/proc")
	if err != nil {
		return process, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return process, err
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

	return process, nil

}

func inspectPidCwd(id string) string {

	file, err := os.Readlink("/proc/" + id + "/cwd")

	path, err := filepath.Abs(file)
	fmt.Println(path)

	if err == nil {
		return "/proc/" + id + "/cwd"
	} else {
		return "no cwd"
	}

}
func inspectPidExe(id string) string {

	_, err := os.Readlink("/proc/" + id + "/exe")
	if err == nil {
		return "/proc/" + id + "/exe\n"
	} else {
		return "no exe\n"
	}
}
