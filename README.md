
# Golang Template




## Using Technologies

**Database:** Postgres

**Cache:** Redis

**Frameworks:** Gorm, Gin

**Token:** JWT

  
## Run it on your computer 

clone project

```bash
  git clone git@github.com:yusuftatli/golang-template.git
```

go to the root

```bash
  cd golang-template
```

Install required packages

```bash
  make install
```

run project

```bash
  make run
```
build project

```bash
  make build
```

install swagger documantation
```bash
  make install_swagger
```

  
## API Using

#### Login and get token 

```http
 POST localhost:8000/login username=&username password=&password
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **Required**. username. |
| `password` | `string` | **Required**. password. |

#### Öğeyi getir

```http
  GET /api/refresh_token/
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `token` | `string` | **Required**. token. |

  
## Docker componse

docker build

```bash
  make docker_build
```

  
## Frameworks&Database

- Gorm, Gin
- Postgres, Swagger
- JWT token

  
## Environment Variables

To run this project you will need to add the following environment variables to your .env file

`PORT`

`DEBUG`

`DATABASE_URL`

`REDIS`

`APPLICATION_HOST`

  
## Help

Email yctatli@gmail.com for support or join our Slack channel.
  
## Feedback

If you have any feedback, please contact us at yctatli@gmail.com

  
## Swagger Documantation

Following code runs to generate api models

```bash
  make models
```

  
## References

[Generate readme](https://readme.so/)


  
## Lisans

[MIT](https://choosealicense.com/licenses/mit/)

  