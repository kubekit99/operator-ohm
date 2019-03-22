# Keystone operator

Armada-Operator deployment of keystone

# Prerequisite

Keystone-Operator still requires:
- Helm2 tiller to be deployed in your cluster


# Deploy the Keystone-Operator

## Rebuild the operator if needed and deploy it

```bash
make install

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
helm-toolkit/templates/snippets/_kubernetes_entrypoint_init_container.tpl
helm-toolkit/templates/snippets/_kubernetes_kubectl_params.tpl
helm-toolkit/templates/snippets/_kubernetes_mandatory_access_control_annotation.tpl
helm-toolkit/templates/snippets/_kubernetes_metadata_labels.tpl
helm-toolkit/templates/snippets/_kubernetes_pod_anti_affinity.tpl
helm-toolkit/templates/snippets/_kubernetes_pod_rbac_roles.tpl
helm-toolkit/templates/snippets/_kubernetes_pod_rbac_serviceaccount.tpl
helm-toolkit/templates/snippets/_kubernetes_pod_security_context.tpl
helm-toolkit/templates/snippets/_kubernetes_resources.tpl
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
mariadb/
mariadb/.helmignore
mariadb/Chart.yaml
mariadb/README.rst
mariadb/requirements.yaml
mariadb/values.yaml
mariadb/charts/
mariadb/charts/helm-toolkit-0.1.0.tgz
mariadb/files/
mariadb/files/nginx.tmpl
mariadb/templates/
mariadb/templates/configmap-bin.yaml
mariadb/templates/configmap-etc.yaml
mariadb/templates/configmap-services-tcp.yaml
mariadb/templates/deployment-error.yaml
mariadb/templates/deployment-ingress.yaml
mariadb/templates/job-image-repo-sync.yaml
mariadb/templates/network_policy.yaml
mariadb/templates/pdb-mariadb.yaml
mariadb/templates/secret-db-root-password.yaml
mariadb/templates/secrets-etc.yaml
mariadb/templates/service-discovery.yaml
mariadb/templates/service-error.yaml
mariadb/templates/service-ingress.yaml
mariadb/templates/service.yaml
mariadb/templates/statefulset.yaml
mariadb/templates/bin/
mariadb/templates/bin/_mariadb-ingress-controller.sh.tpl
mariadb/templates/bin/_mariadb-ingress-error-pages.sh.tpl
mariadb/templates/bin/_readiness.sh.tpl
mariadb/templates/bin/_start.py.tpl
mariadb/templates/bin/_stop.sh.tpl
mariadb/templates/etc/
mariadb/templates/etc/_00-base.cnf.tpl
mariadb/templates/etc/_20-override.cnf.tpl
mariadb/templates/etc/_99-force.cnf.tpl
mariadb/templates/etc/_my.cnf.tpl
mariadb/templates/monitoring/
mariadb/templates/monitoring/prometheus/
mariadb/templates/monitoring/prometheus/exporter-configmap-bin.yaml
mariadb/templates/monitoring/prometheus/exporter-deployment.yaml
mariadb/templates/monitoring/prometheus/exporter-job-create-user.yaml
mariadb/templates/monitoring/prometheus/exporter-secrets-etc.yaml
mariadb/templates/monitoring/prometheus/exporter-service.yaml
mariadb/templates/monitoring/prometheus/bin/
mariadb/templates/monitoring/prometheus/bin/_create-mysql-user.sh.tpl
mariadb/templates/monitoring/prometheus/bin/_mysqld-exporter.sh.tpl
mariadb/templates/monitoring/prometheus/secrets/
mariadb/templates/monitoring/prometheus/secrets/_exporter_user.cnf.tpl
mariadb/templates/secrets/
mariadb/templates/secrets/_admin_user.cnf.tpl

sent 74,563 bytes  received 840 bytes  150,806.00 bytes/sec
total size is 161,804  speedup is 2.15
rsync -avz ../memcached-operator/helm-charts/memcached helm-charts
sending incremental file list
memcached/
memcached/.helmignore
memcached/Chart.yaml
memcached/requirements.yaml
memcached/values.yaml
memcached/charts/
memcached/charts/helm-toolkit-0.1.0.tgz
memcached/templates/
memcached/templates/configmap-bin.yaml
memcached/templates/deployment.yaml
memcached/templates/job-image-repo-sync.yaml
memcached/templates/network_policy.yaml
memcached/templates/service.yaml
memcached/templates/bin/
memcached/templates/bin/_memcached.sh.tpl
memcached/templates/monitoring/
memcached/templates/monitoring/prometheus/
memcached/templates/monitoring/prometheus/exporter-configmap-bin.yaml
memcached/templates/monitoring/prometheus/exporter-deployment.yaml
memcached/templates/monitoring/prometheus/exporter-service.yaml
memcached/templates/monitoring/prometheus/bin/
memcached/templates/monitoring/prometheus/bin/_memcached-exporter.sh.tpl

sent 38,632 bytes  received 341 bytes  77,946.00 bytes/sec
total size is 49,410  speedup is 1.27
rsync -avz ../rabbitmq-operator/helm-charts/rabbitmq helm-charts
sending incremental file list
rabbitmq/
rabbitmq/.helmignore
rabbitmq/Chart.yaml
rabbitmq/installIt
rabbitmq/requirements.yaml
rabbitmq/values.yaml
rabbitmq/charts/
rabbitmq/charts/helm-toolkit-0.1.0.tgz
rabbitmq/templates/
rabbitmq/templates/configmap-bin.yaml
rabbitmq/templates/configmap-etc.yaml
rabbitmq/templates/ingress-management.yaml
rabbitmq/templates/job-image-repo-sync.yaml
rabbitmq/templates/network_policy.yaml
rabbitmq/templates/pod-test.yaml
rabbitmq/templates/service-discovery.yaml
rabbitmq/templates/service-ingress-management.yaml
rabbitmq/templates/service.yaml
rabbitmq/templates/statefulset.yaml
rabbitmq/templates/bin/
rabbitmq/templates/bin/_rabbitmq-liveness.sh.tpl
rabbitmq/templates/bin/_rabbitmq-readiness.sh.tpl
rabbitmq/templates/bin/_rabbitmq-start.sh.tpl
rabbitmq/templates/bin/_rabbitmq-test.sh.tpl
rabbitmq/templates/etc/
rabbitmq/templates/etc/_enabled_plugins.tpl
rabbitmq/templates/monitoring/
rabbitmq/templates/monitoring/prometheus/
rabbitmq/templates/monitoring/prometheus/exporter-deployment.yaml
rabbitmq/templates/monitoring/prometheus/exporter-service.yaml
rabbitmq/templates/utils/
rabbitmq/templates/utils/_to_rabbit_config.tpl

sent 46,283 bytes  received 516 bytes  93,598.00 bytes/sec
total size is 68,136  speedup is 1.46
docker build -t kubekit99/keystone-operator:poc -f build/Dockerfile .
Sending build context to Docker daemon  821.8kB
Step 1/2 : FROM kubekit99/armada-operator-dev:latest
 ---> 05ffe3152315
Step 2/2 : COPY helm-charts/ /opt/armada/helm-charts/
 ---> Using cache
 ---> 3ede624b50e3
Successfully built 3ede624b50e3
Successfully tagged kubekit99/keystone-operator:poc
docker tag kubekit99/keystone-operator:poc kubekit99/keystone-operator:latest
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadachartgroup.yaml
customresourcedefinition.apiextensions.k8s.io/armadachartgroups.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadachart.yaml
customresourcedefinition.apiextensions.k8s.io/armadacharts.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadamanifest.yaml
customresourcedefinition.apiextensions.k8s.io/armadamanifests.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadarequest.yaml
customresourcedefinition.apiextensions.k8s.io/armadarequests.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_helmrelease.yaml
customresourcedefinition.apiextensions.k8s.io/helmreleases.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_helmrequest.yaml
customresourcedefinition.apiextensions.k8s.io/helmrequests.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/helm3crd_v1beta1_lifecycleevent.yaml
customresourcedefinition.apiextensions.k8s.io/lifecyleevents.helm3crd.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/helm3crd_v1beta1_lifecycle.yaml
customresourcedefinition.apiextensions.k8s.io/lifecyles.helm3crd.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/helm3crd_v1beta1_manifest.yaml
customresourcedefinition.apiextensions.k8s.io/manifests.helm3crd.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/helm3crd_v1beta1_releaseaudit.yaml
customresourcedefinition.apiextensions.k8s.io/releaseaudits.helm3crd.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/helm3crd_v1beta1_release.yaml
customresourcedefinition.apiextensions.k8s.io/releases.helm3crd.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/helm3crd_v1beta1_values.yaml
customresourcedefinition.apiextensions.k8s.io/values.helm3crd.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/role_binding.yaml
rolebinding.rbac.authorization.k8s.io/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/role.yaml
role.rbac.authorization.k8s.io/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/service_account.yaml
serviceaccount/armada-operator created
kubectl create -f deploy/operator.yaml
deployment.apps/keystone-operator created
```

