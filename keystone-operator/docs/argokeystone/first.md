# ./argokeystone/first.md

```bash
prompt$ kubectl get all --all-namespaces
NAMESPACE     NAME                                            READY   STATUS    RESTARTS   AGE
argo          pod/argo-ui-f57bb579c-swz7h                     1/1     Running   0          10m
argo          pod/argo-workflow-controller-6d987f7fdc-l5cln   1/1     Running   0          10m
kube-system   pod/calico-etcd-k7h9l                           1/1     Running   0          10m
kube-system   pod/calico-kube-controllers-79bd896977-gw9q9    1/1     Running   0          10m
kube-system   pod/calico-node-qjfj8                           1/1     Running   1          10m
kube-system   pod/coredns-fb8b8dccf-kvdg6                     1/1     Running   0          10m
kube-system   pod/coredns-fb8b8dccf-p9zhw                     1/1     Running   0          10m
kube-system   pod/etcd-airship                                1/1     Running   0          9m59s
kube-system   pod/kube-apiserver-airship                      1/1     Running   0          10m
kube-system   pod/kube-controller-manager-airship             1/1     Running   0          10m
kube-system   pod/kube-proxy-dv5ls                            1/1     Running   0          10m
kube-system   pod/kube-scheduler-airship                      1/1     Running   0          10m
kube-system   pod/tiller-deploy-8458f6c667-sgz5z              1/1     Running   0          10m

NAMESPACE     NAME                    TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                  AGE
argo          service/argo-ui         NodePort    10.107.6.139    <none>        80:32711/TCP             10m
default       service/kubernetes      ClusterIP   10.96.0.1       <none>        443/TCP                  11m
kube-system   service/calico-etcd     ClusterIP   10.96.232.136   <none>        6666/TCP                 10m
kube-system   service/kube-dns        ClusterIP   10.96.0.10      <none>        53/UDP,53/TCP,9153/TCP   11m
kube-system   service/tiller-deploy   ClusterIP   10.102.66.38    <none>        44134/TCP                10m

NAMESPACE     NAME                         DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR                     AGE
kube-system   daemonset.apps/calico-etcd   1         1         1       1            1           node-role.kubernetes.io/master=   10m
kube-system   daemonset.apps/calico-node   1         1         1       1            1           beta.kubernetes.io/os=linux       10m
kube-system   daemonset.apps/kube-proxy    1         1         1       1            1           <none>                            11m

NAMESPACE     NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
argo          deployment.apps/argo-ui                    1/1     1            1           10m
argo          deployment.apps/argo-workflow-controller   1/1     1            1           10m
kube-system   deployment.apps/calico-kube-controllers    1/1     1            1           10m
kube-system   deployment.apps/coredns                    2/2     2            2           11m
kube-system   deployment.apps/tiller-deploy              1/1     1            1           10m

NAMESPACE     NAME                                                  DESIRED   CURRENT   READY   AGE
argo          replicaset.apps/argo-ui-f57bb579c                     1         1         1       10m
argo          replicaset.apps/argo-workflow-controller-6d987f7fdc   1         1         1       10m
kube-system   replicaset.apps/calico-kube-controllers-79bd896977    1         1         1       10m
kube-system   replicaset.apps/coredns-fb8b8dccf                     2         2         2       10m
kube-system   replicaset.apps/tiller-deploy-8458f6c667              1         1         1       10m
```
 
```bash
prompt$ kubectl get nodes --show-labels
NAME      STATUS   ROLES    AGE   VERSION   LABELS
airship   Ready    master   11m   v1.14.0   beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=airship,kubernetes.io/os=linux,node-role.kubernetes.io/master=,openstack-control-plane=enabled
```
 
