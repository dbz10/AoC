import sqlite3

from bench import BENCHDB, TABLE


def main():
    con = sqlite3.connect(BENCHDB)
    cur = con.cursor()

    result = cur.execute(f"""
        with solutions_ranked as (
        select 
            *, 
            rank() over (partition by year, day order by duration_seconds) as rnk from {TABLE}
        )
        select year, day, duration_seconds, commit_hash, run_ts
        from solutions_ranked
        where rnk = 1
    """)

    columns = [d[0] for d in result.description]
    data = result.fetchall()
    data_stringified = [[str(v) for v in d] for d in data]

    markdown_header = "| " + " | ".join(columns) + " |"
    header_separator = "| " + " | ".join(["---"] * len(columns)) + " |"
    markdown_rows = "\n".join(
        ["| " + " | ".join(row_data) + " |" for row_data in data_stringified]
    )

    markdown_table = f"""{markdown_header}
    {header_separator}
    {markdown_rows}"""

    with open("solutions.md", "w") as f:
        f.write(markdown_table)


if __name__ == "__main__":
    main()
