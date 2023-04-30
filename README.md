# Servidor Web Go

Servidor Web hecho con lenguaje Go. 

Permite agregar middlewares, manejar rutas con los verbos más comunes. Mientras se desarrolla el servidor aprendemos a usar structs, funciones anónimas, métodos con receptor, mapas y slices y su estilo de programación concurrente.

## Librerías de Go usadas ##

- fmt: Implementa I/O formateada al estilo C.

> **Println():** Imprime usando formatos y siempre agrega una nueva línea.

- log: Implementa un simple logging. Imprime en Standar Error y agrega la fecha y hora de cada mensaje.

> **Println():** Idem al Println de fmt.

- net/http: Provee implementaciones para cliente y servidor HTTP.

> **http.HandlerFunc:** type HandlerFunc func(ResponseWriter, *Request)

> **http.Request:** Requerimiento http enviado al servidor por un cliente.

> **http.ResponseWriter:** Es usado por el handler para dar una respuesta http.

- time: Provee funciones para medir y mostrar el tiempo.

> **Now():** Retorna la hora local actual.

> **Since(t):** Devuelve el tiempo transcurrido desde t.