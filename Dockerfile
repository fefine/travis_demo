FROM ubuntu

WORKDIR /app

COPY build/travis_demo .

EXPOSE 8080

CMD ["./travis_demo"]