package main

/*
This script has to be run before setting Patroni for any given servers.
It checks if the configurations and server details match between the East and Central servers
Also checks if there are existing Clusters with the same name you're trying to setup.
Arguments to be provided from the command line:
c --> Central Hostname
e --> East Hostname
cluster --> Cluster Name
For example:
go run patroniPreChecks.go -c cxxx12345 -e exxx12345 -cluster myDB
*/
