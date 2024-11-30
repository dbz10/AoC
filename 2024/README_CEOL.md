# Summary
Christmas Elf Officer Lite `ceol.py` (name is shamelessly stolen from the far more real project https://github.com/gcalmettes/christmas-elf-officer) is a small CLI tool to facilitate running Advent of Code.

The functionality is not extensive: it will download and cache your puzzle input (provided with a cookie to authenticate yourself) and run your solution for that day. 

# Usage
To use, first install `uv`: `curl -LsSf https://astral.sh/uv/install.sh | sh`

You also need a session cookie from your logged in session. Place this either in an environment variable `AOC_SESSION_COOKIE` or in a file `env.toml` under key `session_cookie`. For example
```toml
# env.toml
session_cookie = "..."
```

Then, run a particular day with `uv run ceol.py`
```
uv run ceol.py --help
Reading inline script metadata from `ceol.py`

 Usage: ceol.py [OPTIONS] DAY

 Run your solution for a given day. Will attempt to download the puzzle input, which requires a
 cookie for authentication, unless `fetch_input` is set to false (alternatively, can pass as a
 flag `--no-fetch-input`).

╭─ Arguments ────────────────────────────────────────────────────────────────────────────────────╮
│ *    day      INTEGER  [default: None] [required]                                              │
╰────────────────────────────────────────────────────────────────────────────────────────────────╯
╭─ Options ──────────────────────────────────────────────────────────────────────────────────────╮
│ --input                              [sample|input]  [default: input]                          │
│ --fetch-input    --no-fetch-input                    [default: fetch-input]                    │
│ --help                                               Show this message and exit.               │
╰────────────────────────────────────────────────────────────────────────────────────────────────╯

```

For example `uv run ceol.py 1 --input sample` to run Day 1 on the sample input, or `uv run ceol.py 1` to run Day 1 on the real puzzle input.
