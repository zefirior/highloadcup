import sys
import json
import pathlib as pl


def is_account_source(file_name):
    return file_name[:9] == "accounts_" and file_name[-5:] == ".json"


def read_accounts(path):
    with open(path, "r") as accounts_file:
        objects = json.load(accounts_file)
        print(objects['accounts'][0])


data_dir = pl.Path(sys.argv[1])

for file in data_dir.iterdir():
    if file.is_file() and is_account_source(file.name):
        read_accounts(file)
