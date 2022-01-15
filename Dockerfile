FROM golang:1.17-alpine AS builder

WORKDIR /app/
COPY . /app/

RUN CGO_ENABLED=0 go build -o /bin/app -mod=mod
ADD ./IPV6-COUNTRY.SAMPLE.BIN /app/IP-COUNTRY-REGION-CITY.BIN
ENV IPDB=/app/IP-COUNTRY-REGION-CITY.BIN
ENTRYPOINT ["/bin/app"]
ARG TOKEN
ARG DB_URL=https://www.ip2location.com/download/?file=DB3LITEBINIPV6&token=$TOKEN
RUN if [ ! -z $TOKEN ]; then \
    curl -o db.zip $DB_URL && \
    unzip db.zip && \
    mv IP2LOCATION-LITE-DB3.IPV6.BIN /app/var/IP-COUNTRY-REGION-CITY.BIN;fi
