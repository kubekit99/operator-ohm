#!/bin/bash
for i in delete install operational planning rollback test trafficdrain trafficrollout upgrade
do
sed -e "s/phase/$i/g" _wf-phase.yaml > wf-${i}.yaml
done
