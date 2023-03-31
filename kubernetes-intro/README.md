# Выполнено ДЗ №1

- [X] Основное ДЗ
- [X] Задание со *

## В процессе сделано

- Проанализировал, какие компоненты k8s обеспечивают отказоустойчивость ресурсов в namespace "kube-system":
  - Отказоустойчивость Pod'ов "kube-controller-manager", "kube-apiserver", "etcd" обеспечивает механизм [static pod](https://kubernetes.io/docs/tasks/configure-pod-container/static-pod>);
  - Отказоустойчивость Pod'ов "core-dns" обеспечивает контроллер Deployment;
  - Отказоустойчивость Pod'ов "kube-proxy" обеспечивает контроллер DaemonSet;
- Написал [Dockerfile](./web/Dockerfile) для сборки простого web-сервера на golang;
- [Написал манифест Pod'а](./web-pod.yaml), состоящего из двух контейнеров "init" и "web" с подключенным volumeMounts. Web-сервер работает на :8000;
- [Написал манифест Pod'а](./frontend-pod-healthy.yaml), состоящего из контейнера "frontend".
