# Summary
Christmas Elf Officer Lite `ceol.py` (name is shamelessly stolen from the far more real project https://github.com/gcalmettes/christmas-elf-officer) is a small CLI tool to facilitate running Advent of Code.

The functionality is not extensive: it will download and cache your puzzle input (provided with a cookie to authenticate yourself) and run your solution for that day. 

(Functionality to automatically get the sample input is still WIP, for now only fetching the real puzzle input is suppported.)

# Usage
CEOL expects to be used from the directory where `ceol.py` lives, and assumes your AoC project is
1. implemented in python
2. laid out with each day as a package, e.g. 
```
tree
.
├── README_CEOL.md
├── __init__.py
├── ceol.py
├── day01
│   ├── __init__.py
│   ├── main.py
│   └── ...
├── day02
│   ├── __init__.py
│   ├── main.py
│   └── ...
├── ...
```
3. Each day's package implements a `main` function which takes a single input argument providing the path to a file containing the puzzle input, and outputs something (probably the answers to parts 1 and 2) to ideally stdout. 

Concretely, besides downloading the puzzle input, it simply imports `dayxx.main` and runs `dayxx.main::main(input_file)`.

CEOL is simplest to run using `uv`, if needed first install it with `curl -LsSf https://astral.sh/uv/install.sh | sh`. Add any dependencies your solutions use to the requirements for CEOL, for example `uv add --script ceol.py 'numpy'` to add a numpy dependency. `uv` will recompile the virtual environemnt on each run if needed.

To run without `uv`, prepare a virtual environment with `typer`, `requests`, and whatever dependencies your solutions need. Then instead of `uv run ceol.py` as shown below, simply `python ceol.py`.

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
