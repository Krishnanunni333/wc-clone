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
    W "app/cmd/workers"
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
        switch {

            case help:
                fmt.Println(help_doc)

            case bytes:
                fmt.Println(W.PrintBytes(args[0]), args[0])

            case chars:
                fmt.Println(W.CountChars(args[0]), args[0])

            case max_line_length:
                fmt.Println(W.MaxLine(args[0]), args[0])

            case words:
                fmt.Println(W.WordsInLines(args[0]), args[0])

            case lines:
                fmt.Println(W.CountLines(args[0]), args[0])

            default:
                fmt.Println("WRONG COMMAND !")
                fmt.Println(help_doc)
        }
	},
}



func init() {
	rootCmd.AddCommand(wcCmd)

    wcCmd.Flags().BoolVarP(&bytes, "bytes", "c", false, "Display the size of the file")
    wcCmd.Flags().BoolVarP(&chars, "chars", "m", false, "Display the number of characters")
    wcCmd.Flags().BoolVarP(&max_line_length, "max-line-length", "L", false, "Display the length of the line having maximum length ")
    wcCmd.Flags().BoolVarP(&words, "words", "w", false, "Display number of words")
    wcCmd.Flags().BoolVarP(&lines, "lines", "l", false, "Display numbe rof lines")
    wcCmd.Flags().BoolVarP(&help, "help", "h", false, "Help")
}
