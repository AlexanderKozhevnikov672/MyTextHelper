package connector

import (
	"bufio"
	"helper"
	"io"
	"strings"
)

func Run(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	h := helper.NewHelper()

	for scanner.Scan() {
		input := strings.Fields(scanner.Text())

		if input[0] == "+" {
			h.AddText(strings.Join(input[1:], " "))
			w.Write([]byte("Added successfully!\n"))

		} else if input[0] == "?" {
			s, ok := h.MakeNewRequest(input[1])
			if !ok {
				w.Write([]byte("No candidates found!\n"))
			} else {
				w.Write([]byte(s + "\n"))
			}

		} else if input[0] == ">" {
			s, ok := h.ExpandRequest(input[1])
			if !ok {
				w.Write([]byte("No candidates found!\n"))
			} else {
				w.Write([]byte(s + "\n"))
			}

		} else {
			w.Write([]byte("Unknown command!\n"))
		}
	}
}
