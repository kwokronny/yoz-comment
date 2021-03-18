declare function md5(str: string): string;
declare const axios: any;
declare let returnCitySN: any;
declare interface KBCommentConfig {
  token: string;
  theme: string;
}

declare interface KBCommentItem {
  id: number;
  createdAt: string;
  articleToken: string;
  parentId: number;
  rId: number;
  nickName: string;
  mail: string;
  site: string;
  content: string;
  ip: string;
  replys: KBCommentItem[] | null;
}

declare interface KBCommentUserInfo {
  nickName: string;
  mail: string;
  site: string;
}
