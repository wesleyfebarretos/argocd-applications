FROM golang:1.26.3-alpine AS build

WORKDIR /src

COPY go.mod main.go ./

RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /server .

FROM gcr.io/distroless/static-debian13:nonroot

COPY --from=build --chown=nonroot:nonroot /server /server

ENV PORT=8080

EXPOSE 8080

ENTRYPOINT ["/server"]
