package utils

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func (ps *Problemset) MakeProblemset(pagehtml string) {
	scanner := bufio.NewScanner(strings.NewReader(pagehtml))
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), fmt.Sprintf("id=\"%s", ps.Chategory)) {
			var lines []string

			shouldscan := true
			for ; shouldscan; scanner.Scan() {
				lines = append(lines, scanner.Text())
				if strings.Contains(scanner.Text(), "&nbsp;pont") {
					shouldscan = false
				}
			}

			task := Task{}
			task.Id = strings.Split(lines[0], "\"")[1]
			lastline := lines[len(lines)-1]
			task.Points, _ = strconv.Atoi(string(lastline[strings.Index(lastline, "(")+1]))

			lastdesclineidx := len(lines) - 2
			if lastdesclineidx >= 0 {
				ln := lines[lastdesclineidx]
				if strings.Contains(ln, "align=\"right\"") {
					task.Author = ln[(strings.Index(ln, "<i>") + 3):strings.Index(ln, "</i>")]
					lastdesclineidx--
				}
			}

			task.Description = "<p align=\"justify\">"
			task.Description += lines[0][strings.Index(lines[0], "</b>")+4:]
			for i := 1; i <= lastdesclineidx; i++ {
				task.Description += lines[i] + "\n"
			}

			ps.Tasks = append(ps.Tasks, task)
		}
	}
}
