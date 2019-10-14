package main

import (
	"fmt"

	"github.com/containers/libpod/cmd/podman/varlink"
	"github.com/varlink/go/varlink"
)

func main() {

	varlinkConnection, err := varlink.NewConnection("unix:/run/user/1000/podman/io.podman")
	if err != nil {
		fmt.Printf("Error1: %v\n", err)
		return
	}
	defer varlinkConnection.Close()

	args := []string{"localhost/helloworld-test"}
	env := []string{"HW_WHO=fred"}
	vol := []string{"/home/max/test:/host/test"}
	net := "test3"

	createOpts := iopodman.Create{Args: args, Env: &env, Volume: &vol, Network: &net}

	containerID, err := iopodman.CreateContainer().Call(varlinkConnection, createOpts)
	if err != nil {
		fmt.Printf("Error2: %v\n", err)
		return
	}
	fmt.Printf("Created %s\n", containerID)

	_, err = iopodman.StartContainer().Call(varlinkConnection, containerID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Started %s\n", containerID)

	info, err := iopodman.GetContainer().Call(varlinkConnection, containerID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Container %v\n", info)

	result, err := iopodman.StopContainer().Call(varlinkConnection, containerID, 120)
	if err != nil {
		fmt.Printf("Error stopping container: %v: %v\n", containerID, err)
		return
	}
	fmt.Printf("Stopped container %v\n", result)

	result, err = iopodman.RemoveContainer().Call(varlinkConnection, containerID, true, false)
	if err != nil {
		fmt.Printf("Error removing container: %v: %v\n", containerID, err)
		return
	}
	fmt.Printf("Removed container %v\n", result)
}
