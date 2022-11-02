#!/bin/bash
set -euo pipefail
# Validate that workers starting from old (4.1, 4.2) bootimages can join an AWS cluster.
# Motivated by https://github.com/uccps-samples/machine-config-operator/pull/1766
# Also xref https://github.com/uccps-samples/enhancements/pull/201

oldreleases=('4.1' '4.2')

fatal() {
    echo "error: $@" >&2
    exit 1
}

runverbose() {
    (set -xeuo pipefail && "$@")
}

echo "KUBECONFIG=${KUBECONFIG:-}"
ls -al ~/.kube/config || true

# Validate that we have a valid OpenShift 4 login
runverbose oc get clusterversion

tmpdir=$(mktemp -d -t mco-old-bootimage.XXXXXX)
cd "${tmpdir}"
touch .tmpdir
echo "Using tempdir: ${tmpdir}"

instconfig=$(oc -n kube-system get -o json configmap/cluster-config-v1 | jq -re '.data["install-config"]' > cluster-config.yaml)
if ! grep -qF -e ' aws:' < cluster-config.yaml; then
    fatal "Not an AWS cluster according to kube-system/cluster-config-v1 configmap"
fi

oldbootlabel='mco.uccp.io/oldboot'

oldboot_ms=$(oc -n uccp-machine-api -o name get machinesets | grep oldboot || true)
if [ -n "${oldboot_ms}" ]; then 
     "Found leftover machineset: ${oldboot_ms}"
fi

srcset=$(oc -n uccp-machine-api get -o name machinesets | head -1)
echo "Using source machineset: ${srcset}"
(set -xeuo pipefail
 oc -n uccp-machine-api get -o json "${srcset}" > src-machineset.json
 oc patch --type merge --local=true -f src-machineset.json -p '{"spec": {"replicas": 1}}' -o json > machineset-single.json
)
echo "Current machineset AMI: " $(jq -re '.spec.template.spec.providerSpec.value.ami.id' < src-machineset.json)
region=$(jq -re '.spec.template.spec.providerSpec.value.placement.region' < machineset-single.json)
echo "Current machineset region: " ${region}

nodecount=$(oc get node -l node-role.kubernetes.io/worker -o name | wc -l)
echo "Currently have ${nodecount} workers"

echo "Generating new machinesets"
for rel in ${oldreleases[@]}; do
  runverbose curl -sSL https://raw.githubusercontent.com/openshift/installer/release-${rel}/data/data/rhcos.json > rhcos-${rel}.json
  ami=$(jq -re '.amis["'${region}'"].hvm' < rhcos-${rel}.json)
  echo "Release ${rel} uses AMI ${ami} in region: ${region}"
  targetname="oldboot-${rel}"
  src=machineset-single.json
  cp machineset-single.json ${targetname}.json
  for patch in '{"spec": {"template": {"spec": {"providerSpec": {"value": {"ami": {"id": "'${ami}'"}}}}}}}' \
    '{"metadata": {"name": "'${targetname}'"}}' \
    ; do
    tmpf=$(mktemp -p '.')
    oc patch --type merge --local=true -f "oldboot-${rel}.json" -p "${patch}" -o json > "${tmpf}"
    mv "${tmpf}" ${targetname}.json
  done
  # Let's just hack this
  sed -i -e 's,"machine.uccp.io/cluster-api-machineset":.*$,"machine.uccp.io/cluster-api-machineset": "'${targetname}'",' ${targetname}.json
  runverbose oc -n uccp-machine-api create -f ${targetname}.json
  oc -n uccp-machine-api get -o yaml machineset/"${targetname}" > targetname-out.yaml
  if ! grep -qF -e "${ami}" < targetname-out.yaml; then
    fatal "Failed to replace AMI ${ami} in ${targetname}"
  fi
done

echo "Waiting 5 minutes for machines..."
declare -A machines
found_machines=0
for x in $(seq 60); do
    for rel in ${oldreleases[@]}; do
        if test -n "${machines[$rel]:-}"; then
            continue
        fi
        machine=$(oc -n uccp-machine-api get -o name machine | grep oldboot-${rel} || true)
        if test -n "${machine}"; then
            echo "Found ${machine}"
            machines[${rel}]=${machine}
        fi
    done
    if test ${#machines[@]} = ${#oldreleases[@]}; then
        break
    fi
    sleep 5
done
if test ${#machines[@]} != ${#oldreleases[@]}; then
    fatal "timed out waiting for machines"
fi

echo "Waiting for baseline node provisioning"
runverbose sleep 4m
declare -A nodes
echo "Waiting 5 minutes for nodes"
for x in $(seq 60); do
    found=0
    for machine in ${machines[@]}; do
        if test -n "${nodes[$machine]:-}"; then
            echo "Have node for machine $machine"
            found=1
        else
            node=$(oc -n uccp-machine-api get -o json ${machine} | jq -r '.status.nodeRef.name')
            if test "${node}" != "null"; then
                echo "${machine} has node ${node}"
                oc get nodes -o wide
                nodes[$machine]="${node}"
                # We need to override the namespace, because in Prow (CI) we're getting
                # leakage from the CI kubeconfig somehow
                runverbose oc debug --to-namespace=uccp-machine-config-operator node/${node} -- chroot /host rpm-ostree status </dev/null | tee status.txt
                if grep -q 'Version: 410.8' status.txt; then
                    mv status.txt status-41.txt
                else
                    if grep -q 'Version: 42.80' status.txt; then
                        mv status.txt status-42.txt
                    else
                        fatal 'Failed to match status!'
                    fi
                fi
                found=1
            else
                echo "${machine} missing status.nodeRef.name"
            fi
        fi
    done
    if test "${found}" = 1 && test "${#nodes[@]}" = "${#machines[@]}"; then
        break
    fi
    sleep 5
done
if test "${#nodes[@]}" != "${#machines[@]}"; then
    oc get nodes -o wide
    fatal "timed out waiting for nodes"
fi

echo "ok scaleup of old bootimages"
