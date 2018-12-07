package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/runeflow/runeflow/client"
	"github.com/runeflow/runeflow/commandhandler"
	"github.com/runeflow/runeflow/message"
	"github.com/runeflow/runeflow/monitor"
	"github.com/runeflow/runeflow/monitor/apache"
	"github.com/runeflow/runeflow/monitor/cpu"
	"github.com/runeflow/runeflow/monitor/disk"
	"github.com/runeflow/runeflow/monitor/hostname"
	"github.com/runeflow/runeflow/monitor/memory"
	"github.com/runeflow/runeflow/monitor/websites"
)

func run(c *client.Client) {
	defer log.Print("run completed")

	cmdHandler := commandhandler.NewCommandHandler()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	reestablish := true

	done := make(chan struct{})
	go func() {
		<-interrupt
		log.Print("received interrupt")
		reestablish = false
		close(done)
	}()

	for {
		log.Print("establishing connection...")
		if err := establishConnection(c, cmdHandler, done); err != nil {
			log.Printf("error establishing connection: %v", err)
		}
		log.Printf("connection ended")
		if !reestablish {
			log.Print("termination requested: should not reestablish connection")
			return
		}
		log.Print("waiting for 5 seconds before retrying...")
		select {
		case <-time.After(5 * time.Second):
			continue
		case <-done:
			return
		}
	}
}

func establishConnection(c *client.Client, cmdHandler *commandhandler.CommandHandler, interrupt <-chan struct{}) error {
	if err := c.Connect(); err != nil {
		return fmt.Errorf("error connecting to server: %v", err)
	}
	defer c.CloseConn()
	log.Print("connection established")

	log.Print("starting monitors")
	monitors := setupMonitors()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			msg, err := c.ReadMessage()
			if err != nil {
				log.Printf("error reading message: %v", err)
				return
			}
			switch msg.Type {
			case message.CmdMessage:
				if payload, err := msg.ParseCmd(); err == nil {
					cmdHandler.HandleCommand(payload.ID, payload.Action)
				}
			}
			log.Printf("got message: %v", msg)
		}
	}()

	statTicker := time.NewTicker(5 * time.Second)
	defer statTicker.Stop()

	for {
		select {
		case <-done:
			return nil
		case <-statTicker.C:
			if err := c.SendStats(collectStats(monitors)); err != nil {
				log.Printf("error sending stats: %v", err)
			}
		case <-interrupt:
			log.Print("sending close message")
			if err := c.SendClose(); err != nil {
				return err
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return nil
		}
	}
}

func setupMonitors() []monitor.Monitor {
	return []monitor.Monitor{
		disk.NewMonitor(),
		memory.NewMonitor(),
		cpu.NewMonitor(),
		websites.NewMonitor(),
		apache.NewMonitor(),
		hostname.NewMonitor(),
	}
}

func collectStats(monitors []monitor.Monitor) *message.StatsPayload {
	stats := &message.StatsPayload{Timestamp: time.Now()}
	for _, mon := range monitors {
		mon.Sample(stats)
	}
	return stats
}
