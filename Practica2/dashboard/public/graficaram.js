// Set new default font family and font color to mimic Bootstrap's default styling
Chart.defaults.global.defaultFontFamily = 'Nunito', '-apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif';
Chart.defaults.global.defaultFontColor = '#858796';

const dataX_ram = [0]
const dataY_ram = [0]

function actualizando_Grafica_ram(){
    fetch('http://34.125.14.101:8000/infoRam')
    .then(response => response.json())
    .then(datos => {
        let ultimoY = datos[datos.length-1].Porcentaje
        let ultimoX = datos[datos.length-1].Tiempo
        dataX_ram.push(ultimoX);
        dataY_ram.push(ultimoY);
        if (dataX_ram.length >1000){
          dataX_ram.shift()
          dataY_ram.shift()
          Eliminar_registro()
        }
        myLineChart_actual_ram.update();
    })
}


function Ciclo_ram(){
  var saludo = function(){
    actualizando_Grafica_ram();
  };
  setInterval(saludo, 2000);
}
Ciclo_ram();

function Eliminar_registro(){
  fetch('http://34.125.14.101:8000/infoRam',{
    method: 'DELETE'
  })
  .then(response => response.json())
  .then(datos => {console.log(datos)})
}

var ctx = document.getElementById("graficaram");

var myLineChart_actual_ram = new Chart(ctx, {
  type: 'line',
  data: {
    labels: dataX_ram,
    datasets: [{
      label: "Porcentaje de utilizacion de la RAM",
      fill: true,
      lineTension: 0.3,
      backgroundColor: "rgba(255, 128, 0, 0.25)",
      borderColor: "rgba(255, 128, 0, 1)",
      pointRadius: 3,
      pointBackgroundColor: "rgba(255, 128, 0, 1)",
      pointBorderColor: "rgba(255, 128, 0, 1)",
      pointHoverRadius: 3,
      pointHoverBackgroundColor: "rgba(255, 128, 0, 1)",
      pointHoverBorderColor: "rgba(255, 128, 0, 1)",
      pointHitRadius: 10,
      pointBorderWidth: 2,
      data: dataY_ram,
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