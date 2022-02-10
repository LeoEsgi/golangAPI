package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// Nous avons utilisé une structure pour permettre de stocker nos processus
// Cette structure nous permet d'optimiser le code
type Process struct {
	Pid string
	Cwd string
	Exe string
}

func main() {
	fmt.Println(pid())
}

//
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

// inspecte le dossier lié au processus pour chercher si le dossier cwd
// existe et si oui retourne son path
func inspectPidCwd(id string) string {

	file, err := os.Readlink("/proc/" + id + "/cwd")

	if err == nil {
		path, _ := filepath.Abs(file)
		return path + "/proc/" + id + "/cwd"
	} else {
		return "no cwd"
	}

}

// inspecte le dossier lié au processus pour chercher si le fichier exe
// existe et si oui retourne son path
func inspectPidExe(id string) string {

	file, err := os.Readlink("/proc/" + id + "/exe")
	if err == nil {
		path, _ := filepath.Abs(file)
		return path + "/proc/" + id + "/exe\n"
	} else {
		return "no exe\n"
	}
}

// fonction d'affichage  avec pour chaque processus , son pid , le chemin s'il existe de cwd et exe
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