```bash
prompt$ make docker-build
tar -C helm-charts/ -xvf helm-charts/keystone/charts/helm-toolkit-0.1.0.tgz
helm-toolkit/Chart.yaml
helm-toolkit/values.yaml
helm-toolkit/templates/endpoints/_authenticated_endpoint_uri_lookup.tpl
helm-toolkit/templates/endpoints/_endpoint_host_lookup.tpl
helm-toolkit/templates/endpoints/_endpoint_port_lookup.tpl
helm-toolkit/templates/endpoints/_host_and_port_endpoint_uri_lookup.tpl
helm-toolkit/templates/endpoints/_hostname_fqdn_endpoint_lookup.tpl
helm-toolkit/templates/endpoints/_hostname_namespaced_endpoint_lookup.tpl
helm-toolkit/templates/endpoints/_hostname_short_endpoint_lookup.tpl
helm-toolkit/templates/endpoints/_keystone_endpoint_name_lookup.tpl
helm-toolkit/templates/endpoints/_keystone_endpoint_path_lookup.tpl
helm-toolkit/templates/endpoints/_keystone_endpoint_scheme_lookup.tpl
helm-toolkit/templates/endpoints/_keystone_endpoint_uri_lookup.tpl
helm-toolkit/templates/endpoints/_service_name_endpoint_with_namespace_lookup.tpl
helm-toolkit/templates/manifests/_ceph-storageclass.tpl
helm-toolkit/templates/manifests/_ingress.tpl
helm-toolkit/templates/manifests/_job-bootstrap.tpl
helm-toolkit/templates/manifests/_job-db-drop-mysql.tpl
helm-toolkit/templates/manifests/_job-db-init-mysql.tpl
helm-toolkit/templates/manifests/_job-db-sync.tpl
helm-toolkit/templates/manifests/_job-ks-endpoints.tpl
helm-toolkit/templates/manifests/_job-ks-service.tpl
helm-toolkit/templates/manifests/_job-ks-user.yaml.tpl
helm-toolkit/templates/manifests/_job-rabbit-init.yaml.tpl
helm-toolkit/templates/manifests/_job-s3-bucket.yaml.tpl
helm-toolkit/templates/manifests/_job-s3-user.yaml.tpl
helm-toolkit/templates/manifests/_job_image_repo_sync.tpl
helm-toolkit/templates/manifests/_network_policy.tpl
helm-toolkit/templates/manifests/_secret-tls.yaml.tpl
helm-toolkit/templates/manifests/_service-ingress.tpl
helm-toolkit/templates/scripts/_create-s3-bucket.sh.tpl
helm-toolkit/templates/scripts/_create-s3-user.sh.tpl
helm-toolkit/templates/scripts/_db-drop.py.tpl
helm-toolkit/templates/scripts/_db-init.py.tpl
helm-toolkit/templates/scripts/_db-pg-init.sh.tpl
helm-toolkit/templates/scripts/_image-repo-sync.sh.tpl
helm-toolkit/templates/scripts/_ks-domain-user.sh.tpl
helm-toolkit/templates/scripts/_ks-endpoints.sh.tpl
helm-toolkit/templates/scripts/_ks-service.sh.tpl
helm-toolkit/templates/scripts/_ks-user.sh.tpl
helm-toolkit/templates/scripts/_rabbit-init.sh.tpl
helm-toolkit/templates/scripts/_rally_test.sh.tpl
helm-toolkit/templates/snippets/_image.tpl
helm-toolkit/templates/snippets/_keystone_openrc_env_vars.tpl
helm-toolkit/templates/snippets/_keystone_secret_openrc.tpl
helm-toolkit/templates/snippets/_keystone_user_create_env_vars.tpl
helm-toolkit/templates/snippets/_kubernetes_apparmor_configmap.tpl
helm-toolkit/templates/snippets/_kubernetes_apparmor_loader_init_container.tpl
helm-toolkit/templates/snippets/_kubernetes_apparmor_volumes.tpl
helm-toolkit/templates/snippets/_kubernetes_container_security_context.tpl
helm-toolkit/templates/snippets/_kubernetes_entrypoint_init_container.tpl
helm-toolkit/templates/snippets/_kubernetes_kubectl_params.tpl
helm-toolkit/templates/snippets/_kubernetes_mandatory_access_control_annotation.tpl
helm-toolkit/templates/snippets/_kubernetes_metadata_labels.tpl
helm-toolkit/templates/snippets/_kubernetes_pod_anti_affinity.tpl
helm-toolkit/templates/snippets/_kubernetes_pod_rbac_roles.tpl
helm-toolkit/templates/snippets/_kubernetes_pod_rbac_serviceaccount.tpl
helm-toolkit/templates/snippets/_kubernetes_pod_security_context.tpl
helm-toolkit/templates/snippets/_kubernetes_resources.tpl
helm-toolkit/templates/snippets/_kubernetes_seccomp_annotation.tpl
helm-toolkit/templates/snippets/_kubernetes_tolerations.tpl
helm-toolkit/templates/snippets/_kubernetes_upgrades_daemonset.tpl
helm-toolkit/templates/snippets/_kubernetes_upgrades_deployment.tpl
helm-toolkit/templates/snippets/_prometheus_pod_annotations.tpl
helm-toolkit/templates/snippets/_prometheus_service_annotations.tpl
helm-toolkit/templates/snippets/_release_uuid.tpl
helm-toolkit/templates/snippets/_rgw_s3_admin_env_vars.tpl
helm-toolkit/templates/snippets/_rgw_s3_secret_creds.tpl
helm-toolkit/templates/snippets/_rgw_s3_user_env_vars.tpl
helm-toolkit/templates/snippets/_values_template_renderer.tpl
helm-toolkit/templates/tls/_tls_generate_certs.tpl
helm-toolkit/templates/utils/_comma_joined_service_list.tpl
helm-toolkit/templates/utils/_configmap_templater.tpl
helm-toolkit/templates/utils/_daemonset_overrides.tpl
helm-toolkit/templates/utils/_dependency_resolver.tpl
helm-toolkit/templates/utils/_hash.tpl
helm-toolkit/templates/utils/_host_list.tpl
helm-toolkit/templates/utils/_image_sync_list.tpl
helm-toolkit/templates/utils/_joinListWithComma.tpl
helm-toolkit/templates/utils/_joinListWithPrefix.tpl
helm-toolkit/templates/utils/_joinListWithSpace.tpl
helm-toolkit/templates/utils/_merge.tpl
helm-toolkit/templates/utils/_template.tpl
helm-toolkit/templates/utils/_to_ini.tpl
helm-toolkit/templates/utils/_to_k8s_env_vars.tpl
helm-toolkit/templates/utils/_to_kv_list.tpl
helm-toolkit/templates/utils/_to_oslo_conf.tpl
helm-toolkit/requirements.yaml
rsync -avz ../mariadb-operator/helm-charts/mariadb helm-charts
sending incremental file list

sent 1,871 bytes  received 27 bytes  3,796.00 bytes/sec
total size is 186,862  speedup is 98.45
rsync -avz ../memcached-operator/helm-charts/memcached helm-charts
sending incremental file list

sent 753 bytes  received 23 bytes  1,552.00 bytes/sec
total size is 54,824  speedup is 70.65
rsync -avz ../rabbitmq-operator/helm-charts/rabbitmq helm-charts
sending incremental file list

sent 1,180 bytes  received 24 bytes  2,408.00 bytes/sec
total size is 86,816  speedup is 72.11
docker build -t kubekit99/keystone-armada-operator:poc -f build/Dockerfile.armada-operator .
Sending build context to Docker daemon  1.322MB
Step 1/2 : FROM kubekit99/armada-operator-dev:latest
 ---> f5bc56ef607b
Step 2/2 : COPY helm-charts/ /opt/armada/helm-charts/
 ---> Using cache
 ---> bc7d9ab7c350
Successfully built bc7d9ab7c350
Successfully tagged kubekit99/keystone-armada-operator:poc
docker tag kubekit99/keystone-armada-operator:poc kubekit99/keystone-armada-operator:latest
docker build -t kubekit99/keystone-oslc-operator:poc -f build/Dockerfile.oslc-operator .
Sending build context to Docker daemon  1.322MB
Step 1/2 : FROM kubekit99/openstacklcm-operator-dev:latest
 ---> 16b6fab9fb90
Step 2/2 : COPY helm-charts /opt/openstacklcm-operator/helm-charts/
 ---> Using cache
 ---> a58113a8da0e
Successfully built a58113a8da0e
Successfully tagged kubekit99/keystone-oslc-operator:poc
docker tag kubekit99/keystone-oslc-operator:poc kubekit99/keystone-oslc-operator:latest
```
 
