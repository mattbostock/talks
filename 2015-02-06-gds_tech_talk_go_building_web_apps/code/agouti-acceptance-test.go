package main

import (
	"log"
	"os"
	"strings"
	"testing"

	agouti "github.com/sclevine/agouti/core"
)

var (
	driver agouti.WebDriver
	page   agouti.Page
)

// We'd not normally include a main() function in a test suite, but it's
// a necessary hack to show the tests in this example using the Go present tool
func main() {
	t := testing.InternalTest{Name: "TestGOVUKHomepage", F: TestGOVUKHomepage}
	f := func(pat, str string) (result bool, err error) {
		return true, nil
	}
	m := testing.MainStart(f, []testing.InternalTest{t}, []testing.InternalBenchmark{}, []testing.InternalExample{})
	TestMain(m)
}

func TestMain(m *testing.M) {
	var err error

	driver, err = agouti.PhantomJS()
	driver.Start()
	page, err = driver.Page(agouti.Use().Browser("chrome"))

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	result := m.Run()

	driver.Stop()

	os.Exit(result)
}

// START OMIT
func TestGOVUKHomepage(t *testing.T) {
	t.Parallel()

	_ = page.Navigate("https://www.gov.uk/")
	bodyText, _ := page.Find("body").Text()

	if !strings.Contains(bodyText, "Simpler, clearer, faster") {
		t.Error("Missing body text")
	}
}

// END OMIT
