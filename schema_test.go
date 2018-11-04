package s2gen

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

// TestSchemaGeneration tests, whether we emit a runnable code snippet.
func TestSchemaGeneration(t *testing.T) {
	t.Logf("(1/5) building s2gen executable")
	cmd := exec.Command("go", "build", "-o", "s2gen", "cmd/s2gen/main.go")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	// Generate snippet from schema and store in a temporary file.
	schema := "fixtures/schema.xml"
	t.Logf("(2/5) generating code from %s", schema)

	cmd = exec.Command("./s2gen")

	f, err := os.Open(schema)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	cmd.Stdin = f

	fo, err := ioutil.TempFile("", "s2gen-test-*.go")
	if err != nil {
		t.Fatal(err)
	}

	cmd.Stdout = fo
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	t.Logf("(3/5) temporary snippet written to %s", fo.Name())

	// Running on sample input.
	sample := "fixtures/docs01.ndj"

	t.Logf("(4/5) running on input %s", sample)
	cmd = exec.Command("go", "run", fo.Name())

	f, err = os.Open(sample)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	cmd.Stdin = f
	cmd.Stdout = ioutil.Discard
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	// If we succeed, remove temporary files.
	fo.Close()
	os.Remove(fo.Name())

	t.Logf("(5/5) ok")
}