```bash
prompt$ make install-armada && make install-oslc
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadachartgroup.yaml
customresourcedefinition.apiextensions.k8s.io/armadachartgroups.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadachart.yaml
customresourcedefinition.apiextensions.k8s.io/armadacharts.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadamanifest.yaml
customresourcedefinition.apiextensions.k8s.io/armadamanifests.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/role_binding.yaml
rolebinding.rbac.authorization.k8s.io/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/role.yaml
role.rbac.authorization.k8s.io/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/service_account.yaml
serviceaccount/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/argo_armada_role.yaml
serviceaccount/armada-argo-sa created
role.rbac.authorization.k8s.io/armada-argo-role created
rolebinding.rbac.authorization.k8s.io/armada-argo-rolebinding created
kubectl create -f deploy/armada-operator.yaml
deployment.apps/keystone-armada-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_deletephase.yaml
customresourcedefinition.apiextensions.k8s.io/deletephases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_installphase.yaml
customresourcedefinition.apiextensions.k8s.io/installphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_operationalphase.yaml
customresourcedefinition.apiextensions.k8s.io/operationalphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_planningphase.yaml
customresourcedefinition.apiextensions.k8s.io/planningphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_rollbackphase.yaml
customresourcedefinition.apiextensions.k8s.io/rollbackphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_testphase.yaml
customresourcedefinition.apiextensions.k8s.io/testphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_trafficdrainphase.yaml
customresourcedefinition.apiextensions.k8s.io/trafficdrainphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_trafficrolloutphase.yaml
customresourcedefinition.apiextensions.k8s.io/trafficrolloutphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_upgradephase.yaml
customresourcedefinition.apiextensions.k8s.io/upgradephases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_oslc.yaml
customresourcedefinition.apiextensions.k8s.io/oslcs.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/role_binding.yaml
rolebinding.rbac.authorization.k8s.io/openstacklcm-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/roles.yaml
role.rbac.authorization.k8s.io/openstacklcm-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/service_account.yaml
serviceaccount/openstacklcm-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/argo_openstacklcm_role.yaml
serviceaccount/openstacklcm-argo-sa created
role.rbac.authorization.k8s.io/openstacklcm-argo-role created
rolebinding.rbac.authorization.k8s.io/openstacklcm-argo-rolebinding created
kubectl create -f deploy/oslc-operator.yaml
deployment.apps/keystone-oslc-operator created
```
 
