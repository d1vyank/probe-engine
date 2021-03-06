package oonimkall

import (
	"testing"
)

func TestUnitDisabledEvents(t *testing.T) {
	out := make(chan *eventRecord)
	emitter := newEventEmitter([]string{"log"}, out)
	go func() {
		emitter.Emit("log", eventLog{Message: "foo"})
		close(out)
	}()
	var count int64
	for ev := range out {
		if ev.Key == "log" {
			count++
		}
	}
	if count > 0 {
		t.Fatal("cannot disable events")
	}
}

func TestUnitEmitFailureStartup(t *testing.T) {
	out := make(chan *eventRecord)
	emitter := newEventEmitter([]string{}, out)
	go func() {
		emitter.EmitFailureStartup("mocked error")
		close(out)
	}()
	var found bool
	for ev := range out {
		if ev.Key == "failure.startup" {
			evv := ev.Value.(eventFailureGeneric) // panic if not castable
			if evv.Failure == "mocked error" {
				found = true
			}
		}
	}
	if !found {
		t.Fatal("did not see expected event")
	}
}

func TestUnitEmitStatusProgress(t *testing.T) {
	out := make(chan *eventRecord)
	emitter := newEventEmitter([]string{}, out)
	go func() {
		emitter.EmitStatusProgress(0.7, "foo")
		close(out)
	}()
	var found bool
	for ev := range out {
		if ev.Key == "status.progress" {
			evv := ev.Value.(eventStatusProgress) // panic if not castable
			if evv.Message == "foo" && evv.Percentage == 0.7 {
				found = true
			}
		}
	}
	if !found {
		t.Fatal("did not see expected event")
	}
}