## Check the installation of the CRD and operator

```bash
kubectl get all

NAME                                   READY   STATUS    RESTARTS   AGE
pod/keystone-operator-c547fdc5-nklnd   1/1     Running   0          74s

NAME                        TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
service/keystone-operator   ClusterIP   10.97.210.61   <none>        8383/TCP   73s
service/kubernetes          ClusterIP   10.96.0.1      <none>        443/TCP    54m

NAME                                READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-operator   1/1     1            1           74s

NAME                                         DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-operator-c547fdc5   1         1         1       74s
```

```bash
kubectl get amf
No resources found.
```

```bash
kubectl get acg
No resources found.
```

```bash
kubectl get act
No resources found.
```


# Create the ArmadaChart CRD

```bash
make installmanifest
```

or

```bash
kubectl label nodes airship openstack-control-plane=enabled --overwrite
kubectl apply -f examples/keystone/simple_local.yaml
```

Result should be:

```bash
armadachart.armada.airshipit.org/helm-toolkit created
armadachart.armada.airshipit.org/mariadb created
armadachart.armada.airshipit.org/memcached created
armadachart.armada.airshipit.org/rabbitmq created
armadachart.armada.airshipit.org/keystone created
armadachartgroup.armada.airshipit.org/keystone-infra-services created
armadachartgroup.armada.airshipit.org/openstack-keystone created
armadamanifest.armada.airshipit.org/armada-manifest created
```

