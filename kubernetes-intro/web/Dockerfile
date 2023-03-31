FROM golang:1.20.2-alpine3.17
COPY ./golang-server /src/golang-server
WORKDIR /app
COPY ./app .
ARG UID=1001
ARG USERNAME=system-user
RUN adduser --system --uid "${UID}" "${USERNAME}"
USER ${USERNAME}
EXPOSE 8000
ENTRYPOINT [ "go", "run" ]
CMD [ "/src/golang-server/server.go" ]
