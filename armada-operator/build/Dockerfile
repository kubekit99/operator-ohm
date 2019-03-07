FROM alpine:3.8

ENV OPERATOR=/usr/local/bin/armada-operator \
    USER_UID=1001 \
    USER_NAME=armada-operator

# install operator binary
COPY build/_output/bin/armada-operator ${OPERATOR}

COPY build/bin /usr/local/bin
COPY helm-charts/ /opt/armada/helm-charts/
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
