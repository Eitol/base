import os
import re
import sys

GO_MOD_FILE_NAME = "go.mod"
GO_FILE_EXTENSION = ".go"
DEFAULT_ROOT_DIR = "./"
MIN_NAME_LEN = 2
MAX_NAME_LEN = 30


def change_module_name_in_file(dir_: str, file_name: str, current_name: str, new_name: str) -> None:
    if not file_name.endswith(GO_FILE_EXTENSION) and file_name != GO_MOD_FILE_NAME:
        return
    file_path = "{}/{}".format(dir_, file_name)
    with open(file_path, "r+")as f:
        content = f.read()
        if file_name == GO_MOD_FILE_NAME:
            content_new = content.replace("module {}".format(current_name), "module {}".format(new_name))
        elif file_name.endswith(GO_FILE_EXTENSION):
            content_new = content.replace("\"{}/".format(current_name), "\"{}/".format(new_name))
        else:
            return
        if content == content_new:
            return
        f.seek(0)
        f.write(content_new)
        print("Changed file:{}".format(file_path))


def get_module_name(root_path_: str) -> str:
    file = "{}/{}".format(root_path_, GO_MOD_FILE_NAME)
    try:
        with open("{}/{}".format(root_path_, GO_MOD_FILE_NAME), "r") as f:
            line = f.readline()
            if line.startswith("module "):
                return line.strip().split()[1]
    except FileNotFoundError:
        print("ERROR: not {} file in {}".format(GO_MOD_FILE_NAME, file))
        exit(1)
    return ""


def change_module_name_in_dir(dir_: str, current_name: str, new_name: str):
    change_module_name_in_file(dir_, GO_MOD_FILE_NAME, current_name, new_name)
    for root, dirs, files in os.walk("{}/{}".format(dir_, "app"), topdown=True):
        for file in files:
            change_module_name_in_file(root, file, current_name, new_name)


def validate_new_name(new_name: str) -> (bool, str):
    """ return the validity of the name and the error core if exist  """

    if len(new_name) < MIN_NAME_LEN:
        return False, "Min len must be > {}".format(MIN_NAME_LEN)
    if len(new_name) > MAX_NAME_LEN:
        return False, "Max len must be < {}".format(MAX_NAME_LEN)
    pattern = r'[a-z][a-z1-9_/]{0,{MAX-1}}[a-z1-9]{{MIN},{MAX-1}}$' \
        .replace("{MIN}", str(MIN_NAME_LEN)) \
        .replace("{MAX}", str(MAX_NAME_LEN)) \
        .replace("{MAX-1}", str(MAX_NAME_LEN - 1))
    result = re.match(pattern, new_name)
    if not result:
        return False, "invalid module_name".format(MAX_NAME_LEN)
    return True, ""


def search_go_mod() -> str:
    search_path = [DEFAULT_ROOT_DIR, "."]
    for path in search_path:
        try:
            if os.path.isfile(path):
                return path
            print("Not {} found in {}".format(GO_MOD_FILE_NAME, path))
        except FileExistsError:
            pass


def get_root_path() -> str:
    if len(sys.argv) >= 3:
        return sys.argv[2]
    elif len(sys.argv) == 2:
        return search_go_mod()
    else:
        return ""


if __name__ == '__main__':
    if len(sys.argv) == 1:
        print("ERROR: new name is empty")
        exit(1)
    root_path = get_root_path()
    new_name_ = sys.argv[1]
    if new_name_ == "invalid":
        print("ERROR: invalid name")
        print("Usage:   make init new_name=NEW_NAME")
        print("Example: make init new_name=payment_api")
        exit(1)
    is_valid_name, error = validate_new_name(new_name_)
    if not is_valid_name:
        print("Invalid module name '{}'. Error: {}", new_name_, error)
        exit(1)
    current_name_ = get_module_name(DEFAULT_ROOT_DIR)
    if current_name_ == new_name_:
        print("Is the same module name '{}'. Error: {}", new_name_)
        exit(1)
    print("Current module name == {}".format(current_name_))
    change_module_name_in_dir(DEFAULT_ROOT_DIR, current_name_, new_name_)
