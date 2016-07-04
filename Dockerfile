FROM scratch
ADD linux-bin/docker-go-scratchhw /docker-go-scratchhw
EXPOSE 8080
ENTRYPOINT [ "/docker-go-scratchhw"]
