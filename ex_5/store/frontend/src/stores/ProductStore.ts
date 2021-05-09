import {RootStore} from './RootStore';
import {makeAutoObservable} from 'mobx';
import {ProductProps} from '../components/Product/Product';
import {listProducts} from '../services/product';

interface IProductStore {
	products: ProductProps[]
}

export class ProductStore implements IProductStore {
	private rootStore: RootStore | undefined;

	products: ProductProps[] = [];

	constructor(rootStore?: RootStore) {
		makeAutoObservable(this)
		this.rootStore = rootStore;
	}
	
	listProducts = async () => {
		if (this.products.length === 0) {
			const productList = await listProducts()
			this.products = productList.data.data as ProductProps[]
		}
		return this.products
	}
}
