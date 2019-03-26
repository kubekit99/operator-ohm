
ArmadaManifest
===============

ArmadaManifest CRD
------------------

The ArmadaManifest defintion used in production is available here: `production <https://github.com/openstack/airship-armada/blob/master/armada/schemas/armada-chart-schema.yaml>`_

The CRD ArmadaManifest definition is available here:

1. Its Spec which is update through kubectl: `Spec <https://github.com/kubekit99/operator-ohm/blob/master/armada-operator/pkg/apis/armada/v1alpha1/armadachart_types.go#L27>`_
2. Its Status which is updated by the operator and accessible through kubectl describe: `Status <https://github.com/kubekit99/operator-ohm/blob/master/armada-operator/pkg/apis/armada/v1alpha1/common_types.go#L161>`_
3. Its definition made out of the two above components: `Definition <https://github.com/kubekit99/operator-ohm/blob/master/armada-operator/pkg/apis/armada/v1alpha1/armadachart_types.go#L109>`_
4. The yaml version of the CRD: `Yaml <https://github.com/kubekit99/operator-ohm/blob/master/armada-operator/chart/templates/armada_v1alpha1_armadachart.yaml>`_

ArmadaManifest Controller
-------------------------

TBD

.. toctree::
   :maxdepth: 2
