async function actualizando_tabla_info(){
    let response = await fetch('http://34.125.14.101:8000/infoCpu')
    let infocpu = await response.json()

    console.log(infocpu)
    var str = "";

    infocpu.forEach(item =>{
        str += "<tr><td id='td2'>" + item.ProcesosEjecucion + "</td><td id='td2'>" + item.ProcesosSuspendidos + "</td><td id='td2'>" + item.ProcesosDetenidos + "</td><td id='td2'>" + item.ProcesosZombies + "</td><td id='td2'>" + item.ProcesosDesconocidos + "</td><td id='td2'>" + item.TotalProcesos + "</td></tr>"; 
    });
    document.getElementById("tabinfo").innerHTML = str;
}

actualizando_tabla_info();