import { Fragment, useState} from 'react';
import './App.css';
import Boton from './components/Boton';
import Pantalla from './components/Pantalla';
import { Link } from 'react-router-dom';

const App = () => {
  const [expresion, setExpresion] = useState("");
  const [resultado, setResultado] = useState("");

  const operacion = (valor) => {
    setExpresion((expresion) => [...expresion, valor]);
  }

  const borrar = () => {
    setExpresion("");
    setResultado("");
  }

  const resultFinal = () => {
    let numPivote = 0
    let operador = ""
    let listval1 = []
    let listval2 = []
    let val1 = 0
    let val2 = 0

    for(let i = 0; i < expresion.length; i++){
      if(expresion[i] === "+" || expresion[i] === "-" || expresion[i] === "*" || expresion[i] === "/"){
        numPivote = i
        operador = expresion[i]
      }
    }
    
    for(let i = 0; i < numPivote; i++){
      listval1.push(expresion[i])
    }

    for(let i = numPivote + 1; i < expresion.length; i++){
      listval2.push(expresion[i])
    }

    val1 = listval1.join("")
    val2 = listval2.join("")

    let Estructura = {
      Val1: parseFloat(val1),
      Operador: operador,
      Val2: parseFloat(val2),
      Resultado: 0,
      Fecha: "",
      Bandera: true,
      Mensaje: ""
    }

    fetch('http://localhost:8000/operacion', {
      method: 'POST',
      headers:{
          'Content-Type': 'application/json'
      },
      body: JSON.stringify(Estructura)
    })
    .then(response => response.json())
    .then(datos => {
      setExpresion("");
      if(datos.Bandera === false){
        setResultado("Error!");  
      }else{
        setResultado(datos.Resultado);
      }
    })
  }

  return (
    <Fragment>
      <nav>
        <ul>
          <li>
            <Link to={"/"} className='linkref'>Calculadora</Link>
          </li>
          <li>
            <Link to={"/historial"} className='linkref'>Historial</Link>
          </li>
        </ul>
      </nav>
      <div className="App">
        <div className='calc-contorno'>
          <Pantalla expresion={expresion} resultado={resultado}/>
          <div className='fila'>
            <Boton symbol= "7" handleClick={operacion}/>
            <Boton symbol= "8" handleClick={operacion}/>
            <Boton symbol= "9" handleClick={operacion}/>
            <Boton symbol= "+" handleClick={operacion} color="#F39C12"/>
          </div>
          <div className='fila'>
            <Boton symbol= "4" handleClick={operacion}/>
            <Boton symbol= "5" handleClick={operacion}/>
            <Boton symbol= "6" handleClick={operacion}/>
            <Boton symbol= "-" handleClick={operacion} color="#F39C12"/>
          </div>
          <div className='fila'>
            <Boton symbol= "1" handleClick={operacion}/>
            <Boton symbol= "2" handleClick={operacion}/>
            <Boton symbol= "3" handleClick={operacion}/>
            <Boton symbol= "/" handleClick={operacion} color="#F39C12"/>
          </div>
          <div className='fila'>
            <Boton symbol= "0" handleClick={operacion}/>
            <Boton symbol= "." handleClick={operacion}/>
            <Boton symbol= "=" handleClick={resultFinal}/>
            <Boton symbol= "*" handleClick={operacion} color="#F39C12"/>
          </div>
          <Boton symbol= "Limpiar" handleClick={borrar} color="#E67E22"/>
        </div>
      </div>
    </Fragment>
  );
}

export default App;