package day11

import (
	"aoc2025/utils"
	"maps"
	"strings"
)

type Device struct {
	name    string
	outputs []*Device
}

func (d *Device) addOutput(device *Device) {
	d.outputs = append(d.outputs, device)
}

type Network struct {
	Devices []*Device
}

func (n *Network) find(deviceName string) *Device {
	for _, device := range n.Devices {
		if device.name == deviceName {
			return device
		}
	}

	return nil
}

func (n *Network) add(deviceName string) *Device {
	newDevice := &Device{name: deviceName}

	n.Devices = append(n.Devices, newDevice)

	return newDevice
}

func (n *Network) ensureExist(deviceName string) *Device {
	device := n.find(deviceName)

	if device != nil {
		return device
	}

	return n.add(deviceName)
}

func traverse(device *Device) int {
	if device.name == "out" {
		return 1
	}

	sum := 0

	for _, output := range device.outputs {
		sum += traverse(output)
	}

	return sum
}

func buildNetwork(inputPath string) *Network {
	lr, _ := utils.NewLineReader(inputPath)

	network := Network{make([]*Device, 0)}

	for {
		line, ok := lr.Next()

		if !ok {
			break
		}

		parts := strings.Split(line, " ")

		from := network.ensureExist(parts[0][:len(parts[0])-1])

		for _, to := range parts[1:] {
			toDevice := network.ensureExist(to)

			from.addOutput(toDevice)
		}
	}

	return &network
}

func Part1(inputPath string) int {
	network := buildNetwork(inputPath)

	return traverse(network.find("you"))
}

func traverseWithHistory(device *Device, toFind []*Device, visited map[*Device]struct{}) int {
	if device.name == "out" {
		for _, deviceToFind := range toFind {
			_, ok := visited[deviceToFind]

			if !ok {
				return 0
			}
		}

		return 1
	}

	_, seen := visited[device]

	if seen {
		return 0
	}

	newVisited := maps.Clone(visited)

	newVisited[device] = struct{}{}

	sum := 0

	for _, output := range device.outputs {
		sum += traverseWithHistory(output, toFind, newVisited)
	}

	return sum
}

func Part2(inputPath string) int {
	network := buildNetwork(inputPath)

	start := network.find("svr")

	toFind := make([]*Device, 2)
	toFind[0] = network.find("dac")
	toFind[1] = network.find("fft")

	return traverseWithHistory(start, toFind, make(map[*Device]struct{}))
}
