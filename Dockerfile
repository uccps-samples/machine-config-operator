# THIS FILE IS GENERATED FROM Dockerfile DO NOT EDIT
FROM openshift/golang-builder@sha256:4820580c3368f320581eb9e32cf97aeec179a86c5749753a14ed76410a293d83 AS builder
ENV __doozer=update BUILD_RELEASE=202202160023.p0.g14a1ca2.assembly.stream BUILD_VERSION=v4.10.0 OS_GIT_MAJOR=4 OS_GIT_MINOR=10 OS_GIT_PATCH=0 OS_GIT_TREE_STATE=clean OS_GIT_VERSION=4.10.0-202202160023.p0.g14a1ca2.assembly.stream SOURCE_GIT_TREE_STATE=clean 
ENV __doozer=merge OS_GIT_COMMIT=14a1ca2 OS_GIT_VERSION=4.10.0-202202160023.p0.g14a1ca2.assembly.stream-14a1ca2 SOURCE_DATE_EPOCH=1643743595 SOURCE_GIT_COMMIT=14a1ca2cb91ff7e0faf9146b21ba12cd6c652d22 SOURCE_GIT_TAG=unreleased-master-1241-g14a1ca2c SOURCE_GIT_URL=https://github.com/uccps-samples/machine-config-operator 
WORKDIR /go/src/github.com/uccps-samples/machine-config-operator
COPY . .
# FIXME once we can depend on a new enough host that supports globs for COPY,
# just use that.  For now we work around this by copying a tarball.
RUN make install DESTDIR=./instroot && tar -C instroot -cf instroot.tar .

FROM openshift/ose-base:v4.10.0.20220216.010142
ENV __doozer=update BUILD_RELEASE=202202160023.p0.g14a1ca2.assembly.stream BUILD_VERSION=v4.10.0 OS_GIT_MAJOR=4 OS_GIT_MINOR=10 OS_GIT_PATCH=0 OS_GIT_TREE_STATE=clean OS_GIT_VERSION=4.10.0-202202160023.p0.g14a1ca2.assembly.stream SOURCE_GIT_TREE_STATE=clean 
ENV __doozer=merge OS_GIT_COMMIT=14a1ca2 OS_GIT_VERSION=4.10.0-202202160023.p0.g14a1ca2.assembly.stream-14a1ca2 SOURCE_DATE_EPOCH=1643743595 SOURCE_GIT_COMMIT=14a1ca2cb91ff7e0faf9146b21ba12cd6c652d22 SOURCE_GIT_TAG=unreleased-master-1241-g14a1ca2c SOURCE_GIT_URL=https://github.com/uccps-samples/machine-config-operator 
COPY --from=builder /go/src/github.com/uccps-samples/machine-config-operator/instroot.tar /tmp/instroot.tar
RUN cd / && tar xf /tmp/instroot.tar && rm -f /tmp/instroot.tar
COPY install /manifests
RUN if ! rpm -q util-linux; then yum install -y util-linux && yum clean all && rm -rf /var/cache/yum/*; fi
COPY templates /etc/mcc/templates
ENTRYPOINT ["/usr/bin/machine-config-operator"]

LABEL \
        io.openshift.release.operator="true" \
        name="openshift/ose-machine-config-operator" \
        com.redhat.component="ose-machine-config-operator-container" \
        io.openshift.maintainer.product="OpenShift Container Platform" \
        io.openshift.maintainer.component="Machine Config Operator" \
        release="202202160023.p0.g14a1ca2.assembly.stream" \
        io.openshift.build.commit.id="14a1ca2cb91ff7e0faf9146b21ba12cd6c652d22" \
        io.openshift.build.source-location="https://github.com/uccps-samples/machine-config-operator" \
        io.openshift.build.commit.url="https://github.com/uccps-samples/machine-config-operator/commit/14a1ca2cb91ff7e0faf9146b21ba12cd6c652d22" \
        version="v4.10.0"

