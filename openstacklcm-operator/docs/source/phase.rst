
Phase
================

Phase CRD
--------------------

The CRD Phase definition is available here:

1. Its Spec which is update through kubectl: `Spec <https://github.com/kubekit99/operator-ohm/blob/master/openstacklcm-operator/pkg/apis/openstacklcm/v1alpha1/osphases_types.go#L27>`_
2. Its Status which is updated by the operator and accessible through kubectl describe: `Status <https://github.com/kubekit99/operator-ohm/blob/master/openstacklcm-operator/pkg/apis/openstacklcm/v1alpha1/common_types.go#L161>`_
3. Its definition made out of the two above components: `Definition <https://github.com/kubekit99/operator-ohm/blob/master/openstacklcm-operator/pkg/apis/openstacklcm/v1alpha1/osphases_types.go#L109>`_
4. The yaml version of the CRD: `Yaml <https://github.com/kubekit99/operator-ohm/blob/master/openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_osphases.yaml>`_

Phase Controller
---------------------------

TBD

.. toctree::
   :maxdepth: 2
