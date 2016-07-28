package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/**
* Interfaces
**/

type Logger interface {
	Log(...interface{})
}

type SuspendResumer interface {
	Suspend() error
	Resume() error
}

type Job interface {
	Logger
	SuspendResumer
	Run() error
}

type ServerPoller interface {
	PollServer() (string, error)
}

/*
 * Structs that actually implement interfaces
 */

type PollerLogger struct{}

func (p *PollerLogger) Log(args ...interface{}) {
	log.Println(args...)
}

type URLServerPoller struct {
	resourceUrl string
}

func (p *URLServerPoller) PollServer() (string, error) {
	resp, err := http.Get(p.resourceUrl)
	if err != nil {
		return "", err
	}

	// p.Log(p.resourceUrl, "--", resp.Status)

	return fmt.Sprintf("%v --> %v", p.resourceUrl, resp.Status), nil
}

type PollSuspendResumer struct {
	SuspendCh chan bool
	ResumeCh  chan bool
}

func (p *PollSuspendResumer) Suspend() error {
	p.SuspendCh <- true
	return nil
}

func (p *PollSuspendResumer) Resume() error {
	p.ResumeCh <- true
	return nil
}

type PollerJob struct {
	WaitDuration time.Duration
	ServerPoller
	Logger
	*PollSuspendResumer
}

func (p PollerJob) Run() error {
	for {
		select {
		case <-p.SuspendCh:
			<-p.ResumeCh
		default:
			// var resp string
			if resp, err := p.PollServer(); err != nil {
				p.Log("Error trying to get state: ", err)
			} else {
				p.Log(resp)
			}

			// p.Log(p.resourceUrl, "--", resp.Status)

			time.Sleep(1 * p.WaitDuration)
		}
	}
}

func NewPollerJob(resourceUrl string, waitDuration time.Duration) PollerJob {
	return PollerJob{
		WaitDuration: waitDuration,
		Logger:       &PollerLogger{},
		ServerPoller: &URLServerPoller{
			resourceUrl: resourceUrl,
		},
		PollSuspendResumer: &PollSuspendResumer{
			SuspendCh: make(chan bool),
			ResumeCh:  make(chan bool),
		},
	}
}

func main() {
	p := NewPollerJob(fmt.Sprintf("https://api.github.com/repos/%s/releases", "docker/machine"), time.Second*1)
	go p.Run()
	time.Sleep(5 * time.Second)

	p.Log("Suspending monitoring of server for 5 seconds...")
	p.Suspend()
	time.Sleep(5 * time.Second)

	p.Log("Resuming job...")
	p.Resume()

	// Wait for a bit before exiting
	time.Sleep(5 * time.Second)
}
