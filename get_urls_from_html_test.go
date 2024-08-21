package main

import (
	"fmt"
	"golang.org/x/net/html"
	"reflect"
	"strings"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<html>
					<body>
						<a href="/path/one">
							<span>Boot.dev</span>
						</a>
						<a href="https://other.com/path/one">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
				`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "deeply embedded html tree ",
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<html>
					<body>
						<div class="block"> 
							<p>Some text</p>
							<div>
								<a href="/path/one">
									<span>Boot.dev</span>
								</a>
							</div>
							<a href="https://other.com/path/one">
								<span>Boot.dev</span>
							</a>
						</div>
					</body>
				</html>
				`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "another urls",
			inputURL: "https://google.com",
			inputBody: `
				<html>
					<body>
						<a href="/search">
							<span>Google</span>
						</a>
						<a href="https://blog.boot.dev">
							<span>Boot.dev Blog</span>
						</a>
					</body>
				</html>
				`,
			expected: []string{"https://google.com/search", "https://blog.boot.dev"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if reflect.DeepEqual(actual, tc.expected) == false {
				t.Errorf("Test %v - '%s' FAIL: actual result not equal to expected result: %v - %v", i, tc.name, actual, tc.expected)
				return
			}
		})
	}
}

func TestTraverseNodes(t *testing.T) {
	htmlReader := strings.NewReader(`
				<html>
					<body>
						<a href="/path/one">
							<span>Boot.dev</span>
						</a>
						<a href="https://other.com/path/one">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
				`)
	htmlTree, err := html.Parse(htmlReader)
	if err != nil {
		fmt.Println(err)
	}

	htmlReader1 := strings.NewReader(`
				<html>
					<body>
						<div class="block"> 
							<p>Some text</p>
							<div>
								<a href="/path/one">
									<span>Boot.dev</span>
								</a>
							</div>
							<a href="https://other.com/path/one">
								<span>Boot.dev</span>
							</a>
						</div>
					</body>
				</html>
				`)
	htmlTree1, err := html.Parse(htmlReader1)
	if err != nil {
		fmt.Println(err)
	}

	htmlReader2 := strings.NewReader(`
				<html>
					<body>
						<div class="block"> 
							<p>Some text</p>
							<div>
							</div>
						</div>
					</body>
				</html>
				`)
	htmlTree2, err := html.Parse(htmlReader2)
	if err != nil {
		fmt.Println(err)
	}

	tests := []struct {
		name           string
		deep           int
		result         *[]string
		inputBody      *html.Node
		expectedResult *[]string
	}{
		{
			name:           "check correct getting href atribut",
			deep:           0,
			result:         &[]string{},
			inputBody:      htmlTree,
			expectedResult: &[]string{"/path/one", "https://other.com/path/one"},
		},
		{
			name:           "deeply embedded html tree",
			deep:           0,
			result:         &[]string{},
			inputBody:      htmlTree1,
			expectedResult: &[]string{"/path/one", "https://other.com/path/one"},
		},
		{
			name:           "without links",
			deep:           0,
			result:         &[]string{},
			inputBody:      htmlTree2,
			expectedResult: &[]string{},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			traverseNodes(tc.inputBody, tc.deep, tc.result)
			if reflect.DeepEqual(tc.result, tc.expectedResult) == false {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
		})
	}
}
