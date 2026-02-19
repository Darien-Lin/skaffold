ARG BASE_PREFIX=mirror.gcr.io/library/
ARG BASE=busybox

FROM ${BASE} as parent

FROM ${BASE_PREFIX}alpine
COPY --from=parent /app .
CMD ["./app"]