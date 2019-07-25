### Helmcharts
==============


##### 部署时注意事项
-----------
1. 要预先判断namespace是否存在，如果不存在，要创建namespace
2. 带字符串的模板转义
3. namespace要放在values.yaml中，而不是在helm命令行里传入


#### heml安装
------------------

镜像准备：


解决helm安装问题：

```bash
helm list

Error: configmaps is forbidden: User "system:serviceaccount:kube-system:default" cannot list configmap
```

创建serviceaccount
```bash
kubectl create serviceaccount --namespace kube-system tiller
```

创建cluster rolebinding
```bash
kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
```

kubernetes path
```bash
kubectl patch deploy --namespace kube-system tiller-deploy -p '{"spec":{"template":{"spec":{"serviceAccount":"tiller"}}}}'
```

helm init
```bash
helm init --service-account tiller --upgrade
```

```bash
helm list
```

```bash
```