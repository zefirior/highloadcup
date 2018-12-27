import sys
import json
import logging
import pathlib as pl
from contextlib import contextmanager

from mysql.connector import connection

logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)

log_handler = logging.StreamHandler()
log_handler.setLevel(logging.DEBUG)
log_formatter = logging.Formatter("%(levelname)s:%(message)s")
log_handler.setFormatter(log_formatter)

logger.addHandler(log_handler)


def connect():
    return connection.MySQLConnection(
        user='python_app', password='password',
        host='127.0.0.1', port=3306,
        database='highload'
    )


@contextmanager
def after_close(context):
    yield context
    context.close()


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


def load_account(cursor, indict: dict):
    query = """
      insert into account (id, fname, sname, email, phone, sex, status, birth, joined, country, city)
        values (
          %(id)s, %(fname)s, %(sname)s, %(email)s, %(phone)s, %(sex)s, 
          %(status)s, %(birth)s, %(joined)s, %(country)s, %(city)s
        )
    """

    res_dict = {key: indict[key] if key in indict else None
                for key in ["id", "fname", "sname", "email", "phone",
                            "birth", "joined", "country", "city"]}
    for code in ["sex", "status"]:
        res_dict[code] = value_map[code][indict[code]]
    cursor.execute(query, res_dict)


def load_interest(cursor, indict: dict):
    query = """
      insert into interest (account_id, interest)
        values (%s, %s)
    """

    if "interests" not in indict:
        return

    account_id = indict["id"]
    cursor.executemany(query, [(account_id, interest) for interest in indict["interests"]])


def load_premium(cursor, indict: dict):
    query = """
      insert into premium (account_id, start, finish)
        values (%s, %s, %s)
    """

    if "premium" not in indict:
        return

    premium = indict["premium"]
    cursor.execute(query, (indict["id"], premium["start"], premium["finish"]))


def load_likes(cursor, indict: dict):
    query = """
      insert into likes (account_id, ts, like_account_id)
        values (%s, %s, %s)
    """

    if "likes" not in indict:
        return

    likes = indict["likes"]
    account = indict["id"]
    params = [(account, like['ts'], like['id']) for like in likes]
    cursor.executemany(query, params)


def load_json_2_db(jobj):
    handlers = [
        load_account,
        load_interest,
        load_premium,
        load_likes
    ]

    with after_close(connect()) as cnx:
        with after_close(cnx.cursor()) as cursor:
            for acc_data in jobj['accounts']:
                try:
                    for handler in handlers:
                        handler(cursor, acc_data)
                except Exception as e:
                    logger.info(str(acc_data))
                    raise e
            cnx.commit()


def is_account_source(file_name):
    return file_name[:9] == "accounts_" and file_name[-5:] == ".json"


def read_data(path):
    with open(path, "r") as accounts_file:
        jobj = json.load(accounts_file)
        load_json_2_db(jobj)


logger.info('Start parse')

try:
    data_dir = pl.Path(sys.argv[1])
    for file in data_dir.iterdir():
        if file.is_file() and is_account_source(file.name):
            read_data(file)
except Exception as e:
    logger.error(str(e))
else:
    logger.info('End parse')
