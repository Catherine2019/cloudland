#!/bin/bash

cd `dirname $0`
source ../cloudrc

[ $# -lt 3 ] && echo "$0 <vm_ID> <vlan> [vm_mac]" && exit -1

vm_ID=$1
vlan=$2
vm_mac=$3
nic_name=tap$(echo $vm_mac | cut -d: -f4- | tr -d :)
vm_br=br$vlan
./create_link.sh $vlan
virsh attach-interface $vm_ID bridge $vm_br --model virtio --mac $vm_mac --config --target $nic_name