# Check the CRD and the deployment of the underlying helm chart

## Check the ArmachaChart Custom Resource

```bash
kubectl describe amf/armada-manifest
kubectl describe acg/keystone-infra-services
kubectl describe acg/openstack-keystone
kubectl describe act/helm-toolkit
kubectl describe act/keystone
kubectl describe act/mariadb
kubectl describe act/memcached
kubectl describe act/rabbitmq
```

```bash
kubectl describe act/keystone

Name:         keystone
Namespace:    default
Labels:       <none>
Annotations:  kubectl.kubernetes.io/last-applied-configuration:
                {"apiVersion":"armada.airshipit.org/v1alpha1","kind":"ArmadaChart","metadata":{"annotations":{},"name":"keystone","namespace":"default"},"...
API Version:  armada.airshipit.org/v1alpha1
Kind:         ArmadaChart
Metadata:
  Creation Timestamp:  2019-03-07T16:51:01Z
  Finalizers:
    uninstall-helm-release
  Generation:        2
  Resource Version:  6412
  Self Link:         /apis/armada.airshipit.org/v1alpha1/namespaces/default/armadacharts/keystone
  UID:               3089c290-40f9-11e9-a001-0800272e6982
Spec:
  Chart Name:  keystone
  Dependencies:
    helm-toolkit
  Namespace:  openstack
  Release:    keystone
  Source:
    Location:    /opt/armada/helm-charts/keystone
    Reference:   master
    Subpath:     .
    Type:        local
  Target State:
  Test:
  Upgrade:
    No Hooks:  false
    Pre:
      Delete:
        Labels:
        Name:  keystone-bootstrap
        Type:  job
  Values:
  Wait:
    Labels:
    Timeout:  100
Status:
  Actual State:
  Conditions:
    Last Transition Time:  2019-03-07T16:51:06Z
    Status:                True
    Type:                  Initialized
    Last Transition Time:  2019-03-07T16:51:08Z
    Reason:                UpdateSuccessful
    Resource Name:         keystone
    Resource Version:      120
    Status:                True
    Type:                  Deployed
  Succeeded:               true
Events:
  Type    Reason    Age                    From          Message
  ----    ------    ----                   ----          -------
  Normal  Deployed  7m54s                  act-recorder  InstallSuccessful
  Normal  Deployed  7m6s (x24 over 7m50s)  act-recorder  UpdateSuccessful
```

