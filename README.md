# GoNewsAgretor
# Агрегатор
# Агрегатор обрабатывает запросы от API Gateway, позволяя осуществлять поиск по новостям и предоставляет функции пагинации результатов.

1. Путь к исходному коду: ./aggregator
 Порт: 8081
2. Dockerfile: Включён в директорию службы.
Как запустить
3. Перейдите в папку, содержащую Dockerfile Aggregator:
   cd aggregator
4. Соберите Docker образ:
   docker build -t aggregator .
5. Запустите контейнер:
   docker run -p 8081:8081 aggregator
