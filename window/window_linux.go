// +build linux

package window

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetWindowTitles() []string {
	out, err := exec.Command("wmctrl", "-l").Output()
	if err != nil {
		fmt.Println(err)
	}
	s := strings.Split(string(out), "\n")

	info, err := exec.Command("uname", "-a").Output()
	dist := strings.Split(string(info), " ")[1]
	titles := make([]string, 0)
	for _, v := range s {
		if strings.Contains(v, dist) {
			titles = append(titles, strings.Split(v, dist + " ")[1])
		}
	}
	return titles
}