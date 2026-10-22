// SPDX-FileCopyrightText: © 2023 W-A-T EU Operations Oü
// SPDX-License-Identifier: MIT
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package cliapp

import (
	"github.com/spf13/cobra"
	"github.com/WATgroup/errors"
	"os"
)



type command = cobra.Command

type Commander interface {
	Execute() error
	AddCommand(...*cobra.Command)
}




// Designed to integrate with spf13/cobra
func setupCobra(a *App, c Commander) error {

	if nil == c {
		return errors.New("Invalid 'commander' given)")
	}

	c.AddCommand(&cobra.Command{
		Use: "version",
		Short: "display version info",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			// 			c.HelpFunc()(c,[]string{})
			printExtendedInfo(a, os.Stdout)
			os.Exit(0)
		},
	})
	return nil
}
