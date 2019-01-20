import socket

BUFFER_SIZE = 1024


class CoreClient:

    def __init__(self, address: str):
        self.address = address
        self.fd = None

    def connect(self):
        if self.fd is not None:
            raise Exception("already connected")

        self.fd = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
        self.fd.connect(self.address)

    def _send(self, data):
        if isinstance(data, str):
            data = data.encode()
        self.fd.send(self._mess_pack(data))

    def _recv(self):
        resp = self.fd.recv(BUFFER_SIZE).decode()
        while " " not in resp:
            resp += self.fd.recv(BUFFER_SIZE).decode()
        length, resp = resp.split(" ", 1)
        length = int(length)
        while length > len(resp):
            resp += self.fd.recv(BUFFER_SIZE).decode()
        return resp

    @staticmethod
    def _mess_pack(message: bytes):
        length_prefix = "{} ".format(len(message))
        return length_prefix.encode() + message


class ApiCore(CoreClient):
    def _rpc_method(self, method, data=""):
        self._send("{} {}".format(method, data))
        return self._recv()

    def method_AADD(self, data):
        return self._rpc_method('AADD', data)

    def method_CACC(self):
        return self._rpc_method("CACC")

    def method_PLST(self):
        return self._rpc_method("PLST")


if __name__ == "__main__":
    acc_data = "id 123 b 223445345 fn asdgkajdew p 42 235 s 0 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 2345 678683 st 1 ph 8(912)6290012 e alegfbhhh@afdh.ru sn qweasd co country ci city_new"

    sock_file = "/tmp/server-socket.sock"
    core = ApiCore(sock_file)
    print("connect")
    core.connect()
    print("connected")

    import time
    start = time.time()
    for _ in range(1002):
        if core.method_AADD(acc_data) != "ok":
            raise Exception

    print(time.time() - start)
    print(core.method_CACC())
    print(core.method_PLST())
