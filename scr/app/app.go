package app

import "syscall"

type App struct {
	WindowList []Win
}

func NewApp() App {
	return App{}
}

func (app App) Close() {
	syscall.Exit(0)
}
