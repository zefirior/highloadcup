#!/bin/bash
uwsgi --http :80 --wsgi-file app.py --callable app --process 4