# Rendering regression samples

Each `<name>_w<width>.md` is rendered by `TestRegression` (`regression_test.go`)
the way the CLI renders a file — dark style, word wrap at `<width>` — and
compared, with ANSI sequences removed, against `<name>_w<width>.golden`.

```sh
go test -run TestRegression .            # check
go test -run TestRegression -update .    # regenerate the goldens after an intended change
```

A golden equals the piped output of `glow -s dark -w <width> <file>` with its
ANSI sequences removed, so a sample can also be checked by hand.

## What each sample guards

Every sample was checked to fail when the listed fix is reverted, and to pass
with all fixes applied.

| Sample | Fix (fork commit) | Without the fix |
|---|---|---|
| `quote_indent_width_w80` | glamour `34d9f73` ansi: subtract the display width of indent tokens from the block width | a wrapped piece of the quote loses its `│` prefix |
| `hyphen_wrap_w40` | reflow `336f7d8` wordwrap: count breakpoint runes toward the line limit | hyphenated words overflow the width |
| `table_content_height_w60` | lipgloss `0ae2b3c` table: detect content height with cellbuf.Wrap | the last line of a tall cell (`/ ircomm_hidden`) is dropped |
| `table_wrapped_rows_w60` | glamour `344ecca` ansi: separate table rows with a rule when any row wraps | no rule between wrapped rows, or `├`/`┤` at the rule ends |
| `table_wrapped_rows_w60` | lipgloss `c564bd1` table: count row separators only between rows | an extra blank line below the table |
| `table_header_rule_w80` | glamour `cd66a6d` ansi: draw the table header rule with ═ and ╪ | the header rule is drawn with `─`/`┼` like the row rules |
| `table_full_width_w80` | lipgloss `0cc14f4` table: count only enabled outer borders in the resizer | the table is two cells narrower than the width |
| `table_overflow_w120` | lipgloss `09058b5` table: include borders when choosing to expand or shrink | the right edge is cut off: `(3 절)` renders as `(3 절` |
| `table_emoji_vs16_w80` | glow: strip U+FE0F before rendering | the `│` after `ℹ️`/`⏭️` sits one cell left of the header's |

The table samples also exercise the other table fixes, so reverting one table
fix usually fails several table samples; the sample listed above is the one
written for it.
