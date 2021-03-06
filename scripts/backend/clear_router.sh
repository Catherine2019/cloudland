#!/bin/bash

cd $(dirname $0)
source ../cloudrc

[ $# -lt 2 ] && echo "$0 <router> <vrrp_vni>" && exit -1

router=router-$1
vrrp_vni=$2
[ -z "$router" ] && exit 1

suffix=${router%%-*}
ip netns exec $router ip link set lo down
ip link del te-$suffix
#apply_vnic -D te-$suffix
ip link del ti-$suffix
#apply_vnic -D ti-$suffix
./clear_link.sh $vrrp_vni
interfaces=$(cat)
i=0
n=$(jq length <<< $interfaces)
while [ $i -lt $n ]; do
    vni=$(jq -r .[$i].vni <<< $interfaces)
    ip link del ln-$vni
    ./clear_link.sh $vni
    let i=$i+1
done
udevadm settle
ip netns del $router
router_dir=/opt/cloudland/cache/router/$router
kill $(cat $router_dir/keepalived.pid)
rm -rf $router_dir
