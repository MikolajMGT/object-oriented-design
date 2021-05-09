import React, {FC} from 'react';
import { CartStyled } from './CartStyled';
import {CartItem} from './CartItem/CartItem';
import {Button} from '@material-ui/core';
import {inject, observer} from 'mobx-react';
import {RootStore} from '../../stores/RootStore';
import {v4 as uuidv4} from 'uuid';

export const Cart: FC<{store?: RootStore}> = inject("store")(observer(({store}) => {
	const cartStore = store?.cartStore

	const prepareProducts = () => {
		return cartStore?.listProducts().map(product => {
			return (
				<CartItem
					key={uuidv4()} id={product.id} image={product.image} name={product.name} price={product.price}
					quantity={product.quantity}	amount={product.price * product.quantity} currency={product.currency}
				/>
			)
		});
	}

	const calculateTotal = () => {
		let total = 0
		cartStore?.listProducts().forEach(product => total += product.price * product.quantity)
		return total
	}

	return (
		<CartStyled>
			<div className='mt-5 entries'>
				<div className='head'>
					<h4 className='section' style={{marginLeft: '17vw'}}>Product</h4>
					<h4 className='section'>Price</h4>
					<h4 className='section'>Quantity</h4>
					<h4 className='section'>Amount</h4>
				</div>
				{prepareProducts()}
				<div className='p-4'>
					<h4 className='summary' style={{marginLeft: '51vw', display: 'inline'}}>Total: {calculateTotal().toFixed(2)} PLN</h4>
					<Button style={{height: '40px', borderRadius: '20px'}} variant="contained" color='secondary'>Complete Purchase</Button>
				</div>
			</div>
		</CartStyled>
	);
}));
