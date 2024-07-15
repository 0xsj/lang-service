import { Column, Entity, PrimaryGeneratedColumn } from 'typeorm';

export interface IDocument {
  id: string;
  type: string;
}

@Entity()
export class Document implements IDocument {
  @PrimaryGeneratedColumn()
  id: string;

  @Column()
  type: string;

  constructor(id: string, type: string) {
    this.id = id;
    this.type = type;
  }
}
