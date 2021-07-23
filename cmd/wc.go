/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
    "os"
    "bufio"
    "strings"
)

var bytes, chars, lines, max_line_length, words, help bool
var help_doc = `Usage: wc [OPTION]... [FILE]...
  or:  wc [OPTION]... --files0-from=F
Print newline, word, and byte counts for each FILE, and a total line if
more than one FILE is specified.  A word is a non-zero-length sequence of
characters delimited by white space.

With no FILE, or when FILE is -, read standard input.

The options below may be used to select which counts are printed, always in
the following order: newline, word, character, byte, maximum line length.
  -c, --bytes        	print the byte counts
  -m, --chars        	print the character counts
  -l, --lines        	print the newline counts
  	--files0-from=F	read input from the files specified by
                       	NUL-terminated names in file F;
                       	If F is - then read names from standard input
  -L, --max-line-length  print the maximum display width
  -w, --words        	print the word counts
  	--help 	display this help and exit
  	--version  output version information and exit

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/wc>
or available locally via: info '(coreutils) wc invocation'
`
// wcCmd represents the wc command
var wcCmd = &cobra.Command{
	Use:   "wc",
	Short: "A brief description of your command",
	Long: help_doc,
	Run: func(cmd *cobra.Command, args []string) {
        if help == true{
            fmt.Println(help_doc)
            return
        }
        var fileName string
        fileName = args[0]
        if bytes{
            fi, err := os.Stat(fileName)
            if err != nil {
                return
            }

            fmt.Println(fi.Size())
            return
        }
         file, err := os.Open(fileName)
         if err != nil {
             fmt.Println("Err ", err)
         }
         scanner := bufio.NewScanner(file)
         liness, wordss, characters, max := 0, 0, 0, 0
         for scanner.Scan() {
             liness++

             line := scanner.Text()
             characters += len(line)
             if len(line) > max{
                 max = len(line)
            }

             splitLines := strings.Split(line, " ")
             wordss += len(splitLines)
         }
         if chars{
             fmt.Println(characters, fileName)
        }else if max_line_length{
            fmt.Println(max, fileName)
        }else if words{
            fmt.Println(wordss, fileName)
        }else if lines{
            fmt.Println(liness, fileName)
        }else{
            fmt.Println("WRONG COMMAND !")
            fmt.Println(help_doc)
        }
	},
}

func init() {
	rootCmd.AddCommand(wcCmd)

    wcCmd.Flags().BoolVarP(&bytes, "bytes", "c", false, "usage")
    wcCmd.Flags().BoolVarP(&chars, "chars", "m", false, "usage")
    wcCmd.Flags().BoolVarP(&max_line_length, "max-line-length", "L", false, "usage")
    wcCmd.Flags().BoolVarP(&words, "words", "w", false, "usage")
    wcCmd.Flags().BoolVarP(&lines, "lines", "l", false, "usage")
    wcCmd.Flags().BoolVarP(&help, "help", "h", false, "usage")
}