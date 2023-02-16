import { Fragment, useState } from "react";
import { Link } from 'react-router-dom';
import "./Historial.css";

const Historial = () => {

    const [lista, setLista] = useState([]);

    fetch('http://localhost:8000/historial')
    .then(response => response.json())
    .then(datos => {
        setLista(datos)
    })

    return(
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
        <div>
            <table className="tablaLogs">
                <thead>
                    <tr>
                        <th className="col">Valor_1</th>
                        <th className="col">Operador</th>
                        <th className="col">Valor_2</th>
                        <th className="col">Resultado</th>
                        <th className="col">Fecha</th>
                        <th className="col">Mensaje</th>
                    </tr>
                </thead>
                <tbody>
                    {lista.map((item, index) => {
                                return(
                                    <tr key={index}>
                                        <td>{item.Val1}</td>
                                        <td>{item.Operador}</td>
                                        <td>{item.Val2}</td>
                                        <td>{item.Resultado}</td>
                                        <td>{item.Fecha}</td>
                                        <td>{item.Mensaje}</td>
                                    </tr>
                                )
                            }
                        )}
                </tbody>
            </table>
        </div>
        </Fragment>
    )
};

export default Historial;