## Check the Keystone Service deployment

```bash
kubectl get all

NAME                                              READY   STATUS      RESTARTS   AGE
pod/keystone-api-5f6bb99ccc-jbbkk                 1/1     Running     0          5m15s
pod/keystone-bootstrap-4hsrz                      0/1     Completed   0          5m15s
pod/keystone-credential-setup-6qn46               0/1     Completed   0          5m15s
pod/keystone-db-init-wsmhg                        0/1     Completed   0          5m15s
pod/keystone-db-sync-q58k5                        0/1     Completed   0          5m15s
pod/keystone-domain-manage-wmfls                  0/1     Completed   0          5m15s
pod/keystone-fernet-setup-wtdzs                   0/1     Completed   0          5m15s
pod/keystone-operator-c547fdc5-nklnd              1/1     Running     0          8m8s
pod/keystone-rabbit-init-fnzjx                    0/1     Completed   0          5m15s
pod/mariadb-ingress-75cb44bc8c-wqclx              1/1     Running     0          5m19s
pod/mariadb-ingress-error-pages-8f44b444b-hspg6   1/1     Running     0          5m19s
pod/mariadb-server-0                              1/1     Running     0          5m19s
pod/memcached-memcached-5bc79f976c-snwq2          1/1     Running     0          5m17s
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          5m16s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone                      ClusterIP   10.102.214.109   <none>        80/TCP,443/TCP                 5m15s
service/keystone-api                  ClusterIP   10.109.2.8       <none>        5000/TCP                       5m15s
service/keystone-operator             ClusterIP   10.97.210.61     <none>        8383/TCP                       8m7s
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        61m
service/mariadb                       ClusterIP   10.111.142.153   <none>        3306/TCP                       5m19s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              5m19s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         5m19s
service/mariadb-server                ClusterIP   10.106.90.89     <none>        3306/TCP                       5m19s
service/memcached                     ClusterIP   10.97.168.13     <none>        11211/TCP                      5m17s
service/rabbitmq                      ClusterIP   10.98.23.69      <none>        5672/TCP,25672/TCP,15672/TCP   5m16s
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   5m16s
service/rabbitmq-mgr-7b1733           ClusterIP   10.96.200.211    <none>        80/TCP,443/TCP                 5m16s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-api                  1/1     1            1           5m15s
deployment.apps/keystone-operator             1/1     1            1           8m8s
deployment.apps/mariadb-ingress               1/1     1            1           5m19s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           5m19s
deployment.apps/memcached-memcached           1/1     1            1           5m17s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-api-5f6bb99ccc                 1         1         1       5m15s
replicaset.apps/keystone-operator-c547fdc5              1         1         1       8m8s
replicaset.apps/mariadb-ingress-75cb44bc8c              1         1         1       5m19s
replicaset.apps/mariadb-ingress-error-pages-8f44b444b   1         1         1       5m19s
replicaset.apps/memcached-memcached-5bc79f976c          1         1         1       5m17s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     5m19s
statefulset.apps/rabbitmq-rabbitmq   1/1     5m16s

NAME                                  COMPLETIONS   DURATION   AGE
job.batch/keystone-bootstrap          1/1           2m17s      5m15s
job.batch/keystone-credential-setup   1/1           14s        5m15s
job.batch/keystone-db-init            1/1           66s        5m15s
job.batch/keystone-db-sync            1/1           87s        5m15s
job.batch/keystone-domain-manage      1/1           2m3s       5m15s
job.batch/keystone-fernet-setup       1/1           15s        5m15s
job.batch/keystone-rabbit-init        1/1           29s        5m15s

NAME                                       SCHEDULE       SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob.batch/keystone-credential-rotate   0 0 1 * *      False     0        <none>          5m15s
cronjob.batch/keystone-fernet-rotate       0 */12 * * *   False     0        <none>          5m15s
```

