package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

const (
	start = "start"
	stop  = "stop"
)

type Transformer struct {
	LoadedTasks Tasks
}

func (transformer *Transformer) Transform() map[string]string {
	transformedTasks := map[string]string{}
	tasks := transformer.LoadedTasks.Items
	for _, task := range tasks {
		if _, inMap := transformedTasks[task.getIdentifier()]; inMap {
			continue
		}
		taskSeconds, isActive := transformer.TrackingToSeconds(task.getIdentifier())
		humanTime := transformer.SecondsToHuman(taskSeconds)

		status := ""
		if isActive {
			status = "(running)"
		}
		transformedTask := fmt.Sprintf("%s     %s %s", humanTime, task.getIdentifier(), status)
		transformedTasks[task.getIdentifier()] = transformedTask
	}
	return transformedTasks
}

func (transformer *Transformer) SecondsToHuman(totalSeconds int) string {
	hours := math.Floor(float64(((totalSeconds % 31536000) % 86400) / 3600))
	minutes := math.Floor(float64((((totalSeconds % 31536000) % 86400) % 3600) / 60))
	seconds := (((totalSeconds % 31536000) % 86400) % 3600) % 60

	return fmt.Sprintf("%dh:%dm:%ds", int(hours), int(minutes), int(seconds))
}

func (transformer *Transformer) TrackingToSeconds(identifier string) (int, bool) {
	nextAction := "start"
	var durationInSeconds float64
	var startTime, stopTime time.Time

	tasks := transformer.LoadedTasks.getByIdentifier(identifier)
	for _, task := range tasks.Items {
		if task.getAction() == start && nextAction == start {
			nextAction = stop
			startTime = parseTime(task.getAt())
		}
		if task.getAction() == stop && nextAction == stop {
			nextAction = start
			stopTime = parseTime(task.getAt())
			durationInSeconds += stopTime.Sub(startTime).Seconds()
		}
	}

	if isActive(nextAction) {
		durationInSeconds += time.Since(startTime).Seconds()
	}

	return int(durationInSeconds), isActive(nextAction)
}

func isActive(nextAction string) bool {
	return nextAction == stop
}

func parseTime(at string) time.Time {
	then, err := time.Parse(time.RFC3339, at)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return then
}
