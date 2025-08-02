
// +build plugin

package main

import "fmt"

// Name returns the name of the plugin
func Name() string {
    return "example_plugin"
}

// Init initializes the plugin
func Init() {
    fmt.Println("Example plugin loaded.")
}
