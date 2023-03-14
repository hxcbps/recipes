# Proyecto de API Restful de recetas en Gin de Go

Este es un proyecto de API Restful en Gin de Go que proporciona endpoints para obtener, crear, actualizar y eliminar recetas en formato JSON. La API admite las operaciones GET, POST, PUT y DELETE para las recetas.

## Requisitos

Para ejecutar este proyecto, necesita tener instalado Go en su sistema. Puede descargar e instalar Go desde el siguiente enlace: https://golang.org/dl/

## Ejecución del proyecto

Para ejecutar el proyecto, siga los siguientes pasos:

1. Clone este repositorio en su sistema:
2. Vaya al directorio del proyecto:
3. Ejecute el comando `go run main.go` para iniciar la aplicación:
4. La aplicación se ejecutará en `http://localhost:8080`.
## Endpoints

La API expone los siguientes endpoints para interactuar con las recetas:

### Obtener todas las recetas
`GET /recipes`

Este endpoint devuelve todas las recetas disponibles en formato JSON.

Ejemplo de respuesta:
```json
[
  {
    "id": 1,
    "name": "Ensalada de frutas",
    "description": "Una deliciosa ensalada de frutas",
    "ingredients": [
      "Manzanas",
      "Peras",
      "Fresas",
      "Arándanos"
    ]
  },
  {
    "id": 2,
    "name": "Hamburguesa con queso",
    "description": "Una jugosa hamburguesa con queso",
    "ingredients": [
      "Carne de res",
      "Queso cheddar",
      "Lechuga",
      "Tomate",
      "Pan de hamburguesa"
    ]
  }
]
```

### Obtener una respuesta específica

`GET /recipes/:id`

Este endpoint devuelve una receta específica en formato JSON, donde `:id` es el ID de la receta.

Ejemplo de respuesta:

```json
{
   "id":1,
   "name":"Ensalada de frutas",
   "description":"Una deliciosa ensalada de frutas",
   "ingredients":[
      "Manzanas",
      "Peras",
      "Fresas",
      "Arándanos"
   ]
}
```
### Crear una nueva receta

`POST /recipes`

Este endpoint crea una nueva receta y devuelve la receta creada en formato JSON. El cuerpo de la solicitud debe contener los siguientes campos:

- `name`: el nombre de la receta (cadena de texto)
- `description`: la descripción de la receta (cadena de texto)
- `ingredients`: los ingredientes de la receta (lista de cadenas de texto)

Ejemplo de cuerpo de solicitud:
