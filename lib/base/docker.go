package base

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"os"
)

type Docker struct {
	cli *client.Client
}

func NewDocker() (*Docker, error) {
	var (
		err    error
		docker Docker
	)
	docker.cli, err = client.NewClient("tcp://127.0.0.1:1990", "1.18.2", nil, map[string]string{"Content-type": "application/x-tar"})
	if err != nil {
		panic(err)
	}
	return &docker, err
}

func (e *Docker) ImageList() []types.ImageSummary {
	imgs, err := e.cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	return imgs
}

func (e *Docker) ImagePull(imgName string) error {
	events, err := e.cli.ImagePull(context.Background(), imgName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer events.Close()
	_, err = io.Copy(os.Stdout, events)
	return err
}

func (e *Docker) ImagePullAuth(imgName string, userName string, password string) error {
	authConfig := types.AuthConfig{
		Username: userName,
		Password: password,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	events, err := e.cli.ImagePull(context.Background(), imgName, types.ImagePullOptions{RegistryAuth: authStr})
	if err != nil {
		panic(err)
	}
	defer events.Close()
	_, err = io.Copy(os.Stdout, events)
	return err
}

func (e *Docker) ImageLoad(fileName string) (types.ImageLoadResponse, error) {
	fileObj, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fileObj.Close()
	response, err := e.cli.ImageLoad(context.Background(), fileObj, true)
	if err != nil {
		panic(err)
	}
	return response, err
}

func (e *Docker) ImageRemove(imgID string) ([]types.ImageDelete, error) {
	response, err := e.cli.ImageRemove(context.Background(), imgID, types.ImageRemoveOptions{})
	return response, err
}

func (e *Docker) ImageSave(imgID string, fileName string) {
	var idList []string
	idList = append(idList, imgID)
	writer, err := e.cli.ImageSave(context.Background(), idList)
	if err != nil {
		panic(err)
	}
	outfile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_SYNC|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	_, err = io.Copy(outfile, writer)
	if err != nil {
		panic(err)
	}
}

func (e *Docker) ImageTag(src string, target string) error {
	err := e.cli.ImageTag(context.Background(), src, target)
	return err
}

func (e *Docker) ContainerList() []types.Container {
	containers, err := e.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	return containers
}

func (e *Docker) ContainerCommit(containerID string) (types.IDResponse, error) {
	response, err := e.cli.ContainerCommit(context.Background(), containerID, types.ContainerCommitOptions{})
	if err != nil {
		panic(err)
	}
	return response, err
}

func (e *Docker) ContainerKillNotRemove(containerID, signal string) error {
	err := e.cli.ContainerKill(context.Background(), containerID, signal)
	if err != nil {
		panic(err)
	}
	return err
}

func (e *Docker) ContainerRemove(containerID string) error {
	err := e.cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{})
	if err != nil {
		panic(err)
	}
	return err
}

func (e *Docker) ContainerLogs(container string) (io.ReadCloser, error) {
	reader, err := e.cli.ContainerLogs(context.Background(), container, types.ContainerLogsOptions{ShowStdout: true})
	return reader, err
}
