FROM registry.svc.ci.openshift.org/openshift/release:golang-1.13 AS builder
WORKDIR /go/src/github.com/openshift/subscription-injection-webhook
COPY . .
RUN make build

FROM registry.svc.ci.openshift.org/origin/4.6:base
COPY --from=builder /go/src/github.com/openshift/subscription-injection-webhook/_output/subscription-injection-webhook /usr/bin/
ENTRYPOINT []
CMD ["/usr/bin/subscription-injection-webhook"]