#! /bin/sh

PHY_NAME=phy1
CAP_IF_NAME=mon0 # visual interface name
CAP_CHANNEL=11


# for the node that short of RAM
sysctl -w vm.overcommit_memory=1


#  now we create a  visual interface on monitor mode #

# del exist interface
if iw ${CAP_IF_NAME} info >/dev/null
then 
    iw ${CAP_IF_NAME} del
fi


#  add visual interface, and set it on monitor mode 
iw phy ${PHY_NAME} interface add ${CAP_IF_NAME} type monitor

# turn visual interface on
ifconfig ${CAP_IF_NAME} up

# switch to the spec channel 
iw ${CAP_IF_NAME} set channel ${CAP_CHANNEL}

echo "Successfully add ${CAP_IF_NAME} and set channel to ${CAP_CHANNEL}"

