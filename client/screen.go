package main

import "fmt"

type Screen struct {
	Info []string
}

func (s *Screen) Print() {
	for _, info := range s.Info {
		fmt.Println(info)
	}
}

func (s *Screen) Append(input ...string) {
	for _, str := range input {
		s.Info = append(s.Info, str)
	}
}

func (s *Screen) Clf() {
	s.Info = []string{}
}
