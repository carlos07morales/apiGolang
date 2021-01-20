# apiGolang
Api CRUD en Go

Para ejecutar corremos el comando go run . desde una terminar

![ScreenShot](https://raw.githubusercontent.com/carlos07morales/apiGolang/main/screenshots/runGo.png)

localhost:3000
todas  las consultas son al puerto 3000 y los parámetros enviados desde POST, PUT y DELETE se envían en json Body/raw

![ScreenShot](https://raw.githubusercontent.com/carlos07morales/apiGolang/main/screenshots/setting.png)


|         Url         |   Method  |             Parameters(JSON)                 |     Return(JSON)    |
| :-----------------: |:---------:| :-------------------------------------------:|:-------------------:|
| /getAllTickets      |  GET      |                                              | Todos los Tickets   | 
| /getTicket          |  GET      |  ?id=1                                       | Un ticket           |
| /setTicket          |  POST     | { "user": "Carlos", "status": "abierto" }    | OK                  |
| /updateTicket       |  PUT      | { "id": 1, "status": "cerrado" }             | OK                  |
| /deleteTicket       |  DELETE   | { "user": "Carlos", "status": "abierto" }    | OK                  |

![ScreenShot](https://raw.githubusercontent.com/carlos07morales/apiGolang/main/screenshots/example.png)
