# drone-runner-kube

This is a special runtime to execute the drone pipeline on virtual-kubelet (sure on k8s).

## Differences With Official Version

* No Update (wont update the pod when running pipeline, so this will be little faster)

* All Images In The Step Should Have Shell
