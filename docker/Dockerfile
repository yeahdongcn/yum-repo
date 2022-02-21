FROM golang:1.17.6 as builder

WORKDIR /go/src/github.com/yeahdongcn/yum-repo
ADD . .
RUN make build

FROM centos:7
LABEL maintainer="R0CKSTAR <yeahdongcn@gmail.com>"

RUN yum -y install epel-release && \
    yum -y update && \
    yum -y install supervisor createrepo yum-utils nginx && \
    yum clean all

RUN mkdir /repo && \
    chmod 777 /repo && \
    mkdir -p /logs

COPY config/nginx.conf /etc/nginx/nginx.conf
COPY config/supervisord.conf /etc/supervisord.conf
COPY entrypoint.sh /root/
COPY --from=builder /go/src/github.com/yeahdongcn/yum-repo/bin/yum-repo /root/

RUN chmod 700 /root/yum-repo
RUN chmod 700 /root/entrypoint.sh

ENV SERVE_FILES true

EXPOSE 80

ENTRYPOINT ["/root/entrypoint.sh"]