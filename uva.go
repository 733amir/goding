package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

const (
	help = `UVa script is an assistant to work with UVa online problems.
UVa usage:
	./uva get PNUM [DIRPATH]
		It will download problem description and test cases of problem number PNUM to directory DIRPATH. If DIRPATH is 
		not specified it will download the necessary files to current working directory.
	./uva -h
		It will print this help message.`
	rawPdfUrl    = "https://uva.onlinejudge.org/external/%d/p%d.pdf"
	rawUDebugUrl = "https://www.udebug.com/UVa/%d"
)

type uDebugInput struct {
	Value string `json:"input_value"`
}

func main() {
	cmdArgs := os.Args[1:]
	if len(cmdArgs) == 0 {
		fmt.Println(help)
		os.Exit(1)
	}

	switch strings.ToLower(cmdArgs[0]) {
	case "get":
		cmdArgs = cmdArgs[1:]
		if len(cmdArgs) == 0 {
			fmt.Println(help)
			os.Exit(1)
		}
		problemNumber, err := strconv.Atoi(cmdArgs[0])
		if err != nil {
			fmt.Println("UVa: The problem number should be an integer.")
			os.Exit(1)
		}

		cmdArgs = cmdArgs[1:]
		var dirPath string
		if len(cmdArgs) != 0 {
			dirPath = cmdArgs[0]
			pathInfo, err := os.Stat(dirPath)
			if err != nil {
				// FIXME: The error should not be printed directly. Use the error to print more informative message for user.
				fmt.Println(err)
				os.Exit(1)
			}
			if !pathInfo.IsDir() {
				fmt.Println("UVa: The specified path is not pointing to a directory.")
				os.Exit(1)
			}
			if mode := pathInfo.Mode().String(); mode[2] != 'w' && mode[5] != 'w' && mode[8] != 'w' {
				fmt.Println("UVa: The specified directory is not writable.")
				os.Exit(1)
			}
		}

		err = getDescription(problemNumber, dirPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = getTestCases(problemNumber, dirPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println(help)
		os.Exit(1)
	}
}

func getDescription(problemNumber int, dirPath string) error {
	switch {
	case 100 <= problemNumber && problemNumber <= 1760:
		pdfUrl := fmt.Sprintf(rawPdfUrl, problemNumber/100, problemNumber)
		DownloadFile(path.Join(dirPath, fmt.Sprintf("%05d.pdf", problemNumber)), pdfUrl)
		return nil
	case 10000 <= problemNumber && problemNumber <= 13303:
		pdfUrl := fmt.Sprintf(rawPdfUrl, problemNumber/100, problemNumber)
		DownloadFile(path.Join(dirPath, fmt.Sprintf("%05d.pdf", problemNumber)), pdfUrl)
		return nil
	}

	return errors.New("UVa: problem number is not supported")
}

func getTestCases(problemNumber int, dirpath string) error {
	testCaseIds := make([]int, 0, 5)
	var problemNID, formBuildID string

	// Finding all test cases for this problem.
	c := colly.NewCollector()
	c.OnHTML("a[data-id]", func(anchor *colly.HTMLElement) {
		if anchor.ChildText("span") != "" {
			return
		}

		id, err := strconv.Atoi(anchor.Attr("data-id"))
		if err != nil {
			return
		}

		testCaseIds = append(testCaseIds, id)
	})
	c.OnHTML("input[name='problem_nid']", func(input *colly.HTMLElement) {
		problemNID = input.Attr("value")
	})
	c.OnHTML("form[id='udebug-custom-problem-view-input-output-form'] input[name='form_build_id']", func(input *colly.HTMLElement) {
		formBuildID = input.Attr("value")
	})
	err := c.Visit(fmt.Sprintf(rawUDebugUrl, problemNumber))
	if err != nil {
		return err
	}

	// Downloading all input cases for the problem.
	inputCases := make([]string, len(testCaseIds))
	expectedCases := make([]string, len(testCaseIds))
	for i := range testCaseIds {
		// Get input of test case
		c := colly.NewCollector()
		c.OnResponse(func(response *colly.Response) {
			temp := uDebugInput{}
			err = json.Unmarshal(response.Body, &temp)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			inputCases[i] = temp.Value
		})
		c.Post(
			"https://www.udebug.com/udebug-custom-get-selected-input-ajax",
			map[string]string{
				"input_nid": strconv.Itoa(testCaseIds[i]),
			},
		)

		// Get output of test case
		c = colly.NewCollector()
		c.OnHTML("textarea[id='edit-output-data']", func(textarea *colly.HTMLElement) {
			expectedCases[i] = textarea.Text
		})
		c.OnHTML("form[id='udebug-custom-problem-view-input-output-form'] input[name='form_build_id']", func(input *colly.HTMLElement) {
			formBuildID = input.Attr("value")
		})
		err := c.Post(
			fmt.Sprintf(rawUDebugUrl, problemNumber),
			map[string]string{
				"problem_nid":   problemNID,
				"input_data":    inputCases[i],
				"form_build_id": formBuildID,
				"form_id":       "udebug_custom_problem_view_input_output_form",
			},
		)
		if err != nil {
			return err
		}
	}

	for i := range inputCases {
		input, err := os.Create(path.Join(dirpath, fmt.Sprintf("input-%d", i)))
		if err != nil {
			return err
		}
		_, err = input.WriteString(inputCases[i])
		if err != nil {
			return err
		}

		output, err := os.Create(path.Join(dirpath, fmt.Sprintf("expected-%d", i)))
		if err != nil {
			return err
		}
		_, err = output.WriteString(expectedCases[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filePath string, url string) error {

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
