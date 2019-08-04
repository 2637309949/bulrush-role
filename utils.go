// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package role

import "strings"

// TransformAction action tran
func TransformAction(action string) []Action {
	actStrs := strings.Split(action, ";")
	actions := make([]Action, 0)
	for _, act := range actStrs {
		rpStr := strings.Split(act, "@")
		if len(rpStr) == 1 {
			actions = append(actions, Action{Roles: strings.Split(rpStr[0], ",")})
		} else if len(rpStr) == 2 {
			actions = append(actions, Action{Roles: strings.Split(rpStr[0], ","), permits: strings.Split(rpStr[1], ",")})
		}
	}
	return actions
}
