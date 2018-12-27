from flask import Flask, request, abort
from http import HTTPStatus

from filter_api.main import ApiParam


app = Flask(__name__)


@app.route("/accounts/filter/", methods=["GET"])
def api_filter():
    limit = request.args["limit"]
    for arg in request.args:
        if arg in ("limit", "query_id"):
            continue
        if not ApiParam.exists_param(arg):
            abort(HTTPStatus.BAD_REQUEST)
    return "Hello world!"


if __name__ == "__main__":
    app.run('0.0.0.0', 8005, debug=True)
