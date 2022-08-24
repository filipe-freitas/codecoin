import { TableForeignKey } from "typeorm";
import { Table } from "typeorm"
import { MigrationInterface, QueryRunner } from "typeorm"

export class CreateOrderItemsTable1661271028951 implements MigrationInterface {

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.createTable(
            new Table({
                name: 'order_items',
                columns: [{
                    name: 'id',
                    type: 'uuid',
                    isPrimary: true,
                },{
                    name: 'order_id',
                    type: 'uuid',
                },{
                    name:'product_id',
                    type: 'uuid',
                }, {
                    name: 'quantity',
                    type: 'int',
                }, {
                    name: 'price',
                    type: 'double precision',
                }]
            })
        );

        await queryRunner.createForeignKey(
            'order_items',
            new TableForeignKey({
                name: 'order_items_order_id_FK',
                referencedTableName: 'orders',
                referencedColumnNames: ['id'],
                columnNames: ['order_id']
            }));

        await queryRunner.createForeignKey(
            'order_items',
            new TableForeignKey({
                name: 'order_items_product_id_FK',
                referencedTableName: 'products',
                referencedColumnNames: ['id'],
                columnNames: ['product_id']
            }));
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.dropForeignKey('order_items', 'order_items_product_id_FK');
        await queryRunner.dropForeignKey('order_items', 'order_items_order_id_FK');
        await queryRunner.dropTable('order_items');
    }

}
