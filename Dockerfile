FROM golang:1.22.3-bookworm as build

WORKDIR app
COPY . .
RUN go mod tidy
RUN go build
ENTRYPOINT ./tigrinhobet
#
# FROM ubuntu
# COPY --from=build /bin/app tigrinhobet
# RUN ls
# RUN chmod +x tigrinhobet
# CMD ["sleep", "30000000"]
