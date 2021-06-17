FROM golang:alpine AS build-env
RUN apk --no-cache add build-base gcc
ADD . /src
RUN cd /src && go build -o file-data-svc

# final stage
FROM alpine
WORKDIR /app
ENV PORT=4001
ENV GIN_MODE=debug
COPY --from=build-env /src/file-data-svc /app/
EXPOSE 4001
CMD ["./file-data-svc"]