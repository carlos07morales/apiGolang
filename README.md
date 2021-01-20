# apiGolang
Api crud en Go

Para ejecutar corremos el comando go run . desde una terminar

localhost:3000

todas  las consultas son al puerto 3000 y los parametros enviados desde POST, PUT y DELETE se envian en json Body/raw

|         Url         |   Method  |             Parameters(JSON)                 |     Return(JSON)    |
| :-----------------: |:---------:| :-------------------------------------------:|:-------------------:|
| /getAllTickets      |  GET      |                                              | Todos los Tickets   | 
| /getTicket          |  GET      |  ?id=1                                       | Un ticket           |
| /setTicket          |  POST     | { "user": "Carlos", "status": "abierto" }    | OK                  |
| /updateTicket       |  PUT      | { "id": 1, "status": "cerrado" }             | OK                  |
| /deleteTicket       |  DELETE   | { "user": "Carlos", "status": "abierto" }    | OK                  |


