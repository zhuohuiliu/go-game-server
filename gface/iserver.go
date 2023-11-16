package gface

type IServer interface {
	Start()
	Stop()
	Serve()
}
