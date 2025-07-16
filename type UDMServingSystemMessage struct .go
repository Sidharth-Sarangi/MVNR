package main
type UDMServingSystemMessage struct {
    SUPI                SUPI  `asn1:"tag:1"`
    PEI                 *PEI  `asn1:"tag:2,optional"`
    GPSI                *GPSI `asn1:"tag:3,optional"`
    GUAMI               *GUAMI `asn1:"tag:4,optional"`
    GUMMEI              *GUMMEI `asn1:"tag:5,optional"`
    PLMNID              *PLMNID `asn1:"tag:6,optional"`
    ServingSystemMethod UDMServingSystemMethod `asn1:"tag:7"`
    ServiceID           *ServiceID `asn1:"tag:8,optional"`
    RoamingIndicator    *RoamingIndicator `asn1:"tag:9,optional"`
}