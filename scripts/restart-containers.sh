
#! /bin/bash
# Modified from https://github.com/GoogleCloudPlatform/container-vm-guestbook-redis-python/blob/master/config.sh
# Run from project root

set -e

source scripts/config.sh


echo
echo "Setting new containers into VM"
gcloud compute instances add-metadata --zone ${ZONE} ${VM_NAME} \
  --metadata-from-file google-container-manifest=containers.yaml

echo
echo "Rebooting VM to restart containers"
gcloud compute instances reset --zone ${ZONE} ${VM_NAME}

wait_vm_ready
