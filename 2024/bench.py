import sqlite3
from time import perf_counter
from datetime import datetime
import subprocess
import functools

BENCHDB = "bench.db"
TABLE = "aoc_bench"


def init_check_db():
    """
    Create the DB if it's the first time running.
    Pray to god the schema is correct from the start and I never want to do a migration lol lmao
    """
    con = sqlite3.connect(BENCHDB)
    cursor = con.cursor()
    table_exists = (
        cursor.execute(
            f"select * from sqlite_master where type = 'table' and name = '{TABLE}'"
        ).fetchone()
        is not None
    )
    if not table_exists:
        cursor.execute(f"""
        CREATE TABLE {TABLE}(
                       year INT,
                       day INT,
                       commit_hash TEXT,
                       run_ts DATETIME,
                       duration_seconds FLOAT
                       )
        """)
    con.close()


def bench(year, day):
    def decorator_bench(func):
        @functools.wraps(func)
        def inner(*args, **kwargs):
            tic = perf_counter()
            func(*args, **kwargs)
            toc = perf_counter()

            commit_hash = (
                subprocess.run(
                    ["git", "rev-parse", "--short", "HEAD"],
                    capture_output=True,
                )
                .stdout.decode()
                .strip()
            )
            duration_seconds = toc - tic

            data = {
                "year": year,
                "day": day,
                "commit_hash": commit_hash,
                "run_ts": datetime.now(),
                "duration_seconds": duration_seconds,
            }
            init_check_db()
            con = sqlite3.connect(BENCHDB)
            cursor = con.cursor()
            cursor.execute(
                f"""
                INSERT INTO {TABLE}(year, day, commit_hash, run_ts, duration_seconds) 
                VALUES(:year, :day, :commit_hash, :run_ts, :duration_seconds)
                """,
                data,
            )
            con.commit()  # why do i need to do this...? where is autocommit?!?!
            con.close()

        return inner

    return decorator_bench
