#!/usr/bin/env bash

# Build the MCO image and push it to the cluster registry,
# then directly patch the deployments/daemonsets to use that image.
# This is generally faster than building a new release image and
# upgrading to it, but also consequently doesn't work the same way as production
# upgrades work.
#
# To use this, you must first run `cluster-push-prep.sh` (once) for your cluster.
#
# Assumptions: You have set KUBECONFIG to point to your cluster.

set -xeuo pipefail

do_build=1
if [ "${1:-}" = "-n" ]; then
    do_build=0
fi

registry=$(oc get -n uccp-image-registry -o json route/image-registry | jq -r ".spec.host")
curl -k --head https://"${registry}" >/dev/null

imgname=machine-config-operator
LOCAL_IMGNAME=localhost/${imgname}:latest
REMOTE_IMGNAME=uccp-machine-config-operator/${imgname}
if [ "${do_build}" = 1 ]; then
    ./hack/build-image
fi
builder_secretid=$(oc get -n uccp-machine-config-operator secret | egrep '^builder-token-'| head -1 | cut -f 1 -d ' ')
secret="$(oc get -n uccp-machine-config-operator -o json secret/${builder_secretid} | jq -r '.data.token' | base64 -d)"

if [[ "${podman:-}" =~ "docker" ]]; then
  imgstorage="docker-daemon:"
else
  imgstorage="containers-storage:"
fi
skopeo copy --dest-tls-verify=false --dest-creds unused:${secret} "${imgstorage}${LOCAL_IMGNAME}" "docker://${registry}/${REMOTE_IMGNAME}"

digest=$(skopeo inspect --creds unused:${secret} --tls-verify=false docker://${registry}/${REMOTE_IMGNAME} | jq -r .Digest)
imageid=${REMOTE_IMGNAME}@${digest}

oc project uccp-machine-config-operator

IN_CLUSTER_NAME=image-registry.uccp-image-registry.svc:5000/${imageid}

# Scale down the operator now to avoid it racing with our update.
oc scale --replicas=0 deploy/machine-config-operator

# Patch the images.json
tmpf=$(mktemp)
oc get -o json configmap/machine-config-operator-images > ${tmpf}
outf=$(mktemp)
python3 > ${outf} <<EOF
import sys,json
cm=json.load(open("${tmpf}"))
images = json.loads(cm['data']['images.json'])
for k in images:
  if k.startswith('machineConfig'):
    images[k] = "${IN_CLUSTER_NAME}"
cm['data']['images.json'] = json.dumps(images)
json.dump(cm, sys.stdout)
EOF
oc replace -f ${outf}
rm ${tmpf} ${outf}

for x in operator controller server daemon; do
patch=$(mktemp)
cat >${patch} <<EOF
spec:
  template:
     spec:
       containers:
         - name: machine-config-${x}
           image: ${IN_CLUSTER_NAME}
EOF

# And for speed, patch the deployment directly rather
# than waiting for the operator to start up and do leader
# election.
case $x in
    controller|operator)
        target=deploy/machine-config-${x}
        ;;
    daemon|server)
        target=daemonset/machine-config-${x}
        ;;
    *) echo "Unhandled $x" && exit 1
esac

oc patch "${target}" -p "$(cat ${patch})"
rm ${patch}
echo "Patched ${target}"
done
oc scale --replicas=1 deploy/machine-config-operator
