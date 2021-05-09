import React, {FC} from 'react';
import {StorePageStyled} from './StorePageStyled';
import { ProductList } from '../../components/ProductList/ProductList';
import {Page} from '../../components/Page/Cart';
import {v4 as uuidv4} from 'uuid';

export const StorePage: FC = () => {

	return (
		<Page>
			<StorePageStyled key={uuidv4()}>
				<ProductList/>
			</StorePageStyled>
		</Page>
	);
};
