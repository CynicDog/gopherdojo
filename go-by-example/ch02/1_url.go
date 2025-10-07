package main

import "fmt"

type URL struct {
	Scheme string
	Host   string
	Path   string
}

func Parse(rawURL string) (*URL, error) {
	return &URL{
		Scheme: "https",
		Host:   "github.com",
		Path:   "cynicdog",
	}, nil
}

func (u *URL) String() string {
	return fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, u.Path)
}
