package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"
)

func main() {
	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Add wait group to wait for all examples to finish
	var wg sync.WaitGroup

	for _, entry := range entries {
		if !entry.IsDir() {
			filename := entry.Name()
			if isTestFile(filename) {
				// Add to wait group
				wg.Add(1)

				// go func() {

				// Process the file in a goroutine
				fmt.Println("Processing file:", filename)
				if err := processTestFile(filename); err != nil {
					fmt.Println("Error processing test file:\n", err)
					return // Stop processing further if any example fails
				}

				fmt.Println()

				// time.Sleep(1 * time.Second)

				// Done with the file
				wg.Done()

				// }()
			}
		}
	}

	// Wait for all examples to finish
	wg.Wait()
}

func isTestFile(filename string) bool {
	return regexp.MustCompile(`_test\.go$`).MatchString(filename)
}

func findExampleFunctions(filename string) ([]string, error) {
	// ReRead the content after each example
	fileContentBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", filename, err)
	}
	fileContent := string(fileContentBytes)

	var examples []string
	r := regexp.MustCompile(`func Example[^(]*\(\)`)
	matches := r.FindAllString(fileContent, -1)

	for _, match := range matches {
		exampleName := match[5:] // Remove "func " prefix to get the example function name
		examples = append(examples, exampleName)
	}

	return examples, nil
}

func processTestFile(filename string) error {
	examples, err := findExampleFunctions(filename)
	if err != nil {
		return fmt.Errorf("error finding example functions in file %s: %w", filename, err)
	}

	// Test the filename as a hole first
	// go test -run ^(strings.Join(examples, "|"))$
	TestPassCmd := exec.Command("go", "test", "-run", "^("+strings.Join(examples, "|")+")$")
	_, err = TestPassCmd.CombinedOutput()
	if err == nil {
		fmt.Println("Success")
		// fmt.Printf("file %s passed successfully on first run\n", filename)
		return nil
	}

	for _, example := range examples {
		// ReRead the content after each example
		content, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("error reading file %s: %w", filename, err)
		}

		if err := runAndUpdateExample(filename, example, string(content)); err != nil {
			return err // Return the error to stop further processing
		}

		// Give some space in the output
		fmt.Println()
	}
	return nil
}

func runAndUpdateExample(filename, exampleName, content string) error {
	fmt.Println("Testing: ", exampleName)
	firstRunCmd := exec.Command("go", "test", "-run", exampleName+"$")
	firstOutput, err := firstRunCmd.CombinedOutput()
	if err == nil {
		fmt.Println("Success")
		// fmt.Printf("example %s in file %s passed successfully on first run\n", exampleName, filename)
		return nil
	}

	// fmt.Printf("firstOutput: %s\n", string(firstOutput))

	// If the example failed, output the error and continue
	fmt.Println("Failed, running update")
	// fmt.Printf("example %s in file %s failed on first run: %s\n", exampleName, filename, err)

	// Parse the output to get "got" and "want"
	got, want := parseOutputForGotAndWant(string(firstOutput))
	if got == "" || want == "" {
		return fmt.Errorf("example %s in file %s failed, but unable to parse got/want: %s", exampleName, filename, err)
	}

	// fmt.Printf("Got:\n%s\n\nWant:\n%s\n", got, want)

	// Update the file with "got" as new output
	newContent, err := updateExampleOutputInFileContent(content, exampleName, got)
	if err != nil {
		return fmt.Errorf("error updating example %s in file %s: %s", exampleName, filename, err)
	}
	if newContent == "" {
		return fmt.Errorf("failed to update example %s in file %s with new output", exampleName, filename)
	}

	if err := os.WriteFile(filename, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("error writing updated file %s: %s", filename, err)
	}

	// time.Sleep(1 * time.Second)

	// Rerun the example after updating
	secondRunCmd := exec.Command("go", "test", "-run", exampleName+"$")
	secondOutput, err := secondRunCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("example %s in file %s failed after update: %s\nOutput:\n%s", exampleName, filename, err, string(secondOutput))
	}

	fmt.Println("Success after update")
	// fmt.Printf("example %s in file %s passed successfully after update\n", exampleName, filename)
	return nil
}

func parseOutputForGotAndWant(output string) (got, want string) {
	gotRegex := regexp.MustCompile(`(?s)got:\n(.*?)\nwant:`)
	wantRegex := regexp.MustCompile(`(?s)want:\n(.*)`)

	gotMatch := gotRegex.FindStringSubmatch(output)
	wantMatch := wantRegex.FindStringSubmatch(output)

	if len(gotMatch) > 1 {
		got = gotMatch[1]
	}
	if len(wantMatch) > 1 {
		want = wantMatch[1]
	}

	return strings.TrimSpace(got), strings.TrimSpace(want)
}

func updateExampleOutputInFileContent(content, exampleName, newOutput string) (string, error) {
	lines := strings.Split(content, "\n")
	foundExample := false
	foundOutputMarker := false
	exampleStart := fmt.Sprintf("func %s", exampleName) // func ExampleCusip()
	fmt.Println(exampleStart)
	outputMarker := "// Output:"

	// Prepare the new output lines, each prefixed with "// "
	outputLines := strings.Split(newOutput, "\n")
	for i, line := range outputLines {
		// Add // Output:  prefix to the first line
		if i == 0 {
			// Starting with [tab]// Output:
			outputLines[i] = "	" + outputMarker + " " + line
			continue
		}

		// Starting with [tab]//[space]
		outputLines[i] = "	// " + line
	}

	startComment := 0
	endComment := 0
	for i, line := range lines {
		// If you have started the comment figure out the end
		if foundExample && foundOutputMarker {
			// Check if current row starts with // if so continue
			if strings.HasPrefix(strings.TrimSpace(line), "//") {
				continue
			} else {
				// If we are here we are at the end of the comment
				endComment = i
				break
			}
		}

		// Found example function
		if strings.Contains(line, exampleStart) {
			foundExample = true
		}

		// Found output marker
		if foundExample && strings.Contains(line, outputMarker) {
			// Mark that we found the output marker and start inserting new output lines
			foundOutputMarker = true
			startComment = i

			continue
		}
	}

	if !foundExample {
		return "", errors.New(exampleName + " not found in example function")
	}

	if !foundOutputMarker {
		return "", errors.New("output marker not found in example function")
	}

	// If startComment and endComment are 0 then error
	if startComment == 0 && endComment == 0 {
		return "", errors.New("start and end comment not found")
	}

	// fmt.Printf("Start: %d, End: %d\n\n", startComment, endComment)

	// Loop through the lines deleting the old output
	// and inserting the new output
	lines = append(lines[:startComment], append(outputLines, lines[endComment:]...)...)

	return strings.Join(lines, "\n"), nil
}
