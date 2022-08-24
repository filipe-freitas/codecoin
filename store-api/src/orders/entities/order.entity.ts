import { BeforeInsert, Column, CreateDateColumn, Entity, OneToMany, PrimaryGeneratedColumn } from "typeorm";
import { v4 as uuidv4 } from 'uuid';
import { CreditCard } from "./credit-card.embedded";
import { OrderItem } from "./order-item.entity";

export enum OrderStatus {
  Approved = "approved",
  Pending = "pending",
}

@Entity({ name: 'orders' })
export class Order {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Column()
  total: number;

  @Column()
  status: OrderStatus = OrderStatus.Pending;

  @Column(() => CreditCard, { prefix: '' })
  credit_card: CreditCard;

  @CreateDateColumn({ type: 'timestamp' })
  created_at: Date

  @OneToMany(() => OrderItem, item => item.order, { cascade: true })
  items: OrderItem[];

  @BeforeInsert()
  beforeInsertActions() {
    this.generateId();
    this.calculateTotal();
  }

  generateId(){
    if (this.id) {
      return;
    }
    this.id = uuidv4();
  }

  calculateTotal() {
    return (this.total = this.items.reduce((sum, item) => {
      return sum + item.quantity * item.price;
    }, 0));
  }
}
