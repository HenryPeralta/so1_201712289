const mysql = require('mysql');

const mysqlConnection = mysql.createConnection({
    host: '34.27.243.144',
    user: 'root',
    password: 'familia',
    database: 'Prac2SO1'
});

mysqlConnection.connect(function (err) {
    if(err){
        console.log(err);
        return;
    }else{
        console.log("Base de Datos conectada")
    }
});

module.exports = mysqlConnection;