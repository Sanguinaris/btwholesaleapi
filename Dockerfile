FROM golang AS build

WORKDIR /source
COPY . ./
RUN mkdir /app && go build -o /app/api  

FROM busybox:1

WORKDIR /app
COPY --from=build /app ./
ENTRYPOINT ["api"]
