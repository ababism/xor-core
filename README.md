## `xor-core` - репозиторий со всем бэкендом `xority`

### Архитектура
  - `.github/workflows` - все пайплайны для `github-actions`
  - `xor-java` - монорепа со всеми `java` проектами
  - `xor-go` - монорепа со всеми `go` проектами

### Нейминг
  - сервисы называются во множественном числе без приставки `[service-name]`, например, сервис `payments`
  - папки проектов соотвествуют названиям проектов 1 в 1
  - check-еры для монореп `[mono-repo]-ci.yml`, например, `xor-java-ci.yml`
  - пайплайн со сборкой сервиса  `[service-folder]-[deploy-stage].yml`, например, `users-build-docker-image.yml`
