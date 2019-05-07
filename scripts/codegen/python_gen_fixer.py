import os
import re
import sys
from typing import List

PYTHON_GEN_FILE_END = "_pb2.py"
r = re.compile("from\s(?P<pkg>[^\s]*)\simport\s(?P<member>[^\s]*)\sas\s(?P<alias>[^\s]*)\n")


def get_files(path: str, extension: str) -> List[str]:
    """ Return all files from a path and filter by extension """
    files = []
    # r=root, d=directories, f = files
    for r, d, f in os.walk(path):
        for file in f:
            if file.endswith(extension):
                files.append(os.path.join(r, file))
    return files


def fix_file(filename: str, root_path: str):
    """
    Fix protoc python generated files
    For example:
    If the generated file file is in api/gen/python and  contains imports like this:
    from api.protos import brand_pb2 as api_dot_protos_dot_brand__pb2
    then only change to import like this:
    from api.gen.python.api.protos import fuel_type_pb2 as api_dot_protos_dot_fuel__type__pb2
    Then you use in your code:
    import api.gen.python.api.protos.fuel_type_pb2 as ft
    """
    root_path = root_path.replace("/", ".")
    with open(filename) as file:
        file_str = file.read()
        match = r.findall(file_str)
        if not match:
            return
        for m in match:
            if "google." in m[0]:
                continue
            to_rep = "from {} import {} as {}".format(m[0], m[1], m[2])
            m0 = root_path + "." + m[0]
            new_str = "from {} import {} as {}".format(m0, m[1], m[2])
            print("NEW: " + new_str)
            file_str = file_str.replace(to_rep, new_str)
    with open(filename, "w") as file:
        file.write(file_str)


if __name__ == '__main__':
    # py_files = get_files("api/gen/python/api/protos", PYTHON_GEN_FILE_END)
    py_files = get_files(sys.argv[1] + "/" + sys.argv[2], PYTHON_GEN_FILE_END)
    for filename_ in py_files:
        fix_file(filename_, sys.argv[1])
