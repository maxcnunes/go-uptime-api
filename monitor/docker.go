package monitor

import (
	"log"
	"strings"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/maxcnunes/monitor-api/monitor/data"
)

var (
	client *docker.Client
)

const (
	dockerCreate = "create"
	endpoint     = "unix:///tmp/docker.sock"
)

// Container ...
type Container struct {
	URL  string
	Name string
}

// StartEventListener ...
func StartEventListener(data *data.DataMonitor) {
	client, _ = docker.NewClient(endpoint)

	dockerEvents := make(chan *docker.APIEvents)
	log.Println("Starting docker events...")
	go func() {
		// add our channel as an event listener for docker events
		if err := client.AddEventListener(dockerEvents); err != nil {
			log.Fatalf("Unable to register docker events listener, error: %s", err)
		}

		// start the event loop and wait for docker events
		log.Print("Entering into the docker events loop")
		for {
			select {
			case event := <-dockerEvents:
				log.Printf("Received docker event status: %s, id: %s", event.Status, event.ID)

				// only cares to created containers
				if event.Status == "create" {
					virtualHost := getVirtualHost(event.ID)
					if virtualHost != "" {
						// assumes all virtual host are http for while
						data.Target.Create("http://" + virtualHost)
					}
				}
			}
		}
	}()
}

// LoadAllVirtualHosts ...
func LoadAllVirtualHosts(data *data.DataMonitor) {
	filters := make(map[string][]string)
	filters["status"] = []string{"running"}

	containers, _ := client.ListContainers(docker.ListContainersOptions{All: false, Filters: filters})
	for _, container := range containers {
		log.Printf("ID: %s", container.ID)
		log.Printf("Names: %s", container.Names)

		virtualHost := getVirtualHost(container.ID)

		if virtualHost != "" {
			// assumes all virtual host are http for while
			data.Target.Create("http://" + virtualHost)
		}
	}
}

func getVirtualHost(containerID string) string {
	info, err := client.InspectContainer(containerID)
	if err != nil {
		log.Fatalf("Unable to inspect container %s, error: %s", containerID, err)
	}

	// log.Print(">>> Container Info <<<")
	// log.Printf("Args %s", info.Args)
	// log.Printf("Name %s", info.Name)
	// log.Printf("Env", info.Config.Env)
	var virtualHost string
	for _, env := range info.Config.Env {
		// log.Printf(">>> env %s", env)

		if strings.Contains(env, "VIRTUAL_HOST=") {
			virtualHost = strings.Replace(env, "VIRTUAL_HOST=", "", -1)
			break
		}
	}

	return virtualHost
}
