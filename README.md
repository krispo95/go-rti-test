# go-rti-testing
##### Настройка переменных окружения:
Для запуска нужно настроить переменную окружения APP_PORT. Например:
```bash
export APP_PORT=8099
```
##### Запуск сервиса:
Из рабочей директории выполнить команду:
```bash
docker-compose up
```

#### Пример  
Запрос на формирование оффера отправляется по роуту 
```/calculate_offer```

##### Запрос с условиями:   
```json
[
  {
    "ruleName": "technology",
    "value": "xpon"
  },
  {
    "ruleName": "internetSpeed",
    "value": "200"
  }
]
```  
##### Ответ сервиса:  
```json
{
  "name": "Игровой",
  "components": [
    {
      "name": "Интернет",
      "isMain": true,
      "prices": [
        {
          "cost": 765
        }
      ]
    }
  ],
  "totalCost": {
    "cost": 765
  }
}
```   

