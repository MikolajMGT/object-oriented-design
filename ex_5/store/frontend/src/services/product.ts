import axios from 'axios';

const HOST = 'http://localhost:8000'

export const listProducts = async () => {
	return axios.get(`${HOST}/product/list`)
}