```bash
prompt$
```
# ./argokeystone/first.md
```bash
prompt$ kubectl get all --all-namespaces
NAMESPACE     NAME                                            READY   STATUS    RESTARTS   AGE
argo          pod/argo-ui-f57bb579c-swz7h                     1/1     Running   0          10m
argo          pod/argo-workflow-controller-6d987f7fdc-l5cln   1/1     Running   0          10m
kube-system   pod/calico-etcd-k7h9l                           1/1     Running   0          10m
kube-system   pod/calico-kube-controllers-79bd896977-gw9q9    1/1     Running   0          10m
kube-system   pod/calico-node-qjfj8                           1/1     Running   1          10m
kube-system   pod/coredns-fb8b8dccf-kvdg6                     1/1     Running   0          10m
kube-system   pod/coredns-fb8b8dccf-p9zhw                     1/1     Running   0          10m
kube-system   pod/etcd-airship                                1/1     Running   0          9m59s
kube-system   pod/kube-apiserver-airship                      1/1     Running   0          10m
kube-system   pod/kube-controller-manager-airship             1/1     Running   0          10m
kube-system   pod/kube-proxy-dv5ls                            1/1     Running   0          10m
kube-system   pod/kube-scheduler-airship                      1/1     Running   0          10m
kube-system   pod/tiller-deploy-8458f6c667-sgz5z              1/1     Running   0          10m

NAMESPACE     NAME                    TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                  AGE
argo          service/argo-ui         NodePort    10.107.6.139    <none>        80:32711/TCP             10m
default       service/kubernetes      ClusterIP   10.96.0.1       <none>        443/TCP                  11m
kube-system   service/calico-etcd     ClusterIP   10.96.232.136   <none>        6666/TCP                 10m
kube-system   service/kube-dns        ClusterIP   10.96.0.10      <none>        53/UDP,53/TCP,9153/TCP   11m
kube-system   service/tiller-deploy   ClusterIP   10.102.66.38    <none>        44134/TCP                10m

NAMESPACE     NAME                         DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR                     AGE
kube-system   daemonset.apps/calico-etcd   1         1         1       1            1           node-role.kubernetes.io/master=   10m
kube-system   daemonset.apps/calico-node   1         1         1       1            1           beta.kubernetes.io/os=linux       10m
kube-system   daemonset.apps/kube-proxy    1         1         1       1            1           <none>                            11m

NAMESPACE     NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
argo          deployment.apps/argo-ui                    1/1     1            1           10m
argo          deployment.apps/argo-workflow-controller   1/1     1            1           10m
kube-system   deployment.apps/calico-kube-controllers    1/1     1            1           10m
kube-system   deployment.apps/coredns                    2/2     2            2           11m
kube-system   deployment.apps/tiller-deploy              1/1     1            1           10m

NAMESPACE     NAME                                                  DESIRED   CURRENT   READY   AGE
argo          replicaset.apps/argo-ui-f57bb579c                     1         1         1       10m
argo          replicaset.apps/argo-workflow-controller-6d987f7fdc   1         1         1       10m
kube-system   replicaset.apps/calico-kube-controllers-79bd896977    1         1         1       10m
kube-system   replicaset.apps/coredns-fb8b8dccf                     2         2         2       10m
kube-system   replicaset.apps/tiller-deploy-8458f6c667              1         1         1       10m
```
 
