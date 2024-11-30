# /// script
# requires-python = ">=3.11"
# dependencies = [
#     "requests",
#     "typer",
# ]
# ///

import typer
import importlib
from enum import StrEnum, auto
import tomllib
import requests
import logging
from os import environ
from bench import bench

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


# SAMPLE_SELECTOR = "body > main > article:nth-child(1) > pre:nth-child(8) > code"  how to do this consistently across different days?
INPUT_URL_TEMPLATE = "https://adventofcode.com/{year}/day/{day}/input"


def get_input(year: int, day: int) -> None:
    """
    It seemed messy to get the sample input since it's not obvious what the consistent way to fetch it would be.
    For now, content myself with fetching the real input. Maybe I'll come back to this.
    """
    # check whether input is already populated. for now just assume if there's any content,
    # it has been correctly retrieved
    target_file = f"day{day:02d}/input.txt"
    try:
        with open(target_file, "r") as f:
            lines = len(f.readlines())
            if lines > 0:
                logging.info(
                    "Skipping fetching input since it seems to already be materialized. To force re-fetching it, remove the target file: `rm %s`",
                    target_file,
                )
                return
    except FileNotFoundError:
        pass

    input_url = INPUT_URL_TEMPLATE.format(year=year, day=day)

    # We need a session cookie from somewhere
    cookie = environ.get("AOC_SESSION_COOKIE")
    if cookie is None:
        try:
            with open("env.toml", "rb") as f:
                cookie = tomllib.load(f)["session_cookie"]
        except FileNotFoundError:
            pass
    if cookie is None:
        raise ValueError(
            "AoC session cookie was not found in either `AOC_SESSION_COOKIE` environment variable or env.toml::session_cookie."
        )

    resp = requests.get(input_url, cookies={"session": cookie})
    resp.raise_for_status()

    puzzle_input = resp.text
    with open(target_file, "w") as f:
        f.write(puzzle_input)


if __name__ == "__main__":
    typer.run(main)
