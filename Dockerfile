
FROM scratch
COPY panic-linter /usr/bin/panic-linter
ENTRYPOINT ["/usr/bin/panic-linter"]
