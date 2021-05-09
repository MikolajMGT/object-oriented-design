'use strict';

let express = require('express');
const products = require('../controller/product');

exports.prepareRouter = () => {
	let router = express.Router();

	router.get('/product/list', products.listProducts);

	return router;
}
