package main

import "fmt"

type Device struct {
	ID   string
	Name string
	Type string
}

type Connection struct {
	From string
	To   string
}

type Topology struct {
	Devices     []Device
	Connections []Connection
}

func printTopology(t Topology) {
	fmt.Println("Devices : ")
	for _, device := range t.Devices {
		fmt.Printf("- %s (%s)\n", device.Name, device.Type)
	}
	fmt.Println("\nConnections : ")

	for _, connection := range t.Connections {
		fmt.Printf("%s --> %s\n", connection.From, connection.To)
	}
}

func pointToPoint() Topology {
	return Topology{
		Devices: []Device{
			{"1", "PC", "Computer"},
			{"2", "Router", "Router"},
		},

		Connections: []Connection{
			{"PC", "Router"},
		},
	}
}

func daisyChain() Topology {
	return Topology{
		Devices: []Device{
			{"1", "Sensor1", "Sensor"},
			{"2", "Sensor2", "Sensor"},
			{"3", "Sensor3", "Sensor"},
			{"4", "Gateway", "Gateway"},
		},
		Connections: []Connection{
			{"Sensor1", "Sensor2"},
			{"Sensor2", "Sensor3"},
			{"Sensor3", "Gateway"},
		},
	}
}

func busTopology() Topology {
	return Topology{
		Devices: []Device{
			{"1", "Bus", "Cable"},
			{"2", "PC1", "PC"},
			{"3", "PC2", "PC"},
			{"4", "PC3", "PC"},
			{"5", "PC4", "PC"},
		},
		Connections: []Connection{
			{"Bus", "PC1"},
			{"Bus", "PC2"},
			{"Bus", "PC3"},
			{"Bus", "PC4"},
		},
	}
}

func ringTopology() Topology {
	return Topology{
		Devices: []Device{
			{"1", "A", "Switch"},
			{"2", "B", "Switch"},
			{"3", "C", "Switch"},
			{"4", "D", "Switch"},
		},
		Connections: []Connection{
			{"A", "B"},
			{"B", "C"},
			{"C", "D"},
			{"D", "A"},
		},
	}
}

func starTopology() Topology {
	return Topology{
		Devices: []Device{
			{"1", "Switch", "Switch"},
			{"2", "PC1", "PC"},
			{"3", "PC2", "PC"},
			{"4", "PC3", "PC"},
		},
		Connections: []Connection{
			{"Switch", "PC1"},
			{"Switch", "PC2"},
			{"Switch", "PC3"},
		},
	}
}

func meshTopology() Topology {
	return Topology{
		Devices: []Device{
			{"1", "A", "Router"},
			{"2", "B", "Router"},
			{"3", "C", "Router"},
			{"4", "D", "Router"},
		},
		Connections: []Connection{
			{"A", "B"},
			{"A", "C"},
			{"A", "D"},
			{"B", "C"},
			{"B", "D"},
			{"C", "D"},
		},
	}
}

func main() {

	fmt.Println("===== Point-to-Point =====")
	printTopology(pointToPoint())

	fmt.Println("\n===== Daisy Chain =====")
	printTopology(daisyChain())

	fmt.Println("\n===== Bus =====")
	printTopology(busTopology())

	fmt.Println("\n===== Ring =====")
	printTopology(ringTopology())

	fmt.Println("\n===== Star =====")
	printTopology(starTopology())

	fmt.Println("\n===== Mesh =====")
	printTopology(meshTopology())

}
