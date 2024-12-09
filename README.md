###  RETO TECNICO - DragonBall
```bash
-Disponibilizar un servicio go en local
-Conectar a una base de datos (al gusto)
-Crear un endpoint que permita crear un personaje con 3 datos basicos + el id de personaje
-Consumir un servicio externo que triaga la información del personaje y guardar 3 campos adicionales
-Obtenga el nombre del personaje desde el campo name... suponga que aveces el nombre llega como string en el campo "character" y prepare el código para que lo pueda obtener de las dos formas.
-Diagrame la solución
 
https://web.dragonball-api.com/

Según lo conversado, solo se pedirá el nombre en el endpoint (POST) y se agregaran 2 datos más obtenidos desde la API externa

```
## Requisitos

Para ejecutar este proyecto, necesitas tener instalados los siguientes componentes:

- [Go 1.21.10](https://go.dev/dl/) 
- [MongoDB](https://www.mongodb.com/) 
- [Docker](https://www.docker.com/) 

## Instalación

### 1. Clonar el repositorio

```bash
git clone https://github.com/mauri247/db-api.git
cd db-api
```
### 2. Levantar el proyecto
```bash
docker-compose up --build  
```

### 3. Probar endpoint
```bash
curl --location 'http://localhost:8080/characters' \
--header 'Content-Type: application/json' \
--data '{
    "name":"krillin"
}'

VERIFICAR METODO POST
```

### 4. Endpoint adicional (GET)
```bash
curl --location 'http://localhost:8080/characters/search?name=GOku'

VERIFICAR METODO GET
```

### 5. Revisar base de datos 
```bash
docker exec -it mongodb mongosh
use dragonball-db
db.characters.find().pretty()
```

