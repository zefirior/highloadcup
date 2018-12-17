import sys
import json
import pathlib as pl

from mysql.connector import connection
cnx = connection.MySQLConnection(user='root', password='',
                                 host='127.0.0.1',
                                 database='')


# def get


def is_account_source(file_name):
    return file_name[:9] == "accounts_" and file_name[-5:] == ".json"


def read_accounts(path):
    with open(path, "r") as accounts_file:
        objects = json.load(accounts_file)
        with
        for acc in objects['accounts']:



data_dir = pl.Path(sys.argv[1])

for file in data_dir.iterdir():
    if file.is_file() and is_account_source(file.name):
        read_accounts(file)
