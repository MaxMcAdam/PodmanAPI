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

	podCreate := iopodman.PodCreate{Name: "pod1"}
	podId, err := iopodman.CreatePod().Call(varlinkConnection, podCreate)
	if err != nil {
		fmt.Printf("Error create pod1: %v", err)
		return
	}
	fmt.Printf("Pod1 created: %v\n", podId)

	args := []string{"localhost/helloworld-test"}
	env := []string{"HW_WHO=fred"}
	pod := "pod1"
	hostname1 := "container1"
	name := "container1"

	createOpts := iopodman.Create{Args: args, Env: &env, Pod: &pod, Hostname: &hostname1, Name: &name}

	containerID, err := iopodman.CreateContainer().Call(varlinkConnection, createOpts)
	if err != nil {
		fmt.Printf("Error creating container 1: %v\n", err)
		return
	}
	fmt.Printf("Created %s\n", containerID)

	hostname2 := "container2"
	name2 := "conatiner2"

	createOpts2 := iopodman.Create{Args: args, Env: &env, Pod: &pod, Hostname: &hostname2, Name: &name2}

	container2ID, err := iopodman.CreateContainer().Call(varlinkConnection, createOpts2)
	if err != nil {
		fmt.Printf("Error creating container 2: %v\n", err)
		return
	}
	fmt.Printf("Created %s\n", containerID)

	_, err = iopodman.StartContainer().Call(varlinkConnection, containerID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Started %s\n", containerID)

	_, err = iopodman.StartContainer().Call(varlinkConnection, container2ID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Started %s\n", container2ID)

	containerInfo, err := iopodman.GetContainer().Call(varlinkConnection, containerID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Container %v\n", containerInfo)

	container2Info, err := iopodman.GetContainer().Call(varlinkConnection, container2ID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Container %v\n", container2Info)

	/*
		execOpts := iopodman.ExecOpts{Name: name, Cmd: []string{"/bin/ping " + name2}}

		err = iopodman.ExecContainer().Call(varlinkConnection, execOpts)
		if err != nil {
			fmt.Printf("Error pinging container1 %v, %v\n", execOpts, err)
			return
		}
		fmt.Printf("Execed %v\n", execOpts)
	*/

	result, err := iopodman.RemoveContainer().Call(varlinkConnection, containerID, true, false)
	if err != nil {
		fmt.Printf("Error removing container: %v: %v\n", containerID, err)
		return
	}
	fmt.Printf("Removed container %v\n", result)

	result, err = iopodman.RemoveContainer().Call(varlinkConnection, container2ID, true, false)
	if err != nil {
		fmt.Printf("Error removing container: %v: %v\n", container2ID, err)
	}
	fmt.Printf("Removed container %v\n", result)

	result, err = iopodman.RemovePod().Call(varlinkConnection, podId, true)
	if err != nil {
		fmt.Printf("Error removing pod1: %v\n", err)
	}
	fmt.Printf("Removed pod: %v\n", result)

}
