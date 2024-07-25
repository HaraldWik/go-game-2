package app

import "syscall"

type App struct {
	WindowList []Win
}

func New() App {
	return App{}
}

func (app App) Close() {
	syscall.Exit(0)
}
