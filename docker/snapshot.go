package docker

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/portainer/agent"
	portainer "github.com/portainer/portainer/api"
)

// This seems alot simpler than the original in portainer/portainer...
func CreateSnapshot() (*portainer.DockerSnapshot, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion(agent.SupportedDockerAPIVersion))
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	_, err = cli.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	snapshot := &portainer.DockerSnapshot{
		StackCount: 0,
	}

	err = snapshotInfo(snapshot, cli)
	if err != nil {
		log.Printf("[WARN] [docker,snapshot] [message: unable to snapshot engine information] [err: %s]", err)
	}

	if snapshot.Swarm {
		err = snapshotSwarmServices(snapshot, cli)
		if err != nil {
			log.Printf("[WARN] [docker,snapshot] [message: unable to snapshot Swarm services] [err: %s]", err)
		}

		err = snapshotNodes(snapshot, cli)
		if err != nil {
			log.Printf("[WARN] [docker,snapshot] [message: unable to snapshot Swarm nodes] [err: %s]", err)
		}
	}

	err = snapshotContainers(snapshot, cli)
	if err != nil {
		log.Printf("[WARN] [docker,snapshot] [message: unable to snapshot containers] [err: %s]", err)
	}

	err = snapshotImages(snapshot, cli)
	if err != nil {
		log.Printf("[WARN] [docker,snapshot] [message: unable to snapshot images] [err: %s]", err)
	}

	err = snapshotVolumes(snapshot, cli)
	if err != nil {
		log.Printf("[WARN] [docker,snapshot] [message: unable to snapshot volumes] [err: %s]", err)
	}

	err = snapshotNetworks(snapshot, cli)
	if err != nil {
		log.Printf("[WARN] [docker,snapshot] [message: unable to snapshot networks] [err: %s]", err)
	}

	err = snapshotVersion(snapshot, cli)
	if err != nil {
		log.Printf("[WARN] [docker,snapshot] [message: unable to snapshot engine version] [err: %s]", err)
	}

	snapshot.Time = time.Now().Unix()
	return snapshot, nil
}

func snapshotInfo(snapshot *portainer.DockerSnapshot, cli *client.Client) error {
	info, err := cli.Info(context.Background())
	if err != nil {
		return err
	}

	snapshot.Swarm = info.Swarm.ControlAvailable
	snapshot.DockerVersion = info.ServerVersion
	snapshot.TotalCPU = info.NCPU
	snapshot.TotalMemory = info.MemTotal
	snapshot.SnapshotRaw.Info = info
	return nil
}

func snapshotNodes(snapshot *portainer.DockerSnapshot, cli *client.Client) error {
	nodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		return err
	}
	var nanoCpus int64
	var totalMem int64
	for _, node := range nodes {
		nanoCpus += node.Description.Resources.NanoCPUs
		totalMem += node.Description.Resources.MemoryBytes
	}
	snapshot.TotalCPU = int(nanoCpus / 1e9)
	snapshot.TotalMemory = totalMem
	snapshot.NodeCount = len(nodes)
	return nil
}

func snapshotSwarmServices(snapshot *portainer.DockerSnapshot, cli *client.Client) error {
	stacks := make(map[string]struct{})

	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		return err
	}

	for _, service := range services {
		for k, v := range service.Spec.Labels {
			if k == "com.docker.stack.namespace" {
				stacks[v] = struct{}{}
			}
		}
	}

	snapshot.ServiceCount = len(services)
	snapshot.StackCount += len(stacks)
	return nil
}

func snapshotContainers(snapshot *portainer.DockerSnapshot, cli *client.Client) error {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return err
	}

	runningContainers := 0
	stoppedContainers := 0
	healthyContainers := 0
	unhealthyContainers := 0
	stacks := make(map[string]struct{})
	for _, container := range containers {
		if container.State == "exited" {
			stoppedContainers++
		} else if container.State == "running" {
			runningContainers++
		}

		if strings.Contains(container.Status, "(healthy)") {
			healthyContainers++
		} else if strings.Contains(container.Status, "(unhealthy)") {
			unhealthyContainers++
		}

		for k, v := range container.Labels {
			if k == "com.docker.compose.project" {
				stacks[v] = struct{}{}
			}
		}
	}

	snapshot.RunningContainerCount = runningContainers
	snapshot.StoppedContainerCount = stoppedContainers
	snapshot.HealthyContainerCount = healthyContainers
	snapshot.UnhealthyContainerCount = unhealthyContainers
	snapshot.StackCount += len(stacks)
	snapshot.SnapshotRaw.Containers = containers
	return nil
}

func snapshotImages(snapshot *portainer.DockerSnapshot, cli *client.Client) error {
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return err
	}

	snapshot.ImageCount = len(images)
	snapshot.SnapshotRaw.Images = images
	return nil
}

func snapshotVolumes(snapshot *portainer.DockerSnapshot, cli *client.Client) error {
	volumes, err := cli.VolumeList(context.Background(), filters.Args{})
	if err != nil {
		return err
	}

	snapshot.VolumeCount = len(volumes.Volumes)
	snapshot.SnapshotRaw.Volumes = volumes
	return nil
}

func snapshotNetworks(snapshot *portainer.DockerSnapshot, cli *client.Client) error {
	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		return err
	}
	snapshot.SnapshotRaw.Networks = networks
	return nil
}

func snapshotVersion(snapshot *portainer.DockerSnapshot, cli *client.Client) error {
	version, err := cli.ServerVersion(context.Background())
	if err != nil {
		return err
	}
	snapshot.SnapshotRaw.Version = version
	return nil
}