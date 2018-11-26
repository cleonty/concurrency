package main

import "context"

func Open() *APIConnection {
	return &APIConnection{}
}

type APIConnection struct {
}

func (a *APIConnection) ReadFile(ctx context.Context) error {
	return nil
}

func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	return nil
}
