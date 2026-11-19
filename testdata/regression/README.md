# Rendering regression samples

Each `<name>_w<width>.md` is rendered by `TestRegression` (`regression_test.go`)
the way the CLI renders a file — dark style, word wrap at `<width>` — and
compared, with ANSI sequences removed, against `<name>_w<width>.golden`.

```sh
go test -run TestRegression .            # check
go test -run TestRegression -update .    # regenerate the goldens after an intended change
```

A golden equals the piped output of `glow -s dark -w <width> <file>` with its
ANSI sequences removed, so a sample can also be checked by hand. The one
exception is `link_korean_anchor_w60`: the CLI resolves a relative link against
the file's directory, while the test renders with an empty base URL. glamour v2
pads inline code with NO-BREAK SPACE (U+00A0), so goldens contain it.

## What each sample guards

The fixes live on `v3.0.0-custom` of this fork and on the ikasty forks of
glamour (`v2.0.1-custom`) and lipgloss (`v2.0.6-custom`). Each was checked
by reverting it alone: the samples in the last column then fail, and all
samples pass with every fix applied.

| Fix (fork commit) | Without the fix | Samples that fail |
|---|---|---|
| glow `fda1d36` utils: strip U+FE0F before rendering | the `│` after `ℹ️`/`⏭️` sits one cell left of the header's | `table_emoji_vs16_w80` |
| glamour `18c5465` ansi: subtract the display width of indent tokens from the block width (glamour#560) | a wrapped piece of the quote loses its `│` prefix | `quote_indent_width_w80` |
| glamour `e1d66fc` ansi: separate table rows with a rule when any row wraps | no rule between wrapped rows | `table_wrapped_rows_w60`, `table_content_height_w60`, `table_full_width_w80`, `table_overflow_w120` |
| glamour `f8929c7` ansi: draw the table header rule with ═ and ╪ | the header rule is drawn with `─`/`┼` like the row rules | every table sample |
| glamour `e61ca49` ansi: do not emit OSC 8 hyperlinks | a Korean anchor leaks out of the broken OSC 8 sequence | `link_korean_anchor_w60` |
| lipgloss `1554ff6` table: include borders when choosing to expand or shrink | the right edge is cut off: `(3 절)` renders as `(3 절` | `table_overflow_w120` |

Two lipgloss fork commits are kept but no sample fails without them under v3:

- `bb2c0ac` table: count row separators only between rows — lipgloss v2 uses
  the computed height only as a maximum, so the extra row it counted no
  longer shows up as a blank line.
- `f61a447` detect content height with the same Wrap the cell renderer uses.

Samples that guarded v2.1.1-custom fixes not carried over:

- `hyphen_wrap_w40` — the reflow fork (breakpoint runes count toward the line
  limit). Upstream reflow v0.3.0 renders it correctly under v3.
- `table_full_width_w80` — lipgloss `0cc14f4` (count only enabled outer
  borders) is upstream in v2 (baa7611).
