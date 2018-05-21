// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"fmt"
	"strconv"

	"github.com/jackzampolin/invasion/invasion"
	"github.com/spf13/cobra"
)

const numTurns = 10000

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play [mapFile path] [numAliens]",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Read in the map file and create a *Map from it
		m, err := invasion.NewMapFromFile(args[0])
		if err != nil {
			fmt.Println("Failed loading map:", err)
			fmt.Println("Try using the newMap command to create a new map")
			return
		}

		// Parse out the desired number of aliens
		numAliens, err := strconv.ParseInt(args[1], 10, 16)
		if err != nil {
			fmt.Println("numAliens is not a valid integer:", args[1])
			return
		}

		// Add the aliens to the map
		m.NewAliens(int(numAliens))

		// Play the game
		m.Play()

		fmt.Printf("\nENDING MAP (%d):\n\n%s\nENDING ALIENS:\n\n", len(m.Cities), m)
		for _, a := range m.Aliens {
			fmt.Println(a)
		}
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
