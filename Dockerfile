FROM scratch
COPY {{.applicationName}} /
ENTRYPOINT ["/{{.applicationName}}"]
