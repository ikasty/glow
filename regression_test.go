package main

import (
	"flag"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"charm.land/glamour/v2"
	"charm.land/glow/v3/utils"
	"github.com/charmbracelet/x/ansi"
)

// Rendering regression tests for the fixes carried by this fork and by the
// ikasty forks of glamour and lipgloss (see testdata/regression/README.md).
//
// Each testdata/regression/<name>_w<width>.md is rendered the way the CLI
// renders a file (dark style, word wrap at <width>) and compared, with ANSI
// sequences stripped, against <name>_w<width>.golden. A golden equals
// the piped output of `glow -s dark -w <width> <name>_w<width>.md` with its
// ANSI sequences removed.
// Regenerate the goldens with: go test -run TestRegression -update
var updateRegression = flag.Bool("update", false, "update regression golden files")

var widthSuffix = regexp.MustCompile(`_w(\d+)\.md$`)

// renderLikeCLI mirrors the renderer options of executeCLI for a markdown
// file. glamour v2 always emits styled output and leaves color downsampling
// to the writer, so the result does not depend on the terminal. The style
// sequences reach the table cells and some width bugs only show up with
// them, so they are stripped only after rendering.
func renderLikeCLI(t *testing.T, md []byte, width int) string {
	t.Helper()
	r, err := glamour.NewTermRenderer(
		utils.GlamourStyle("dark", false),
		glamour.WithWordWrap(width),
		glamour.WithBaseURL(""),
		glamour.WithPreservedNewLines(),
	)
	if err != nil {
		t.Fatal(err)
	}
	out, err := r.Render(utils.RemoveEmojiVariationSelector(string(utils.RemoveFrontmatter(md))))
	if err != nil {
		t.Fatal(err)
	}
	return ansi.Strip(out)
}

func TestRegression(t *testing.T) {
	files, err := filepath.Glob(filepath.Join("testdata", "regression", "*.md"))
	if err != nil {
		t.Fatal(err)
	}
	if len(files) == 0 {
		t.Fatal("no regression samples found")
	}
	for _, f := range files {
		if filepath.Base(f) == "README.md" {
			continue
		}
		m := widthSuffix.FindStringSubmatch(f)
		if m == nil {
			t.Fatalf("%s: file name must end with _w<width>.md", f)
		}
		width, _ := strconv.Atoi(m[1])
		name := strings.TrimSuffix(filepath.Base(f), ".md")

		t.Run(name, func(t *testing.T) {
			md, err := os.ReadFile(f)
			if err != nil {
				t.Fatal(err)
			}
			got := renderLikeCLI(t, md, width)

			golden := strings.TrimSuffix(f, ".md") + ".golden"
			if *updateRegression {
				if err := os.WriteFile(golden, []byte(got), 0o644); err != nil { //nolint:gosec
					t.Fatal(err)
				}
				return
			}
			want, err := os.ReadFile(golden)
			if err != nil {
				t.Fatalf("%v (run with -update to create it)", err)
			}
			if got != string(want) {
				t.Errorf("rendering of %s changed\n--- want\n%s\n--- got\n%s", f, visible(string(want)), visible(got))
			}
		})
	}
}

// visible marks line ends so that trailing spaces show up in failure output.
func visible(s string) string {
	return strings.ReplaceAll(s, "\n", "⏎\n")
}
