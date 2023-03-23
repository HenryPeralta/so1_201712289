var listacompleta = []

async function actualizando_tabla(){
 
    let response = await fetch('http://34.118.240.56:8000/Cpu')
    let datoscpu = await response.json()

    let response2 = await fetch('http://34.118.240.56:8000/Hijos')
    let datoshijos = await response2.json()

    console.log(datoscpu)

    var str = "";
    var contador = 0
    var contador2 = 0
    var bandera = false
    datoscpu.forEach(item =>{
        for(let i = 0; i < datoshijos.length; i++){
            if(item.Id_Cpu == datoshijos[i].Id_Cpu){
                bandera = true
                break
            }
            
            bandera = false
        }
        if(bandera == true){
            var listahijos = []
            str += "<tr><td><button id='btn' class='desplegar' onclick=\"toggle(this)\"></button>" + "</td><td>" + item.Id_Cpu + "</td><td>" + item.Pidp + "</td><td>" + item.Nombrep + "</td><td>" + item.Estado + "</td><td>" + item.Usuario + "</td><td>" + item.Ram + "</td></tr>"; 
            str += "<div id='contenthijos' class='contenthijos'>"
            str += "<tr id='mytr" + contador + "' class='tr-hijos'><th class='th-hijos'>Id_Cpu</th><th class='th-hijos'>Pid Hijo</th><th class='th-hijos'>Nombre Hijo</th></tr>" 
            for(let i = 0; i < datoshijos.length; i++){
                if(item.Id_Cpu == datoshijos[i].Id_Cpu){
                    str += "<tr id='myth" + contador2 + "' class='tr-hijos'><td>" + datoshijos[i].Id_Cpu + "</td><td>" + datoshijos[i].Pidh + "</td><td class='td-hijos'>" + datoshijos[i].Nombreh  + "</td></tr>"; 
                    listahijos.push(contador2)
                    contador2 = contador2 + 1
                }
            }
            listacompleta.push(listahijos)
            str += "</div>"
            contador = contador + 1
        }else{
            str += "<tr><td>" + "</td><td>" + item.Id_Cpu + "</td><td>" + item.Pidp + "</td><td>" + item.Nombrep + "</td><td>" + item.Estado + "</td><td>" + item.Usuario + "</td><td>" + item.Ram + "</td></tr>"; 
        }
    });
    document.getElementById("tabla").innerHTML = str;
    console.log(listacompleta)

}

let toggle = button => {
    var coll = document.getElementsByClassName("desplegar");
    var val = 0;
    for (let i = 0; i < coll.length; i++) {
    coll[i].addEventListener("click", function() {
        val = i
        console.log(val)
        let id = "mytr" + val
        console.log(id)
        let element = document.getElementById(id);
        let hidden = element.getAttribute("hidden");
    
        if (hidden) {
           element.removeAttribute("hidden");
           button.innerText = "Cerrar";
        } else {
           element.setAttribute("hidden", "hidden");
           button.innerText = "Abrir";
        }
        for(let i = 0; i < listacompleta[val].length; i++){
            console.log(listacompleta)
            let ids = "myth" + listacompleta[val][i]
            console.log(ids)
            let element = document.getElementById(ids);
            let hidden = element.getAttribute("hidden");
    
            if (hidden) {
                element.removeAttribute("hidden");
                button.innerText = "Cerrar";
            } else {
                element.setAttribute("hidden", "hidden");
                button.innerText = "Abrir";
            }
        }
    })
    }
}

actualizando_tabla();

