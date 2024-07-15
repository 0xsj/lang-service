export enum UserRole {
  ADMIN = 'admin',
  USER = 'user',
  EMPLOYEE = 'employee',
  GUEST = 'guest',
}

export interface IUser {
  id: string;
  firstName: string;
  lastName: string;
  role: UserRole;
}
