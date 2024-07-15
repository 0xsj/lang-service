import { Entity, PrimaryGeneratedColumn, Column } from 'typeorm';

export interface IEntity {
  id: string;
  name: string;
  contact: string;
}

@Entity()
export class Entities implements IEntity {
  @PrimaryGeneratedColumn()
  id: string;

  @Column()
  name: string;

  @Column()
  contact: string;

  constructor(id: string, name: string, contact: string) {
    this.id = id;
    this.name = name;
    this.contact = contact;
  }
}
