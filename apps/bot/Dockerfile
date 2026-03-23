FROM scratch

ARG APP_VERSION
ENV APP_VERSION=$APP_VERSION
ENV SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt

COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY fonts /usr/share/fonts
COPY bot /bot

ENTRYPOINT ["/bot"]
