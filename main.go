package main

import (
	"context"
	"github.com/maximorov/auditor"
	"time"
)

type myRepository struct {
}

func (r *myRepository) CreateMany(_ context.Context, _ []auditor.Valuable) (int, error) {
	return 3, nil
}

type some struct {
	text string
}

func (s *some) Value() string {
	return `Some ` + s.text
}

func main() {
	a := auditor.New(&myRepository{})
	a.Update(&some{`text1`})
	a.Update(&some{`text2`})
	a.Update(&some{`text3`})

	time.Sleep(time.Second * 5)
}
