from contextlib import contextmanager

from mysql.connector import connection


@contextmanager
def closer(context):
    yield context
    context.close()


config = {
    "user": "python_app",
    "password": "password",
    "host": "0.0.0.0",
    "database": "accounts",
}
with closer(connection.MySQLConnection(**config)) as cnx:
    with closer(cnx.cursor()) as cursor:
        cursor.execute("select * from test")
        message = repr(cursor.fetchall())
        with open("/app/src/test_message.txt", "+w") as mess_f:
            mess_f.write(message)

