![build](https://github.com/khostya/backend-bootcamp-assignment-2024/actions/workflows/build.yml/badge.svg?event=push)

```sh
docker compose up -d postgres
make migration-up
docker compose up -d server
```

1. в задании /login должен происходить по почте, а в api указан id. я выбрал вариант api
2. ничего не сказано об уникальности дома. наверное, потому что это очевидно
3. при dummy login подписка на добавление квартир не работает, потому что сообщение отправляется на почту
4. во многих местах api, некоторые поля не указаны как required. 
решил сделать на свое усмотрение, при /register без почты точно не обойтись.
5. номер квартиры = количество квартир + 1

# Тестовое задание для отбора на Avito Backend Bootcamp
## Сервис домов

На Авито ежедневно публикуются тысячи объявлений о продаже или аренде недвижимости. Они попадают в каталог домов, в котором пользователь может выбрать жильё по нужным параметрам в понравившемся доме.  

Прежде чем попасть в каталог, объявление проходит тщательную модерацию, чтобы в нём не было недопустимого контента. 
## Описание задачи

Разработайте backend-сервис, с помощью которого пользователь сможет продать квартиру, загрузив объявление на Авито. 
### Функционал сервиса: 
1. **Авторизация пользователей:**
   1. Используя ручку /dummyLogin и передав в неё желаемый тип пользователя (client, moderator), сервис в ответе вернет токен с соответствующим уровнем доступа — обычного пользователя или модератора. Этот токен нужно передавать во все endpoints, требующие авторизации. **[OK]**  
    2. Регистрация и авторизация пользователей по почте и паролю:
      - При регистрации используется endpoint /register. В базе создаётся и сохраняется новый пользователь желаемого типа: обычный пользователь (client) или модератор (moderator). **[OK]**
      - У созданного пользователя появляется токен endpoint /login. При успешной авторизации по почте и паролю возвращается токен для пользователя с соответствующим ему уровнем доступа. **[OK]**
2. **Создание дома:**
   1. Только модератор имеет возможность создать дом используя endpoint /house/create. В случае успешного запроса возвращается полная информация о созданном доме **[OK]**
3. **Создание квартиры:**
   1. Создать квартиру может любой пользователь, используя endpoint /flat/create. При успешном запросе возвращается полная информация о квартире. **[OK]**  
    2. Если жильё успешно создано через endpoint /flat/create, то объявление получает статус модерации created. **[OK]**  
    3. У дома, в котором создали новую квартиру, обновляется дата последнего добавления жилья. **[OK]**
4. **Модерация квартиры:**
   1. Статус модерации квартиры может принимать одно из четырёх значений: created, approved, declined, on moderation. **[OK]**  
    2. Только модератор может изменить статус модерации квартиры с помощью endpoint /flat/update. При успешном запросе возвращается полная информация об обновленной квартире. **[OK]**
5. **Получение списка квартир по номеру дома:**
   1. Используя endpoint /house/{id}, обычный пользователь и модератор могут получить список квартир по номеру дома. Только обычный пользователь увидит все квартиры со статусом модерации approved, а модератор — жильё с любым статусом модерации. **[OK]**
6. **Дополнительно.** **Подписка на уведомления о новых квартирах в доме по его номеру. Обычный пользователь может подписаться на такие уведомления с помощью endpoint /house/{id}/subscribe.**

### Общие вводные

У сущности **«Дом»** есть:  
- уникальный номер дома **[OK]**
- адрес **[OK]**
- год постройки **[OK]**
- застройщик (у 50% домов) **[OK]**
- дата создания дома в базе **[OK]**
- дата последнего добавления новой квартиры дома **[OK]**

У сущности **«Квартира»** есть:
- номер квартиры **[OK]**
- цена (целое число) **[OK]**
- количество комнат **[OK]**

**Связи между сущностями:**
1. Каждая квартира может иметь только одно соответствие с домом (один к одному). **[OK]**
2. Номер дома служит уникальным идентификатором самого дома. **[OK]**
3. **Номер квартиры не является уникальным идентификатором. Например, квартира №1 может  находиться как в доме №1, так и в доме №2, и в этом случае это будут разные квартиры.**  **[OK]**

Список квартир в доме — ключевая функция, которой пользуются: 
- Модераторы — получают полный список всех объявлений в доме вне зависимости от статуса модерации.  **[OK]**
- Пользователи — получают список только прошедших модерацию объявлений.  **[OK]**
  Важно гарантировать быстрый отклик endpoint для пользователей.  **[OK]**
## Условия
1. Используйте этот [API](https://github.com/avito-tech/backend-bootcamp-assignment-2024/blob/main/api.yaml). **[OK]**
2. Реализация пользовательской авторизаций не является обязательным условием.  Поэтому токен авторизации можете получить из метода /dummyLogin, описанного в пункте 1 API. В параметрах запроса можно выбрать роль пользователя: модератор или обычный пользователь. В зависимости от роли будет сгенерирован токен с определённым уровнем доступа. **[OK]**
3. Всего квартир (до 20kk), RPS — 1k, SLI времени ответа — 50 мс, SLI успешности ответа — 99.99%
4. Для авторизации доступов должны использоваться два вида токенов: обычного пользователя и модератора. Получение списка квартир и публикация нового объявления может происходить с помощью токена пользователя или модератора, остальные действия могут выполняться только с помощью токена модератора. **[OK]**
5. Разработайте интеграционные и модульные тесты для сценариев получения списка квартир и процесса публикации новой квартиры.
6. Квартира может не пройти модерацию. В таком случае её видят только модераторы. **[OK]** 
7. Работать с сервисом могут несколько модераторов. При этом конкретную квартиру может проверять только один модератор. Перед началом работы нужно перевести квартиру в статус on moderate — тем самым запретив брать её на проверку другим модераторам. В конце квартиру переводят в статус approved или declined. **[OK]**
## Дополнительные задания:

*Не являются обязательными, но дадут вам преимущество перед другими кандидатами.*
1. Реализовать пользовательскую авторизацию по методам /register и /login. **[OK]**
2. Реализовать асинхронный механизм уведомления пользователя о появлении новых квартир в доме по почте, метод /house/{id}/subscribe. **[OK]**
   Саму отправку писем разрабатывать не нужно, вместо этого стоит использовать готовый интерфейс:

```go
package sender

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Sender struct{}

func New() *Sender {
	return &Sender{}
}

func (s *Sender) SendEmail(ctx context.Context, recipient string, message string) error {
	// Имитация отправки сообщения
	duration := time.Duration(rand.Int63n(3000)) * time.Millisecond
	time.Sleep(duration)

  	// Имитация неуспешной отправки сообщения
	errorProbability := 0.1
	if rand.Float64() < errorProbability {
		return errors.New("internal error")
	}

	fmt.Printf("send message '%s' to '%s'\n", message, recipient)

	return nil
}
```

3. Настроить CI (например, через github actions или travis), по итогам которого:
   1. в README.md корня вашего репозитория будет отображён бейдж со статусом билда (успех или неудача) в текущей ветке **[OK]**
   2. в README.md корня вашего репозитория отображен бейдж с процентом покрытия кода тестами в текущей ветке: <img width="398" alt="image" src="https://github.com/user-attachments/assets/b8a2838a-078e-4164-ba9a-477468cf4c48">
4. **Настроить логгер**
5. **Настроить кодогенерацию DTO endpoint'ов по openapi схеме** **[OK]**
## Требования по стеку

- **Язык сервиса:** – Go. **[OK]**
- **База данных:** – предпочтительно PostgreSQL, но можно выбрать другую удобную вам. Нельзя использовать ORM для взаимодействия с базой. Допустимо использовать билдеры для запросов, например, такой: https://github.com/Masterminds/squirrel **[OK]**
- **Для деплоя зависимостей и самого сервиса** нужно использовать Docker или Docker & DockerCompose. **[OK]**
## Дополнения к решению

Если у вас возникнут вопросы, ответов на которые нет в условиях, то принимать решения можете самостоятельно.  В таком случае приложите к проекту Readme-файл со списком вопросов и объяснениями своих решений.
## Оформление и отправка решения

Создайте публичный git-репозиторий на любом хосте (GitHub, GitLab и другие), содержащий в master/main ветке:
1. Код сервиса;
2. Docker или Docker & DockerCompose или описанную в Readme.md инструкцию по запуску;
3. Описанные в Readme.md вопросы или проблемы, с которыми вы столкнулись, и описание своих решений.

Прикрепите ссылку на репозиторий в письмо на почте от организаторов программы.