# Remove the CRD 

```bash
make deletemanifest
```

or

```bash
kubectl delete -f examples/keystone/simple_local.yaml

armadachart.armada.airshipit.org "helm-toolkit" deleted
armadachart.armada.airshipit.org "mariadb" deleted
armadachart.armada.airshipit.org "memcached" deleted
armadachart.armada.airshipit.org "rabbitmq" deleted
armadachart.armada.airshipit.org "keystone" deleted
armadachartgroup.armada.airshipit.org "keystone-infra-services" deleted
armadachartgroup.armada.airshipit.org "openstack-keystone" deleted
armadamanifest.armada.airshipit.org "armada-manifest" deleted
```

Check the corresponding resources are being deleted

```bash
kubectl get all

NAME                                              READY   STATUS        RESTARTS   AGE
pod/keystone-api-5f6bb99ccc-jbbkk                 1/1     Terminating   0          9m40s
pod/keystone-operator-c547fdc5-nklnd              1/1     Running       0          12m
pod/mariadb-ingress-75cb44bc8c-wqclx              1/1     Terminating   0          9m44s
pod/mariadb-ingress-error-pages-8f44b444b-hspg6   1/1     Terminating   0          9m44s
pod/mariadb-server-0                              1/1     Terminating   0          9m44s
pod/memcached-memcached-5bc79f976c-snwq2          1/1     Terminating   0          9m42s
pod/rabbitmq-rabbitmq-0                           1/1     Terminating   0          9m41s

NAME                        TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
service/keystone-operator   ClusterIP   10.97.210.61   <none>        8383/TCP   12m
service/kubernetes          ClusterIP   10.96.0.1      <none>        443/TCP    66m

NAME                                READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-operator   1/1     1            1           12m

NAME                                         DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-operator-c547fdc5   1         1         1       12m
```

Check the ArmadaManifest, ChartGroup and Chart are no longer present
```bash
kubectl get amf
No resources found.
```

```bash
kubectl get acg
No resources found.
```

```bash
kubectl get act
No resources found.
```

# Purge artifacts from Kubernetes cluster

```bash
make purge

kubectl delete -f deploy/operator.yaml
deployment.apps "keystone-operator" deleted
kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadachartgroup.yaml
customresourcedefinition.apiextensions.k8s.io "armadachartgroups.armada.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadachart.yaml
customresourcedefinition.apiextensions.k8s.io "armadacharts.armada.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadamanifest.yaml
customresourcedefinition.apiextensions.k8s.io "armadamanifests.armada.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadarequest.yaml
customresourcedefinition.apiextensions.k8s.io "armadarequests.armada.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_helmrelease.yaml
customresourcedefinition.apiextensions.k8s.io "helmreleases.armada.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_helmrequest.yaml
customresourcedefinition.apiextensions.k8s.io "helmrequests.armada.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/helm3crd_v1beta1_lifecycleevent.yaml
customresourcedefinition.apiextensions.k8s.io "lifecyleevents.helm3crd.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/helm3crd_v1beta1_lifecycle.yaml
customresourcedefinition.apiextensions.k8s.io "lifecyles.helm3crd.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/helm3crd_v1beta1_manifest.yaml
customresourcedefinition.apiextensions.k8s.io "manifests.helm3crd.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/helm3crd_v1beta1_releaseaudit.yaml
customresourcedefinition.apiextensions.k8s.io "releaseaudits.helm3crd.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/helm3crd_v1beta1_release.yaml
customresourcedefinition.apiextensions.k8s.io "releases.helm3crd.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/helm3crd_v1beta1_values.yaml
customresourcedefinition.apiextensions.k8s.io "values.helm3crd.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/role_binding.yaml
rolebinding.rbac.authorization.k8s.io "armada-operator" deleted
kubectl delete -f ../armada-operator/chart/templates/role.yaml
role.rbac.authorization.k8s.io "armada-operator" deleted
kubectl delete -f ../armada-operator/chart/templates/service_account.yaml
serviceaccount "armada-operator" deleted
```

