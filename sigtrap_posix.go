// Copyright 2015 Light Code Labs, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !windows && !plan9 && !nacl && !js
// +build !windows,!plan9,!nacl,!js

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// trapSignalsPosix captures POSIX-only signals.
func trapSignalsPosix() {
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

		for sig := range sigchan {
			switch sig {
			case syscall.SIGQUIT:
				fmt.Println("[INFO] SIGQUIT: Quitting process immediately")
				serverInstance.Stop()
				os.Exit(0)
			case syscall.SIGTERM:
				fmt.Println("[INFO] SIGTERM: Shutting down servers then terminating")
				exitCode := 3
				err := serverInstance.Stop()
				if err != nil {
					fmt.Printf("[ERROR] SIGTERM stop: %v", err)
					exitCode = 3
				}
				os.Exit(exitCode)
			case syscall.SIGUSR1:
				fmt.Println("[INFO] SIGQUIT: SIGUSR1")
			case syscall.SIGUSR2:
				fmt.Println("[INFO] SIGQUIT: SIGUSR2")
			default:
				fmt.Println("[INFO] SIGQUIT: default")
			}
		}
	}()
}
