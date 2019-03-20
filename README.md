# operator-ohm

Goal is to compare the behavior/usefullness of Helm CRD based operators and the
benefits they would bring to the Airship ecosystem.

## Design documents

[Operator](https://github.com/kubekit99/operator-ohm/blob/master/docs/source/operator.rst)

## Armada Operator POC

### Armada operator

This Operator uses the Airship CRDs (ArmadaChart) and
leverage Helm and enable the sequencement of those deployments using either ArmadaChartGroup
or Argo Workflows.
This operator make extensive use of golang, the operator-framework and kubebuilder.
This operator aims at dealing with the differences between Helm v2 and Helm v3 (no tiller, Release CRD).

[Armada](https://github.com/kubekit99/operator-ohm/lob/master/armada-operator/README.md)

### RabbitMQ operator

This Operator simply embeeds the RabbitMQ Openstack HELM chart to Armada-Operator container image.
There is no additional golang code specific code for this operator.
`make install` followed by `make installmanifest` will deploy RabbitMQ on your K8s cluster.

[RabbitMQ](https://github.com/kubekit99/operator-ohm/blob/master/rabbitmq-operator/README.md)

### Memcached operator

This Operator simply embeeds the Memcached Openstack HELM chart to Armada-Operator container image.
There is no additional golang code specific code for this operator.
`make install` followed by `make installmanifest` will deploy Memcached on your K8s cluster.

[Memcached](https://github.com/kubekit99/operator-ohm/blob/master/memcached-operator/README.md)

### MariaDB operator

This Operator simply embeeds the MariaDB Openstack HELM chart to Armada-Operator container image.
There is no additional golang code specific code for this operator.
`make install` followed by `make installmanifest` will deploy MariaDB on your K8s cluster.

[MariaDB](https://github.com/kubekit99/operator-ohm/blob/master/mariadb-operator/README.md)

### Keystone operator 

This Operator simply embeeds the RabbitMQ, Memcached, MariaDB and Keystone Openstack HELM charts 
to Armada-Operator container image.
There is no additional golang code specific code for this operator.
`make install` followed by `make installmanifest` will deploy Keystone on your K8s cluster.

[Keystone](https://github.com/kubekit99/operator-ohm/blob/master/keystone-operator/README.md)

## Other operators POC

## OpenstackLCM operator 

This POC is leveraging the CRD framework with Argo framework in an attempt to control the 
LCM operations applicable to an Openstack cluster.

[OpenstackLCM](https://github.com/kubekit99/operator-ohm/blob/master/openstacklcm-operator/README.md)

## OpenstackHelm operator 

This POC is the first prototype of an Operator leveraging Helm and Argo.

[OpenstackHelm](https://github.com/kubekit99/operator-ohm/blob/master/openstackhelm-operator/README.md)