```
kubectl get all

NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   25m
```

Delete the configmaps related to keystone if needed
```
kubectl get configmaps

NAME                              DATA   AGE
mariadb-mariadb-mariadb-ingress   0      12m
mariadb-mariadb-state             5      12m
```

# Deployment using argo workflow

## Create the ArmadaChart and ArgoWorkflow in K8s
```bash
kubectl apply -f examples/argo/

armadachart.armada.airshipit.org/mariadb created
armadachart.armada.airshipit.org/memcached created
armadachart.armada.airshipit.org/rabbitmq created
armadachart.armada.airshipit.org/keystone created
workflow.argoproj.io/armada-manifest created
```

## Check the sequencement of the deployment

```bash
argo get armada-manifest

Name:                armada-manifest
Namespace:           default
ServiceAccount:      armada-argo-sa
Status:              Running
Created:             Fri Mar 22 11:23:24 -0500 (20 seconds ago)
Started:             Fri Mar 22 11:23:24 -0500 (20 seconds ago)
Duration:            20 seconds

STEP                          PODNAME                     DURATION  MESSAGE
 ● armada-manifest
 ├-● keystone-infra-services
 | ├---✔ enable-mariadb       armada-manifest-842022813   2s
 | ├---✔ mariadb-ready        armada-manifest-361794910   4s
 | └---◷ enable-memcached     armada-manifest-4290341812  3s        ContainerCreating
 └-✔ openstack-keystone
   ├---✔ enable-keystone      armada-manifest-2789817682  3s
   └---✔ keystone-ready       armada-manifest-587324697   5s
```

```bash
argo get armada-manifest
Name:                armada-manifest
Namespace:           default
ServiceAccount:      armada-argo-sa
Status:              Running
Created:             Fri Mar 22 11:23:24 -0500 (22 seconds ago)
Started:             Fri Mar 22 11:23:24 -0500 (22 seconds ago)
Duration:            22 seconds

STEP                          PODNAME                     DURATION  MESSAGE
 ● armada-manifest
 ├-● keystone-infra-services
 | ├---✔ enable-mariadb       armada-manifest-842022813   2s
 | ├---✔ mariadb-ready        armada-manifest-361794910   4s
 | ├---✔ enable-memcached     armada-manifest-4290341812  3s
 | └---◷ memcached-ready      armada-manifest-13034731    0s
 └-✔ openstack-keystone
   ├---✔ enable-keystone      armada-manifest-2789817682  3s
   └---✔ keystone-ready       armada-manifest-587324697   5s
```

```bash
argo get armada-manifest
Name:                armada-manifest
Namespace:           default
ServiceAccount:      armada-argo-sa
Status:              Running
Created:             Fri Mar 22 11:23:24 -0500 (25 seconds ago)
Started:             Fri Mar 22 11:23:24 -0500 (25 seconds ago)
Duration:            25 seconds

STEP                          PODNAME                     DURATION  MESSAGE
 ● armada-manifest
 ├-● keystone-infra-services
 | ├---✔ enable-mariadb       armada-manifest-842022813   2s
 | ├---✔ mariadb-ready        armada-manifest-361794910   4s
 | ├---✔ enable-memcached     armada-manifest-4290341812  3s
 | └---● memcached-ready      armada-manifest-13034731    3s
 └-✔ openstack-keystone
   ├---✔ enable-keystone      armada-manifest-2789817682  3s
   └---✔ keystone-ready       armada-manifest-587324697   5s
```

