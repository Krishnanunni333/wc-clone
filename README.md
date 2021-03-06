# wc command using golang
A clone of wc Unix command developed using go & cobra.

Syntax <br>
$ app wc [OPTION]... [FILE]...

**File name is must**

Options
1. -l or --lines
This option prints the number of lines present in a file. With this option wc command displays two-columnar output, 1st column shows number of lines present in a file and 2nd itself represent the file name.

2. -w or --words
This option prints the number of words present in a file. With this option wc command displays two-columnar output, 1st column shows number of words present in a file and 2nd is the file name.

3. -c or --bytes
This option displays count of bytes present in a file. With this option it display two-columnar output, 1st column shows number of bytes present in a file and 2nd is the file name.

4. -m or --chars
Using -m option ‘wc’ command displays count of characters from a file.

5. -L or --max-line-length
The ‘wc’ command allow an argument -L, it can be used to print out the length of longest (number of characters) line in a file.

6. –h or --help
This option is used to display the help message.

# Example
$ app wc -m apple.txt </br>
  27 apple.txt

# Testing
$ go test -v

```console
=== RUN   TestWc
=== RUN   TestWc/counting_words_in_a_file
=== RUN   TestWc/counting_lines_in_a_file
=== RUN   TestWc/counting_characters_in_a_file
=== RUN   TestWc/counting_bytes_in_a_file
=== RUN   TestWc/counting_the_max_line_length_in_a_file
--- PASS: TestWc (0.00s)
    --- PASS: TestWc/counting_words_in_a_file (0.00s)
    --- PASS: TestWc/counting_lines_in_a_file (0.00s)
    --- PASS: TestWc/counting_characters_in_a_file (0.00s)
    --- PASS: TestWc/counting_bytes_in_a_file (0.00s)
    --- PASS: TestWc/counting_the_max_line_length_in_a_file (0.00s)
PASS
ok      app     0.050s
foo
```


