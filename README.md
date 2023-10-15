# vtbMapAPI

API для нахождения оптимального офиса исходя из данных по загруженности

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ret0rn/vtbMapAPI?style=plastic)](https://github.com/ret0rn/vtbMapAPI/blob/master/go.mod)

## Client

### <a href="https://github.com/eduardpawlow/vtb-offices">Fronted client</a>

##

## Запуск

### Makefile
По дефолту  при старте makefile'ов выставляется флаг `--build`

Вы можете добавить нужные вам <a href="https://docs.docker.com/engine/reference/commandline/compose_up/#options">флаги</a>
посредством выставлении переменной `DOKERFLAGS`

 - *-d : Detached mode: Run containers in the background*

### <img src="https://www.svgrepo.com/download/474385/laptop1.svg" alt="generate short url" width="25"> Develop
Запустите make команду
```shell
make develop -sk DOKERFLAGS="-d"
```

###	<img src="https://www.svgrepo.com/download/474394/server.svg" alt="generate short url" width="25"> Production
Для запуска на сервере понадобится создать фалы переменных окружения

```shell
mkdir ./env/prod/ && \
cp ./env/develop/* ./env/prod/
```
Поставьте надежный пароль для базы данных 
и поменяйте путь подключения в конфигурации приложения 
```shell
nano ./env/prod/database.env ## изменение конфигов базы данных
nano ./env/prod/app.env ## изменение конфигов приложения
```
Запустите make команду
```shell
make prod -sk DOKERFLAGS="-d"
```


## Использование
При локальном запуске с документацией можно ознакомится по ссылке
```
http://localhost:8070/doc/index.html
```

Или с <a href="http://185.244.51.218:8070/doc/index.html">запущенным на сервере вариантом</a>


## Остановка приложения
Для остановки приложения и базы данных пропишите данную команду
```shell
make stop
```

## Tips 
Вы можете изменить лимиты потребления 
ресурсов у приложения и базы данных, посредством выставления 
нужные вам значений в файле `docker-compose.yaml` в разделе `deploy`
```shell
deploy:
  resources:
    limits:
      cpus: "1"
      memory: "1g"
```
