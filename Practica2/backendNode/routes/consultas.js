const express = require('express');
const router = express.Router();
const mysqlConnection = require('../database');

router.get('/infoRam', (req, res) => {
    mysqlConnection.query('select * from Ram;', (err, rows, fields) => {
        if(!err){
            res.json(rows);
        }else{
            console.log(err);
        }
    });
});

router.delete('/infoRam', (req, res) => {
    mysqlConnection.query('delete from Ram limit 500;', (err, rows, fields) => {
        if(!err){
            res.json(rows);
        }else{
            console.log(err);
        }
    });
});

router.get('/usoCpu', (req, res) => {
    mysqlConnection.query('select * from UsoCpu;', (err, rows, fields) => {
        if(!err){
            res.json(rows);
        }else{
            console.log(err);
        }
    });
});

router.delete('/usoCpu', (req, res) => {
    mysqlConnection.query('delete from UsoCpu limit 500;', (err, rows, fields) => {
        if(!err){
            res.json(rows);
        }else{
            console.log(err);
        }
    });
});

router.get('/procesosCpu', (req, res) => {
    mysqlConnection.query('SELECT cp.Id_Cpu, cp.Pidp, cp.Nombrep, cp.Estado, cp.Usuario, cp.Ram, hi.Pidh, hi.Nombreh FROM Cpu cp LEFT JOIN Hijos hi on cp.Id_Cpu = hi.Id_Cpu;', (err, rows, fields) => {
        if(!err){
            res.json(rows);
        }else{
            console.log(err);
        }
    });
});

router.get('/infoCpu', (req, res) => {
    mysqlConnection.query('select * from InfoCpu;', (err, rows, fields) => {
        if(!err){
            res.json(rows);
        }else{
            console.log(err);
        }
    });
});

router.get('/Cpu', (req, res) => {
    mysqlConnection.query('select * from Cpu;', (err, rows, fields) => {
        if(!err){
            res.json(rows);
        }else{
            console.log(err);
        }
    });
});

router.get('/Hijos', (req, res) => {
    mysqlConnection.query('select * from Hijos;', (err, rows, fields) => {
        if(!err){
            res.json(rows);
        }else{
            console.log(err);
        }
    });
});


module.exports = router;