'use strict';

const express = require('express');
const bodyParser = require('body-parser');

/**
	Configuration
 */

const config = {
	'service.name': 'store',
	'service.host': '127.0.0.1',
	'service.port': '8000',
};

/**
	Preparation
 */

const app = express()
app.disable("x-powered-by");

const {prepareRouter} = require('../router/paths');
const cors = require('cors')

const corsOptions = {
	origin: 'http://localhost:3000',
	optionsSuccessStatus: 200
};
app.use(cors(corsOptions));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use('/', prepareRouter());

/**
	Start
 */

const server = app.listen(config['service.port'], () => {

	let host = server.address().address;
	let port = server.address().port;

	console.log('App listening at http://%s:%s', host, port);
})

