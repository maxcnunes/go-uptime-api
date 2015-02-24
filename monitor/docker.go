package monitor

import (
	docker "github.com/fsouza/go-dockerclient"
	"log"
	"strings"
)

var (
	client *docker.Client
)

const (
	dockerStart   = "start"
	dockerDestroy = "destroy"
	endpoint      = "unix:///tmp/docker.sock"
)

// Container ...
type Container struct {
	URL  string
	Name string
}

// StartEventListener ...
func StartEventListener(data *DataMonitor) {
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

				switch event.Status {
				case dockerStart:
					LoadVirtualHostsToURLS(data)
				case dockerDestroy:
					data.RemoveTarget("http://twitter.com/")
				}
			}
		}
		log.Fatalf("Exitting the docker events loop")
	}()
}

// LoadVirtualHostsToURLS ...
func LoadVirtualHostsToURLS(data *DataMonitor) {
	filters := make(map[string][]string)
	filters["status"] = []string{"running"}

	containers, _ := client.ListContainers(docker.ListContainersOptions{All: false, Filters: filters})
	for _, container := range containers {
		log.Printf("ID: %s", container.ID)
		log.Printf("Names: %s", container.Names)

		virtualHost := getVirtualHost(container.ID)

		if virtualHost != "" {
			// assumes all virtual host are http for while
			data.AddTarget("http://" + virtualHost)
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
