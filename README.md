## `xor-core` - репозиторий со всем бэкендом `xority`

### Архитектура
  - `.github/workflows` - все пайплайны для `github-actions`
  - `xor-go` - монорепа со всеми `go` сервисами
  - `xor-java` - монорепа со всеми `java` сервисами
  - `xor-py` - монорепа со всеми `py` сервисами

### .proto зависимости
Все `.proto` файлы хранятся в корне репозитория в папке `./proto/service-name`. Чтобы собрать необходимые
зависимости, необходимо выполнить скрипт `./xor-python/scripts/prepare_protos.py`, который соберет протобуфы
и положит в подходящее место.

Доступные флаги:
- `--xor-go` - копирует содержимое корневой папки `./proto` в `./xor-go/proto` и собирает протобуфы.
- `--xor-java` - копирует содержимое корневой папки в соотвествующие сборочные проекты с постфиксом `-proto`.
Пример: `./proto/idm/..` -> `./xor-java/idm-proto/src/main/proto/idm/..`.

### Документация 
 - `xor-go`: [Go_general_readme.md](https://github.com/xority-space/xor-core/blob/master/xor-go/XOR_GO_README.md)

### Нейминг
  - сервисы называются во множественном числе без приставки `[service-name]`, например, сервис `payments`
  - папки проектов соотвествуют названиям проектов 1 в 1
  - check-еры для монореп `[mono-repo]-ci.yml`, например, `xor-java-ci.yml`
  - пайплайн со сборкой сервиса  `[service-folder]-[deploy-stage].yml`, например, `users-build-docker-image.yml`
