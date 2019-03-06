
all: setup

setup:
ifndef GOPATH
	$(error GOPATH not defined, please define GOPATH. Run "go help gopath" to learn more about GOPATH)
endif


# Brute force installation of the charts to validate current charts are working properly
manualhelminstall: setup
	kubectl label nodes airship openstack-control-plane=enabled --overwrite
	helm install --name memcached --namespace openstack memcached-operator/helm-charts/memcached
	helm install --name rabbitmq --namespace openstack rabbitmq-operator/helm-charts/rabbitmq
	helm install --name mariadb --namespace openstack mariadb-operator/helm-charts/mariadb/
	kubectl get all -n openstack

manualhelmpurge: setup
	helm delete --purge memcached
	helm delete --purge rabbitmq
	helm delete --purge mariadb
	kubectl delete configmap --all -n openstack
	kubectl delete secret --all -n openstack
	kubectl get all -n openstack

manualkeystoneinstall: setup
	helm install --name keystone --namespace openstack keystone-operator/helm-charts/keystone
	kubectl get all -n openstack

manualkeystonepurge: setup
	helm delete --purge keystone
	kubectl get all -n openstack
