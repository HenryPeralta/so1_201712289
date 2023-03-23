Universidad San Carlos de Guatemala <br>
Facultad de Ingeniería <br>
Escuela de Ciencias y Sistemas <br>
Sistemas Operativos 1 - Sección A <br>
Ing. Jesús Alberto Guzmán Polanco <br>
Aux. German Jose Paz Cordon <br>

### Manual Técnico Practica 2 SO1

| Nombre | Carnet |
|--|--|
| Henry Gabriel Peralta Martinez | 201712289  |

## 
### Objetivos

 - Conocer el kernel de Linux y sus módulos.
 - Comprender y utilizar herramientas de Google Cloud Platform.
 - Utilizar máquinas virtuales de GCP.
 - Creación de módulos.
 - Utilizar CloudSQL.
 
### Arquitectura

![enter image description here](https://i.ibb.co/4VGkyBR/Captura-de-pantalla-de-2023-03-22-19-29-35.png)

### Aplicación
### Frontend

![enter image description here](https://i.ibb.co/ZXK4k3Q/Captura-de-pantalla-de-2023-03-22-19-20-03.png)

![enter image description here](https://i.ibb.co/0M2jpWY/Captura-de-pantalla-de-2023-03-22-19-20-11.png)

![enter image description here](https://i.ibb.co/2kQYhcb/Captura-de-pantalla-de-2023-03-22-19-20-16.png)

Para la creacion del frontend se utilizo react js, se creo un dashbord el cual simula el gestor de tareas, se crearon 2 graficas las cuales muestran el porcentaje de uso de la memoria Ram y el uso del Cpu. Tambien se creo una tabla la cual muestra todos los estados de los procesos de la Cpu. y por ultimo hay una tabla que muestra todos los procesos activos y tambien muestra sus sub procesos.
Esta aplicación se levanta en el puerto 3000 y esta aplicación hace peticiones Get y Delete a otro servidor que esta montado en el puerto 8000.

### Backend Node
Para este backend se utilizo node js, el cual obtiene los datos de la base de datos, y de ahi este crea endpoints para que el frontend consuma los datos. Este se levanta en el puerto 8000.

### Backend Go
Para este backend se utilizo go, el cual ejecuta la accion de cat para obtener los datos de la cpu, despues de obtener los resultados este los inserta a la base de datos de mysql y como esta en un bucle pasa recibiendo datos cada segundo.

### Base de datos
Para la base de datos se utilizo mysql y esta esta creada en google cloud plattaform, esta tiene tablas las cuales guardan los datos recibidos de go y para que posteriormente node obtengas los datos.

### Modulos 
Se crearon 2 modulos uno para la cpu y otro para la ram. se utulizaron librerias de C para obtener los datos del kernel de linux

### Google Cloud Plataform
Se utilizo el servicio de la nube de google (GCP) para crear 2 instancias vm y 1 instancia SQL para que el proyecto se ejecute.  

![enter image description here](https://i.ibb.co/JsGzzRh/Captura-de-pantalla-de-2023-03-22-19-20-32.png)
