import type { User } from "./user";

export interface AuthState {
  user: User | null,
  accesToken: string | null,
  refreshToken: string | null,
}
//TODO: add expiration time
