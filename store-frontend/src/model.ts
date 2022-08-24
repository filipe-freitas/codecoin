export interface Product {
  id: string;
  name: string;
  description: string;
  image_url: string;
  slug: string;
  price: number;
  created_at: string;
}

export const products: Product[] = [
  {
    id: 'uuid1',
    name: 'Produto 1',
    description: 'Lorem ipsum dolor sit amet consectetur adipisicing elit.',
    price: 12.27,
    image_url: 'https://source.unsplash.com/random?product,1',
    slug: 'product1',
    created_at: '2022-06-06T00:00:00'
  },
  {
    id: 'uuid2',
    name: 'Produto 2',
    description: 'Lorem ipsum dolor sit amet consectetur adipisicing elit.',
    price: 24.59,
    image_url: 'https://source.unsplash.com/random?product,2',
    slug: 'product2',
    created_at: '2022-06-06T00:00:00'
  },
  {
    id: 'uuid3',
    name: 'Produto 3',
    description: 'Lorem ipsum dolor sit amet consectetur adipisicing elit.',
    price: 37.90,
    image_url: 'https://source.unsplash.com/random?product,3',
    slug: 'product3',
    created_at: '2022-06-06T00:00:00'
  },
  {
    id: 'uuid4',
    name: 'Produto 4',
    description: 'Lorem ipsum dolor sit amet consectetur adipisicing elit.',
    price: 40.00,
    image_url: 'https://source.unsplash.com/random?product,4',
    slug: 'product4',
    created_at: '2022-06-06T00:00:00'
  },
]