```bash
prompt$ kubectl get nodes --show-labels
NAME      STATUS   ROLES    AGE   VERSION   LABELS
airship   Ready    master   11m   v1.14.0   beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=airship,kubernetes.io/os=linux,node-role.kubernetes.io/master=,openstack-control-plane=enabled
```
 
```bash
prompt$ make docker-build
tar -C helm-charts/ -xvf helm-charts/keystone/charts/helm-toolkit-0.1.0.tgz
helm-toolkit/Chart.yaml
helm-toolkit/values.yaml
helm-toolkit/templates/endpoints/_authenticated_endpoint_uri_lookup.tpl
helm-toolkit/templates/endpoints/_endpoint_host_lookup.tpl
helm-toolkit/templates/endpoints/_endpoint_port_lookup.tpl
helm-toolkit/templates/endpoints/_host_and_port_endpoint_uri_lookup.tpl
helm-toolkit/templates/endpoints/_hostname_fqdn_endpoint_lookup.tpl
helm-toolkit/templates/endpoints/_hostname_namespaced_endpoint_lookup.tpl
helm-toolkit/templates/endpoints/_hostname_short_endpoint_lookup.tpl
helm-toolkit/templates/endpoints/_keystone_endpoint_name_lookup.tpl
helm-toolkit/templates/endpoints/_keystone_endpoint_path_lookup.tpl
helm-toolkit/templates/endpoints/_keystone_endpoint_scheme_lookup.tpl
helm-toolkit/templates/endpoints/_keystone_endpoint_uri_lookup.tpl
helm-toolkit/templates/endpoints/_service_name_endpoint_with_namespace_lookup.tpl
helm-toolkit/templates/manifests/_ceph-storageclass.tpl
helm-toolkit/templates/manifests/_ingress.tpl
helm-toolkit/templates/manifests/_job-bootstrap.tpl
helm-toolkit/templates/manifests/_job-db-drop-mysql.tpl
helm-toolkit/templates/manifests/_job-db-init-mysql.tpl
helm-toolkit/templates/manifests/_job-db-sync.tpl
helm-toolkit/templates/manifests/_job-ks-endpoints.tpl
helm-toolkit/templates/manifests/_job-ks-service.tpl
helm-toolkit/templates/manifests/_job-ks-user.yaml.tpl
helm-toolkit/templates/manifests/_job-rabbit-init.yaml.tpl
helm-toolkit/templates/manifests/_job-s3-bucket.yaml.tpl
helm-toolkit/templates/manifests/_job-s3-user.yaml.tpl
helm-toolkit/templates/manifests/_job_image_repo_sync.tpl
helm-toolkit/templates/manifests/_network_policy.tpl
helm-toolkit/templates/manifests/_secret-tls.yaml.tpl
helm-toolkit/templates/manifests/_service-ingress.tpl
helm-toolkit/templates/scripts/_create-s3-bucket.sh.tpl
helm-toolkit/templates/scripts/_create-s3-user.sh.tpl
helm-toolkit/templates/scripts/_db-drop.py.tpl
helm-toolkit/templates/scripts/_db-init.py.tpl
helm-toolkit/templates/scripts/_db-pg-init.sh.tpl
helm-toolkit/templates/scripts/_image-repo-sync.sh.tpl
helm-toolkit/templates/scripts/_ks-domain-user.sh.tpl
helm-toolkit/templates/scripts/_ks-endpoints.sh.tpl
helm-toolkit/templates/scripts/_ks-service.sh.tpl
helm-toolkit/templates/scripts/_ks-user.sh.tpl
helm-toolkit/templates/scripts/_rabbit-init.sh.tpl
helm-toolkit/templates/scripts/_rally_test.sh.tpl
helm-toolkit/templates/snippets/_image.tpl
helm-toolkit/templates/snippets/_keystone_openrc_env_vars.tpl
helm-toolkit/templates/snippets/_keystone_secret_openrc.tpl
helm-toolkit/templates/snippets/_keystone_user_create_env_vars.tpl
helm-toolkit/templates/snippets/_kubernetes_apparmor_configmap.tpl
helm-toolkit/templates/snippets/_kubernetes_apparmor_loader_init_container.tpl
helm-toolkit/templates/snippets/_kubernetes_apparmor_volumes.tpl
helm-toolkit/templates/snippets/_kubernetes_container_security_context.tpl
helm-toolkit/templates/snippets/_kubernetes_entrypoint_init_container.tpl
helm-toolkit/templates/snippets/_kubernetes_kubectl_params.tpl
helm-toolkit/templates/snippets/_kubernetes_mandatory_access_control_annotation.tpl
helm-toolkit/templates/snippets/_kubernetes_metadata_labels.tpl
helm-toolkit/templates/snippets/_kubernetes_pod_anti_affinity.tpl
helm-toolkit/templates/snippets/_kubernetes_pod_rbac_roles.tpl
helm-toolkit/templates/snippets/_kubernetes_pod_rbac_serviceaccount.tpl
helm-toolkit/templates/snippets/_kubernetes_pod_security_context.tpl
helm-toolkit/templates/snippets/_kubernetes_resources.tpl
helm-toolkit/templates/snippets/_kubernetes_seccomp_annotation.tpl
helm-toolkit/templates/snippets/_kubernetes_tolerations.tpl
helm-toolkit/templates/snippets/_kubernetes_upgrades_daemonset.tpl
helm-toolkit/templates/snippets/_kubernetes_upgrades_deployment.tpl
helm-toolkit/templates/snippets/_prometheus_pod_annotations.tpl
helm-toolkit/templates/snippets/_prometheus_service_annotations.tpl
helm-toolkit/templates/snippets/_release_uuid.tpl
helm-toolkit/templates/snippets/_rgw_s3_admin_env_vars.tpl
helm-toolkit/templates/snippets/_rgw_s3_secret_creds.tpl
helm-toolkit/templates/snippets/_rgw_s3_user_env_vars.tpl
helm-toolkit/templates/snippets/_values_template_renderer.tpl
helm-toolkit/templates/tls/_tls_generate_certs.tpl
helm-toolkit/templates/utils/_comma_joined_service_list.tpl
helm-toolkit/templates/utils/_configmap_templater.tpl
helm-toolkit/templates/utils/_daemonset_overrides.tpl
helm-toolkit/templates/utils/_dependency_resolver.tpl
helm-toolkit/templates/utils/_hash.tpl
helm-toolkit/templates/utils/_host_list.tpl
helm-toolkit/templates/utils/_image_sync_list.tpl
helm-toolkit/templates/utils/_joinListWithComma.tpl
helm-toolkit/templates/utils/_joinListWithPrefix.tpl
helm-toolkit/templates/utils/_joinListWithSpace.tpl
helm-toolkit/templates/utils/_merge.tpl
helm-toolkit/templates/utils/_template.tpl
helm-toolkit/templates/utils/_to_ini.tpl
helm-toolkit/templates/utils/_to_k8s_env_vars.tpl
helm-toolkit/templates/utils/_to_kv_list.tpl
helm-toolkit/templates/utils/_to_oslo_conf.tpl
helm-toolkit/requirements.yaml
rsync -avz ../mariadb-operator/helm-charts/mariadb helm-charts
sending incremental file list

sent 1,871 bytes  received 27 bytes  3,796.00 bytes/sec
total size is 186,862  speedup is 98.45
rsync -avz ../memcached-operator/helm-charts/memcached helm-charts
sending incremental file list

sent 753 bytes  received 23 bytes  1,552.00 bytes/sec
total size is 54,824  speedup is 70.65
rsync -avz ../rabbitmq-operator/helm-charts/rabbitmq helm-charts
sending incremental file list

sent 1,180 bytes  received 24 bytes  2,408.00 bytes/sec
total size is 86,816  speedup is 72.11
docker build -t kubekit99/keystone-armada-operator:poc -f build/Dockerfile.armada-operator .
Sending build context to Docker daemon  1.322MB
Step 1/2 : FROM kubekit99/armada-operator-dev:latest
 ---> f5bc56ef607b
Step 2/2 : COPY helm-charts/ /opt/armada/helm-charts/
 ---> Using cache
 ---> bc7d9ab7c350
Successfully built bc7d9ab7c350
Successfully tagged kubekit99/keystone-armada-operator:poc
docker tag kubekit99/keystone-armada-operator:poc kubekit99/keystone-armada-operator:latest
docker build -t kubekit99/keystone-oslc-operator:poc -f build/Dockerfile.oslc-operator .
Sending build context to Docker daemon  1.322MB
Step 1/2 : FROM kubekit99/openstacklcm-operator-dev:latest
 ---> 16b6fab9fb90
Step 2/2 : COPY helm-charts /opt/openstacklcm-operator/helm-charts/
 ---> Using cache
 ---> a58113a8da0e
Successfully built a58113a8da0e
Successfully tagged kubekit99/keystone-oslc-operator:poc
docker tag kubekit99/keystone-oslc-operator:poc kubekit99/keystone-oslc-operator:latest
```
 
