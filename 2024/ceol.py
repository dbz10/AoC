# /// script
# requires-python = ">=3.11"
# dependencies = [
#     "requests",
#     "typer",
#     "numpy",
#     "matplotlib",
# ]
# ///

import typer
import importlib
from enum import StrEnum, auto
from typing_extensions import Annotated

from bench import bench
from fetch import get_input
import solution_performance

YEAR = 2024


class InputType(StrEnum):
    sample = auto()
    input = auto()


def main(
    day: int,
    sample: Annotated[
        bool,
        typer.Option(
            help="Flag to control whether the solution is run against the sample input or the real input. Targets the real input by default."
        ),
    ] = False,
    fetch_input: Annotated[
        bool,
        typer.Option(
            help="Flag to control fetching the input or not. True by default. Input is saved between repeated runs so will still only be fetched once, unless the file is removed. Set to False if e.g. you don't have a session cookie handy"
        ),
    ] = True,
    benchmark: Annotated[
        bool,
        typer.Option(help="Flag to control whether benchmarking will be run or not."),
    ] = False,
) -> None:
    """
    Run your solution for a given day. Will attempt to download the puzzle input, which requires a cookie for authentication,
    unles `fetch_input` is set to false (alternatively, can pass as a flag `--no-fetch-input`).
    """
    if fetch_input:
        get_input(YEAR, day)
    target_pkg = f"day{day:02d}.main"
    if sample:
        target_file = f"day{day:02d}/sample.txt"
    else:
        target_file = f"day{day:02d}/input.txt"
    day_runner = importlib.import_module(target_pkg)

    if benchmark:

        @bench(YEAR, day)
        def run_day():
            day_runner.main(target_file)

        run_day()
    else:
        day_runner.main(target_file)

    # update timing sheet
    solution_performance.main()


if __name__ == "__main__":
    typer.run(main)
