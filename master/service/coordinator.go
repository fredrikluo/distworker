package service

import "context"

type Task struct {
	TaskId  string
	Agentid string
	TaskCmd string
}

type HearBeatData struct {
	TaskId        string
	AgentId       string
	ProcessedTime int64
}

type ReportData struct {
}

type Coordinator struct {
	ToDoTask      []Task
	FinishedTask  []Task
	InProcessTask map[string]Task
}

func NewCoordinator() *Coordinator {
	return &Coordinator{}
}

func (co *Coordinator) AddTask(ctx context.Context, tasks []Task) error {
	return nil
}

func (co *Coordinator) SendInReport(ctx context.Context, report ReportData) error {
	return nil
}

func (co *Coordinator) GetReport(ctx context.Context) (ReportData, error) {
	return ReportData{}, nil
}

func (co *Coordinator) GetTask(ctx context.Context) (Task, error) {
	return Task{}, nil
}

func (co *Coordinator) HeartBeat(ctx context.Context, heartBeat HearBeatData) (Task, error) {
	return Task{}, nil
}
