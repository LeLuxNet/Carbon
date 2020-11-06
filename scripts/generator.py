import pathlib
import re
from typing import Dict, List

HEADER = """/* WARNING!
   This file is automatically generated through scripts/generator.py
   DON'T EDIT IT MANUALLY. ALL CHANGES WILL BE LOST! */

"""

value = List[List[str]]


def _set_var(var: Dict[str, value], match) -> str:
    var[match.group(1)] = [[b.strip() for b in a.split(":")] for a in match.group(2).split(",")]
    return ""


def process(text: str) -> str:
    parts = re.split(r"\n#split\n", text)
    global_res = HEADER
    for part in parts:
        var = {}
        res = re.sub(r"\n#def ([a-z]) (.+)\n", lambda x: _set_var(var, x), part)
        for key, values in var.items():
            round_res = ""
            for v in values:
                value_res = res
                for i, a in enumerate(v):
                    value_res = value_res.replace("<%s%d>" % (key, i + 1), a)
                round_res += value_res
            res = round_res
        global_res += res
    return global_res


def process_file(path: pathlib.Path):
    src = path.read_text()

    res = process(src)

    dst = path.with_suffix('')
    dst.write_text(res)

    print("Generated", dst.relative_to('.'))


for path in pathlib.Path('.').rglob('*.template'):
    process_file(path)