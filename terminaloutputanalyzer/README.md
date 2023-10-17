## How to Use
- Open terminal
- Move position to the folder ```cd terminaloutputanalyzer```
- Run this command to see the result of the terminaloutputanalyzer.txt
    ```go run main.go```
- To change the input, you can either:
    - Add a new file and change the code on main.go: `textFile, err := src.ReadLines("terminaloutputanalyzer.txt")` change the "terminaloutputanalyzer.txt" into the name of your new file. Note that the file should be on root of the project.
    - Or you can simply modify the `terminaloutputanalyzer.txt` and then run `go run main.go` again.