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
	"os"
	"strconv"

	"github.com/jackzampolin/invasion/invasion"
	"github.com/spf13/cobra"
)

var numCities int

// newMapCmd represents the newMap command
var newMapCmd = &cobra.Command{
	Use:   "newMap [name]",
	Short: "Generates a new map file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mapFile := fmt.Sprintf("%s.map", args[0])
		if _, err := os.Stat(mapFile); err == nil {
			fmt.Printf("Map file %s already exists, try another file name...\n", mapFile)
			return
		}
		m := invasion.NewMap()
		fmt.Printf("Generating %d random city names of %d length...\n", numCities, invasion.CityNameLength)
		cities := m.CityBytes(numCities, invasion.CityNameLength)
		f, _ := os.OpenFile(mapFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		for i, c := range cities {
			if i%1000 == 0 {
				fmt.Printf("\r%s percent of cities written to %s...", strconv.FormatFloat((float64(i)/float64(numCities))*100, 'f', 2, 64), mapFile)
			}
			city := m.NewCityBytes(string(c), cities)
			f.Write(city.Bytes())
		}
		fmt.Printf("\nCreated map file %s in current directory...\n", mapFile)
	},
}

func init() {
	rootCmd.AddCommand(newMapCmd)
	newMapCmd.Flags().IntVar(&numCities, "numCities", 10, "A number of cites to have in the generated map")
}
