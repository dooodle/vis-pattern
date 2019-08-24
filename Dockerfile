FROM scratch
COPY server /
CMD ["/server","-serve"]
