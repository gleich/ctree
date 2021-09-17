FROM golang:1.15

# Meta data:
LABEL maintainer="email@mattglei.ch"
LABEL description="🎄 A Christmas tree right from your terminal!"

# Copying over all the files:
COPY . /usr/src/app
WORKDIR /usr/src/app

CMD ["make", "local-test"]
