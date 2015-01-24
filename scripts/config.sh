# This file holds common parameters for the VM start/stop scripts

VM_NAME=goreport
ZONE=asia-east1-a
MACHINE_TYPE=f1-micro

function wait_vm_ready() {
    VM_IP=$(gcloud compute instances list ${VM_NAME} --format=text | grep natIP | cut -f 2 -d ':' | tr -d ' ')
    echo
    echo "Waiting for VM to be ready on ${VM_IP}."
    echo
    echo "  This will continually check to see if the app is up and "
    echo "  running. If something goes wrong this could hang forever."
    echo
    until $(curl --output /dev/null --silent --head --max-time 1 --fail "http://${VM_IP}"); do
        printf "."
        sleep 2
    done

    echo
    echo "Container VM is now up and running"
    echo
    echo "  http://${VM_IP}"
    echo
}