```bash
argo get armada-manifest
Name:                armada-manifest
Namespace:           default
ServiceAccount:      armada-argo-sa
Status:              Running
Created:             Fri Mar 22 11:23:24 -0500 (33 seconds ago)
Started:             Fri Mar 22 11:23:24 -0500 (33 seconds ago)
Duration:            33 seconds

STEP                          PODNAME                     DURATION  MESSAGE
 ● armada-manifest
 ├-● keystone-infra-services
 | ├---✔ enable-mariadb       armada-manifest-842022813   2s
 | ├---✔ mariadb-ready        armada-manifest-361794910   4s
 | ├---✔ enable-memcached     armada-manifest-4290341812  3s
 | ├---✔ memcached-ready      armada-manifest-13034731    4s
 | ├---✔ enable-rabbitmq      armada-manifest-1873087745  3s
 | └---● rabbitmq-ready       armada-manifest-3014489540  3s
 └-✔ openstack-keystone
   ├---✔ enable-keystone      armada-manifest-2789817682  3s
   └---✔ keystone-ready       armada-manifest-587324697   5s
```

```bash
argo get armada-manifest

Name:                armada-manifest
Namespace:           default
ServiceAccount:      armada-argo-sa
Status:              Succeeded
Created:             Fri Mar 22 11:23:24 -0500 (35 seconds ago)
Started:             Fri Mar 22 11:23:24 -0500 (35 seconds ago)
Finished:            Fri Mar 22 11:23:58 -0500 (1 second ago)
Duration:            34 seconds

STEP                          PODNAME                     DURATION  MESSAGE
 ✔ armada-manifest
 ├-✔ keystone-infra-services
 | ├---✔ enable-mariadb       armada-manifest-842022813   2s
 | ├---✔ mariadb-ready        armada-manifest-361794910   4s
 | ├---✔ enable-memcached     armada-manifest-4290341812  3s
 | ├---✔ memcached-ready      armada-manifest-13034731    4s
 | ├---✔ enable-rabbitmq      armada-manifest-1873087745  3s
 | └---✔ rabbitmq-ready       armada-manifest-3014489540  3s
 └-✔ openstack-keystone
   ├---✔ enable-keystone      armada-manifest-2789817682  3s
   └---✔ keystone-ready       armada-manifest-587324697   5s
```

## Check the state of the keystone service

Check that keystone is now up and running

