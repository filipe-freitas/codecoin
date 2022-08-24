import { Avatar, Box, Button, Grid, ListItem, ListItemAvatar, ListItemText, TextField, Typography } from '@material-ui/core'
import axios from 'axios'
import type { GetServerSideProps, NextPage } from 'next'
import Head from 'next/head'
import React from 'react'
import http from '../../../http'
import { Product } from '../../../model'

interface OrderPageProps {
  product: Product
}

const ProductDetailPage: NextPage<OrderPageProps> = ({ product }) => {
  return (
    <div>
      <Head>
        <title>Pagamento</title>
      </Head>

      <Typography
        component="h1"
        variant='h3'
        color='textPrimary'
        gutterBottom
      >
        Checkout
      </Typography>

      <ListItem>
        <ListItemAvatar>
          <Avatar src={ product.image_url }/>
        </ListItemAvatar>
        <ListItemText
          primary={product.name}
          secondary={`R$ ${product.price}`}
        />
      </ListItem>

      <Typography
        component="h2"
        variant='h6'
        gutterBottom
      >
        Pague com cartão de crédito
      </Typography>

      <form>
        <Grid
          container
          spacing={3}
        >
          <Grid item xs={12} md={6}>
            <TextField
              label="Nome"
              required
              fullWidth
            />
          </Grid>
          <Grid item xs={12} md={6}>
            <TextField
              label="Número do cartão"
              required
              inputProps={{ maxLength: 16 }}
              fullWidth
            />
          </Grid>
          <Grid item xs={12} md={6}>
            <TextField
              label="CVV"
              required
              type="number"
              fullWidth
            />
          </Grid>
          <Grid item xs={12} md={6}>
            <Grid container spacing={3}>
              <Grid item xs={6} md={6}>
                <TextField
                  label="Expiração mês"
                  required
                  type="number"
                  fullWidth
                />
              </Grid>
              <Grid item xs={6} md={6}>
                <TextField
                  label="Expiração ano"
                  required
                  type="number"
                  fullWidth
                />
              </Grid>
            </Grid>
          </Grid>
        </Grid>
        <Box marginTop={1}>
          <Button
            type="submit"
            variant="contained"
            color="primary"
            fullWidth
          >
            Pagar
          </Button>
        </Box>
      </form>
    </div>
  )
}

export default ProductDetailPage

export const getServerSideProps: GetServerSideProps<OrderPageProps, { slug: string }> = async(context) => {
  const { slug } = context.params!;
  try {
    const response = await http.get(`products/${slug}`);
    return {
      props: {
        product: response.data,
      }
  }}
  catch (e) {
    if (axios.isAxiosError(e) && e.response?.status === 404) {
      return { notFound: true }
    }
    throw e;
  }
}
