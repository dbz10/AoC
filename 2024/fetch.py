import tomllib
import requests
import logging
from os import environ

INPUT_URL_TEMPLATE = "https://adventofcode.com/{year}/day/{day}/input"
PUZZLE_URL_TEMPLATE = "https://adventofcode.com/{year}/day/{day}"


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


def get_example(year: int, day: int) -> None:
    """
    Ok, so the example usually follows right after the occurrence of the words "For example:"
    """
    # check whether sample is already populated. for now just assume if there's any content,
    # it has been correctly retrieved
    target_file = f"day{day:02d}/sample.txt"
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

    target_url = PUZZLE_URL_TEMPLATE.format(year=year, day=day)
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

    resp = requests.get(target_url, cookies={"session": cookie})
    resp.raise_for_status()

    puzzle_input = resp.body
    print(puzzle_input)
    # with open(target_file, "w") as f:
    #     f.write(puzzle_input)


if __name__ == "__main__":
    get_example(2023, 10)
