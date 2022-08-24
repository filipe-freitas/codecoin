import { Button, Card, CardActions, CardContent, CardHeader, CardMedia, Typography } from '@material-ui/core'
import axios from 'axios'
import type { GetStaticPaths, GetStaticProps, NextPage } from 'next'
import Head from 'next/head'
import Link from 'next/link'
import React from 'react'
import http from '../../../http'
import { Product } from '../../../model'

interface ProductDetailPageProps {
  product: Product
}

const ProductDetailPage: NextPage<ProductDetailPageProps> = ({ product }) => {
  return (
    <div>
      <Head>
        <title>{product.name} - Detalhes do produto</title>
      </Head>
      <Card>
        <CardHeader
          title={product.name.toUpperCase()}
          subheader={`R$ ${product.price}`} />

        <CardActions>
          <Link
            href="/products/[slug]/order"
            as={`/products/${product.slug}/order`}
            passHref
          >
            <Button
              size='small'
              color='primary'
              component='a'>
              Comprar
            </Button>
          </Link>
        </CardActions>

        <CardMedia
          style={{ paddingTop: "56%" }}
          image={product.image_url} />

        <CardContent>
          <Typography
            component='p'
            variant='body2'
            color="textSecondary"
          >
            {product.description}
          </Typography>
        </CardContent>

      </Card>
    </div>
  )
}

export default ProductDetailPage

export const getStaticProps: GetStaticProps<ProductDetailPageProps, { slug: string }> = async(context) => {
  const { slug } = context.params!;
  try {
    const response = await http.get(`products/${slug}`);
    return {
      props: {
        product: response.data,
      },
      revalidate: 1 * 60 * 2
  }}
  catch (e) {
    if (axios.isAxiosError(e) && e.response?.status === 404) {
      return { notFound: true }
    }
    throw e;
  }
}

export const getStaticPaths: GetStaticPaths = async(context) => {
  const response = await http.get('products');
  const paths = response.data.map((product: Product) => ({
    params: {slug: product.slug}
  }));

  return { paths, fallback: 'blocking' }
}
