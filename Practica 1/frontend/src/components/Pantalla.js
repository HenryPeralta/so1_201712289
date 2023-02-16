import "./Pantalla.css";

const Pantalla = ({expresion, resultado}) => {
    return <div className="pantalla-contorno">
        <div className="expresion">
            <h3>{expresion}</h3>
        </div>
        <div className="resultado">
            <h1>{resultado}</h1>
        </div>
    </div>
}

export default Pantalla;