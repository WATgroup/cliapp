// SPDX-FileCopyrightText: © 2023 W-A-T EU Operations Oü
// SPDX-License-Identifier: MIT
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package cliapp

import (
	"fmt"
	"github.com/WATgroup/errors"
	"os"
	"path/filepath"
)

const (
	msgApp_NotInit = "Application object not initialized properly"
)

type AppInfo struct {
	Name      string
	Version   string
	GitTag    string
	Copyright string
	Email     string
}

type App struct {
	Name     string
	info     AppInfo
	cmder    Commander
	authors  []Author
	metadata map[string]interface{}
	init     bool
}

func New(i AppInfo) *App {

	var n string

	if n = i.Name; "" == n {
		n = filepath.Base(os.Args[0])
	}

	return &App{
		init: true,
		Name: n,
		info: i,
	}
}

func (a *App) SetCommander(c Commander) error {
	if !a.init {
		return errors.New(msgApp_NotInit)
	}
	a.cmder = c
	return setupCobra(a, c)
}

func (a *App) SetMetadata(mm map[string]interface{}) error {
	if !a.init {
		return errors.New(msgApp_NotInit)
	}
	a.metadata = mm
	return nil
}

func (x *App) SetAuthors(aa ...Author) {

	// if nil == x.authors {
	// 	x.authors = make([]Author,0,1)
	// }

	for _, a := range aa {
		x.authors = append(x.authors, a)
	}
}

func (a *App) Run() { // NORETURN

	if !a.init {
		//		fmt.Println("FATAL: Application object not initialized properly")
		panic(msgApp_NotInit)
		// os.Exit(127)
	}
	r := os.Args

	if 2 == len(r) {
		switch r[1] {
		case "-V":
			fallthrough
		case "-version":
			fallthrough
		case "--version":
			printVersionInfo(a, os.Stdout)
			os.Exit(0)
		case "-h":
			if hf, ok := (a.cmder).(interface{ HelpFunc(*command, []string) }); ok {
				hf.HelpFunc((a.cmder).(*command), []string{})
			}
			os.Exit(0)
		}
	}

	if err := a.cmder.Execute(); err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
