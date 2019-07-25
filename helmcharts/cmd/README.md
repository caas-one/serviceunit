#### ycectl支持的命令
--------------

##### setup: 根据CICD环境配置生成profile文件
--------------
```bash
ycectl setup -o ${path_to_profile}
```

#####  build: 根据profile生成对应的values/templates/charts等文件
---------------
```bash
ycectl build -f ${path_to_profile}
```

##### init:  执行helm init
-----------
```bash
ycectl init
```

##### update: 修改profile.yaml文件后，升级应用
----------
```bash
ycectl update
```

##### clean: 清除现在已经安装的helm/charts
-----------
```bash
ycectl clean
```
