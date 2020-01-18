// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package compiler

import (
	"os"
	"strings"

	"github.com/drone-runners/drone-runner-kube/engine"
	"github.com/drone-runners/drone-runner-kube/engine/compiler/shell"
	"github.com/drone-runners/drone-runner-kube/engine/resource"
)

// helper function configures the pipeline script for the
// target operating system.
func (c *Compiler) setupScript(src *resource.Step, dst *engine.Step, isService bool) {
	if len(src.Commands) == 0 && len(src.Entrypoint) == 0 && !isService {
		src.Commands = []string{getCommand(src.Image)}
	}
	if len(src.Commands) > 0 {
		setupScriptPosix(c.envCommands, src.Commands, dst)
	}

	if len(src.Entrypoint) > 0 {
		cmds := []string{
			strings.Join(append(src.Entrypoint, src.Command...), " "),
		}
		setupScriptPosix(c.envCommands, cmds, dst)
	}
}

// helper function configures the pipeline script for the
// linux operating system.
func setupScriptPosix(before func() string, commands []string, dst *engine.Step) {
	dst.Entrypoint = []string{"sh", "-c"}
	// dst.Command = []string{`echo "$DRONE_SCRIPT" | sh`}
	dst.Command = []string{"sleep 7200"}
	dst.Envs["DRONE_SCRIPT"] = shell.Script(before, append(commands, "kill $heartid"))
}

func getCommand(image string) string {
	temp := getImageName(image)
	temp = strings.ReplaceAll(temp, "-", "_")
	temp = strings.ToUpper(temp)
	command := os.Getenv("STEP_" + temp)
	if command != "" {
		return command
	}
	return getImageName(image)
}

func getImageName(image string) string {
	noTagImage := strings.Split(image, ":")[0]
	split := strings.Split(noTagImage, "/")
	return split[len(split)-1]
}
