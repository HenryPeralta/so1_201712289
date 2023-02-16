import "./Boton.css";

const Boton = ({symbol, color, handleClick}) => {
    return (
        <div onClick={() => handleClick(symbol)} className="boton-contorno" style={{backgroundColor: color}}>
            {symbol}
        </div>
    );
}

export default Boton;