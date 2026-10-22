// SPDX-FileCopyrightText: © 2023 W-A-T EU Operations Oü
// SPDX-License-Identifier: MIT
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package cliapp

import (
	"fmt"
)

// Someone who has contributed to a project.
type Author struct {
        Name  string
        Email string
}

// Implement the Stringer interface
func (a Author) String() string {
	mail := ""  
	if "" != a.Email {
		mail = fmt.Sprintf(" <%s>", a.Email)
	}
	return a.Name+mail
}
