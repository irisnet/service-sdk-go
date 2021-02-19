FROM irisnet/irishub:1.0.0-rc0

COPY . /scripts

RUN sh /scripts/setup.sh

EXPOSE 26657
EXPOSE 9090

CMD iris start