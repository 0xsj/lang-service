import { Entity, PrimaryGeneratedColumn, Column } from 'typeorm';

export interface IBrand {
  id: number;
  name: string;
}

@Entity()
export class Brand implements IBrand {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  name: string;
  constructor(id: number, name: string) {
    this.id = id;
    this.name = name;
  }
}
