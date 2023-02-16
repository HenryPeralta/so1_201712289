create database EjemploDB;
use EjemploDB;

create table calculadora(
	val1 float,
    operador varchar(3),
    val2 float,
    resultado float,
    fecha varchar(100),
    bandera bool,
    mensaje varchar(100)
);