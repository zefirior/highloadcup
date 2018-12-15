FROM python:3.6

WORKDIR /src


### supervizord

RUN apt-get update && apt-get install -y supervisor
RUN mkdir -p /var/log/supervisor
COPY ./supervisord.conf /etc/supervisor/conf.d/supervisord.conf


### mysql
RUN apt-cache search mysql-server
RUN apt-get install -y mysql-server


### app
COPY ./requirements.txt ./requirements.txt
COPY ./app.py ./app.py

RUN pip install -U pip && pip install -r requirements.txt

EXPOSE 80 3306

CMD /usr/bin/supervisord
