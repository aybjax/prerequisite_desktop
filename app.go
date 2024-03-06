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

	bs, err := json.Marshal(input.ToOutput())
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

type Input struct {
	Units       []Unit       `json:"units"`
	Topics      []Topic      `json:"topics"`
	Microtopics []Microtopic `json:"microtopics"`
	Edges       []Edge       `json:"edges"`
}
type Unit struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type Topic struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Parent string `json:"parent"`
}
type Microtopic struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Parent string `json:"parent"`
}
type Edge struct {
	Source string `json:"source,omitempty"`
	Target string `json:"target,omitempty"`
}

type Output struct {
	Data Node `json:"data"`
}
type Node struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Bg     string `json:"bg"`
	Color  string `json:"color"`
	Group  string `json:"group"`
	Parent string `json:"parent"`
	//
	Source string `json:"source,omitempty"`
	Target string `json:"target,omitempty"`
}

func (i Input) ToOutput() []*Output {
	var outputs []*Output

	for _, unit := range i.Units {
		var output Output

		output.Data.Group = "units"
		output.Data.Bg = "rgb(255, 172, 48)"
		output.Data.Name = unit.Name
		output.Data.Id = unit.Id
		output.Data.Color = "white"

		outputs = append(outputs, &output)
	}

	for _, topic := range i.Topics {
		var output Output

		output.Data.Group = "topics"
		output.Data.Bg = "rgb(252, 245, 235)"
		output.Data.Name = topic.Name
		output.Data.Id = topic.Id
		output.Data.Color = "white"
		output.Data.Parent = topic.Parent

		outputs = append(outputs, &output)
	}

	for _, microtopic := range i.Microtopics {
		var container Output
		var data Output

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
		var data Output

		data.Data.Source = edge.Source
		data.Data.Target = edge.Target
		data.Data.Id = fmt.Sprint(rand.Int())

		outputs = append(outputs, &data)
	}

	return outputs
}
