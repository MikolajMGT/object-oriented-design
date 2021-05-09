'use strict';

const data = require('../data/products')

/**
 * LIST
 */

exports.listProducts = async (req, res) => {

	res.setHeader('Content-Type', 'application/json');
	res.end(JSON.stringify(data));

	console.log(`response:\n${JSON.stringify(data)}`)

	return res;
}
