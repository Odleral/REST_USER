# User REST API
REST сервер стартует на стандартном порте :80, PostgreSQL :5432
Дефолтные значения сервера находятся в файле окружения /docker/rest/.env
>SERVER_PORT=80  
SERVER_TIMEOUT_READ=5s  
SERVER_TIMEOUT_WRITE=10s  
SERVER_TIMEOUT_IDLE=15s  
DEBUG=true  
SERVER_WAIT_DB=7s  
DB_HOST_POSTGRES=postgres  
DB_PORT_POSTGRES=5432  
DB_NAME_POSTGRES=users  
DB_USER_POSTGRES=user  
DB_PASSWORD_POSTGRES=postgres

**Запуск**


    docker-compose build
    docker-compose up

>Переменная окружения SERVER_WAIT_DB задает время ожидание до запуска PostgreSQL, если возникает ошибка подключение к базе, следует увеличить время ожидание.


# Route

**GET** /api/v1/user/{uuid}

    {
    	"uuid":  "39d8f8e6-1ded-4507-b6d8-ea29813717ab",
    	"firstName":  "FirstName",
    	"lastName":  "LastName",
    	"email":  "example@gmail.com",
    	"age":  25,
    	"create_at":  "2022-02-14T14:19:08.847702Z"
    }
**POST** /api/v1/user   
**RAW body: json**

    {
	    "firstName":"FirstName",
	    "lastName":"LastName",
	    "email":"example@gmail.com",
	    "age":25
    }
**PUT** /api/v1/user/{uuid}
**RAW body: json**

    {
	    "firstName":"FirstName",
	    "lastName":"LastName",
	    "email":"example@gmail.com",
	    "age":25
    }
**DELETE**  /api/v1/user/{uuid}
