Universidad San Carlos de Guatemala <br>
Facultad de Ingeniería <br>
Escuela de Ciencias y Sistemas <br>
Sistemas Operativos 1 - Sección A <br>
Ing. Jesús Alberto Guzmán Polanco <br>
Aux. German Jose Paz Cordon <br>
<br>
<br>
### Manual Técnico Practica 1 SO1
<br>
<br>

| Nombre | Carnet |
|--|--|
| Henry Gabriel Peralta Martinez | 201712289  |

<br>

## 
### Objetivos

 - Comprender cómo funcionan los contenedores.
 - Practicar comandos de Docker.
 - Implementar Docker-Compose.
 - Generar persistencia de datos mediante Docker Volume.
 - Utilizar Docker Hub.
 - Desarrollar scripts usando comandos de bash.

### Arquitectura
![enter image description here](https://i.ibb.co/zGK1JCz/Captura-de-pantalla-de-2023-02-19-22-21-38.png)

### Aplicación
### Frontend
![enter image description here](https://i.ibb.co/KGTCBTz/Captura-de-pantalla-de-2023-02-19-22-32-00.png)

![enter image description here](https://i.ibb.co/vDXPCXy/Captura-de-pantalla-de-2023-02-19-22-35-49.png)

Para el frontend se utilizo React como lenguaje y para este se realizaron dos pestañas, en una esta la aplicación que es una calculadora la cual puede sumar, restar, multiplicar y dividir. La segunda pestaña es una tabla en la cual se guardan todos los logs realizados. Esta aplicación se levanta en el puerto 3000 y esta aplicación hace peticiones Post y Get a otro servidor que esta montado en el puerto 8000.

### Backend
En la parte del backend se utilizo Go como lenguaje y en este se crean dos endpoints, uno de ellos es utilizado para recibir los datos enviados por la calculadora y después de recibirlos manda el resultado correcto al hacer la operación. El otro endopoint es el encargado de mostrar todos los datos guardados en la base de datos la cual se va llenando cada ves que se hace un Post al servidor. También en el servidor se crea un archivo.txt el cual servirá para que el contenedor del script pueda leerlo y poder ver los reportes. Este backend es levantado en el puerto 8000.

### Base de datos
Para la base de datos se utilizo Mysql como lenguaje. en la base de datos se creo una tabla llamada calculadora la cual es la encargada de guardar todos los datos enviados por el backend.

### Script
Para los scripts se creo un archivo.sh el cual es el encargado de leer los datos mandados desde el backend. en este muestra en consola todos los reportes requeridos por la practica.

![enter image description here](https://i.ibb.co/LDydkCT/Captura-de-pantalla-de-2023-02-19-22-55-30.png)

### Imágenes
Se necesito que crear imágenes de docker las cuales se subieron a mi repositorio de docker hub. se crearon 3 imágenes una para el frontend otra para el backend y otra para el script.

### Dockerfiles y Docker Compose
Para cada parte de la aplicación se creo un dockerfile el cual especifica las herramientas y dependencias de cada parte y estos se vuelven las imágenes ya mencionadas.
Para el docker compose se crea un archivo docker-compose.yml el cual lleva el manejo de las imágenes y como estas depende de cada una para su funcionamiento. también en este se crean los volúmenes los cuales guardan la información cada ves que este se detiene.

