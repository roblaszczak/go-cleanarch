package main

import "fmt"

type sliceFlag []string

func (s *sliceFlag) String() string {
	return fmt.Sprintf("%s", s)
}

func (s *sliceFlag) Set(v string) error {
	*s = append(*s, v)
	return nil
}
