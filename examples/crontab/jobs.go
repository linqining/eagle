package main

import (
	"github.com/linqining/eagle/pkg/log"
)

type GreetingJob struct {
	Name string
}

func (g GreetingJob) Run() {
	log.Info("Hello ", g.Name)
}

type SendEmail struct {
	Name string
}

func (g SendEmail) Run() {
	log.Info("Send mail to... ", g.Name)
}
