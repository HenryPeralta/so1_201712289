// Set new default font family and font color to mimic Bootstrap's default styling
Chart.defaults.global.defaultFontFamily = 'Nunito', '-apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif';
Chart.defaults.global.defaultFontColor = '#858796';

const dataX_cpu = [0]
const dataY_cpu = [0]

function actualizando_Grafica_cpu(){
    fetch('http://34.118.240.56:8000/usoCpu')
    .then(response => response.json())
    .then(datos => {
        let ultimoYcpu = datos[datos.length-1].Porcentaje
        let ultimoXcpu = datos[datos.length-1].Tiempo
        dataX_cpu.push(ultimoXcpu);
        dataY_cpu.push(ultimoYcpu);
        if (dataX_cpu.length >1000){
          dataX_cpu.shift()
          dataY_cpu.shift()
          Eliminar_registro_cpu()
        }
        myLineChart_actual_cpu.update();
    })
}

function Ciclo_cpu(){
  var saludocpu = function(){
    actualizando_Grafica_cpu();
  };
  setInterval(saludocpu, 2000);
}
Ciclo_cpu();

function Eliminar_registro_cpu(){
  fetch('http://34.118.240.56:8000/usoCpu',{
    method: 'DELETE'
  })
  .then(response => response.json())
  .then(datos => {console.log(datos)})
}

var ctxcpu = document.getElementById("graficacpu");

var myLineChart_actual_cpu = new Chart(ctxcpu, {
  type: 'line',
  data: {
    labels: dataX_cpu,
    datasets: [{
      label: "Porcentaje de utilizacion de la CPU",
      fill: true,
      lineTension: 0.3,
      backgroundColor: "rgba(255, 255, 0, 0.25)",
      borderColor: "rgba(255, 255, 0, 1)",
      pointRadius: 3,
      pointBackgroundColor: "rgba(255, 255, 0, 1)",
      pointBorderColor: "rgba(255, 255, 0, 1)",
      pointHoverRadius: 3,
      pointHoverBackgroundColor: "rgba(255, 255, 0, 1)",
      pointHoverBorderColor: "rgba(255, 255, 0, 1)",
      pointHitRadius: 10,
      pointBorderWidth: 2,
      data: dataY_cpu,
    }],
  },
  options: {
    maintainAspectRatio: false,
    scales: {
      xAxes: [{
        scaleLabel: {
          display: true,
          labelString: 'Tiempo actual',
          fontSize: 15
        },
      }],
      yAxes: [{
        scaleLabel: {
          display: true,
          labelString: 'Porcentaje de uso %',
          fontSize: 15
        }
      }]
    }
  }
});