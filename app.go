package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})

	// fmt.Println(path)
	// fmt.Println(err)
}

func (a *App) OpenFiles() (string, error) {
	var input Input

	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	if err = json.Unmarshal(data, &input); err != nil {
		return "", err
	}

	output, err := input.ToOutput()
	if err != nil {
		return "", err
	}
	bs, err := json.Marshal(output)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

type Input struct {
	Units       []InputUnit       `json:"units"`
	Topics      []InputTopic      `json:"topics"`
	Microtopics []InputMicrotopic `json:"microtopics"`
	Edges       []InputEdge       `json:"edges"`
}
type InputUnit struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type InputTopic struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Parent string `json:"parent"`
}
type InputMicrotopic struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Parent string `json:"parent"`
}
type InputEdge struct {
	Source string `json:"source,omitempty"`
	Target string `json:"target,omitempty"`
}

type OutputNode struct {
	Data CytoNode `json:"data"`
}
type OutputEdge struct {
	Data CytoEdge `json:"data"`
}
type CytoNode struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Bg     string `json:"bg"`
	Color  string `json:"color"`
	Group  string `json:"group"`
	Parent string `json:"parent"`
}
type CytoEdge struct {
	Id     string `json:"id"`
	Source string `json:"source,omitempty"`
	Target string `json:"target,omitempty"`
}

func (i Input) ToOutput() ([]any, error) {
	var outputs []any
	//exists := map[string]struct{}{}

	//for _, unit := range i.Units {
	//	exists[unit.Id] = struct{}{}
	//
	//	var output OutputNode
	//
	//	output.Data.Group = "units"
	//	output.Data.Bg = "rgb(255, 172, 48)"
	//	output.Data.Name = unit.Name
	//	output.Data.Id = unit.Id
	//	output.Data.Color = "white"
	//
	//	outputs = append(outputs, &output)
	//}

	//for _, topic := range i.Topics {
	//	if topic.Parent == "" {
	//		return nil, fmt.Errorf("topic(%s) has no parent", topic.Id)
	//	}
	//	if _, ok := exists[topic.Parent]; !ok {
	//		return nil, fmt.Errorf("topic(%s)'s parent does not exists (%s)", topic.Id, topic.Parent)
	//	}
	//
	//	exists[topic.Id] = struct{}{}
	//
	//	var output OutputNode
	//
	//	output.Data.Group = "topics"
	//	output.Data.Bg = "rgb(252, 245, 235)"
	//	output.Data.Name = topic.Name
	//	output.Data.Id = topic.Id
	//	output.Data.Color = "white"
	//	output.Data.Parent = topic.Parent
	//
	//	outputs = append(outputs, &output)
	//}

	for _, microtopic := range i.Microtopics {
		//if microtopic.Parent == "" {
		//	return nil, fmt.Errorf("microtopic(%s) has no parent", microtopic.Id)
		//}
		//if _, ok := exists[microtopic.Parent]; !ok {
		//	return nil, fmt.Errorf("microtopic(%s)'s parent does not exists (%s)", microtopic.Id, microtopic.Parent)
		//}

		//exists[microtopic.Id] = struct{}{}

		var container OutputNode
		var data OutputNode

		container.Data.Group = "microtopic_container"
		container.Data.Bg = "rgb(164, 233, 133)"
		container.Data.Id = microtopic.Id
		//container.Data.Color = "white"
		container.Data.Parent = microtopic.Parent

		data.Data.Group = "microtopics"
		//data.Data.Bg = "rgb(164, 233, 133)"
		data.Data.Name = microtopic.Name
		data.Data.Id = microtopic.Id + "_data"
		data.Data.Color = "white"
		data.Data.Parent = microtopic.Id

		outputs = append(outputs, &container)
		outputs = append(outputs, &data)
	}

	for _, edge := range i.Edges {
		//if _, ok := exists[edge.Source]; !ok {
		//	return nil, fmt.Errorf("microtopic(%s) does not exists for edge", edge.Source)
		//} else if _, ok := exists[edge.Target]; !ok {
		//	return nil, fmt.Errorf("microtopic(%s) does not exists for edge", edge.Target)
		//}

		var data OutputEdge

		data.Data.Source = edge.Source
		data.Data.Target = edge.Target
		data.Data.Id = fmt.Sprintf("%d-%d", rand.Int(), rand.Int())

		outputs = append(outputs, &data)
	}

	return outputs, nil
}
