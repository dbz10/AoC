# /// script
# requires-python = ">=3.11"
# dependencies = [
#     "requests",
#     "typer",
#     "numpy",
# ]
# ///

import typer
import importlib
from enum import StrEnum, auto

from bench import bench
from fetch import get_input
import solution_performance

YEAR = 2024


class InputType(StrEnum):
    sample = auto()
    input = auto()


def main(
    day: int,
    input: InputType = InputType.input,
    fetch_input: bool = True,
    benchmark: bool = False,
) -> None:
    """
    Run your solution for a given day. Will attempt to download the puzzle input, which requires a cookie for authentication,
    unles `fetch_input` is set to false (alternatively, can pass as a flag `--no-fetch-input`).
    """
    if fetch_input:
        get_input(YEAR, day)
    target_pkg = f"day{day:02d}.main"
    target_file = f"day{day:02d}/{input.name}.txt"
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
