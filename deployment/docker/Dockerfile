# Set the base image
FROM alpine:3.3

# RUN apk --update upgrade && apk add ca-certificates && update-ca-certificates
COPY ./kubelet.linux /bin/kubelet 

ENTRYPOINT ["/bin/kubelet"]
