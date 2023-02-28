#!/bin/bash

archivo="./reportes/logs.txt"
#archivo="prueba.txt"

echo "Operaciones"
while IFS= read -r linea
do
    echo "$linea"
done < "$archivo"

echo ""
echo "+-------------------------+"
echo "|        REPORTE 1        |"
echo "+-------------------------+"
echo ""
size=$(wc -l < $archivo)
echo "Cantidad total de logs registrados: $size"
echo ""
echo "+-------------------------+"
echo "|        REPORTE 2        |"
echo "+-------------------------+"
echo ""
row=""
contador=0
while IFS= read -r linea
do
    row=$linea
    if [[ "${row:(-4):4}" == "cero" ]]; then
        contador=$((contador+1))
    fi
done < "$archivo"
echo "Cantidad total de operaciones que resultaron en error: $contador"
echo ""
echo "+-------------------------+"
echo "|        REPORTE 3        |"
echo "+-------------------------+"
fila=""
suma=0
resta=0
division=0
multi=0
operaciones=()
while IFS= read -r linea
do
    operaciones+=("$linea")
done < "$archivo"

for i in "${operaciones[@]}"; do
    if [[ "${i:(-4):4}" == "cero" ]]; then
        tam=${#i}
        for ((j = 0 ; j < tam - 45 ; j++)); do
            if [[ "${i:j:1}" == "+" ]]; then
                suma=$((suma+1))
            elif [[ "${i:j:1}" == "-" ]]; then
                resta=$((resta+1))
            elif [[ "${i:j:1}" == "*" ]]; then
                multi=$((multi+1))
            elif [[ "${i:j:1}" == "/" ]]; then
                division=$((division+1))
            fi
        done
    else
        tam=${#i}
        for ((j = 0 ; j < tam - 28 ; j++)); do
            if [[ "${i:j:1}" == "+" ]]; then
                suma=$((suma+1))
            elif [[ "${i:j:1}" == "-" ]]; then
                resta=$((resta+1))
            elif [[ "${i:j:1}" == "*" ]]; then
                multi=$((multi+1))
            elif [[ "${i:j:1}" == "/" ]]; then
                division=$((division+1))
            fi
        done
    fi
done
echo ""
echo "La cantidad de sumas son: $suma"
echo "La cantidad de restas son: $resta"
echo "La cantidad de multiplicaciones son: $multi"
echo "La cantidad de divisiones son: $division"
echo ""
echo "+-------------------------+"
echo "|        REPORTE 4        |"
echo "+-------------------------+"
echo ""
echo "Logs del dia de hoy:"
echo ""
line=""
diaoperaciones=()
while IFS= read -r line
do
    if [[ "${line:(-4):4}" == "cero" ]]; then
        if [[ "${line:(-45):10}" == "2023-02-21" ]]; then
            diaoperaciones+=("$line")     
        fi
    else
        if [[ "${line:(-28):10}" == "2023-02-21" ]]; then
            diaoperaciones+=("$line")     
        fi
    fi
done < "$archivo"

for i in "${diaoperaciones[@]}"; do
    if [[ "${i:(-4):4}" == "cero" ]]; then
        tam=${#i}
        echo "${i:0:tam-45}"
    else
        tam=${#i}
        echo "${i:0:tam-28}"
    fi
done
echo ""