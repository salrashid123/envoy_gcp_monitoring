FROM envoyproxy/envoy-dev:2e6db8378477a4a63740746c5bfeb264cd76bc34
COPY envoy.yaml /etc/envoy/envoy.yaml
WORKDIR /app
ADD certs /certs
EXPOSE 9000
EXPOSE 443
RUN chmod -R go+r /certs

RUN chmod go+r /etc/envoy/envoy.yaml
