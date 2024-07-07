Выполнил Скобелкин Ярослав в рамказ [летней практики в mediasoft](https://docs.google.com/document/d/1-laS0wKfca9m3r0FOBkMI1GuZ6HSyC73/edit)

> Для запуска требуется Docker

Запуск приложения:
1. Скачайте приложение
```
   git clone https://github.com/SkobelkinYaroslav/summer_practice.git
```
2. Перейдите в директорию проекта
```
cd summer_practice
```
3. Соберите Docker образ
```
docker build -t mediasoft_summer_skobelkin . 
```

4. Запустите Docker контейнер
```
docker run -p 8080:8080 mediasoft_summer_skobelkin
```
-p 8080:8080 определяет проброс портов, где 8080 после двоеточия - порт на вашем локальном компьютере, а перед двоеточием - порт внутри контейнера. Если вам нужно изменить порт на вашем компьютере, замените первое вхождение 8080 на желаемый порт.


> Для удобства тестирования API предоставлена коллекция Postman, которую можно импортировать в Postman для быстрой проверки функциональности.