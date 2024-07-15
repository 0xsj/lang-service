import { Column, Entity, PrimaryGeneratedColumn } from 'typeorm';

export interface IUser {
  id: number;
  firstName: string;
  lastName: string;
  email: string;
  role: string;
}

@Entity()
export class User implements IUser {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  firstName: string;

  @Column()
  lastName: string;

  @Column()
  email: string;

  @Column()
  role: string;

  constructor(
    id: number,
    firstName: string,
    lastName: string,
    email: string,
    role: string
  ) {
    this.id = id;
    this.firstName = firstName;
    this.lastName = lastName;
    this.email = email;
    this.role = role;
  }
}
