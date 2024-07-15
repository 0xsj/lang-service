export enum IDocumentType {
  W4 = 'w4',
  HANDBOOK = 'handbook',
  APPLICATION = 'application',
}

export enum IDocumentStatus {
  APPROVED = 'approved',
  REJECTED = 'rejected',
  CANCELLED = 'cancelled',
}

export interface IDocument {
  id: string;
  type: IDocumentType;
  status: IDocumentStatus;
}
