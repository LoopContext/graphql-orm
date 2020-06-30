package main

import (
	"github.com/loopcontext/graphql-orm/cmd"
	"github.com/loopcontext/graphql-orm/events"
)

func main() {
	cmd.Execute()
}

// this is just for importing the events package and adding it to the go modules
var _ events.Event
