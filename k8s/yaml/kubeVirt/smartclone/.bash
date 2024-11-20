#!/bin/bash
VG_NAME="openebs_vg_1"    # Replace with your volume group name
SNAPSHOT_NAME="1236a1ba-efea-4295-bcbe-26ebb26ce0bf"   # Replace with your snapshot name
NEW_LV_NAME="new_lv_name"       # Replace with the desired name for the new LVM volume
NEW_LV_SIZE="20G"                # Replace with the desired size of the new LVM volume

# Step 1: Check if the snapshot exists
if ! lvs | grep -q "$SNAPSHOT_NAME"; then
    echo "Error: Snapshot $SNAPSHOT_NAME does not exist."
    exit 1
fi

# Step 2: Create a new logical volume from the snapshot
lvcreate --name "$NEW_LV_NAME" --size "$NEW_LV_SIZE" --snapshot "$VG_NAME/$SNAPSHOT_NAME" "$VG_NAME"

if [ $? -ne 0 ]; then
    echo "Error: Failed to create new logical volume from snapshot."
    exit 1
fi

echo "Successfully created new LVM volume: $NEW_LV_NAME"

# Step 3: Format the new LVM volume (optional, depending on your needs)
# Uncomment the following line if you need to format the volume (e.g., as ext4)
# mkfs.ext4 /dev/$VG_NAME/$NEW_LV_NAME

# Step 4: Optionally, you can add this volume to a new KVM instance
# Replace `<vm_name>` with your desired VM name
# virt-install --name <vm_name> --memory 2048 --vcpus 1 --disk path=/dev/$VG_NAME/$NEW_LV_NAME,format=raw --cdrom /path/to/installer.iso

echo "New LVM volume $NEW_LV_NAME is ready to be used as a KVM storage backend."