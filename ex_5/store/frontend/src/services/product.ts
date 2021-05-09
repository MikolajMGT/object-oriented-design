import axios from 'axios';

const HOST = 'http://localhost:8000'

export const listProducts = async () => {
	return await axios.get(`${HOST}/product/list`)
}
