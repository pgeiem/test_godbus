package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
)

func main() {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	//var s string
	obj := conn.Object("com.iem.livio.gui", "/fields/username")
	propertyObject, err := obj.GetProperty("com.iem.livio.gui.formfield.value")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to call function (is the server example running?):", err)
		os.Exit(1)
	}

	if propertyObject.Value() == nil {
		fmt.Fprintln(os.Stderr, "Unexpected nil value received when reading property")
		os.Exit(1)
	}

	propertyValue, ok := propertyObject.Value().(string)
	if !ok {
		fmt.Fprintln(os.Stderr, "Received unexpected type as value, expected string got '%T'", propertyValue)
		os.Exit(1)
	}

	fmt.Println("Result from calling Foo function on com.github.guelfey.Demo interface:")
	fmt.Println(propertyValue)
}