```bash
prompt$ make install-armada && make install-oslc
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadachartgroup.yaml
customresourcedefinition.apiextensions.k8s.io/armadachartgroups.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadachart.yaml
customresourcedefinition.apiextensions.k8s.io/armadacharts.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadamanifest.yaml
customresourcedefinition.apiextensions.k8s.io/armadamanifests.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/role_binding.yaml
rolebinding.rbac.authorization.k8s.io/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/role.yaml
role.rbac.authorization.k8s.io/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/service_account.yaml
serviceaccount/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/argo_armada_role.yaml
serviceaccount/armada-argo-sa created
role.rbac.authorization.k8s.io/armada-argo-role created
rolebinding.rbac.authorization.k8s.io/armada-argo-rolebinding created
kubectl create -f deploy/armada-operator.yaml
deployment.apps/keystone-armada-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_deletephase.yaml
customresourcedefinition.apiextensions.k8s.io/deletephases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_installphase.yaml
customresourcedefinition.apiextensions.k8s.io/installphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_operationalphase.yaml
customresourcedefinition.apiextensions.k8s.io/operationalphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_planningphase.yaml
customresourcedefinition.apiextensions.k8s.io/planningphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_rollbackphase.yaml
customresourcedefinition.apiextensions.k8s.io/rollbackphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_testphase.yaml
customresourcedefinition.apiextensions.k8s.io/testphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_trafficdrainphase.yaml
customresourcedefinition.apiextensions.k8s.io/trafficdrainphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_trafficrolloutphase.yaml
customresourcedefinition.apiextensions.k8s.io/trafficrolloutphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_upgradephase.yaml
customresourcedefinition.apiextensions.k8s.io/upgradephases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_oslc.yaml
customresourcedefinition.apiextensions.k8s.io/oslcs.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/role_binding.yaml
rolebinding.rbac.authorization.k8s.io/openstacklcm-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/roles.yaml
role.rbac.authorization.k8s.io/openstacklcm-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/service_account.yaml
serviceaccount/openstacklcm-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/argo_openstacklcm_role.yaml
serviceaccount/openstacklcm-argo-sa created
role.rbac.authorization.k8s.io/openstacklcm-argo-role created
rolebinding.rbac.authorization.k8s.io/openstacklcm-argo-rolebinding created
kubectl create -f deploy/oslc-operator.yaml
deployment.apps/keystone-oslc-operator created
```
