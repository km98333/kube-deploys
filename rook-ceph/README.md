kubectl create -f common.yaml
kubectl create -f operator.yaml
kubectl create -f cluster.yaml
kubectl create -f storageclass.yaml
kubectl patch storageclass rook-ceph-block -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'
kubectl create -f filesystem.yaml
kubectl create -f toolbox.yaml
ceph -s

#Running PODS 3 NODES
rook-ceph-agent                                     	1/1     Running     0          
rook-ceph-agent                                     	1/1     Running     0         
rook-ceph-agent                                     	1/1     Running     0          
rook-ceph-mds-fs-a		                    	1/1     Running     0          
rook-ceph-mds-fs-b               	     		1/1     Running     0          
rook-ceph-mgr-a                         		1/1     Running     0          
rook-ceph-mon-a                          		1/1     Running     0          
rook-ceph-mon-b                          		1/1     Running     0          
rook-ceph-mon-c                          		1/1     Running     0          
rook-ceph-operator                      		1/1     Running     0          
rook-ceph-osd-0                          		1/1     Running     0          
rook-ceph-osd-1                         		1/1     Running     0          
rook-ceph-osd-2                          		1/1     Running     0          
rook-ceph-osd-prepare   				0/2     Completed   0          
rook-ceph-osd-prepare   				0/2     Completed   0          
rook-ceph-osd-prepare   				0/2     Completed   0          
rook-ceph-tools                         		1/1     Running     0          
rook-discover                                      	1/1     Running     0          
rook-discover                                      	1/1     Running     0          
rook-discover                                      	1/1     Running     0
