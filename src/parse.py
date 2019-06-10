import sys
import json
import logging
import pathlib as pl

from core_api import ApiCore

logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)

log_handler = logging.StreamHandler()
log_handler.setLevel(logging.DEBUG)
log_formatter = logging.Formatter("%(levelname)s:%(message)s")
log_handler.setFormatter(log_formatter)

logger.addHandler(log_handler)


value_map = {
    "sex": {
        "m": 1,
        "f": 0,
    },
    "status": {
        "свободны": 0,
        "заняты": 1,
        "всё сложно": 2,
    }
}


def is_account_source(file_name):
    return file_name[:9] == "accounts_" and file_name[-5:] == ".json"


def read_data(path):
    with open(path, "r") as accounts_file:
        return json.load(accounts_file)


def encode_account(obj: dict):
    result = []
    if "status" in obj:
        value = obj["status"]
        result += ["st", str(value_map["status"][value])]

    if "sex" in obj:
        value = obj["sex"]
        result += ["s", str(value_map["sex"][value])]

    attr_table = (
        ("birth", "b"),
        ("email", "e"),
        ("country", "co"),
        ("city", "ci"),
        ("id", "id"),
        ("phone", "ph"),
        ("fname", "fn"),
        ("sname", "sn"),
        ("joined", "jo"),
    )
    for key, marker in attr_table:
        if key in obj:
            result += [marker, str(obj[key])]

    if "likes" in obj:
        result += ["l", str(len(obj["likes"]))]
        for like in obj["likes"]:
            result += [str(like["id"]), str(like["ts"])]

    if "interests" in obj:
        for interest in obj["interests"]:
            result += ["in", interest]

    if "premium" in obj:
        premium = obj["premium"]
        result += ["p", str(premium["start"]), str(premium["finish"])]

    return "|".join(result)

logger.info('Start parse')

sock_file = "/tmp/server-socket.sock"
core = ApiCore(sock_file)
print("connect")
core.connect()
print("connected")

import time

try:
    print("path=", sys.argv[1])
    data_dir = pl.Path(sys.argv[1])
    start = time.time()
    for _ in range(1500000 // 10000 // 3):
        for file in data_dir.iterdir():
            if file.is_file() and is_account_source(file.name):
                jobj = read_data(file)
                for account in jobj["accounts"]:
                    account_data = encode_account(account)
                    core.method_AADD(account_data)
                print(core.method_CACC())
                print(core.method_PLST())
    print(time.time() - start)

except Exception as e:
    logger.error(str(e))
else:
    logger.info('End parse')


from collections import Counter
import string


class Item:
    def __init__(self, ch, num, flag):
        self.ch = ch
        self.num = num
        self.flag = flag

    @staticmethod
    def filter(item):
        return item.num > 1

    def __str__(self):
        return "{}:{}".format(self.flag, self.ch * self.num)


def construct_item(ch, max1, max2):
    if max1 == max2:
        return Item(ch, max1, "=")
    return [Item(ch, max1, "1"), Item(ch, max2, "2")][max1 < max2]


def mix(s1, s2):
    c1, c2 = Counter(s1), Counter(s2)
    data = [construct_item(ch, c1.get(ch, 0), c2.get(ch, 0)) for ch in string.ascii_lowercase]
    data = filter(Item.filter, data)
    data = sorted(data, key=lambda item: str(item))
    data = sorted(data, key=lambda item: item.num, reverse=True)
    return "/".join(str(item) for item in data)
