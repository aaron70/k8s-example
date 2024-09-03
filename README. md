# Docker

## 1. Crear imagenes y contenedores

Para crear contenedores con Docker necesitamos crear 
un [Dockerfile](https://docs.docker.com/reference/dockerfile/) que contenga las instruciones 
necesarias para crear la imagen.

```Dockerfile
FROM python:3-alpine
WORKDIR /app
COPY ./requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt
COPY . .
CMD [ "python", "./main.py"]
```

Para construir la imagen, utilizamos el siguiente comando:

```bash
docker build -t {NOMBRE}:{TAG} .
# docker -> Nombre del CLI docker 
# build  -> Sub-comando para crear imagenes de docker
# -t {NOMBRE}:{TAG} -> Flag para especificar un nombre para la imagen
# . -> Contexto de creación para la imagen
```

Sustituya `{NOMBRE}` por el nombre de la imagen y `{TAG}` por la version o un tag personalizado para la imagen.

> Si el tag no es especificado, se utiliza `latest` por defecto.

El contexto es la carpeta a la cual Docker tiene acceso durante la creación de la imagen para poder copiar archivos o carpetas desde la máquina **host** al **contendor**. El punto `.` significa la carpeta actual donde se está ejecutando el comando. 

## 2. Publicar imagenes a DockerHub

[DockerHub](https://hub.docker.com/) es un repositorio de imagenes, donde puedes encontrar imagenes tanto oficiales como desarrolladas por la comunidad sobre herramientas o productos (como lenguajes o bases de datos).

> Si inicias sesión puede publicar tus propias imagenes para utilizarlas con Kubernetes.

Para poder publicar tus imagenes a tu repositorio personal primero debes iniciar sesión con el siguiente comando.

```bash
docker login
# Ingresar tu usuario y contraseña
```

Luego debes nombrar tu imagen con el siguiente formato `{NAMESPACE}/{NOMBRE}:{TAG}`. Al iniciar sesión por defecto se crea un _**namespace**_ con el nombre de usuario. 

```bash
docker image tag {NOMBRE}:{TAG} {NAMESPACE}/{NOMBRE}:{TAG}
```

Por último para publicar la imagen se utiliza el siguiente comando:

```bash
docker push {NAMESPACE}/{NOMBRE}:{TAG}
```

> Puedes ver tus imagenes desde la página de DockerHub

# Kubernetes

## 1. Aplicar cambios al cluster

Para aplicar cambios al cluster de K8s utilizamos el siguiente comando:

```bash
kubectl apply -f "{PATH}"
```

Sustituya `{PATH}` por la ruta del archivo `.yaml` que contiene los recursos a crear en el cluster.

```yaml
apiVersion: apps/v1
kind: Deployment # Tipo del recurso (Pod, Service, Deployment)
metadata: 
  name: client-app # Nombre único del recurso
  labels:
    app: client # Etiqueta personalizada
spec:
  replicas: 1 # Cantidad de contenedores
  selector: 
    matchLabels: # Monitorea aquellos pods que tengan los labels aquí mostrados
      app: client 
  template:
    metadata:
      labels: # Labels para hacer match con el ReplicaSet
        app: client
    spec: # Especificación del container
      containers:
      - name: client-container
        image: aaronv70/client-app:latest
```

## 2. Ver cambios 

Una vez haya creado recursos en el cluster, puede listarlos con alguno de los siguientes comandos:

```bash
# Lista los Pods del cluster
kubectl get pods
```

```bash
# Lista los Deployments del cluster
kubectl get deployments
```

```bash
# Lista los Services del cluster
kubectl get services
```

Para ver en detalle un recurso en específico puede utilizar `describe`


```bash
# Describe un Pod del cluster
kubectl describe pod {NAME}
```

```bash
# Describe un Deployment del cluster
kubectl describe deployment {NAME}
```

```bash
# Describe un Service del cluster
kubectl describe service {NAME}
```

## 3. Eliminar recurso del cluster

Para eliminar un recurso en específico del cluster utilice:

```bash
# Elimina un Pod del cluster
kubectl delete pod {NAME}
```

```bash
# Elimina un Deployment del cluster
kubectl delete deployment {NAME}
```

```bash
# Elimina un Service del cluster
kubectl delete service {NAME}
```

Para eliminar todos los recursos de un mismo tipo pase el flag `--all`

```bash
# Elimina todos los Pods del cluster
kubectl delete pods --all
```

```bash
# Elimina todos los Deployments del cluster
kubectl delete deployments --all
```

```bash
# Elimina todos los Services del cluster
kubectl delete services --all 
```

Pude encontrar más información sobre [Kubernetest aquí](https://kubernetes.io/docs/concepts/).