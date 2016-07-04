FROM scratch
ADD linux-bin/docker-go-scratchhw /docker-go-scratchhw
ENTRYPOINT [ "/docker-go-scratchhw"]
