FROM postgres:9.4

ADD init-shippy-db.sh /docker-entrypoint-initdb.d/
RUN chmod a+x /docker-entrypoint-initdb.d/init-shippy-db.sh