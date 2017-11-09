FROM debian:stretch-slim

RUN apt-get install mariadb-server \
  sed -i 's/^\(bind-address\s.*\)/# \1/' /etc/mysql/my.cnf && \
  echo "mysqld_safe &" > /tmp/config && \
  echo "mysqladmin --silent --wait=30 ping || exit 1" >> /tmp/config && \
  echo "mysql -e 'GRANT ALL PRIVILEGES ON *.* TO \"root\"@\"%\" WITH GRANT OPTION;'" >> /tmp/config && \
  bash /tmp/config && \
  rm -f /tmp/config

VOLUME ["/etc/mysql", "/var/lib/mysql"]

WORKDIR /home

CMD ["mysqld_safe"]

ENTRYPOINT ["/bin/bash"]

EXPOSE
3306