```bash
NAME                                              READY   STATUS      RESTARTS   AGE
pod/armada-manifest-13034731                      0/1     Completed   0          3m27s
pod/armada-manifest-1873087745                    0/1     Completed   0          3m23s
pod/armada-manifest-2789817682                    0/1     Completed   0          3m49s
pod/armada-manifest-3014489540                    0/1     Completed   0          3m19s
pod/armada-manifest-361794910                     0/1     Completed   0          3m46s
pod/armada-manifest-4290341812                    0/1     Completed   0          3m32s
pod/armada-manifest-587324697                     0/1     Completed   0          3m44s
pod/armada-manifest-842022813                     0/1     Completed   0          3m49s
pod/keystone-api-5f6bb99ccc-kzs58                 1/1     Running     0          3m44s
pod/keystone-bootstrap-8dk7l                      0/1     Completed   0          3m44s
pod/keystone-credential-setup-8gxkr               0/1     Completed   0          3m44s
pod/keystone-db-init-sksvx                        0/1     Completed   0          3m44s
pod/keystone-db-sync-cddcd                        0/1     Completed   0          3m44s
pod/keystone-domain-manage-7njql                  0/1     Completed   0          3m44s
pod/keystone-fernet-setup-6ltxd                   0/1     Completed   0          3m44s
pod/keystone-operator-c547fdc5-rhslg              1/1     Running     0          9m14s
pod/keystone-rabbit-init-6r4tf                    0/1     Completed   0          3m44s
pod/mariadb-ingress-75cb44bc8c-6cx7v              1/1     Running     0          3m46s
pod/mariadb-ingress-error-pages-8f44b444b-npk44   1/1     Running     0          3m46s
pod/mariadb-server-0                              1/1     Running     0          3m46s
pod/memcached-memcached-5bc79f976c-gklsw          1/1     Running     0          3m27s
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          3m17s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone                      ClusterIP   10.104.46.98     <none>        80/TCP,443/TCP                 3m44s
service/keystone-api                  ClusterIP   10.97.85.189     <none>        5000/TCP                       3m44s
service/keystone-operator             ClusterIP   10.110.250.132   <none>        8383/TCP                       9m12s
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        67m
service/mariadb                       ClusterIP   10.103.9.46      <none>        3306/TCP                       3m46s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              3m46s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         3m46s
service/mariadb-server                ClusterIP   10.110.232.126   <none>        3306/TCP                       3m46s
service/memcached                     ClusterIP   10.109.129.158   <none>        11211/TCP                      3m27s
service/rabbitmq                      ClusterIP   10.107.124.53    <none>        5672/TCP,25672/TCP,15672/TCP   3m17s
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   3m17s
service/rabbitmq-mgr-7b1733           ClusterIP   10.98.214.134    <none>        80/TCP,443/TCP                 3m17s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-api                  1/1     1            1           3m44s
deployment.apps/keystone-operator             1/1     1            1           9m14s
deployment.apps/mariadb-ingress               1/1     1            1           3m46s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           3m46s
deployment.apps/memcached-memcached           1/1     1            1           3m27s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-api-5f6bb99ccc                 1         1         1       3m44s
replicaset.apps/keystone-operator-c547fdc5              1         1         1       9m14s
replicaset.apps/mariadb-ingress-75cb44bc8c              1         1         1       3m46s
replicaset.apps/mariadb-ingress-error-pages-8f44b444b   1         1         1       3m46s
replicaset.apps/memcached-memcached-5bc79f976c          1         1         1       3m27s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     3m46s
statefulset.apps/rabbitmq-rabbitmq   1/1     3m17s

NAME                                  COMPLETIONS   DURATION   AGE
job.batch/keystone-bootstrap          1/1           3m14s      3m44s
job.batch/keystone-credential-setup   1/1           15s        3m44s
job.batch/keystone-db-init            1/1           100s       3m44s
job.batch/keystone-db-sync            1/1           2m22s      3m44s
job.batch/keystone-domain-manage      1/1           3m1s       3m44s
job.batch/keystone-fernet-setup       1/1           15s        3m44s
job.batch/keystone-rabbit-init        1/1           2m1s       3m44s

NAME                                       SCHEDULE       SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob.batch/keystone-credential-rotate   0 0 1 * *      False     0        <none>          3m44s
cronjob.batch/keystone-fernet-rotate       0 */12 * * *   False     0        <none>          3m44s
```

## Cleanup

```bash
kubectl delete -f examples/argo

armadachart.armada.airshipit.org "mariadb" deleted
armadachart.armada.airshipit.org "memcached" deleted
armadachart.armada.airshipit.org "rabbitmq" deleted
armadachart.armada.airshipit.org "keystone" deleted
workflow.argoproj.io "armada-manifest" deleted
```

```bash
argo list


```bash
kubectl get all

NAME                                   READY   STATUS    RESTARTS   AGE
pod/keystone-operator-c547fdc5-rhslg   1/1     Running   0          16m

NAME                        TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
service/keystone-operator   ClusterIP   10.110.250.132   <none>        8383/TCP   16m
service/kubernetes          ClusterIP   10.96.0.1        <none>        443/TCP    74m

NAME                                READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-operator   1/1     1            1           16m

NAME                                         DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-operator-c547fdc5   1         1         1       16m
```