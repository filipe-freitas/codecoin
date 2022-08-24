import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { EntityNotFoundError, Repository } from 'typeorm';
import { validate as uuidValidate } from 'uuid';
import { CreateProductDto } from './dto/create-product.dto';
import { UpdateProductDto } from './dto/update-product.dto';
import { Product } from './entities/product.entity';

@Injectable()
export class ProductsService {

  constructor(
    @InjectRepository(Product) private productRepository: Repository<Product>
  ){}

  create(createProductDto: CreateProductDto) {
    const product = this.productRepository.create(createProductDto);
    return this.productRepository.save(product);
  }

  findAll() {
    return this.productRepository.find();
  }

  async findOne(idOrSlug: string) {
    const where = uuidValidate(idOrSlug)
      ? { id: idOrSlug }
      : { slug: idOrSlug };
    const product = await this.productRepository.findOneBy(where);
    if (!product) {
      throw new EntityNotFoundError(Product, idOrSlug);
    }
    return product;
  }

  async update(id: string, updateProductDto: UpdateProductDto) {
    const updateResult = await this.productRepository.update(id, UpdateProductDto);
    if (!updateResult.affected) {
      throw new EntityNotFoundError(Product, id);
    }
    return this.productRepository.findOneBy({
      id: id,
    });
  }

  async remove(id: string) {
    const deleteResult = await this.productRepository.delete(id);
    if (!deleteResult.affected) {
      throw new EntityNotFoundError(Product, id);
    }
  }
}
