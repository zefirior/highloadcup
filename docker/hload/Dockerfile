FROM pymysql

WORKDIR /app

### app
COPY ./requirements.txt ./requirements.txt
COPY ./src ./src
COPY ./db ./db
COPY ./hload-entrypoint.sh ./hload-entrypoint.sh

RUN chmod 777 ./hload-entrypoint.sh

RUN pip install -U pip && pip install -r requirements.txt

EXPOSE 80 3306

#ENTRYPOINT /bin/bash

RUN apt-get install -y unzip

CMD ["./hload-entrypoint.sh"]
