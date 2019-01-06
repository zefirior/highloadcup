import socket

def mess_pack(message: bytes):
    length_prefix = "{} ".format(len(message))
    return length_prefix.encode() + message


sock_file = "/tmp/server-socket.sock"
sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
sock.connect(sock_file)

def get_message(sock):
    resp = sock.recv(1024).decode()
    while " " not in resp:
        resp += sock.recv(1024).decode()
    length, resp = resp.split(" ", 1)
    length = int(length)
    while length > len(resp):
        resp += sock.recv(1024).decode()
    return resp

import time

try:
    start = time.time()
    for i in range(1):
        sock.send(mess_pack(b'AADD ' + str(i).encode() + b'q' * 1200))
        resp = get_message(sock)
        # print(i, resp)
    print(time.time() - start)
except Exception as e:
    print(i)
    raise e

print(i, resp)
